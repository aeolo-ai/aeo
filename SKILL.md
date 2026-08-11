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

## Command Reference

> **Aliases** (normalized on the dashboard-chat + MCP agent surfaces; the raw `aeo` terminal binary ships only the canonical form): `diagnose` is canonical for `visibility`/`audit`; `measure` for `metrics` (`metrics overview` = `measure overview`); `account` for `billing`/`whoami`. `aeo publish` mirrors `content deploy|redeploy|preview`. Noun-first plurals (`posts`, `channels`) and legacy verbs (`brand update`) route too; prefer the canonical form in shell examples.

### Retired surfaces — TBD (2026-07-27)

Three output channels are sunset. **Do not offer them or call their output commands** — route every content request to a blog article. Full detail in [retired.md](references/retired.md).

| Retired | Commands | Instead |
|---------|----------|---------|
| Card news | `carousel create`, `carousel update` | blog article |
| Channel posts / Threads | `post write`, `post import`, `post preview`, `post approve`, `post publish` | blog article |
| Shortform video output | `video generate`, `video poll` | blog article |

> **Only output generation is retired — reference input is fully live.** `post analyze`, `reference style`, `channel voice`, and read/cleanup verbs on existing rows all stay.

### aeo domain — Domain selection & brand metadata

| Command | What it does | Reference |
|---------|-------------|-----------|
| `/aeo domain list` | List accessible domains | [brand.md](references/brand.md) |
| `/aeo domain add <url>` | Onboard a new brand; crawls in the background. Never re-analyzes an existing domain — use `domain rescan`. | [brand.md](references/brand.md) |
| `/aeo domain switch [id]` | Switch active domain | [brand.md](references/brand.md) |
| `/aeo domain brand update` | Update brand context fields (bare `aeo brand update` still routes as a legacy alias) | [brand.md](references/brand.md) |
| `/aeo domain audit` | Show latest audit report | [brand.md](references/brand.md) |
| `/aeo domain channels` | List connected channels (platform, status, ID) | [channels.md](references/channels.md) |
| `/aeo domain setup` | Onboarding checklist — market/language, topics, prompts, first check, Google, somewhere to publish. Drive/strategy sit under Next and don't count. | [setup.md](references/setup.md) |
| `/aeo domain rescan` | Re-crawl the domain and refresh its brand snapshot (name, category, value prop, typography/colors). Not credit-metered; synchronous. | [brand.md](references/brand.md) |

### aeo agent — Agent context

| Command | What it does | Reference |
|---------|-------------|-----------|
| `/aeo agent context` | Show the same default brand operating context used by the dashboard agent | [brand.md](references/brand.md) |

### aeo channel — Channel management

