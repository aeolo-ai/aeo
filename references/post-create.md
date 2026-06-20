# Channel Post Creation

Write platform-native social media posts optimized for both human engagement and AI engine citation.

**Core principle**: "This person following best practices would have written it like this." — platform structure is the goal, user voice is the constraint.

---

## /aeo post write — Write a channel post

Agent writes the post directly using the guidelines below + [channel-washing.md](channel-washing.md) platform rules. After writing, import via `/aeo post import`.

---

### Pre-flight Checklist

**Mandatory gate before writing a single word.**

```
## Post Brief

### Platform
- [ ] Target platform: [threads / linkedin / reddit]

### Source Material
- [ ] Entry path: [hub-to-spoke / original / data-insight]
- [ ] If hub-to-spoke: content ID or article URL
- [ ] Topic / angle

### Brand Context
- [ ] Brand profile loaded (`/aeo domain brand`)
- [ ] Writing style profile loaded from `writing_styles` when available
- [ ] Voice examples loaded from `brand_voice_examples` when available
```

#### Entry Paths

| Path | When | Input |
|------|------|-------|
| **Hub-to-spoke** | Article exists, distribute to channel | `/aeo content get <id>` → pick angle → wash for platform |
| **Original** | No article, standalone channel post | Topic + brand context |
| **Data-insight** | Stat, finding, or insight to frame | Data point + context |

If brand profile not loaded → fetch via `/aeo domain brand` first.
If no writing style profile is available → ask user (same gate as content-create.md Step 1.5). Treat legacy Tone & Voice notes in `brand_context` as fallback context only.

---

### Step 1 — Load Platform Rules

Read **[channel-washing.md](channel-washing.md)** and extract the rules for the target platform. The platform-specific sections are the source of truth for:

