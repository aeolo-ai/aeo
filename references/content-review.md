# Content Review

Reviews existing content based on GEO domain expertise.

---

## /aeo content review <id> — Content review from a GEO perspective

### Flow

1. **Load context** — Fetch the following 4 items in parallel:
   ```bash
   aeo content get <id> > /tmp/aeo_review_article.md &
   aeo agent context > /tmp/aeo_review_brand.md &   # /aeo agent context
   aeo domain audit > /tmp/aeo_review_audit.md &    # /aeo domain audit
   aeo content --status=published --limit=50 > /tmp/aeo_review_published.md &   # for Audit #0b
   wait
   ```

2. **Perform review** — Evaluate the article against the checklist below.

3. **Output report** — Present results in the format below.

4. **Suggest next actions** — If edits are needed, suggest patches via `/aeo content update <id>`.

---

### Review Checklist

#### 0. Fabrication & Citation Audit (BLOCKING — run this first)

Run this **before** the quality checklist below, and run it **adversarially**: assume the draft contains at least one fabricated quote and one unsourced statistic, and go find them. This pass is not about confirming the draft is good — it is about proving that every quote and every number is real and traceable. A draft that fails any check here is **NOT deploy-ready**, no matter how well it scores on categories 1–6.

| Item | Adversarial check | Verdict if it fails |
|------|-------------------|---------------------|
| **Quote traceability** | For EVERY direct quotation (anything in quotation marks attributed to a person, study, or organization), name the exact source URL — gathered during research — that contains those words. Can you point to the URL? Do those words actually appear there? | If a quote has no traceable source URL, it is a **FABRICATION**. Remove it, or convert it to an unquoted paraphrase of a real cited source. Never pass a draft with an unverifiable quote. |
| **Named-person check** | For every named expert / advisor / customer, does a real citable source establish that this person said this? An invented name (e.g. "advisor Julio Pina") with no source is a fabrication. | Cut the name and quote, or replace with a real, cited expert. |
| **Quantified-claim sourcing** | For EVERY number, percentage, "#1" / "best" / "most" superlative, "studies show", or factual/spec claim — is there an inline `[Source Name](URL)` link next to it that traces to a real source? | If a quantified or superlative claim has no inline source, **cut it or hedge it** (drop the specific number / soften to a non-quantified statement). Do not ship unsourced stats. |
| **Source-to-claim match** | Open the cited sources. Does each source directly support the attached claim at the same strength and scope? Check every mechanism, practical recommendation, causal inference, and generalized conclusion — not only quotes and numbers. | If the evidence is narrower than the claim, narrow or remove the claim. If the source does not support it, fix the claim or find a real source. |

**How to record the verdict:** the Fabrication & Citation Audit must itemize every quote and every quantified claim with its traceable source URL (or "NO SOURCE → removed/hedged"). "Issues found: none" is only acceptable when you have explicitly walked each quote and each stat and shown its source. A blanket "looks good / deploy-ready" without this itemized trace is itself a review failure.

#### 0b. Cannibalization Audit (BLOCKING — run this second)

A draft can be flawless on every other axis and still be a defect: if the brand already published this topic, the new article splits the brand's own citation signal instead of adding to it. Two pages competing for one query means AI engines see a weaker, more diffuse answer for both.

Run this against `/tmp/aeo_review_published.md`. **Compare topics, not title strings** — the failure mode is two different-sounding headlines covering the same ground:

> "Is K-Beauty Sunscreen Safe for Sensitive Skin?" vs "HAESKN Sunscreen Stick for Sensitive Skin Guide"
>
> Different strings. Same article. An exact-title check passes this; a topic check does not.

