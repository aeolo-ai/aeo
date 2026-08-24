# Content Management

---

## Content Creation — Rationale-Driven

When creating content (via `/aeo content generate` or `/aeo content import`), always include a **rationale** explaining why this content should exist:
- Which visibility gaps does it address?
- Which prompts/keywords is it targeting?
- What stage of the customer journey does it serve?

The rationale is stored in `content_history.rationale` and helps prioritize content in the pipeline. Discover gaps through conversation (`aeo visibility show`, `aeo domain audit`) and let the rationale emerge naturally from the analysis.

---

> The full `aeo content` verb set is `list | feed | get <id> | review <id> | import | generate | jobs | update <id> | preview <id> | deploy <id> | redeploy <id>`. This file covers list/feed/get/update/preview/deploy/redeploy; `content generate`/`content jobs`/`content import` live in [content-create.md](content-create.md) and `content review` in [content-review.md](content-review.md).

## /aeo content list — List content items

```bash
aeo content [--status=draft|review|published] [--limit=N] [--offset=N]
```

Optional filters:
- `--status` — filter by status (comma-separated OK: `--status=draft,review`)
- `--limit` — max items to return (default: 20, max: 200)
- `--offset` — skip first N items (default: 0, for pagination)

Response: `text/markdown` — table with `id`, title, status, type, words, keywords, meta description, date. Includes total count (e.g., "Showing 20 of 61"). Use `--limit=20 --offset=20` for page 2.

---

## /aeo content feed — Content Feed for your own site

```bash
aeo content feed
```

Returns your published articles as a delivery contract you can render on the customer's **own domain** (Tier A — the GEO-correct owned-media path). Response: `text/markdown` with:

- **Feed URL** — one route, no id in the path: `GET /v1/connector/feed.json`. The key names the channel it may read, so an unbound key is refused (`TOKEN_NOT_CHANNEL_BOUND`) rather than served the wrong slice. Append `?base=https://yourdomain.com/blog` — canonical is built as `<base>/<slug>`, so the base must be the path the articles actually live at. Connecting the channel mints a bound key; more via `POST /domains/:id/api-keys` with `channel_id`.
- **Item shape** — standard **JSON Feed 1.1** + `_geo` extension: `content_html` (server-rendered body), `title`/`summary`/`image`/`date_published`/`tags`, `_geo.slug`, `_geo.schema_jsonld` (inline-ready schema.org Article), `_geo.canonical`.
- **How to consume** — fetch server-side (SSR/SSG), render `content_html` in your layout, inline the JSON-LD, set the page canonical to your own URL. A client-only widget defeats GEO; the body must be in server HTML on your origin. Incremental pulls: `?limit=&since=`.

Use this when the customer wants Aeolo articles on their existing site in their own design rather than the hosted `*.aeolo.blog` blog or a Shopify deploy.

---

## /aeo content get <id> — Read article content

```bash
aeo content get <id>              # the whole body
aeo content get <id> --head       # metadata + first ~600 chars — a state check
aeo content get <id> --blocks     # the article as addressable blocks
aeo content get <id> --block b3   # one block, verbatim
```

Response: `text/markdown`.

**Read the least you need.** The default returns the whole article, and on a 6,000-character body that is thousands of characters of context spent on text you will not use.

