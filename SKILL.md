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

> **Noun.verb aliases**: the connector accepts noun-first aliases for many commands. `diagnose` is the canonical visibility/audit noun (`diagnose visibility`/`visibility run`/`visibility poll <jobId>`, `diagnose audit`/`audit run`/`audit poll <jobId>`) and `visibility`/`audit` are the friendly aliases. `measure` is canonical and `metrics` is the alias (`metrics overview` = `measure overview`). `account` is the canonical billing noun and `billing`/`whoami` are aliases (`billing subscription` = `account subscription`, `whoami` = `account whoami`). `aeo publish` is its own binary group (`publish preview|deploy|redeploy`) that mirrors `content deploy|redeploy|preview`. `posts`/`channels` noun forms also route. The CLI image/thumbnail nouns are `image` and `video` (the `content thumbnail`/`media` forms are connector-internal). Both forms route to the same endpoint — use whichever reads better.
>
> **Terminal vs agent surface**: the noun-first plural (`posts`, `channels`) and legacy-verb (`brand update`) aliases are normalized by the connector command registry, so they resolve on the **dashboard-chat and MCP agent surfaces**. The raw `aeo` terminal binary only ships the canonical forms — in shell examples prefer `post` (not `posts`) and `domain brand update` (not `brand update`).

### Retired surfaces — TBD (2026-07-27)

Three channels are sunset. **Do not offer them to the user, and do not call their output commands.**

| Retired | Commands | Instead |
|---------|----------|---------|
| Card news | `carousel create`, `carousel update` | blog article |
| Channel posts / Threads | `post write`, `post import`, `post preview`, `post approve`, `post publish` | blog article |
| Shortform video output | `video generate`, `video poll` | blog article |

The loop never closed for any of the three: Threads publishing was gated on a Meta app review that never landed, and reels drive Google SERP discovery but not chatbot citations. Blog articles are the only channel where diagnose → write → deploy → index → citation closes inside the product, so route every content request to `content generate` (or draft directly + `create import`) → `publish deploy`.

> **Only output generation is retired — reference input is fully live.** Analyzing someone else's Reel, TikTok, or post as *reference* to brief a blog article is unaffected: `diagnose references analyze`, `diagnose video analyze`, `reference style`, `channel voice`, and `post analyze` all stay. So do reads and cleanup on rows that already exist (`carousel list`/`get`/`delete`, `post list`/`get`/`delete`).

### aeo domain — Domain selection & brand metadata

| Command | What it does | Reference |
|---------|-------------|-----------|
| `/aeo domain list` | List accessible domains | this file |
| `/aeo domain switch [id]` | Switch active domain | this file |
| `/aeo domain brand update` | Update brand context fields (bare `aeo brand update` still routes as a legacy alias) | [brand.md](references/brand.md) |
| `/aeo domain audit` | Show latest audit report | this file |
| `/aeo domain channels` | List connected channels (platform, status, ID) | this file |
| `/aeo domain setup` | Show setup checklist (integrations status) | this file |
| `/aeo domain rescan` | Re-crawl the domain and refresh its brand snapshot (name, category, value prop, typography/colors). Not credit-metered; synchronous. | [brand.md](references/brand.md) |

### aeo agent — Agent context

| Command | What it does | Reference |
|---------|-------------|-----------|
| `/aeo agent context` | Show the same default brand operating context used by the dashboard agent | [brand.md](references/brand.md) |

### aeo channel — Channel management

