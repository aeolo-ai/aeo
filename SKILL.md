---
name: aeo
description: |
  Aeolo is an organic content engine: turn deep brand understanding into marketing
  content that compounds without ad spend. Two channels — articles cited in AI search
  (GEO: ChatGPT, Perplexity, Gemini) and short-form video (analyze Reels/TikToks to
  brief content). Start from the brand, then write, deploy, and measure. It loads real
  Aeolo brand data and writes back changes (brand updates, article generation, Shopify
  deployment) to run the full organic content cycle autonomously.
  Use whenever the user mentions brand understanding, organic content, content strategy,
  article writing or performance, AI-search visibility, audit scores, brand tone,
  short-form/Reel/TikTok analysis, content deployment, or GSC indexing. Triggers: /aeo,
  "understand my brand", "what should I write today", "write an article", "review my
  content", "check my AI search visibility", "analyze this Reel/TikTok", "deploy to
  Shopify", "onboard my brand", "domain setup", "weekly report", "GSC 인덱싱",
  "색인 요청", "request indexing".
---

> **Requires**: `aeo` CLI — [Install/update](https://github.com/kithlabs/aeo)
> ```
> brew install kithlabs/aeo/aeo                        # Homebrew
> curl -fsSL https://skills.tryaeolo.com | sh          # Direct install
> ```
> Run `aeo --version` to check for updates.

# Aeolo GEO Co-pilot

Read and write live Aeolo data across the full GEO execution cycle.

## Command Reference

### aeo domain — Domain profile & metadata

| Command | What it does | Reference |
|---------|-------------|-----------|
| `/aeo domain list` | List accessible domains | this file |
| `/aeo domain switch [id]` | Switch active domain | this file |
| `/aeo domain brand update` | Update brand context fields | [brand.md](references/brand.md) |
| `/aeo domain audit` | Show latest audit report | this file |
| `/aeo domain channels` | List connected channels (platform, status, ID) | this file |
| `/aeo domain setup` | Show setup checklist (integrations status) | this file |

### aeo agent — Agent context

| Command | What it does | Reference |
|---------|-------------|-----------|
| `/aeo agent context` | Show the same default brand operating context used by the dashboard agent | [brand.md](references/brand.md) |

### aeo channel — Channel management

| Command | What it does | Reference |
|---------|-------------|-----------|
| `/aeo channel list` | List connected channels | this file |
| `/aeo channel add` | Add a channel (--url required, --type, --label) | this file |
| `/aeo channel update <id>` | Update a channel (--url, --type, --label) | this file |
| `/aeo channel delete <id>` | Delete a non-primary channel | this file |
| `/aeo channel connect <id>` | OAuth connect — opens browser for threads/linkedin/reddit | this file |
| `/aeo channel disconnect <id>` | Disconnect OAuth integration from a channel | this file |

### aeo visibility — Visibility data & checks

| Command | What it does | Reference |
|---------|-------------|-----------|
| `/aeo visibility show` | Show the last visibility snapshot | [visibility.md](references/visibility.md) |
| `/aeo visibility check run` | Run a credit-metered visibility check | [visibility.md](references/visibility.md), [polling.md](references/polling.md) |
| `/aeo visibility check poll <jobId>` | Poll check status | [visibility.md](references/visibility.md), [polling.md](references/polling.md) |

### aeo audit — Site foundation checks

| Command | What it does | Reference |
|---------|-------------|-----------|
| `/aeo audit run` | Start a site foundation audit (uses production credits) | [polling.md](references/polling.md) |
| `/aeo audit poll <jobId>` | Poll an audit job | [polling.md](references/polling.md) |

### aeo strategy — Content strategy

| Command | What it does | Reference |
|---------|-------------|-----------|
| `/aeo strategy show` | Show current content strategy (manifest + schedule) | [strategy.md](references/strategy.md) |
| `/aeo strategy update` | Create or update content strategy | [strategy.md](references/strategy.md) |

### aeo content — Content lifecycle

| Command | What it does | Reference |
|---------|-------------|-----------|
| `/aeo content list` | List content items (--status, --limit, --offset) | [content-manage.md](references/content-manage.md) |
| `/aeo content get <id>` | Read full article content (markdown) | [content-manage.md](references/content-manage.md) |
| `/aeo content generate` | Start a server-side content generation job (uses production credits) | [content-create.md](references/content-create.md), [polling.md](references/polling.md) |
| `/aeo content jobs` | List active content generation jobs | [polling.md](references/polling.md) |
| `/aeo content update <id>` | Update a content item (status, title, meta, keywords, body via patches or full replace) | [content-manage.md](references/content-manage.md) |
| `/aeo content preview <id>` | Generate preview link and open in browser | [content-manage.md](references/content-manage.md) |
| `/aeo content deploy <id>` | Deploy an article to the connected Shopify channel | [content-manage.md](references/content-manage.md) |
| `/aeo content redeploy <id>` | Update an already-deployed Shopify article in-place (keeps URL) | [content-manage.md](references/content-manage.md) |
| `/aeo content import` | Push an already-written draft to content history | [content-create.md](references/content-create.md) |
| `/aeo content review <id>` | GEO content review (structure, trust, freshness, brand, engine fit) | [content-review.md](references/content-review.md) |

> Use `content generate` for the server-side paid generation job and `content import` for already-written drafts.

### aeo post — Channel posts (social media distribution)

| Command | What it does | Reference |
|---------|-------------|-----------|
| `/aeo post analyze --url <URL>` | Analyze one channel/reference URL and propose task-specific voice evidence | [tov-extract.md](references/tov-extract.md) |
| `/aeo post write` | Write a channel post (agent writes directly → review → import) | [post-create.md](references/post-create.md) |
| `/aeo post list` | List channel posts (--platform, --status, --limit, --offset) | [channel-washing.md](references/channel-washing.md) |
| `/aeo post get <id>` | Get a channel post (full body + metadata) | [channel-washing.md](references/channel-washing.md) |
| `/aeo post import` | Import a channel post draft (--platform, --body required) | [channel-washing.md](references/channel-washing.md) |

### aeo reference / video — Analysis & generation

| Command | What it does | Reference |
|---------|-------------|-----------|
| `/aeo reference analyze --url <url> --media <type>` | Analyze a reference URL as a background job (uses production credits) | [tov-extract.md](references/tov-extract.md) |
| `/aeo reference poll <jobId>` | Poll a reference analysis job | [polling.md](references/polling.md) |
| `/aeo video analyze --url <url>` | Analyze a short-form video URL synchronously (uses production credits) | this file |
| `/aeo video generate --prompt <text>` | Generate short-form video(s) for Reels/TikTok (uses production credits). `--model seedance-2-fast\|seedance-2\|kling-3\|grok-video`, `--sweep N` (1-8 candidate variations), `--aspect`, `--duration`, `--ref`, `--audio`, `--wait`. Async — returns job IDs. | this file |
| `/aeo video poll <jobId...>` | Check status + result URLs of video generation jobs | this file |

### aeo metrics — Article & site performance

| Command | What it does | Reference |
|---------|-------------|-----------|
| `/aeo metrics overview` | Show deployed articles with GA4 + GSC stats (last 30 days) | [metrics.md](references/metrics.md) |
| `/aeo metrics article <id>` | Detailed per-article stats (traffic sources, top queries) | [metrics.md](references/metrics.md) |
| `/aeo metrics traffic` | Site-level GSC traffic: top queries, pages, country, device (--days=7\|14\|30\|90) | [metrics.md](references/metrics.md) |

### aeo prompts — Tracked prompts

| Command | What it does | Reference |
|---------|-------------|-----------|
| `/aeo prompts list` | List prompts grouped by stage | [brand.md](references/brand.md) |
| `/aeo prompts add` | Add a manual prompt to brand_prompts | [brand.md](references/brand.md) |
| `/aeo prompts update <id>` | Edit an existing prompt (text, stage, query_form) | [brand.md](references/brand.md) |
| `/aeo prompts delete <id>` | Soft-delete a prompt by ID | [brand.md](references/brand.md) |

### aeo drive — Google Drive files

| Command | What it does | Reference |
|---------|-------------|-----------|
| `/aeo drive list` | List files in connected Google Drive folder (--folder) | [drive.md](references/drive.md) |
| `/aeo drive read <file_id>` | Read a file from Google Drive | [drive.md](references/drive.md) |

> **Supported types**: Google Docs/Sheets, txt/json/md/csv, **PDF**, **XLSX/XLS** (all sheets, 200-row cap each), **DOCX**, images (≤5MB base64). Not supported: `.doc`, `.pptx`, `.pages`, `.numbers`, `.key` — see [drive.md](references/drive.md).

### aeo gsc — Google Search Console (browser automation)

| Command | What it does | Reference |
|---------|-------------|-----------|
| `/aeo gsc index` | Bulk indexing request via browser automation (requires Chrome + GSC login) | [gsc-indexing.md](references/gsc-indexing.md) |
| `/aeo gsc index --domain <domain>` | Specify target domain (skips domain prompt) | [gsc-indexing.md](references/gsc-indexing.md) |

> **Environment requirement**: This command uses Claude in Chrome browser automation, NOT the aeo CLI. Requires: (1) Chrome + Claude in Chrome extension, (2) GSC login in browser. If unavailable, the agent guides the user to set up or do it manually.

### aeo config — Agent configuration

| Command | What it does | Reference |
|---------|-------------|-----------|
| `/aeo config data-sources` | Show configured data sources for research | [data-access.md](references/data-access.md) |
| `/aeo config data-sources update` | Update data source pointers (--data-sources) | [data-access.md](references/data-access.md) |

### aeo auth — Authentication

| Command | What it does | Reference |
|---------|-------------|-----------|
| `/aeo auth login` | Authenticate via browser (device flow) | this file |
| `/aeo auth status` | Show current stored credentials | this file |
| `/aeo auth logout` | Clear stored credentials | this file |

### aeo billing — Subscription and credits

| Command | What it does | Reference |
|---------|-------------|-----------|
| `/aeo billing subscription` | Show current tier and credit summary | this file |
| `/aeo billing credits` | Show current credit balance | this file |
| `/aeo billing ledger` | Show recent credit ledger entries | this file |

### Utilities

| Command | What it does | Reference |
|---------|-------------|-----------|
| `/aeo` | Load full GEO context (brand + audit + visibility) | this file |
| `/aeo report` | Submit an error report when a command fails | this file |

---

## Workflows — Autonomous GEO Optimization

These workflows enable you to run a full GEO optimization cycle — from brand setup to daily content production to weekly performance review. Read [workflows.md](references/workflows.md) for detailed decision logic and quality gates.

| Workflow | When | What |
|----------|------|------|
| **Onboarding** | New brand, first setup | Assess setup → auto-fill what you can → guide user for OAuth/permissions → verify 5/5 |
| **Daily Content** | Every day (cron or manual) | Pick topic from priority queue → write → deploy → distribute to channels |
| **Weekly Report** | Every week (cron or manual) | Visibility check → performance analysis → strategy adjustment → report to user |

Start with `aeo domain setup` to see where you are.

---

### /aeo domain setup — Setup checklist

```bash
aeo domain setup
```

Returns a 5-item checklist showing which integrations are complete:

1. **Brand Context** — domain analyzed or value proposition set
2. **Publishing Channel (Shopify)** — Shopify OAuth connected with API token
3. **Analytics (GA4 + GSC)** — Google OAuth + GA4 property + GSC site selected
4. **Data Source (Drive)** — Google Drive folder connected via SA viewer invite
5. **Content Strategy** — strategy manifest created

Use before starting automation to verify all prerequisites are met. The daily/weekly loops should not start until all 5 items are complete.

---

Read the relevant reference file before executing any command.

## CUD Rule

**Always get explicit user confirmation before any Create / Update / Delete operation.**

Applies to: visibility check run, content generate, content import, content update, content deploy, content redeploy, audit run, reference analyze, video analyze, video generate, brand update, strategy update, prompts add, prompts update, prompts delete, post import.

Never call a write API without confirmation. Always show what you're about to do and ask "Proceed?" first.

## Communication Rules

- **UUID is internal only.** User-facing messages must use `title`, `name`, `domain`, `canonical`, etc. Example: `"bc2ef290-..." updated` → `"Best Project Management Tools for Startups" updated`
- **Skill workflows** (`content review`, `post write`): These require LLM reasoning and are not the canonical server-side generation command. `aeo content generate` starts a server-side content generation job and spends production credits; use `aeo content jobs` or the relevant poll command for background status.
- **Explicit verbs required**: `aeo content list`, `aeo visibility show`, `aeo strategy show`, etc. Running `aeo <command>` without a verb shows sub-help. Exception: `aeo content --limit 5` (bare flags = implicit list).

Before writing or generating any content (`/aeo content generate` or manual draft/import), always read [geo-strategy.md](references/geo-strategy.md) and [strategy.md](references/strategy.md) first.

---

## Setup check

Before any command, check if the `aeo` CLI is installed:

```bash
aeo --version
```

If `aeo` is not found, guide installation first:

```
## aeo CLI Installation

Install with one command (no Go or Node.js required — it's a single binary):

curl -fsSL https://skills.tryaeolo.com | sh

Or with Homebrew:

brew install kithlabs/aeo/aeo

After install, verify: `aeo --version`

## Update

aeo update              # self-update to latest
brew upgrade aeo        # if installed via Homebrew
```

Then verify the agent is authenticated:

```bash
aeo auth status
```

If not logged in, guide the user through authentication:

```
## Aeolo Authentication

1. Run `aeo auth login` — this opens a browser for authentication
2. After login, your API key and default domain are saved automatically
3. To switch domains: `aeo domain switch` or `--domain <id>` flag
```

---

## /aeo — Load GEO context

Fetch agent context, audit report, and visibility data in parallel:

```bash
aeo agent context  > /tmp/aeo_brand.md &
aeo domain audit  > /tmp/aeo_audit.md &
aeo visibility show > /tmp/aeo_visibility.md &
wait
```

If any file is empty or starts with `{` (JSON error), show a helpful message and stop.

Present as a unified briefing:

```
## Aeolo GEO Briefing — {domain}

{agent-context content}

---

{audit-report content}

---

{visibility content}

---
> Data loaded from Aeolo. Ready for GEO work.
```

After presenting, note 1-2 sentences on the highest-leverage opportunity (critical audit item, visibility gap cluster, or brand mismatch). Then ask what the user wants to work on.

---

## /aeo domain list — List accessible domains

```bash
aeo domain list
```

Shows all domains the user has access to (owner + member). Useful for multi-domain setups.

---

## /aeo domain switch — Switch active domain

```bash
aeo domain switch [id]
```

Requires a domain ID. Run `aeo domain list` first to find the ID. The selected domain is persisted in `~/.config/aeo/config.json`.

---

## /aeo domain audit — Show audit report

```bash
aeo domain audit
```

Response: `text/markdown` — audit scores and recommendations. See [geo-strategy.md](references/geo-strategy.md) for how to interpret audit data.

---

## /aeo channel — Channel management

### /aeo channel list (or /aeo domain channels)

Show all channels connected to the current domain. Returns a markdown table with label, platform, URL, and channel ID. Primary channel is marked with star.

```bash
aeo channel list
aeo domain channels   # alias
```

### /aeo channel add

Add a new channel to the current domain. Type is auto-detected from URL if not specified.

```bash
aeo channel add --url https://www.threads.net/@mybrand --type threads --label "Threads Main"
```

Types: `shopify`, `vercel`, `linkedin`, `threads`, `reddit`, `instagram`, `x`, `website`

### /aeo channel update

Update an existing channel's URL, type, or label.

```bash
aeo channel update <channel-id> --label "New Label" --type linkedin
```

### /aeo channel delete

Delete a non-primary channel. Primary channels cannot be deleted.

```bash
aeo channel delete <channel-id>
```

### /aeo channel connect

Generate OAuth URL and open browser for social platform authorization (threads, linkedin, reddit).

```bash
aeo channel connect <channel-id>
```

The browser opens the platform's OAuth page. On success, redirects to the dashboard.

### /aeo channel disconnect

Remove OAuth integration from a channel without deleting the channel row.

```bash
aeo channel disconnect <channel-id>
```

---

## /aeo auth — Authentication

### /aeo auth login

```bash
aeo auth login
```

Opens a browser for device-flow authentication. On success, saves API key and default domain to `~/.config/aeo/config.json`.

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

---

## /aeo report — Error reporting

When any `aeo` command fails, automatically submit a diagnostic report so failures are tracked.

```bash
aeo report \
  --command "visibility-check run --engines=gemini,grok" \
  --status-code 500 \
  --response-body "Internal Server Error: worker timeout" \
  --context "Running scheduled visibility check after content deploy"
```

**Fields:**
- `--command` (required) — the command that was attempted
- `--status-code` — HTTP status or error code received
- `--response-body` — raw response or error message (max 2000 chars)
- `--context` — one sentence: what the agent was trying to do

The API automatically logs the domain ID and a masked token hint (first 8 chars).

> **Agent auto-report rule:** When any `aeo` CLI call fails, catch the error and call `aeo report` with the failure details before surfacing the error to the user. This is fire-and-forget — do not block on the report response.
