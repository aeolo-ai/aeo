# GEO Strategy — Domain Knowledge

This file contains domain knowledge for the agent to read GEO data and make content strategy decisions.
After loading `/aeo`, always refer to this before interpreting data or deciding content direction.

---

## What is GEO

GEO work aims to improve supported brand representation, citations and recommendations in AI answers. Crawlability, indexing, retrieval, citation, recommendation and customer action are separate outcomes. Clear content and correct technical delivery are controllable inputs; they do not guarantee downstream results.

Treat the recommendations below as editorial decisions and testable hypotheses. Do not turn a citation share observed in one sample into a universal writing rule.

---

## Core GEO Inputs

```
Brand Context      → "What does this domain do?"
Content Strategy   → "What direction is already approved?"
Visibility Gaps    → "Where is AI not mentioning this brand?"
Existing Content    → "What can be refreshed, linked, or extended?"
```

Use Agent/Brand Context, content strategy, visibility gaps, and existing content to decide the next content action.

---

## How to Read Visibility Data

### A gap is a scoped observation

Read the metric actually measured: brand mention, recommendation and a linked citation are different. Retain engine, query, market/language, date and repeat count. Missing results or unmeasured engines are not zero visibility.

Before prioritizing, ask whether the question represents a customer decision, whether an eligible product/service can support an answer, what evidence is available, and whether existing content already answers it. Competitor appearance and repeated absence help prioritize within that scope; neither proves demand, conversion potential or the required format.

Comparison, use-case and foundational questions serve different reader decisions. Branded pre-purchase checks can also deserve content, while remaining separate from unaided discovery measurement.

### Discovery Prompts vs. Branded Diagnostics

The default tracked Prompt portfolio measures **unaided discovery**: would an AI
engine recommend the customer brand when the user did not name it first?

- Do not generate or retain the customer brand name, domain, product name, or a
  recognizable variant in the default visibility/SOV Prompt portfolio.
- Treat self-included Prompts such as "Is [Brand] good for sensitive skin?" as
  branded diagnostics, not discovery sensors. They may test answer accuracy or
  brand understanding, but they do not prove discoverability, Share of Voice,
  or competitive preference.
- Create branded diagnostics only when the user explicitly asks for a separate
  diagnostic set. Label and report that set separately from discovery Prompts.
- Never average branded-diagnostic results into discovery visibility, SOV, gap
  prioritization, or content-opportunity judgments.
- If an existing tracked portfolio mixes both surfaces, call out the
  contamination and recommend removing or untracking self-included Prompts from
  the discovery portfolio before interpreting the aggregate.

### Engine Priority

Use the customer's configured measurement scope and priorities. Compare only measured engines under the same relevant conditions; record the reason for an engine-specific experiment rather than assuming a universal engine ranking.

---

## What Aeolo's Own Measurements Show

Observed on Aeolo customer data (2026-08 to 2026-10). These are measured patterns, not guarantees. When you explain why a material, a fact or a planning change helps AI search, say only what this section supports, tie it to the brand's own numbers (tracked questions, their mention rate, what a material covers), and say plainly when something is unknown.

1. **AI engines never see the planning documents.** Brand Notes, Content Strategy, Source Guide and the status on each fact line (확인 / 공개 / 기록 / 충돌) are internal. They decide what the brand's articles may state as fact and how. Raising a line from 기록 to 확인 lets an article state it; it does not make any engine trust the brand more. Engines only see published pages.
2. **Engines cite pages, and a brand's own official pages are often its most cited pages.** In one measured brand, ChatGPT cited the home, about and promotion pages far more than any article, while Google-family engines and DeepSeek picked up the articles. Read citation coverage per engine, never as one average. An official page that states a fact clearly is both a 공개 source for Brand Notes and a citation target, so keeping it accurate matters.
3. **Effects travel through sources.** When the brand's content first appears among one engine's cited sources, the other engines' cited sources and mentions of the brand rise over the following weeks; mention changes without a source change were close to zero. A material helps when it becomes, or improves, a page engines can read. A fact that stays only in internal notes moves nothing by itself.
4. **What did and did not predict citation.** The domain itself was the largest single factor. Within a domain, a higher brand mention rate on the target question before publishing, English, and the canonical original (rather than a locale edition) predicted earlier citation. Outbound link count, a first-party-only link policy, article length, number of targeted prompts and llms.txt did not.
5. **Access is measured, not declared.** A robots.txt that allows everything can still sit behind a firewall that blocks citation crawlers (OAI-SearchBot, PerplexityBot, ClaudeBot) on article pages. GPTBot is a training crawler; reaching it proves nothing about citation.
6. **Citation rarely turns into a visit.** AI referral visits from cited pages were a low single-digit percentage. Talk about mentions and citations as the outcome; do not promise traffic or sales from a planning change.

How to use this in planning: when the user asks why a material or fact matters, answer with (a) which tracked questions it helps answer, (b) those questions' current mention rate, and (c) whether it would become or inform a published page. Do not claim any other effect.

---

## Gap → Content Decision Framework

> **Gaps are just one of many content triggers.** Client briefs, specific prompt targets, existing content refreshes (rewriting on the same topic), and other paths exist. Regardless of the path, all converge into the Pre-flight → Outline flow in `content-create.md`.

### Phase 1: Gap Clustering

Semantically group multiple gaps that can be covered by a single article. Criteria:
- **Is the same intent being asked in different ways?** → Same cluster
  - "best sunscreen stick for runners" + "top SPF sticks for outdoor sports" → one listicle
- **Same knowledge domain?** → Hub article opportunity
  - "what is GEO" + "how does AI citation work" + "why AI search matters" → foundational hub
- **Missing from the same prompt across engines?** → Can cover multiple engines with a single article

