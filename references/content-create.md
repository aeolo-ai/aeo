# Content Creation

Default draft directory: `.aeo/` in the working directory. Add to `.gitignore`.
File naming: `{AEOLO_DOMAIN_ID}_{slug}.md`. See Step 5 for path resolution.

---

## /aeo content write — Write a GEO-optimized article

Agent writes the article directly using the guidelines below. No Mastra workflow API call — the agent is the writer. After writing, save the draft and import via `/aeo content import`.

---

### Pre-flight: Article Brief

**Mandatory gate before proceeding to the outline.** Regardless of the entry path, this checklist must be filled.

#### Trigger Paths (agent determines automatically)

| Path | Trigger | Already Available | Additionally Needed |
|------|---------|-------------------|---------------------|
| **Gap-driven** | `/aeo` visibility gap or `/aeo content propose` result | Brand data, target prompts, gap context | Competitor facts |
| **Client-brief** | User provides client materials/meeting notes/appeal points | Angle, materials (raw state) | Claims extraction, competitor context |
| **Prompt-targeted** | "I want us to show up for this prompt" | Target prompt | Brand claims, competitors |
| **Content Refresh** | Existing article's recency expired (91+ days) | Existing article structure + content | Latest data, updated competitor context |

Content Refresh is **writing a new article**, not a "patch." Read the existing article via `/aeo content get <id>` and use it as reference input, but the subsequent flow is the same as other Paths. A patch that only changes `dateModified` does not count as a GEO recency signal.

#### Article Brief Checklist

```
## Article Brief

### Trigger
- Starting point: [gap / client brief / prompt target / content refresh]
- Related gap ID, prompt, or existing content ID: [fill in if available]

### Brand Ammunition
- [ ] Brand Claims (minimum 3): claim + proof point + source
- [ ] Competitor Context: facts about comparison targets (required/optional depending on articleType)

### 1st-Party Experience (authenticity material)
- [ ] Internal tests/usage experience: who, how many times, under what conditions? (be specific)
- [ ] Customer testimonials/cases: is there actual user feedback?
- [ ] Domain-specific context: are there unique facts specific to this topic? (no generic statements)
- [ ] Positioning: will the article be transparent about being written by the brand?

> If the above items are empty, request them from the user. To avoid reading like "written by someone who's never used the product," at least 1 first-party material is needed.

### Article Direction
- [ ] Topic / H1 candidate
- [ ] articleType
- [ ] Target keywords
- [ ] Target engine(s)
```

If Brand Ammunition is empty, refer to **[brand-ammunition.md](brand-ammunition.md)**:
- Brand Claims → Extract from `/aeo domain brand` + session documents, then convert to Ammunition format
- Competitor Context → Web research or request from user

> **Authority Sources are not a Pre-flight target.** External third-party data (statistics, studies, etc.) for article credibility is gathered in Step 3 (Research).

---

### Step 1 — Collect inputs

Ask the user (or infer from context):
- **Topic** (required)
- **Article type** — choose from format matrix below; `blog` is default
- **Target keywords** (required, 1–20)
- **Language** — `en` (default) | `ko` | `ja` | `zh` | `ar`
- **Word count** — default 1500

If brand profile isn't loaded yet, fetch it first (`/aeo domain brand`) — brand context shapes the entire article.

### Step 1.5 — Brand Tone & Voice (MANDATORY GATE)

**Must be completed before writing a single word.**

Once the brand profile is loaded, look for:

- the default tone profile from `writing_styles`;
- GOOD/BAD/reference examples from `brand_voice_examples`;
- any legacy **Tone & Voice** notes in `brand_context` only as fallback context.

#### If a tone profile exists

Extract in the format below and apply consistently throughout all subsequent steps (Outline heading naming → Writing tone → FAQ style):

```
## Extracted Tone Profile
- Formality: [formal / balanced / casual]
- Voice characteristics: [extracted adjectives]
- Phrases to use: [example sentences/expressions]
- Phrases to avoid: [prohibited expressions]
- Other notes: [additional style instructions]
```

After extraction, confirm with the user in one line: _"This article will be written in a [characteristics] tone. Does that sound right?"_

#### If no tone profile exists

**Stop and ask the user directly. Never decide the tone arbitrarily.**

```
No writing style profile found for this brand.
I have a few questions to determine the article's tone:

1. Overall tone — Is it closer to formal (professional/polished) / casual (friendly/conversational) / authoritative?
2. Are there expressions or tones to avoid? (e.g., exaggerated marketing language, excessive technical jargon, etc.)
3. If you have example sentences or copy that you think represent the brand well, please share them.
```

After receiving answers, produce a reviewed writing style patch for
`writing_styles`, and include any concrete sample copy as proposed
`brand_voice_examples`. Do not bury voice examples inside `brand_context`.

