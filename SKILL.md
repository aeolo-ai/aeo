---
name: aeo
description: |
  Aeolo is an organic content engine: turn deep brand understanding into marketing
  content that compounds without ad spend. One output channel — blog articles
  cited in AI search (GEO: ChatGPT, Perplexity, Gemini). Reels/TikToks and competitor
  posts stay in scope as reference input you analyze to brief those articles, never as
  output. It loads real Aeolo brand data and writes back changes (brand updates,
  article generation, Shopify deployment) to run the full organic content cycle
  autonomously.
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

> **Aliases** (legacy spellings work on every surface: the CLI handles them or forwards them to the router; prefer canonical forms in new examples): `diagnose` = `visibility`/`audit` · `measure` = `metrics` · `account` = `billing`/`whoami` · `publish` mirrors `content deploy|redeploy|preview`. Noun-first plurals (`posts`, `channels`) and legacy verbs (`brand update`) route too.

> **Retired surfaces (2026-07-27)** — 410: `carousel create`, `post import` (and the `post write` workflow), `video generate`. **Never offer or call these** — route every content request to a blog article. Existing rows only, never for new work: `carousel update`, `post preview/approve/publish`. Analysis input stays live (`post analyze`, `reference *`, `video analyze`, `video poll`), as do read/cleanup verbs. [retired.md](references/retired.md)

