# Polling Guide — Long-running Aeolo Jobs

Some Aeolo commands kick off async jobs that take minutes to complete.
Trigger returns a `jobId`; you must poll a status command to get results.

---

## Jobs that require polling

| Command | CLI command | Typical duration |
|---------|------------|-----------------|
| `aeo visibility check run` | `aeo visibility check run` | 3–8 min |

> Future jobs (server-side content generation, bulk audit) will follow the same pattern.

---

## Polling flow

**Step 1 — Trigger, get jobId**

```
aeo visibility check run
# → { "data": { "jobId": "abc-123", "status": "pending" } }
```

**Step 2 — Poll in the background**

Schedule a recurring poll every 1 minute using whatever scheduling mechanism your runtime provides:
- Claude Code: `CronCreate` with `"*/1 * * * *"`
- NullClaw/ZeroClaw: `cron add` or `cron add-agent`
- Other runtimes: any scheduler that can run `aeo visibility check poll {jobId}` periodically

The poll should self-cancel on completion or error.

**Step 3 — Confirm to user**

```
Visibility check triggered (job: {jobId}). Polling every minute in the background.
You can keep working — I'll report back when it's done.
```

---

## Status response reference

| Response | Meaning | Action |
|----------|---------|--------|
| `{ "status": "pending"\|"running" }` | In progress | Wait |
| `text/markdown` (starts with `#`) | Complete — full report | Stop polling, present report |
| `{ "code": "CHECK_FAILED" }` | Job failed | Stop polling, report error |
| `{ "code": "NOT_FOUND" }` | jobId invalid or expired | Stop polling, re-trigger |
