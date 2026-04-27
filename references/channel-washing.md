# Channel Washing — Multi-Channel Content Adaptation

Transform hub articles, original ideas, or data insights into channel-native format that AI engines crawl and cite. This is the "washing layer" — same core claim, different voice per platform, optimized for both human engagement and AI extraction.

---

## When to Use

| Input | Flow |
|-------|------|
| Hub article exists | Read article → pick angle → wash for channel |
| Original idea (no article) | Write directly in channel format |
| Data/insight snippet | Frame as channel-native post |

Works with `/aeo content write` (articles) and standalone channel content.

---

## Universal Rules (All Channels)

1. **One canonical URL per hub article** — all spokes link back to it. AI actively crawls link graphs; every spoke→hub backlink strengthens the canonical page's citation probability
2. **Core claim consistency** — the factual claim is identical across channels; only framing changes
3. **Brand density decreases as informality increases**: Blog 15-25% → LinkedIn 10-20% → Reddit 5-10% → Threads <5%
4. **Never copy-paste between channels** — each must feel native. AI engines detect duplicate content across domains
5. **1st-party experience framing** — "I/we tested" not "Brand X offers"
6. **Every piece must contain an AI Citable Unit** — a self-contained 2-3 sentence block that answers a query without surrounding context (TL;DR, BLUF, Key Takeaway)

### Hub → Spoke Timing

| Day | Action | Channel | GEO Rationale |
|-----|--------|---------|---------------|
| D+0 | Hub article live | Blog (canonical) | Establishes canonical URL for link graph |
| D+0~1 | Reddit value post | Reddit | Perplexity re-indexes Reddit within 24-48h; early posting captures citation window |
| D+1~2 | LinkedIn post + article | LinkedIn | ChatGPT crawls LinkedIn articles; Article indexed by search engines |
| D+2~3 | Thread series | Threads | Meta indexing pipeline; brand signal amplification |
| D+7 | Reddit follow-up (answer comments, add insights) | Reddit | Comment engagement boosts post visibility → higher upvotes → citation threshold |
| D+14 | LinkedIn follow-up post (results, additional analysis) | LinkedIn | Recency signal for AI engines |
| D+30 | Performance review across channels | All | Citation drift = 40-60% monthly; refresh cycle decision |

---

## Reddit

### GEO Mechanics

| AI Engine | Reddit Citation Rate | How It Crawls |
|-----------|---------------------|---------------|
| Perplexity | **46.7%** of social citations | Direct crawl; TL;DR extracted as "default explanation" |
| Google AI Overview | **21%** | Google indexing pipeline; Reddit threads rank on page 1 |
| ChatGPT | Top 10 domain | GPTBot active crawl; prefers posts with structured findings |
| Gemini | Indirect | Via Google index; inherits Google's Reddit ranking |

- Reddit saw **450% growth** in AI search citations (2025 Mar-Jun)
- Domains with millions of Reddit/Quora mentions are **4x more likely** to be cited
- Posts with **50+ upvotes** cross the AI citation threshold — below this, rarely cited
- **TL;DR is the single most cited element** — AI engines use it as the default answer snippet
- Each numbered finding in a post can be independently extracted by AI as a citation unit
- Comparison tables are gold for AI extraction (structured data in unstructured platform)

### Post Structure → AI Citation Mapping

```
[Question-form title matching "What is the best X for Y?" pattern]
   → AI maps this to user queries

TL;DR: [2-3 sentences, standalone answer]          ← AI CITABLE UNIT #1
   → Perplexity extracts this verbatim

Key Findings (numbered):                            ← AI CITABLE UNIT #2-N
   1. [Claim] — [evidence]                          → Each independently citable
   2. [Claim] — [evidence]

Comparison table (if applicable):                   ← AI CITABLE UNIT (structured)
   → Gemini/ChatGPT extract tabular comparisons

"When we tested..." (1st-party experience)          ← Brand integration point
   → Natural, earned mention
```

### Tone Rules

| Do | Don't |
|----|-------|
| First person ("I tested", "In my experience") | Third person ("Brand X provides") |
| Acknowledge limitations ("This won't work if...") | Oversell |
| Mention alternatives including competitors | Only mention your product |
| Community-native vocabulary | Marketing jargon |
| Humble expert (share knowledge, not pitch) | Authority flex |

