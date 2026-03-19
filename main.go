package main

import (
	"bytes"
	"embed"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"time"
)

//go:embed references/*.md
var referencesFS embed.FS

var version = "0.3.0"

// ── Config ──────────────────────────────────────────────────────────────────

type Config struct {
	APIKey   string `json:"apiKey,omitempty"`
	DomainID string `json:"domainId,omitempty"`
	APIBase  string `json:"apiBase,omitempty"`
}

func configPath() string {
	home, _ := os.UserHomeDir()
	return filepath.Join(home, ".config", "aeo", "config.json")
}

func readConfig() Config {
	data, err := os.ReadFile(configPath())
	if err != nil {
		return Config{}
	}
	var cfg Config
	json.Unmarshal(data, &cfg)
	return cfg
}

func writeConfig(cfg Config) error {
	dir := filepath.Dir(configPath())
	os.MkdirAll(dir, 0o700)
	data, _ := json.MarshalIndent(cfg, "", "  ")
	return os.WriteFile(configPath(), append(data, '\n'), 0o600)
}

// ── Credentials ─────────────────────────────────────────────────────────────

type Credentials struct {
	APIBase  string
	APIKey   string
	DomainID string
}

func resolveCredentials() Credentials {
	cfg := readConfig()
	return Credentials{
		APIBase:  envOr("AEOLO_API_BASE", strOr(cfg.APIBase, "https://api.tryaeolo.com")),
		APIKey:   envOr("AEOLO_API_KEY", cfg.APIKey),
		DomainID: envOr("AEOLO_DOMAIN_ID", cfg.DomainID),
	}
}

func envOr(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}

func strOr(a, b string) string {
	if a != "" {
		return a
	}
	return b
}

// ── HTTP Client ─────────────────────────────────────────────────────────────

func callConnector(path, method string, body []byte, domainOverride string) (string, error) {
	creds := resolveCredentials()
	if creds.APIKey == "" {
		return "", fmt.Errorf("not authenticated. Run: aeo auth login")
	}

	did := domainOverride
	if did == "" {
		did = creds.DomainID
	}

	var url string
	if path == "/domains" {
		url = creds.APIBase + "/v1/connector/domains"
	} else {
		if did == "" {
			return "", fmt.Errorf("domain ID required. Set AEOLO_DOMAIN_ID or use --domain")
		}
		url = creds.APIBase + "/v1/connector/domains/" + did + path
	}

	var bodyReader io.Reader
	if body != nil {
		bodyReader = bytes.NewReader(body)
	}

	req, err := http.NewRequest(method, url, bodyReader)
	if err != nil {
		return "", err
	}
	req.Header.Set("Authorization", "Bearer "+creds.APIKey)
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	if resp.StatusCode >= 400 {
		var errObj struct{ Message string `json:"message"` }
		if json.Unmarshal(respBody, &errObj) == nil && errObj.Message != "" {
			return "", fmt.Errorf("%s", errObj.Message)
		}
		return "", fmt.Errorf("HTTP %d: %s", resp.StatusCode, truncate(string(respBody), 200))
	}

	// Pretty-print JSON responses
	if strings.Contains(resp.Header.Get("Content-Type"), "application/json") {
		var buf bytes.Buffer
		if json.Indent(&buf, respBody, "", "  ") == nil {
			return buf.String(), nil
		}
	}

	return string(respBody), nil
}

func truncate(s string, n int) string {
	if len(s) <= n {
		return s
	}
	return s[:n]
}

// ── CLI Helpers ─────────────────────────────────────────────────────────────

func run(path, method string, body []byte, domainID string) {
	result, err := callConnector(path, method, body, domainID)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err)
		os.Exit(1)
	}
	fmt.Println(result)
}

func runSilent(path, method string, body []byte, domainID string) {
	result, _ := callConnector(path, method, body, domainID)
	if result != "" {
		fmt.Println(result)
	}
}

func findFlag(args []string, names ...string) string {
	for i, arg := range args {
		for _, name := range names {
			if arg == name && i+1 < len(args) {
				return args[i+1]
			}
		}
	}
	return ""
}

