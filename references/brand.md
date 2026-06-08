# Brand & Prompts

---

## /aeo domain brand — Show current brand profile

```bash
aeo domain brand
```

Response: `text/markdown` — brand profile including name, category, industry, value proposition, key features, language, and `brand_context` (free-form GEO strategy markdown, up to 50,000 chars).

If empty or JSON error, suggest running setup and checking the domain ID.

---

## /aeo domain brand update — Update brand profile fields

1. Fetch current profile first (show it to the user)
2. Ask the user what they want to change
3. Confirm before writing

Accepted fields:

| Field | Type | Notes |
|-------|------|-------|
| `name` | string | Brand display name |
| `category` | string | e.g. `"B2B SaaS"` |
| `industry` | string | e.g. `"Developer Tools"` |
| `value_proposition` | string (max 2000) | Core positioning statement |
| `key_features` | string[] (max 20) | Feature list for brand mentions |
| `primary_language` | ISO 639-1 | e.g. `"en"`, `"ko"`, `"ja"` |
| `brand_context` | string (max 50000) | Free-form GEO strategy (see template below) |

```bash
aeo domain brand update --name="..." --category="..." --value-proposition="..."
```

Partial update — unset fields are preserved.

### brand_context template

`brand_context` is free-form markdown. Suggest this structure when helping a user build it from scratch:

```markdown
## Brand Overview
[Brand mission, positioning, and what makes it different]

## Target Audience
[Who the brand serves — personas, pain points, jobs-to-be-done]

## GEO Strategy
[Which AI engines to prioritize, content angles to emphasize, competitive positioning]

## Key Narratives
[3–5 core messages the brand wants AI engines to associate with]

## Competitive Context
[Main competitors, how to frame comparisons, where the brand wins]

## Tone & Voice
[Writing style, formality level, phrases to use/avoid]
```

---

## /aeo prompts list — List prompts grouped by stage

```bash
aeo prompts
```

Response: `text/markdown` — table grouped by stage (foundational → comparison → use-case → implementation), showing language, query form, prompt text, visibility score, and last checked date.

After displaying:
- Note which stages are sparse or have low visibility scores
- Suggest adding prompts where coverage is thin
- If no prompts exist, prompt the user to add foundational ones first

---

## /aeo prompts add — Add a manual prompt

Ask the user for the prompt details, then run:

```bash
aeo prompts add --prompt="best project management tools" --stage=comparison --language=en --query_form=conversational
```

Accepted fields:

| Field | Type | Required | Default | Example |
|-------|------|----------|---------|---------|
| `canonical` | string | ✅ | — | `"best project management tools"` (CLI: `--prompt`) |
| `localized_prompt` | string | — | same as canonical | `"最好的项目管理工具"` |
| `stage` | `foundational` \| `comparison` \| `use-case` \| `implementation` | — | `foundational` | `"comparison"` |
| `language` | `en` \| `ko` \| `ja` \| `zh` \| `ar` | — | `en` | `"zh"` |
| `query_form` | `short-tail` \| `long-tail` \| `conversational` | — | `conversational` | `"long-tail"` |

Confirm the details with the user before submitting. After success, ask whether to run `/aeo visibility check run` to measure the new prompt set; it reserves credits based on prompt x engine count.

---

## /aeo prompts update — Edit an existing prompt

```bash
aeo prompts update <promptId> --prompt="updated text" --stage=use-case
```

Accepted fields (all optional, at least one required):

| Field | Type | Notes |
|-------|------|-------|
| `canonical` | string | English prompt text |
| `localized_prompt` | string | Native-language prompt |
| `stage` | `foundational` \| `comparison` \| `use-case` \| `implementation` | Move to different stage |
| `query_form` | `short-tail` \| `long-tail` \| `conversational` | Update query form |

Use `/aeo prompts list` first to get the prompt ID. Confirm with user before updating.

---

## /aeo prompts delete — Delete a prompt (soft delete)

1. Run `/aeo prompts list` to show current prompts with IDs
2. Ask the user to confirm which prompt(s) to delete
3. Show the list and ask "Proceed?" before calling DELETE

```bash
aeo prompts delete <promptId>
```

Response: `{ success: true, deleted: { id, canonical } }`

- Soft delete — sets `deleted_at`, data is preserved
- `404` if prompt not found or already deleted
- For bulk deletion (multiple prompts), confirm the full list once then loop calls sequentially