### Format: Value Post (300-800 words)

```
**Title**: [Question-form or insight-share] — match subreddit culture

**Body**:

**TL;DR** (top or bottom, 2-3 sentences)              ← AI CITABLE UNIT
[Standalone summary — must answer the title question without any other context]

## Context
[Why this matters — 2-3 sentences]

## Key Findings (numbered, each independently citable)  ← AI CITABLE UNITS
1. **Finding 1** — [data/evidence]
2. **Finding 2** — [data/evidence]
3. **Finding 3** — [data/evidence]

## My Experience
[1st-party: "When we tested..." / "After 6 months of..."]
[Brand touch here — natural, earned, not pitched]

## Takeaway
[Actionable conclusion]
[Discussion question: "What's your experience with...?"]

---
*Full analysis: [canonical URL] (in comment or here, per subreddit rules)*
*Updated [date]* ← recency signal for AI
```

### Title Formulas (Proven High-Engagement)

| Pattern | Example | GEO Value |
|---------|---------|-----------|
| Number + time period | "After 3 years of X, here's what I learned" | Experience signal |
| Contrarian hook | "Stop doing X — it's actually hurting your Y" | High engagement → upvotes |
| Specific result | "I went from A to B doing X — full breakdown" | Quantified claim for citation |
| Comparison | "X vs Y — I tried both for 30 days" | Matches comparison-stage queries |
| Question | "What's one thing you wish you knew before starting X?" | Matches foundational queries |
| PSA/TIL | "PSA: You can do X by doing Y" | How-to extraction by AI |

### Brand Integration (Reddit-Specific)

- **Brand density: 5-10%** — brand appears in "My Experience" section only, never in title or TL;DR
- **The Comment Strategy**: Brand mentioned in reply comments, not the post itself. "Someone asked what tool I used — it's [Brand]. Here's why..." This feels organic and avoids spam filters
- **90/10 rule**: 90% pure value, 10% brand-related across all your activity
- **3:1 ratio**: 3 helpful comments on others' posts for every 1 self-post
- **Never lead with brand**: The post is about the problem/solution, brand is incidental

### Subreddit Execution

| Rule | Detail |
|------|--------|
| Pre-posting research | Read sidebar + wiki + pinned posts before first post |
| Karma warmup | 2-4 weeks of comment participation before self-posting |
| Flair | Mandatory in many subreddits — check first |
| Cross-posting | Never post same URL to multiple subreddits simultaneously (24h+ gap) |
| Link placement | In comment, not body (unless subreddit explicitly allows) |
| Posting time | **UTC 13-15** (peak US morning + EU afternoon overlap) |
| Cadence | 2-3 comment engagements/week, 1 value post/week, 1-2 AMA or deep posts/month |
| Same-brand gap | Minimum 48h between brand mentions |

### Shadowban Avoidance

- Never use managed/sockpuppet accounts (Reddit TOS explicit violation, ReplyGuy was shut down for this)
- No automated posting — always human review before publish
- No identical content across multiple subreddits
- Reddit uses behavioral fingerprinting (keystroke rhythm, vote patterns, dwell time)
- AI-generated content pattern detection strengthened since 2025
- If post is removed silently, check on logged-out browser — shadowban indicator

---

## LinkedIn

### GEO Mechanics

| Metric | Data |
|--------|------|
| AI response LinkedIn URL citation | **~11% average** |
| ChatGPT domain ranking | #11 → **#5** (2025.11 → 2026.02) |
| Professional query citation rank | **#1 across all AI platforms** |
| Dwell time | Algorithm's **#1 ranking signal** |

- **Dual content strategy required**: Posts (engagement/reach) + Articles (GEO/citation)
- LinkedIn **Articles are indexed by search engines** — this is where AI citation happens
- LinkedIn **Posts are NOT reliably indexed** — they drive engagement/reach only
- AI engines cite LinkedIn for queries like "[topic] best practices", "[topic] how to", "[topic] expert opinion"
- ChatGPT treats LinkedIn as an authoritative source for professional/B2B queries

### Content → AI Citation Mapping