func buildJSON(fields map[string]string) []byte {
	m := make(map[string]string)
	for k, v := range fields {
		if v != "" {
			m[k] = v
		}
	}
	data, _ := json.Marshal(m)
	return data
}

// ── Main ────────────────────────────────────────────────────────────────────

const usage = `aeo — manage your brand visibility from the terminal

USAGE:
  aeo <command> [options]

COMMANDS:
  domain list              List accessible domains
  domain brand             Show brand profile
  domain brand update      Update brand profile
  domain audit             Show latest audit report
  domain channels          List connected channels
  visibility               Show last visibility snapshot
  visibility check run     Trigger a new visibility check
  visibility check poll    Poll check status
  strategy                 Show content strategy
  strategy update          Update content strategy
  content                  List content items
  content get <id>         Get content item
  content update <id>      Update content item
  content preview <id>     Generate preview link
  content deploy <id>      Deploy to Shopify
  content redeploy <id>    Redeploy to Shopify
  content propose          Generate proposals
  prompts                  List prompts
  prompts add              Add a prompt
  prompts update <id>      Update a prompt
  prompts delete <id>      Delete a prompt
  metrics                  Article performance overview
  metrics article <id>     Detailed article stats
  auth login               Authenticate via browser
  auth status              Show credentials
  auth logout              Clear credentials
  report                   Submit error report

OPTIONS:
  -d, --domain <id>        Override domain ID
  --version                Show version
  --help                   Show this help
`

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Print(usage)
		return
	}

	domainID := findFlag(args, "-d", "--domain")
	cmd := args[0]

	switch cmd {
	case "--version", "-V":
		fmt.Printf("aeo %s (native)\n", version)
	case "--help", "help":
		fmt.Print(usage)

	// ── domain ──
	case "domain", "domains":
		if len(args) < 2 || cmd == "domains" {
			run("/domains", "GET", nil, domainID)
			return
		}
		switch args[1] {
		case "list":
			run("/domains", "GET", nil, domainID)
		case "brand":
			if len(args) >= 3 && args[2] == "update" {
				run("/brand-profile", "PATCH", buildJSON(map[string]string{
					"name":              findFlag(args, "--name"),
					"industry":          findFlag(args, "--industry"),
					"category":          findFlag(args, "--category"),
					"value_proposition": findFlag(args, "--value-proposition"),
					"brand_context":     findFlag(args, "--brand-context"),
				}), domainID)
			} else {
				run("/brand-profile", "GET", nil, domainID)
			}
		case "audit":
			run("/audit-report", "GET", nil, domainID)
		case "channels":
			run("/channels", "GET", nil, domainID)
		case "switch":
			if len(args) < 3 {
				fmt.Fprintln(os.Stderr, "Usage: aeo domain switch <id>")
				fmt.Fprintln(os.Stderr, "Run `aeo domain list` to see available domains.")
				os.Exit(1)
			}
			cfg := readConfig()
			cfg.DomainID = args[2]
			writeConfig(cfg)
			fmt.Printf("✓ Switched to domain %s\n", args[2])
		default:
			fmt.Fprintf(os.Stderr, "Unknown domain command: %s\n", args[1])
			os.Exit(1)
		}

	// ── visibility ──
	case "visibility":
		if len(args) < 2 {
			run("/visibility", "GET", nil, domainID)
			return
		}
		if args[1] == "check" {
			if len(args) < 3 {
				fmt.Fprintln(os.Stderr, "Usage: aeo visibility check <run|poll>")
				os.Exit(1)
			}
			switch args[2] {
			case "run":
				engines := findFlag(args, "--engines")
				if engines == "" {
					engines = "chatgpt,gemini,perplexity,grok"
				}
				engineParts := strings.Split(engines, ",")
				for i, p := range engineParts {
					engineParts[i] = strings.TrimSpace(p)
				}
				enginesJSON, _ := json.Marshal(map[string][]string{"engines": engineParts})
				run("/visibility-check", "POST", enginesJSON, domainID)
			case "poll":
				if len(args) < 4 {
					fmt.Fprintln(os.Stderr, "Usage: aeo visibility check poll <jobId>")
					os.Exit(1)
				}
				run("/visibility-check/"+args[3], "GET", nil, domainID)
			default:
				fmt.Fprintf(os.Stderr, "Unknown visibility check command: %s\n", args[2])
				os.Exit(1)
			}
		} else {
			run("/visibility", "GET", nil, domainID)
		}

	// ── strategy ──
	case "strategy":
		if len(args) >= 2 && args[1] == "update" {
			manifest := findFlag(args, "--manifest")
			freq := findFlag(args, "--frequency")
			if manifest == "" && freq == "" {
				fmt.Fprintln(os.Stderr, "Error: --manifest or --frequency required")
				os.Exit(1)
			}
			body := map[string]any{}
			if manifest != "" {
				body["manifest"] = manifest
			}
			if freq != "" {
				sc := map[string]any{"frequency": freq}
				if apc := findFlag(args, "--articles-per-cycle"); apc != "" {
					sc["articles_per_cycle"] = apc
				}
				if pd := findFlag(args, "--preferred-days"); pd != "" {
					sc["preferred_days"] = pd
				}
				if findFlag(args, "--auto-propose") != "" {
					sc["auto_propose"] = true
				}
				body["schedule_config"] = sc
			}
			data, _ := json.Marshal(body)
			run("/strategy", "PUT", data, domainID)
		} else {
			run("/strategy", "GET", nil, domainID)
		}

	// ── content ──
	case "content":
		if len(args) < 2 {
			status := findFlag(args, "--status")
			limit := findFlag(args, "--limit")
			path := "/content"
			qs := ""
			if status != "" {
				qs += "status=" + status
			}
			if limit != "" {
				if qs != "" {
					qs += "&"
				}
				qs += "limit=" + limit
			}
			if qs != "" {
				path += "?" + qs
			}
			run(path, "GET", nil, domainID)
			return
		}
		sub := args[1]
		switch sub {
		case "get":
			requireArg(args, 2, "aeo content get <id>")
			run("/content/"+args[2], "GET", nil, domainID)
		case "update":
			requireArg(args, 2, "aeo content update <id>")
			run("/content/"+args[2], "PATCH", buildJSON(map[string]string{
				"status":        findFlag(args, "--status"),
				"deploy_status": findFlag(args, "--deploy-status"),
				"title":         findFlag(args, "--title"),
			}), domainID)
		case "preview":
			requireArg(args, 2, "aeo content preview <id>")
			run("/content/"+args[2]+"/preview-link", "POST", nil, domainID)
		case "deploy":
			requireArg(args, 2, "aeo content deploy <id>")
			run("/content/"+args[2]+"/deploy", "POST", buildJSON(map[string]string{
				"channel_id": findFlag(args, "--channel"),
			}), domainID)
		case "redeploy":
			requireArg(args, 2, "aeo content redeploy <id>")
			run("/content/"+args[2]+"/redeploy", "PUT", nil, domainID)
		case "propose":
			lang := findFlag(args, "--language")
			if lang == "" {
				lang = "en"
			}
			run("/content-queue", "POST", buildJSON(map[string]string{"language": lang}), domainID)
		default:
			// Might be a content ID: aeo content <uuid>
			run("/content/"+sub, "GET", nil, domainID)
		}

	// ── prompts ──
	case "prompts":
		if len(args) < 2 {
			run("/prompts", "GET", nil, domainID)
			return
		}
		switch args[1] {
		case "add":
			prompt := findFlag(args, "--prompt")
			if prompt == "" {
				fmt.Fprintln(os.Stderr, "Error: --prompt required")
				os.Exit(1)
			}
			lang := findFlag(args, "--language")
			if lang == "" {
				lang = "en"
			}
			run("/prompts", "POST", buildJSON(map[string]string{
				"canonical":  prompt,
				"language":   lang,
				"stage":      findFlag(args, "--stage"),
				"query_form": findFlag(args, "--query-form"),
			}), domainID)
		case "update":
			requireArg(args, 2, "aeo prompts update <id>")
			run("/prompts/"+args[2], "PATCH", buildJSON(map[string]string{
				"canonical":  findFlag(args, "--prompt"),
				"stage":      findFlag(args, "--stage"),
				"query_form": findFlag(args, "--query-form"),
			}), domainID)
		case "delete":
			requireArg(args, 2, "aeo prompts delete <id>")
			run("/prompts/"+args[2], "DELETE", nil, domainID)
		default:
			run("/prompts", "GET", nil, domainID)
		}

	// ── metrics ──
	case "metrics":
		if len(args) >= 2 && args[1] == "article" {
			requireArg(args, 2, "aeo metrics article <id>")
			run("/metrics/article/"+args[2], "GET", nil, domainID)
		} else {
			run("/metrics/overview", "GET", nil, domainID)
		}

	// ── auth ──
	case "auth":
		if len(args) < 2 {
			fmt.Print(usage)
			return
		}
		switch args[1] {
		case "status":
			cfg := readConfig()
			envKey := os.Getenv("AEOLO_API_KEY")
			envDomain := os.Getenv("AEOLO_DOMAIN_ID")
			envBase := os.Getenv("AEOLO_API_BASE")

			activeKey := strOr(envKey, cfg.APIKey)
			activeDomain := strOr(envDomain, cfg.DomainID)
			activeBase := strOr(envBase, strOr(cfg.APIBase, "https://api.tryaeolo.com"))

			if activeKey == "" {
				fmt.Println("Not logged in. Run: aeo auth login")
				return
			}

			src := func(env, cfgVal string) string {
				if env != "" {
					return "env"
				}
				if cfgVal != "" {
					return "config"
				}
				return "—"
			}

			hint := activeKey
			if len(hint) > 12 {
				hint = hint[:12]
			}

			fmt.Println("Logged in")
			fmt.Printf("  API key:  %s...  (%s)\n", hint, src(envKey, cfg.APIKey))
			if activeDomain == "" {
				fmt.Printf("  Domain:   (not set)  (%s)\n", src(envDomain, cfg.DomainID))
			} else {
				fmt.Printf("  Domain:   %s  (%s)\n", activeDomain, src(envDomain, cfg.DomainID))
			}
			fmt.Printf("  API base: %s  (%s)\n", activeBase, src(envBase, cfg.APIBase))
			fmt.Printf("  Config:   %s\n", configPath())

			if envDomain != "" && cfg.DomainID != "" && envDomain != cfg.DomainID {
				fmt.Printf("\n  ⚠ Config domain (%s) overridden by AEOLO_DOMAIN_ID env var.\n", cfg.DomainID)
				fmt.Println("    Use --domain to override, or unset the env var.")
			}
			if envKey != "" && cfg.APIKey != "" && envKey != cfg.APIKey {
				fmt.Println("  ⚠ Config API key overridden by AEOLO_API_KEY env var.")
			}

		case "logout":
			writeConfig(Config{})
			fmt.Println("Credentials cleared.")

		case "login":
			apiBase := findFlag(args, "--api-base")
			if apiBase == "" {
				apiBase = "https://api.tryaeolo.com"
			}
			authLogin(apiBase)

		default:
			fmt.Fprintf(os.Stderr, "Unknown auth command: %s\n", args[1])
			os.Exit(1)
		}

	// ── report ──
	case "report":
		reportCmd := findFlag(args, "--command")
		if reportCmd == "" {
			fmt.Fprintln(os.Stderr, "Error: --command required")
			os.Exit(1)
		}
		reportBody := map[string]any{"command": reportCmd}
		if sc := findFlag(args, "--status-code"); sc != "" {
			var code int
			fmt.Sscanf(sc, "%d", &code)
			reportBody["statusCode"] = code
		}
		if rb := findFlag(args, "--response-body"); rb != "" {
			reportBody["responseBody"] = rb
		}
		if ctx := findFlag(args, "--context"); ctx != "" {
			reportBody["context"] = ctx
		}
		reportJSON, _ := json.Marshal(reportBody)
		runSilent("/report", "POST", reportJSON, domainID)

	// ── aliases ──
	case "brand-profile":
		run("/brand-profile", "GET", nil, domainID)
	case "audit-report":
		run("/audit-report", "GET", nil, domainID)
	case "channels":
		run("/channels", "GET", nil, domainID)

	default:
		fmt.Fprintf(os.Stderr, "Unknown command: %s\n", cmd)
		fmt.Print(usage)
		os.Exit(1)
	}
}

