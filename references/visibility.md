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

**Flags:**
- `--engines <list>` — comma-separated. Supported set: `chatgpt,gemini,perplexity,google-ai-mode,google-aio,amazon`. Default (when the flag is omitted) is `chatgpt,gemini,perplexity`. `amazon` (Amazon Rufus, GA as of 2026-07-01) is opt-in — pass it explicitly; it costs more credits per query than the ChatGPT/Gemini/Perplexity engines. `grok` is currently **disabled** — never pass it.
- `--limit <n>` — cap the number of prompts checked
- `--prompt-ids <id,id>` — check only specific prompts

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
aeo visibility show
```

Response: `text/markdown` — **overall composite score + per-engine scores**, mention rate by engine (with avg position), stage breakdown, **Share of Voice** (which brands AI recommends, incl. competitors — populated even at 0% mentions), **citation-type distribution**, visibility gaps, cited sources, top queries, competitors. Enough to narrate the full score breakdown ("X% of AI recommendations go to CeraVe, Pure'AM 0%").

If the response is empty or shows no check has been run, ask whether to run `/aeo visibility check run`.