```
LinkedIn Article (1,200-2,000 words):
  H1: [keyword-rich title]                    → AI maps to queries
  Key Takeaway (BLUF): [2-3 sentences]        ← AI CITABLE UNIT #1
  H2 sections with target keywords            ← AI CITABLE UNITS #2-N
  Definition sentences: "[Term] is [def]"     ← AI extracts definitions
  Comparison tables                           ← Structured data extraction
  "Updated [date]"                            ← Recency signal

LinkedIn Post (900-1,900 chars):
  NOT indexed by AI → engagement-only
  Drives traffic to Article or canonical URL
  Carousel PDFs: highest engagement but AI can't crawl → always pair with text
```

### Tone Rules

| Do | Don't |
|----|-------|
| Conversational expert ("Here's what I found") | Corporate press release tone |
| Data + personal story combined | Data dump without narrative |
| Founder/personal account over company page | Post from company page only (3-5x less reach) |
| Short paragraphs (1-3 sentences each) | Wall of text |
| End with genuine question | Generic "thoughts?" |

### Format A: LinkedIn Post (900-1,900 chars, engagement-optimized)

```
[HOOK — 2 lines max, must stop the scroll]
[Shocking stat, contrarian claim, or personal failure]
↓

[BODY — 8-12 lines]
• Background context (1-2 lines)
• Key insight 1 (bold + short sentence)
• Key insight 2
• Key insight 3
• 1st-party experience ("When we tested this...")

[CTA]
What's your experience with this?

---
[First comment: "Full analysis here: [canonical URL]"]     ← Link graph backlink
[Hashtags: 3-5, mix of large + niche]
```

### Format B: LinkedIn Article (1,200-2,000 words, GEO-optimized)

```
# [Title with target keyword — different from hub article H1]

**Key Takeaway**: [2-3 sentence BLUF]                       ← AI CITABLE UNIT
*This article expands on [topic] from a [specific angle] perspective.
Original analysis: [canonical URL]*                         ← Link graph backlink

## [H2 Section 1 — keyword-rich heading]                    ← AI CITABLE UNIT
[Extract from hub + add personal professional context]
[Definition sentence: "[Term] is [definition]"]

## [H2 Section 2]                                           ← AI CITABLE UNIT
[Data visualization / comparison table]

## [H2 Section 3]
[Practical application for professionals]

## Conclusion
[1 actionable takeaway]
[Discussion question]

*Originally published at [canonical URL]*
*Updated [date]*                                            ← Recency signal
**Tags**: #Tag1 #Tag2 #Tag3
```

### Brand Integration (LinkedIn-Specific)

- **Brand density: 10-20%** — higher than Reddit because professional context makes brand mention natural
- Founder personal account is the primary vehicle (3-5x more reach than company page)
- Brand appears in experience-based claims: "At [Brand], we found that..."
- Company page reposts founder content (not the other way around)
- LinkedIn audience expects expertise claims — brand mention as credential is natural

### Critical Rules

| Rule | Detail |
|------|--------|
| External links in posts | **NO** — algorithm penalizes 40-50% reach reduction. Links in first comment only |
| External links in articles | OK — Articles don't get penalized |
| Carousel PDFs | Highest engagement (6.60%) but AI can't crawl → always pair with text article |
| Posting time | Tue-Thu 8-10am target timezone (60-90 min before audience peak) |
| Optimal post length | **900-1,900 characters** |
| Formatting | No markdown in posts. Unicode bold/italic sparingly. Line breaks as primary formatting. → or • for lists. 0-3 emojis max, never as bullet points |
| Author profile | Optimize headline + About section with expertise keywords — AI indexes these |
| Cadence | 3-5 posts/week (founder + company combined), 1 article biweekly |

---

## Threads

### GEO Mechanics

| Metric | Status |
|--------|--------|
| AI engine direct citation | **Minimal data** — Threads crawling/citation still limited |
| Indirect effect | Meta ecosystem (Instagram + Facebook) → brand signal amplification |
| Growth trajectory | 400M MAU, projected to surpass X by 2026 Q4 |
| Fediverse/ActivityPub | Opt-in available → future-proofing for decentralized crawlability |

