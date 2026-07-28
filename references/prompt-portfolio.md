# Prompt Portfolio (safe restructure)

Use this agent-only workflow for `/aeo prompts portfolio <domain_id>`. It is not a
bare `aeo` CLI verb. This is the **only** safe way to restructure a live tracked
Prompt set — recover a shrunk portfolio, swap categories, or hit an exact target
count — because it previews an exact diff and gets explicit confirmation *before*
any write.

> **Why this exists.** `aeo prompts generate` proposes *and* mutates in one shot: it
> saves straight to `tracked`, its `--count` is a generation hint (not an exact
> contract), and it computes no keep/retire/replace diff. Run against a live
> portfolio it has shrunk a real one 30 → 14 (AEO-208/AEO-212). Never use
> `prompts generate` to clean up or restructure a live portfolio, and never
> bulk-untrack existing Prompts without the full plan below. `prompts generate` is
> safe only to seed a brand-new/empty portfolio.

## Input

Require exactly one `domain_id` UUID and a **target tracked count** (and, optionally,
a per-Topic allocation). Do not resolve a brand name, hostname, or active-domain
fallback. The target count is an **exact contract** — if you cannot compose the
portfolio to hit it exactly, do not write; report why and stop.

## 1 — Read current state

Run independent reads in parallel when possible:

```bash
aeo prompts list --status tracked -d <domain_id>
aeo prompts list --status untracked -d <domain_id>
aeo topics list -d <domain_id>
aeo strategy show -d <domain_id>
aeo agent context -d <domain_id>
```

Capture, before proposing anything:

- every tracked Prompt (id, text, stage, language, segment, Topic assignment)
- untracked Prompts that could be **reactivated** instead of re-created
- active Topics and their current Prompt counts
- strategy/brand **claim guardrails** (unsupported claims the Prompts must not assert)
- which tracked Prompts are **benchmark sensors** (the durable discovery probes you
  compare over time)

Record the current tracked total and per-Topic counts — these are the baseline the
final verify checks against.

## 2 — Propose an exact diff

Classify **every** existing Prompt into exactly one bucket, and list new candidates:

| Bucket | Meaning |
|--------|---------|
| `preserve` | tracked, stays as-is (all benchmark sensors default here) |
| `reactivate` | an untracked Prompt brought back to `tracked` |
| `update` | text/stage/segment/Topic edited, stays tracked |
| `add` | a brand-new Prompt to create |
| `untrack` | a tracked Prompt moved to `untracked` (soft removal) |

Rules:

- **Preserve benchmark sensors** unless the user explicitly names one to retire.
  Never silently drop a sensor to make room.
- Decide each candidate's **Topic before writing** — every added/reactivated/updated
  Prompt must land on an active Topic. Honor the per-Topic allocation if given.
- Compose the buckets so the resulting tracked total **equals the target exactly**.
  If it cannot (e.g. too few valid candidates), stop and report — do not write a
  close-enough portfolio.

## 3 — Audit the proposal (before preview)

For the proposed final set, check:

- **Validity** — each Prompt is a real, non-leading, brand-answerable question with a
  clear Topic fidelity. Reuse the judgment model in [prompt-audit.md](prompt-audit.md).
- **Duplication** — no two tracked Prompts are near-duplicates (same intent /
  paraphrase). Drop or merge before preview.
- **Claim safety** — no Prompt asserts or presupposes a claim the brand can't support
  (from the strategy/brand guardrails). A guardrail-violating candidate must not reach
  the write step.

## 4 — Preview and get explicit confirmation

Show the **full** plan before any write — do not write first and summarize after:

```markdown
## Prompt Portfolio — proposed change (domain: {name})

Baseline: {N} tracked across {T} Topics  →  Target: {target}

**Add ({k})** — [Topic] question …
**Reactivate ({k})** — …
**Update ({k})** — before → after …
**Untrack ({k})** — question (reason) …
**Preserve ({k})** — {count summary; list benchmark sensors explicitly}

**Final per-Topic allocation**
| Topic | count |
| … | … |
**Expected tracked total: {target}**  (= sum of Topic counts above)

Portfolio revision at read: {revision/snapshot marker}
```

- State the **expected final tracked total and per-Topic totals** so the user
  confirms the exact shape, not a vibe.
- Note the portfolio **revision** you read at (e.g. a hash of the current tracked id
  set, or the read timestamp). This identifies the portfolio the plan was built on.
- Ask **"Proceed?"** and wait for explicit confirmation. No production write happens
  before it (CUD Rule).

## 5 — Write, then verify

Only after confirmation, apply the diff. Order the writes so the portfolio is never
left below target mid-flight, and treat the batch as all-or-nothing:

```bash
# reactivate / update existing rows first, then add, then untrack last
aeo prompts update <id> --status tracked --stage <s> --segment <tags> -d <domain_id>
aeo prompts add --prompts-json '[{"prompt":"…","stage":"…"}]' -d <domain_id>
aeo topics assign-prompts <topicId> --prompt-ids <id,id,…> -d <domain_id>
aeo prompts update <id> --status untracked -d <domain_id>   # removals LAST
```

- **Re-check the revision first.** Re-read `prompts list --status tracked` and confirm
  the tracked id set still matches what you previewed. If it changed (someone else
  edited the portfolio since preview), **abort** — the preview is stale; re-run from
  step 1. Do not write a stale plan.
- **Apply removals last** so a mid-sequence failure never leaves the portfolio
  shrunk below target with nothing added back.
- If any step fails, **stop and roll back** the writes already made (reactivate what
  you untracked, delete what you added) rather than leaving a partial result. Report
  the partial state explicitly.

Then verify with fresh reads — do not report success on the plan alone:

```bash
aeo prompts list --status tracked -d <domain_id>
aeo topics list -d <domain_id>
```

The portfolio is correct only when **all** hold:

```
tracked total == sum(active Topic prompt counts) == target count
```

and there are **no** unassigned tracked Prompts, no duplicates, and no
guardrail-violating Prompts. If any invariant fails, say so and do not claim success.

## Output

Lead with the outcome against the contract, then the deltas:

```markdown
## Prompt Portfolio — done

Tracked: {before} → {after} (target {target}) ✅
Per-Topic: {Topic}: {n} · …
Changed: +{added} added, {reactivated} reactivated, {updated} updated, {untracked} untracked
Benchmark sensors preserved: {list}
```

Name Topics, Prompts, and articles in user-facing copy; never expose UUIDs. A new
visibility check, content generation, or deployment is out of scope here — each needs
its own explicit confirmation.

> **Known limit (honest).** Until the CLI/API ship a server-side `--dry-run` and an
> atomic revisioned write, the "atomicity" and "stale-preview conflict" guarantees
> above are enforced by *this procedure* (ordered writes, re-read-before-write,
> manual rollback), not by the database. Follow the steps exactly; do not shortcut
> the re-read or the removals-last ordering.
