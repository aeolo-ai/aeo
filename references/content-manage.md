# Content Management

---

## Content Creation — Rationale-Driven

When creating content (via `/aeo content generate` or `/aeo content import`), always include a **rationale** explaining why this content should exist:
- Which visibility gaps does it address?
- Which prompts/keywords is it targeting?
- What stage of the customer journey does it serve?

The rationale is stored in `content_history.rationale` and helps prioritize content in the pipeline. Discover gaps through conversation (`aeo visibility show`, `aeo domain audit`) and let the rationale emerge naturally from the analysis.

---

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

## /aeo content get <id> — Read full article content

```bash
aeo content get <id>
```

Response: `text/markdown` — full article content. Use this to review a draft before updating or deploying.

---

## /aeo content update <id> — Update a content item

```bash
# Metadata update
aeo content update <id> --status=review --title="New Title"

# Full body replace
aeo content update <id> --content "# New Article\n\nFull markdown content..."

# SEO fields
aeo content update <id> --meta-description="Updated description" --keywords="seo,geo,brand"
```

All flags are optional — send only what you want to change.

| Field | CLI Flag | Type | Notes |
|-------|----------|------|-------|
| `title` | `--title` | string | Article title |
| `meta_description` | `--meta-description` | string (max 320) | SEO meta description |
| `status` | `--status` | `draft` \| `review` \| `published` \| `archived` | Workflow status |
| `deploy_status` | `--deploy-status` | string | Deployment status |
| `target_keywords` | `--keywords` | string[] | Comma-separated: `"seo,geo,brand"` |
| `content` | `--content` / `--body` | string | Full body replacement |

### Body editing workflow

Always read the current body first, modify it, then send the full replacement:

```bash
# 1. Read current content
aeo content get <id>

# 2. Send updated body
aeo content update <id> --content "# Updated Title\n\nRevised full markdown..."
```

---

## /aeo content preview <id> — Preview in browser

```bash
aeo content preview <id>
```

Generates a preview link and automatically opens it in the browser. Use the `--no-open` flag to output the link only.

Response: `{ "success": true, "data": { "content_id": "...", "title": "...", "preview_url": "https://tryaeolo.com/preview/{share_token}", "share_token": "..." } }`

Idempotent — calling multiple times returns the same link. If the browser cannot be opened in the current environment, output the URL and inform the user.

---

## /aeo content deploy <id> — Deploy to Shopify

```bash
aeo content deploy <id>
```

`channel_id` is optional — if omitted, the first active Shopify integration for the domain is used automatically.

Response: `{ "success": true, "data": { "shopifyArticleId": "...", "blogId": "...", "url": "..." } }`

After success, show the published URL. If deploy fails because no Shopify channel is connected, direct the user to Aeolo dashboard → Integrations.

**Post-deploy hint**: After a successful deploy, suggest GSC indexing:
> "배포 완료! 빠른 색인을 위해 `/aeo gsc index`로 GSC 색인 요청도 진행할까요? (브라우저 자동화 필요)"

---

## /aeo content redeploy <id> — Update existing Shopify article

```bash
aeo content redeploy <id>
```

Updates the body, title, tags, and schema of an **already-deployed** Shopify article in-place. The URL handle is preserved — no need to delete and recreate.

Use when:
- Article content was edited after initial deploy
- A bug was fixed in the rendering pipeline (e.g., markdown → HTML conversion)
- Schema or metadata needs updating

Prerequisite: the article must have been deployed at least once (`deploy_metadata.shopify_article_id` exists).

Response: `{ "success": true, "data": { "shopify_article_id": ..., "handle": "...", "published_url": "..." } }`
