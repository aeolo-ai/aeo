package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"os/exec"
	"reflect"
	"strings"
	"sync/atomic"
	"testing"
)

// Exercise main in a subprocess: these regressions concern dispatch, not just
// payload helpers. The HTTP stub never contacts a real account or creates jobs.
func TestCommandContractDispatch(t *testing.T) {
	if raw := os.Getenv("AEO_CONTRACT_CHILD_ARGS"); raw != "" {
		var args []string
		if err := json.Unmarshal([]byte(raw), &args); err != nil {
			t.Fatal(err)
		}
		os.Args = append([]string{"aeo"}, args...)
		main()
		return
	}
	promptFile := t.TempDir() + "/prompt.txt"
	if err := os.WriteFile(promptFile, []byte("  article with\nmultiple lines  "), 0600); err != nil {
		t.Fatal(err)
	}
	tests := []struct {
		name        string
		args        []string
		wantArgv    []string
		wantContent string
		wantRead    string
		status      int
	}{
		{name: "generate overrides", args: []string{"content", "generate", "--prompt", "article", "--hero", "false", "--hero-source", "gallery", "--hero-sweep", "3", "--article-length", "long", "--locale-fanout", "false", "--locales", "ja,ko", "--target-channel", "channel", "--folder", "/guides"}},
		{name: "prompt file", args: []string{"content", "write", "--prompt", "ignored", "--prompt-file=" + promptFile, "--hero", "false"}, wantArgv: []string{"content", "write", "--hero", "false", "--prompt", "article with\nmultiple lines"}},
		{name: "publish target", args: []string{"channel", "update", "channel", "--publish-target", "category"}},
		{name: "bare products", args: []string{"products"}, wantRead: "/products"},
		{name: "product list", args: []string{"products", "list"}, wantRead: "/products"},
		{name: "bare channels", args: []string{"channels"}, wantRead: "/channels"},
		{name: "channel list", args: []string{"channels", "list"}, wantRead: "/channels"},
		{name: "product discovery", args: []string{"products", "discover"}},
		{name: "product refresh", args: []string{"products", "rescan"}},
		{name: "future product verb", args: []string{"products", "future-verb", "--option", "value"}},
		{name: "plural channel indexing", args: []string{"channels", "indexing", "channel", "--enabled", "true", "--backfill"}},
		{name: "plural channel update", args: []string{"channels", "update", "channel", "--publish-target", "category"}},
		{name: "plural channel delete", args: []string{"channels", "delete", "channel"}},
		{name: "legacy import body", args: []string{"content", "import", "--title", "title", "--body", "article body"}, wantContent: "article body"},
		{name: "import canonical content", args: []string{"content", "import", "--title", "title", "--content", "article body"}, wantContent: "article body"},
		{name: "server rejection", args: []string{"content", "generate", "--prompt", "article", "--hero", "false"}, status: 400},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var requests atomic.Int32
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				requests.Add(1)
				var body map[string]any
				if tt.wantRead == "" {
					if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
						t.Error(err)
					}
				}
				if tt.wantRead != "" {
					if r.Method != "GET" || r.URL.Path != "/v2/connector/domains/audit-domain"+tt.wantRead {
						t.Errorf("wrong read: %s %s", r.Method, r.URL.Path)
					}
				} else if tt.wantContent != "" {
					if r.Method != "POST" || r.URL.Path != "/v2/connector/domains/audit-domain/content/import" || body["content"] != tt.wantContent {
						t.Errorf("wrong import: %s %s %#v", r.Method, r.URL.Path, body)
					}
				} else {
					if r.Method != "POST" || r.URL.Path != "/v2/connector/execute" {
						t.Errorf("wrong route: %s %s", r.Method, r.URL.Path)
					}
					want := tt.wantArgv
					if want == nil {
						want = tt.args
					}
					raw, _ := json.Marshal(body["argv"])
					var got []string
					if err := json.Unmarshal(raw, &got); err != nil {
						t.Error(err)
					}
					if !reflect.DeepEqual(got, want) || body["domainId"] != "audit-domain" {
						t.Errorf("lost command/options/scope: %#v; want %v", body, want)
					}
				}
				w.Header().Set("Content-Type", "application/json")
				if tt.status != 0 {
					w.WriteHeader(tt.status)
					_, _ = w.Write([]byte(`{"message":"contract rejected"}`))
				} else {
					_, _ = w.Write([]byte(`{"success":true,"data":"accepted"}`))
				}
			}))
			defer server.Close()
			raw, _ := json.Marshal(tt.args)
			cmd := exec.Command(os.Args[0], "-test.run=^TestCommandContractDispatch$")
			cmd.Env = append(os.Environ(), "AEO_CONTRACT_CHILD_ARGS="+string(raw), "AEOLO_API_BASE="+server.URL, "AEOLO_API_KEY=audit-placeholder", "AEOLO_DOMAIN_ID=audit-domain")
			out, err := cmd.CombinedOutput()
			if tt.status != 0 {
				if err == nil || !strings.Contains(string(out), "contract rejected") {
					t.Fatalf("server failure lost: %v %s", err, out)
				}
			} else if err != nil {
				t.Fatalf("command failed: %v %s", err, out)
			}
			if requests.Load() != 1 {
				t.Errorf("got %d requests, want 1", requests.Load())
			}
		})
	}
}

func TestContentGeneratePromptValidation(t *testing.T) {
	for _, args := range [][]string{{"content", "generate"}, {"content", "generate", "--prompt", "  "}} {
		if _, err := contentGenerateArgs(args); err == nil {
			t.Errorf("empty prompt accepted: %v", args)
		}
	}
	if _, err := contentGenerateArgs([]string{"content", "generate", "--prompt-file", t.TempDir() + "/missing"}); err == nil {
		t.Fatal("unreadable prompt file accepted")
	}
	file := t.TempDir() + "/prompt with spaces.txt"
	if err := os.WriteFile(file, []byte("  local prompt  "), 0600); err != nil {
		t.Fatal(err)
	}
	got, err := contentGenerateArgs([]string{"content", "generate", "--prompt-file", file, "--locales", "ko,ja"})
	want := []string{"content", "generate", "--locales", "ko,ja", "--prompt", "local prompt"}
	if err != nil || !reflect.DeepEqual(got, want) {
		t.Fatalf("prompt file lost: %v %v", got, err)
	}
}
