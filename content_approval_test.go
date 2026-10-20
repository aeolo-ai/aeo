package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"os/exec"
	"path/filepath"
	"reflect"
	"strings"
	"sync/atomic"
	"testing"
)

const approvalTestFingerprint = "0123456789abcdef0123456789abcdef"

func TestContentApprovalCLI(t *testing.T) {
	// Build the shipped entry point: a test-only main helper could miss dispatch
	// or argument handling differences in the actual executable.
	binary := filepath.Join(t.TempDir(), "aeo")
	if out, err := exec.Command("go", "build", "-o", binary, ".").CombinedOutput(); err != nil {
		t.Fatalf("build CLI: %v\n%s", err, out)
	}
	home := t.TempDir()
	configDir := filepath.Join(home, ".config", "aeo")
	if err := os.MkdirAll(configDir, 0o700); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(configDir, "config.json"), []byte(`{"domainId":"config-domain"}`), 0o600); err != nil {
		t.Fatal(err)
	}
	runCLI := func(serverURL, envDomain string, argv []string) (string, error) {
		t.Helper()
		cmd := exec.Command(binary, argv...)
		for _, entry := range os.Environ() {
			key, _, _ := strings.Cut(entry, "=")
			if key != "HOME" && key != "AEO_NO_UPDATE_CHECK" && !strings.HasPrefix(key, "AEOLO_") {
				cmd.Env = append(cmd.Env, entry)
			}
		}
		cmd.Env = append(cmd.Env, "HOME="+home, "AEOLO_API_BASE="+serverURL,
			"AEOLO_API_KEY=test-token", "AEOLO_DOMAIN_ID="+envDomain, "AEO_NO_UPDATE_CHECK=1")
		out, err := cmd.CombinedOutput()
		return string(out), err
	}

	t.Run("proxy requests", func(t *testing.T) {
		for _, action := range []string{"approval", "approve"} {
			for _, scope := range []struct {
				name      string
				flags     []string
				envDomain string
				want      string
			}{
				{"long", []string{"--domain", "explicit-domain"}, "env-domain", "explicit-domain"},
				{"long equals", []string{"--domain=explicit-domain"}, "env-domain", "explicit-domain"},
				{"short", []string{"-d", "short-domain"}, "env-domain", "short-domain"},
				{"short equals", []string{"-d=short-domain"}, "env-domain", "short-domain"},
				{"environment", nil, "env-domain", "env-domain"},
				{"saved config", nil, "", "config-domain"},
			} {
				for _, placement := range []string{"after id", "before id", "before content", "before verb", "confirm before id"} {
					if action != "approve" && placement == "confirm before id" {
						continue
					}
					name := action + "/" + scope.name + "/" + placement
					t.Run(name, func(t *testing.T) {
						flagsFirst := placement != "after id"
						argv := []string{"content", action}
						flags := append([]string{}, scope.flags...)
						if action == "approve" {
							if flagsFirst {
								flags = append(flags, "--confirm", "--fingerprint="+approvalTestFingerprint)
							} else {
								flags = append(flags, "--fingerprint", approvalTestFingerprint, "--confirm")
							}
						}
						if flagsFirst {
							argv = append(argv, flags...)
							argv = append(argv, strings.ToUpper(reviewTestID))
						} else {
							argv = append(argv, reviewTestID)
							argv = append(argv, flags...)
						}
						wantArgv := append([]string{}, argv...)
						if placement == "before content" {
							argv = append(append([]string{}, scope.flags...), "content", action)
							argv = append(argv, wantArgv[2+len(scope.flags):]...)
						} else if placement == "before verb" {
							argv = append([]string{"content"}, scope.flags...)
							argv = append(argv, action)
							argv = append(argv, wantArgv[2+len(scope.flags):]...)
						} else if placement == "confirm before id" {
							argv = []string{"content", "approve", "--confirm", strings.ToUpper(reviewTestID), "--fingerprint", approvalTestFingerprint}
							argv = append(argv, scope.flags...)
						}
						if action == "approve" {
							id := reviewTestID
							if flagsFirst {
								id = strings.ToUpper(id)
							}
							wantArgv = []string{"content", "approve", id, "--fingerprint", approvalTestFingerprint}
							if len(scope.flags) > 0 {
								wantArgv = append(wantArgv, "--domain", scope.want)
							}
							wantArgv = append(wantArgv, "--confirm")
						}
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
							if body.DomainID != scope.want || !reflect.DeepEqual(body.Argv, wantArgv) {
								t.Errorf("request changed: %#v, want %q / %s", body, wantArgv, scope.want)
							}
							w.Header().Set("Content-Type", "application/json")
							_, _ = w.Write([]byte(`{"success":true,"data":{"fingerprint":"` + approvalTestFingerprint + `","publishedByThisCommand":false}}`))
						}))
						defer server.Close()
						// Each invocation (including a caller retry) forwards the exact
						// version once, without an extra read, review or publish request.
						for attempt := int32(1); attempt <= 2; attempt++ {
							out, err := runCLI(server.URL, scope.envDomain, argv)
							if err != nil || !strings.Contains(out, approvalTestFingerprint) || !strings.Contains(out, `"publishedByThisCommand": false`) {
								t.Fatalf("command failed: %v: %s", err, out)
							}
							if got := hits.Load(); got != attempt {
								t.Fatalf("expected one request per invocation, got %d after %d attempts", got, attempt)
							}
						}
					})
				}
			}
		}
	})

	t.Run("validation before request", func(t *testing.T) {
		cases := []struct {
			name   string
			action string
			args   []string
			want   string
		}{
			{"missing fingerprint", "approve", []string{reviewTestID, "--confirm"}, "--fingerprint is required"},
			{"missing confirm", "approve", []string{reviewTestID, "--fingerprint", approvalTestFingerprint}, "--confirm is required"},
			{"missing fingerprint value", "approve", []string{reviewTestID, "--confirm", "--fingerprint"}, "--fingerprint requires a value"},
			{"empty fingerprint", "approve", []string{reviewTestID, "--fingerprint=", "--confirm"}, "--fingerprint requires a value"},
			{"blank fingerprint", "approve", []string{reviewTestID, "--fingerprint", " ", "--confirm"}, "--fingerprint requires a value"},
			{"flag as fingerprint", "approve", []string{reviewTestID, "--fingerprint", "--confirm"}, "--fingerprint requires a value"},
			{"duplicate fingerprint", "approve", []string{reviewTestID, "--fingerprint", approvalTestFingerprint, "--fingerprint=" + approvalTestFingerprint, "--confirm"}, "provide --fingerprint only once"},
			{"duplicate confirm", "approve", []string{reviewTestID, "--fingerprint", approvalTestFingerprint, "--confirm", "--confirm"}, "provide --confirm only once"},
		}
		for _, bad := range []string{"abcd", approvalTestFingerprint[:31], approvalTestFingerprint + "0", strings.ToUpper(approvalTestFingerprint), "g" + approvalTestFingerprint[1:], " " + approvalTestFingerprint, approvalTestFingerprint + "\n"} {
			cases = append(cases, struct {
				name, action string
				args         []string
				want         string
			}{"invalid fingerprint " + bad, "approve", []string{reviewTestID, "--fingerprint", bad, "--confirm"}, "--fingerprint must be 32 lowercase hex characters"})
		}
		for _, value := range []string{"false", "true", "", "1", "0"} {
			cases = append(cases, struct {
				name, action string
				args         []string
				want         string
			}{"confirm equals " + value, "approve", []string{reviewTestID, "--fingerprint", approvalTestFingerprint, "--confirm=" + value}, "--confirm must be a bare flag"})
		}
		for _, action := range []string{"approval", "approve"} {
			base := []string{}
			if action == "approve" {
				base = []string{"--fingerprint", approvalTestFingerprint, "--confirm"}
			}
			for _, invalid := range []struct {
				name string
				args []string
				want string
			}{
				{"missing id", nil, "exactly one content ID is required"},
				{"empty id", []string{""}, "content ID must be a UUID"},
				{"invalid id", []string{"article"}, "content ID must be a UUID"},
				{"nonhex id", []string{"z23e4567-e89b-42d3-a456-426614174000"}, "content ID must be a UUID"},
				{"unhyphenated id", []string{strings.ReplaceAll(reviewTestID, "-", "")}, "content ID must be a UUID"},
				{"padded id", []string{reviewTestID + " "}, "content ID must be a UUID"},
				{"extra id", []string{reviewTestID, reviewTestKey}, "exactly one content ID is required"},
				{"extra value", []string{reviewTestID, "false"}, "exactly one content ID is required"},
				{"unknown", []string{reviewTestID, "--unknown=value"}, "unsupported flag --unknown"},
				{"short flag", []string{reviewTestID, "-x"}, "unsupported flag -x"},
				{"separator", []string{reviewTestID, "--"}, "unsupported flag --"},
				{"force", []string{reviewTestID, "--force"}, "unsupported flag --force"},
				{"repair", []string{reviewTestID, "--repair=false"}, "unsupported flag --repair"},
				{"publish", []string{reviewTestID, "--publish"}, "unsupported flag --publish"},
				{"status", []string{reviewTestID, "--status=approved"}, "unsupported flag --status"},
				{"review key", []string{reviewTestID, "--key", reviewTestKey}, "unsupported flag --key"},
				{"review job", []string{reviewTestID, "--job", reviewTestJob}, "unsupported flag --job"},
				{"missing domain value", []string{reviewTestID, "--domain"}, "--domain requires a value"},
				{"missing short domain value", []string{reviewTestID, "-d"}, "--domain requires a value"},
				{"empty domain", []string{reviewTestID, "--domain="}, "--domain requires a value"},
				{"blank domain", []string{reviewTestID, "-d", " "}, "--domain requires a value"},
				{"flag as domain", []string{reviewTestID, "--domain", "--confirm"}, "--domain requires a value"},
				{"duplicate domain aliases", []string{reviewTestID, "--domain", "one", "-d=two"}, "provide --domain only once"},
				{"duplicate long domain", []string{reviewTestID, "--domain=one", "--domain", "one"}, "provide --domain only once"},
				{"duplicate short domain", []string{reviewTestID, "-d", "one", "-d=one"}, "provide --domain only once"},
			} {
				cases = append(cases, struct {
					name, action string
					args         []string
					want         string
				}{action + " " + invalid.name, action, append(append([]string{}, base...), invalid.args...), invalid.want})
			}
		}
		for _, flag := range []string{"--fingerprint", "--fingerprint=" + approvalTestFingerprint, "--confirm", "--confirm=false"} {
			name, _, _ := strings.Cut(flag, "=")
			cases = append(cases, struct {
				name, action string
				args         []string
				want         string
			}{"read rejects " + flag, "approval", []string{reviewTestID, flag}, "unsupported flag " + name})
		}
		for _, tc := range cases {
			t.Run(tc.name, func(t *testing.T) {
				var hits atomic.Int32
				server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					hits.Add(1)
					_, _ = w.Write([]byte(`{"success":true}`))
				}))
				defer server.Close()
				argv := append([]string{"content", tc.action}, tc.args...)
				out, err := runCLI(server.URL, "env-domain", argv)
				if err == nil || !strings.Contains(out, "Error: "+tc.want) {
					t.Errorf("want local error %q, got %v: %s", tc.want, err, out)
				}
				if got := hits.Load(); got != 0 {
					t.Fatalf("invalid command sent %d requests", got)
				}
			})
		}
	})

	t.Run("global domain validation before request", func(t *testing.T) {
		for _, action := range []string{"approval", "approve"} {
			for _, tc := range []struct {
				name   string
				prefix []string
				tail   []string
				want   string
			}{
				{"empty prefix domain", []string{"--domain="}, nil, "--domain requires a value"},
				{"empty short prefix domain", []string{"-d", ""}, nil, "--domain requires a value"},
				{"flag as prefix domain", []string{"--domain", "--confirm"}, nil, "--domain requires a value"},
				{"duplicate prefix aliases", []string{"--domain", "one", "-d=two"}, nil, "provide --domain only once"},
				{"duplicate prefix and tail", []string{"-d=one"}, []string{"--domain", "two"}, "provide --domain only once"},
				{"unknown with prefix", []string{"--domain=one"}, []string{"--unknown"}, "unsupported flag --unknown"},
			} {
				t.Run(action+"/"+tc.name, func(t *testing.T) {
					var hits atomic.Int32
					server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
						hits.Add(1)
						_, _ = w.Write([]byte(`{"success":true}`))
					}))
					defer server.Close()
					argv := append(append([]string{}, tc.prefix...), "content", action, reviewTestID)
					if action == "approve" {
						argv = append(argv, "--fingerprint", approvalTestFingerprint, "--confirm")
					}
					argv = append(argv, tc.tail...)
					out, err := runCLI(server.URL, "env-domain", argv)
					if err == nil || !strings.Contains(out, "Error: "+tc.want) || hits.Load() != 0 {
						t.Fatalf("want local error %q, got %v, requests=%d: %s", tc.want, err, hits.Load(), out)
					}
				})
			}
		}
	})

	t.Run("API refusals", func(t *testing.T) {
		for _, action := range []string{"approval", "approve"} {
			for _, status := range []int{http.StatusForbidden, http.StatusConflict} {
				t.Run(action+"/"+http.StatusText(status), func(t *testing.T) {
					var hits atomic.Int32
					server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
						hits.Add(1)
						w.Header().Set("Content-Type", "application/json")
						w.WriteHeader(status)
						_, _ = w.Write([]byte(`{"message":"Approval refused; read the current version again"}`))
					}))
					defer server.Close()
					argv := []string{"content", action, reviewTestID}
					if action == "approve" {
						argv = append(argv, "--fingerprint", approvalTestFingerprint, "--confirm")
					}
					out, err := runCLI(server.URL, "env-domain", argv)
					if err == nil || !strings.Contains(out, "Approval refused") || hits.Load() != 1 {
						t.Fatalf("API refusal was not preserved: %v, requests=%d, %s", err, hits.Load(), out)
					}
				})
			}
		}
	})

	t.Run("version", func(t *testing.T) {
		out, err := runCLI("http://127.0.0.1:1", "env-domain", []string{"--version"})
		if err != nil || !strings.Contains(out, "aeo "+version+" (native)") {
			t.Fatalf("CLI version mismatch: %v, %s", err, out)
		}
	})

	t.Run("help", func(t *testing.T) {
		for _, argv := range [][]string{{"--help"}, {"content", "--help"}} {
			out, err := runCLI("http://127.0.0.1:1", "env-domain", argv)
			if err != nil || !strings.Contains(out, "approval <id>") || !strings.Contains(out, "approve <id>") || !strings.Contains(out, "--fingerprint") || !strings.Contains(out, "--confirm") {
				t.Fatalf("approval help missing: %v, %s", err, out)
			}
		}
	})
}
