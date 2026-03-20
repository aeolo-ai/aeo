# Content Management

---

## /aeo content propose — Generate content proposals from visibility gaps

```bash
aeo content propose
```

Response: `{ "success": true, "data": { "items": [...] } }` — array of proposals, each with `id`, `title`, `articleType`, `priority`, `rationale`, `targetKeywords`.

Present as a table:

```
## Content Proposals — {N} new items

| # | Title | Type | Priority | Rationale |
|---|-------|------|----------|-----------|
| 1 | ...   | blog | high     | ...       |
...

Proposals saved to content history (status=proposed).
```

Then ask if the user wants to write any of them now (`/aeo content write`). Or suggest running `/aeo strategy update` to update the Priority Queue in the manifest.

---

## /aeo content list — List content items

```bash
aeo content [--status=draft|review|published]
```

Optional `status` filter. Response: `text/markdown` — table with `id`, title, status, type, language, created date.

---

## /aeo content get <id> — Read full article content

```bash
aeo content get <id>
```

Response: `text/markdown` — full article content. Use this to review a draft before updating or deploying.

---

## /aeo content update <id> — Update a content item

```bash
aeo content update <id> --status=review --title="New Title"
```

Accepted fields (all optional, partial update).

> **CLI vs API:** The `aeo` CLI supports `--status`, `--deploy-status`, and `--title` flags. The `patches` and `content` fields below are **API-only** — use direct HTTP calls or agent-level API calls for body editing.

| Field | Type | Notes |
|-------|------|-------|
| `title` | string | Article title |
| `meta_description` | string (max 320) | SEO meta description |
| `status` | `draft` \| `review` \| `published` \| `archived` | Workflow status |
| `target_keywords` | string[] | Update keyword targets |
| `content` | string | Full article body replacement (markdown) |
| `schema_types` | `string[]` | Override schema types (`Article`, `BlogPosting`, `FAQPage`, `HowTo`, `ItemList`) |
| `patches` | `{search, replace}[]` (max 20) | Targeted edits — preferred over full replacement |

### Body editing — patches vs full replace

**Use `patches` (preferred):** Send only the changed sections. First match of each `search` string is replaced in order.

```json
{
  "patches": [
    {
      "search": "## Why SPF Sticks Work for Padel\n\nPadel is a fast-growing sport",
      "replace": "## Why SPF Sticks Are Essential for Padel\n\nPadel is one of the fastest-growing sports"
    }
  ]
}
```

Response includes `patch_result: { applied: [0], failed: [] }`. If a patch index appears in `failed`, the search string wasn't found — read the current body first with `/aeo content get <id>` and retry with the correct string.

**Use `content` (full replace):** When rewriting large portions or restructuring. Always read the current body first.

---

## /aeo content preview <id> — Preview in browser

```bash
aeo content preview <id>
```

Generates a preview link and automatically opens it in the browser. Use the `--no-open` flag to output the link only.

Response: `{ "success": true, "data": { "content_id": "...", "title": "...", "preview_url": "https://tryaeolo.com/preview/...", "share_token": "..." } }`

Idempotent — calling multiple times returns the same link. If the browser cannot be opened in the current environment, output the URL and inform the user.

---

## /aeo content deploy <id> — Deploy to Shopify

```bash
aeo content deploy <id>
```

`channel_id` is optional — if omitted, the first active Shopify integration for the domain is used automatically.

Response: `{ "success": true, "data": { "shopifyArticleId": "...", "blogId": "...", "url": "..." } }`

After success, show the published URL. If deploy fails because no Shopify channel is connected, direct the user to Aeolo dashboard → Integrations.

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
