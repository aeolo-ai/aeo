# Content Idea

Use this agent-only workflow for `/aeo content idea <domain_id>`. It is not a bare `aeo` CLI verb: the external agent runs existing read commands and judges the evidence.

## Input

Require exactly one `domain_id` UUID. Do not resolve a brand name, hostname, or active-domain fallback. If the argument is absent or invalid, request the `domain_id` and stop.

## Read live evidence

Run independent reads in parallel when possible:

```bash
aeo agent context -d <domain_id>
aeo strategy show -d <domain_id>
aeo content list --limit 200 -d <domain_id>
aeo prompts list -d <domain_id>
aeo measure visibility -d <domain_id>
aeo measure traffic --days 28 -d <domain_id>
aeo measure overview -d <domain_id>
```

Also inspect current Topic-to-Prompt assignments through an authenticated read surface when available. If the CLI output does not expose Topic assignments and no other read surface is available, state that limitation; never invent an existing Topic.

Do not run a visibility check, generate content, create a Topic, reassign a Prompt, or write data.

## Separate observations

Keep these buckets distinct before judging:

1. **Brand** — offering, category, customer situation, priority CEP, differentiation, negative positioning, allowed claims, language, and market.
2. **GEO** — active business Topics, Topic-scoped active Prompts, latest visibility gaps, and current date.
3. **Content** — separate `liveCoverage` from `inFlightContent`.
   - Count only a verified owned canonical URL or deploy fact as live coverage.
   - Treat `draft`, `review`, and `approved` as production-collision signals only.
   - Do not infer deployment from a status label alone.
   - PBN publication does not close owned coverage.
4. **Demand/outcome** — use GSC and GA only when present.
   - GSC queries and product pages show observed search demand, not AI Prompt volume.
   - Use GA conversion evidence only for verified live article landing pages.
   - Missing GSC/GA must not stop the recommendation.

Treat prior recommendations as comparison material, never ground truth. Re-read current evidence and judge afresh.

## Choose one article idea

Return exactly one new article idea for the next production slot.

- Place it in a verified active Topic when one fits; otherwise propose a new parent Topic without saving it.
- Never return `refresh_existing`, `no_recommendation`, or a skipped slot.
- Avoid cannibalization through distinct search intent, explanatory frame, audience situation, or first-party evidence.
- Prefer a customer situation where the brand can naturally be recommended, compared, or cited.
- Avoid generic explainers when the brand has no credible reason to answer.
- Use timeliness only when the current date materially changes the question. Never add or remove a year mechanically.
- Stay inside verified facts. Turn unsupported efficacy, certification, testing, ingredient, or safety claims into cautions.
- Do not compute a universal score or apply hard thresholds. Explain the decisive tradeoff from the whole context.

## Output

```markdown
## Content Idea

**Topic placement:** existing | proposed
**Parent Topic:** ...
**Next article:** ...

**Why this article now:** One concise paragraph explaining the decisive tradeoff.

**Evidence**
- [brand] ...
- [visibility] ...
- [content] ...
- [GSC] ...
- [GA] ...

**Missing signals**
- ...

**Caution**
- ...
```

Omit evidence rows that do not exist. Keep the parent Topic distinct from the one-article idea. Do not expose the input UUID in user-facing copy. Stop after the recommendation unless the user separately confirms a mutation.