**Honest assessment**: Threads has low direct GEO impact today. Its value is:
1. Brand awareness + community building
2. Content testing ground → what resonates here gets expanded to Reddit/LinkedIn
3. Meta indexing acceleration (eventual Google/AI crawlability expected)
4. Text Attachments with structured content may become indexable

### Tone Rules

| Do | Don't |
|----|-------|
| "I" voice — write as a person | Corporate "we are pleased to announce" |
| Hot takes, opinions, personality | Neutral, safe, fence-sitting |
| Casual humor, self-deprecation | Forced professionalism |
| Behind-the-scenes, process sharing | Polished final results only |
| Intentional imperfection (authentic) | Over-edited, AI-sounding |

### Format: Thread Series (5-7 posts, 50-300 chars each)

```
**Post 1 (Hook)**:
[Contrarian claim or surprising data point — 1-2 lines max]
"Did you know [shocking fact]?"

**Post 2-3 (Core Insight)**:
[Most interesting finding from hub article]
• Number/data-driven
• Short sentences, accessible language

**Post 4 (Brand Touch — natural)**:
"We actually tried this and..."
→ 1st-party experience, NOT an ad
→ If it feels like marketing, rewrite

**Post 5-6 (Actionable Takeaway)**:
"So what can you actually do?"
• 3 tips or checklist
• Immediately applicable

**Post 7 (CTA + Close)**:
"Full analysis in profile link"
+ Discussion prompt: "How are you handling this?"
```

### Brand Integration (Threads-Specific)

- **Brand density: <5%** — brand appears in at most 1 post out of 7 (Post 4)
- Frame as personal experience, never as promotion
- If it feels like an ad, cut it. Threads audience is allergic to marketing
- Profile bio carries the brand — posts don't need to

### Engagement & Algorithm Rules

| Rule | Detail |
|------|--------|
| Self-reply | **+42% engagement boost** — always add context in first reply |
| Carousel posts | **21.77% engagement** (3x other formats) |
| Reply chains | Primary algorithm signal — replies > likes |
| Optimal length | 500 chars ceiling, not target. Short and punchy wins |
| Text Attachment | 10,000 chars available for long-form (supports bold/italic/lists) |
| Topic tags | Exactly 1 per post (not traditional hashtags) |
| Link posts | **Lowest engagement (2.34%)** — put links in replies, NEVER in the main post body |
| Posting time | **Wed-Thu 7-9am** target timezone. Avoid weekends/evenings |
| 70/30 rule | 70% engaging with others' content, 30% your own posts |
| Instagram cross-post | Doesn't get algorithm bonus — optimize separately |
| Cadence | 3-5 posts/week, 1 hub-derived thread biweekly |
| No follow-for-follow | **Never do mutual follow/like exchanges.** Algorithm detects spam-like reciprocal patterns → suppresses reach |

### Content Mix (Weekly)

| Type | Share | Examples |
|------|-------|---------|
| Conversation starters (questions) | 30% | "A vs B?" / "Is it just me?" |
| Education/insights | 25% | Industry data, tips |
| Behind-the-scenes | 20% | Team life, process, mistakes |
| Hot takes / trend-jacking | 15% | Industry opinions, trend reactions |
| Memes / humor | 10% | Relevant memes, self-deprecating humor |

### Content Purpose Strategy (Reach vs Conversion)

Not all posts serve the same goal. Separate **reach content** from **conversion content** — mixing them kills both.

