package main

import (
	"encoding/json"
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
