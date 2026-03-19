# Visibility — check & show

---

## /aeo visibility check — Trigger a new visibility check

Run a fresh check against all tracked prompts. Use this after publishing content to measure GEO impact, or to get current gap data before generating new proposals.

### Step 1 — Trigger

```bash
aeo visibility check run
```

Response: `{ "success": true, "data": { "jobId": "uuid", "status": "pending"|"running" } }`

If a check is already running, the same jobId is returned — just poll it.

### Step 2 — Poll until complete

> See **[polling.md](polling.md)** for agent-specific polling instructions (Claude Code vs opencode).

Poll command:
```bash
aeo visibility check poll <jobId>
```

Typical duration: 3–8 minutes depending on number of prompts and engines.

### Step 3 — Present and analyze

Display the visibility report. Then:
1. Note new gaps vs previous check (if you have prior context)
2. Suggest running `/aeo content propose` to generate proposals for newly discovered gaps

---

## /aeo visibility show — Show last visibility snapshot

Returns the cached result of the most recent visibility check. Faster than running a new check — use this when you just need current gap context.

```bash
aeo visibility
```

Response: `text/markdown` — mention rate by engine, visibility gaps, top performing keywords.

If the response is empty or shows no check has been run, suggest `/aeo visibility check run`.