- Tone rules (Do / Don't table)
- Post format & structure template
- Brand density limit
- Engagement & algorithm rules
- GEO mechanics (how AI engines crawl this platform)

**Do NOT skip this step or rely on memory.** The rules are specific and detailed.

| Platform | channel-washing.md Section | Brand Density | Key Format |
|----------|---------------------------|---------------|------------|
| Threads | "## Threads" | **<5%** | Thread series 5-7 posts, 50-300 chars each |
| LinkedIn | "## LinkedIn" | **10-20%** | Post (900-1,900 chars) OR Article (1,200-2,000 words) |
| Reddit | "## Reddit" | **5-10%** | Value post 300-800 words with TL;DR |

---

### Step 1.5 — Load Voice Examples (few-shot)

Load voice examples for the target platform from the API:

```bash
aeo post examples --platform {platform}
```

This returns GOOD, BAD, and Reference examples stored in the DB. Extract:
- **GOOD examples** — these are the few-shot exemplars. Match their tone, sentence structure, energy, and vocabulary when writing.
- **BAD examples** (anti-patterns) — if your draft resembles any of these, rewrite.
- **Reference examples** — techniques borrowed from benchmark accounts.

If no examples are returned, skip this step. The post will rely on the writing style profile and any legacy voice notes only. Suggest running `/aeo post analyze --url <channel_url>` to generate voice examples.

**How to use the examples during writing:**
- Read each GOOD example before writing. Internalize the rhythm, word choice, and energy.
- Do NOT copy them — write new content that *sounds like* them.
- After writing, compare your draft against the GOOD examples. Would they feel like they came from the same person?
- Check against BAD examples. If your draft has any of those patterns (corporate tone, feature lists, discount codes), rewrite.

---

### Step 2 — Determine Content Archetype

Not every post is a hub-to-spoke conversion. Pick the archetype that fits:

| Archetype | Description | Best For |
|-----------|-------------|----------|
| **Hub-to-spoke** | Adapt a hub article's angle for the platform | After publishing a blog article |
| **Value-share** | Share a useful insight, tip, or experience | Building authority + engagement |
| **Commentary** | React to industry news/trend with an opinion | Hot takes, trend-jacking |
| **Event recap** | Share experience from an event, meetup, race | Behind-the-scenes, community |
| **Question/poll** | Ask a genuine question to spark discussion | Conversation starters |
| **Data-drop** | Lead with a surprising stat or finding | Education, credibility |

For hub-to-spoke: read the source article first (`/aeo content get <id>` or provided URL), then select 1-2 most interesting angles — do NOT summarize the whole article.

---

### Step 3 — Write

Follow the platform's format template from channel-washing.md strictly. Apply these universal rules:

#### Universal Rules (from channel-washing.md)

1. **Single language only** — use the language specified in the brief (default: English). Never mix languages within a post. No code-switching, no untranslated phrases, no foreign-language flourishes. If the brand voice includes bilingual elements, ignore them for channel posts — platform nativeness requires linguistic consistency
2. **AI Citable Unit required** — every post must contain a self-contained 2-3 sentence block that answers a query without surrounding context. This block must be **definitive** (no hedging with "maybe", "roughly", "perhaps") and **extractable** (an AI engine could quote it verbatim as an answer). Examples:
   - TL;DR for Reddit
   - BLUF for LinkedIn
   - One core-insight post for Threads (typically Post 2) — write it as a standalone fact, not a conversational aside
3. **Core claim consistency** — if adapting from hub article, the factual claim is identical; only framing changes
4. **1st-party experience framing** — "I/we tested" not "Brand X offers"
5. **Never copy-paste** between channels or from the hub article — each must feel native
6. **Brand density per platform** — Threads <5%, Reddit 5-10%, LinkedIn 10-20%

#### Platform-Specific Writing Gates

**Threads:**
- [ ] **Single language** — no mixing. English thread = all English. Korean thread = all Korean
- [ ] "I" voice — write as a person, NOT corporate
- [ ] Hot takes, opinions, personality — NOT neutral/safe
- [ ] Brand appears in at most 1 post out of 5-7 (Post 4 position)
- [ ] If it feels like marketing → rewrite
- [ ] Self-reply planned for Post 1 (+42% engagement boost)
- [ ] Each post 50-300 chars (500 chars ceiling)
- [ ] No links in posts — link in profile or in a sub-reply (NEVER in main post body)
- [ ] **1 topic tag per post** — e.g. `#running`, `#skincare` (must be present on every post, not optional)
- [ ] **Last post references profile link** — "Link in bio" or "Full breakdown in profile" (drives traffic without algorithm penalty)
- [ ] **Reach vs conversion**: Is this a reach post or conversion post? Reach → regular feed. Conversion → pin to profile. For reach posts, attach CTA as a sub-reply to preserve virality
- [ ] **No follow-for-follow** — never do mutual follow/like exchanges (algorithm suppresses reach)

**LinkedIn:**
- [ ] Choose format: Post (engagement) OR Article (GEO) — often both
- [ ] Post: hook in first 2 lines (must stop the scroll), 900-1,900 chars
- [ ] Article: BLUF/Key Takeaway at top, 1,200-2,000 words, H2 sections
- [ ] NO external links in posts (40-50% reach penalty) — links in first comment
- [ ] External links OK in articles
- [ ] End with genuine question (not generic "thoughts?")
- [ ] Founder/personal account voice, not company page tone
- [ ] 3-5 hashtags (1 large + 2 medium + 1-2 niche)
- [ ] No markdown formatting in posts — Unicode bold/italic sparingly, line breaks for structure

**Reddit:**
- [ ] TL;DR at top or bottom (2-3 sentences, standalone answer) — this is what AI extracts
- [ ] Title in question form: "What is the best X for Y?" pattern
- [ ] Numbered findings (each independently citable by AI)
- [ ] Brand appears in "My Experience" section only, never in title or TL;DR
- [ ] Acknowledge limitations, mention alternatives including competitors
- [ ] Community-native vocabulary, no marketing jargon
- [ ] Add canonical URL (in body or comment, per subreddit rules)
- [ ] Add "Updated [date]" for recency signal

---

### Step 4 — Self-Review

Before presenting to the user, check:

| Check | Pass Criteria |
|-------|--------------|
| **Language purity** | Single language throughout? Zero mixed-language phrases? |
| Platform nativeness | Would a regular user of this platform post this? Not an ad? |
| Brand density | Within platform limit? Brand not in title/hook? |
| AI Citable Unit | Is there a **definitive** 2-3 sentence block (no hedging) an AI could extract verbatim? |
| 1st-party framing | "I/we" not "Brand X"? |
| Marketing smell | Zero discount codes, CTAs to buy, product feature lists? |
| Tone match | Matches the user's writing style profile and approved voice constraints? |
| **Voice examples match** | Does this sound like the GOOD examples? Would it feel like the same person wrote it? |
| **Anti-example check** | Does this resemble any BAD examples? (corporate tone, feature lists, discount codes, generic AI phrasing) |
| Format compliance | Follows the exact structure template from channel-washing.md? |
| **Topic tags present** | Every post has exactly 1 topic tag? (Threads/LinkedIn) |
| **Profile link reference** | Last post mentions "link in bio" / "full breakdown in profile"? (Threads) |

**If 2+ checks fail**, revise before showing to the user. Language purity failure alone is grounds for full rewrite.

---

### Step 5 — Present & Import

Show the post to the user with:
1. The full post content
2. Platform + archetype used
3. Any notes (e.g., "first comment should contain the canonical URL")

On user approval, import:

```
aeo post import --platform <platform> --body "<content>"
```

Optional flags:
- `--title "<title>"` (Reddit: post title)
- `--post-type <type>` (value_post, commentary, event_recap, question, data_drop, hub_to_spoke)
- `--content-id <id>` (link to source hub article)
- `--channel-id <id>` (target channel, from `/aeo domain channels`)

---

## Quick Reference: Platform Comparison

| | Threads | LinkedIn | Reddit |
|---|---------|----------|--------|
| **Primary GEO value** | Brand signal amplification (indirect) | Professional authority citation | Direct AI citation (Perplexity 46.7%) |
| **Format** | Thread series 5-7 posts | Post (900-1,900 chars) + Article (1,200-2,000 words) | Value post 300-800 words |
| **Brand density** | <5% | 10-20% | 5-10% |
| **Voice** | Personal, casual, opinionated | Conversational expert | Humble community member |
| **Links** | In profile only | In first comment (post) or body (article) | In body or comment per subreddit rules |
| **AI Citable Unit** | Hook + core insight | BLUF / Key Takeaway | TL;DR (verbatim extracted by Perplexity) |
| **Posting time** | Wed-Thu 7-9am | Tue-Thu 8-10am | UTC 13-15 |
| **Cadence** | 3-5 posts/week | 3-5 posts/week + 1 article biweekly | 1 value post/week + 2-3 comments |

---

*For detailed platform rules, engagement mechanics, and hub-to-spoke conversion checklists, see [channel-washing.md](channel-washing.md).*
