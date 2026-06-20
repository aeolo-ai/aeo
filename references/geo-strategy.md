# GEO Strategy — Domain Knowledge

This file contains domain knowledge for the agent to read GEO data and make content strategy decisions.
After loading `/aeo`, always refer to this before interpreting data or deciding content direction.

---

## What is GEO

While SEO focused on search result "rankings," GEO aims to be **cited, recommended, and trusted** by AI engines like ChatGPT, Claude, Perplexity, and Gemini. It's a fundamentally different paradigm.

| Traditional SEO | GEO |
|-----------------|-----|
| Page ranking optimization | Citation & recommendation optimization |
| Volume-first | Consistency-first |
| Create once and leave | Continuous updates on 30–60 day cycles |
| DA (Domain Authority) focused | Third-party mention frequency + structured content |
| Single search engine | 4+ AI engines simultaneously |

**GEO rewards speed and structural clarity. Both are variables the brand can control.**

---

## Aeolo 4-Step Pipeline

```
Step 1: Brand    → "What does this domain do?"
Step 2: Prompts  → "What should we ask the AI engines?"
Step 3: Visibility → "Is the AI mentioning this brand?"
Step 4: Content  → "What's missing, and what should we write?"
```

The agent is the executor of Step 4 — filling gaps with content. First, read and analyze the data from Steps 1–3.

---

## How to Read Visibility Data

### Gap = "This engine did not cite our brand for this prompt"

What to check in the visibility report:
1. **Which engine** was the brand missing from? → Strategy differs by engine
2. **Which stage** prompt was it missing from? → Determines content type and angle
3. **Are competitors appearing?** → If yes, comparison content opportunity; if nobody appears, foundational content opportunity
4. **How many engines simultaneously missed it?** → More engines = higher priority

### Meaning and Priority by Stage

| Stage | Meaning | Query Example | GEO Priority | Reason |
|-------|---------|---------------|--------------|--------|
| `comparison` | Comparing options — right before purchase | "best CRM for startups", "which CRM for a B2B startup" | **1st priority** | High-intent queries without brand names. AI recommendations directly influence purchase |
| `use-case` | Specific situation-based | "CRM for 10-person B2B team" | **2nd priority** | Long-tail, low competition, high conversion |
| `foundational` | Concept/need exploration | "what is CRM", "why use CRM" | **3rd priority** | Brand awareness building. Long-term effect |
| `implementation` | Post-purchase / pre-purchase verification | "HubSpot pricing", "how to use [BrandName]" | **Skip** | Brand name already in the query → not a GEO target |

### Engine Priority

Unless otherwise specified: **ChatGPT > Gemini > Perplexity > Grok**

However, two exceptions:
- If the user specifies a particular engine, follow their preference
- If visibility checks were only run on some engines, evaluate only among the checked engines

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

Decide based on the gap's stage + query pattern:

| Gap Pattern | Recommended articleType |
|-------------|------------------------|
| comparison stage + "best X", "top X" | `ranked_list` — **always consider first** |
| comparison stage + "X vs Y" | `comparison` |
| use-case stage + "best X for [situation]" | `ranked_list` or `how_to` |
| foundational stage + "what is X" | `guide` |
| foundational stage + "how to X" | `how_to` |
| Multiple stages simultaneously | `faq` (serves as a hub) |

**Listicle (ranked_list) is the default**: 53% of AI citations come from listicles. When the format isn't clear, go with ranked_list.

### Phase 3: Apply Engine-Specific Strategy

When writing content, tailor it to the characteristics of the engine where the gap exists:

| Engine | Primary Citation Sources | Tone & Strategy |
|--------|--------------------------|-----------------|
| **ChatGPT** | Wikipedia, Global news sites, Blogs | Practical, conversational, how-to. Prefers encyclopedic tone. Average cited page is **2,800 words** |
| **Gemini** | **YouTube (category leader in most categories)**, Blogs, News sites | Schema-enriched, structured data. Consider YouTube content in parallel |
| **Perplexity** | Blog/editorial, News, Expert reviews | Niche expertise + preference for content from **within the last 90 days** |
| **Grok** | X (Twitter), real-time news | Real-time trends, community reaction incorporation |

If gaps exist across multiple engines simultaneously, default to the ChatGPT + Gemini combination ("practical + structured").