| Command | What it does | Reference |
|---------|-------------|-----------|
| `/aeo channel list` | List connected channels | [channels.md](references/channels.md) |
| `/aeo channel add` | Add a channel (--url required, --type, --label) | [channels.md](references/channels.md) |
| `/aeo channel update <id>` | Update a channel: `--url`, `--type`, `--label`; `--auto-publish true\|false`; `--publish-target <id>` (shopify blog / wordpress category / cafe24 board, picked by the channel's platform) | [channels.md](references/channels.md) |
| `/aeo channel indexing <id> --enabled true\|false` | Toggle IndexNow auto-indexing; add `--backfill` to submit existing published articles | [channels.md](references/channels.md) |
| `/aeo channel delete <id>` | Delete a non-primary channel | [channels.md](references/channels.md) |
| `/aeo channel connect <id>` | OAuth connect — opens browser for threads/linkedin/reddit | [channels.md](references/channels.md) |
| `/aeo channel disconnect <id>` | Disconnect OAuth integration from a channel | [channels.md](references/channels.md) |
| `/aeo channel voice` | Read selected channel-voice / reference-style evidence (`--provider`, `--url`) — same data as `reference style` | [tov-extract.md](references/tov-extract.md) |
| `/aeo blog bind <host>` | Bind hosted blog to their host | [channels.md](references/channels.md) |

### aeo visibility — Visibility data & checks

| Command | What it does | Reference |
|---------|-------------|-----------|
| `/aeo visibility show` | Show the last visibility snapshot | [visibility.md](references/visibility.md) |
| `/aeo visibility history` | Score over time, one row per check — use for trends (`--limit <n>`) | [visibility.md](references/visibility.md) |
| `/aeo visibility check run` | Run a credit-metered visibility check (`--engines <list>`, `--location KR\|US\|EU\|JP`, `--limit <n>`, `--prompt-ids <id,id>`). Unknown engines error out with the valid set; unsupported `--location` is rejected. | [visibility.md](references/visibility.md), [polling.md](references/polling.md) |
| `/aeo visibility check poll <jobId>` | Poll check status | [visibility.md](references/visibility.md), [polling.md](references/polling.md) |

### aeo audit — Site foundation checks

| Command | What it does | Reference |
|---------|-------------|-----------|
| `/aeo audit run` | Start a site foundation audit (`--max-pages <n>` default 5, credit cost is dynamic and scales with pages crawled; `--channel-id <id>`) | [polling.md](references/polling.md) |
| `/aeo audit poll <jobId>` | Poll an audit job | [polling.md](references/polling.md) |

### aeo strategy — Content strategy

| Command | What it does | Reference |
|---------|-------------|-----------|
| `/aeo strategy show` | Show current content strategy (manifest) | [strategy.md](references/strategy.md) |
| `/aeo strategy update` | Create or update content strategy | [strategy.md](references/strategy.md) |
| `/aeo strategy visual update` | Update the visual style guide used by image/video generation (`--description <text>`, `--keywords a,b,c`) | this file |

### aeo automation — Automation schedules

Drives the Automation page (`/deployment-calendar`): the two tracks Aeolo runs on a cadence — **visibility** (the weekly AI auto-scan) and **content** (the autonomous writing loop). Both upsert on `(domain_id, task_kind)`.

| Command | What it does | Reference |
|---------|-------------|-----------|
| `/aeo automation schedules` | Show both automation tracks (state, frequency, days, timezone; content adds autonomy, articles/run, hero image) | [workflows.md](references/workflows.md) |
| `/aeo automation schedule set --track visibility\|content` | Configure one track — `--enabled true\|false`, `--frequency daily\|weekly\|biweekly`, `--days mon,tue`, `--timezone <tz>`; content-only `--autonomy draft\|auto`, `--articles-per-run N` (1-14), `--hero true\|false` | [workflows.md](references/workflows.md) |

> Partial updates are safe: `automation schedule set` inherits existing settings for any omitted flag (so `--enabled false` alone just pauses the track). The visibility track is also how you toggle the weekly auto-scan.

### aeo content — Content lifecycle

| Command | What it does | Reference |
|---------|-------------|-----------|
| `/aeo content idea <domain_id>` | Agent-only, read-only workflow: inspect live Aeolo signals and recommend one next article idea. This is not a bare CLI verb. | [content-idea.md](references/content-idea.md) |
| `/aeo content list` | List content items (--status, --limit, --offset) | [content-manage.md](references/content-manage.md) |
| `/aeo content feed` | Content Feed URLs + JSON Feed contract to render Aeolo articles on your own site | [content-manage.md](references/content-manage.md) |
| `/aeo content get <id>` | Read full article content (markdown) | [content-manage.md](references/content-manage.md) |
| `/aeo content import` | Push an already-written draft to content history | [content-create.md](references/content-create.md) |
| `/aeo content generate` | Explicit-only server-side content generation job (5 credits) | [content-create.md](references/content-create.md), [polling.md](references/polling.md) |
| `/aeo content jobs` | List active content generation jobs | [polling.md](references/polling.md) |
| `/aeo content update <id>` | Update a content item (`--status`, `--title`, `--meta-description`, `--keywords`, `--body`/`--body-file` full replace or `--patch "search>>>replace"` targeted edit, `--thumbnail-url`, `--clear-thumbnail`) | [content-manage.md](references/content-manage.md) |
| `/aeo content preview <id>` | Generate a shareable preview link (prints the URL; does not open a browser) | [content-manage.md](references/content-manage.md) |
| `/aeo content deploy <id> [--target shopify\|blog\|wordpress\|cafe24]` | Deploy an approved article to a publish destination (default `shopify`; `blog` = hosted Aeolo blog) | [content-manage.md](references/content-manage.md) |
| `/aeo content redeploy <id> [--target shopify\|blog\|wordpress\|cafe24]` | Push the current article back in-place (keeps URL); target auto-detected, `--target` forces one. Hosted `blog` needs no push. | [content-manage.md](references/content-manage.md) |
| `/aeo content unpublish <id>` | Remove a deployed article from its platform (auto-detected; `--target` to force) and reset it to draft | [content-manage.md](references/content-manage.md) |
| `/aeo content delete <id>` | Soft-delete an article (drops out of lists; restorable). Deployed copies stay live — `content unpublish` first to remove them. | [content-manage.md](references/content-manage.md) |
| `/aeo content job cancel <jobId>` | Delete a finished writing job + its events (blocked while pending/running — let active jobs finish or fail) | [polling.md](references/polling.md) |
| `/aeo content review <id>` | GEO content review (structure, trust, freshness, brand, engine fit) | [content-review.md](references/content-review.md) |

### aeo carousel — RETIRED

Card news is sunset (generation returns HTTP 410, no credits charged); read/cleanup verbs `carousel list`/`get`/`delete` stay. See [retired.md](references/retired.md).

### aeo post — RETIRED (output only) · `post analyze` is LIVE

Channel-post publishing is sunset. `post analyze` (reference voice evidence) and `post list`/`get`/`delete` (read/cleanup) stay live. See [retired.md](references/retired.md).

| Command | What it does | Reference |
|---------|-------------|-----------|
| `/aeo post analyze --url <URL>` | **LIVE** (reference input) — analyze one channel/reference URL for voice evidence (`--provider blog\|threads\|tiktok\|instagram`, `--mode owned\|reference`, `--limit`) | [tov-extract.md](references/tov-extract.md) |
| `/aeo post list` / `/aeo post get <id>` / `/aeo post delete <id>` | Read/cleanup on existing rows (--platform, --status, --limit, --offset) | [channel-washing.md](references/channel-washing.md) |

### aeo reference / video — Analysis (LIVE); generation retired

> **Analysis in, generation out.** Every command that *reads* an external Reel/TikTok/post to brief an article is live. `video generate`/`video poll` are retired — see [retired.md](references/retired.md).

| Command | What it does | Reference |
|---------|-------------|-----------|
| `/aeo reference analyze --url <url> --media <type>` | **LIVE** — analyze a reference URL as a background job. `--media linkedin_post\|threads_post\|visual_asset\|instagram_reels\|tiktok_reels`, `--language` optional | [tov-extract.md](references/tov-extract.md) |
| `/aeo reference style --url <url>` | **LIVE** — read reference style evidence (`--provider blog\|threads\|linkedin\|instagram\|tiktok`) | [tov-extract.md](references/tov-extract.md) |
| `/aeo reference poll <jobId>` | Poll a reference analysis job | [polling.md](references/polling.md) |
| `/aeo reference delete <jobId>` | Delete a reference/style/channel-voice job (auto-detected; soft-delete) | [tov-extract.md](references/tov-extract.md) |
| `/aeo video analyze --url <url>` | **LIVE** (reference input) — analyze a short-form video URL synchronously (15 credits). `--media instagram_reels\|tiktok_reels` | [tov-extract.md](references/tov-extract.md) |

### aeo measure / metrics — Article & site performance

Reference: [metrics.md](references/metrics.md)

> `measure` is canonical; `metrics` is the alias (`metrics overview` = `measure overview`, `metrics article <id>` = `measure content <id>`). `measure visibility` = same data as `aeo visibility show`.
>
> **Two "report" commands:** `diag report` (alias `measure report`) submits command-failure diagnostics to the team; `report snapshot` creates a shareable customer report link (see `aeo report` below).

| Command | What it does |
| --------- | ------------- |
| `/aeo measure overview` (alias `metrics overview`) | Show deployed articles with GA4 + GSC stats (last 30 days) |
| `/aeo measure content <id>` (alias `metrics article <id>`) | Detailed per-article stats (traffic sources, top queries) |
| `/aeo measure traffic` (alias `metrics traffic`) | Site-level GSC traffic: top queries, pages, country, device (--days=7\|14\|30\|90) |
| `/aeo measure visibility` | Show last visibility snapshot (same data as `aeo visibility show`) |
| `/aeo measure attribution` | First-touch AI attribution (Traffic/Attribution page): attributed sessions, revenue, CVR, AOV + sessions by AI source (`--days 7\|30\|90`, default 30) |
| `/aeo diag report --command <cmd>` (alias `measure report`) | Submit command-failure diagnostics (`--status-code`, `--response-body`, `--context`) |

### aeo report — Shareable report links

| Command | What it does | Reference |
|---------|-------------|-----------|
| `/aeo report snapshot` | Create a shareable report link. Freeze a rendered report and mint a secret/password share URL: `--type audit\|visibility\|traffic\|prompts\|performance` (required), `--html <rendered report HTML>` (required, ≥100 chars — the connector does not render server-side), `--access-mode secret_link\|password`, `--password <pw>`, `--expires 7\|30\|90\|never`, `--locale en\|ko`. | this file |

### aeo market-map — Market map (upstream of prompts/topics)

Journey-stage demand, homeground vocabulary, surface holders, topic rail of
judged prompt candidates — read [market-map.md](references/market-map.md) first.

| Command | What it does |
|---------|-------------|
| `/aeo market-map` | Show the map (`--market KR\|US\|JP\|TW\|HK\|CN\|GB\|ES\|MX`, latest when omitted) |
| `/aeo market-map run` | Build (~3 min job); needs snapshot products + `identity.exclusions` |
| `/aeo market-map poll <jobId>` | Poll the build job |
| `/aeo market-map populate` | Rail → Topics; prompts follow via a job. `--topics "…"` (default: all non-held), `--prompts N` caps the count |

### aeo prompts — Tracked prompts

| Command | What it does | Reference |
|---------|-------------|-----------|
| `/aeo prompts portfolio <domain_id>` | Agent-only workflow: the **only** safe way to restructure the tracked Prompt set (recover a shrunk portfolio, swap categories, hit an exact count). read → propose diff → audit → preview → confirm → atomic write → verify. Not a bare CLI verb. | [prompt-portfolio.md](references/prompt-portfolio.md) |
| `/aeo prompts audit <domain_id>` | Agent-only, read-only workflow: evaluate each tracked Prompt category's validity and whether its live content is working, then give a natural-language operating recommendation. This is not a bare CLI verb. | [prompt-audit.md](references/prompt-audit.md) |
| `/aeo prompts list` | List prompts grouped by stage; optionally filter with `--status tracked\|untracked` | [brand.md](references/brand.md) |
| `/aeo prompts add` | Add prompts to brand_prompts — one (`--prompt`) or up to 30 at once (`--prompts-json`); also `--stage`, `--language`, `--segment foo,bar` | [brand.md](references/brand.md) |
| `/aeo prompts update <id>` | Edit an existing prompt (`--prompt`, `--stage`, `--query-form`, `--segment foo,bar`, `--status tracked\|untracked`) | [brand.md](references/brand.md) |
| `/aeo prompts delete <id>` | Soft-delete a prompt by ID | [brand.md](references/brand.md) |
| `/aeo prompts generate` | Generate a CEP-based prompt set from brand context (free; needs brand analysis). `--count N`, `--languages en,ko`, `--instruction "..."`. **Not exact-count; saves straight to `tracked`. Never use it to restructure a live portfolio (can shrink 30→14) — use `prompts portfolio`. Safe only to seed an empty one.** | [prompt-portfolio.md](references/prompt-portfolio.md) |

### aeo topics — Stable customer situations

Two senses of "topic" on this noun: `topics list/create/update/…` manage durable
Topic **entities**; `topics next` suggests the next **article angle** and saves nothing.

| Command | What it does | Reference |
|---------|-------------|-----------|
| `/aeo topics next` | One next article angle + the evidence for it. Serves an unwritten research brief, else ranks tracked queries by citation gap, demand, intent, coverage. Free, read-only, deterministic; same picker the writing cron uses. | [workflows.md](references/workflows.md) |
| `/aeo topics suggest <domain_id>` | Agent-only, read-only workflow: derive 3–5 durable business Topics from live brand, strategy, product, Prompt, visibility, traffic, and owned-content context. This is not a bare CLI verb. | [topic-suggest.md](references/topic-suggest.md) |
| `/aeo topics list` | List Topics with status, revision, and Prompt counts (`--include-archived`) | [topics.md](references/topics.md) |
| `/aeo topics create` | Create a Topic (`--name` required, `--description` optional) | [topics.md](references/topics.md) |
| `/aeo topics update <id>` | Rename or redefine a Topic with required `--revision` | [topics.md](references/topics.md) |
| `/aeo topics archive <id>` | Archive an empty Topic with required `--revision`; move Prompts first | [topics.md](references/topics.md) |
| `/aeo topics restore <id>` | Restore an archived Topic with required `--revision` | [topics.md](references/topics.md) |
| `/aeo topics assign-prompts <id>` | Atomically reassign Prompts via `--prompt-ids id1,id2` | [topics.md](references/topics.md) |

### aeo segments — Segment tags

| Command | What it does | Reference |
|---------|-------------|-----------|
| `/aeo segments list` | List segment tags with prompt counts (tags are metadata/filtering only) | [brand.md](references/brand.md) |

### aeo products / image — Thumbnail pipeline (Pexels + product swap)

Reference: [image-thumbnails.md](references/image-thumbnails.md)

| Command | What it does |
| --------- | ------------- |
| `/aeo products` (or `/aeo product list`) | List the product catalog (IDs + image status) used as swap sources |
| `/aeo product add --pdp <url>` | Add a product by PDP URL (scrapes title/image/price) |
| `/aeo products discover` | Crawl the domain sitemap for candidate PDP URLs to add (read-only; flags URLs already in the catalog) |
| `/aeo products rescan` | Re-scrape every product PDP (up to 30) to backfill/refresh media, title, image, price |
| `/aeo image search <query>` | Search Pexels for reference scenes (--per-page, --page) |
| `/aeo image swap --content <id> --product <id> --reference <url>` | Generate a thumbnail by swapping a product into a reference scene (5 credits) |
| `/aeo image upload --file <path>` | Upload a local image (≤25 MP) to the thumbnail bucket (--content to pin) |
| `/aeo image generate --prompt <text>` | Generate image(s) from a text prompt for thumbnails/gallery (credit-metered; cost scales with model + `--sweep`). `--model nano-banana-pro\|gpt-image-2\|grok-image`, `--sweep N` (1-8), `--aspect`, `--resolution`, `--ref`, `--brand-style`. Async. |
| `/aeo image poll <jobId...>` | Check status + result URLs of image generation jobs |
| `/aeo image gallery [--type image\|video] [--limit N]` | List generated gallery assets for the domain (default `image`, limit 20) |
| `/aeo image gallery delete <assetId>` | Delete a generated gallery asset by ID |

### aeo feedback — Send feedback to the team

| Command | What it does | Reference |
|---------|-------------|-----------|
| `/aeo feedback [<message>]` | Send feedback (bug, idea, anything) to the Aeolo team; bare form opens $EDITOR (requires $EDITOR/$VISUAL set) | this file |

### aeo drive — Google Drive files

Reference: [drive.md](references/drive.md)

| Command | What it does |
| --------- | ------------- |
| `/aeo drive list` | List files in connected Google Drive folder (--folder) |
| `/aeo drive read <file_id>` | Read a file from Google Drive |
| `/aeo drive download <file_id>` | Stream raw bytes to disk (pptx, large/binary files); `-o <path>` to set output |

> **Supported types**: Google Docs/Sheets, txt/json/md/csv, **PDF**, **XLSX/XLS** (all sheets, 200-row cap each), **DOCX**, images (≤5MB base64). Not supported: `.doc`, `.pptx`, `.pages`, `.numbers`, `.key` — see [drive.md](references/drive.md).

### aeo integrations google — GA4 / GSC selection

Reference: [metrics.md](references/metrics.md)

| Command | What it does |
| --------- | ------------- |
| `/aeo integrations google status` | Is Google usable here? `ready` means a property/site is **selected**, not merely connected — connected-with-nothing-selected loads no data. |
| `/aeo integrations google properties` | List GA4 properties accessible to the connected Google account (for selection) |
| `/aeo integrations google sites` | List Search Console sites accessible to the connected Google account (for selection) |
| `/aeo integrations google set [--ga4-property <id>] [--gsc-site <url>]` | Select the GA4 property and/or GSC site for this domain (powers `measure overview`/`measure traffic`) |

> Connecting Google is a browser sign-in (**Settings → domain integrations**), not a CLI command. If not connected, these return a pointer to Settings — connect first, then pick a property/site.

### aeo agency — Execution-agency matching waitlist

| Command | What it does | Reference |
|---------|-------------|-----------|
| `/aeo agency request --types seo,pr,blog,reddit,youtube,influencer [--other "..."]` | Join the execution-agency matching waitlist (Earned Media early access). Matching isn't live yet; this records interest and pings the team. Idempotent per (domain, type). | this file |

### aeo gsc — Google Search Console (browser automation)

| Command | What it does | Reference |
|---------|-------------|-----------|
| `/aeo gsc index` | Bulk indexing request via browser automation (requires Chrome + GSC login) | [gsc-indexing.md](references/gsc-indexing.md) |
| `/aeo gsc index --domain <domain>` | Specify target domain (skips domain prompt) | [gsc-indexing.md](references/gsc-indexing.md) |

> **Environment**: Uses Claude in Chrome browser automation, NOT the aeo CLI — requires Chrome + the Claude in Chrome extension and a GSC login. If unavailable, guide the user to do it manually.

### aeo config — Agent configuration

| Command | What it does | Reference |
|---------|-------------|-----------|
| `/aeo config data-sources` | Show configured data sources for research | [data-access.md](references/data-access.md) |
| `/aeo config data-sources update` | Update data source pointers (--data-sources) | [data-access.md](references/data-access.md) |

### aeo auth — Authentication

Reference: [setup.md](references/setup.md)

| Command | What it does |
| --------- | ------------- |
| `/aeo auth login` | Authenticate via browser (device flow) |
| `/aeo auth status` | Show current stored credentials |
| `/aeo auth logout` | Clear stored credentials |

### aeo account — Account & subscription

> `account` is canonical; `billing`/`whoami` are aliases (`aeo billing subscription` = `aeo account subscription`, bare `aeo whoami` = `aeo account whoami`).

| Command | What it does |
| --------- | ------------- |
| `/aeo account whoami` (or `/aeo whoami`) | Show current user (email, tier, trial days) |
| `/aeo account subscription` | Show current subscription, tier, and credit summary |
| `/aeo account credits` | Show current credit balance |
| `/aeo account ledger` | Show credit ledger entries (`--days` default 30, `--limit` default 50) |

### Utilities

| Command | What it does | Reference |
|---------|-------------|-----------|
| `/aeo` | Load full GEO context (brand + audit + visibility) | [workflows.md](references/workflows.md) |
| `/aeo report` | Submit an error report when a command fails | this file |

---

## Workflows — Autonomous GEO Optimization

Run the full GEO cycle — brand setup → daily content → weekly review. See [workflows.md](references/workflows.md) for decision logic and quality gates.

| Workflow | When | What |
|----------|------|------|
| **Onboarding** | New brand, first setup | Assess setup → auto-fill what you can → guide user for OAuth/permissions → verify 5/5 |
| **Daily Content** | Every day (cron or manual) | Pick topic from priority queue → write → deploy → request indexing (channel-post distribution is retired) |
| **Weekly Report** | Every week (cron or manual) | Visibility check → performance analysis → strategy adjustment → report to user |

Start with `aeo domain setup` — a 5-item integration checklist (Brand, Shopify, GA4+GSC, Drive, Strategy). Install/auth detail in [setup.md](references/setup.md); don't start the loops until all 5 are complete.

---

Read the relevant reference file before executing any command.

## CUD Rule

**Always get explicit user confirmation before any Create / Update / Delete operation.**

Applies to **every write/execute/external command** — anything that is not a pure read; the registry riskLevel is the source of truth (content, prompts, topics, brand/strategy/config, channels, images, `market-map run`, `audit run`, …).

Never call a write API without confirmation. Always show what you're about to do and ask "Proceed?" first. Be extra explicit for the irreversible ones: `content deploy`/`content redeploy`/`content unpublish` change the customer's live site.

Retired writes (`carousel create/update`, `post import/preview/approve/publish`, `video generate`) are **not** on this list — do not offer or call them at all. See [Retired surfaces](#retired-surfaces--tbd-2026-07-27).

## Communication Rules

- **UUID is internal only.** User-facing messages must use `title`, `name`, `domain`, `canonical`, etc. Example: `"bc2ef290-..." updated` → `"Best Project Management Tools for Startups" updated`
- **Agent-only workflows** (no bare CLI verb — they need external-agent reasoning): `/aeo topics suggest <domain_id>` proposes a durable Topic architecture without saving; `/aeo content idea <domain_id>` recommends one next article; `/aeo prompts audit <domain_id>` judges Prompt validity/utility read-only; `/aeo prompts portfolio <domain_id>` safely restructures the tracked Prompt set (preview → confirm → atomic write → verify). Manual drafting writes directly then `aeo content import`. `aeo content generate` is the explicit paid server-side job; `aeo content review <id>` is a real wired command. The `post write` → `post import` workflow is retired.
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
