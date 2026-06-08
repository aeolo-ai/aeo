# Visibility — check & show

---

## /aeo visibility check — Trigger a credit-metered visibility check

Run a fresh check against tracked prompts. This reserves production credits before work starts: 1 credit per prompt x engine.

### Step 1 — Trigger

```bash
aeo visibility check run
```

Optional:

```bash
aeo visibility check run --engines=chatgpt,gemini --limit=10
aeo visibility check run --prompt-ids=id1,id2 --engines=chatgpt
```

Response includes `jobId`, `promptCount`, selected engines, and reserved credits.

### Step 2 — Poll until complete

> See **[polling.md](polling.md)** for agent-specific polling instructions.

```bash
aeo visibility check poll <jobId>
```

Typical duration: 3-8 minutes depending on number of prompts and engines.

### Step 3 — Present and analyze

Display the visibility report. Then:
1. Note new gaps vs previous check (if you have prior context)
2. Suggest a content or audit action that spends credits only after user confirmation

---

## /aeo visibility show — Show last visibility snapshot

Use the latest snapshot to understand current gap context.

```bash
aeo visibility
```

Response: `text/markdown` — mention rate by engine, visibility gaps, top performing keywords.

If the response is empty or shows no check has been run, ask whether to run `/aeo visibility check run`.