- Interactive CLI/operator flow: after explicit approval, apply it with
  the writing style / voice example management surface when available.
- Background writing job or chat flow: do not write product memory directly.
  Include the patch in the final response for review.

This way, the same questions can be promoted into durable memory without
silently changing product context from an autonomous writing pass.

### Step 1.6 — Load Voice Examples (few-shot)

Load voice examples for blog content:

```bash
aeo post examples --platform blog
```

If examples exist:
- **GOOD examples** — internalize the rhythm, vocabulary, sentence structure. Match them.
- **BAD examples** — if your draft resembles any of these, rewrite.
- After writing, compare your draft against the GOOD examples. Would they feel like the same brand wrote it?

If no examples are returned, skip this step and rely on the writing style profile from Step 1.5 only.

> Voice examples provide concrete few-shot guidance that complements the abstract Tone & Voice profile. The profile says "casual and friendly"; the examples show *exactly how* that looks in practice.

### Step 2 — Outline

Build a structured outline:
- H1: question-based title matching actual AI search queries (see 10 commandments below)
- H2/H3: hierarchy where each section is independently quotable
- Note which sections will include comparison tables, expert quotes, or authority citations
- Identify 3–5 FAQ questions to cover at the end

Show the outline to the user and confirm before writing.

### Step 3 — Research (MANDATORY GATE)

**Research must be proactive and exhaustive before writing begins.** Read [data-access.md](data-access.md) for the full source configuration.

#### 3.1 — Load data sources

```bash
aeo config data-sources show
```

Read the configured sources. If none exist, use the default sources from data-access.md.

#### 3.2 — 1st-Party material sweep (ALWAYS FIRST)

Search for topic-relevant material in this order:

1. **Google Drive**: `aeo drive list` → scan folders mentioned in data_sources config → `aeo drive read <id>` for matches
2. **Custom URLs**: If data_sources lists specific URLs, read/browse them
3. **Brand context**: `aeo domain brand` → extract relevant claims and proof points
4. **Existing content**: `aeo content list --status=published` → check for related articles (avoid overlap, identify cross-link targets)

Extract 1st-party facts: product test results (who, how many, conditions), customer testimonials, internal data.

#### 3.3 — Ask user for gaps

If 1st-party material is insufficient for the topic:

```
"I found [what you found]. Do you have additional materials on [topic]?
(e.g., test reports, customer stories, internal data)"
```

If the user has nothing, note the gap and proceed — but flag it in Step 4.5 authenticity check.

#### 3.4 — External authority research

After 1st-party is exhausted, gather external sources:
- Recent statistics (.edu/.gov preferred)
- Expert quotes (name + title + affiliation)
- Industry reports and reviews
- Competitor specs (for comparison/ranked_list types)

Collect `{ name, url, description }` for each source — these become inline citations.

#### 3.5 — Research validation gate

Before proceeding to Step 4:
- [ ] Authority sources ≥ 3 (1st-party + external combined)
- [ ] 1st-party material ≥ 1 (if available for this topic)
- [ ] Every claim has a prepared citation
- [ ] Competitor facts verified (for comparison/ranked_list types)

**If validation fails, do not proceed to writing.** Request additional material from the user.

### Step 4 — Write

Write the full article following the **GEO Writing Instructions** below. Key rules:
- **Apply the Tone Profile extracted in Step 1.5 consistently across all sections** — including heading naming, sentence length, and tone
- **Title ≠ H1 (HARD RULE)** — These two titles must always be separate:
  - **Title** (`/aeo content import` title parameter): SEO-optimized long version (question-based, keyword-rich)
    Example: `"How to Reapply Sunscreen During a Long Run Without Stopping — A Stick Sunscreen Guide for Runners"`
  - **H1** (first `#` in the body): Short and punchy reader-facing heading
    Example: `"Reapply Sunscreen Mid-Run: The Runner's Guide"`
  - The body must start with `# {H1}`, and the title text must not be repeated anywhere in the body
  - The first paragraph after H1 should go straight to the BLUF answer
  - Do not use the same text for Title and H1
  - Do not repeat the H1 text in the intro
  - Example: `# Do You Need Sunscreen for Padel?` → `Glass walls on a padel court block wind, not UV. Here's what dermatologists say...`
- **Title length ≤ 60 chars (HARD RULE — deploy gate enforces this)**:
  - Mobile SERP truncates around 55 chars; desktop ~70. A 60-char ceiling lands every device cleanly.
  - Anything longer = "…" in the SERP listing, search-intent signals get cut, CTR collapses.
  - The Shopify deploy step rejects with HTTP 422 / `INVALID_TITLE_LENGTH` if `title.length > 60` — agent must shorten and retry.
  - Good: `"Best No-White-Cast SPF Sticks for Sports (2026)"` (47 chars)
  - Bad: `"What's the Best Sunscreen for NYC Half Marathon Runners in 2026? Race-Morning Picks That Won't Sting, Slip, or Leave a White Cast"` (133 chars — drop everything after the question mark)
