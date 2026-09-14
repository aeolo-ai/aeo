package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"os/exec"
	"reflect"
	"strings"
	"testing"
)

func TestTopicDemandDispatchPreservesBudgetAndScope(t *testing.T) {
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
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"status":"partial","googleVolume":null,"aiEstimate":12}`))
	}))
	defer server.Close()
	t.Setenv("AEOLO_API_BASE", server.URL)
	t.Setenv("AEOLO_API_KEY", "test-token")
	original := os.Args
	defer func() { os.Args = original }()
	for _, args := range [][]string{
		{"topics", "demand", "--domain", "saalgu"},
		{"topics", "demand", "run", "--max-credits", "1", "--domain", "saalgu"},
		{"topics", "demand", "run", "--max-credits=0", "--domain", "saalgu"},
		{"topics", "demand", "run", "--max-credits=100", "--domain", "saalgu"},
		{"topics", "demand", "poll", "job-123", "--domain", "saalgu"},
	} {
		os.Args = append([]string{"aeo"}, args...)
		main()
		if got.DomainID != "saalgu" || !reflect.DeepEqual(got.Argv, args) {
			t.Errorf("lost command or scope: %#v; want %q", got, args)
		}
	}
}

// Exercise the executable boundary: an invalid budget must fail before any
// authenticated request, instead of silently dispatching a paid job.
func TestTopicDemandRejectsUnsafeArgsBeforeRequest(t *testing.T) {
	if raw := os.Getenv("AEO_TEST_DEMAND_ARGS"); raw != "" {
		var args []string
		if err := json.Unmarshal([]byte(raw), &args); err != nil {
			panic(err)
		}
		os.Args = append([]string{"aeo", "topics", "demand"}, args...)
		main()
		return
	}
	for _, args := range [][]string{
		{"run"}, {"run", "--max-credits"}, {"run", "--max-credits="},
		{"run", "--max-credits", "-1"}, {"run", "--max-credits=1.5"},
		{"run", "--max-credits=101"}, {"run", "--max-credits=oops"},
		{"poll"}, {"poll", "--domain", "saalgu"},
	} {
		raw, _ := json.Marshal(args)
		cmd := exec.Command(os.Args[0], "-test.run=^TestTopicDemandRejectsUnsafeArgsBeforeRequest$")
		cmd.Env = append(os.Environ(), "AEO_TEST_DEMAND_ARGS="+string(raw), "AEOLO_API_BASE=http://127.0.0.1:1", "AEOLO_API_KEY=test-token")
		out, err := cmd.CombinedOutput()
		if err == nil {
			t.Fatalf("accepted unsafe args %q", args)
		}
		want := "--max-credits"
		if args[0] == "poll" {
			want = "jobId"
		}
		if !strings.Contains(string(out), want) || strings.Contains(string(out), "connection refused") {
			t.Fatalf("did not reject locally for %q: %s", args, out)
		}
	}
}
