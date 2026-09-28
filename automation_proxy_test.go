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

func TestAutomationProxyExitAndScope(t *testing.T) {
	if os.Getenv("AEO_AUTOMATION_PROXY_CHILD") == "1" {
		os.Args = []string{"aeo", "automation", "schedules", "--format", "json", "--domain", "skin-test"}
		main()
		return
	}
	for _, status := range []int{200, 400, 401} {
		t.Run(http.StatusText(status), func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				if r.Method != "POST" || r.URL.Path != "/v2/connector/execute" {
					t.Errorf("incorrect route: %s %s", r.Method, r.URL.Path)
				}
				var body struct {
					DomainID string   `json:"domainId"`
					Argv     []string `json:"argv"`
				}
				if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
					t.Error(err)
				}
				if body.DomainID != "skin-test" || !reflect.DeepEqual(body.Argv, []string{"automation", "schedules", "--format", "json", "--domain", "skin-test"}) {
					t.Errorf("lost scope or flags: %#v", body)
				}
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(status)
				if status == 200 {
					_ = json.NewEncoder(w).Encode(map[string]any{"success": true, "data": "{\n  \"domainId\": \"skin-test\",\n  \"repeat\": 3\n}"})
				} else {
					_, _ = w.Write([]byte(`{"message":"schedule access rejected"}`))
				}
			}))
			defer server.Close()
			cmd := exec.Command(os.Args[0], "-test.run=^TestAutomationProxyExitAndScope$")
			cmd.Env = append(os.Environ(), "AEO_AUTOMATION_PROXY_CHILD=1", "AEOLO_API_BASE="+server.URL, "AEOLO_API_KEY=test-token")
			out, err := cmd.CombinedOutput()
			if status == 200 {
				if err != nil || !strings.Contains(string(out), `"repeat": 3`) {
					t.Fatalf("successful command failed: %v, %s", err, out)
				}
			} else {
				if err == nil || !strings.Contains(string(out), "schedule access rejected") {
					t.Fatalf("failure incorrectly succeeded: %v, %s", err, out)
				}
			}
		})
	}
}
