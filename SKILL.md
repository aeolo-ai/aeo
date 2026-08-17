---
name: aeo
description: |
  Aeolo is an organic content engine: turn deep brand understanding into marketing
  content that compounds without ad spend. One supported output channel — blog articles
  cited in AI search (GEO: ChatGPT, Perplexity, Gemini). Reels/TikToks and competitor
  posts stay in scope as reference input you analyze to brief those articles, never as
  output. Start from the brand, then write, deploy, and measure. It loads real
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

> **Requires**: `aeo` CLI — [Install/update](https://github.com/aeolo-ai/aeo)
> ```
> curl -fsSL https://skills.aeolo.io | sh
> ```
> Run `aeo --version` to check for updates.

# Aeolo GEO Co-pilot

Read and write live Aeolo data across the full GEO execution cycle.

## Command Index

Names only, on purpose: flags, credit costs, sharp edges, and output contracts live in the linked reference — **read it before executing anything** — and `aeo <noun> --help` prints the canonical flags. Commands marked † are agent-only workflows (no bare CLI verb; see Communication Rules).

> **Aliases** (dashboard-chat + MCP surfaces; the raw terminal binary ships only canonical forms): `diagnose` = `visibility`/`audit` · `measure` = `metrics` · `account` = `billing`/`whoami` · `publish` mirrors `content deploy|redeploy|preview`. Noun-first plurals (`posts`, `channels`) and legacy verbs (`brand update`) route too; prefer canonical forms in shell examples.

> **Retired surfaces (2026-07-27)** — card news (`carousel create/update`), channel posts/Threads (`post write/import/preview/approve/publish`), short-form video output (`video generate/poll`). **Never offer or call these** — route every content request to a blog article. Only output generation is retired: analysis input (`post analyze`, `reference *`, `video analyze`) and read/cleanup verbs on existing rows stay live. [retired.md](references/retired.md)

| Area | Commands | Reference |
|------|----------|-----------|
| `domain` | `list` · `add <url>` (new brands only — never re-analyzes; existing ones use `rescan`) · `switch [id]` · `brand update` (`--markets ko-KR,en-US` sets the whole reach; `--family-json` replaces the brand-family roster) · `brand aliases` (roster + alias candidates; suggests only) · `audit` · `channels` · `setup` (onboarding checklist) · `rescan` | [brand.md](references/brand.md), [setup.md](references/setup.md) |
| `agent` | `context` | [brand.md](references/brand.md) |
| `channel` | `list` · `add` · `update <id>` · `indexing <id>` · `delete <id>` · `connect <id>` · `disconnect <id>` · `voice` · `blog bind <host>` / `blog unbind` | [channels.md](references/channels.md), [tov-extract.md](references/tov-extract.md) |
| `visibility` | `show` · `history` · `check run` · `check poll <jobId>` | [visibility.md](references/visibility.md), [polling.md](references/polling.md) |
| `audit` | `run` (dynamic credits, scales with pages) · `poll <jobId>` | [polling.md](references/polling.md) |
| `strategy` | `show` · `update` · `visual` (style guide + mood boards; read it before editing — the boards are URL-keyed) · `visual update` (`--description` / `--keywords`; `--add-images` / `--remove-images` [`--product <id>`]; `--definitions` JSON merge) | [strategy.md](references/strategy.md) |
| `automation` | `schedules` · `schedule set --track visibility\|content` (partial updates safe — omitted flags inherit; the visibility track is the weekly auto-scan toggle) | [workflows.md](references/workflows.md) |
| `content` | `idea <domain_id>`† · `list` · `feed` · `get <id>` · `import` · `generate` (explicit-only paid job; needs a destination — `--target-channel`, or the domain's sole publish surface) · `jobs` · `update <id>` · `preview <id>` · `deploy <id>` · `redeploy <id>` · `unpublish <id>` · `editions <id>` (locale fan-out) · `delete <id>` (deployed copies stay live — `unpublish` first) · `job cancel <jobId>` · `review <id>` | [content-manage.md](references/content-manage.md), [content-create.md](references/content-create.md), [content-idea.md](references/content-idea.md), [content-review.md](references/content-review.md) |
| `post` | `analyze --url` (LIVE reference input) · `list` / `get <id>` / `delete <id>` (read/cleanup) | [tov-extract.md](references/tov-extract.md), [channel-washing.md](references/channel-washing.md) |
| `reference` / `video` | `reference analyze --url --media` · `reference style --url` · `reference poll <jobId>` · `reference delete <jobId>` · `video analyze --url` (15 credits, synchronous) | [tov-extract.md](references/tov-extract.md), [polling.md](references/polling.md) |
| `measure` (alias `metrics`) | `overview` · `content <id>` · `traffic` · `visibility` · `attribution` · `diag report` (command-failure diagnostics — distinct from `report snapshot`) | [metrics.md](references/metrics.md) |
| `report` | `snapshot` — freeze a rendered report into a secret/password share link (requires the rendered `--html`; the connector does not render server-side) | this file |
| `market-map` | bare `market-map` (show) · `run` · `poll <jobId>` · `populate` (rail → Topics; prompts follow via a background job) | [market-map.md](references/market-map.md) |
| `prompts` | `portfolio <domain_id>`† (the **only** safe way to restructure a live tracked set) · `audit <domain_id>`† · `list` · `add` · `update <id>` (`--regions US,GB` = standing target markets, measured once per market per check) · `delete <id>` · `prompts generate` (seeds an **empty** set only — previews nothing, saves straight to tracked, and has shrunk a live portfolio 30→14; restructuring goes through `prompts portfolio`) | [brand.md](references/brand.md), [prompt-portfolio.md](references/prompt-portfolio.md), [prompt-audit.md](references/prompt-audit.md) |
| `topics` | `next` (one article angle; saves nothing) · `suggest <domain_id>`† (agent judges an architecture) · `candidates` / `candidates poll <jobId>` (server job: demand-priced candidates + a `--demand-token`) · `prompts` / `prompts poll [jobId]` (writes tracked Prompts into empty Topics — **use this, not `prompts generate`**) · `list` · `create` (pass `--demand-token` to keep measured demand) · `update <id>` · `archive <id>` / `restore <id>` (both need `--revision`) · `assign-prompts <id>` | [topics.md](references/topics.md), [topic-suggest.md](references/topic-suggest.md), [setup.md](references/setup.md) |
| `segments` | `list` (tags are metadata/filtering only) | [brand.md](references/brand.md) |
| `products` / `image` | `products` · `product add --pdp <url>` · `products discover` · `products rescan` · `image search` · `image swap` (5 credits) · `image upload` · `image generate` (credit-metered, async) · `image poll <jobId…>` · `image gallery [delete <assetId>]` · `media set-thumbnail <id> --url` | [image-thumbnails.md](references/image-thumbnails.md) |
| `feedback` | `feedback [<message>]` (bare form opens `$EDITOR`) | this file |
| `drive` | `list` · `read <file_id>` · `download <file_id>` (supported/unsupported file types in reference) | [drive.md](references/drive.md) |
| `integrations google` | `status` (`ready` = a property/site is **selected**, not merely connected) · `properties` · `sites` · `set` — connecting Google itself is a browser sign-in in Settings, not a CLI command | [metrics.md](references/metrics.md) |
| `agency` | `request --types seo,pr,…` (waitlist; idempotent per domain+type) | this file |
| `gsc` | `index [--domain <domain>]` — Claude in Chrome browser automation, **not** the aeo CLI; needs Chrome + GSC login, else guide the user to do it manually | [gsc-indexing.md](references/gsc-indexing.md) |
| `config` | `data-sources [update]` · `glossary [update]` · `reference-policy [update]` | [data-access.md](references/data-access.md) |
| `auth` | `login` · `status` · `logout` | [setup.md](references/setup.md) |
| `account` (alias `billing`/`whoami`) | `whoami` · `subscription` · `credits` · `ledger` | this file |
| — | bare `/aeo` (full GEO context) · `/aeo report` (error reporting — see below) | [workflows.md](references/workflows.md) |

---

## Workflows — Autonomous GEO Optimization

Run the full GEO cycle — brand setup → daily content → weekly review. See [workflows.md](references/workflows.md) for decision logic and quality gates.

| Workflow | When | What |
|----------|------|------|
| **Onboarding** | New brand, first setup | Assess setup → auto-fill what you can → guide user for OAuth/permissions → verify 7/7 |
| **Daily Content** | Every day (cron or manual) | Pick topic from priority queue → write → deploy → request indexing (channel-post distribution is retired) |
| **Weekly Report** | Every week (cron or manual) | Visibility check → performance analysis → strategy adjustment → report to user |

Start with `aeo domain setup` — a 7-item checklist (family & aliases, market, topics, prompts, first check, Google, somewhere to publish) plus a "Next" section that does NOT count toward completion. Every ⬜ row carries the exact command that clears it; run the hint, then re-read the list. Install/auth and the per-step walkthrough are in [setup.md](references/setup.md); don't start the loops until all 7 are done.

---

Read the relevant reference file before executing any command.

## CUD Rule

**Always get explicit user confirmation before any Create / Update / Delete operation.**

Applies to **every write/execute/external command** — anything that is not a pure read; the registry riskLevel is the source of truth (content, prompts, topics, brand/strategy/config, channels, images, `market-map run`, `audit run`, …).

Never call a write API without confirmation. Always show what you're about to do and ask "Proceed?" first. Be extra explicit for the irreversible ones: `content deploy`/`content redeploy`/`content unpublish` change the customer's live site.

Retired writes (`carousel create/update`, `post import/preview/approve/publish`, `video generate`) are **not** on this list — do not offer or call them at all. See the retired-surfaces note in the Command Index.

## Communication Rules

- **UUID is internal only.** User-facing messages must use `title`, `name`, `domain`, `canonical`, etc. Example: `"bc2ef290-..." updated` → `"Best Project Management Tools for Startups" updated`
- **Agent-only workflows** (no bare CLI verb — they need external-agent reasoning): `/aeo topics suggest <domain_id>` proposes a durable Topic architecture without saving (distinct from `aeo topics candidates`, the server job that generates and demand-prices candidates); `/aeo content idea <domain_id>` recommends one next article; `/aeo prompts audit <domain_id>` judges Prompt validity/utility read-only; `/aeo prompts portfolio <domain_id>` safely restructures the tracked Prompt set (preview → confirm → atomic write → verify). Manual drafting writes directly then `aeo content import`. `aeo content generate` is the explicit paid server-side job; `aeo content review <id>` is a real wired command. The `post write` → `post import` workflow is retired.
- **Explicit verbs required**: `aeo content list`, `aeo visibility show`, etc. Bare `aeo <command>` shows sub-help (exception: `aeo content --limit 5` = implicit list).

Before writing/generating content, always read [geo-strategy.md](references/geo-strategy.md) and [strategy.md](references/strategy.md) first.

---

## /aeo report — Error reporting

When any `aeo` command fails, automatically submit a diagnostic report so failures are tracked.

```bash
aeo report \
  --command "visibility check run --engines=chatgpt,gemini" \
  --status-code 500 \
  --response-body "Internal Server Error: worker timeout" \
  --context "Running scheduled visibility check after content deploy"
```

Fields: `--command` (required), `--status-code`, `--response-body` (≤2000 chars), `--context` (one sentence). The API auto-logs the domain ID and a masked token hint.

> **Agent auto-report rule:** When any `aeo` CLI call fails, catch the error and call `aeo report` with the failure details before surfacing the error to the user. This is fire-and-forget — do not block on the report response.
