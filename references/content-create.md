# Content Creation

This document owns the interactive workflow and output contract, not a second quality checklist.

Default draft directory: `.aeo/` in the working directory. Add to `.gitignore`.
File naming: `{AEOLO_DOMAIN_ID}_{slug}.md`. See Step 5 for path resolution.

---

## /aeo content generate — Server-side content generation

Starts a credit-metered background job. Use this when the user explicitly wants
Aeolo to generate the article server-side; confirm topic, keywords, language,
article type, and any task-specific reference/tone inputs first. Track progress
with `/aeo content jobs`.

**Writing requires a destination (AEO-531).** The destination decides wording,
locale set, and reference policy, so it cannot be picked after the text exists:

- Pass `--target-channel <channelId>` (ids from `aeo channels list`). If the
  domain has exactly ONE publish surface, omitting it is fine — that surface is
  adopted automatically.
- **Several channels + no flag → refused** (`CHANNEL_CHOICE_REQUIRED`). Ask the
  user which surface the article is for and retry with `--target-channel`. Do
  not guess and do not retry blind — the refusal costs nothing (it fires before
  credits are reserved).
- **No publish surface at all → refused** (`NO_ACTIVE_CHANNEL`). The fix is a
  channel: guide the user to connect one (or activate the hosted blog — an
  explicit user action in the dashboard). Never work around this by importing a
  channel-less draft instead; that just moves the dead end to deploy time.

The background job must treat reference material, voice examples, and broad
tone notes as task context. Do not silently update brand memory from a
generation pass.

---

## Manual Draft/Import Path — GEO-optimized article

This is the default external-agent writing path. Draft the article locally in the agent reasoning loop; after writing, save the draft and import via `aeo content import`.

Use `aeo content generate` only when the user explicitly wants Aeolo to run a server-side paid generation job.

> **Destination note:** pass the destination on import: `--target-channel <id>`
> (ids from `aeo channels list`), plus `--folder /path` or `--folder-id <id>` on
> managed sites. Know where the article will be published BEFORE you draft it;
> the destination decides wording and locales. Before go-live you can
> re-address it with `aeo content update <id> --target-channel <id>`; once the
> article is live that is refused (`DESTINATION_LOCKED`).

---

### Pre-flight: Article Brief

**Mandatory gate before proceeding to the outline.** Regardless of the entry path, this checklist must be filled.

#### Trigger Paths (agent determines automatically)

| Path | Trigger | Already Available | Additionally Needed |
|------|---------|-------------------|---------------------|
| **Gap-driven** | `aeo visibility show` gap + content strategy priority | Brand context, target prompts, gap context | Competitor facts |
| **Client-brief** | User provides client materials/meeting notes/appeal points | Angle, materials (raw state) | Claims extraction, competitor context |
| **Prompt-targeted** | "I want us to show up for this prompt" | Target prompt | Brand claims, competitors |
| **Content Refresh** | Changed information or an unanswered reader need in an existing article | Existing article structure + content | Latest data, updated competitor context |

For Content Refresh, read the existing article via `/aeo content get <id>` and confirm whether the authorized task is an update or a distinct new article. Do not create a competing original merely because time elapsed. Changing only `dateModified` is not a substantive update.

#### Article Brief Checklist

```
## Article Brief

### Trigger
- Starting point: [gap / client brief / prompt target / content refresh]
- Related gap ID, prompt, or existing content ID: [fill in if available]

### Brand Ammunition
- [ ] Relevant brand claims: claim + proof point + source, without a quota
- [ ] Competitor Context: facts about comparison targets (required/optional depending on articleType)

### 1st-Party Experience (authenticity material)
- [ ] Internal tests/usage experience: who, how many times, under what conditions? (be specific)
- [ ] Customer testimonials/cases: is there actual user feedback?
- [ ] Domain-specific context: are there unique facts specific to this topic? (no generic statements)
- [ ] Positioning: will the article be transparent about being written by the brand?

> Use first-party experiences only when supplied, relevant and authorized. Their absence does not prevent a supported educational or product explanation. Ask for missing evidence only when it is necessary to answer the commissioned question.

### Article Direction
- [ ] Topic / H1 candidate
- [ ] articleType
- [ ] Target keywords
- [ ] Target engine(s)
```

If Brand Ammunition is empty:
- Brand Claims → Extract from `/aeo agent context` + session documents, then convert to Ammunition format
- Competitor Context → Web research or request from user

