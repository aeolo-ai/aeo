# Topic Suggest

Use this agent-only workflow for `/aeo topics suggest <domain_id>`. It is not a bare `aeo` CLI verb. The external agent reads live Aeolo context and proposes a Topic architecture without writing data.

## Input

Require exactly one `domain_id` UUID. Do not resolve a brand name, hostname, or active-domain fallback. If the argument is absent or invalid, request the `domain_id` and stop.

## Read live evidence

Run independent reads in parallel when possible:

```bash
aeo agent context -d <domain_id>
aeo strategy show -d <domain_id>
aeo products -d <domain_id>
aeo topics list -d <domain_id>
aeo prompts list --status tracked -d <domain_id>
aeo content list --limit 200 -d <domain_id>
aeo measure visibility -d <domain_id>
aeo measure traffic --days 30 -d <domain_id>
```

If Topics are unavailable, say so and propose a fresh topology. Do not substitute legacy segment tags for Topics. Do not run a visibility check, generate Prompts, create or change Topics, or write data.

Use the `language` from live brand context as the sole language for the entire user-facing response, including headings, labels, Topic names and definitions, Prompt territories, evidence, guardrails, and allocations. Do not infer the output language from the conversation or hard-code a locale pair. If live brand context has no language, state that limitation and ask the user which language to use before proposing Topics.

## Judge the whole context

Treat current Topics, segment tags, Prompts, and prior proposals as evidence, never ground truth. Derive the topology afresh from:

- the offer, priority products, customer situations, differentiation, negative positioning, language, and market;
- reviewed strategy, launch direction, claims policy, and operating cadence;
- verified product and first-party evidence the brand can credibly answer from;
- tracked Prompt intent, visibility gaps, GSC demand, and owned live coverage when present.

Do not compute a universal score or apply hard thresholds.

## Propose 3–5 durable Topics

A Topic is a stable business theme or customer situation that can support multiple Prompts and articles. It is not one Prompt, article title, format, channel, language, campaign, or temporary content slot.

For every proposed Topic:

- connect it to a real customer situation and a reason this brand can answer;
- give it at least two distinct Prompt territories in the live brand language;
- keep it useful beyond one product launch or calendar year;
- keep unsupported efficacy, certification, clinical, safety, or regulatory claims out of the definition;
- keep language or regional variants inside the same Topic unless the underlying business situation truly differs by market;
- merge overlapping Topics instead of preserving a noisy current taxonomy.

Prefer the smallest coherent topology. A missing evidence source becomes a guardrail, not a fabricated Topic.

## Output

Use this semantic structure, translating every heading and label into the live brand language. Bracketed text describes a slot; do not copy it literally.

```markdown
## [Topic Suggestions]

**[Topology judgment]:** [one concise paragraph]

### 1. [Topic name]
**[Action]:** [keep | rename | merge | new]
**[Definition]:** ...
**[Why this brand]:** ...
**[Prompt territories]:**
- ...
- ...
**[Evidence]:** ...
**[Missing / guardrails]:** ...

## [Proposed Prompt Allocation]
- [Topic name]: [n Prompts]

## [Not Topics]
- ...
```

Return 3–5 Topics. Do not draft the full final Prompt set in this workflow. Do not expose the input UUID. Stop after the proposal unless the user separately confirms a mutation.
