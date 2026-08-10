# Channel Management

Connect, update, and disconnect the destinations Aeolo publishes to (or pulls a
feed from). All channel commands are scoped to the active domain.

## /aeo channel list (or /aeo domain channels)

Show all channels connected to the current domain. Returns a markdown table with
label, platform, URL, and channel ID. The primary channel is marked with a star.

```bash
aeo channel list
aeo domain channels   # alias
```

## /aeo channel add

Add a new channel to the current domain. Type is auto-detected from the URL if
`--type` is omitted.

```bash
aeo channel add --url https://www.threads.net/@mybrand --type threads --label "Threads Main"
```

Types: `shopify`, `vercel`, `linkedin`, `threads`, `reddit`, `instagram`, `x`, `website`, `custom`

**`--type custom` — connect the customer's own site as a Content Feed pull-channel.**
Use this when they publish on a site we don't push to (their own blog, a static or
hand-built site). It provisions the feed in one step:

```bash
aeo channel add --type custom --url https://yourdomain.com/blog
```

- mints a read-scoped API key (returned once) to pull the feed,
- returns the authed feed URL with `?base=` so canonical points at their domain,
- returns an ownership verify `<meta>` tag to paste into their site `<head>`.

Then wire the feed **server-side** per the guide it links (render `content_html`, inline
`_geo.schema_jsonld`, set the page canonical to `_geo.canonical`). A client-only
widget defeats GEO. See also `aeo content feed` for the raw feed URLs + JSON Feed contract.

## /aeo channel update

Update an existing channel's URL, type, or label. Also sets publish behavior:
`--auto-publish true|false` and `--publish-target <id>` (shopify blog / wordpress
category / cafe24 board, picked by the channel's platform).

```bash
aeo channel update <channel-id> --label "New Label" --type linkedin
```

## /aeo channel indexing

Toggle IndexNow auto-indexing for a channel. Add `--backfill` to submit existing
published articles.

```bash
aeo channel indexing <channel-id> --enabled true --backfill
```

## /aeo channel delete

Delete a non-primary channel. Primary channels cannot be deleted.

```bash
aeo channel delete <channel-id>
```

## /aeo channel connect

Generate an OAuth URL and open the browser for social-platform authorization
(threads, linkedin, reddit).

```bash
aeo channel connect <channel-id>
```

The browser opens the platform's OAuth page. On success, it redirects to the dashboard.

## /aeo channel disconnect

Remove the OAuth integration from a channel without deleting the channel row.

```bash
aeo channel disconnect <channel-id>
```

## /aeo channel voice

Read the selected channel-voice / reference-style evidence for the channel
(`--provider`, `--url`) — the same data as `reference style`. See
[tov-extract.md](tov-extract.md).

## /aeo blog show

Report the hosted Aeolo blog for the active domain: its home URL, any bound custom
host, and — when a host is bound but not yet verified — the DNS record still to add.

```bash
aeo blog show
```

The hosted blog is opt-in. If it was never created this says so rather than minting
a URL that 404s; create it from **Channels** in the dashboard or by deploying an
article with `aeo content deploy <id> --target blog`.

## /aeo blog bind

Serve the hosted blog on a host the customer owns, e.g. `pages.yourbrand.com`.

```bash
aeo blog bind pages.yourbrand.com
```

Returns the CNAME record to add at their DNS provider. **The customer adds that
record themselves** — it lives in their DNS zone, and nothing we run can write it
for them. Until it resolves, the command reports `waiting for DNS`; that is pending,
not failed, so do not unbind and retry. Re-run `aeo blog show` to check.

Three things to tell them:

- **On Cloudflare, keep the record DNS only (grey cloud).** Proxying it makes
  certificate verification need an extra TXT record, and Cloudflare then caches the
  HTML so our page-view and AI-crawler telemetry stops recording hits.
- **Their main site is untouched.** This is a separate subdomain record, so it works
  whether they run Shopify, Cafe24, WordPress, or anything else.
- Once bound, the host becomes the canonical address: every canonical tag, `og:url`,
  sitemap entry, and newly published article URL uses it. The `{sub}.aeolo.blog`
  address keeps resolving and permanently redirects, so anything an AI engine
  already cited still reaches the article.

This is a subdomain binding, not subdirectory hosting. Serving at
`brand.com/blog/*` needs a reverse proxy inside the customer's own edge, which we
cannot provision for them — if they publish on their own CMS, connect it as a
channel instead (`aeo channel add`) and publish there directly.

## /aeo blog unbind

Stop serving on the custom host. The blog reverts to `{sub}.aeolo.blog` rather than
going dark — the reserved name is never released, which is why re-binding later
returns the same address.

```bash
aeo blog unbind
```

Tell them to remove the CNAME record afterwards; it no longer points anywhere we serve.

> **CUD Rule applies** to `channel add`, `channel update`, `channel delete`,
> `channel connect`, `channel disconnect`, `channel indexing`, `blog bind`, and
> `blog unbind`. Show what you are about to change and ask "Proceed?" before calling.