| Area | Commands | Reference |
|------|----------|-----------|
| `domain` | `list` · `add <url>` (new brands only — never re-analyzes; existing ones use `rescan`) · `switch [id]` · `brand update` (`--markets ko-KR,en-US` sets the whole reach; `--family-json` replaces the brand-family roster; `--voice-json` / `--competitors-json` / `--ceps-json` / `--description` replace those Brand Understanding fields) · `brand aliases` (roster + alias candidates; suggests only) · `audit` · `channels` · `setup [--format json]` (server checklist and canonical strategy) · `rescan` | [brand.md](references/brand.md), [setup.md](references/setup.md) |
| `agent` | `context` · `ask "<message>"` (the Aeolo agent judges; paid plans) · `approve --thread <id> "<approval>"` · `result <turnId>` | [brand.md](references/brand.md) |
| `channel` | `list` · `add` · `update <id>` · `indexing <id>` · `delete <id>` · `connect <id>` (legacy social OAuth) · `disconnect <id>` · `voice` · `blog show` · `blog bind <host>` / `blog unbind` | [channels.md](references/channels.md), [tov-extract.md](references/tov-extract.md) |
| `pages` | `access` · `channels` · `list/get` · `create/update` · DOM `elements/edit` · `editions/review` · `publication/publish` · `trash/restore` · `share/unshare` (read-only preview_url; does not publish) · `builds` · `research` · `comments` (active Custom domain members; editor required for page changes; explicit reviewed-revision publishing) | [pages.md](references/pages.md) |
| `site folders` | `list` (managed-site folder paths and IDs; pass `--folder /path` on generate/import) | [channels.md](references/channels.md) |
| `site template` | `show` (where the article slot sits — warns when it landed in the header/footer) · `read` · `edit` (exact-match patches, or ops by offsets from a preview selection; chat/MCP via `templateUpdate`, confirm first) · `revisions` · `rollback <n>` · `recapture` | [channels.md](references/channels.md) |
| `visibility` | `summary [--market KR\|US] [--engine chatgpt] [--check-id UUID] [--format json]` · `show` · `history` · `check run` · `check poll <jobId>` | [visibility.md](references/visibility.md), [polling.md](references/polling.md) |
| `audit` | `run` (dynamic credits, scales with pages) · `poll <jobId>` | [polling.md](references/polling.md) |
| `strategy` | `show` · `update` · `visual` (style guide + mood boards; read it before editing — the boards are URL-keyed) · `visual update` (`--description` / `--keywords`; `--add-images` / `--remove-images` [`--product <id>`]; `--definitions` JSON merge) | [strategy.md](references/strategy.md) |
| `automation` | `schedules [--format json]` · `schedule set --track visibility\|content` (partial updates safe — omitted flags inherit; visibility adds `--repeat N`; reads show saved/effective repeat and plan limit; manual run repeat is separate) · `notify` / `notify set` (draft-review Slack/email rails) | [workflows.md](references/workflows.md) |
| `content` | `list` · `feed` · `get <id>` · `import` · `generate` (explicit-only paid job; needs a destination — `--target-channel`, or the domain's sole publish surface) · `jobs` · `update <id>` · `slug <id> <new-slug>` (moves a live URL) · `thumbnail <id> --url` · `preview <id>` · `deploy <id>` · `redeploy <id>` · `unpublish <id>` · `editions <id>` (locale fan-out; the canonical must be approved first) · `delete <id>` (deployed copies stay live — `unpublish` first) · `job cancel <jobId>` · `review <id>` | [content-manage.md](references/content-manage.md), [content-create.md](references/content-create.md), [content-review.md](references/content-review.md) |
| `post` | `analyze --url` (LIVE reference input) · `list` / `get <id>` / `delete <id>` (read/cleanup) | [tov-extract.md](references/tov-extract.md) |
| `reference` / `video` | `reference analyze --url --media` · `reference style [--url <url>] [--provider <p>]` (no `--url` = the domain's own channel voice) · `reference poll <jobId>` · `reference delete <jobId>` · `video analyze --url` (15 credits, synchronous) | [tov-extract.md](references/tov-extract.md), [polling.md](references/polling.md) |
| `measure` (alias `metrics`) | `overview` · `content <id>` · `traffic` · `visibility` | [metrics.md](references/metrics.md) |
| `diag` | `report` (command-failure diagnostics; legacy `aeo report`, `aeo measure report`; distinct from `report snapshot`) | [metrics.md](references/metrics.md) |
| `report` | `snapshot` — freeze a rendered report into a secret/password share link (requires the rendered `--html`; the connector does not render server-side) | this file |
| `market-map` | bare `market-map` (show) · `run` · `poll <jobId>` · `populate` (rail → Topics; prompts follow via a background job) | [market-map.md](references/market-map.md) |
| `prompts` | `portfolio <domain_id>`† (the **only** safe way to restructure a live tracked set) · `list` · `add` · `update <id>` (`--regions US,GB` = standing target markets, measured once per market per check) · `delete <id>` · `prompts generate` (seeds an **empty** set only — previews nothing, saves straight to tracked, and has shrunk a live portfolio 30→14; restructuring goes through `prompts portfolio`) | [brand.md](references/brand.md), [prompt-portfolio.md](references/prompt-portfolio.md) |
| `topics` | `demand` (saved demand + freshness + credit quote) · `demand run --max-credits 1` (one stale batch, up to 30 Topics) / `demand poll <jobId>` · `next` (one article angle; saves nothing) · `candidates` / `candidates poll <jobId>` (server job: demand-priced candidates + a `--demand-token`) · `prompts` / `prompts poll [jobId]` (writes tracked Prompts into empty Topics — **use this, not `prompts generate`**) · `list` · `create` (pass `--demand-token` to keep measured demand) · `update <id>` · `archive <id>` / `restore <id>` (both need `--revision`) · `assign-prompts <id>` | [topics.md](references/topics.md), [setup.md](references/setup.md) |
| `segments` | `list` (tags are metadata/filtering only) | [brand.md](references/brand.md) |
| `products` / `image` | Covers are **generated then curated**: `image generate` (5 credits per candidate — `--brand-style` + `--product <id>` for brand look with the real product, `--sweep N` then pick by eye, or `--recommend <contentId>` to compose from the article's own brief) · `image poll <jobId…>` · `content thumbnail <id> --url` (pin) · `image upload` · `image gallery [delete <assetId>]` · `image search` · `image swap` (5 credits — reproduces the packshot verbatim, so a box packshot makes a box ad) · `products` · `product add --pdp <url>` · `products discover` · `products rescan`. Read `strategy visual` first — it is the only channel brand knowledge has into images. | [image-thumbnails.md](references/image-thumbnails.md) |
| `source` | `youtube <url>` (verify a video **before** embedding — returns the canonical snippet + transcript; at most one embed per article) · `products` (catalog product/SKU tokens + the exact `product-carousel` marker; at most one carousel, ≤3 tokens) | [content-create.md](references/content-create.md) |
| `feedback` | `feedback [<message>]` (bare form opens `$EDITOR`) | this file |
| `drive` | `list` · `read <file_id>` · `download <file_id>` (supported/unsupported file types in reference) | [drive.md](references/drive.md) |
| `integrations google` | `status` (`ready` = a property/site is **selected**, not merely connected) · `properties` · `sites` · `set` — connecting Google itself is a browser sign-in on Traffic (`/traffic-analysis`) or `/connect/google?domain=<id>`, not a CLI command | [setup.md](references/setup.md) |
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

Never call a write API without confirmation. Always show what you're about to do and ask "Proceed?" first. Be extra explicit for live-site writes: `content deploy`/`redeploy`/`unpublish`/`slug`, `pages publish`, `site template edit`/`rollback`/`recapture`, `blog bind`.

Retired writes (`carousel create`, `post import`, `video generate`) are **not** on this list — do not offer or call them at all. See the retired-surfaces note in the Command Index.

## Communication Rules

- **UUID is internal only.** User-facing messages must use `title`, `name`, `domain`, `canonical`, etc. Example: `"bc2ef290-..." updated` → `"Best Project Management Tools for Startups" updated`
- **Agent-only workflows** (no bare CLI verb — they need external-agent reasoning): for a judgment (which article next and why, a Topic architecture, a Prompt validity/utility audit, why visibility moved) run `aeo agent ask "<request>"`: the Aeolo agent answers read-only from the brand understanding and strategy it holds (paid plans; on a plan error, tell the user). Relay proposed changes and run `aeo agent approve --thread <id> "<the user's words>"` only after the user approves; `/aeo prompts portfolio <domain_id>` safely restructures the tracked Prompt set (preview → confirm → atomic write → verify). Manual drafting writes directly then `aeo content import`. `aeo content generate` is the explicit paid server-side job; `aeo content review <id>` loads the review workspace; staff can run an independent review with `aeo content review-run <id> --key <uuid>` and read its saved result with `aeo content review-result <id>`; authorized editors read `aeo content approval <id>` and record explicit version-bound approval with `aeo content approve <id> --fingerprint <hash> --confirm` (see [content-review.md](references/content-review.md)). The `post write` → `post import` workflow is retired.
- **Prefer explicit verbs**: `aeo content list`, `aeo visibility show`, etc. Some bare nouns read (`aeo strategy`, `aeo prompts`, `aeo channels`, `aeo market-map`, `aeo blog`, `aeo content --limit 5`); others show sub-help.
- **Dashboard screens** to point users to: AI visibility `/geo-analysis` (tabs incl. SEO rankings `?tab=seo-lab`, no CLI command; site audit `?tab=audit`) · Traffic `/traffic-analysis` · Sites `/channels` · Content `/contents/blog` · Pages `/contents/pages` · Images `/contents/visual?tab=gallery` · Writing settings `/contents/writing-settings` · Strategist `/contents/strategist` · Writing sources `/contents/writing-sources` · Automation `/deployment-calendar` · Reports `/reports`.

Before editing brand documents, read [document-roles.md](references/document-roles.md), the generated canonical ownership contract. Preserve original text and approval/hold conditions; propose destination and exact before/after before using existing save commands.

For the next article, start from `aeo topics next`; for topic planning, read [strategy.md](references/strategy.md). Do not claim why a material, a fact or a planning change helps AI search beyond what the brand's own Aeolo measurements show. Writing, factual review, editorial suggestions and publishing decisions have separate roles; use [content-review.md](references/content-review.md) for that boundary.

---

## /aeo report — Error reporting

When any `aeo` command fails, automatically submit a diagnostic report so failures are tracked.

```bash
aeo diag report \
  --command "visibility check run --engines=chatgpt,gemini" \
  --status-code 500 \
  --response-body "Internal Server Error: worker timeout" \
  --context "Running scheduled visibility check after content deploy"
```

Fields: `--command` (required), `--status-code`, `--response-body` (≤2000 chars), `--context` (one sentence). The API auto-logs the domain ID and a masked token hint.

> **Agent auto-report rule:** When any `aeo` CLI call fails, catch the error and call `aeo diag report` with the failure details before surfacing the error to the user. This is fire-and-forget — do not block on the report response.
