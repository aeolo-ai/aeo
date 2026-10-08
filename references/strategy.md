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
  --manifest "## Reader Decision\n..."
```

**Flags:**
| Flag | Type | Description |
|------|------|-------------|
| `--manifest` | string | Full strategy manifest (markdown, normalized body max 5,000 characters; detailed evidence/history belongs in linked documents) |

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

Use [writing-inputs.md](writing-inputs.md) for the boundary between facts, voice and strategy. Replace prompts below with decisions before saving; do not copy the brand introduction or accumulate a worklog.

```markdown
## Reader Decision
[Whose decision should the content help, in which market and situation?]

## Priorities
[Rank the next opportunities; explain demand, offering fit and available evidence.]

## Scope
[Included markets/topics and strategic exclusions; link to the facts and policies that constrain them.]

## Next Action
[What should the reader be able to do next, and which eligible destination supports it?]
```

The normalized manifest body is limited to 5,000 characters on new saves. Existing longer manifests remain readable; do not silently truncate or overwrite them. The separately bounded planning-settings envelope is not extra space for prose. Preserve existing structured planning settings when changing only the body.

## Initial Strategy Creation Guide

1. Read current brand identity, offerings, notes and source/market constraints.
2. Identify a real customer decision the offering can support. Use demand, observed search/visibility and existing content as evidence; absence of a brand mention alone is not enough.
3. Rank opportunities with reasons and relevant source links. Distinguish new discovery from branded purchase checks and from unsupported product claims.
4. Choose scope and intended reader action. Do not impose a universal article mix or repeat voice rules.
5. Review the proposed replacement, then save with `aeo strategy update --manifest "..."` under the normal write authorization.

## When to Update the Manifest

Update when approved priorities, market eligibility, evidence or the next reader action change. Completing an article can change the priority queue; publishing does not require appending a changelog. Keep detailed rationale/history in linked records, and retain only current decisions in the active input. A single visibility fluctuation is a review signal, not an automatic strategy rewrite.

---

## API Reference

| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/v2/connector/domains/:domainId/strategy` | Get strategy (markdown) |
| PUT | `/v2/connector/domains/:domainId/strategy` | Create/replace strategy |

PUT body:
```json
{
  "manifest": "## Reader Decision\n..."
}
```
