# Setup, Auth & Onboarding Checklist

Everything needed before running GEO commands: install the CLI, authenticate,
and walk a new brand through the seven onboarding steps.

## Setup check — run before any command

Check that the `aeo` CLI is installed:

```bash
aeo --version
```

If `aeo` is not found, guide installation first (single binary — no Go or Node.js required):

```
## aeo CLI Installation

Install with one command:

curl -fsSL https://skills.aeolo.io | sh

After install, verify: `aeo --version`

## Update

aeo update              # self-update to latest
```

Then verify the agent is authenticated:

```bash
aeo auth status
```

If not logged in, guide the user through authentication (see below).

## /aeo auth — Authentication

### /aeo auth login

```bash
aeo auth login
```

Opens a browser for device-flow authentication. On success, saves the API key and
default domain to `~/.config/aeo/config.json`.

```
## Aeolo Authentication

1. Run `aeo auth login` — this opens a browser for authentication
2. After login, your API key and default domain are saved automatically
3. To switch domains: `aeo domain switch` or `--domain <id>` flag
```

### /aeo auth status

```bash
aeo auth status
```

Shows current credentials (API key hint, active domain, source: config vs env).

### /aeo auth logout

```bash
aeo auth logout
```

Clears stored credentials.

## /aeo domain setup — Onboarding checklist

```bash
aeo domain setup
```

Returns a 7-item checklist. **The same list the dashboard's Brand Setup screen
renders** — one definition, two renderers, so the two surfaces can never
disagree about the same brand.

| # | Item | Clear it with |
|---|------|---------------|
| 1 | Brand family & aliases | `aeo domain brand aliases`, then `aeo domain brand update --family-json '[...]'` |
| 2 | Market & language | `aeo domain brand update --markets ko-KR,en-US` |
| 3 | Topics | `aeo topics candidates` → `aeo topics create` |
| 4 | Prompts | `aeo topics prompts` |
| 5 | First visibility check | `aeo visibility check run` |
| 6 | Google (property selected) | `aeo integrations google set --ga4-property <id> --gsc-site <url>` |
| 7 | Somewhere to publish | `aeo channel add --url <site>` |

**Every ⬜ row carries the exact command that clears it.** Run the hint, then
re-read `aeo domain setup` to confirm the row flipped. If a row does not flip,
say so — do not run the same hint again in a loop.

A "Next" section lists Drive and content strategy. Those are real but do NOT
count toward completion: they are not what stands between a new brand and its
first article, and counting them would leave every healthy brand permanently
short of done.

If the header says **⏳ Still analyzing the site**, the `domain add` crawl is
still building the brand understanding. Wait a minute and re-run rather than
acting on a half-built brand.

## Walking the steps

### 1. Brand family & aliases

The identity roster the mention axis matches on. Read it first — the command
also proposes alias candidates for the markets already selected:

```bash
aeo domain brand aliases
```

It suggests only; nothing is saved. Confirm the spellings with the user, then
write the whole roster:

```bash
aeo domain brand update --family-json '[
  {"name":"EO","role":"self","aliases":{"ko-KR":["이오","EO"]},"contested":["EO"]},
  {"name":"EO planet","role":"product","urls":["eoplanet.com"]}
]'
```

- `role` is one of `self`, `parent`, `sibling`, `product`, `channel`. `sibling`
  is recorded but NOT counted as ours on the mention axis.
- The write **replaces** the roster wholesale. Send the full list every time.
- A spelling that also names an unrelated entity goes in `contested`, never in
  `aliases` — contested spellings are never auto-matched.
- **`'[]'` is a real submission**, not an empty one: it records "reviewed, this
  brand has no family members" and completes the step. Use it for single-brand
  companies rather than leaving the row undone.

### 2. Market & language

```bash
aeo domain brand update --markets ko-KR,en-US,ja-JP
```

Sets the brand's whole reach in one call — the markets it competes in and the
content units it publishes. The FIRST entry is the primary. Bare halves work
too (`--markets ko,US`) and are normalized to canonical units; the reply prints
what was actually written.

`--markets` replaces `--region` / `--language` rather than combining with them.
Passing both is refused, because a silently-ignored region and a silently-widened
reach are both states you would have to re-read the row to discover.

### 3. Topics

Two lanes. The fast one is a server job:

```bash
aeo topics candidates                    # ~2-3 min, then:
aeo topics candidates poll <jobId>
```

Each candidate comes back with its evidence and a `--demand-token`. **Nothing is
created.** Confirm the picks with the user, then create only those:

```bash
aeo topics create --name "레티놀 입문" --demand-token <token>
```

Pass the token — it carries demand the job already paid DataForSEO for. Creating
without it triggers a second, billed measurement.

The deeper lane is the market map (`aeo market-map run` → `populate`), which is
grounded in market terrain rather than the brand snapshot. See
[market-map.md](market-map.md).

> `aeo topics candidates` is NOT `/aeo topics suggest`. The latter is an
> agent-only workflow where you read live context and judge a 3–5 Topic
> architecture ([topic-suggest.md](topic-suggest.md)). Same word, different
> product.

### 4. Prompts

```bash
aeo topics prompts                       # then:
aeo topics prompts poll <jobId>
```

Writes tracked Prompts for every active Topic that has none yet — 20~30 per run,
capped by the plan's remaining slots. Progressive: run it again and it targets
whatever is still empty.

🚨 **Do not use `aeo prompts generate` for this step.** That is the FLAT
generator: it writes prompts with no `topic_id`, so the Topics you just
confirmed stay empty while unattached rows pile up beside them. `prompts
generate` seeds a brand-level set on a domain with no Topics; `topics prompts`
is the onboarding step.

### 5. First visibility check

```bash
aeo visibility check run                 # then:
aeo visibility check poll <jobId>
```

Credit-metered. Needs at least one tracked prompt and a set market.

### 6. Google

```bash
aeo integrations google status           # connected? property selected?
aeo integrations google properties       # GA4 properties to choose from
aeo integrations google sites            # GSC sites to choose from
aeo integrations google set --ga4-property <id> --gsc-site <url>
```

`ready` means SELECTED, not merely connected — a connection with nothing
selected turns the status dots green while every read still returns nothing.

The OAuth grant itself needs a browser. If `status` reports not connected, ask
the user to connect Google in dashboard Settings; you cannot do it from here.

### 7. Somewhere to publish

```bash
aeo channel add --url <site>
```

Wider than "a push integration": a `custom` channel publishes by pull, and the
hosted Aeolo blog publishes with no channel row at all. See
[channels.md](channels.md).

## After 7/7

Only then start the loops — daily content and the weekly report. A loop started
on a half-set-up brand measures a market nobody chose and writes against an
empty topic set.
