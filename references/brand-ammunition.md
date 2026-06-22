# Brand Ammunition Guide

A guide for preparing "things you can say" before writing an article.

**This file is referenced during the Pre-flight stage.** When Brand Claims and Competitor Context are empty in `content-create.md`'s Pre-flight: Article Brief, use this guide to fill them.

---

## Brand Claims — Claims you can make + supporting evidence

When writing about the brand in an article, every claim must have a proof point and source. **Claims without a source must not be used.**

### Format

```
| Claim | Proof Point | Source | Applicable Context |
|-------|------------|--------|-------------------|
| [claim] | [evidence — numbers, specs, facts] | [source URL/document name] | [which article types is this suitable for] |
```

### Writing Rules

- No exaggeration — use "one of the top-rated" instead of "best"
- Prioritize quantitative data (pricing, specs, numbers)
- Founder/team credentials should be used as expert quotes ("real name + title + quote" — 10 Commandments #6)
- Customer reviews/test results require attribution

### Claims Source Priority

1. Product pages/official specs (most reliable)
2. Founder/team interviews, press releases
3. Third-party reviews/test results
4. Client-provided materials (Google Drive, meeting notes, etc.)

### Brand Context → Ammunition Conversion

Brand context data retrieved via `/aeo agent context` is in raw form. It needs to be converted to the Ammunition format.

| Brand Context Field | Ammunition Conversion Method |
|---------------------|------------------------------|
| `value_proposition` | Break down into core claim statements → find proof point + source for each claim |
| `key_features` | Convert each feature into a Claim → verify proof from specs/product pages |
| `category` / `industry` | For positioning context — not direct Claims |

If brand context alone lacks sufficient proof:
1. Supplement from session documents (Google Drive, meeting notes, product specs)
2. Crawl official website/product pages
3. If still insufficient, request from the user

---

## Competitor Context — Comparison table material

Facts for when the brand appears alongside competitors in comparison/list articles. Spec-based, not opinion-based.

### Required/Optional by articleType

| articleType | Competitor Context |
|---|---|
| `ranked_list`, `comparison` | **Required** (5+ recommended) |
| `blog`, `how_to`, `guide` | Optional |
| `thought_leadership`, `faq`, `case_study` | Optional (nice to have) |

If required for the article type but Competitor Context is empty, fill via web research or request from the user.

### Format

```
| Competitor | Price | Key Specs | Factual Differences vs Ours |
|------------|-------|-----------|----------------------------|
| [name] | [$] | [specs] | [fact-based differences] |
```

### Writing Rules

- No competitor bashing — list facts only, let the reader judge
- Pricing based on list price + date verified
- No direct "we're better" comparisons → parallel spec listing
- For `ranked_list`: 5+ items; for `comparison`: minimum 2 direct comparison targets

### Competitor Context Sources

1. Aeolo visibility data — reference competitor domains from `topCompetitors`
2. Competitor official websites/product pages
3. Third-party reviews/comparison articles

---

## Content Refresh (Path D) — Reading existing articles

When using existing content as reference input:

```bash
# If you have an existing content ID, read it first
aeo content get <id>
```

What to check in the existing article:
- Sections/structures to keep (claims still valid, competitor comparisons still accurate)
- Data to replace (outdated statistics, discontinued products, changed pricing)
- Content to remove (positioning that no longer aligns with strategy changes)

The subsequent flow is the same as other Paths — fill the Pre-flight checklist and proceed to Outline.

> **A patch that only changes `dateModified` does not count as a GEO recency signal.** Content Refresh means writing a new article.