> **Authority Sources are not a Pre-flight target.** External third-party data (statistics, studies, etc.) for article credibility is gathered in Step 3 (Research).

---

### Step 1 — Collect inputs

Ask the user (or infer from context):
- **Topic** (required)
- **Article type** — choose from format matrix below; `blog` is default
- **Target keywords** (required, 1–20)
- **Language** — `en` (default) | `ko` | `ja` | `zh-Hans` | `zh-Hant` | `es` | `id` | `vi` | `th` | `de` | `fr` | `hi` | `ar` | `it` | `nl` | `pl` | `ru`
- **Length**: follow an explicit request where supported; otherwise supported depth decides (no fixed word count)

If brand context isn't loaded yet, fetch it first (`/aeo agent context`) — brand context shapes the entire article.

### Step 1.5 — Task Tone / Reference Guidance

Do not read deprecated voice-example stores by default. Apply the content
strategy's `Editorial Voice` and language-specific guidance as the default.
When the user selects a task-specific reference, use it to refine that default
without overriding verified brand facts or claim boundaries.

Look for:

- explicit reference-analysis output attached to this task;
- concrete sample copy or URLs supplied by the user;
- language-specific `Editorial Voice` guidance in the content strategy;
- broad tone notes inside `brand_context`, only as fallback context.

#### If task-specific tone/reference evidence exists

Extract in the format below and apply consistently throughout all subsequent steps (Outline heading naming → Writing tone → FAQ style):

```
## Extracted Voice Guidance
- Formality: [formal / balanced / casual]
- Voice characteristics: [extracted adjectives]
- Phrases to use: [example sentences/expressions]
- Phrases to avoid: [prohibited expressions]
- Other notes: [additional style instructions]
```

> Approved channel requirements own voice and output conventions. An analyzed style sample is a reference, not authority for facts or mandatory article structure. Use real markdown section headings; if a FAQ is useful, use a `## FAQ` H2.

After extraction, confirm with the user in one line: _"This article will be written in a [characteristics] tone. Does that sound right?"_

#### If no task-specific tone/reference evidence exists

Ask briefly if tone matters for this task. If the user wants default behavior, proceed with the brand context and content strategy.

```
No selected tone/reference evidence is attached to this task.
If tone matters here, I need one of:

1. a reference URL or analyzed reference result;
2. example copy to match or avoid;
3. a short tone instruction.
```

Do not write product memory directly from this flow. If the answer should become durable brand context, summarize the proposed patch and ask for review.

If no task-specific tone/reference evidence is attached or provided, skip this step and rely on brand context plus content strategy.

### Step 2 — Outline

Build a structured outline:
- H1: a clear reader-facing heading for the commissioned question
- H2/H3: hierarchy where each section is independently quotable
- Identify evidence needed for each section; use tables or quotations only where useful and supported
- Add FAQ only for supported adjacent questions not already answered

Show the outline to the user and confirm before writing.

### Step 3 — Research (MANDATORY GATE)

**Research must be proactive and exhaustive before writing begins.** Read [data-access.md](data-access.md) for the full source configuration.

#### 3.1 — Load data sources

```bash
aeo config data-sources
```

Read the configured sources. If none exist, use the default sources from data-access.md.

#### 3.2 — 1st-Party material sweep (ALWAYS FIRST)

Search for topic-relevant material in this order:

1. **Google Drive**: `aeo drive list` → scan folders mentioned in data_sources config → `aeo drive read <id>` for matches
2. **Custom URLs**: If data_sources lists specific URLs, read/browse them
3. **Brand context**: `aeo agent context` → extract relevant claims and proof points

Extract 1st-party facts: product test results (who, how many, conditions), customer testimonials, internal data.

#### 3.2.1 — Optional inline catalog image

Keep articles text-only by default. Evaluate an inline image only for the manual
draft/import path and only when the article contains real product or catalog
copy. For `aeo content generate`, do not add image markup yourself; the
server-side writing job applies its own deterministic catalog guard.

Qualify one candidate in this order:

1. Read the catalog in `aeo agent context` and capture the exact product title
   plus its verified first-party PDP URL.
2. Verify that the PDP belongs to the live brand domain or its subdomain. The
   image asset itself may live on the brand's commerce CDN; first-party
   provenance comes from the PDP, not the image host.
