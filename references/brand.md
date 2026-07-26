# Brand & Prompts

---

## /aeo agent context — Show current agent context

```bash
aeo agent context
```

Response: `text/markdown` — the default brand operating context used by the agent, including brand snapshot, brand notes, content strategy, source policy, publishing channels, measurement status, provenance, and warnings.

If empty or JSON error, suggest running setup and checking the domain ID.

---

## /aeo domain brand update — Update brand context fields

1. Fetch current Agent/Brand Context with `aeo agent context` first (show it to the user)
2. Ask the user what they want to change
3. Confirm before writing

Accepted fields:

| Field | Type | Notes |
|-------|------|-------|
| `name` | string | Brand display name |
| `category` | string | e.g. `"B2B SaaS"` |
| `industry` | string | e.g. `"Developer Tools"` |
| `value_proposition` | string (max 2000) | Core positioning statement |
| `key_features` | string[] (max 20) | Feature list for brand mentions |
| `primary_language` | ISO 639-1 | e.g. `"en"`, `"ko"`, `"ja"` |
| `brand_context` | string (max 50000) | Free-form brand positioning and durable notes (see template below) |

```bash
aeo domain brand update --name="..." --category="..." --value-proposition="..."
```

> Canonical form is `aeo domain brand update` (matches `aeo domain` help). Bare `aeo brand update` routes only on the **dashboard-chat / MCP agent surface** (the connector registry normalizes it to `domain brand update`); the raw `aeo` terminal binary does not accept it, so always use `aeo domain brand update` in shell examples.

Partial update — unset fields are preserved.

### brand_context template

`brand_context` is free-form markdown for durable brand facts, positioning,
audience, narratives, and constraints. Do not store reference-analysis dumps or
one-off voice examples here; those should stay with the task/reference analysis
that produced them. Only durable tone constraints that should affect all future
work belong here.

Suggest this structure when helping a user build it from scratch:

```markdown
## Brand Overview
[Brand mission, positioning, and what makes it different]

## Target Audience
[Who the brand serves — personas, pain points, jobs-to-be-done]

## GEO Strategy
[Which AI engines to prioritize, content angles to emphasize, competitive positioning]

## Key Narratives
[3–5 core messages the brand wants AI engines to associate with]

## Competitive Context
[Main competitors, how to frame comparisons, where the brand wins]

## Constraints
[Compliance requirements, claims to avoid, source preferences, positioning guardrails]
```

---

## /aeo prompts list — List prompts grouped by stage

```bash
aeo prompts list
aeo prompts list --status tracked
```

> Use the explicit `list` verb. Bare `aeo prompts` prints sub-help in the terminal binary (it only lists via the agent surface), so `aeo prompts list` is the form that works everywhere.

Response: `text/markdown` — table grouped by stage (foundational → comparison → use-case → implementation), showing language, query form, prompt text, visibility score, and last checked date.

Optional filter: `--status tracked|untracked`. Omit it to list both statuses.

After displaying:
- Note which stages are sparse or have low visibility scores
- Suggest adding prompts where coverage is thin
- If no prompts exist, prompt the user to add foundational ones first

---

## /aeo prompts generate — Generate and save a prompt set

```bash
aeo prompts generate
aeo prompts generate --count 12 --languages en,ko --instruction "Prioritize comparison prompts."
```

Generates a stage-balanced CEP prompt set from the active domain's saved brand context and immediately persists the result as tracked Prompts. This is free, but it is still a write action: show the requested languages/count/instruction and get explicit confirmation before running it.

Accepted flags:

| Flag | Type | Notes |
|------|------|-------|
| `--count` | positive integer | Generation hint, not a hard cap; the server keeps the set stage-balanced |
| `--languages` | comma-separated language codes | Omit to use the domain's primary language |
| `--instruction` | string | Extra generation guidance |

The terminal binary calls the dashboard's authenticated `POST /score/prompts` REST endpoint directly. The agent command uses the connector surface; both reach the same server-side generation and persistence implementation.

If the command reports `Run brand analysis first`, complete Brand Understanding before retrying. After success, review the generated set with `aeo prompts list` before starting a visibility check.

---

## /aeo prompts health — Audit the tracked prompt set

```bash
aeo prompts health
```

Free, read-only. The topic engine optimises **within** the tracked set, so a wrong set makes it optimise faithfully toward the wrong target — and nothing downstream reports that, because every measurement is taken against the same set. This audits the set itself.

Returns:

| Section | What it answers |
|---------|-----------------|
| What is working | Theme-level verdicts (several prompts sharing a word), not per-prompt noise. Stays silent when too little content exists to judge. |
| Biggest openings | Queries where engines DO recommend a brand and it is not us — with how crowded the answer is and where we rank on Google |
| Wasted effort | Prompts with several articles and still no citation; brand-name prompts holding tracked slots |
| Demand you are not tracking | Real search demand with no prompt behind it, phrasings folded together and summed |
| Every tracked prompt | Status + recommended action |

