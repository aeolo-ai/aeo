package main

import (
	"encoding/json"
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