---

## How to Read Audit Data

Audit score measures "how well AI engines can read and trust this site."

### Item Meaning and Actions

| Audit Item | What It Means When Missing | Immediate Action |
|------------|----------------------------|------------------|
| **Schema (FAQ, HowTo, Article)** | AI cannot distinguish content type | Add FAQ/HowTo JSON-LD |
| **H1 empty** | AI cannot identify the page's core topic | Add text with keywords to H1 |
| **No TL;DR/BLUF** | AI cannot cite a "specific answer" | Add a 2–3 sentence summary at the top of every article |
| **No datePublished/Modified** | AI cannot assess freshness → citation plummets | Add `<time datetime="">` + OG meta |
| **Insufficient internal links** | AI cannot connect related content | Hub-and-spoke structure + cross-links between articles |
| **No listicle/comparison content** | Missing 53% of AI citations | Prioritize ranked_list + comparison articles |
| **No author byline** | AI has difficulty assessing credibility | Add name + title + credentials |

### Audit Priority Assessment

- **HIGH (immediate)**: No schema, H1 empty, no TL;DR, dates not structured
- **HIGH (content)**: No listicles, no Hub-and-Spoke, insufficient internal links
- **MED**: No author byline, insufficient external links

Audit issues = even well-written articles can't be read by AI. **Address HIGH items before writing articles.**

### Technical Crawler Accessibility (AI crawlers differ from Googlebot)

Before schema/content, verify that AI can actually read the site:

| Item | What It Means When Failing | How to Verify |
|------|----------------------------|---------------|
| **SSR (Server-Side Rendering)** | AI crawlers often can't execute JS → content not recognized | Check if body text exists in the HTML source |
| **robots.txt AI crawler access** | If GPTBot, anthropic-ai, PerplexityBot, etc. are blocked → invisible | Ensure `Allow: /` is specified in robots.txt |
| **Page load under 2 seconds** | AI crawlers don't wait as long as Google | Check Core Web Vitals |

If any of these 3 are blocked, all other optimizations are meaningless. If technical issues are suspected, flag them to the user first.

---

## Using Brand Context

From data loaded via `/aeo domain brand`:

- **competitors** usage: Include alongside competitors in comparison content for naturalness
- **key_features + value_proposition**: Source for fact-based descriptions when mentioning the brand
- **brand_context**: Market positioning, target audience, core narrative — used for determining article angle
- **content_strategy.manifest**: Reviewed priorities, target angles, and publishing direction. Prefer it over generated snapshot/analysis fallbacks when present.

---

## Additional Context to Gather from the User

Things that may be insufficient from `/aeo` data alone. Check before writing:

**Product/service related:**
- Latest specs, pricing, launch dates (may not be in API data)
- Actual customer testimonials or case studies
- Competitive advantages over competitors (fact-based)

**Content strategy related:**
- Is there a target engine specification? (If not, default to ChatGPT + Gemini)
- Is there a target language/market specification?
- Are there previously published related articles? (for internal linking)
- Publishing channel: own blog, Shopify, or external media?

**Research related:**
- Have external documents (PDF, Google Drive, web research) already been retrieved? → Use as `externalResearch`
- Are there specific statistics or data sources that must be used?

→ When these are clear, the research step can proceed quickly and the article's credibility increases.

---

## Advanced GEO Concepts (Reference for decision-making)

- **Semantic Neighbourhood**: What concepts and brands AI associates with the brand. Frequent co-occurrence with desired keywords strengthens the association
- **Hub-and-Spoke**: 1 hub (long-form guide) + N spokes (specialized articles). Helps AI understand site structure and connect related content
- **Citation drift**: AI citations fluctuate 40–60% monthly. Trends matter more than a single snapshot. 30–60 day content refresh cycles recommended
- **Trust Spine**: The 5–10 core high-authority sources AI references when citing. Building these stabilizes visibility

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

### Cross-Posting Flow (Reference)

Distribution sequence to increase AI model trust and mention frequency after article publication:

**Blog (canonical)** → LinkedIn article → Medium (with canonical tag) → Substack → Reddit (relevant subreddits, authentic engagement)

Cross-linking between all channels is mandatory. AI actively crawls link graphs.