// ── Auth Login (device flow) ────────────────────────────────────────────────

func authLogin(apiBase string) {
	// Step 1: request device code
	resp, err := http.Post(apiBase+"/auth/device/code", "application/json", nil)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	if resp.StatusCode >= 400 {
		fmt.Fprintf(os.Stderr, "Error: HTTP %d: %s\n", resp.StatusCode, truncate(string(body), 200))
		os.Exit(1)
	}

	var deviceResp struct {
		DeviceCode      string `json:"device_code"`
		UserCode        string `json:"user_code"`
		VerificationURI string `json:"verification_uri"`
		Interval        int    `json:"interval"`
	}
	json.Unmarshal(body, &deviceResp)
	if deviceResp.Interval == 0 {
		deviceResp.Interval = 5
	}

	// Step 2: open browser
	activateURL := deviceResp.VerificationURI + "?code=" + deviceResp.UserCode
	fmt.Printf("\n  Opening browser to %s\n", activateURL)
	fmt.Println("  If it didn't open, visit the URL above.")
	fmt.Printf("\n  Code: \033[1m%s\033[0m\n\n", deviceResp.UserCode)
	openBrowser(activateURL)

	// Step 3: poll for token
	fmt.Print("  Waiting for authentication...")
	for {
		time.Sleep(time.Duration(deviceResp.Interval) * time.Second)
		fmt.Print(".")

		tokenBody, _ := json.Marshal(map[string]string{"device_code": deviceResp.DeviceCode})
		tokenResp, err := http.Post(apiBase+"/auth/device/token", "application/json", bytes.NewReader(tokenBody))
		if err != nil {
			continue
		}

		tokenRespBody, _ := io.ReadAll(tokenResp.Body)
		tokenResp.Body.Close()

		if tokenResp.StatusCode == 428 {
			// authorization_pending — keep polling
			continue
		}

		if tokenResp.StatusCode >= 400 {
			fmt.Fprintf(os.Stderr, "\nError: HTTP %d: %s\n", tokenResp.StatusCode, truncate(string(tokenRespBody), 200))
			os.Exit(1)
		}

		// Success — parse token + domains
		var result struct {
			Data struct {
				Token   string `json:"token"`
				Domains []struct {
					ID     string `json:"id"`
					Domain string `json:"domain"`
					Name   string `json:"name"`
				} `json:"domains"`
			} `json:"data"`
		}
		json.Unmarshal(tokenRespBody, &result)

		cfg := Config{
			APIKey:   result.Data.Token,
			DomainID: "",
		}
		if apiBase != "https://api.tryaeolo.com" {
			cfg.APIBase = apiBase
		}
		if len(result.Data.Domains) > 0 {
			cfg.DomainID = result.Data.Domains[0].ID
		}
		writeConfig(cfg)

		fmt.Println()
		fmt.Println("✓ Logged in")
		if len(result.Data.Token) > 12 {
			fmt.Printf("  API key: %s...\n", result.Data.Token[:12])
		}
		if len(result.Data.Domains) > 0 {
			d := result.Data.Domains[0]
			name := d.Name
			if name == "" {
				name = d.Domain
			}
			fmt.Printf("  Domain:  %s (%s)\n", name, d.ID)
		}
		fmt.Printf("  Config:  %s\n\n", configPath())
		return
	}
}

func openBrowser(url string) {
	var cmd string
	var cmdArgs []string
	switch runtime.GOOS {
	case "darwin":
		cmd = "open"
		cmdArgs = []string{url}
	case "windows":
		cmd = "cmd"
		cmdArgs = []string{"/c", "start", "", url}
	default:
		cmd = "xdg-open"
		cmdArgs = []string{url}
	}
	exec.Command(cmd, cmdArgs...).Start()
}

func requireArg(args []string, idx int, usage string) {
	if len(args) <= idx {
		fmt.Fprintf(os.Stderr, "Usage: %s\n", usage)
		os.Exit(1)
	}
}