3. Run `aeo products` and require an exact product-title match with a real Image
   URL. Do not guess that two similarly named products are the same SKU.
4. Require the final body to name that exact product or link its verified
   first-party PDP.

When all checks pass, insert at most one inline image immediately after the
first explanatory paragraph for that product:

```markdown
![Exact product title](verified-image-url)
```

If the draft already contains an image, do not add another. Treat the image as
presentation, never as evidence for a claim or a replacement for an inline
citation. Do not use `aeo image search`, `aeo image generate`, web-search image
results, stock images, marketplace PDPs, or arbitrary external URLs for inline
body images. If any check fails, keep the article text-only.

#### 3.2.2 — Optional rich blocks: YouTube embed · product carousel

Two inline blocks are allowed besides the catalog image, both **verified-only
and optional**. The author supplies data; Aeolo renders the markup on every
surface (hosted blog, Cafe24/Shopify/WordPress/Pango Lingo deploys, feeds,
export) — never hand-write iframes or card HTML, never add `style`/`class` of
your own.

**YouTube embed — at most ONE per article.** When research surfaced a directly
relevant video (an official demo, tutorial, or talk that materially helps the
reader), verify it first:

```bash
aeo source youtube <url>
```

It confirms the video exists and allows embedding, returns the exact embed
snippet with the real title, and best-effort fetches the caption transcript so
the body can discuss what the video actually says. Paste the returned snippet on
its own line at the point where it helps most. Never construct an embed from an
unverified or guessed VIDEO_ID; if verification fails or no clearly relevant
video exists, embed nothing.

**Product carousel — at most ONE per article, up to 3 tokens.** When the body
genuinely discusses specific catalog products or SKUs (a recommendation,
comparison, or routine that names them), get the tokens first:

```bash
aeo source products
```

It lists this domain's catalog products and SKU tokens (`productId`, or
`productId:variantId` for one SKU) together with the exact marker syntax. Place
the marker on its own line right after the section it supports:

```html
<figure class="product-carousel" data-products="TOKEN1,TOKEN2"></figure>
```

Use only tokens the command returned — never invent or edit one — and prefer
the exact SKUs the section discusses. When the article walks through several
catalog products (a routine, a comparison), put them together in the one
carousel (up to 3) rather than featuring only one. When no catalog product
clearly matches the article, add no carousel. Aeolo renders the cards (image ·
name · price · localized CTA linking to the PDP) at publish time; a token that
is not in the domain catalog simply does not render.

#### 3.3 — Ask user for gaps

If 1st-party material is insufficient for the topic:

```
"I found [what you found]. Do you have additional materials on [topic]?
(e.g., test reports, customer stories, internal data)"
```

If the user has nothing, note the gap and proceed — but flag it in Step 4.5 authenticity check.

#### 3.4 — External authority research

After 1st-party is exhausted, gather external sources:
- Primary research, official regulators, and professional institutions for
  scientific, medical, safety, or regulatory claims
- Recent statistics from authoritative sources
- Expert quotes (name + title + affiliation)
- Industry reports and reviews
- Competitor specs (for comparison/ranked_list types)

Commercial explainers, ingredient suppliers, and brand blogs may help discover
terms or primary sources, but they are not sufficient evidence for scientific,
medical, safety, or regulatory claims when an authoritative source is
available. Use reputable secondary media for industry context only within the
scope it directly supports.

Collect `{ name, url, description }` for each source — these become inline citations.

#### 3.5 — Research validation gate

Before proceeding to Step 4:
- [ ] Every claim traces to a source you actually read. Under the `standard` reference policy, import needs ≥3 external source URLs (see [data-access.md](data-access.md))
- [ ] 1st-party material ≥ 1 (if available for this topic)
- [ ] Every claim has a prepared citation
- [ ] Competitor facts verified (for comparison/ranked_list types)

**If validation fails, do not proceed to writing.** Request additional material from the user.

### Step 4 — Write

Write the article to the approved channel requirements. Output rules:
- **Apply the Tone Profile extracted in Step 1.5 consistently across all sections** — including heading naming, sentence length, and tone
- **Title ≠ H1 (HARD RULE)** — These two titles must always be separate:
  - **Title** (`/aeo content import` title parameter): SEO-optimized long version (question-based, keyword-rich)
    Example: `"How to Reapply Sunscreen During a Long Run Without Stopping — A Stick Sunscreen Guide for Runners"`
  - **H1** (first `#` in the body): Short and punchy reader-facing heading
    Example: `"Reapply Sunscreen Mid-Run: The Runner's Guide"`
  - The body must start with `# {H1}`, and the title text must not be repeated anywhere in the body
  - Open in the approved channel voice and answer the reader promptly; no compulsory summary block
  - Do not use the same text for Title and H1
  - Do not repeat the H1 text in the intro
