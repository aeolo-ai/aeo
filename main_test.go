package main

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestBuildPromptsListPath(t *testing.T) {
	tests := []struct {
		name    string
		args    []string
		want    string
		wantErr bool
	}{
		{name: "unfiltered", args: []string{"prompts", "list"}, want: "/prompts"},
		{name: "tracked", args: []string{"prompts", "list", "--status", "tracked"}, want: "/prompts?status=tracked"},
		{name: "untracked", args: []string{"prompts", "list", "--status=untracked"}, want: "/prompts?status=untracked"},
		{name: "invalid", args: []string{"prompts", "list", "--status", "paused"}, wantErr: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := buildPromptsListPath(tt.args)
			if tt.wantErr {
				if err == nil {
					t.Fatal("expected an error")
				}
				return
			}
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if got != tt.want {
				t.Fatalf("got %q, want %q", got, tt.want)
			}
		})
	}
}

func TestBuildPromptsGenerateBody(t *testing.T) {
	body, err := buildPromptsGenerateBody([]string{
		"prompts", "generate",
		"--count", "12",
		"--languages", "en, ko,en",
		"--instruction", "Prioritize comparison prompts.",
	}, "domain-1")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	var got map[string]any
	if err := json.Unmarshal(body, &got); err != nil {
		t.Fatalf("invalid JSON: %v", err)
	}
	if got["domainId"] != "domain-1" {
		t.Fatalf("got domainId %#v", got["domainId"])
	}
	languages, ok := got["languages"].([]any)
	if !ok || len(languages) != 2 || languages[0] != "en" || languages[1] != "ko" {
		t.Fatalf("got languages %#v", got["languages"])
	}
	if got["instruction"] != "Prioritize comparison prompts. Aim for about 12 prompts." {
		t.Fatalf("got instruction %#v", got["instruction"])
	}
}

func TestBuildPromptsGenerateBodyRejectsInvalidInput(t *testing.T) {
	tests := []struct {
		name     string
		args     []string
		domainID string
	}{
		{name: "missing domain", args: []string{"prompts", "generate"}},
		{name: "zero count", args: []string{"prompts", "generate", "--count", "0"}, domainID: "domain-1"},
		{name: "invalid count", args: []string{"prompts", "generate", "--count", "many"}, domainID: "domain-1"},
		{name: "empty languages", args: []string{"prompts", "generate", "--languages", ",,"}, domainID: "domain-1"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if _, err := buildPromptsGenerateBody(tt.args, tt.domainID); err == nil {
				t.Fatal("expected an error")
			}
		})
	}
}

