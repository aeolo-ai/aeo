# Content Strategy Reference

## Commands

### `/aeo strategy show`
Show the current content strategy for the active domain.

```bash
aeo strategy show
```

Returns the strategy manifest (markdown). If no strategy exists, returns a template.

### `/aeo strategy update`
Create or update the content strategy. Uses PUT (atomic replace via upsert).

```bash
aeo strategy update \
  --manifest "## Brand Positioning\n..."
```

**Flags:**
| Flag | Type | Description |
|------|------|-------------|
| `--manifest` | string | Full strategy manifest (markdown, max 100K chars) |

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

## Initial Strategy Creation Guide

When creating a strategy for the first time:

1. **Load context first**: Run `/aeo` to get brand profile + audit + visibility
2. **Identify gaps**: Look at visibility gaps — which engines, which topics are underserved?
3. **Check brand profile**: Ensure `brand_context` is filled for durable
   positioning/audience/narratives, and check `writing_styles` /
   `brand_voice_examples` separately for tone.
4. **Draft manifest**: Use the template above. Focus on:
   - What makes this brand unique (positioning)
   - What content types work best for the gaps (balance)
   - Top 3–5 topics to write next (priority queue)
5. **Document cadence in the manifest if needed**: Match to the team's capacity. Start conservative (for example, weekly with 2 articles), but keep cadence as plain strategy text unless a separate scheduler is configured.
6. **Propose/apply**:
   - Interactive CLI/operator flow: after explicit approval, save with
     `aeo strategy update --manifest "..."`
   - Background writing job or chat flow: do not write product memory directly.
     Return a reviewed `content_strategy.manifest` patch for the user/operator
     to apply.

---

## When to Update the Manifest

- **After proposals are generated**: Add accepted topics to Priority Queue, remove completed ones
- **After publishing an article**: Update Changelog, adjust Content Balance if mix shifted
- **After a visibility check**: New gaps may surface — update Priority Queue
- **After brand profile changes**: Positioning section may need alignment
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
  "manifest": "## Brand Positioning\n..."
}
```
