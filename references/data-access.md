# Data Access — Research Source Configuration

## Default sources (always available)

These are available to all aeo users without configuration:

| Source | Command | What to look for |
|--------|---------|-----------------|
| Google Drive | `aeo drive list`, `aeo drive read <id>` | Product specs, test results, customer feedback, internal docs |
| Brand Context | `aeo agent context` | Value proposition, key features, competitors |
| Published Content | `aeo content list --status=published` | Existing articles (avoid overlap, find cross-link targets) |
| Visibility Data | `aeo visibility show` | Gap queries, competitor mentions |
| Prompt set health | `aeo prompts health` | Whether the tracked prompts are the right target at all |
| Next topic | `aeo topics next` | One article angle ranked by citation gap, with the evidence behind it (free, read-only, deterministic) |

## Custom sources (per-domain)

Run `aeo config data-sources` to see this domain's configured sources.

If no custom sources are configured, ask the user during onboarding or before the first article:
"Where does your team keep product data, test results, and customer feedback?
(Google Drive folders, specific URLs, internal wikis, etc.)"

Turn their answer into a Source Policy patch.

- Interactive CLI/operator flow: after explicit approval, apply it with
  `aeo config data-sources update --data-sources "..."`
- Background writing job or chat flow: do not write product memory directly.
  Return the `domains.data_sources` patch for review.

## Reference policy — how evidence gets rendered

`aeo config reference-policy` shows whether this brand publishes with outbound
citation links.

| Policy | Body renders evidence as | Import validator |
|--------|--------------------------|------------------|
| `standard` (default) | inline `[Source](URL)` markdown links | requires >= 3 external source URLs |
| `first_party` | prose attribution ("according to the CDC") with no link | source minimum waived; external links stripped to plain text |

**Grounding is identical under both.** You still research, still trace every
quote and number to a real source, still cut what you cannot source. Only the
rendered hyperlink differs. Links to the brand's own domain always survive.

`standard` does not merely allow links — it **requires** them, and an import
with fewer than 3 will be rejected. So a brand that should not link out (a
clinic, a law firm, anyone publishing as the sole authority) must be moved to
`first_party` explicitly; leaving it on the default produces the opposite of
what they want.

After explicit approval (CUD Rule), apply with
`aeo config reference-policy update --policy first_party`.

Flipping the policy affects **articles written from then on**. Already-published
bodies keep their links — the strip runs at import time only.

## Terminology glossary — the brand's approved name per market

`aeo config glossary` shows the renderings this brand requires in each locale.

A locale edition is not a script conversion. 써마지 is `热玛吉` in the mainland
and `鳳凰電波` in Taiwan — different **names**, not different characters, so no
converter produces one from the other. The mapping is data the brand supplies.

```
aeo config glossary update --glossary '{"version":1,"terms":[
  {"id":"thermage","renderings":{"ko":["써마지"],"zh-Hans":["热玛吉"],"zh-Hant":["鳳凰電波"]}}
]}'
```

Renderings are most-preferred first; the rest are accepted synonyms. **Omit a
locale when the brand has no special name there** — for an ordinary noun where
translation or script conversion is correct, omitting it keeps the term out of
the writer's prompt and out of the post-write check.

Optional `matchKeys` changes what counts as the term OCCURRING in a locale
without changing what gets written. Reach for it when a name is a homonym of an
ordinary word:

```json
{ "id": "onda",
  "renderings": { "ko": ["온다"], "zh-Hant": ["Onda"] },
  "matchKeys": { "ko": ["온다(Onda)", "Onda"] } }
```

The device 온다 is spelled exactly like the verb 온다 ("comes"), so without this
every article containing 나온다 or 돌아온다 would be told to name a treatment it
never mentioned. matchKeys are matched literally; renderings also match with a
trailing parenthetical dropped.

**If a term cannot be matched honestly in any form, leave it out.** A demand the
article can never satisfy either corrupts the text or fails the job — worse than
not checking. Symptom to listen for: "it keeps telling us to add a word we
didn't use." That is a glossary edit, never a code change.

When a locale edition is written, the job injects only the terms its source
article actually uses, then verifies each approved rendering is present in the
finished body and refuses to complete while one is missing. Writing a term's
generic category word instead of the approved name does not satisfy it.

Ask for a glossary when a brand publishes in more than one locale and sells
named products, treatments, or packages. After explicit approval (CUD Rule),
apply with the command above — the update REPLACES the whole document, so send
the full term list, not a delta.

## Research order

When researching for an article:

1. **Read custom sources** — `aeo config data-sources` → follow each pointer
2. **Search Drive** — `aeo drive list` with topic-relevant keywords
3. **Check brand context** — `aeo agent context` for claims and positioning
4. **Ask the user** — if 1st-party material is insufficient for the topic
5. **External research** — web search, authority sources, competitor sites

Never skip steps 1-3. External research (step 5) supplements 1st-party material, not replaces it.
