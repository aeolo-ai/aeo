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

// `domain brand update` used to build its own REST body from six whitelisted
// flags, so --key-features, --markets, --family-json and the Brand
// Understanding JSON flags never reached the server. It now hands argv to the
// server router intact.
func TestDomainBrandUpdateProxiesArgvIntact(t *testing.T) {
	argv := []string{
		"domain", "brand", "update",
		"--key-features", "Agent builder,Integrations",
		"--voice-json", `{"tone":"calm","avoid":["Hype"]}`,
		"--competitors-json", `[{"name":"Lindy","domain":"lindy.ai"}]`,
		"--ceps-json", `["Automate follow-up"]`,
		"--domain", "dom-test",
	}
	if os.Getenv("AEO_BRAND_UPDATE_PROXY_CHILD") == "1" {
		os.Args = append([]string{"aeo"}, argv...)
		main()
		return
	}
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
		if body.DomainID != "dom-test" || !reflect.DeepEqual(body.Argv, argv) {
			t.Errorf("lost scope or flags: %#v", body)
		}
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(map[string]any{"success": true, "data": "# Brand Context Updated"})
	}))
	defer server.Close()
	cmd := exec.Command(os.Args[0], "-test.run=^TestDomainBrandUpdateProxiesArgvIntact$")
	cmd.Env = append(os.Environ(), "AEO_BRAND_UPDATE_PROXY_CHILD=1", "AEOLO_API_BASE="+server.URL, "AEOLO_API_KEY=test-token", "HOME="+t.TempDir())
	out, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("command failed: %v, %s", err, out)
	}
}