Statuses:

| Status | Meaning | Action |
|--------|---------|--------|
| `STUCK_HARD` | Tried repeatedly, never cited — but competitors ARE recommended here | **Change the approach**, do not retire. The seat exists; the format or angle failed |
| `STUCK_DEAD` | Tried, never cited, and the engines recommend nobody | Retire — there is no list to be on |
| `UNTRIED` | No article targets it yet | Prioritise. A zero here is untouched, not failed |
| `RISING` / `WINNING` | Citation climbing or already held | Keep |
| `HOLDING` / `UNMEASURED` | Not enough evidence yet | Keep, revisit |

The distinction that matters: a prompt at 0% with five articles behind it and one at 0% with none look identical in the numbers and need opposite treatment.

---

## /aeo prompts add — Add manual prompts

Ask the user for the prompt details, then run:

```bash
aeo prompts add --prompt="best project management tools" --stage=comparison --language=en --segment foo,bar
```

### Adding several at once

**Whenever you have more than one prompt, send them in a single call** — pass a JSON array instead of calling `add` repeatedly (max 30 per call):

```bash
aeo prompts add --prompts-json='[
  {"prompt":"best project management tools","stage":"comparison"},
  {"prompt":"what is a kanban board","stage":"foundational"},
  {"prompt":"how do I migrate off jira","stage":"implementation"}
]'
```

Entries may be objects or plain strings. Per-item keys win; the top-level `--stage`/`--language`/`--query-form`/`--segment` flags fill in whatever an item omits:

```bash
aeo prompts add --language=ko --stage=comparison --prompts-json='["노션 vs 에어테이블","먼데이 vs 아사나"]'
```

Accepted flags (binary `aeo prompts add`):

| Flag | Type | Required | Default | Example |
|------|------|----------|---------|---------|
| `--prompt` | string | ✅ unless `--prompts-json` | — | `"best project management tools"` |
| `--prompts-json` | JSON array of objects or strings, max 30 | ✅ unless `--prompt` | — | `'[{"prompt":"...","stage":"comparison"}]'` |
| `--stage` | `foundational` \| `comparison` \| `use-case` \| `implementation` | — | `foundational` | `comparison` |
| `--language` | `en` \| `ko` \| `ja` \| `zh` \| `ar` | — | `en` | `zh` |
| `--query-form` | `short-tail` \| `long-tail` \| `conversational` | — | `conversational` | `long-tail` |
| `--segment` | comma-separated tags | — | — | `foo,bar` |

> Use `--prompts-json`, never a comma-separated list of prompts — prompt text routinely contains commas ("best CRM for startups, 2026") and would be shredded.

Batch adds are best-effort: valid prompts are inserted, and the response reports anything skipped with a reason (duplicate, too short/long, invalid stage). Duplicates are matched both against prompts already tracked on the domain and within the batch itself, so re-running a batch will not stack copies. Fix only the reported items and re-send those.

Confirm the details with the user before submitting. After success, suggest running `/aeo visibility check run` to measure visibility for the new prompts.

---

## /aeo prompts update — Edit an existing prompt

```bash
aeo prompts update <promptId> --prompt="updated text" --stage=use-case
```

Accepted flags (all optional, at least one required):

| Flag | Type | Notes |
|------|------|-------|
| `--prompt` | string | Single Prompt text used immediately for display, dedupe, writing, and visibility measurement |
| `--stage` | `foundational` \| `comparison` \| `use-case` \| `implementation` | Move to different stage |
| `--query-form` | `short-tail` \| `long-tail` \| `conversational` | Update query form (kebab-case) |
| `--segment` | comma-separated tags | Replaces the prompt's tags |
| `--status` | `tracked` \| `untracked` | Toggle measurement tracking |

Use `/aeo prompts list` first to get the prompt ID. Confirm with user before updating.

After a text update, the response returns the stored Prompt text and its new revision. There is no separate canonical/localized text to update.

---

## /aeo prompts delete — Delete a prompt (soft delete)

1. Run `/aeo prompts list` to show current prompts with IDs
2. Ask the user to confirm which prompt(s) to delete
3. Show the list and ask "Proceed?" before calling DELETE

```bash
aeo prompts delete <promptId>
```

Response: `{ success: true, deleted: { id, canonical } }`

- Soft delete — sets `deleted_at`, data is preserved
- `404` if prompt not found or already deleted
- For bulk deletion (multiple prompts), confirm the full list once then loop calls sequentially
