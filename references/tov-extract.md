# Tone of Voice Extraction — Channel Analysis

Crawl a single URL, analyze the posts, and extract a Tone of Voice profile.
Store concrete GOOD/BAD/reference samples in `brand_voice_examples`. Keep any
structured voice rules as proposed replacement Voice profile material; do not
write new `writing_styles` records. Treat broad Tone & Voice notes in
`brand_context` as temporary context only. Run multiple times with different
URLs to build a complete picture.

---

## /aeo post analyze --url <URL>

One URL per invocation. Works for both your own channels and reference/benchmark accounts.

```
/aeo post analyze --url https://www.threads.com/@yourbrand.official    ← own channel
/aeo post analyze --url https://www.linkedin.com/company/yourbrand/    ← own channel
/aeo post analyze --url https://www.threads.com/@competitor.account    ← reference
```

### When to Use

- First time setting up a brand's social content pipeline
- Quarterly refresh (tone evolves as the brand matures)
- Benchmarking a competitor or reference account

---

### Step 1 — Detect Platform & Crawl

Detect the platform from the URL, then crawl with **gstack browse**. Social platforms lazy-load content — you must scroll to load more posts.

```bash
B=$(git rev-parse --show-toplevel 2>/dev/null)/.claude/skills/browse/dist/browse
$B goto <URL>

# Scroll to load more posts — repeat until no new content appears
$B scroll
# wait 2s for content to load
$B scroll
# wait 2s
$B scroll
# wait 2s

# Now extract all text
$B text
```

**Scroll loop**: Keep scrolling until the text output stops growing or you hit 15+ posts. Social platforms (Threads, LinkedIn) load 4-6 posts initially and need 3-5 scrolls to reach 15-20. If text length doesn't increase after a scroll, you've hit the bottom or a login wall.

**Verification**: After crawling, count the posts in the output. If fewer than 10, scroll more. If fewer than 5 after max scrolling, note "limited data — login wall" and proceed with what's available.

| URL Pattern | Platform | Data Available |
|------------|----------|----------------|
| `threads.com/@handle` | Threads | Posts, likes, replies, reposts, bio |
| `linkedin.com/company/name/` | LinkedIn (company) | Posts, likes, comments, about |
| `linkedin.com/in/name/` | LinkedIn (personal) | Recent posts, headline, about |

**Collect:**
- All visible posts (minimum 10, aim for 15-20+)
- Engagement numbers (likes, replies, reposts) per post
- Bio/about section
- Topic tags/hashtags used

---

### Step 2 — Determine Context

Is this the brand's own channel or a reference account?

| Signal | Own Channel | Reference |
|--------|------------|-----------|
| URL matches a registered channel (`/aeo domain channels`) | Own | — |
| User specified "reference" or "benchmark" | — | Reference |
| URL is for a different brand | — | Reference |
| Unclear | Ask the user | — |

This determines how the analysis is used in Step 4.

---

### Step 3 — Analyze Patterns

Extract these dimensions from the crawled posts:

#### A. Voice Attributes

| Dimension | What to look for | Example values |
|-----------|-----------------|----------------|
| **Person** | 1st person (I/we) vs 3rd person vs impersonal | "I tested" vs "Brand offers" vs "Studies show" |
| **Formality** | Casual → balanced → formal | Casual / Balanced / Formal |
| **Energy** | Low-key → moderate → high-energy | "let's gooo!!!!" = high |
| **Sentence length** | Short punchy vs long narrative | Avg chars per sentence |
| **Emoji usage** | None → sparse → heavy | Count per post average |
| **Language** | Single language only | English / Korean (mixing discouraged) |
| **Humor** | None → dry/subtle → overt | Self-deprecating, industry jokes |
| **Authority stance** | Humble learner → peer → expert | "I'm still figuring this out" vs "After 10 years..." |

#### B. Content Patterns

| Dimension | What to look for |
|-----------|-----------------|
| **Topic distribution** | What % is product, lifestyle, education, community, personal? |
| **Post types** | Questions, hot takes, tips, behind-the-scenes, event recaps? |
| **CTA style** | "What do you think?" vs "Link in bio" vs no CTA |
| **Brand mention frequency** | How often does brand name appear? In what context? |
| **Engagement correlation** | Which post types get most replies? Most reposts? |

