# Topics — Stable customer situations and their measurement Prompts

## Identity model

- **Topic** is the stable business theme or customer situation the brand must answer.
- **Prompt** is one concrete question used to measure that Topic in AI engines.
- **Prompt stage** (`foundational`, `comparison`, `use-case`, `implementation`) is an orthogonal funnel lens, not a Topic.
- **Unassigned** is a system migration bucket, not a business Topic.

## Read before writing

```bash
aeo topics list
aeo topics list --include-archived
```

The list includes the current `revision`. Every update, archive, and restore must use that revision. If the server returns `REVISION_CONFLICT`, list again, show the changed state, and ask the user to confirm again. Never auto-retry a stale write.

## Measure existing Topic demand

```bash
aeo topics demand
aeo topics demand run --max-credits 1
aeo topics demand poll <jobId>
```

Read first: `topics demand` returns saved demand per active Topic, freshness, and the next batch's credit estimate. Reading and polling cost **0 credits** and never start a measurement. Add `--domain <domainId>` to scope any command explicitly.

Demand is a **category/head-term proxy for a Topic**, not the exact search volume of its conversational Prompts. Keep Google monthly search volume separate from the AI demand estimate; the latter is an estimate, not an observed count of questions asked in AI engines. A missing/provider-unavailable value is `null`, never zero. Report the measurement source, market, freshness, and any partial/unavailable result alongside numbers. Do not sum overlapping Topics or reuse one Topic's volume as the volume of each Prompt.

Each explicit run refreshes **up to 30 stale Topics for 1 credit**. The cache lasts **30 days**. Fresh cache and reuse of an already active job cost **0 credits**. `--max-credits` is mandatory, accepts integers **0–100**, and caps this invocation's customer credit spend. `--max-credits 0` permits only zero-cost reuse; a needed paid batch is refused. A larger budget does not request more than one batch. The server enforces the budget again, including for MCP callers.

Show the next batch estimate and use the user's authorized budget before running. Poll the returned job to a terminal outcome; job acceptance is not completed measurement. Distinguish complete, partial, and unavailable evidence, and read the refreshed snapshot before ranking Topics. If more than 30 Topics are stale, read the next quote after each batch; never loop paid runs beyond the user's total authorization.

Existing automatic lifecycle sweeps remain **platform-funded**; creating/editing Topics must not be described as silently charging a customer for these sweeps. Provider costs are tracked internally and are separate from the customer credit tariff. Candidate generation has its own workflow: preserve its signed `--demand-token` when creating the selected Topic to reuse already measured evidence.

## Create and update

```bash
aeo topics create --name "Sensitive-skin sunscreen" --description "Low-irritation daily SPF decisions"
aeo topics update <topicId> --revision 2 --name "Sensitive-skin daily SPF"
```

Topic names are unique among active Topics after case/whitespace normalization. `Unassigned` is reserved. Show the proposed name and description and get explicit confirmation before writing.

## Reassign Prompts

```bash
aeo topics assign-prompts <topicId> --prompt-ids <promptId1>,<promptId2>
```

Assignment is atomic for 1–100 active same-domain Prompts. Before confirmation, show Topic name and Prompt text—not bare UUIDs. Reassignment changes the current Prompt organization; it does not rewrite historical visibility results.

## Archive and restore

```bash
aeo topics archive <topicId> --revision 3
aeo topics restore <topicId> --revision 4
```

- Move every active Prompt before archiving; otherwise the server returns `TOPIC_HAS_PROMPTS`.
- The system Unassigned Topic cannot be renamed or archived.
- Archive is the Topic delete lifecycle. There is no destructive hard-delete command.
