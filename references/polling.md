# Polling Guide — Long-running Aeolo Jobs

Some Aeolo commands kick off async jobs that take minutes to complete.
Trigger returns a `jobId`; you must poll a status command to get results.

---

## Jobs that require polling

| Command | CLI command | Typical duration |
|---------|------------|-----------------|
| Site audit | `aeo audit run` | 3–8 min |
| AI writing | `aeo content write` | 2–8 min |
| Reference analysis | `aeo reference analyze` | 2–6 min |

---

## Polling flow

**Step 1 — Trigger, get jobId**

```
aeo audit run --max-pages 5
# → { "data": { "jobId": "abc-123", "status": "pending" } }
```

**Step 2 — Poll in the background**

Poll `aeo audit poll {jobId}`, `aeo reference poll {jobId}`, or `aeo content jobs` every 60 seconds using your runtime's timer or scheduling mechanism. Stop polling on completion or error.

**Step 3 — Confirm to user**

```
Job triggered (job: {jobId}). Polling every minute in the background.
You can keep working — I'll report back when it's done.
```

---

## Status response reference

| Response | Meaning | Action |
|----------|---------|--------|
| `{ "status": "pending"\|"running" }` | In progress | Wait |
| result/status JSON | Complete — full report or result payload | Stop polling, present report |
| `{ "code": "...FAILED" }` | Job failed | Stop polling, report error |
| `{ "code": "NOT_FOUND" }` | jobId invalid or expired | Stop polling, re-trigger |
