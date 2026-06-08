# Visibility — show scheduled results

---

## /aeo visibility show — Show last visibility snapshot

Visibility checks are scheduled/admin-only. Use the latest snapshot to understand current gap context.

```bash
aeo visibility
```

Response: `text/markdown` — mention rate by engine, visibility gaps, top performing keywords.

After displaying the visibility report:
1. Note new gaps vs previous check (if you have prior context)
2. Suggest a content or audit action that spends credits only after user confirmation
3. If the response is empty, explain that scheduled visibility data is not available yet and proceed from brand/strategy context