| Item | Adversarial check | Verdict if it fails |
|------|-------------------|---------------------|
| **Recent duplicate** | Does any article published in the **last 14 days** answer the same core question, even under a different headline? Strip brand names and format words ("best", "guide", "review") from both titles and compare what remains. | **BLOCKING.** Do not deploy a second article on a topic covered within two weeks. Recommend updating the existing article instead (`aeo content update <existing-id>`). |
| **Cluster saturation** | How many published articles already cover this topic cluster? | **3 or more → BLOCKING.** The problem is not coverage, it is that the existing pages are not being cited. Another article will not fix that; strengthening the best existing one might. |
| **Spoke justification** | If the draft claims to be a "spoke", a "different angle", or to "add new context" on an existing topic — name the specific question it answers that the existing article does **not**. Then check: does the existing article already answer it in a section? | If you cannot name a question the existing article leaves unanswered, this is **not a spoke — it is a duplicate**. "Connects two existing pieces" and "approaches it through a different lens" are rationalizations, not new coverage. |
| **Hub link** | If it genuinely is a spoke, does the draft link back to the hub article, and does the hub have room to link forward to it? | Add the cross-link before deploy. An unlinked spoke reads to an engine as a competing page, not a supporting one. |

**Why this section exists:** the review pass previously had no cannibalization item at all, so a duplicate published eight days after its twin received a clean "ZERO defects, deploy-ready" verdict. The reviewer was not wrong about what it checked — the check simply did not exist. If you find yourself writing a justification for why a near-duplicate is acceptable, that is the signal to fail it, not to pass it.

#### 1. Structure & Quotability

| Item | Criteria | Reference |
|------|----------|-----------|
| **BLUF** | Is the core answer in the first 2–3 sentences? | 10 Commandments #1 |
| **H1** | Is it a question-based title? Does it match actual AI queries? | 10 Commandments #2 |
| **H2/H3 hierarchy** | Can each section be independently quoted? | 10 Commandments #3 |
| **Markdown headings** (blog articles only) | Is every section heading a real markdown `## `/`### ` heading — not bold text (`**Section**`) or a bare numbered line? Not applicable to social posts. | 10 Commandments #3 |
| **Comparison tables** | If comparison data exists, is it structured as a table? | 10 Commandments #4 |
| **FAQ section** | Are there 3–5 supported adjacent questions at the bottom that the main body has not already answered? | Replace or remove any FAQ that merely restates a body conclusion, or whose answer introduces unsupported advice. |
| **FAQ is a `## ` H2** (blog articles only) | Is the FAQ section under a real markdown `## FAQ` / `## 자주 묻는 질문` H2 (not a bold "FAQ" line)? Deploy-time FAQPage schema only fires when it is. Not applicable to social posts. | 10 Commandments #7 |
| **Schema hints** | Is the recommended schema type specified? | 10 Commandments #8 |

#### 2. Trust & Authority

