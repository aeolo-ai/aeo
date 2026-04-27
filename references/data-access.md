# Data Access — Research Source Configuration

## Default sources (always available)

These are available to all aeo users without configuration:

| Source | Command | What to look for |
|--------|---------|-----------------|
| Google Drive | `aeo drive list`, `aeo drive read <id>` | Product specs, test results, customer feedback, internal docs |
| Brand Context | `aeo domain brand` | Value proposition, key features, competitors |
| Published Content | `aeo content list --status=published` | Existing articles (avoid overlap, find cross-link targets) |
| Visibility Data | `aeo visibility show` | Gap queries, competitor mentions |

## Custom sources (per-domain)

Run `aeo config data-sources show` to see this domain's configured sources.

If no custom sources are configured, ask the user during onboarding or before the first article:
"Where does your team keep product data, test results, and customer feedback?
(Google Drive folders, specific URLs, internal wikis, etc.)"

Save their answer: `aeo config data-sources update --data-sources "..."`

## Research order

When researching for an article:

1. **Read custom sources** — `aeo config data-sources show` → follow each pointer
2. **Search Drive** — `aeo drive list` with topic-relevant keywords
3. **Check brand context** — `aeo domain brand` for claims and positioning
4. **Ask the user** — if 1st-party material is insufficient for the topic
5. **External research** — web search, authority sources, competitor sites

Never skip steps 1-3. External research (step 5) supplements 1st-party material, not replaces it.
