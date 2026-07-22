package main

import "testing"

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
