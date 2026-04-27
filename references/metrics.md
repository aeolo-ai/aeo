# aeo metrics — Article & Site Performance

Track deployed article performance and site-level search traffic via Google Analytics (GA4) and Google Search Console (GSC).

## Prerequisites

- Domain must have Google Analytics and/or Search Console connected (Dashboard → Settings → Integrations)
- Articles must be deployed (`deploy_status = 'deployed'`) with a `published_url`

## Commands

### `aeo metrics overview`

```bash
aeo metrics overview [-d <domainId>]
```

Returns a markdown table of all deployed articles (up to 20) with:
- **Page Views** (GA4)
- **Clicks** (GSC)
- **Impressions** (GSC)
- Published date

Period: last 30 days.

If Google is not connected, shows article list without stats + a setup prompt.

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

### `aeo metrics traffic`

```bash
aeo metrics traffic [-d <domainId>] [--days=30]
```

Site-level GSC traffic overview (not per-article). Returns:

**Top Queries** — queries driving clicks/impressions to the entire site
**Top Pages** — pages receiving the most search traffic
**Country Breakdown** — traffic by country
**Device Breakdown** — desktop vs mobile vs tablet

`--days` accepts 7, 14, 30, or 90 (default: 30).

Requires Google Search Console connected.

## Interpreting Results

- **AI traffic**: Look for `sourceMedium` containing `chatgpt`, `perplexity`, `gemini`, or `you.com` in the traffic sources table (article-level)
- **High impressions + low clicks** = ranking but not getting clicked → improve title/meta
- **Top queries** show what people actually search to find the article — compare with `target_keywords`
- **Site-level traffic** (`metrics traffic`) shows overall search presence — useful for spotting new keyword opportunities and tracking GEO impact across all pages

## Connector API

| Endpoint | Method |
|----------|--------|
| `/v1/connector/domains/:domainId/metrics/overview` | GET |
| `/v1/connector/domains/:domainId/metrics/article/:contentId` | GET |
| `/v1/connector/domains/:domainId/metrics/traffic?days=N` | GET |

All are read-only (all token scopes).