- `--head` when you only need to identify or triage an article, or to check whether an edit saved.
- `--blocks` lists the article as `b1`, `b2`, … with a one-line preview each. Tables stay whole, so a comparison table is one block rather than a pile of fragments.
- `--block <id>` returns that block's source verbatim. **That source is exactly what a `--patch` anchor needs** — read the block, patch against what it returned.
- No flag when you genuinely need the full body (a sweeping rewrite, or a read where you don't yet know which part matters).

Block ids are positional and derived per read, so they are valid for the article as it was when you listed them. List, then act — don't cache ids across turns where the article may have changed.

---

## /aeo content update <id> — Update a content item

```bash
# Metadata update
aeo content update <id> --status=review --title="New Title"

# Targeted body edit (preferred — surgical)
# Shows the before/after and asks before writing. Pass --yes to skip the prompt
# (scripted use); the prompt is skipped automatically when stdin isn't a terminal.
aeo content update <id> --patch "old sentence>>>new sentence"

# Full body replace (from a file)
aeo content update <id> --body-file ./draft.md

# SEO fields
aeo content update <id> --meta-description="Updated description" --keywords="seo,geo,brand"

# Thumbnail
aeo content update <id> --thumbnail-url https://example.com/og.png
aeo content update <id> --clear-thumbnail
```

All flags are optional — send only what you want to change.

| Field | CLI Flag | Type | Notes |
|-------|----------|------|-------|
| `title` | `--title` | string | Article title |
| `meta_description` | `--meta-description` | string (max 320) | SEO meta description |
| `status` | `--status` | `draft` \| `review` \| `published` \| `archived` | Workflow status |
| `deploy_status` | `--deploy-status` | string | Deployment status |
| `target_keywords` | `--keywords` | string[] | Comma-separated: `"seo,geo,brand"` |
| body | `--body` / `--body-file` | string / file path | Full body replacement (`--body` is the legacy `--content` alias) |
| body patch | `--patch` | `"search>>>replace"` | Targeted edit — replaces the first match without resending the whole body |
| thumbnail | `--thumbnail-url` | url | Pin an external thumbnail directly |
| thumbnail | `--clear-thumbnail` | flag | Drop the existing thumbnail |

### Body editing workflow

For small changes prefer `--patch "search>>>replace"` — it edits in place without resending the whole article. For larger rewrites, read the current body, edit it locally, and send the full replacement via `--body-file`:

```bash
# 1. Read current content
aeo content get <id>

# 2a. Surgical edit
aeo content update <id> --patch "outdated stat>>>updated stat"

# 2b. Or full replacement from a file
aeo content update <id> --body-file ./revised-draft.md
```

---

## /aeo content preview <id> — Generate a preview link

```bash
aeo content preview <id>
```

Generates a shareable preview link and prints it. It does not open a browser (there is no `--no-open` flag) — surface the URL to the user yourself.

Response: `{ "success": true, "data": { "content_id": "...", "title": "...", "preview_url": "https://aeolo.io/preview/{share_token}", "share_token": "..." } }`

Idempotent — calling multiple times returns the same link.

---

## /aeo content deploy <id> — Deploy to a publish destination

```bash
aeo content deploy <id>                     # destination picked as described below
aeo content deploy <id> --target wordpress
```

### Targets

`--target` takes one of five values; anything else is refused by name.

| Target | Destination | Channel required |
|--------|-------------|------------------|
| `shopify` | the connected Shopify blog | yes — auto-resolved when the brand has exactly one store |
| `blog` | the hosted Aeolo blog | no — always available |
| `wordpress` | a connected WordPress site | yes |
| `cafe24` | a connected Cafe24 mall board | yes |
| `pangolingo` | a connected Pango Lingo newsroom — ships the article **and every locale edition as one unit** | yes |

How the destination is decided: a named `channel_id` decides it from that channel's own `type`, and a `--target` that disagrees is refused naming both rather than one silently winning. Only when neither a channel nor a `--target` is given does it fall back to `shopify`.

`channel_id` is optional — for Shopify, the domain's single active integration is used automatically.

### Deploy gates

Both apply to **every** target, and both refuse before anything is written.

1. **Approved status** — Shopify requires status exactly `approved`; the other targets accept `approved` or `published`. Otherwise: "Only approved articles can be deployed." Fix with `aeo content update <id> --status approved`.
2. **SERP metadata** — the title must be non-empty and ≤ 60 chars, and the meta description must be present and 50–160 chars. Failures return `INVALID_TITLE_LENGTH` / `INVALID_META_DESCRIPTION` as a `**Deploy blocked — …**` message, which the connector API answers with **HTTP 422 `DEPLOY_BLOCKED`**. Fix with `aeo content update <id> --title "…"` / `--meta-description "…"`, then retry deploy.

Response: `{ "success": true, "data": { "shopifyArticleId": "...", "blogId": "...", "url": "..." } }`

After success, show the published URL. If deploy fails because the target's channel is not connected, direct the user to Aeolo dashboard → Owned Media.

**Post-deploy hint**: After a successful deploy, suggest GSC indexing:
> "배포 완료! 빠른 색인을 위해 `/aeo gsc index`로 GSC 색인 요청도 진행할까요? (브라우저 자동화 필요)"

---

## /aeo content redeploy <id> — Push the current article back in place

```bash
aeo content redeploy <id>
```

Updates the body, title, tags, and schema of an **already-deployed** article in-place. The URL handle is preserved — no need to delete and recreate.

Use when:
- Article content was edited after initial deploy
- A bug was fixed in the rendering pipeline (e.g., markdown → HTML conversion)
- Schema or metadata needs updating

**Destination is auto-detected**, read off `deploy_metadata` in this order: `shopify_article_id` → `wp_post_id` → `cafe24_article_no` → `pangolingo_article_id`, then `blog` if the article is `published`. Nothing detected means it was never deployed — run `content deploy` first. The `aeo` CLI sends no `--target` on redeploy (auto-detect only); the agent/MCP surface accepts one.

Target-specific notes:
- `blog` — the hosted blog renders straight from the article, so there is nothing to push; the edit is already live once the page cache rolls over.
- `pangolingo` — the Korean original body is create-only on their API, so only the locale editions are pushed.

**Redeploy is not the deploy gate.** There is no approved-status check, and only the **Shopify** path revalidates title/meta — WordPress, Cafe24, and Pango Lingo redeploy without rechecking. When the Shopify path does refuse, the message opens `**Redeploy blocked — …**`, which the connector route does **not** map to 422: it answers HTTP 200 `{"success": true}`. Read the returned message, not the status code, before reporting a redeploy as live.

Response: `{ "success": true, "data": { "shopify_article_id": ..., "handle": "...", "published_url": "..." } }`
