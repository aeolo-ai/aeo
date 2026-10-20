package main

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"os/exec"
	"reflect"
	"strings"
	"sync/atomic"
	"testing"
)

const reviewTestID = "123e4567-e89b-42d3-a456-426614174000"
const reviewTestKey = "a5c41b3e-78f2-4a61-8df9-2c131b8d0659"
const reviewTestJob = "6597d7ec-4ec5-4dd9-b3e6-23aa148f2145"

// Exercise main in a child process so validation must exit before any request.
func TestContentReviewCLIProcess(t *testing.T) {
	if raw := os.Getenv("AEO_CONTENT_REVIEW_CHILD"); raw != "" {
		var argv []string
		if err := json.Unmarshal([]byte(raw), &argv); err != nil {
			t.Fatal(err)
		}
		os.Args = append([]string{"aeo"}, argv...)
		main()
	}
}

func contentReviewCLI(t *testing.T, serverURL string, argv []string) (string, error) {
	t.Helper()
	encoded, err := json.Marshal(argv)
	if err != nil {
		t.Fatal(err)
	}
	cmd := exec.Command(os.Args[0], "-test.run=^TestContentReviewCLIProcess$")
	cmd.Env = append(os.Environ(), "AEO_CONTENT_REVIEW_CHILD="+string(encoded),
		"AEOLO_API_BASE="+serverURL, "AEOLO_API_KEY=test-token",
		"AEOLO_DOMAIN_ID=env-domain", "AEO_NO_UPDATE_CHECK=1", "HOME="+t.TempDir())
	out, err := cmd.CombinedOutput()
	return string(out), err
}

func TestContentReviewProxyRequests(t *testing.T) {
	cases := []struct {
		name   string
		argv   []string
		domain string
	}{
		{"run stable key", []string{"content", "review-run", reviewTestID, "--key", reviewTestKey, "--domain", "explicit-domain"}, "explicit-domain"},
		{"run equals flags first", []string{"content", "review-run", "--domain=explicit-domain", "--key=" + strings.ToUpper(reviewTestKey), strings.ToUpper(reviewTestID)}, "explicit-domain"},
		{"run short domain", []string{"content", "review-run", reviewTestID, "--key", reviewTestKey, "-d", "short-domain"}, "short-domain"},
		{"run env domain", []string{"content", "review-run", reviewTestID, "--key=" + reviewTestKey}, "env-domain"},
		{"result", []string{"content", "review-result", reviewTestID, "--domain", "explicit-domain"}, "explicit-domain"},
		{"result domain equals", []string{"content", "review-result", reviewTestID, "--domain=explicit-domain"}, "explicit-domain"},
		{"result short domain equals", []string{"content", "review-result", "-d=short-domain", reviewTestID}, "short-domain"},
		{"result env domain", []string{"content", "review-result", reviewTestID}, "env-domain"},
		{"result exact job", []string{"content", "review-result", reviewTestID, "--job", reviewTestJob, "--domain", "explicit-domain"}, "explicit-domain"},
		{"result exact job equals", []string{"content", "review-result", "--job=" + strings.ToUpper(reviewTestJob), reviewTestID}, "env-domain"},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			var hits atomic.Int32
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				hits.Add(1)
				if r.Method != http.MethodPost || r.URL.Path != "/v2/connector/execute" || r.URL.RawQuery != "" {
					t.Errorf("unexpected request: %s %s", r.Method, r.URL)
				}
				if r.Header.Get("Authorization") != "Bearer test-token" || r.Header.Get("X-Client-Version") != version {
					t.Error("request did not use the existing authenticated connector")
				}
				var body struct {
					Argv     []string `json:"argv"`
					DomainID string   `json:"domainId"`
				}
				decoder := json.NewDecoder(r.Body)
				decoder.DisallowUnknownFields()
				if err := decoder.Decode(&body); err != nil {
					t.Error(err)
				}
				if body.DomainID != tc.domain || !reflect.DeepEqual(body.Argv, tc.argv) {
					t.Errorf("command, stable key or domain changed: %#v, want %q / %s", body, tc.argv, tc.domain)
				}
				w.Header().Set("Content-Type", "application/json")
				_, _ = w.Write([]byte(`{"success":true,"data":"review response"}`))
			}))
			defer server.Close()
			// A caller retry must send the same key, without generating a new one,
			// following the result, repairing content, or publishing anything.
			for attempt := int32(1); attempt <= 2; attempt++ {
				out, err := contentReviewCLI(t, server.URL, tc.argv)
				if err != nil || !strings.Contains(out, "review response") {
					t.Fatalf("command failed: %v, %s", err, out)
				}
				if got := hits.Load(); got != attempt {
					t.Fatalf("expected one request per invocation, got %d after %d attempts", got, attempt)
				}
			}
		})
	}
}