- **Title length ≤ 60 chars (HARD RULE — import enforces this)**:
  - This is Aeolo's current title validation constraint, not a guarantee of search display or performance.
  - `content import` rejects with `INVALID_TITLE_LENGTH` if `title.length > 60`; shorten and retry. At deploy an over-long title is only advice.
  - Good: `"Best No-White-Cast SPF Sticks for Sports (2026)"` (47 chars)
  - Bad: `"What's the Best Sunscreen for NYC Half Marathon Runners in 2026? Race-Morning Picks That Won't Sting, Slip, or Leave a White Cast"` (133 chars — drop everything after the question mark)
- **Search description** — Author it separately from the fuller opening answer. Lead with the answer and preserve the article's distinctive scope. Comparisons need the set/count and criteria; routines need roles, order and conditions; explainers need the defining distinction and how to check it. Never replace useful detail with a generic brand routine or boilerplate cautions.
  - Omitted metadata is permitted. Through the CLI/connector a nonempty description is capped at 320 characters (import/update reject longer). Around 160 characters is a display advisory, not a search-engine limit or a quality score. Do not mechanically shorten an accurate, useful explanation.
  - Google may choose a query-dependent snippet from the body or the description. Automatic and authored descriptions can both be useful. Do not claim a guaranteed CTR increase, ranking gain, or AI citation preference from the field or its length.
  - Keep comparison counts, verified price context, distinctions and product roles when they materially answer the question. Reverify changing prices before using them; do not restore unsupported safety claims to add detail.
- **Existing review pass** — Within the usual review, check the answer and description against body, tables and FAQ for lost scope or contradictory claims. Apply concrete corrections or acknowledge a clean draft. Do not add a review/ack loop after every edit or a separate summary report.
- Render citations according to the selected channel's reference policy; keep evidence traceable even when public links are disabled
- **Inline product images are catalog-only and optional** — follow Step 3.2.1; otherwise keep the body text-only
- **Rich blocks are verified-only and optional** — at most one YouTube embed via `aeo source youtube <url>` and one product carousel via `aeo source products` (Step 3.2.2); never hand-write iframes or card HTML
- Mention the brand where it helps answer the question. First-party explanations can stand alone; do not impose a mention quota or manufacture a comparison.
- FAQ is optional; include only useful, supported adjacent questions

**Output format:**
- **Body**: Author in markdown for the import tool; the system converts it to canonical HTML. Do not reinterpret stored HTML as an instruction to change authoring syntax
- **Metadata**: Displayed separately from the body

### Step 4.5 — Review the draft

Use [content-review.md](content-review.md) with the saved or draft version actually being reviewed. Fix factual or explicit request defects within the authorized task; keep editing suggestions separate from factual defects. A single unsupported material claim can need correction, while several optional style suggestions do not make the facts false. Do not invent experiences, outside experts or weaknesses to satisfy a checklist. Self-review is not independent approval.

### Step 5 — Save draft

Determine the slug from the title (kebab-case, max 60 chars).

**How to determine the save path** — Since `pwd` may differ across environments (Cowork VM, local, etc.), follow this order:

1. If the user specified a save path, use it as-is
2. If a path was already used in the current session, use the same path
3. Otherwise, ask the user: "Where should I save the draft? (e.g., `~/Documents/aeo-drafts/`, current working directory, etc.)"

File name: `{AEOLO_DOMAIN_ID}_{slug}.md`

Tell the user the exact file path and that they can review/edit before importing.
Then suggest running `/aeo content import <path>`.

---

## /aeo content import [path] — Push draft to content history

Push a completed draft to the Aeolo content history for review and publishing.

### When to use

- After writing an article manually — import to Aeolo content history
- Any externally-written draft you want to register

### Flow

