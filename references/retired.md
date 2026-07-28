# Retired Surfaces (TBD 2026-07-27)

Three output channels are sunset. **Do not offer them to the user, and do not call
their output commands.** Route every content request to a blog article
(`content generate` or draft directly + `content import` → `publish deploy`).

| Retired | Commands | Instead |
|---------|----------|---------|
| Card news | `carousel create`, `carousel update` | blog article |
| Channel posts / Threads | `post write`, `post import`, `post preview`, `post approve`, `post publish` | blog article |
| Shortform video output | `video generate`, `video poll` | blog article |

The loop never closed for any of the three: Threads publishing was gated on a Meta
app review that never landed, and reels drive Google SERP discovery but not chatbot
citations. Blog articles are the only channel where diagnose → write → deploy →
index → citation closes inside the product.

> **Only output generation is retired — reference input is fully live.** Analyzing
> someone else's Reel, TikTok, or post as *reference* to brief a blog article is
> unaffected: `diagnose references analyze`, `diagnose video analyze`,
> `reference style`, `channel voice`, and `post analyze` all stay. So do reads and
> cleanup on rows that already exist (`carousel list`/`get`/`delete`,
> `post list`/`get`/`delete`).

## Card news (carousel) — RETIRED

Generation is closed server-side (HTTP 410, no credits charged); the read/cleanup
verbs stay so existing decks remain visible and removable.

| Command | What it does |
|---------|-------------|
| `/aeo carousel list [--limit N]` | List carousel decks that already exist for the domain |
| `/aeo carousel get <jobId>` | Get an existing deck (slides + caption) by job ID |
| `/aeo carousel create [--topic "..."]` | **RETIRED** — returns 410. Write a blog article instead. |
| `/aeo carousel update <jobId> --caption "..."` | **RETIRED** — do not edit card news captions |
| `/aeo carousel delete <jobId>` | Soft-delete a deck (cancels active work; assets stay recoverable) |

## Channel posts (social distribution) — RETIRED (output only)

Threads led this surface and its publishing was gated on a Meta app review that
never landed, so the loop never closed. The `post` group is platform-agnostic, so
the retirement covers every platform it served (threads / linkedin / reddit /
instagram / x). Do not draft, import, approve, or publish channel posts.

`post analyze` is **not** retired: it reads an owned or competitor post as reference
voice evidence for an article brief. `post list`/`get`/`delete` stay for reading and
cleaning up rows that already exist.

| Command | What it does |
|---------|-------------|
| `/aeo post analyze --url <URL>` | **LIVE** (reference input) — analyze one channel/reference URL for voice evidence |
| `post write` | **RETIRED** — draft a blog article instead |
| `/aeo post import` | **RETIRED** — do not import new channel-post drafts |
| `/aeo post preview <id>` | **RETIRED** — with the post pipeline |
| `/aeo post approve <id>` | **RETIRED** — with the post pipeline |
| `/aeo post publish <id>` | **RETIRED** — publishing to social platforms is closed |
| `/aeo post list` / `/aeo post get <id>` / `/aeo post delete <id>` | Read/cleanup on existing rows |

## Shortform video output — RETIRED

Reels drive Google SERP discovery, not chatbot citations. Only generation is closed;
every command that *reads* an external Reel/TikTok to brief an article stays live
(see `reference analyze`, `video analyze` in the main command reference).

| Command | What it does |
|---------|-------------|
| `/aeo video generate --prompt <text>` | **RETIRED 2026-07-27** — brief a blog article instead |
| `/aeo video poll <jobId...>` | **RETIRED** — only reaches jobs created before the sunset |