func TestContentReviewValidationRejectsBeforeRequest(t *testing.T) {
	cases := []struct {
		name string
		argv []string
		want string
	}{
		{"run missing key", []string{"content", "review-run", reviewTestID}, "--key is required"},
		{"run missing key value", []string{"content", "review-run", reviewTestID, "--key"}, "--key requires a value"},
		{"run empty key", []string{"content", "review-run", reviewTestID, "--key="}, "--key requires a value"},
		{"run flag as key", []string{"content", "review-run", reviewTestID, "--key", "--domain", "x"}, "--key requires a value"},
		{"run malformed key", []string{"content", "review-run", reviewTestID, "--key", "random"}, "--key must be a UUID"},
		{"run duplicate key", []string{"content", "review-run", reviewTestID, "--key", reviewTestKey, "--key=" + reviewTestKey}, "provide --key only once"},
		{"result does not accept key", []string{"content", "review-result", reviewTestID, "--key", reviewTestKey}, "unsupported flag --key"},
		{"result does not accept equals key", []string{"content", "review-result", reviewTestID, "--key=" + reviewTestKey}, "unsupported flag --key"},
		{"run does not accept job", []string{"content", "review-run", reviewTestID, "--key", reviewTestKey, "--job", reviewTestJob}, "unsupported flag --job"},
		{"run does not accept equals job", []string{"content", "review-run", reviewTestID, "--key", reviewTestKey, "--job=" + reviewTestJob}, "unsupported flag --job"},
		{"result missing job value", []string{"content", "review-result", reviewTestID, "--job"}, "--job requires a value"},
		{"result empty job", []string{"content", "review-result", reviewTestID, "--job="}, "--job requires a value"},
		{"result blank job", []string{"content", "review-result", reviewTestID, "--job", " "}, "--job requires a value"},
		{"result flag as job", []string{"content", "review-result", reviewTestID, "--job", "--domain", "x"}, "--job requires a value"},
		{"result malformed job", []string{"content", "review-result", reviewTestID, "--job", "random"}, "--job must be a UUID"},
		{"result nonhex job", []string{"content", "review-result", reviewTestID, "--job", "z597d7ec-4ec5-4dd9-b3e6-23aa148f2145"}, "--job must be a UUID"},
		{"result unhyphenated job", []string{"content", "review-result", reviewTestID, "--job", strings.ReplaceAll(reviewTestJob, "-", "")}, "--job must be a UUID"},
		{"result padded job", []string{"content", "review-result", reviewTestID, "--job", reviewTestJob + " "}, "--job must be a UUID"},
		{"result duplicate job", []string{"content", "review-result", reviewTestID, "--job", reviewTestJob, "--job=" + reviewTestJob}, "provide --job only once"},
		{"result job without content id", []string{"content", "review-result", "--job", reviewTestJob}, "exactly one content ID is required"},
	}
	for _, verb := range []string{"review-run", "review-result"} {
		base := []string{"content", verb}
		if verb == "review-run" {
			base = append(base, "--key", reviewTestKey)
		}
		for _, invalid := range []struct {
			name string
			tail []string
			want string
		}{
			{"missing id", nil, "exactly one content ID is required"},
			{"empty id", []string{""}, "content ID must be a UUID"},
			{"malformed id", []string{"article"}, "content ID must be a UUID"},
			{"nonhex id", []string{"g23e4567-e89b-42d3-a456-426614174000"}, "content ID must be a UUID"},
			{"unhyphenated id", []string{strings.ReplaceAll(reviewTestID, "-", "")}, "content ID must be a UUID"},
			{"padded id", []string{" " + reviewTestID}, "content ID must be a UUID"},
			{"extra id", []string{reviewTestID, reviewTestID}, "exactly one content ID is required"},
			{"extra argument", []string{reviewTestID, "extra"}, "exactly one content ID is required"},
			{"repair", []string{reviewTestID, "--repair"}, "unsupported flag --repair"},
			{"repair equals", []string{reviewTestID, "--repair=false"}, "unsupported flag --repair"},
			{"publish", []string{reviewTestID, "--publish", "true"}, "unsupported flag --publish"},
			{"force", []string{reviewTestID, "--force"}, "unsupported flag --force"},
			{"unknown", []string{reviewTestID, "--unknown=value"}, "unsupported flag --unknown"},
			{"format", []string{reviewTestID, "--format", "json"}, "unsupported flag --format"},
			{"id flag", []string{"--id", reviewTestID}, "unsupported flag --id"},
			{"short flag", []string{reviewTestID, "-x"}, "unsupported flag -x"},
			{"separator", []string{reviewTestID, "--"}, "unsupported flag --"},
			{"missing domain value", []string{reviewTestID, "--domain"}, "--domain requires a value"},
			{"empty domain", []string{reviewTestID, "--domain="}, "--domain requires a value"},
			{"blank domain", []string{reviewTestID, "--domain", " "}, "--domain requires a value"},
			{"duplicate domain", []string{reviewTestID, "--domain", "one", "-d=two"}, "provide --domain only once"},
		} {
			cases = append(cases, struct {
				name string
				argv []string
				want string
			}{verb + " " + invalid.name, append(append([]string{}, base...), invalid.tail...), invalid.want})
		}
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			var hits atomic.Int32
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				hits.Add(1)
				_, _ = w.Write([]byte(`{"success":true,"data":"unexpected request"}`))
			}))
			defer server.Close()
			out, err := contentReviewCLI(t, server.URL, tc.argv)
			if err == nil || !strings.Contains(out, "Error: "+tc.want) {
				t.Errorf("want local validation error %q, got %v: %s", tc.want, err, out)
			}
			if hits.Load() != 0 {
				t.Fatalf("invalid command sent %d requests", hits.Load())
			}
		})
	}
}