| Command | What it does | Reference |
|---------|-------------|-----------|
| `/aeo channel list` | List connected channels | this file |
| `/aeo channel add` | Add a channel (--url required, --type, --label) | this file |
| `/aeo channel update <id>` | Update a channel: `--url`, `--type`, `--label`; `--auto-publish true\|false`; `--publish-target <id>` (shopify blog / wordpress category / cafe24 board, picked by the channel's platform) | this file |
| `/aeo channel indexing <id> --enabled true\|false` | Toggle IndexNow auto-indexing; add `--backfill` to submit existing published articles | this file |
| `/aeo channel delete <id>` | Delete a non-primary channel | this file |
| `/aeo channel connect <id>` | OAuth connect — opens browser for threads/linkedin/reddit | this file |
| `/aeo channel disconnect <id>` | Disconnect OAuth integration from a channel | this file |
| `/aeo channel voice` | Read selected channel-voice / reference-style evidence (`--provider`, `--url`) — same data as `reference style` | [tov-extract.md](references/tov-extract.md) |

### aeo visibility — Visibility data & checks

| Command | What it does | Reference |
|---------|-------------|-----------|
| `/aeo visibility show` | Show the last visibility snapshot | [visibility.md](references/visibility.md) |
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

> Partial updates are safe: `automation schedule set` inherits the track's existing settings for any flag you omit, so `--enabled false` alone just pauses the track without resetting its cadence. Configuring the visibility track is also how you turn the weekly visibility auto-scan on or off.

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
| `/aeo content redeploy <id>` | Update an already-deployed Shopify article in-place (keeps URL) | [content-manage.md](references/content-manage.md) |
| `/aeo content unpublish <id>` | Remove a deployed article from its platform (auto-detected; `--target` to force) and reset it to draft | [content-manage.md](references/content-manage.md) |
| `/aeo content delete <id>` | Soft-delete an article (drops out of lists; restorable). Deployed copies stay live — `content unpublish` first to remove them. | [content-manage.md](references/content-manage.md) |
| `/aeo content job cancel <jobId>` | Delete a finished writing job + its events (blocked while pending/running — let active jobs finish or fail) | [polling.md](references/polling.md) |
| `/aeo content review <id>` | GEO content review (structure, trust, freshness, brand, engine fit) | [content-review.md](references/content-review.md) |

### aeo carousel — Instagram carousel decks — RETIRED

> **Retired 2026-07-27 — card news is no longer a supported channel.** Do not offer it. Generation is closed server-side (HTTP 410, no credits charged); the read/cleanup verbs stay so existing decks remain visible and removable.

| Command | What it does | Reference |
|---------|-------------|-----------|
| `/aeo carousel list [--limit N]` | List carousel decks that already exist for the domain | [content-manage.md](references/content-manage.md) |
| `/aeo carousel get <jobId>` | Get an existing deck (slides + caption) by job ID | [content-manage.md](references/content-manage.md) |
| `/aeo carousel create [--topic "..."]` | **RETIRED** — card news generation is closed (returns 410). Write a blog article instead. | — |
| `/aeo carousel update <jobId> --caption "..."` | **RETIRED** — do not edit card news captions; the channel is sunset | — |
| `/aeo carousel delete <jobId>` | Soft-delete a deck (cancels active work; assets stay recoverable) | [content-manage.md](references/content-manage.md) |

> Default external-agent writing path: draft directly, then use `content import`. Use `content generate` only when the user explicitly wants an Aeolo server-side paid generation job.
>
> **Deploy targets**: `content deploy --target` picks the destination — `shopify` (default, connected Shopify blog), `blog` (the always-available hosted Aeolo blog, no connection needed), `wordpress`, or `cafe24`. If a target's channel isn't connected, the error names the target and points to aeolo.io → Owned Media.
>
> **Deploy gate**: `content deploy`/`content redeploy` reject with HTTP 422 unless the title is ≤ 60 chars and the meta description is present and 50–160 chars. The approved-status gate and this SERP title/meta gate apply to every target. Check with `content get` and fix with `content update` before deploying.

### aeo post — Channel posts (social media distribution) — RETIRED (output only)

> **Retired 2026-07-27 — channel-post publishing is no longer a supported channel.** Threads led this surface and its publishing was gated on a Meta app review that never landed, so the loop never closed. The `post` group is platform-agnostic, so the retirement covers every platform it served (threads / linkedin / reddit / instagram / x). Do not draft, import, approve, or publish channel posts — take the request to a blog article instead.
>
> `post analyze` is **not** retired: it reads an owned or competitor post as reference voice evidence for an article brief. `post list`/`get`/`delete` stay for reading and cleaning up rows that already exist.

| Command | What it does | Reference |
|---------|-------------|-----------|
| `/aeo post analyze --url <URL>` | **LIVE** (reference input) — analyze one channel/reference URL and propose task-specific voice evidence (`--provider blog\|threads\|tiktok\|instagram`, `--mode owned\|reference`, `--limit`) | [tov-extract.md](references/tov-extract.md) |
| `post write` | **RETIRED** — the channel-post writing workflow is sunset. Draft a blog article instead. | [post-create.md](references/post-create.md) |
| `/aeo post list` | List channel posts that already exist (--platform, --status, --limit, --offset) | [channel-washing.md](references/channel-washing.md) |
| `/aeo post get <id>` | Get an existing channel post (full body + metadata) | [channel-washing.md](references/channel-washing.md) |
| `/aeo post import` | **RETIRED** — do not import new channel-post drafts | — |
| `/aeo post preview <id>` | **RETIRED** — with the post pipeline | — |
| `/aeo post approve <id>` | **RETIRED** — with the post pipeline | — |
| `/aeo post publish <id>` | **RETIRED** — publishing to social platforms is closed | — |
| `/aeo post delete <id>` | Delete a channel post (cleanup on existing rows) | [channel-washing.md](references/channel-washing.md) |

### aeo reference / video — Analysis (live) & generation (retired)

> **Analysis in, generation out.** Every command that *reads* an external Reel, TikTok, or post to brief an article is live. Only `video generate` (and its `video poll`) is retired — shortform video output is no longer a supported channel because reels drive Google SERP discovery, not chatbot citations.

| Command | What it does | Reference |
|---------|-------------|-----------|
| `/aeo reference analyze --url <url> --media <type>` | **LIVE** (reference input) — analyze a reference URL as a background job (credit cost varies by media type). `--media linkedin_post\|threads_post\|visual_asset\|instagram_reels\|tiktok_reels`, `--language` optional. The reels/post media types stay supported: they are input, not output. | [tov-extract.md](references/tov-extract.md) |
| `/aeo reference style --url <url>` | **LIVE** (reference input) — read selected reference style evidence (--provider blog\|threads\|linkedin\|instagram\|tiktok) | [tov-extract.md](references/tov-extract.md) |
| `/aeo reference poll <jobId>` | Poll a reference analysis job | [polling.md](references/polling.md) |
| `/aeo reference delete <jobId>` | Delete a reference analysis job (or a reference-style/channel-voice job — auto-detected). Soft-delete; cancels active work first. | [tov-extract.md](references/tov-extract.md) |
| `/aeo video analyze --url <url>` | **LIVE** (reference input) — analyze a short-form video URL synchronously (15 credits). `--media instagram_reels\|tiktok_reels`, `--mime-type` optional | this file |
| `/aeo video generate --prompt <text>` | **RETIRED 2026-07-27** — shortform video output is sunset. Do not offer it; brief a blog article instead. | — |
| `/aeo video poll <jobId...>` | **RETIRED** with `video generate` — only reaches jobs created before the sunset | — |

### aeo measure / metrics — Article & site performance

> Canonical noun is `measure`; `metrics` is the accepted alias. `metrics overview` = `measure overview`, `metrics article <id>` = `measure content <id>`, `metrics traffic` = `measure traffic`. The `measure` noun also carries `measure visibility` (last visibility snapshot, same data as `aeo visibility show`).
>
> **Two different "report" commands — don't confuse them:** `diag report` (canonical; `measure report` is the deprecated alias) **submits command-failure diagnostics** to the Aeolo team. `report snapshot` **creates a shareable report link** for a customer. See the `aeo report` section below for `report snapshot`.

| Command | What it does | Reference |
|---------|-------------|-----------|
| `/aeo measure overview` (alias `metrics overview`) | Show deployed articles with GA4 + GSC stats (last 30 days) | [metrics.md](references/metrics.md) |
| `/aeo measure content <id>` (alias `metrics article <id>`) | Detailed per-article stats (traffic sources, top queries) | [metrics.md](references/metrics.md) |
| `/aeo measure traffic` (alias `metrics traffic`) | Site-level GSC traffic: top queries, pages, country, device (--days=7\|14\|30\|90) | [metrics.md](references/metrics.md) |
| `/aeo measure visibility` | Show last visibility snapshot (same data as `aeo visibility show`) | [metrics.md](references/metrics.md) |
| `/aeo measure attribution` | First-touch AI attribution (Traffic/Attribution page): attributed sessions, revenue, CVR, AOV + sessions by AI source (`--days 7\|30\|90`, default 30) | [metrics.md](references/metrics.md) |
| `/aeo diag report --command <cmd>` (alias `measure report`) | Submit command-failure diagnostics (`--status-code`, `--response-body`, `--context`) | [metrics.md](references/metrics.md) |

### aeo report — Shareable report links

| Command | What it does | Reference |
|---------|-------------|-----------|
| `/aeo report snapshot` | Create a shareable report link. Freeze a rendered report and mint a secret/password share URL: `--type audit\|visibility\|traffic\|prompts\|performance` (required), `--html <rendered report HTML>` (required, ≥100 chars — the connector does not render server-side), `--access-mode secret_link\|password`, `--password <pw>`, `--expires 7\|30\|90\|never`, `--locale en\|ko`. | this file |

### aeo prompts — Tracked prompts

| Command | What it does | Reference |
|---------|-------------|-----------|
| `/aeo prompts audit <domain_id>` | Agent-only, read-only workflow: evaluate each tracked Prompt category's validity and whether its live content is working, then give a natural-language operating recommendation. This is not a bare CLI verb. | [prompt-audit.md](references/prompt-audit.md) |
| `/aeo prompts list` | List prompts grouped by stage; optionally filter with `--status tracked\|untracked` | [brand.md](references/brand.md) |
| `/aeo prompts add` | Add prompts to brand_prompts — one (`--prompt`) or up to 30 at once (`--prompts-json`); also `--stage`, `--language`, `--segment foo,bar` | [brand.md](references/brand.md) |
| `/aeo prompts update <id>` | Edit an existing prompt (`--prompt`, `--stage`, `--query-form`, `--segment foo,bar`, `--status tracked\|untracked`) | [brand.md](references/brand.md) |
| `/aeo prompts delete <id>` | Soft-delete a prompt by ID | [brand.md](references/brand.md) |
| `/aeo prompts generate` | Generate a CEP-based prompt set from brand context and save it to tracked prompts (free; needs brand analysis first). `--count N`, `--languages en,ko`, `--instruction "..."` | [brand.md](references/brand.md) |

### aeo topics — Stable customer situations

| Command | What it does | Reference |
|---------|-------------|-----------|
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

| Command | What it does | Reference |
|---------|-------------|-----------|
| `/aeo products` (or `/aeo product list`) | List the product catalog (IDs + image status) used as swap sources | [image-thumbnails.md](references/image-thumbnails.md) |
| `/aeo product add --pdp <url>` | Add a product by PDP URL (scrapes title/image/price) | [image-thumbnails.md](references/image-thumbnails.md) |
| `/aeo products discover` | Crawl the domain sitemap for candidate PDP URLs to add (read-only; flags URLs already in the catalog) | [image-thumbnails.md](references/image-thumbnails.md) |
| `/aeo products rescan` | Re-scrape every product PDP (up to 30) to backfill/refresh media, title, image, price | [image-thumbnails.md](references/image-thumbnails.md) |
| `/aeo image search <query>` | Search Pexels for reference scenes (--per-page, --page) | [image-thumbnails.md](references/image-thumbnails.md) |
| `/aeo image swap --content <id> --product <id> --reference <url>` | Generate a thumbnail by swapping a product into a reference scene (5 credits) | [image-thumbnails.md](references/image-thumbnails.md) |
| `/aeo image upload --file <path>` | Upload a local image (≤25 MP) to the thumbnail bucket (--content to pin) | [image-thumbnails.md](references/image-thumbnails.md) |
| `/aeo image generate --prompt <text>` | Generate image(s) from a text prompt for thumbnails/gallery (credit-metered; cost scales with model and `--sweep` count). `--model nano-banana-pro\|gpt-image-2\|grok-image`, `--sweep N` (1-8 candidates), `--aspect`, `--resolution`, `--ref`, `--brand-style`. Async — returns job IDs. | [image-thumbnails.md](references/image-thumbnails.md) |
| `/aeo image poll <jobId...>` | Check status + result URLs of image generation jobs | [image-thumbnails.md](references/image-thumbnails.md) |
| `/aeo image gallery [--type image\|video] [--limit N]` | List generated gallery assets for the domain (default `image`, limit 20) | [image-thumbnails.md](references/image-thumbnails.md) |
| `/aeo image gallery delete <assetId>` | Delete a generated gallery asset by ID | [image-thumbnails.md](references/image-thumbnails.md) |

### aeo feedback — Send feedback to the team

| Command | What it does | Reference |
|---------|-------------|-----------|
| `/aeo feedback [<message>]` | Send feedback (bug, idea, anything) to the Aeolo team; bare form opens $EDITOR (requires $EDITOR/$VISUAL set) | this file |

### aeo drive — Google Drive files

| Command | What it does | Reference |
|---------|-------------|-----------|
| `/aeo drive list` | List files in connected Google Drive folder (--folder) | [drive.md](references/drive.md) |
| `/aeo drive read <file_id>` | Read a file from Google Drive | [drive.md](references/drive.md) |
| `/aeo drive download <file_id>` | Stream raw bytes to disk (pptx, large/binary files); `-o <path>` to set output | [drive.md](references/drive.md) |

> **Supported types**: Google Docs/Sheets, txt/json/md/csv, **PDF**, **XLSX/XLS** (all sheets, 200-row cap each), **DOCX**, images (≤5MB base64). Not supported: `.doc`, `.pptx`, `.pages`, `.numbers`, `.key` — see [drive.md](references/drive.md).

### aeo integrations google — GA4 / GSC selection

| Command | What it does | Reference |
|---------|-------------|-----------|
| `/aeo integrations google properties` | List GA4 properties accessible to the connected Google account (for selection) | [metrics.md](references/metrics.md) |
| `/aeo integrations google sites` | List Search Console sites accessible to the connected Google account (for selection) | [metrics.md](references/metrics.md) |
| `/aeo integrations google set [--ga4-property <id>] [--gsc-site <url>]` | Select the GA4 property and/or GSC site for this domain (powers `measure overview`/`measure traffic`) | [metrics.md](references/metrics.md) |

> Connecting Google is a browser sign-in (**Settings → domain integrations**), not a CLI command. If Google isn't connected these return an actionable pointer to Settings — connect there first, then pick a property/site.

### aeo agency — Execution-agency matching waitlist

| Command | What it does | Reference |
|---------|-------------|-----------|
| `/aeo agency request --types seo,pr,blog,reddit,youtube,influencer [--other "..."]` | Join the execution-agency matching waitlist (Earned Media early access). Matching isn't live yet; this records interest and pings the team. Idempotent per (domain, type). | this file |

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

### aeo account — Account & subscription

> `account` is the canonical noun. `billing` and `whoami` are aliases: `aeo billing subscription` = `aeo account subscription`, `aeo billing credits` = `aeo account credits`, `aeo billing ledger` = `aeo account ledger`, and bare `aeo whoami` = `aeo account whoami`.

| Command | What it does | Reference |
|---------|-------------|-----------|
| `/aeo account whoami` (or `/aeo whoami`) | Show current user (email, tier, trial days) | this file |
| `/aeo account subscription` | Show current subscription, tier, and credit summary | this file |
| `/aeo account credits` | Show current credit balance | this file |
| `/aeo account ledger` | Show credit ledger entries (`--days` default 30, `--limit` default 50) | this file |

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
| **Daily Content** | Every day (cron or manual) | Pick topic from priority queue → write → deploy → request indexing (channel-post distribution is retired) |
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

Applies to: visibility check run, content generate, content import, content update, content deploy, content redeploy, content unpublish, content delete, content job cancel, carousel delete, audit run, reference analyze, reference delete, video analyze, image swap, image generate, image upload, image gallery delete, brand update, strategy update, strategy visual update, config data-sources update, automation schedule set, topics create, topics update, topics archive, topics restore, topics assign-prompts, prompts add, prompts update, prompts delete, prompts generate, report snapshot, post analyze, post delete, channel add, channel update, channel delete, channel connect, channel disconnect, channel indexing, product add, products rescan, domain rescan, agency request, integrations google set.

Never call a write API without confirmation. Always show what you're about to do and ask "Proceed?" first. Be extra explicit for the irreversible ones: `content deploy`/`content redeploy`/`content unpublish` change the customer's live site.

Retired writes (`carousel create`, `carousel update`, `post import`, `post preview`, `post approve`, `post publish`, `video generate`) are **not** on this list because confirmation is not the gate — do not offer or call them at all. See [Retired surfaces](#retired-surfaces--tbd-2026-07-27).

## Communication Rules

- **UUID is internal only.** User-facing messages must use `title`, `name`, `domain`, `canonical`, etc. Example: `"bc2ef290-..." updated` → `"Best Project Management Tools for Startups" updated`
- **Agent workflows** (`/aeo topics suggest <domain_id>`, `/aeo content idea <domain_id>`, `/aeo prompts audit <domain_id>`, and manual article drafting → `aeo content import`): These have no bare CLI verb — they require LLM reasoning in the external agent. `topics suggest` proposes a durable Topic architecture without saving it; `content idea` recommends one next article; `prompts audit` judges validity and utility without changing Prompts. Drafting workflows write directly and import. `aeo content generate` is only for explicit server-side generation jobs and spends production credits. (`aeo content review <id>` is a real wired command, not a workflow.) The `post write` → `post import` workflow is **retired** — see [Retired surfaces](#retired-surfaces--tbd-2026-07-27).
- **Explicit verbs required**: `aeo content list`, `aeo visibility show`, `aeo strategy show`, etc. Running `aeo <command>` without a verb shows sub-help. Exception: `aeo content --limit 5` (bare flags = implicit list).

Before writing or generating any content (manual draft/import or explicit `/aeo content generate`), always read [geo-strategy.md](references/geo-strategy.md) and [strategy.md](references/strategy.md) first.

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

curl -fsSL https://skills.aeolo.io | sh

After install, verify: `aeo --version`

## Update

aeo update              # self-update to latest
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

Types: `shopify`, `vercel`, `linkedin`, `threads`, `reddit`, `instagram`, `x`, `website`, `custom`

**`--type custom` — connect the customer's own site as a Content Feed pull-channel.**
Use this when they publish on a site we don't push to (their own blog, a static or
hand-built site). It provisions the feed in one step:

```bash
aeo channel add --type custom --url https://yourdomain.com/blog
```

- mints a read-scoped API key (returned once) to pull the feed,
- returns the authed feed URL with `?base=` so canonical points at their domain,
- returns an ownership verify `<meta>` tag to paste into their site `<head>`.

Then wire the feed **server-side** per the guide it links (render `content_html`, inline
`_geo.schema_jsonld`, set the page canonical to `_geo.canonical`). A client-only
widget defeats GEO. See also `aeo content feed` for the raw feed URLs + JSON Feed contract.

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
  --command "visibility check run --engines=chatgpt,gemini" \
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
