# Prompt Audit

Use this agent-only workflow for `/aeo prompts audit <domain_id>`. It is not a bare `aeo` CLI verb. Judge Prompt validity and observed utility in one read-only diagnosis; do not collapse them into one score.

## Input

Require exactly one `domain_id` UUID. Do not resolve a brand name, hostname, or active-domain fallback.

## Read live evidence

Run independent reads in parallel when possible:

```bash
aeo agent context -d <domain_id>
aeo strategy show -d <domain_id>
aeo prompts list -d <domain_id>
aeo content list --limit 200 -d <domain_id>
aeo measure visibility -d <domain_id>
aeo measure traffic --days 28 -d <domain_id>
aeo measure overview -d <domain_id>
```

Use `aeo measure content <content_id> -d <domain_id>` for articles that plausibly target the Prompt. Inspect Topic assignments through an authenticated read surface when available; otherwise state that limitation. Do not run a new visibility check or mutate data.

## Judge two concepts together

For each material Prompt or coherent Prompt group, make two separate judgments from the whole context:

- **Validity** — Is this a question the intended audience could genuinely ask, and can this brand credibly answer or be recommended without stretching its positioning or claims? Consider its distinct role within the Topic and whether another Prompt already covers the same intent.
- **Utility** — What has this Prompt actually helped reveal or cause? Consider visibility observations, competitor/category intelligence, owned-content lineage, GSC demand, GA outcomes, and how long evidence has had to mature.

Do not equate validity with current visibility. A valid Prompt can score zero. Do not call a Prompt useful merely because an article exists, or useless because a young article has little traffic. GSC is evidence of search demand, not AI Prompt volume, and correlation is not causation.

A Prompt may be valuable as a **research sensor** even when it should not drive an article. Preserve this distinction instead of forcing every tracked Prompt into the content queue.

## Recommend a role

Use best judgment; do not apply universal weights or hard thresholds.

- **content target** — credible brand fit and an actionable unanswered content opportunity
- **research sensor** — useful for monitoring category language, competitors, or recommendation behavior, but weak as a direct writing target
- **revise** — underlying intent matters, but the phrasing, scope, Topic placement, or claim burden is wrong
- **retire** — neither a credible target nor a useful sensor, or redundant with a better Prompt
- **insufficient evidence** — keep the judgment open when utility cannot yet be observed; say what is missing

## Output

Start with the portfolio-level conclusion, then show only the material groups and actionable exceptions. Do not dump every healthy Prompt.

```markdown
## Prompt Audit

**Portfolio conclusion:** ...

| Prompt or group | Validity | Observed utility | Recommended role | Why |
|---|---|---|---|---|
| ... | ... | ... | ... | ... |

**What to change now**
- ...

**Evidence limits**
- ...
```

Name Prompts in user-facing copy; do not expose UUIDs. Stop after the read-only recommendation. Prompt edits, Topic reassignment, tracking changes, or new visibility checks require separate explicit confirmation under the CUD Rule.
