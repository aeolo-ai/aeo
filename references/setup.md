# Setup, Auth & Onboarding Checklist

Everything needed before running GEO commands: install the CLI, authenticate, and
verify the domain's integrations are complete.

## Setup check — run before any command

Check that the `aeo` CLI is installed:

```bash
aeo --version
```

If `aeo` is not found, guide installation first (single binary — no Go or Node.js required):

```
## aeo CLI Installation

Install with one command:

curl -fsSL https://skills.aeolo.io | sh

After install, verify: `aeo --version`

## Update

aeo update              # self-update to latest
```

Then verify the agent is authenticated:

```bash
aeo auth status
```

If not logged in, guide the user through authentication (see below).

## /aeo auth — Authentication

### /aeo auth login

```bash
aeo auth login
```

Opens a browser for device-flow authentication. On success, saves the API key and
default domain to `~/.config/aeo/config.json`.

```
## Aeolo Authentication

1. Run `aeo auth login` — this opens a browser for authentication
2. After login, your API key and default domain are saved automatically
3. To switch domains: `aeo domain switch` or `--domain <id>` flag
```

### /aeo auth status

```bash
aeo auth status
```

Shows current credentials (API key hint, active domain, source: config vs env).

### /aeo auth logout

```bash
aeo auth logout
```

Clears stored credentials.

## /aeo domain setup — Setup checklist

```bash
aeo domain setup
```

Returns a 5-item checklist showing which integrations are complete:

1. **Brand Context** — domain analyzed or value proposition set
2. **Publishing Channel (Shopify)** — Shopify OAuth connected with API token
3. **Analytics (GA4 + GSC)** — Google OAuth + GA4 property + GSC site selected
4. **Data Source (Drive)** — Google Drive folder connected via SA viewer invite
5. **Content Strategy** — strategy manifest created

Use before starting automation to verify all prerequisites are met. The
daily/weekly loops should not start until all 5 items are complete.
