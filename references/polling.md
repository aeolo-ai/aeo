# Polling Guide — Long-running Aeolo Jobs

Some Aeolo commands kick off async jobs that take minutes to complete.
Trigger returns a `jobId`; you must poll a status command to get results.

---

## Jobs that require polling

| Command | CLI command | Typical duration |
|---------|------------|-----------------|
| Visibility check | `aeo visibility check run` | 3-8 min |
| Site audit | `aeo audit run` | 3–8 min |
| Content generation | `aeo content generate` | 2–8 min |
| Reference analysis | `aeo reference analyze` | 2–6 min |
| ~~Video generation~~ | **RETIRED 2026-07-27** — `aeo video generate` / `aeo video poll` are sunset (shortform output is no longer a supported channel). Do not start new jobs. | — |

---

## Polling flow

**Step 1 — Trigger, get jobId**

```
aeo visibility check run
# → { "data": { "jobId": "abc-123", "status": "pending" } }
```

**Step 2 — Poll in the background**

Poll `aeo visibility check poll {jobId}`, `aeo audit poll {jobId}`, `aeo reference poll {jobId}`, or `aeo content jobs` every 60 seconds using your runtime's timer or scheduling mechanism. Stop polling on completion or error.

**Step 3 — Confirm to user**

```
Job triggered (job: {jobId}). Polling every minute in the background.
You can keep working — I'll report back when it's done.
```

---

## Site audit: current scoring and completion

`aeo audit run` uses the same default 10-page budget as the dashboard. Use
`--max-pages N` for an integer from 1 to 50 and `--channel-id <uuid>` to target
a channel. The run spends credits; a report read and polling do not launch a new run.
Poll the returned job ID with `aeo audit poll <jobId>` until terminal status, then
read `aeo domain audit`. That report selects the domain's latest saved run: compare
its run ID, channel, measurement URLs and timestamps with the job you requested.

The current total is the rounded equal-weight average of Lighthouse SEO,
Performance and Accessibility category means, matching the dashboard. Best
practices is separate. Missing measurements are N/A, never zero; legacy readiness
is a different historical evaluation and must not substitute for the current total.
Inspect technical SEO issue codes, affected URLs, evidence, device scope and sample
coverage before recommending a fix. Sampled lab scores do not establish whole-site
performance, indexing, AI citations or conversion. On read failure, retry the read;
do not start another paid audit automatically.

## General job statuses

| Response | Meaning | Action |
|----------|---------|--------|
| `{ "status": "pending"\|"running" }` | In progress | Wait |
| result/status JSON | Complete — full report or result payload | Stop polling, present report |
| `{ "code": "...FAILED" }` | Job failed | Stop polling, report error |
| `{ "code": "NOT_FOUND" }` | jobId invalid or expired | Stop polling, verify job ID and saved results before considering a new paid run |
