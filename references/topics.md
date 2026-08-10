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