- **Meta description is REQUIRED, 50–160 chars (HARD RULE — deploy gate enforces this)**:
  - The deploy step rejects with HTTP 422 / `INVALID_META_DESCRIPTION` if missing or out of range. Always pass `--meta-description` in `/aeo content import` and `/aeo content update`.
  - **Why it matters**:
    1. *SERP hook*: when missing, Shopify auto-fills from the body's first sentence — usually a generic intro that doesn't motivate clicks. A purposeful meta gives 2–5× CTR over fallback.
    2. *CTR feedback loop*: low CTR signals to Google "users don't find this useful" → article drops further in rank → impressions vanish. After honeymoon ends, recovery requires good CTR.
    3. *GEO citations*: ChatGPT / Perplexity sometimes lift the meta as the source summary when citing your URL — empty meta forces them to extract from the body, reducing citation accuracy.
  - **What to write (target 120–155 chars)**:
    - Sentence 1 — lead with the answer/result (BLUF style, mirrors the title intent)
    - Sentence 2 — one differentiator: founder credentials, first-party test, scope, or comparative angle
    - Optional close — specific value or invitation ("compared", "tested by…", "with…")
    - Never copy the body's opening sentence verbatim (LLMs default to this; resist)
  - Good: `"Sunscreen breaks down faster than you think. AAD says 2 hours — but sweat, water, and friction reset the clock. Here's the science, plus how stick formats fix mid-day reapplication."` (192 chars — too long, trim) → `"Sunscreen breaks down faster than 2 hours when sweat or water hits. Here's the AAD-backed science and why stick formats fix the mid-day reapplication gap."` (155 chars ✓)
  - Bad: `"Sunscreen is essential for athletes who spend time outdoors..."` (generic intro fallback — no CTR hook)
- BLUF in first 2–3 sentences
- Inline citations as `[Source Name](URL)` throughout
- Brand mentions at 15–25% density, always as part of a list (never solo promo)
- FAQ section at the end (3–5 questions)

**Output format:**
- **Body**: Pure markdown only (from H1 to FAQ)
- **Metadata**: Displayed separately from the body

### Step 4.5 — Semantic Authenticity Check (MANDATORY GATE)

**Must be performed before saving the article.** Self-check against the "Semantic Authenticity" checklist from [content-review.md](content-review.md):

| Check Item | Pass Criteria |
|------------|--------------|
| Positioning honesty | If written by the brand, does it avoid pretending to be an independent review? |
| Own product bias | Are own product weaknesses addressed as honestly as competitor weaknesses? |
| Independent expert voice | Is there at least 1 independent expert beyond founder/internal quotes? |
| Experience specificity | Do test/experience claims specify who, how many times, and under what conditions? |
| Domain-specific context | Does the article contain unique facts specific to this topic? (not just generic statements) |
| 1st-party data | Has the 1st-party material collected during Pre-flight been actually reflected in the article? |

**If 2 or more items fail**, show the results to the user and confirm whether to revise. A single failure shows a warning only and can proceed.

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

- After writing an article (via `/aeo content write` or manually) — import to Aeolo pipeline
- Any externally-written draft you want to register

### Flow