#### C. Platform-Specific (extract what applies)

**Threads:** Thread series vs single posts ratio, self-reply usage, topic tag patterns
**LinkedIn:** Post vs Article ratio, hook line style, hashtag strategy, company vs personal

---

### Step 4 — Select Example Posts

From the crawled posts, select the **top 3-5 by engagement** as GOOD examples. These are the few-shot exemplars that `post write` will use to match tone.

For each selected post, record:
- Original text (verbatim)
- Engagement numbers
- Why it worked (1 line)

Then generate 1-2 **BAD examples** — rewrite one of the GOOD posts in generic AI tone to show what to avoid. This teaches the LLM the contrast between authentic brand voice and generic output.

---

### Step 5 — Save Voice Examples via API

Save each example to the DB via CLI. This makes examples accessible to all agents (local, Fly.io, etc.).

```bash
# GOOD example (high engagement post)
aeo post examples add \
  --platform threads \
  --type good \
  --body "We are invited to activate + promote padel at Coachella..." \
  --source-url "https://www.threads.com/@yourbrand.official" \
  --note "Community CTA — 936 likes, highest engagement"

# BAD example (rewrite in generic AI tone to show contrast)
aeo post examples add \
  --platform threads \
  --type bad \
  --body "Our SPF 50 Sunstick is specifically designed for athletes..." \
  --note "Corporate tone, feature list, discount code — everything to avoid"
```

For reference accounts:
```bash
aeo post examples add \
  --platform threads \
  --type reference \
  --body "{competitor's high-engagement post verbatim}" \
  --source-url "https://www.threads.com/@competitor" \
  --note "Borrowed technique: question-based CTA drives 5x replies"
```

Verify saved examples: `aeo post examples --platform threads`

---

### Step 6 — Propose Voice Updates

Load existing `brand_voice_examples` and any `brand_context` Tone & Voice
section. Then propose updates to the voice-specific stores:

**If own channel:**
- Add concrete GOOD/BAD samples to `brand_voice_examples`.
- Return Core Voice + Channel Modifiers as proposed replacement Voice profile
  material, not as a `writing_styles` patch.

**If reference account:**
- Add reference techniques as proposed `brand_voice_examples` with `type=reference`.
- Keep benchmark notes out of `brand_context` unless they affect durable brand positioning.

#### Replacement Voice Profile Candidate (compact — no examples here)

```markdown
## Core Voice
- **Person**: [extracted]
- **Energy**: [extracted]
- **Authority**: [extracted]
- **Humor**: [extracted]
- **Language**: [extracted]

## Vocabulary DNA
- **Use**: [5-10 phrases]
- **Avoid**: [phrases]

## Channel Modifiers
### Threads
- Formality: [extracted]
- Post style: [extracted]
- Best performing: [type]
### LinkedIn
- Formality: [extracted]
- Post style: [extracted]
- Best performing: [type]
```

---

### Step 7 — Confirm & Save

1. Show the user: proposed Voice profile candidate + proposed `brand_voice_examples`
2. Ask: "Does this capture your brand's voice? Anything to adjust?"
3. On approval:
   - Interactive CLI/operator flow: save concrete examples through the voice examples management surface when available
   - Background writing job or chat flow: do not write product memory directly;
     return reviewed `brand_voice_examples` patches and keep structured profile
     material pending until the replacement Voice profile exists

---

## Multiple Runs

```
Run 1: /aeo post analyze --url https://threads.com/@mybrand     → Core Voice + Threads modifier + Threads examples
Run 2: /aeo post analyze --url https://linkedin.com/company/x   → adds LinkedIn modifier + LinkedIn examples
Run 3: /aeo post analyze --url https://threads.com/@competitor   → adds Reference Benchmark + reference examples
```

Each run appends reviewed examples to `brand_voice_examples` and refines the
replacement Voice profile candidate.

---

*For channel-specific writing rules, see [channel-washing.md](channel-washing.md).*
*For post generation workflow, see [post-create.md](post-create.md).*