1. Read the draft file at `[path]` (default: the path used in the current session's manual draft step)
2. Extract or confirm: `title`, `targetKeywords`, `articleType`, `language`
3. Show the user a summary and confirm
4. Run the CLI command:

```bash
aeo content import \
  --title "Article Title" \
  --body-file ./path/to/draft.md \
  --type blog \
  --keywords "keyword1, keyword2, keyword3" \
  --language en \
  --rationale "Why this article exists" \
  --meta-description "SEO meta description"
```

Flags:

| Flag | Type | Required | Default |
|------|------|----------|---------|
| `--title` | string ≤60 chars | Yes | — |
| `--body` or `--body-file` | string / file path | Yes (one of) | — |
| `--type` | enum (blog, ranked_list, comparison, how_to, guide, faq, thought_leadership, case_study) | — | `blog` |
| `--keywords` | comma-separated | — | — |
| `--language` | enum | — | `en` |
| `--rationale` | string | — | — |
| `--meta-description` | string, up to 320 chars | Optional; preserve answer and article scope | — |
| `--sources` | JSON array `[{"name":"...","url":"..."}]` | — | — |
| `--target-channel` | channel id (from `aeo channels list`) | | |
| `--folder` | managed-site folder path (`/guide`) | | channel default |
| `--folder-id` | managed-site folder id (not with `--folder`) | | channel default |

> **Validation**: an omitted meta description is permitted. Import rejects a title over 60 characters or a description over 320. `/aeo content deploy` blocks only an empty title. The 160-character display advisory is not a publishing failure. Drafts may still need factual and editorial review even when length validation passes.

5. On success: "Imported → View in Aeolo dashboard → Content (`/contents/blog`)"

---

## GEO Writing Instructions — Agent Writing Guidelines

### Format Selection Matrix

Choose a form that resolves the reader's decision using available evidence. No default listicle, citation-rate quota or engine-specific word count is justified here. Supported depth determines length.

The rules below include Aeolo output conventions, not proven requirements for AI citation.

### Shared writing criteria

Do not add local quotas or assume a structure guarantees AI citations.

### Output conventions

Use real markdown H2/H3 headings when authoring for import. If a FAQ is useful, use `## FAQ` (or its localized equivalent) and H3 questions; do not add one just for schema. Keep metadata outside the body. Set the supported article type and let the platform render structured metadata; do not write JSON-LD or publication dates into article prose. Preserve verbatim quotations when applying house style.

### Engine Fit and Freshness

Use the approved audience and voice. Engine-specific experiments need dated, scoped evidence from the selected queries, market and engine; they are not universal tone or length rules. Review time-sensitive facts against current sources. Refresh when information changes or the page no longer resolves the reader's question, not merely because a fixed number of days elapsed. Keep publication/modification dates accurate; editing a date is not a substantive refresh.

### Brand Mention Principles

Represent the publisher honestly. A first-party product explanation can stand alone; comparison content needs a real decision and defensible criteria. Do not insert competitors or target a mention percentage solely to appear independent. Match product/service claims to the applicable catalog entry and evidence; brand voice is not a source for those claims.

### Metadata Generation

Generate the following metadata upon article completion (used in the import payload):

- `title` — SEO-optimized long version (question-based, keyword-rich). **Must differ from H1** (see Step 4 HARD RULE)
- `metaDescription` — Independently authored answer and article-specific scope, up to 320 characters; 160 is a display advisory
- `targetKeywords` — 1–20
- `articleType` — Based on the format matrix
- `estimatedRefreshDate` — An editorial review date based on the subject's expected change rate, not a promised citation lifespan

**Important: Do not include metadata in the article body.**
- When saving the draft: save only the pure markdown body to the file
- The agent tracks metadata separately and uses it only in the API payload during import
- When presenting to the user, show metadata separately from the body (e.g., "Here's the article I wrote + metadata shown separately")


### Managed-site folder destination

For a proxy destination, run `aeo site folders list`, then pass `--folder /ingredients` (or `--folder-id <id>`) with `--target-channel <managed-channel-id>` to `aeo content generate` or `aeo content import`. Both accept the same destination flags. The server resolves a path to its stable folder ID before enqueue/import, so renaming the path does not reroute a queued job. Omitted folder uses the default. Unknown, deleted, or foreign folders fail with available paths; ask for a valid choice instead of dropping the flag. Locale editions inherit their original article's folder.

Structured `content_import` uses `targetChannelId` plus `hostedFolderId` or `hostedFolderPath`. Never pass both folder selectors. Automation saves `config.hostedFolderId`; execution revalidates the stored selection and does not silently replace a deleted folder.
