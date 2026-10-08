# Writing Input Authoring

Read this when creating or cleaning brand inputs, voice, strategy or briefs. These are authoring responsibilities for existing inputs, not new nodes or a requirement to load this whole guide into every writing job. Keep Markdown where it is editable; keep IDs, locale, glossary and CTA fields structured.

## One owner per fact or instruction

| Input / existing owner | Write here | Keep elsewhere |
|---|---|---|
| Brand identity: profile name, category, description, value proposition, key features | What is offered; who it serves and in which situations; supported distinction and limits | Monthly priorities, duplicated catalog, worklogs |
| Scoped facts and corrections: catalog + source documents; `brand_context` for durable supplemental corrections | Product/service/location/market identity, corrected fact, supporting source and confirmation state | Tone instructions, unverified promotional promises |
| Brand voice; channel voice/reference; task-selected examples | Speaker, register, explanation style, do/don't examples. Brand default and channel/task overrides remain distinct | Ingredients, efficacy, prices, professional credentials or claims masquerading as style |
| `content_strategy.manifest` | Current reader decisions, ranked priorities with reasons, market/scope exclusions, intended next action | Brand introduction, voice rules, source dumps, historical changelog |
| Reference policy, source configuration, Drive/documents | Approved sources and how evidence is presented; documents that support specific questions | Archive of file downloads, upload status or meeting logistics |
| Selected research brief + user request | This question, audience, angle, candidate offering and fit rationale, required checks, format and CTA | Full brand/strategy copied again; guesses presented as confirmed catalog matches |

A reference to an owner is not evidence that the writer can access it. Check the saved value and actual runtime input/tool availability separately. If a correction has no dedicated field, retain it in scoped brand notes with a source link until a supported migration exists; never drop it merely to make notes shorter.

## Facts: enough scope to prevent reuse in the wrong place

For consequential claims record the subject, applicable market/SKU/location, source, and confirmation state. Add source date or check date when freshness matters. A light Markdown entry is sufficient:

```markdown
## Product A / KR / 50 ml — correction
Claim: [exact corrected fact]
Source: [document or customer message + date]
Status: customer-provided; original test report not yet reviewed
Scope: this SKU only; do not transfer to the US formulation
Replaces: [old claim, if known]
```

This is a template, not a product fact. Distinguish customer-provided, source-verified and unresolved information. Source-verified means the exact supporting passage was read, not just that a folder is connected. Do not convert an omitted claim into proof of absence. Resolve conflicting sources for the same scope; do not blend them. Preserve safety constraints and supersession history in linked records, with the current correction in the active input.

The same rule covers services: a qualification belongs to a named practitioner, a service to a location and market. It does not automatically apply to every branch.

## Voice: explain known facts without supplying new ones

Avoid: “Always mention centella, panthenol and ceramides.” This is a product instruction, not voice, and could invent ingredients for a selected SKU.

Prefer: “Explain the benefit first, then explain why using only ingredients and properties verified for the selected product.”

For a clinic, “Explain calmly and define unfamiliar terms” is voice; practitioner credentials and treatment availability belong to scoped facts. Examples demonstrate expression, not reusable efficacy claims. See [brand.md](brand.md) for fields and [content-create.md](content-create.md) for task references.

## Strategy: make a choice

A useful priority states whose decision matters, which supported offering could help, why this opportunity comes first and the next action. For example: “Prioritize first-time buyers comparing daily hydration options in the approved market; answer texture and layering questions using available product evidence; link to the eligible product page.” This is an illustrative decision, not a customer instruction.

A brand introduction or “Define the primary audience” placeholder does not make that decision. Replace template prompts with actual choices before saving. Keep the normalized manifest body within 5,000 characters; move detailed evidence and historical decisions to accessible linked documents. Existing long manifests remain readable. See [strategy.md](strategy.md).

## Assignment and evidence: different roles

The selected brief carries the planning decision directly. Its linked research is evidence to retrieve under the job's allowed tools. Candidate product/service IDs, relevance reasons, market fit and unresolved checks belong to the offering handoff; nomination is not verified recommendation. The writer checks current catalog/source material and records use, exclusion or unresolved fit through the supported workflow. An informational article may legitimately need no offering.

A missing optional research attachment is different from a missing required brief. Never invent either. Respect the job's required-input failure path and tool on/off settings; this guide grants no additional access or generation authorization.

## Cleanup and verification

1. Read current saved inputs and capture a reversible snapshot with source/time.
2. Assign each passage an owner; retain facts, customer corrections and meaningful restrictions. Move work history to linked records and remove duplicated introductions or unfilled templates.
3. Compare the proposed diff for changed scope, lost negation, unsupported certainty and cross-node contradiction. Read back any approved save and inspect the assembled prompt separately.
4. Compare first drafts with the same briefs, model and evidence access. Evaluate factual/scope errors, question resolution, offering fit, reader clarity and next action. Report review quality/latency separately. Guide cleanup or token reduction alone is not evidence of better output.

## Evidence behind this approach

[Anthropic's context engineering guidance](https://www.anthropic.com/engineering/effective-context-engineering-for-ai-agents) supports clear boundaries, sufficient relevant context and avoiding redundant instructions. It does not prove this exact node taxonomy is optimal. The taxonomy above is an Aeolo implementation choice to test with actual drafts. GEO citation/traffic outcomes remain a separate longitudinal measure; see [geo-strategy.md](geo-strategy.md).