1. Read the draft file at `[path]` (default: the path used in the current session's content write step)
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
| `--meta-description` | string 50–160 chars | **Required for deploy** | — |
| `--sources` | JSON array `[{"name":"...","url":"..."}]` | — | — |

> **Deploy gate**: `/aeo content import` itself accepts a missing or oversized meta (drafts iterate freely), but `/aeo content deploy` will return HTTP 422 with `INVALID_TITLE_LENGTH` or `INVALID_META_DESCRIPTION` if the article fails the SERP-friendly limits at deploy time. Always supply both at import so the article is deploy-ready.

5. On success: "Imported → View in Aeolo dashboard → Content Queue"

---

## GEO Writing Instructions — Agent Writing Guidelines

### Format Selection Matrix

When in doubt, **default to ranked_list** — 53% of AI citations come from listicles.

| Prompt Pattern | articleType | AI Citation Rate | Length Guide |
|---|---|---|---|
| "best X", "top X", "X recommendations" | `ranked_list` | 32% | 2,000–4,000 words |
| "X vs Y", "compare X and Y" | `comparison` | 18% | 1,200–2,500 words |
| "how to X", "step by step X" | `how_to` | 15% | 1,500–3,000 words |
| "what is X", "why X", "X explained" | `guide` | — | 800–1,500 words |
| Complex questions, multiple questions at once | `faq` | 11% | 1,500–2,500 words |
| Industry trends, expert perspectives | `thought_leadership` | — | 1,500–3,000 words |
| Customer stories, adoption results | `case_study` | — | 1,200–2,000 words |

### GEO Writing 10 Commandments

1. **BLUF (Bottom Line Up Front)** — Place the core answer in the first 2–3 sentences. AI cites "specific answers," not entire articles. Don't beat around the bush in the intro.
2. **Title ≠ H1** — Title is the long SEO version (question-based, keyword-rich). The body `#` is a short, punchy reader-facing heading. They must always be different text. Example: Title `"What's the Best SPF Stick for Outdoor Sports in 2026?"` → H1 `"Best SPF Sticks for Outdoor Sports"`
3. **Logical H2/H3 hierarchy** — Semantic HTML5 structure. Each section must be independently quotable. A single H2 should make sense on its own.
4. **Comparison tables** — Comparison data must be in markdown/HTML tables. AI prefers structured data over unstructured text.
5. **Authority signals** — External authority source citations are mandatory. Statistics, research, .edu/.gov sources. Inline citation format: `[Source Name](URL)`. No unsourced claims.
6. **Expert quotes (attributed)** — "Real name + title + quote" format. AI uses this for credibility assessment. No fabricated quotes.
7. **FAQ section** — 3–5 FAQs at the bottom. Cover related questions not addressed in the main body.
8. **Schema markup hints** — Specify the schema type as an HTML comment at the top of the article (before H1). Format: `<!-- schema: Type1, Type2 -->`

   | articleType | Schema Hint |
   |---|---|
   | ranked_list | `<!-- schema: Article, ItemList -->` |
   | comparison | `<!-- schema: Article, ItemList -->` |
   | how_to | `<!-- schema: HowTo, Article -->` |
   | faq | `<!-- schema: FAQPage, Article -->` |
   | guide | `<!-- schema: Article -->` |
   | blog | `<!-- schema: Article, BlogPosting -->` |
   | thought_leadership | `<!-- schema: Article, BlogPosting -->` |
   | case_study | `<!-- schema: Article -->` |
9. **Freshness signals** — Specify `datePublished` and `dateModified` at the top of the article. Content citation rate drops from 100% within 30 days to 18% after 1 year.
10. **Internal + external links** — Cross-link your own content + link to external authority sources. AI actively crawls link graphs.

### Platform-Specific Tone Guide

If not specified, default to a "practical + structured" combination that works for both ChatGPT and Gemini.

| Engine | Primary Citation Sources | Tone & Structure |
|--------|--------------------------|------------------|
| **ChatGPT** | Wikipedia, Global news, Blogs | Practical, conversational, how-to focused. Encyclopedic tone preferred. Average **2,800 words** benchmark. |
| **Claude** | Academic content | 5,000+ character long-form. Technical documentation tone. Academic citations, structured argumentation. Depth-first. |
| **Perplexity** | Blog/editorial, News, Expert reviews | Niche expertise + content from **within the last 90 days**. High fact density. |
| **Gemini** | **YouTube (category leader)**, Blogs, News | Schema-enriched. Structured data. Consider YouTube content in parallel. |

### Content Freshness Rules

| Article Age | Residual Citation Rate | Agent Action |
|-------------|------------------------|--------------|
| 0–30 days | 100% | Optimal time for new article publishing |
| 31–90 days | 73% | Data/statistics update recommended |
| 91–180 days | 51% | Refresh needed — new data + structural improvements |
| 181–365 days | 34% | Major rewrite |
| 1 year+ | 18% | New article recommended |

When updating existing content, always update `dateModified`. Replace with the latest statistics/data when possible.

### Brand Mention Principles

- Brand appears as **part of a list** — never as standalone promotional content
- Mention competitors alongside for naturalness
- Fact-based information only (specs, pricing, review summaries)
- Mention density: **15–25%** of the total article — excessive mentions reduce AI trust
- Use brand VP and key features from the brand profile (`/aeo domain brand`)

### Metadata Generation

Generate the following metadata upon article completion (used in the import payload):

- `title` — SEO-optimized long version (question-based, keyword-rich). **Must differ from H1** (see Step 4 HARD RULE)
- `metaDescription` — Under 150 characters (BLUF-based summary)
- `targetKeywords` — 1–20
- `articleType` — Based on the format matrix
- `estimatedRefreshDate` — Publish date + 60 days

**Important: Do not include metadata in the article body.**
- When saving the draft: save only the pure markdown body to the file
- The agent tracks metadata separately and uses it only in the API payload during import
- When presenting to the user, show metadata separately from the body (e.g., "Here's the article I wrote + metadata shown separately")