### Phase 2: Determine articleType

| Reader decision | Candidate articleType | Required basis |
|---|---|---|
| Choose among options | `ranked_list` | A real comparison set, explicit criteria and evidence for placements |
| Compare named alternatives | `comparison` | Matched dimensions and applicable product/service scope |
| Complete a task | `how_to` | Supported steps, prerequisites and limits |
| Understand a question or mechanism | `guide` | Direct answer, explanation and relevant evidence |
| Resolve several distinct questions | `faq` | Questions with separate useful answers, not repeated filler |
| Assess a perspective | `thought_leadership` | Clearly attributed reasoning and supported claims |
| Understand a documented outcome | `case_study` | Actual case evidence, conditions and limitations |

These are choices, not citation guarantees. A brand-owned product explanation need not pretend to be an independent ranking. Select the supported form after defining the reader's decision; do not default to a listicle because of an unattributed percentage.

### Phase 3: Test Engine-Specific Hypotheses

Use actual cited pages/answers from the configured engines to propose a hypothesis. Record query set, locale, dates, sample and confounders. Cited-page length, format or platform share is an observation, not proof that copying it causes citation. Test content quality first, then measure retrieval/citation and customer actions over time. There is no universal word count, freshness window or platform integration prescribed here.

---

## How to Read Audit Data

An audit reports observable technical/content conditions, not an AI trust or citation guarantee. Prioritize confirmed blockers to the intended page and reader journey.

| Finding | Check and action |
|---|---|
| Inaccessible or empty page | Confirm HTTP response, final URL and whether meaningful content is available to the relevant crawler. |
| Crawler restrictions | Inspect the applicable user agent, robots rules, authentication and indexing directives. Training, search and user-initiated bots can have different roles; do not assume all must be allowed. |
| Unclear page topic or answer | Improve descriptive headings and the answer where useful. Missing a TL;DR does not prove the page cannot be cited. |
| Missing or incorrect structured data | Validate relevant metadata against visible content and the actual renderer. Missing schema alone does not prove unreadability; schema does not guarantee rich results or AI citations. |
| Stale or misleading dates/facts | Correct the underlying information and accurate metadata; do not refresh dates as a substitute for substantive work. |
| Weak navigation/source traceability | Add useful related-page links and appropriate evidence links under the configured source policy. |
| Unclear authorship | Represent the actual publisher and verifiable author qualifications; never invent experience or credentials. |

Measure performance problems directly instead of applying an unsupported universal two-second crawler cutoff. Public HTTP success, indexability and actual indexing remain separate checks.

[Google's structured-data policies](https://developers.google.com/search/docs/appearance/structured-data/sd-policies) require accurate, representative markup and do not guarantee rich-result display. They are not evidence of AI citation uplift.

---

## Using Brand Context

Use [writing-inputs.md](writing-inputs.md) for authoring boundaries:

- Brand profile: supported identity and positioning. Do not treat promotional descriptors as proof of product-level efficacy.
- Catalog and scoped corrections: facts for the selected product, service, market or location; check the relevant sources.
- Brand/channel/task voice: expression, not a source of ingredient or qualification claims.
- `content_strategy.manifest`: current approved priorities and reader decisions; prefer it over generated snapshot/analysis fallbacks when present.
- Competitors: a comparison opportunity when the question requires it, not names to insert for naturalness.
- Brief and offering candidates: the planning decision and proposed fit, followed by evidence checks. A nomination alone does not establish a recommendation.

---

## Additional Context to Gather from the User

Things that may be insufficient from `/aeo` data alone. Check before writing:

**Product/service related:**
- Latest specs, pricing, launch dates (may not be in API data)
- Actual customer testimonials or case studies
- Competitive advantages over competitors (fact-based)

**Content strategy related:**
- Which engines are in the approved measurement or experiment scope?
- Is there a target language/market specification?
- Are there previously published related articles? (for internal linking)
- Publishing channel: own blog, Shopify, or external media?

**Research related:**
- Have external documents (PDF, Google Drive, web research) already been retrieved? → Use as `externalResearch`
- Are there specific statistics or data sources that must be used?

→ When these are clear, the research step can proceed quickly and the article's credibility increases.

---

## Advanced GEO Concepts (Reference for decision-making)

- **Topic relationships**: Connect related questions when that helps readers navigate or understand a decision; no guaranteed brand-association effect is asserted.
- **Hub-and-spoke**: A hub can organize distinct related questions. Do not create thin duplicate pages merely to satisfy the structure.
- **Citation variability**: Repeated measurements help distinguish a stable pattern from a snapshot. No fixed monthly fluctuation percentage is assumed.
- **Source concentration**: Repeatedly cited sources may help identify evidence gaps. Authority and relevance must be assessed for the actual claim, not just a domain label.



### GEO Recommended Strategy Types (Reference)

The angle for which gaps to fill first may vary depending on the brand's current situation:

| Strategy | Core Question | Suitable When |
|----------|---------------|---------------|
| `product-discovery` | "What's the best X?" / "X vs Y?" | Missing from comparison/recommendation queries. Most common |
| `thought-leadership` | "How do I do X?" | When seeking citations as an expert/authority source |
| `trust-reviews` | "Should I trust X?" | When trust barriers are high or reviews are lacking |
| `local-authority` | "Best X in [location]?" | When visibility is needed for location-based queries |
| `brand-awareness` | "What is X?" | When AI doesn't even know the brand exists — prerequisite for other strategies |

Use alongside `content_strategy.manifest` and current visibility gaps to determine content angle.

### Distribution (Reference)

Choose additional channels only when their audience and publishing rules fit an authorized distribution plan. Preserve attribution and canonical intent where supported. Cross-posting and cross-linking are not mandatory and do not establish increased AI trust by themselves.
