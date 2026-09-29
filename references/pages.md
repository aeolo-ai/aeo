# Pages — product detail pages and DOM review

Pages are a separate HTML workspace, not article content. Use `aeo pages` for PDPs, product comparisons, ingredients and routines. Do not import PDP HTML as an article or use `site template edit` to modify a page: the latter changes captured site chrome.

Access follows the **domain owner's active Custom plan** and domain membership. Your personal plan is irrelevant. Members can read and comment. Owners/editors can create, edit, translate, change review state, ask the comment agent and accept its proposal. Viewers can reply and resolve/reopen/delete their own threads. Staff retain support access. Private edits do not publish or request indexing.

## Scope and data

Start with `aeo pages access --domain <domainId>` and `aeo pages channels --domain <domainId>`. Use returned channel UUID as `--channel-id`. Every page operation uses that domain/channel pair; every language edition has a separate page ID and revision.

CLI accepts `--input-file request.json` (recommended for HTML/quotes/newlines) or `--input-json '<JSON>'`. MCP/chat uses the same `pages … --input-json` commands; never put a local filename in an MCP request. Responses are JSON with a `data` envelope. Errors start with an explicit code. Stale revision/version is a conflict: re-read and reassess rather than overwriting.

## Read and edit

- `pages list [--trash true]`, `pages get <pageId>`
- `pages elements <pageId>` returns exact DOM selections with domainId/channelId/pageId/locale/revision/start/startEnd/closeStart/end. Pick one returned element; never fabricate offsets.
- `pages inspect <pageId> --revision N` returns the inert preview HTML.
- `pages create --input-file page.json`: PageDraftInput plus an idempotency UUID `id`.
- `pages update <pageId> --input-file page.json`: full PageDraftInput plus `expectedRevision` from get. Preserve the existing representation and locale.
- `pages preview --input-file page.json`: renders PageDraftInput without saving.
- `pages edit <pageId> --input-file edit.json`: `{ "selection": <one elements result>, "patches": [{"find":"exact old HTML/text","replace":"new HTML/text"}] }`. Only the selected DOM can change; header/footer/navigation remain protected. Read the stored HTML first. A successful edit consumes the selection and advances revision; reselect for further changes.

PageDraftInput fields: `channelId`, `title`, `slug` (no extension), `kind` (`product|comparison|routine|ingredient|collection`), `destination` (from channels), `locale`, `sourceUrls`, `bodyHtml`. A designed page uses `designPreview: {version:1,path:"/products/product-name",html:"<!doctype html>…"}` and empty bodyHtml. Preserve product identity across translations: language is not a different regional SKU. `productPage` is a separate structured representation; do not combine it with designPreview. See `pages get` for the actual document.

## Languages and review

- `pages editions <sourcePageId>` with `{ "expectedRevision": N, "locale": "ko" }` creates an empty edition. English is the source language.
- `pages builds translate <targetPageId>` with `{ "expectedRevision": N }` starts its translation. Poll `pages builds edition <targetPageId>` or `pages builds get <jobId>`. For a failed job, an explicit retry passes its `retryOf` ID.
- `pages review <pageId>` with `{ "expectedRevision": N, "state": "draft|in_review|reviewed", "sourceFingerprint": "…" }`. Non-English review requires the current English fingerprint. Review does not publish.
- `pages trash <pageId>` / `pages restore <pageId>` with `{ "expectedRevision": N }` are recoverable; no permanent deletion command.

## DOM comments

All commands require `--page-id <pageId>` and `--channel-id <channelId>`.

- `pages comments list`
- `pages comments create` with `{ "selection": <one elements result>, "body": "Requested change" }`
- `pages comments update --thread-id <id>` with `{ "version": N, "action": "reply", "body": "…" }`; also `resolve`, `reopen`, `delete` or `reanchor` (the latter takes a fresh `selection` and needs editor rights).
- `pages comments agent --thread-id <id>` with `{ "version": N }` generates a stored proposal. It may take up to two minutes. A transport error does not prove failure: list the thread and inspect `work.state` before retrying. Working/failed and proposalRevision are durable; expired work is retryable.
- `pages comments apply --thread-id <id>` with `{ "version": N }` accepts the stored proposal, checks page and thread revision atomically and resolves the thread. Do this only after the user asks to apply the proposed change. Orphaned or stale proposals require a fresh selection/proposal.

## Generate from product evidence

- `pages research list|get <jobId>|start --product-id <productId>`. Start input: `{ "id": "<UUID>", "locale": "en" }`.
- `pages builds list|get <jobId>`
- `pages builds analyze` input: `{ "id": "<UUID>", "sourceUrl": "https://…", "locale": "en", "targetQuestion": "…", "productId": "…", "researchJobId": "…" }` (last three optional).
- `pages builds generate <analysisJobId>` input: `{ "id": "<UUID>", "questionId": "<returned question ID>" }`.
- `pages builds apply <generationJobId>` saves the completed result as a private draft.

Generation/research/translation invoke server AI work: use only when requested. Do not repeat a pending job. Reuse the same creation UUID for transport retries, poll the existing job, and inspect its evidence/claims before requesting review. Neither completion nor reviewed status proves publication, indexing or AI citation.