| Item | Criteria | Reference |
|------|----------|-----------|
| **Inline citations** | Are there enough `[Source Name](URL)` inline citations? (1–2 per section) Does every quantified/factual claim carry one? (see Audit #0) | 10 Commandments #5 |
| **Expert quotes** | Are they in "real name + title + quote" format, and does each quote trace to a real source URL? (see Audit #0 — any untraceable quote is a fabrication and must be removed) | 10 Commandments #6 |
| **Authority sources** | For scientific, medical, safety, or regulatory claims, are primary research, official regulators, or professional institutions used instead of commercial explainers or ingredient suppliers? | 10 Commandments #5 |
| **Internal + external links** | Are both internal content links and external authority source links present? | 10 Commandments #10 |

#### 3. Freshness

| Item | Criteria | Reference |
|------|----------|-----------|
| **Freshness metadata** | Is the body free of `datePublished`/`dateModified` lines? (dates live in structured metadata / deploy-time JSON-LD, not the body) | 10 Commandments #9 |
| **Data recency** | Are cited statistics/data less than 1 year old? | Freshness rules |
| **Article age** | What is the residual citation rate based on publish date? (0–30 days 100% → 1 year+ 18%) | Freshness rules |

#### 4. Brand Integration

| Item | Criteria | Reference |
|------|----------|-----------|
| **Mention density** | Is it within 15–25% of the total article? | Brand mention principles |
| **Appears within a list** | Does the brand appear as part of a list, not as standalone promotion? | Brand mention principles |
| **Competitors mentioned together** | Are competitors mentioned alongside for naturalness? | Brand mention principles |
| **Fact-based** | Is only verifiable information used (specs, pricing, review summaries)? | Brand mention principles |
| **Tone consistency** | Is it consistent with the approved task-specific task-specific reference evidence and relevant voice examples? | content-create Step 1.5 |

#### 5. Semantic Authenticity

Catches semantic issues where the article reads like "AI-written content" or "brand advertisement." Even if structure and sources are perfect, failing this category means neither AI engines nor readers will trust it.

| Item | Criteria | Red Flag Example |
|------|----------|------------------|
| **Positioning honesty** | Is the article honest about its identity? If written by the brand, does it pretend to be an independent review? | Ranking own product #1 while listing "Editorial Team" as author → disguised advertising |
| **Own product bias** | Are own product weaknesses addressed as honestly as competitor weaknesses? Not bashing competitors while going easy on own product? | Competitor: "only 40min water resistance" vs Own: "water resistance not independently rated" (cushioned) |
| **Independent expert voice** | Beyond founder/internal quotes, is there at least one independent expert (dermatologist, researcher, etc.)? | All quotes from founders → feels like a press release |
| **Experience specificity** | Are test/experience claims specific? Who, how many times, under what conditions? | "tested through months of sessions" (zero specificity) vs "tested by 3 players across 12 sessions on outdoor courts in 30°C+" |
| **Testing methodology** | For comparison/review articles, are evaluation criteria and methods stated? | Wirecutter: separate methodology section with timing/location/personnel specified. Ours: none |
| **Domain-specific context** | Does the article contain context unique to the topic? Not just generic statements? | Padel article with no padel court characteristics (glass wall glare, match duration) → feels like any sport was substituted in |
| **1st-party data** | Does it include 1st-party experiences such as internal tests, customer testimonials, or usage data? | All external sources only → "written by someone who never used it" |
| **Author E-E-A-T** | Does it have a real author name + bio + credentials? Not anonymous like "Editorial Team"? | "Acme Editorial Team" → unknown identity, suspected AI/ghost writers |

**Core principle**: Even if structure and sources are perfect, without genuine experience and honesty, the article will read as "AI-generated." This category evaluates whether the article stems from real human experience.

#### 6. Engine Fit

Based on the target engines in the brand context or visibility gap data:

| Engine | Check Points |
|--------|-------------|
| **ChatGPT** | Practical/conversational tone, ~2800 words, how-to structure |
| **Gemini** | Schema-enriched, structured data, YouTube integration consideration |
| **Perplexity** | Niche expertise, data from within the last 90 days, high fact density |
| **Grok** | Real-time trends, community reaction incorporation |

---

### Report Format

```
## GEO Content Review — "{article title}"

### Summary
- **Overall**: ✅ Good / ⚠️ Needs Work / ❌ Major Issues
- **Word count**: {n} words
- **Article type**: {type}
- **Cannibalization**: ✅ no overlapping published article / ❌ duplicates "{title}" (published {n}d ago)

### Scores

| Category | Score | Notes |
|----------|-------|-------|
| Fabrication & Citation (#0) | ✅ / ❌ | quotes traced: {n}, stats traced: {n} |
| Cannibalization (#0b) | ✅ / ❌ | closest published article + why it is or is not a duplicate |
| Structure & Quotability | ✅ / ⚠️ / ❌ | ... |
| Trust & Authority | ✅ / ⚠️ / ❌ | ... |
| Freshness | ✅ / ⚠️ / ❌ | ... |
| Brand Integration | ✅ / ⚠️ / ❌ | ... |
| Semantic Authenticity | ✅ / ⚠️ / ❌ | ... |
| Engine Fit | ✅ / ⚠️ / ❌ | ... |

### Issues Found
1. **[Category]** — {specific issue} → {recommended fix}
2. ...

### Recommended Patches
> If edits are needed, include patch suggestions ready to use with `/aeo content update <id>`.
> Apply after user confirmation.
```

---

### Notes

- Reviews are **read-only** — not subject to the CUD Rule. Edits are made via `/aeo content update` after user confirmation.
- Externally written articles (local files) can also be reviewed — if a file path is provided instead of `<id>`, the file is read and the same checklist is applied.
- Reviews can be performed without brand context, but the Brand Integration category is skipped and this is noted explicitly.