func TestContentReviewLegacyReadCompatibility(t *testing.T) {
	for _, id := range []string{reviewTestID, "legacy-id"} {
		t.Run(id, func(t *testing.T) {
			var hits atomic.Int32
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				hits.Add(1)
				if r.Method != http.MethodGet || r.URL.Path != "/v2/connector/domains/legacy-domain/content/"+id+"/review" {
					t.Errorf("legacy review read changed: %s %s", r.Method, r.URL)
				}
				body, _ := io.ReadAll(r.Body)
				if len(body) != 0 {
					t.Errorf("legacy GET sent a body: %s", body)
				}
				w.Header().Set("Content-Type", "application/json")
				_, _ = w.Write([]byte(`{"success":true,"data":"legacy workspace"}`))
			}))
			defer server.Close()
			out, err := contentReviewCLI(t, server.URL, []string{"content", "review", id, "--domain", "legacy-domain"})
			if err != nil || !strings.Contains(out, "legacy workspace") || hits.Load() != 1 {
				t.Fatalf("legacy review failed: %v, requests=%d, %s", err, hits.Load(), out)
			}
		})
	}
}

func TestContentReviewKeepsAPIAccessRestriction(t *testing.T) {
	for _, verb := range []string{"review-run", "review-result"} {
		t.Run(verb, func(t *testing.T) {
			var hits atomic.Int32
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				hits.Add(1)
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusForbidden)
				_, _ = w.Write([]byte(`{"code":"FORBIDDEN","message":"Review requires staff access"}`))
			}))
			defer server.Close()
			argv := []string{"content", verb, reviewTestID}
			if verb == "review-run" {
				argv = append(argv, "--key", reviewTestKey)
			}
			out, err := contentReviewCLI(t, server.URL, argv)
			if err == nil || !strings.Contains(out, "Review requires staff access") || hits.Load() != 1 {
				t.Fatalf("API refusal was not preserved: %v, requests=%d, %s", err, hits.Load(), out)
			}
		})
	}
}
