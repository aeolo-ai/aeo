package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"reflect"
	"testing"
)

// Test main dispatch, not a duplicated flag builder: setup previously used a
// different endpoint and dropped options even when the server supported them.
func TestOnboardingReadsUseSharedCommandRouter(t *testing.T) {
	var got struct {
		Argv     []string `json:"argv"`
		DomainID string   `json:"domainId"`
	}
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/v2/connector/execute" || r.Method != http.MethodPost {
			t.Errorf("unexpected route %s %s", r.Method, r.URL.Path)
		}
		if err := json.NewDecoder(r.Body).Decode(&got); err != nil {
			t.Error(err)
		}
		w.Header().Set("Content-Type", "text/markdown")
		_, _ = w.Write([]byte("{}"))
	}))
	defer server.Close()
	t.Setenv("AEOLO_API_BASE", server.URL)
	t.Setenv("AEOLO_API_KEY", "test-token")
	original := os.Args
	defer func() { os.Args = original }()
	for _, args := range [][]string{
		{"domain", "setup", "--format", "json", "--domain", "clinic"},
		{"visibility", "summary", "--market", "KR", "--engine", "chatgpt", "--format", "json", "--domain", "clinic"},
		{"diagnose", "visibility", "summary", "--market", "US", "--domain", "clinic"},
	} {
		os.Args = append([]string{"aeo"}, args...)
		main()
		if got.DomainID != "clinic" {
			t.Errorf("domain lost: %q", got.DomainID)
		}
		want := args
		if args[0] == "diagnose" {
			want = append([]string{"visibility"}, args[2:]...)
		}
		if !reflect.DeepEqual(got.Argv, want) {
			t.Errorf("got argv %q, want %q", got.Argv, want)
		}
	}
}
