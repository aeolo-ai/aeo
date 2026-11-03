# aeo measure / metrics — Article & Site Performance

Track deployed article performance and site-level search traffic via Google Analytics (GA4) and Google Search Console (GSC).

> **Canonical noun is `measure`; `metrics` is an accepted alias.** `metrics article <id>` = `measure content <id>`. The examples below use the `metrics` alias. Failure diagnostics are `aeo diag report` (see the bottom of this file).
>
> **Retired 2026-11-03:** `measure overview` / `metrics overview`, `measure traffic` / `metrics traffic` / `diagnose traffic`, and `measure visibility`. The spellings still answer, with `traffic pages`, `traffic search-console` and `visibility insights` respectively. Call those directly.

## Prerequisites

- Domain must have Google Analytics and/or Search Console connected and a property/site selected (connect in Traffic, `/traffic-analysis`, or `/connect/google?domain=<id>`; select with `aeo integrations google set`, see [setup.md](setup.md))
- Articles must be deployed (`deploy_status = 'deployed'`) with a `published_url`

## Commands

For every published page with GA4 and Search Console numbers, use `aeo traffic pages --published-only true`. For site-level queries, pages, countries and devices, use `aeo traffic search-console` (both below).

### `aeo metrics article <contentId>`

```bash
aeo metrics article <contentId> [-d <domainId>]
```

Detailed stats for a single article:

**GA4 section:**
- Page Views, Sessions, Active Users
- Traffic Sources breakdown (source/medium + sessions)

**GSC section:**
- Clicks, Impressions, CTR, Avg Position
- Top 10 queries with per-query clicks/impressions/position

Period: last 30 days.

## Interpreting Results

- **AI traffic**: Look for `sourceMedium` containing `chatgpt`, `perplexity`, `gemini`, or `you.com` in the traffic sources table
- **High impressions + low clicks** = ranking but not getting clicked → improve title/meta
- **Top queries** show what people actually search to find the article — compare with `target_keywords`

## Interpreting Results

- **AI traffic**: Look for `sourceMedium` containing `chatgpt`, `perplexity`, `gemini`, or `you.com` in the traffic sources table (article-level)
- **High impressions + low clicks** = ranking but not getting clicked → improve title/meta
- **Top queries** show what people actually search to find the article — compare with `target_keywords`
- **Site-level traffic** (`traffic search-console`) shows overall search presence — useful for spotting new keyword opportunities and tracking GEO impact across all pages

### `aeo traffic <card>` — the Traffic screen, card by card

Each read returns what one part of the dashboard's Traffic page shows, through the same loader and window. Window: `--days 30|60|90|180|365` (ends 3 days ago, like the screen's period picker) or `--start YYYY-MM-DD --end YYYY-MM-DD`. Add `--format json` for the full payload. No credits.

| Screen | Command |
|---|---|
| Overview: GA4 users/sessions/page views, daily trend, channels, countries, AI-assistant sessions by engine + Search Console totals | `aeo traffic summary` |
| Detail cards: sources, landing pages by source, top pages, countries, devices, funnel | `aeo traffic details` |
| AI engine traffic: source → landing → next page, landing ownership | `aeo traffic journey [--view journey\|ownership] [--mode ai\|all]` |
| Purchase/traffic events: summary, recent rows, breakdowns | `aeo traffic events [--kind purchase\|traffic] [--rows true] [--breakdown page\|country\|source\|device]` |
| Custom tracked events | `aeo traffic tracked-events` |
| Published page performance | `aeo traffic pages [--published-only true]` |
| Store inflow from Aeolo content | `aeo traffic attribution [--detail campaign\|key-events\|purchase-trend\|conversions\|content-touches]` |
| AI citations per published article | `aeo traffic citations [--days N]` |
| Indexing tab | `aeo traffic indexing [--verdict <v>] [--q text]` |
| GSC tab: AI Overview / AI Mode impressions | `aeo traffic ai-overview` |
| GSC and GA4 tabs: explorer rankings | `aeo traffic explore --kind gsc\|ga4 [--dimension query\|page\|source\|device] [--country XX]` |
| All tab: headline chart split | `aeo traffic hero [--breakdown channel\|device\|country\|subfolder\|device-daily] [--metric sessions\|totalSessions\|searchSessions\|proxySessions]` |
| Search Console tab: full report (queries, pages, countries, devices, daily, Discover) | `aeo traffic search-console [--source <gscSiteUrl>] [--limit 25]` |
| Managed site / customer-site tracker | `aeo traffic site` |

GA4 detail, journey and event reports are built by a background job, as on the screen. When one is not built yet the read says so and returns the job id; read again in a minute rather than retrying in a loop.

### `aeo diag report --command <cmd>`

```bash
aeo diag report --command "<cmd>" [--status-code <code>] [--response-body <text>] [--context "<one sentence>"]
```

Submit command-execution diagnostics when a command fails. Required flag: `--command`. Optional: `--status-code`, `--response-body`, `--context`. Legacy spellings `aeo report` and `aeo measure report` reach the same command. Not the same as `report snapshot`, which creates a shareable report link.

## Connector API

| Endpoint | Method |
|----------|--------|
| `/v2/connector/domains/:domainId/metrics/article/:contentId` | GET |

All are read-only (all token scopes).