| Type | Goal | Expected Views | Placement |
|------|------|---------------|-----------|
| **Reach content** | Virality, new audience, engagement | High (algorithm-favored) | Regular feed posting |
| **Conversion content** | Drive signups, sales, profile visits | Low (normal — don't panic) | **Pin to profile** |

**Key rules:**

1. **Conversion content gets low views — that's normal.** Don't judge it by reach metrics. Pin it to your profile so profile visitors see it passively
2. **Attach CTA to reach content via replies, not in the post itself.** Write a viral-worthy post → add CTA as a sub-reply. This preserves the post's virality while capturing conversion from engaged readers
3. **Never put CTA + link + pitch in a single top-level post** — people won't engage, algorithm won't distribute

**Reference**: [@ai_hamzzi.mirra](https://www.threads.com/@ai_hamzzi.mirra) — executes this pattern well (reach posts with CTA in replies, conversion pinned to profile)

---

## Cross-Channel GEO: Link Graph & Citation Layer

### How Spokes Strengthen the Hub

| Citation Path | Mechanism | Impact |
|---------------|-----------|--------|
| Blog → Reddit TL;DR | Perplexity cites Reddit TL;DR as "default explanation" | **Direct citation (highest)** |
| Blog → LinkedIn Article | ChatGPT cites LinkedIn as professional authority | **Direct citation (high)** |
| Blog → Reddit comments | Brand mention frequency → semantic neighbourhood | **Indirect (medium)** |
| Blog → Threads | Meta indexing → brand signal amplification | **Indirect (low-medium)** |
| Cross-channel link graph | All spokes → canonical URL = link density increase | **Trust amplification (cumulative)** |

### Trust Spine Expansion

Each spoke adds to the brand's Trust Spine (the 5-10 core sources AI references):

1. **Canonical blog article** = 1st-party authoritative source
2. **Reddit value post** = community-validated source (upvotes = social proof)
3. **LinkedIn Article** = professional authority source
4. **3rd-party citations** = existing GEO strategy (media, reviews)

### Cross-Channel Consistency Matrix

| Element | Blog (Hub) | Reddit | LinkedIn | Threads |
|---------|-----------|--------|----------|---------|
| Brand message | Formal | Community-native | Expert | Casual |
| Core claim | Identical | Identical | Identical | Identical (simplified) |
| Data/statistics | Full | Key excerpts only | Visualized + interpreted | 1 highlight |
| Brand density | 15-25% | 5-10% | 10-20% | <5% |
| Citations/sources | Full inline | 1-2 key sources | 3-5 main sources | Omit or brief |
| CTA | Internal links | Discussion prompt | Comment prompt | Conversation starter |
| AI Citable Unit | BLUF + H2s | TL;DR + numbered findings | BLUF + H2s (article) | None (indirect) |

---

## Hub → Spoke Conversion Checklists

### Hub Article → Reddit

- [ ] Write TL;DR (2-3 sentences, must answer a query standalone — this is what AI extracts)
- [ ] Select target subreddit + verify rules (sidebar, wiki, pinned posts)
- [ ] Rewrite title — prefer "What is the best X for Y?" pattern (most cited by AI)
- [ ] Remove all marketing language — rewrite in community-native voice
- [ ] Reduce brand density to 5-10%, frame as 1st-party experience
- [ ] Add numbered findings (each independently citable by AI)
- [ ] Add comparison table if applicable (gold for AI extraction)
- [ ] Add canonical URL naturally (body or comment, per subreddit policy)
- [ ] Add discussion question at end
- [ ] Add "Updated [date]" for recency signal
- [ ] Verify: would this get upvoted if a random community member posted it?

### Hub Article → LinkedIn

- [ ] Choose format: Post (engagement) or Article (GEO) — usually both
- [ ] Post title ≠ Hub H1 ≠ Article title (all different, all keyword-targeted)
- [ ] Write BLUF/Key Takeaway (AI Citable Unit for the article)
- [ ] Extract 3-4 most relevant sections + add professional context
- [ ] Add definition sentences: "[Term] is [definition]" (AI extracts these)
- [ ] Include comparison table or structured data
- [ ] Create data visualization (carousel or inline)
- [ ] Set hashtags (3-5: 1 large + 2 medium + 1-2 niche)
- [ ] Place canonical URL in first comment (post) or body (article)
- [ ] Add "Updated [date]" for recency
- [ ] Decide: founder personal account or company page? (personal = 3-5x reach)

### Hub Article → Threads

- [ ] Select 1-2 most interesting insights from hub
- [ ] Write hook (first post must stop the scroll)
- [ ] Split into 5-7 thread posts (50-300 chars each)
- [ ] Rewrite entirely in casual tone (no copy-paste from hub)
- [ ] Brand touch limited to 1 post (Post 4), natural framing only
- [ ] Add self-reply to first post (+42% engagement)
- [ ] Set 1 topic tag per post
- [ ] End with profile link reference (not direct URL in post)

---

*Last updated: 2026-03-23*
