# Prompt Audit

Use this agent-only workflow for `/aeo prompts audit <domain_id>`. It is not a bare `aeo` CLI verb. Evaluate Validity and Utility separately for each tracked Prompt category, then give a natural-language operating recommendation. Do not collapse the judgments into a score or fixed decision enum.

## Input

Require exactly one `domain_id` UUID. Do not resolve a brand name, hostname, or active-domain fallback.

## Read live evidence

Run independent reads in parallel when possible:

```bash
aeo agent context -d <domain_id>
aeo strategy show -d <domain_id>
aeo prompts list --status tracked -d <domain_id>
aeo content list --limit 200 -d <domain_id>
aeo measure visibility -d <domain_id>
aeo measure traffic --days 30 -d <domain_id>
aeo measure overview -d <domain_id>
```

Use `aeo measure content <content_id> -d <domain_id>` for live articles that plausibly target a tracked category. Inspect Topic assignments through an authenticated read surface when available. Use actual Topics as the category boundary; use segment tags or a clearly stated semantic grouping only as a fallback.

Do not use untracked Prompts in either judgment. If none are tracked, say that the portfolio cannot be audited and stop. Do not run a new visibility check or mutate data.

## Build category evidence

Group the tracked Prompts before judging. For each category, collect:

- tracked Prompt count, representative questions, latest visibility, and movement over time when history exists
- live owned articles that plausibly answer those tracked questions
- article-level GSC and GA evidence when connected
- publication age, duplicate/cannibalizing coverage, and measurement gaps

Before category judgment, split the portfolio into discovery Prompts and branded
diagnostics. A Prompt is self-included when it contains the customer brand name,
domain, product name, or a recognizable variant. The default tracked portfolio
is a discovery sensor, so treat self-included Prompts as misplaced even when
they sound natural or would produce a useful answer. Do not use their results as
evidence of discoverability, Share of Voice, competitive preference, category
validity, or content utility. Report the contamination explicitly and recommend
removing or untracking those Prompts. Only evaluate them as a separate branded
diagnostic surface when the user explicitly says that such a set was intended.

Treat Prompt-to-article matching as an internal lineage estimate, not the result itself. Count only a verified owned canonical URL or deploy fact as live. Exclude drafts, review/approved-only items, PBN publication, and articles that target only untracked Prompts. If an article lives on a host not covered by the connected GSC/GA property, call it unmeasurable rather than failed.

## Judge two concepts separately

### Validity

Judge whether the category's tracked Prompts form a credible sensor for this brand: real-question plausibility, brand answerability, Topic fidelity, distinctness, non-leading phrasing, claim safety, and sufficient context. Visibility and article performance do not determine Validity.

### Utility

Judge whether writing for the category appears to be working. Read together:

- AI visibility across only the category's tracked Prompts
- GSC demand and outcomes for live category articles
- GA engagement or conversion evidence for those live articles
- publication age and whether multiple articles confound attribution

Separate the observed outcome by surface. A category may work in Google search while showing no AI visibility movement. Do not call a category successful merely because content exists, or failed because a young article has little traffic.

Without a saved pre-publication baseline, Prompt revision history, and explicit Prompt-to-article lineage, describe current evidence as a qualitative retrospective rather than causal lift. GSC is search demand, not AI Prompt volume.

## Output

Lead with the category-level conclusion. Do not dump every Prompt or foreground the internal article matching.

```markdown
## Prompt Audit

**Portfolio read:** ...

| Tracked category | Validity | Utility | Decisive evidence |
|---|---|---|---|
| ... | ... | ... | ... |

**Operating recommendation**
- [Category]: one natural-language recommendation grounded in the two judgments.

**Evidence limits**
- ...
```

Do not force recommendations into `working`, `SEO-only`, `research-sensor`, `revise`, `retire`, or any other fixed enum. Use the language the evidence warrants. Name categories, Prompts, and articles in user-facing copy; do not expose UUIDs.

Stop after the read-only recommendation. Prompt edits, Topic reassignment, tracking changes, or new visibility checks require separate explicit confirmation under the CUD Rule.
