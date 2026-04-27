package main

import (
	"bytes"
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

var version = "0.9.0"

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
	} else if path == "/whoami" {
		url = creds.APIBase + "/v1/connector/whoami"
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
	req.Header.Set("X-Client-Version", version)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	// Check for version update hint from server (only show if latest > current)
	if latest := resp.Header.Get("X-Latest-Version"); latest != "" && semverGreater(latest, version) {
		fmt.Fprintf(os.Stderr, "\n  Update available: %s → %s\n  Run: brew update && brew upgrade aeo\n\n", version, latest)
	}

	if resp.StatusCode >= 400 {
		// Parse generic error envelope (message + optional code/details/upgrade_url)
		var errObj struct {
			Message    string          `json:"message"`
			Code       string          `json:"code"`
			UpgradeURL string          `json:"upgrade_url"`
			Details    json.RawMessage `json:"details"`
		}
		if json.Unmarshal(respBody, &errObj) == nil && errObj.Message != "" {
			msg := errObj.Message

			// G22: include zod validation details when present
			if len(errObj.Details) > 0 && string(errObj.Details) != "null" {
				var pretty bytes.Buffer
				if json.Indent(&pretty, errObj.Details, "", "  ") == nil {
					msg = fmt.Sprintf("Message: %s\nDetails: %s", errObj.Message, pretty.String())
				} else {
					msg = fmt.Sprintf("Message: %s\nDetails: %s", errObj.Message, string(errObj.Details))
				}
			}

			// G23: append upgrade URL on TRIAL_EXPIRED
			if errObj.Code == "TRIAL_EXPIRED" && errObj.UpgradeURL != "" {
				msg = fmt.Sprintf("%s\n\n  → Subscribe at: %s", msg, errObj.UpgradeURL)
			}

			return "", fmt.Errorf("%s", msg)
		}
		return "", fmt.Errorf("HTTP %d: %s", resp.StatusCode, truncate(string(respBody), 200))
	}

	// G21: surface markdown payloads from `{success: true, data: "<markdown>"}` envelopes
	if strings.Contains(resp.Header.Get("Content-Type"), "application/json") {
		var envelope struct {
			Success *bool           `json:"success"`
			Data    json.RawMessage `json:"data"`
		}
		if json.Unmarshal(respBody, &envelope) == nil && envelope.Success != nil && *envelope.Success && len(envelope.Data) > 0 {
			var dataStr string
			if json.Unmarshal(envelope.Data, &dataStr) == nil {
				if strings.HasPrefix(dataStr, "#") || strings.HasPrefix(dataStr, "**") || strings.Contains(dataStr, "\n") {
					return dataStr, nil
				}
			}
		}
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

func wantsHelp(args []string) bool {
	for _, a := range args {
		if a == "--help" || a == "-h" {
			return true
		}
	}
	return false
}

// checkLatestVersion queries GitHub for the latest release and prints an
// upgrade notice if the current binary is outdated. Fails silently.
func checkLatestVersion() {
	client := &http.Client{Timeout: 2 * time.Second}
	resp, err := client.Get("https://api.github.com/repos/kithlabs/aeo/releases/latest")
	if err != nil {
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return
	}
	var release struct {
		TagName string `json:"tag_name"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&release); err != nil {
		return
	}
	latest := strings.TrimPrefix(release.TagName, "v")
	if latest != "" && latest != version {
		fmt.Fprintf(os.Stderr, "\nUpdate available: %s → %s\n", version, latest)
		fmt.Fprintf(os.Stderr, "  brew upgrade aeo\n")
		fmt.Fprintf(os.Stderr, "  curl -fsSL https://skills.tryaeolo.com/aeo | sh\n")
	}
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

// ── Semver ──────────────────────────────────────────────────────────────────

// semverGreater returns true if a > b (both like "0.4.1" or "v0.4.1").
func semverGreater(a, b string) bool {
	parse := func(s string) [3]int {
		s = strings.TrimPrefix(s, "v")
		parts := strings.SplitN(s, ".", 3)
		var v [3]int
		for i, p := range parts {
			if i < 3 {
				fmt.Sscanf(p, "%d", &v[i])
			}
		}
		return v
	}
	va, vb := parse(a), parse(b)
	for i := 0; i < 3; i++ {
		if va[i] > vb[i] {
			return true
		}
		if va[i] < vb[i] {
			return false
		}
	}
	return false
}

// ── Main ────────────────────────────────────────────────────────────────────

const usage = `aeo — manage your brand visibility from the terminal

USAGE:
  aeo <command> <verb> [options]

COMMANDS:
  domain        list | switch <id> | brand | brand update | audit | channels
  channel       list | add | update <id> | delete <id> | connect <id> | disconnect <id>
  visibility    show | check run | check poll <jobId>
  strategy      show | update
  content       list | get <id> | update <id> | preview <id> | deploy <id> | redeploy <id>
  prompts       list | add | update <id> | delete <id>
  metrics       overview | article <id> | traffic [--days]
  post          list | get <id> | import | approve <id> | publish <id>
  drive         list [--folder <id>] | read <file_id>
  auth          login | status | logout
  whoami        Show current user (email, tier, trial days)
  report        --command <cmd>

OPTIONS:
  -d, --domain <id>        Override domain ID
  --version                Show version
  --help                   Show this help

Run 'aeo <command>' without a verb for detailed help.
`

var subUsage = map[string]string{
	"domain": `aeo domain <verb>

  setup             Show setup checklist (integrations status)
  list              List accessible domains
  switch <id>       Switch active domain
  brand             Show brand profile
  brand update      Update brand profile
                    Flags: --name, --industry, --category, --value-proposition, --brand-context
  audit             Show latest audit report
  channels          List connected channels
`,
	"channel": `aeo channel <verb>

  list              List connected channels
  add               Add a channel (--url required, --type, --label)
  update <id>       Update a channel (--url, --type, --label)
  delete <id>       Delete a non-primary channel
  connect <id>      Generate OAuth URL to connect a social channel
  disconnect <id>   Disconnect OAuth integration from a channel

Types: shopify, vercel, linkedin, threads, reddit, instagram, x, website
`,
	"visibility": `aeo visibility <verb>

  show              Show last visibility snapshot
  check run         Trigger a new visibility check
                    Flags: --engines (comma-separated, default: chatgpt,gemini,perplexity,grok)
  check poll <id>   Poll check status
`,
	"strategy": `aeo strategy <verb>

  show              Show content strategy
  update            Update content strategy
                    Flags: --manifest, --frequency, --articles-per-cycle, --preferred-days, --auto-propose
`,
	"content": `aeo content <verb>

  list              List content items
                    Flags: --status, --limit, --offset
  get <id>          Get full article content
  update <id>       Update content item
                    Flags: --status, --deploy-status, --title, --meta-description,
                           --keywords (comma-separated), --body, --body-file, --patch ("search>>>replace")
  preview <id>      Generate preview link
  deploy <id>       Deploy to Shopify (--channel)
  redeploy <id>     Redeploy to Shopify
  import            Import a draft article
                    Required: --title, --body (or --body-file)
                    Optional: --type, --keywords (comma-separated), --language, --rationale,
                              --meta-description, --sources (JSON array)
`,
	"prompts": `aeo prompts <verb>

  list              List prompts grouped by stage
  add               Add a prompt (--prompt, --stage, --language)
  update <id>       Update a prompt (--prompt, --stage, --query-form)
  delete <id>       Delete a prompt
`,
	"metrics": `aeo metrics <verb>

  overview          Article performance overview
  article <id>      Detailed article stats
  traffic           Site-level GSC traffic (--days=7|14|30|90)
`,
	"drive": `aeo drive <verb>

  list              List Google Drive files (--folder <id>)
  read <file_id>    Read a file from Google Drive
`,
	"post": `aeo post <verb>

  list              List channel posts
                    Flags: --platform, --status, --limit, --offset
  get <id>          Get a channel post by ID
  import            Import a channel post draft
                    Required: --platform, --body (or --posts JSON array)
                    Optional: --title, --post-type, --target, --content-id, --channel-id
  preview <id>      Generate preview link
  delete <id>       Delete a channel post
  examples          List voice examples (--platform)
  examples add      Add a voice example
                    Required: --platform, --type (good|bad), --body
                    Optional: --source-url, --note
  examples delete <id>  Delete a voice example
  approve <id>      Approve a draft post for publishing
  publish <id>      Publish an approved post to its platform
`,
	"auth": `aeo auth <verb>

  login             Authenticate via browser (--api-base)
  status            Show credentials
  logout            Clear credentials
`,
	"config": `aeo config <subcommand>

  data-sources         Show configured data sources
  data-sources update  Update data sources (--data-sources)
`,
	"report": `aeo report

  Report command execution to the server.
  Flags: --command (required), --status-code, --response-body, --context
`,
}

func printSubUsage(cmd string) {
	if u, ok := subUsage[cmd]; ok {
		fmt.Print(u)
	} else {
		fmt.Print(usage)
	}
}

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
		checkLatestVersion()
	case "--help", "help":
		fmt.Print(usage)

	// ── domain ──
	case "domain", "domains":
		if len(args) < 2 || cmd == "domains" || wantsHelp(args) {
			printSubUsage("domain")
			return
		}
		switch args[1] {
		case "setup":
			run("/setup-status", "GET", nil, domainID)
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
			printSubUsage("domain")
		}

	// ── channel ──
	case "channel":
		if len(args) < 2 {
			run("/channels", "GET", nil, domainID)
			return
		}
		if wantsHelp(args) {
			printSubUsage("channel")
			return
		}
		switch args[1] {
		case "list":
			run("/channels", "GET", nil, domainID)
		case "add":
			url := findFlag(args, "--url")
			if url == "" {
				fmt.Fprintln(os.Stderr, "Error: --url is required.\nUsage: aeo channel add --url https://... [--type threads] [--label \"My Channel\"]")
				os.Exit(1)
			}
			run("/channels", "POST", buildJSON(map[string]string{
				"url":   url,
				"type":  findFlag(args, "--type"),
				"label": findFlag(args, "--label"),
			}), domainID)
		case "update":
			requireArg(args, 2, "aeo channel update <id> [--url ...] [--type ...] [--label ...]")
			run("/channels/"+args[2], "PATCH", buildJSON(map[string]string{
				"url":   findFlag(args, "--url"),
				"type":  findFlag(args, "--type"),
				"label": findFlag(args, "--label"),
			}), domainID)
		case "delete":
			requireArg(args, 2, "aeo channel delete <id>")
			run("/channels/"+args[2], "DELETE", nil, domainID)
		case "disconnect":
			requireArg(args, 2, "aeo channel disconnect <id>")
			run("/channels/"+args[2]+"/disconnect", "POST", nil, domainID)
		case "connect":
			requireArg(args, 2, "aeo channel connect <id>")
			connectChannel(args[2], domainID)
		default:
			printSubUsage("channel")
		}


	// ── visibility ──
	case "visibility":
		if len(args) < 2 || wantsHelp(args) {
			printSubUsage("visibility")
			return
		}
		switch args[1] {
		case "show":
			run("/visibility", "GET", nil, domainID)
		case "check":
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
		default:
			printSubUsage("visibility")
		}

	// ── config ──
	case "config":
		if len(args) < 2 || wantsHelp(args) {
			printSubUsage("config")
			return
		}
		switch args[1] {
		case "data-sources":
			if len(args) >= 3 && args[2] == "update" {
				ds := findFlag(args, "--data-sources")
				if ds == "" {
					fmt.Fprintln(os.Stderr, "Error: --data-sources is required")
					os.Exit(1)
				}
				data, _ := json.Marshal(map[string]any{"data-sources": ds})
				run("/data-sources", "PUT", data, domainID)
			} else {
				run("/data-sources", "GET", nil, domainID)
			}
		default:
			fmt.Fprintln(os.Stderr, "Unknown config subcommand:", args[1])
			os.Exit(1)
		}

	// ── strategy ──
	case "strategy":
		if len(args) < 2 || wantsHelp(args) {
			printSubUsage("strategy")
			return
		}
		switch args[1] {
		case "show":
			run("/strategy", "GET", nil, domainID)
		case "update":
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
		default:
			printSubUsage("strategy")
		}

	// ── content ──
	case "content":
		contentList := func() {
			status := findFlag(args, "--status")
			limit := findFlag(args, "--limit")
			offset := findFlag(args, "--offset")
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
			if offset != "" {
				if qs != "" {
					qs += "&"
				}
				qs += "offset=" + offset
			}
			if qs != "" {
				path += "?" + qs
			}
			run(path, "GET", nil, domainID)
		}
		if len(args) < 2 || wantsHelp(args) {
			printSubUsage("content")
			return
		}
		if strings.HasPrefix(args[1], "--") {
			// Treat bare flags as implicit list: aeo content --limit 5
			contentList()
			return
		}
		sub := args[1]
		switch sub {
		case "list":
			contentList()
		case "get":
			requireArg(args, 2, "aeo content get <id>")
			run("/content/"+args[2], "GET", nil, domainID)
		case "update":
			requireArg(args, 2, "aeo content update <id>")
			body := map[string]any{}
			if v := findFlag(args, "--status"); v != "" {
				body["status"] = v
			}
			if v := findFlag(args, "--deploy-status"); v != "" {
				body["deploy_status"] = v
			}
			if v := findFlag(args, "--title"); v != "" {
				body["title"] = v
			}
			if v := findFlag(args, "--meta-description"); v != "" {
				body["meta_description"] = v
			}
			if v := findFlag(args, "--keywords"); v != "" {
				var kw []string
				for _, k := range strings.Split(v, ",") {
					if t := strings.TrimSpace(k); t != "" {
						kw = append(kw, t)
					}
				}
				body["target_keywords"] = kw
			}
			if v := findFlag(args, "--body-file"); v != "" {
				raw, err := os.ReadFile(v)
				if err != nil {
					fmt.Fprintf(os.Stderr, "Error reading body file: %s\n", err)
					os.Exit(1)
				}
				body["content"] = string(raw)
			}
			if v := findFlag(args, "--body"); v != "" {
				body["content"] = v
			}
			if v := findFlag(args, "--patch"); v != "" {
				// format: "search>>>replace"
				parts := strings.SplitN(v, ">>>", 2)
				if len(parts) != 2 {
					fmt.Fprintf(os.Stderr, "Error: --patch format must be \"search>>>replace\"\n")
					os.Exit(1)
				}
				body["patches"] = []map[string]string{{"search": parts[0], "replace": parts[1]}}
			}
			data, _ := json.Marshal(body)
			run("/content/"+args[2], "PATCH", data, domainID)
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
		case "import":
			title := findFlag(args, "--title")
			if title == "" {
				fmt.Fprintf(os.Stderr, "Error: --title is required.\nUsage: aeo content import --title \"...\" --body \"...\" [--type blog] [--keywords \"k1,k2\"]\n")
				os.Exit(1)
			}
			bodyContent := findFlag(args, "--body")
			if v := findFlag(args, "--body-file"); v != "" {
				raw, err := os.ReadFile(v)
				if err != nil {
					fmt.Fprintf(os.Stderr, "Error reading body file: %s\n", err)
					os.Exit(1)
				}
				bodyContent = string(raw)
			}
			if bodyContent == "" {
				fmt.Fprintf(os.Stderr, "Error: --body or --body-file is required.\n")
				os.Exit(1)
			}
			importBody := map[string]any{
				"title":   title,
				"content": bodyContent,
			}
			if v := findFlag(args, "--type"); v != "" {
				importBody["articleType"] = v
			}
			if v := findFlag(args, "--keywords"); v != "" {
				var kw []string
				for _, k := range strings.Split(v, ",") {
					if t := strings.TrimSpace(k); t != "" {
						kw = append(kw, t)
					}
				}
				importBody["targetKeywords"] = kw
			}
			if v := findFlag(args, "--language"); v != "" {
				importBody["language"] = v
			}
			if v := findFlag(args, "--rationale"); v != "" {
				importBody["rationale"] = v
			}
			if v := findFlag(args, "--meta-description"); v != "" {
				importBody["metaDescription"] = v
			}
			if v := findFlag(args, "--sources"); v != "" {
				var sources []map[string]any
				if err := json.Unmarshal([]byte(v), &sources); err != nil {
					fmt.Fprintf(os.Stderr, "Error: --sources must be valid JSON array: %s\n", err)
					os.Exit(1)
				}
				importBody["sources"] = sources
			}
			importJSON, _ := json.Marshal(importBody)
			run("/content/import", "POST", importJSON, domainID)
		default:
			// Might be a content ID: aeo content <uuid>
			run("/content/"+sub, "GET", nil, domainID)
		}

	// ── prompts ──
	case "prompts":
		if len(args) < 2 || wantsHelp(args) {
			printSubUsage("prompts")
			return
		}
		switch args[1] {
		case "list":
			run("/prompts", "GET", nil, domainID)
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
			printSubUsage("prompts")
		}

	// ── metrics ──
	case "metrics":
		if len(args) < 2 || wantsHelp(args) {
			printSubUsage("metrics")
			return
		}
		switch args[1] {
		case "overview":
			run("/metrics/overview", "GET", nil, domainID)
		case "article":
			requireArg(args, 2, "aeo metrics article <id>")
			run("/metrics/article/"+args[2], "GET", nil, domainID)
		case "traffic":
			days := findFlag(args[2:], "--days")
			path := "/metrics/traffic"
			if days != "" {
				path += "?days=" + days
			}
			run(path, "GET", nil, domainID)
		default:
			printSubUsage("metrics")
		}

	// ── whoami ──
	case "whoami":
		run("/whoami", "GET", nil, "")

	// ── auth ──
	case "auth":
		if len(args) < 2 || wantsHelp(args) {
			printSubUsage("auth")
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

			// Best-effort: enrich with email / tier / trial from /whoami.
			// Failure is non-fatal (e.g. offline, server down, expired token).
			if whoamiRaw, err := callConnector("/whoami", "GET", nil, ""); err == nil && whoamiRaw != "" {
				var whoami struct {
					Email             string `json:"email"`
					Tier              string `json:"tier"`
					TrialDaysRemaining *int   `json:"trial_days_remaining"`
					Data              struct {
						Email             string `json:"email"`
						Tier              string `json:"tier"`
						TrialDaysRemaining *int   `json:"trial_days_remaining"`
					} `json:"data"`
				}
				if json.Unmarshal([]byte(whoamiRaw), &whoami) == nil {
					email := strOr(whoami.Email, whoami.Data.Email)
					tier := strOr(whoami.Tier, whoami.Data.Tier)
					trial := whoami.TrialDaysRemaining
					if trial == nil {
						trial = whoami.Data.TrialDaysRemaining
					}
					if email != "" {
						fmt.Printf("  Email:    %s\n", email)
					}
					if tier != "" {
						fmt.Printf("  Tier:     %s\n", tier)
					}
					if trial != nil {
						fmt.Printf("  Trial:    %d day(s) remaining\n", *trial)
					}
				}
			}

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
		if wantsHelp(args) {
			printSubUsage("report")
			return
		}
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

	// ── post (channel posts) ──
	case "post":
		postList := func() {
			platform := findFlag(args, "--platform")
			status := findFlag(args, "--status")
			limit := findFlag(args, "--limit")
			offset := findFlag(args, "--offset")
			path := "/channel-posts"
			qs := ""
			if platform != "" {
				qs += "platform=" + platform
			}
			if status != "" {
				if qs != "" {
					qs += "&"
				}
				qs += "status=" + status
			}
			if limit != "" {
				if qs != "" {
					qs += "&"
				}
				qs += "limit=" + limit
			}
			if offset != "" {
				if qs != "" {
					qs += "&"
				}
				qs += "offset=" + offset
			}
			if qs != "" {
				path += "?" + qs
			}
			run(path, "GET", nil, domainID)
		}
		if len(args) < 2 || wantsHelp(args) {
			printSubUsage("post")
			return
		}
		if strings.HasPrefix(args[1], "--") {
			postList()
			return
		}
		switch args[1] {
		case "list":
			postList()
		case "get":
			requireArg(args, 2, "aeo post get <id>")
			run("/channel-posts/"+args[2], "GET", nil, domainID)
		case "import":
			platform := findFlag(args, "--platform")
			if platform == "" {
				fmt.Fprintf(os.Stderr, "Error: --platform is required.\nUsage: aeo post import --platform threads --body \"...\"\n       aeo post import --platform threads --posts '[{\"body\":\"...\"}]'\n")
				os.Exit(1)
			}
			postsJSON := findFlag(args, "--posts")
			body := findFlag(args, "--body")
			if postsJSON == "" && body == "" {
				fmt.Fprintf(os.Stderr, "Error: --body or --posts is required.\n")
				os.Exit(1)
			}
			importBody := map[string]any{
				"platform": platform,
			}
			if postsJSON != "" {
				var posts []map[string]any
				if err := json.Unmarshal([]byte(postsJSON), &posts); err != nil {
					fmt.Fprintf(os.Stderr, "Error: --posts must be valid JSON array: %s\n", err)
					os.Exit(1)
				}
				importBody["posts"] = posts
			} else {
				importBody["body"] = body
				if v := findFlag(args, "--title"); v != "" {
					importBody["title"] = v
				}
			}
			if v := findFlag(args, "--post-type"); v != "" {
				importBody["postType"] = v
			}
			if v := findFlag(args, "--target"); v != "" {
				importBody["target"] = v
			}
			if v := findFlag(args, "--content-id"); v != "" {
				importBody["contentHistoryId"] = v
			}
			if v := findFlag(args, "--channel-id"); v != "" {
				importBody["channelId"] = v
			}
			importJSON, _ := json.Marshal(importBody)
			run("/channel-posts", "POST", importJSON, domainID)
		case "preview":
			requireArg(args, 2, "aeo post preview <id>")
			run("/channel-posts/"+args[2]+"/preview-link", "POST", nil, domainID)
		case "delete":
			requireArg(args, 2, "aeo post delete <id>")
			run("/channel-posts/"+args[2], "DELETE", nil, domainID)
		case "examples":
			if len(args) > 2 && args[2] == "add" {
				platform := findFlag(args, "--platform")
				exType := findFlag(args, "--type")
				body := findFlag(args, "--body")
				if platform == "" || exType == "" || body == "" {
					fmt.Fprintf(os.Stderr, "Error: --platform, --type, and --body are required.\nUsage: aeo post examples add --platform threads --type good --body \"...\"\n")
					os.Exit(1)
				}
				exBody := map[string]any{
					"platform":    platform,
					"exampleType": exType,
					"body":        body,
				}
				if v := findFlag(args, "--source-url"); v != "" {
					exBody["sourceUrl"] = v
				}
				if v := findFlag(args, "--note"); v != "" {
					exBody["note"] = v
				}
				exJSON, _ := json.Marshal(exBody)
				run("/voice-examples", "POST", exJSON, domainID)
			} else if len(args) > 2 && args[2] == "delete" {
				requireArg(args, 3, "aeo post examples delete <id>")
				run("/voice-examples/"+args[3], "DELETE", nil, domainID)
			} else {
				// list
				platform := findFlag(args, "--platform")
				path := "/voice-examples"
				if platform != "" {
					path += "?platform=" + platform
				}
				run(path, "GET", nil, domainID)
			}
		case "approve":
			requireArg(args, 2, "aeo post approve <id>")
			run("/channel-posts/"+args[2]+"/approve", "POST", nil, domainID)
		case "publish":
			requireArg(args, 2, "aeo post publish <id>")
			run("/channel-posts/"+args[2]+"/publish", "POST", nil, domainID)
		default:
			printSubUsage("post")
		}

	// ── drive ──
	case "drive":
		if len(args) < 2 || wantsHelp(args) {
			printSubUsage("drive")
			return
		}
		switch args[1] {
		case "list":
			folder := findFlag(args, "--folder")
			if folder != "" {
				run("/drive?folder="+folder, "GET", nil, domainID)
			} else {
				run("/drive", "GET", nil, domainID)
			}
		case "read":
			fileID := ""
			if len(args) >= 3 {
				fileID = args[2]
			}
			fileID = strOr(fileID, findFlag(args, "--id"))
			if fileID == "" {
				fmt.Fprintln(os.Stderr, "Usage: aeo drive read <file_id>")
				os.Exit(1)
			}
			run("/drive/"+fileID, "GET", nil, domainID)
		default:
			printSubUsage("drive")
		}

	// ── update ──
	case "update":
		selfUpdate()

	// ── aliases ──
	case "brand-profile":
		run("/brand-profile", "GET", nil, domainID)
	case "audit-report":
		run("/audit-report", "GET", nil, domainID)
	case "channels":
		run("/channels", "GET", nil, domainID)
	case "channel-connect":
		if len(args) < 2 {
			fmt.Fprintln(os.Stderr, "Usage: aeo channel connect <id>")
			os.Exit(1)
		}
		connectChannel(args[1], domainID)

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

// ── Channel Connect (OAuth) ────────────────────────────────────────────────

func connectChannel(channelID, domainOverride string) {
	creds := resolveCredentials()
	if creds.APIKey == "" {
		fmt.Fprintln(os.Stderr, "Not logged in. Run: aeo auth login")
		os.Exit(1)
	}

	did := domainOverride
	if did == "" {
		did = creds.DomainID
	}
	if did == "" {
		fmt.Fprintln(os.Stderr, "No domain set. Use --domain or run: aeo domain switch <id>")
		os.Exit(1)
	}

	// Get OAuth URL via connector API
	authURL := fmt.Sprintf("%s/v1/connector/domains/%s/social/auth-url?channelId=%s&platform=auto", creds.APIBase, did, channelID)
	authResp, err := doAPIRequest(authURL, "GET", nil, creds.APIKey)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error getting OAuth URL: %s\n", err)
		os.Exit(1)
	}

	var authResult struct {
		AuthUrl string `json:"authUrl"`
	}
	json.Unmarshal(authResp, &authResult)

	if authResult.AuthUrl == "" {
		fmt.Fprintf(os.Stderr, "Error: No auth URL returned. Response: %s\n", truncate(string(authResp), 200))
		os.Exit(1)
	}

	fmt.Printf("\n  Opening browser to authorize...\n")
	fmt.Printf("  If it didn't open, visit this URL:\n\n")
	fmt.Printf("  %s\n\n", authResult.AuthUrl)
	openBrowser(authResult.AuthUrl)
	fmt.Println("  Complete the authorization in your browser.")
	fmt.Println("  You'll be redirected to the dashboard when done.")
}

func doAPIRequest(url, method string, body []byte, apiKey string) ([]byte, error) {
	var reqBody io.Reader
	if body != nil {
		reqBody = bytes.NewReader(body)
	}
	req, err := http.NewRequest(method, url, reqBody)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+apiKey)
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	respBody, _ := io.ReadAll(resp.Body)
	if resp.StatusCode >= 400 {
		return nil, fmt.Errorf("HTTP %d: %s", resp.StatusCode, truncate(string(respBody), 200))
	}
	return respBody, nil
}

// ── Self Update ────────────────────────────────────────────────────────────

func selfUpdate() {
	fmt.Printf("Current version: %s\n", version)

	// Detect if installed via Homebrew
	exe, _ := os.Executable()
	resolved, _ := filepath.EvalSymlinks(exe)
	if strings.Contains(resolved, "Cellar") || strings.Contains(resolved, "homebrew") {
		fmt.Println("Installed via Homebrew. Run:")
		fmt.Println("  brew upgrade aeo")
		return
	}

	fmt.Println("Downloading latest version...")

	cmd := exec.Command("sh", "-c", "curl -fsSL https://raw.githubusercontent.com/kithlabs/aeo/main/install.sh | sh")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "Update failed. Try:\n  brew update && brew upgrade aeo\n")
		os.Exit(1)
	}

	fmt.Println("✓ Updated successfully")
}
