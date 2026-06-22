# Content Strategy Reference

## Commands

### `/aeo strategy show`
Show the current content strategy for the active domain.

```bash
aeo strategy show
```

Returns the manifest (markdown) and schedule_config table. If no strategy exists, returns a template.

### `/aeo strategy update`
Create or update the content strategy. Uses PUT (atomic replace via upsert).

```bash
aeo strategy update \
  --manifest "## Brand Positioning\n..." \
  --frequency weekly \
  --articles-per-cycle 3 \
  --preferred-days mon,wed,fri
```

**Flags:**
| Flag | Type | Description |
|------|------|-------------|
| `--manifest` | string | Full strategy manifest (markdown, max 100K chars) |
| `--frequency` | enum | `daily`, `weekly`, `biweekly`, `monthly` |
| `--articles-per-cycle` | int | 1–20 articles per publishing cycle |
| `--preferred-days` | list | Comma-separated: `mon,tue,wed,thu,fri,sat,sun` |

---

## Manifest Template

A good manifest has these sections:

```markdown
## Brand Positioning
How the brand should appear in AI search results.
Key differentiators, tone, and authority signals.

## Content Balance
Target mix of article types (e.g., 40% how-to, 30% comparison, 20% thought leadership, 10% FAQ).
Language distribution if multi-language.

## Priority Queue
Highest-priority topics to address next, with rationale.
Link to visibility gaps or competitive intelligence.

## Constraints
Topics to avoid, compliance requirements, tone guidelines.
Competitor mentions policy.

## Changelog
- 2026-03-16 — Initial strategy created based on visibility audit
```

---

## Schedule Config Fields

| Field | Type | Default | Description |
|-------|------|---------|-------------|
| `frequency` | string | — | Publishing cadence |
| `articles_per_cycle` | number | — | How many articles per cycle |
| `preferred_days` | string[] | — | Days of the week to publish |

---

## Initial Strategy Creation Guide

When creating a strategy for the first time:

1. **Load context first**: Run `/aeo agent context` to get brand context + audit + visibility
2. **Identify gaps**: Look at visibility gaps — which engines, which topics are underserved?
3. **Check brand context**: Ensure `brand_context` is filled for durable
   positioning/audience/narratives. Use tone/reference analysis only when a
   task explicitly selects it.
4. **Draft manifest**: Use the template above. Focus on:
   - What makes this brand unique (positioning)
   - What content types work best for the gaps (balance)
   - Top 3–5 topics to write next (priority queue)
5. **Set schedule**: Match to the team's capacity. Start conservative (weekly, 2 articles)
6. **Save**: `aeo strategy update --manifest "..." --frequency weekly --articles-per-cycle 2`

---

## When to Update the Manifest

- **After proposals are generated**: Add accepted topics to Priority Queue, remove completed ones
- **After publishing an article**: Update Changelog, adjust Content Balance if mix shifted
- **After a visibility check**: New gaps may surface — update Priority Queue
- **After brand context changes**: Positioning section may need alignment
- **Monthly review**: Full review of all sections, trim stale items

---

## API Reference

| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/v1/connector/domains/:domainId/strategy` | Get strategy (markdown) |
| PUT | `/v1/connector/domains/:domainId/strategy` | Create/replace strategy |

PUT body:
```json
{
  "manifest": "## Brand Positioning\n...",
  "schedule_config": {
    "frequency": "weekly",
    "articles_per_cycle": 3,
    "preferred_days": ["mon", "wed", "fri"]
  }
}
```
