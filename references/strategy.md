# Content Strategy Reference

## Commands

### `/aeo strategy show`
Show the current content strategy for the active domain.

```bash
aeo strategy show
```

Returns the manifest (markdown). If no strategy exists, returns a template.

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

> Scheduling flags (`--frequency`, `--articles-per-cycle`, `--preferred-days`, `--auto-propose`) were removed. The CLI rejects them — encode any cadence/priority intent inside the manifest instead.

---

## Visual Style Guide

The style guide steers **image generation** — it is applied whenever "apply brand style" is on. Three things ride along:

- **description + keywords** → appended to the prompt as text.
- **board images** → attached to the generation as reference images.
- **definitions** → one `IMAGE N — …` line per attached image, in the order the provider receives them. Without a definition an image arrives as an anonymous attachment, so the model has to guess what it is looking at.

There are two kinds of board. The **brand board** applies to every generation. A **product board** applies only when that product is the subject, and it outranks the brand board for that product. Each board holds up to **24** images (a candidate pool, not the per-generation payload — the model's reference budget picks from it).

### `/aeo strategy visual`

```bash
aeo strategy visual
```

Read this **before** any board edit. Boards are keyed by image URL, so the read is the only way to learn the arguments for `--remove-images` and `--definitions`. It returns the description, keywords, the brand board (each image's URL and its definition), one row per product board, and how many definitions are stored.

### `/aeo strategy visual update`

```bash
# text half
aeo strategy visual update \
  --description "Clean clinical photography, warm neutral background" \
  --keywords "minimal,clinical,warm"

# brand board
aeo strategy visual update --add-images "https://…/1.jpg,https://…/2.jpg"
aeo strategy visual update --remove-images "https://…/1.jpg"

# one product's board
aeo strategy visual update --product <productId> --add-images "https://…/3.jpg"

# per-image definitions (merge; "" deletes one)
aeo strategy visual update \
  --definitions '{"https://…/1.jpg":"product front, white seamless","https://…/2.jpg":""}'
```

**Flags:**
| Flag | Type | Description |
|------|------|-------------|
| `--description` | string | Style directive, ≤2000 chars. Replaces. |
| `--keywords` | csv | Up to 20 keywords of ≤60 chars. Replaces; `--keywords ""` clears. |
| `--add-images` | csv of URLs | Appended to the end of the target board, duplicates dropped |
| `--remove-images` | csv of URLs | Removed from the target board |
| `--product` | product ID (UUID) | Targets that product's board instead of the brand board. Only valid alongside an image flag; get IDs from `aeo brand products` |
| `--definitions` | JSON object | `{"image url": "definition"}`, ≤50 images per update, ≤500 chars each |

**Semantics that differ between flags:**

- **Boards replace, definitions merge.** `--definitions` only touches the keys it names: an absent key keeps its definition, and an empty string (`""`) deletes it. There is no way to clear the whole map in one call, deliberately — the dashboard's annotator sends one image at a time and a replace would wipe the rest.
- **Board edits preserve definitions and vice versa.** Removing an image leaves its definition behind (harmless, and it comes back if the image returns); `aeo strategy visual` reports how many definitions no longer point at any board.
- **The 24-image cap is refused, not truncated.** If an `--add-images` would overflow a board the whole update is rejected and nothing is written. A `--remove-images` in the same call frees its slots first.
- One flag per command: repeated `--add-images` flags do **not** accumulate (the last wins). Pass one comma-separated list.

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

1. **Load context first**: Run `/aeo agent context` to get brand context + audit + visibility
2. **Identify gaps**: Look at visibility gaps — which engines, which topics are underserved?
3. **Check brand context**: Ensure `brand_context` is filled for durable
   positioning/audience/narratives. Use tone/reference analysis only when a
   task explicitly selects it.
4. **Draft manifest**: Use the template above. Focus on:
   - What makes this brand unique (positioning)
   - What content types work best for the gaps (balance)
   - Top 3–5 topics to write next (priority queue)
5. **Save**: `aeo strategy update --manifest "..."`

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
| GET | `/v2/connector/domains/:domainId/strategy` | Get strategy (markdown) |
| PUT | `/v2/connector/domains/:domainId/strategy` | Create/replace strategy |

PUT body:
```json
{
  "manifest": "## Brand Positioning\n..."
}
```