func TestCallAPIOnceUsesDirectScoreEndpoint(t *testing.T) {
	var gotPath string
	var gotAuthorization string
	var gotClientVersion string
	var gotBody []byte

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotPath = r.URL.Path
		gotAuthorization = r.Header.Get("Authorization")
		gotClientVersion = r.Header.Get("X-Client-Version")
		gotBody, _ = io.ReadAll(r.Body)
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"promptCount":1,"persisted":true}`))
	}))
	defer server.Close()

	t.Setenv("AEOLO_API_BASE", server.URL)
	t.Setenv("AEOLO_API_KEY", "test-token")

	body := []byte(`{"domainId":"domain-1"}`)
	if _, err := callAPIOnce("/score/prompts", http.MethodPost, body); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if gotPath != "/score/prompts" {
		t.Fatalf("got path %q, want /score/prompts", gotPath)
	}
	if gotAuthorization != "Bearer test-token" {
		t.Fatalf("got authorization %q", gotAuthorization)
	}
	if gotClientVersion != version {
		t.Fatalf("got client version %q, want %q", gotClientVersion, version)
	}
	if string(gotBody) != string(body) {
		t.Fatalf("got body %s, want %s", gotBody, body)
	}
}

func TestTopicMutationBody(t *testing.T) {
	body, err := topicMutationBody([]string{"--revision", "3", "--name", "Active Sun"}, "")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	var got map[string]any
	if err := json.Unmarshal(body, &got); err != nil {
		t.Fatalf("invalid JSON: %v", err)
	}
	if got["revision"] != float64(3) || got["name"] != "Active Sun" {
		t.Fatalf("unexpected body: %#v", got)
	}

	if _, err := topicMutationBody([]string{"--name", "Active Sun"}, ""); err == nil {
		t.Fatal("expected missing revision error")
	}
	if _, err := topicMutationBody([]string{"--revision", "3"}, ""); err == nil {
		t.Fatal("expected missing update fields error")
	}
}

func TestTopicArchiveBody(t *testing.T) {
	body, err := topicMutationBody([]string{"--revision=4"}, "archived")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	var got map[string]any
	if err := json.Unmarshal(body, &got); err != nil {
		t.Fatalf("invalid JSON: %v", err)
	}
	if got["revision"] != float64(4) || got["status"] != "archived" {
		t.Fatalf("unexpected body: %#v", got)
	}
}

func TestBuildSelfUpdatePlanUsesRunningBinaryDirectory(t *testing.T) {
	plan, err := buildSelfUpdatePlan("/Users/example/.local/bin/aeo", func(path string) (string, error) {
		return path, nil
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if plan.Method != selfUpdateDirect {
		t.Fatalf("got method %q, want %q", plan.Method, selfUpdateDirect)
	}
	if plan.InstallDir != "/Users/example/.local/bin" {
		t.Fatalf("got install dir %q", plan.InstallDir)
	}
	if plan.ExecutablePath != "/Users/example/.local/bin/aeo" {
		t.Fatalf("got executable path %q", plan.ExecutablePath)
	}
}

func TestBuildSelfUpdatePlanDetectsHomebrewSymlink(t *testing.T) {
	plan, err := buildSelfUpdatePlan("/opt/homebrew/bin/aeo", func(string) (string, error) {
		return "/opt/homebrew/Cellar/aeo/2.3.3/bin/aeo", nil
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if plan.Method != selfUpdateHomebrew {
		t.Fatalf("got method %q, want %q", plan.Method, selfUpdateHomebrew)
	}
	if plan.InstallDir != "" {
		t.Fatalf("Homebrew update should not set an install dir: %q", plan.InstallDir)
	}
}

func TestBuildSelfUpdatePlanUsesStableHomebrewOptPath(t *testing.T) {
	plan, err := buildSelfUpdatePlan("/opt/homebrew/Cellar/aeo/2.3.3/bin/aeo", func(path string) (string, error) {
		return path, nil
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if plan.ExecutablePath != "/opt/homebrew/opt/aeo/bin/aeo" {
		t.Fatalf("got executable path %q", plan.ExecutablePath)
	}
}

func TestValidateInstalledVersion(t *testing.T) {
	if err := validateInstalledVersion("aeo 2.3.4 (native)\n", "2.3.4"); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if err := validateInstalledVersion("aeo 2.3.3 (native)\n", "2.3.4"); err == nil || !strings.Contains(err.Error(), "2.3.3") {
		t.Fatalf("expected version mismatch, got %v", err)
	}
}

func TestFindAEOCopiesPreservesPathOrder(t *testing.T) {
	root := t.TempDir()
	first := filepath.Join(root, "first")
	second := filepath.Join(root, "second")
	for _, dir := range []string{first, second} {
		if err := os.Mkdir(dir, 0o755); err != nil {
			t.Fatal(err)
		}
		if err := os.WriteFile(filepath.Join(dir, "aeo"), []byte("binary"), 0o755); err != nil {
			t.Fatal(err)
		}
	}

	got := findAEOCopies(strings.Join([]string{first, second, first}, string(os.PathListSeparator)))
	want := []string{filepath.Join(first, "aeo"), filepath.Join(second, "aeo")}
	if len(got) != len(want) {
		t.Fatalf("got %v, want %v", got, want)
	}
	for i := range want {
		if got[i] != want[i] {
			t.Fatalf("got %v, want %v", got, want)
		}
	}
}
