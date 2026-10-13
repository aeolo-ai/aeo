# Content Review

Review an existing article using [writing-policy.md](writing-policy.md), the approved reader request, channel rules and actual source material. Keep factual findings, request compliance and editorial suggestions separate. A style preference or related topic is not a factual error.

## Answer and scope — within the existing review

Check the saved title, metadata and every body paragraph, list item, table and caption. Compare the answer with the requested question and conditions. Do not review only quotations and numbers. Track unexamined units and missing inputs internally; missing evidence must not become a clean result.

## /aeo content review <id> — Review workflow

### Flow

1. Read the actual saved article with `aeo content get <id>` and brand context with `aeo agent context`. Load the approved channel instructions, corrections and allowed sources through the available reference tools. A context summary or source listing is not the full source. Prefer the exact source text delivered during writing when it is available; identify any later evidence separately.
2. Read [writing-policy.md](writing-policy.md). Apply the factual and request checks below. If required source material is inaccessible, report what is missing and which conclusions remain unverified. Do not silently substitute a smaller evidence packet.
3. Return concise findings using the report below. Suggest targeted edits; review itself does not authorize editing, deleting, overwriting or publishing. Apply only changes already authorized by the user.

### Review Checklist

#### Facts and evidence

Use the source-to-claim criteria in the shared policy. For each factual assertion, preserve subject, relationship, eligibility, market, period, quantity, attribution and certainty. Split a sentence's independent assertions when support differs; one supported clause cannot clear its unsupported neighbour.

- **supported**: the allowed evidence supports the actual assertion at its stated scope and strength.
- **unsupported**: the evidence examined does not establish the assertion. This does not prove it false.
- **contradicted**: applicable evidence directly conflicts with the assertion. Identify that conflict.
- **not checked**: necessary evidence is inaccessible, review coverage is incomplete or the check failed. Hold the affected conclusion; never report it as passed.

Do not treat the absence of a public URL as fabrication. Approved brand documents and readable internal sources can support a claim; preserve source identity and the matching passage internally, respecting disclosure restrictions and the channel's link policy. A URL or source ID alone is not semantic support. Do not label a statement fabricated just because evidence is missing. Quotes need matching original words and attribution; weakening an unsupported assertion does not supply evidence.

Do not require citations for pure questions or ordinary non-factual reader advice, such as asking the clinic about appointment arrangements. A conversational introduction such as “we often get this question” is not by itself a factual defect under the agreed review scope. Concrete claims about actual customers, credentials, services, frequency statistics or measured outcomes still need matching evidence. Separate those claims from surrounding advice.

Record the defective passage, its individual assertion, verdict, source/passage or precise evidence gap, and a minimal remedy. Do not reproduce the entire article and source corpus in the report or ask the agent to reconstruct hashes/IDs the tool already supplies.

#### Request compliance

Check the commissioned question, requested comparisons or conditions, approved narrator/register and explicit channel requirements. Identify an actual unmet instruction; do not replace it with generic GEO conventions. Style examples do not establish credentials or experience. Preserve material limitations without narrating the agent's research process.

#### Editorial suggestions

Use the shared house style and the selected channel's requirements. Suggest changes where repetition, confusing structure or empty process narration makes the answer less useful. Keep preferences separate from factual verdicts; an explicit channel requirement belongs under request compliance. No compulsory opening summary, FAQ count, expert quotation, anecdote, author biography or competitor mention is needed to make an article truthful. Technical metadata, syntax and schema validation belong to the existing tools; report their actual errors instead of inventing another prose checklist.

### Report Format

- **Scope**: article/version reviewed, sources examined, missing inputs and coverage limits.
- **Facts**: supported/no factual defects found within this scope, corrections needed, or not checked. List actual defects with the passage, verdict and source/gap; do not force an error count or invent one.
- **Request**: met requirements and any specific omissions.
- **Editing**: separate suggestions, with the applicable channel rule where required.
- **Next action**: targeted corrections, missing evidence or a planning decision. Do not collapse these into a single “deploy-ready” score.

### Execution boundary

`content_review_ack` is the writer's recorded self-review. Acceptance of its structure and source references does not establish independent factual approval. Claim an independent review only when an actual separate reviewer result is available, and identify its scope. Do not launch an additional model, paid task or legal review merely because this guide mentions review. The existing workflow owns retries, persistence and publication gates; tool failure is an execution issue, not evidence that the article is false. Legal compliance, when requested or required by the workflow, needs its own applicable jurisdiction and current sources; do not certify it from factual support alone.
