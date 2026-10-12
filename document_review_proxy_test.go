package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"os/exec"
	"reflect"
	"testing"
)

// Brand Notes, Content Strategy and Source Guide share one review gate on the
// server: the first save answers DOCUMENT_REVIEW_REQUIRED with a review token,
// and the same command re-run with `--force --review-token <t>` is what saves.
//
// Through v2.3.38 `strategy update` and `config data-sources update` did not
// hand argv to the shared router. Each built its own REST body from exactly one
// flag (`{"manifest": …}` / `{"data-sources": …}`) and PUT it to the legacy
// per-domain route, so --force and --review-token were dropped on the client
// and the server kept answering the same warning (pure'am, 2026-10-12). The
// legacy route does read `force`/`reviewToken` from the body — the CLI simply
// never sent them. `domain brand update` had already moved to proxyCommand,
// which is why the same flags worked there.
//
// This test runs the real `main()` in a child process against a fake API and
// requires the request to be POST /v2/connector/execute with argv byte-for-byte
// intact. It fails on the v2.3.38 source. It can only ever guard the source it
// is compiled with: a stale installed binary is invisible to it, so a passing
// run here means nothing until the fix is tagged and released.
func TestDocumentReviewFlagsReachSharedRouter(t *testing.T) {
	cases := []struct {
		name string
		argv []string
	}{
		{"strategy_force", []string{"strategy", "update", "--manifest", "Scoped strategy", "--force", "--review-token", "receipt", "--domain", "dom-test"}},
		{"strategy_force_eq_true", []string{"strategy", "update", "--manifest", "Scoped strategy", "--force=true", "--review-token=receipt", "--domain", "dom-test"}},
		{"strategy_flags_first", []string{"strategy", "update", "--force", "--review-token", "receipt", "--manifest", "Scoped strategy", "--domain", "dom-test"}},
		{"data_sources_force", []string{"config", "data-sources", "update", "--data-sources", "Private source guide", "--force", "--review-token", "receipt", "--domain", "dom-test"}},
		{"data_sources_force_eq_true", []string{"config", "data-sources", "update", "--data-sources", "Private source guide", "--force=true", "--review-token=receipt", "--domain", "dom-test"}},
	}
	if raw := os.Getenv("AEO_DOCUMENT_REVIEW_CHILD"); raw != "" {
		var argv []string
		_ = json.Unmarshal([]byte(raw), &argv)
		os.Args = append([]string{"aeo"}, argv...)
		main()
		return
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			argv := tc.argv
			hits := 0
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				hits++
				if r.Method != "POST" || r.URL.Path != "/v2/connector/execute" {
					t.Errorf("unshared route: %s %s", r.Method, r.URL.Path)
				}
				var body struct {
					DomainID string   `json:"domainId"`
					Argv     []string `json:"argv"`
				}
				_ = json.NewDecoder(r.Body).Decode(&body)
				if body.DomainID != "dom-test" || !reflect.DeepEqual(body.Argv, argv) {
					t.Errorf("lost review flags or scope: %#v", body)
				}
				for _, want := range []string{"--force", "--review-token"} {
					if !hasFlag(body.Argv, want) {
						t.Errorf("%s missing from argv the server received: %q", want, body.Argv)
					}
				}
				w.Header().Set("Content-Type", "application/json")
				_ = json.NewEncoder(w).Encode(map[string]any{"success": true, "data": "Saved revision"})
			}))
			defer server.Close()
			encoded, _ := json.Marshal(argv)
			cmd := exec.Command(os.Args[0], "-test.run=^TestDocumentReviewFlagsReachSharedRouter$")
			cmd.Env = append(os.Environ(), "AEO_DOCUMENT_REVIEW_CHILD="+string(encoded), "AEOLO_API_BASE="+server.URL, "AEOLO_API_KEY=test-token", "HOME="+t.TempDir())
			if out, err := cmd.CombinedOutput(); err != nil {
				t.Fatalf("%v %s", err, out)
			}
			if hits != 1 {
				t.Fatalf("expected exactly one request to the shared router, got %d", hits)
			}
		})
	}
}
