# Voice Evidence Extraction — Channel Analysis

Crawl or analyze a single URL and extract task-specific tone / format evidence.
Do not write deprecated voice/style records.
The output should stay attached to the analysis result as reusable reference
evidence, then be selected explicitly by the user when a writing/generation task
needs it. Do not merge analysis into `brand_context` automatically.

---

## /aeo post analyze --url <URL>

One URL per invocation. Works for both your own channels and reference/benchmark accounts.

```
/aeo post analyze --url https://www.threads.com/@yourbrand.official --mode owned        ← own channel
/aeo post analyze --url https://www.linkedin.com/company/yourbrand/ --mode owned        ← own channel
/aeo post analyze --url https://www.threads.com/@competitor.account --mode reference    ← reference
```

**Flags:**
- `--url` (required) — the channel/reference URL to analyze
- `--provider blog|threads|tiktok|instagram` — override platform detection
- `--mode owned|reference` — declare whether this is your own channel or a reference/benchmark (drives how the evidence is used in Step 2 below)
- `--limit` — cap the number of posts analyzed

### When to Use

- First time setting up a brand's social content pipeline
- Quarterly refresh (tone evolves as the brand matures)
- Benchmarking a competitor or reference account

---

### Step 1 — Queue Aeolo Analysis

Run `aeo post analyze` with one user-provided URL. Aeolo's background job owns
URL collection and analysis; do not invoke a local browser helper, automate
Chrome, or manually scroll and extract page text as a substitute.

```bash
aeo post analyze --url <URL> --mode owned|reference [--provider <provider>] [--limit <n>]
```

Keep the returned job ID. After the job completes, read the stored result with
`aeo reference style --url <URL>` (add `--provider` when URL inference is not
enough). If Aeolo reports insufficient evidence or an access restriction, state
that limitation instead of attempting a client-side login-wall workaround.

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

Is this the brand's own channel or a reference account? Express the decision with `--mode owned|reference` on the `post analyze` call.

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

From the crawled posts, select the **top 3-5 by engagement** as task-selected references. These are candidates for task-specific few-shot context; the writing task should use only reviewed examples relevant to the current language and task.

For each selected post, record:
- Original text (verbatim)
- Engagement numbers
- Why it worked (1 line)

Then generate 1-2 **BAD examples** — rewrite one of the GOOD posts in generic AI tone to show what to avoid. This teaches contrast without making the examples universal style rules.

---

### Step 5 — Produce Reference Evidence

Return a compact style brief that can be attached to the analysis result and selected later by the user.

```markdown
## Style Brief
- Tone:
- Format:
- Hook:
- Rhythm:
- Do:
- Don't:
- Source URL:
```

This is evidence, not global brand memory. Keep it scoped to the analyzed URL.

---

### Step 6 — Propose Optional Durable Updates

Only propose durable brand-context changes if the analysis reveals a rule that should affect every future task for the brand.

**If own channel:**
- Return Core Voice + Channel Modifiers as proposed notes.
- Ask the user before turning any broad rule into `brand_context`.

**If reference account:**
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

### Step 7 — Confirm Usage

1. Show the user the style brief and representative evidence.
2. Ask whether this reference should be used for the current task.
3. Do not write product memory directly from this flow.

---

## Multiple Runs

```
Run 1: /aeo post analyze --url https://threads.com/@mybrand     → Core Voice + Threads style brief
Run 2: /aeo post analyze --url https://linkedin.com/company/x   → LinkedIn style brief
Run 3: /aeo post analyze --url https://threads.com/@competitor   → Reference benchmark brief
```

Each reviewed run should remain a selectable analysis/reference result. Promote
only durable, brand-wide rules into `brand_context` after explicit review.

---

*For channel-specific writing rules, see [channel-washing.md](channel-washing.md).*
*Post generation itself is retired (2026-07-27) — [post-create.md](post-create.md) is reference only. This analysis feeds blog-article briefs.*
