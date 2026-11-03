# Dashboard screens

<!-- Generated from packages/core/src/connector/screen-commands.ts. Edit that table, then run:
     pnpm --filter @geo/api exec tsx scripts/sync-screen-commands-doc.ts -->

Each command below reads or changes one dashboard screen through the same route the screen calls, so it returns what the screen shows and enforces the screen's rules (owner-only actions stay owner-only). `aeo <group>` alone prints that group's commands. Writes marked _(write)_, _(execute)_ or _(external)_ change something: show the user what will change and get their yes first.

## members — Members screen: team, roles, invitations

- `aeo members list` — Team members of this domain with role and join date. Same data and rules as the dashboard screen.
- `aeo members invitations` — Pending invitations of this domain (email, role, expiry). Same data and rules as the dashboard screen.
- `aeo members invite --email <email> --role editor|viewer|owner` — Invite a person by email (sends the invitation mail). Owner/editor rules as on the Members screen. Confirm the address and role with the user first. _(external)_
- `aeo members invite-link --role editor|viewer [--expires-in-days N]` — Create a shareable invite link for a role (anyone holding it can join). Confirm with the user first. _(write)_
- `aeo members invite-cancel <invitationId>` — Cancel a pending invitation. _(write)_
- `aeo members role <memberId> --role owner|editor|viewer` — Change a member's role (owners only; the last owner cannot be demoted). memberId is the row id from `members list`. _(write)_
- `aeo members remove <memberId>` — Remove a member from this domain, or leave it yourself (your own member id). Irreversible for that member's access; confirm first. _(write)_

## invitations — Invitations addressed to you

- `aeo invitations mine` — Invitations addressed to the signed-in user, across domains (the Home invitation banner).
- `aeo invitations decline <invitationId>` — Decline an invitation addressed to the signed-in user. _(write)_
- `aeo invitations accept --token <token>` — Accept an invitation from its link token (joins that brand). _(write)_

## account chat-usage — Agent chat usage meter

- `aeo account chat-usage` — Agent chat usage meter for the selected domain (used/limit and reset), as on Settings › Billing.

## account plans — Purchasable plans (read only)

- `aeo account plans` — Purchasable plans with prices and credits, as on the Plans screen. Reading only: the agent never starts a checkout.

## report monthly — Stored monthly report (Reports screen)

- `aeo report monthly [--month YYYY-MM | --start YYYY-MM-DD --end YYYY-MM-DD] [--time-zone Area/City]` — The stored report on the Reports screen for a month (default: this month, UTC) or a date range: score, trend, what changed, traffic, work completed, index status, per-engine scores. Reads the stored copy; it does not generate one.

## report delivery — Report email delivery

- `aeo report delivery` — Report email delivery: on/off and recipients.
- `aeo report delivery set --input-json '{"enabled":true,"recipients":["a@x.com"],"locale":"en|ko","timeZone":"Asia/Seoul","revision":N}'` — Turn report emails on/off and replace recipients, language and time zone. Read `report delivery` first and send its revision back. _(write)_

## automation calendar — Automation calendar month

- `aeo automation calendar [--month YYYY-MM]` — The Automation calendar month: every booked writing, visibility, review and deploy event with totals.

## automation schedules — Save Automation schedules with every field

- `aeo automation schedules save --input-json '{"schedules":[...]}'` — Save Automation schedules with every field the screen edits (per-destination target_channel_id, engines, articleLength, locales, reference, hero_sweep, research/locale task rows). Read `automation schedules --format json` first and send back the edited rows. _(write)_

## automation slack — Slack review alerts and channel

- `aeo automation slack` — Slack review alert settings, including the chosen channel.
- `aeo automation slack channels` — Slack channels the connected workspace offers for review alerts.
- `aeo automation slack set --input-json '{"enabled":true,"slack_channel_id":"C…"}'` — Save Slack review alert settings (on/off, channel). Read `automation slack` first. _(write)_
- `aeo automation slack test [--channel-id C…]` — Send a test review alert to Slack. _(external)_

## automation email — Email review alert test

- `aeo automation email test [--recipients a@x.com] [--locale ko|en]` — Send a test review alert email. _(external)_

## research — Research plans and briefs

- `aeo research list` — Research plans and their briefs (status, briefs written, sources).
- `aeo research log <planId>` — One research plan's run log.
- `aeo research write <planId> --brief <index>` — Start writing the article for brief N of a research plan (credit-metered like content generate). _(execute)_
- `aeo research delete <planId>` — Delete a research plan (the screen's rules on which plans may be deleted apply). _(write)_

## planning — Planning allocation, settings, preview, queue, theses

- `aeo planning allocation [--channel-id <id>]` — Planning allocation: the topic mix ratios, expected revision and per-channel overrides.
- `aeo planning allocation set --input-json '{"ratios":{...},"expectedRevision":N}' [--channel-id <id>]` — Save the planning allocation (send the revision you read). _(write)_
- `aeo planning allocation reset --channel-id <id> --input-json '{"expectedRevision":N}'` — Drop a channel's allocation override back to the domain default. _(write)_
- `aeo planning settings [--channel-id <id>]` — The composer's planning settings (method, language, channel defaults).
- `aeo planning settings set --input-json '{...}' [--channel-id <id>]` — Save the composer's planning settings. Read `planning settings` first. _(write)_
- `aeo planning preview --method coverage|competition|external --language <l>` — Next-topic candidates for one planning method, as the Planning screen previews them.
- `aeo planning choices` — Topics queued as the next articles (the Planning screen's queue).
- `aeo planning theses --language <l> --market <code>` — Build (or read cached) topic theses for a language and market, as the Planning screen does. May take minutes. _(execute)_
- `aeo planning strategist [--before-revision N]` — The planning strategist's plan and its conversation (/contents/strategist).
- `aeo planning strategist ask --input-json '{"requestId":"uuid","expectedRevision":N,"message":"…","language":"ko|en"}'` — One turn of the planning strategist (it may revise the plan). _(execute)_

## competitors — Competitor ranking and playbooks

- `aeo competitors list` — Competitor ranking with playbook status (active/archived).
- `aeo competitors show <competitorKey>` — One competitor's insights (what they publish and where they win).
- `aeo competitors learning <competitorKey>` — Stored cross-article learning for one competitor (never generated on read).
- `aeo competitors activate <competitorKey>` — Activate a competitor playbook. _(write)_
- `aeo competitors archive <competitorKey>` — Archive a competitor playbook. _(write)_
- `aeo competitors plan --input-json '{"planId":"…","checkId":"…",…}'` — Create a content plan from a competitor (as the screen's "plan from competitor"). _(execute)_
- `aeo competitors merges` — Competitor merge groups (names merged into one brand), revision and rebuild job.
- `aeo competitors merge --input-json '{"checkId":"…","memberKeys":["a","b"],"representativeKey":"a","expectedRevision":N}'` — Merge competitor names into one brand, as the rankings card. _(write)_
- `aeo competitors unmerge <mergeId> --expected-revision N` — Undo a competitor merge. _(write)_
- `aeo competitors track --input-json '{"name":"…","domain":"example.com"}'` — Track a competitor for comparison (the comparison picker's add). _(write)_

## content history — Article revision history

- `aeo content history <id> [--limit 50]` — An article's revision history (who/when/what changed).

## content revision — Article revisions and restore

- `aeo content revision <id> <revision>` — One stored revision of an article.
- `aeo content revision restore <id> <revision> [--expected-updated-at ISO]` — Restore an article to a stored revision (the current text is kept as a new revision). _(write)_

## content comments — Article comment threads

- `aeo content comments <id>` — Comment threads on an article (anchors, messages, status).
- `aeo content comments create <id> --input-json '{"anchorQuote":"…","body":"…"}'` — Start a comment thread on a quoted passage. _(write)_
- `aeo content comments reply <id> <threadId> --body "…"` — Reply in a comment thread. _(write)_
- `aeo content comments update <id> <threadId> --input-json '{"status":"resolved"}'` — Resolve/reopen or edit a comment thread. _(write)_
- `aeo content comments delete <id> <threadId>` — Delete a comment thread. _(write)_

## content schedule — Timed publishing

- `aeo content schedule --input-json '{"items":[{"content_id":"…","publish_at":"ISO",…}]}'` — Book timed publishing for articles (the Articles screen's schedule). Items as the screen sends them. _(write)_

## content unschedule — Cancel timed publishing

- `aeo content unschedule --input-json '{"items":[{"content_id":"…"}]}'` — Cancel booked timed publishing. _(write)_

## content job — Stop or retry writing jobs

- `aeo content job stop <jobId>` — Stop a writing job that is still running (content job cancel only deletes finished jobs). _(write)_
- `aeo content job retry <jobId>` — Retry a failed writing job with its original inputs (credit-metered). _(execute)_

## content edition — Hand-written locale edition

- `aeo content edition manual <id> --language <locale>` — Add an empty locale edition to write by hand (no credits; content editions generates one). _(write)_

## content batch — Bulk article change/delete

- `aeo content batch --input-json '{"ids":["…"],"updates":{"status":"…"}}'` — Change several articles at once, as the Articles screen's selection bar does. _(write)_
- `aeo content batch delete --input-json '{"ids":["…"]}'` — Delete several articles at once. Confirm the list first. _(write)_

## content legal — Regulatory/legal review state

- `aeo content legal <id>` — An article's regulatory/legal review state: submissions, findings and who approved.

## content og-assets — Cover/sharing-card template assets

- `aeo content og-assets` — The cover frame and sharing-card template assets (logo, background, frame).

## content og-preview — Cover/sharing-card preview

- `aeo content og-preview --input-json '{"sampleTitle":"…",…}'` — Render the sharing-card/cover template preview for a sample title.

## review — Publication review journal, queue and general approval

- `aeo review journal [--queue]` — Publication review journal: what was reviewed, changed and approved, or (--queue) what waits for review.
- `aeo review approve <contentId> --fingerprint <f>` — The review journal's general publication approval for an article (not the professional legal approval). Fingerprint from content legal. _(write)_

## writing — Writing inputs and web research scope

- `aeo writing inputs` — Which inputs writing uses (brand notes, strategy, sources, web search…) and their state.
- `aeo writing inputs set <node> --input-json '{...}'` — Turn one writing input on/off or change it, as the screen's switchboard. _(write)_
- `aeo writing web-search` — Web research scope for writing (allowed/blocked sources).
- `aeo writing web-search set --input-json '{...}'` — Save the web research scope. Read `writing web-search` first. _(write)_
- `aeo writing models` — Which AI vendor writes each writing stage.

## brand settings — Regulatory profile, measured targets, store link UTM, OG template

- `aeo brand settings` — Brand screen settings brand update does not show: regulatory_profile, prompt_targets (measured region × languages), store_link_params, og_template and the curated brand visual assets.
- `aeo brand settings set --input-json '{"regulatory_profile":{…}|"prompt_targets":[…]|"store_link_params":{…}|"og_template":{…}}'` — Save one or more of those Brand screen settings. Only these keys are accepted; everything else stays with brand update. _(write)_

## brand visual-assets — Curated brand visual assets

- `aeo brand visual-assets set --input-json '[{"url":"…","kind":"…",…}]'` — Replace the curated brand visual assets (the crawled images the Brand screen lets you keep or drop). Read `brand settings` first. _(write)_

## brand locations — Brand locations and location check

- `aeo brand locations` — Brand locations (stores/clinics) on the Brand screen.
- `aeo brand locations add --input-json '{"name":"…","locality":"…","streetAddress":"…","telephone":"…","siteUrl":"https://…","gbpUrl":"https://…"}'` — Add a brand location. _(write)_
- `aeo brand locations update <locationId> --input-json '{...}'` — Edit a brand location. _(write)_
- `aeo brand locations delete <locationId>` — Delete a brand location. _(write)_
- `aeo brand locations check` — Start the location visibility check (pilot brands, as on screen). _(execute)_
- `aeo brand locations check poll <jobId>` — Poll a location check.

## site create — Create the managed site

- `aeo site create` — Create the managed site for this brand (Sites screen). _(write)_

## site remove — Remove the managed site

- `aeo site remove` — Remove the managed site. Its pages stop being served; confirm first. _(write)_

## site indexing — Managed site search indexing switch

- `aeo site indexing` — Whether search engines may index the managed site (open/hold).
- `aeo site indexing set --enabled true|false` — Open or hold search-engine indexing of the managed site. _(write)_

## site host — Managed site custom host

- `aeo site host` — The managed site's custom host binding and its DNS/verify state.
- `aeo site host preflight --host <host>` — Check whether a host can be bound (DNS, conflicts).
- `aeo site host bind --host <host> [--mode path]` — Bind a custom host to the managed site. _(write)_
- `aeo site host verify-path --article-path <path>` — Verify a path-mode host serves our articles at the given path.
- `aeo site host unbind` — Remove the managed site's custom host. _(write)_

## site locales — Managed site languages

- `aeo site locales set --locales ko,en,ja` — Set the managed site's languages. _(write)_

## site folders — Create/rename/delete/sync managed site folders

- `aeo site folders create --input-json '{"name":"…","path":"/guides"}'` — Create a managed site folder. _(write)_
- `aeo site folders update <folderId> --input-json '{"name":"…","path":"…"}'` — Rename or move a managed site folder. _(write)_
- `aeo site folders delete <folderId>` — Delete a managed site folder. _(write)_
- `aeo site folders sync` — Sync managed site folders with the destination. _(write)_

## channels primary — Primary destination

- `aeo channels primary <channelId>` — Make a destination the primary one. _(write)_

## channels settings — Destination writing rules and footer

- `aeo channels settings set <channelId> --input-json '{"writing_rules":"…"|"footer":{…}}'` — Save a destination's writing rules or the footer under its articles (the detail tabs channels update does not carry). _(write)_

## channels html — Destination HTML capture

- `aeo channels html <channelId> [--status-only true] [--version-id <id>]` — A destination's captured HTML (or its capture status).
- `aeo channels html capture <channelId>` — Capture a destination's HTML now. _(execute)_
- `aeo channels html apply <channelId> --input-json '{...}'` — Apply a captured HTML version to the destination. _(write)_

## index — Sitemap submit and IndexNow key

- `aeo index sitemap submit [--channel-id <id>]` — Submit the sitemap to Google Search Console for one site. _(external)_
- `aeo index indexnow key --key <key> [--channel-id <id>]` — Adopt the IndexNow key a site already serves (as the Indexing screen's key step). _(write)_
- `aeo index indexnow submit [--channel-id <id>]` — Send a site's article URLs to IndexNow now (any site, managed ones included). _(external)_

## sources — Source library: channels, documents, suggestions, Drive

- `aeo sources list` — The source library: registered channels (owned/other, on/off) the writer may draw from.
- `aeo sources documents [--channel-id <id>]` — Ingested source documents (articles, videos, Drive docs).
- `aeo sources document <documentId>` — One ingested source document.
- `aeo sources suggestions` — Suggested source candidates and which are reserved.
- `aeo sources add --input <url-or-handle> [--owned true|false] [--scope <s>]` — Register a source channel (site, YouTube, Instagram…). _(write)_
- `aeo sources update <sourceId> --owned true|false` — Mark a source as the brand's own channel or someone else's. _(write)_
- `aeo sources usage <sourceId> --enabled true|false` — Switch a source on or off for writing. _(write)_
- `aeo sources delete <sourceId>` — Remove a source channel from the library. _(write)_
- `aeo sources article --url <url> [--channel-id <sourceId>]` — Ingest one web article as a source document. _(execute)_
- `aeo sources video --video-id <id> [--channel-id <sourceId>]` — Ingest one video (transcript) as a source document. _(execute)_
- `aeo sources convert <documentId> --input-json '{...}'` — Turn a source document into an article draft, as the screen's convert. _(execute)_
- `aeo sources reserve <candidateId>` — Reserve a suggested source for the next article. _(write)_
- `aeo sources unreserve <candidateId>` — Release a reserved source. _(write)_
- `aeo sources videos <sourceId> [--continuation <token>]` — A source channel's videos with ingest state (the video picker).
- `aeo sources collect <sourceId>` — Collect an Instagram source channel now. _(execute)_
- `aeo sources reserve-item --input-json '{"channelRowId":"…","externalId":"…","url":"…","title":"…"}'` — Reserve a source item that has no candidate row yet. _(write)_
- `aeo sources brief <documentId> [--input-json '{"targetPromptId":"…","topicId":"…"}']` — Pull a writing brief out of a source document. _(execute)_
- `aeo sources drive-folder --folder-id <id>` — Register a Google Drive folder as a source channel. _(write)_
- `aeo sources drive-document --file-id <id> [--channel-id <sourceId>]` — Import one Google Drive document as a source document. _(execute)_

## pages all — Every page across destinations

- `aeo pages all` — Every page across destinations (the Pages screen's all-channels list).

## pages revisions — Page revisions

- `aeo pages revisions <pageId> [--channel-id <id>]` — A page's stored revisions.

## pages revision — Restore a page revision

- `aeo pages revision restore <pageId> <revision> [--channel-id <id>] --input-json '{"expectedRevision":N}'` — Restore a page to a stored revision. _(write)_

## traffic home — Home traffic card

- `aeo traffic home [--growth true]` — The Home traffic card: AI-referral sessions and, with --growth, the growth comparison the card shows.

## domain setup — Brand setup steps

- `aeo domain setup publish-skip` — Skip the publish step of brand setup (connect a destination later). _(write)_

## domain onboarding-channel — Brand setup destination

- `aeo domain onboarding-channel` — The brand-setup destination choice and its connection state.
- `aeo domain onboarding-channel set --input-json '{"choice":{…},"revision":"…"|null}'` — Save the brand-setup destination choice (send the revision you read). _(write)_
- `aeo domain onboarding-channel complete` — Finish the brand-setup destination step. _(write)_
- `aeo domain onboarding-channel retry` — Retry connecting the brand-setup destination. _(execute)_

## brand family — Brand family discovery

- `aeo brand family discover` — Discover the brand family (parent, sibling, product brands) as brand setup does; save the result with brand update --family-json. _(execute)_

## brand competitors — Competitor discovery

- `aeo brand competitors` — Brand setup's competitor discovery status and the competitors it found.
- `aeo brand competitors discover` — Start or continue brand setup's competitor discovery; save the picks with brand update --competitors-json. _(execute)_

## integrations drive — Google Drive folder connect/disconnect

- `aeo integrations drive disconnect` — Disconnect Google Drive from this brand (Brand › Sources). _(write)_
- `aeo integrations drive connect --folder-id <id>` — Check a shared Drive folder and save it as the brand's folder. _(write)_

## channels preview — Destination previews

- `aeo channels preview <channelId>` — A destination's captured preview status.
- `aeo channels preview capture <channelId> [--source-url <url>]` — Capture a destination preview now. _(execute)_
- `aeo channels preview thumbnails` — Preview thumbnails of every destination.

## channels tone — Destination tone analysis

- `aeo channels tone run --input-json '{"provider":"…","channelId":"…",…}'` — Run a destination's tone analysis, as its Tone tab. _(execute)_
- `aeo channels tone jobs [--provider <p>]` — Tone analysis jobs for this brand.
- `aeo channels tone job <jobId>` — One tone analysis job and its result.
- `aeo channels tone delete <jobId>` — Delete a tone analysis. _(write)_

## channels shopify — Shopify blogs per language

- `aeo channels shopify locale-blog <channelId> --language <l>` — Whether a Shopify destination is ready to publish one language (blog, template).
- `aeo channels shopify locale-blog create <channelId> --language <l> --title "…"` — Create a Shopify blog for one language. _(external)_
- `aeo channels shopify locale-blog set <channelId> --language <l> --blog-id <id>` — Use an existing Shopify blog for one language. _(write)_
- `aeo channels shopify locale-blog template <channelId> --language <l> --template-suffix <s>` — Set the Shopify template suffix for one language's blog. _(external)_
- `aeo channels shopify blogs <channelId>` — The Shopify shop's blogs, to pick a publish target.

## site click-preview — Managed site click map

- `aeo site click-preview --path </page> [--host <host>]` — Where visitors clicked on one managed-site page (the Sites visit view's click map).

## content topic-preview — Composer next topic

- `aeo content topic-preview --language <l> --channel-id <id>` — The topic the composer would write next for one language and destination.

## content family — Article language family

- `aeo content family <id>` — An article's language family: the original and every edition with status (the edition switcher).

## content schema — Article structured data

- `aeo content schema <id>` — The structured data (JSON-LD) the article ships with.

## content images — Images linked to an article

- `aeo content images <id> [--limit 40]` — Images linked to an article (the thumbnail picker).

## content pull-wordpress — Pull the live WordPress post back

- `aeo content pull-wordpress <id>` — Copy the live WordPress post back into the article (the fix when a redeploy refuses with REMOTE_MODIFIED). _(write)_

## content unpublish-pangolingo — Unpublish from Pango Lingo

- `aeo content unpublish-pangolingo <id>` — Withdraw an article and its editions from a Pango Lingo newsroom (content unpublish does not reach Pango Lingo). _(external)_

## audit history — Site audit history

- `aeo audit history [--limit 10] [--channel-id <id>]` — Past site-audit runs, optionally for one site (the audit tab's history).

## report check — Stored report for one check

- `aeo report check <checkId>` — The stored report for one visibility check (reads the stored copy; does not generate).

## report saved — Saved AI-written reports

- `aeo report saved [<sessionId>]` — Saved AI-written reports; with an id, open one.

## report opportunities — Visibility × Search Console opportunities

- `aeo report opportunities [--limit 20]` — Opportunities that cross AI visibility with Search Console (the traffic report's gap-first list).

## prompts deleted — Deleted and untracked prompts

- `aeo prompts deleted [--q text] [--offset N]` — Deleted and untracked prompts (the restore list).

## prompts restore — Restore a prompt

- `aeo prompts restore <promptId>` — Restore a deleted or untracked prompt. _(write)_

## prompts research — Prompt research context

- `aeo prompts research <promptId>` — A prompt's topic and market-map context for its region and language.

## prompts coverage — Tracked-prompt quota

- `aeo prompts coverage` — Tracked-prompt quota (used/limit) and which prompts count toward measurement.

## prompts experiments — Prompt experiments

- `aeo prompts experiments [<jobId>] [--prompt-id <id>]` — Prompt experiment runs and their results (when experiments are on for this brand).

## visibility atlas-products — Atlas product assignment

- `aeo visibility atlas-products --country <code> --concern-id <id>` — Which catalog products Atlas assigns to a concern in a country (staff and Atlas pilot brands, as on screen).

## traffic ai-impressions — Collect Google AI impressions

- `aeo traffic ai-impressions check` — Ask for Google AI impressions to be collected now (the report is traffic ai-overview). _(execute)_
- `aeo traffic ai-impressions disconnect` — Stop collecting Google AI impressions. _(write)_

## reference jobs — Reference analysis jobs

- `aeo reference jobs --media <type>` — Reference analysis jobs for one media type (visual_asset, instagram_reels, …), as the reference panels list them.

## brand product — Edit or remove a product

- `aeo brand product update <productId> --input-json '{"title":"…","og_image_url":"…","price":"…"}'` — Override a product's scraped title, image or price. _(write)_
- `aeo brand product delete <productId>` — Remove a product from the catalog. _(write)_

## channels request — Request an integration

- `aeo channels request --input-json '{"channel_types":["…"],"other_channel":"…"}'` — Ask the Aeolo team to add an integration (the Pro+ request button). _(external)_

## channels expansion — Channel recommendations from cited sources

- `aeo channels expansion --input-json '{"citedSources":[{"domain":"…","citations":N}],"language":"ko"}'` — Recommend more channels from the sources AI cites (the citation moments panel). _(execute)_

## channels social-profile — Connected social profile

- `aeo channels social-profile <channelId>` — The connected social account's profile.

## channels custom-site — Custom site verification tag and check

- `aeo channels custom-site tag <channelId>` — The verification meta tag to paste into a custom site's <head>.
- `aeo channels custom-site check <channelId>` — Check whether a custom site serves our newest article.

## channels wordpress — WordPress site info

- `aeo channels wordpress site <channelId>` — A WordPress site's languages, default edition and translation linking (live probe).
- `aeo channels wordpress categories <channelId>` — The WordPress site's categories, to pick a publish target.

## channels cafe24 — Cafe24 boards, language routes and store SEO

- `aeo channels cafe24 boards <channelId> [--shop-no 1]` — The Cafe24 shop's boards, to pick a publish target.
- `aeo channels cafe24 routes set <channelId> --input-json '{"routes":[{"locale":"ko","language_code":"ko","shop_no":1,"board_no":N,"domain":"…"}],"link_sitemap":true}'` — Route each language to a Cafe24 shop board. _(external)_
- `aeo channels cafe24 seo <channelId>` — The Cafe24 store's SEO settings (title, meta, robots, llms.txt, sitemap).
- `aeo channels cafe24 seo set <channelId> --input-json '{"shop_no":1,"settings":{…}}'` — Save the Cafe24 store's SEO settings (live store). _(external)_
- `aeo channels cafe24 seo restore <channelId> --shop-no N` — Restore the Cafe24 store's original SEO settings from the saved backup. _(external)_
- `aeo channels cafe24 seo link-sitemap <channelId>` — Add our sitemap line to each Cafe24 shop's robots.txt. _(external)_
- `aeo channels cafe24 shops <channelId>` — The Cafe24 mall's shops (languages) for routing and SEO.

## brand social-links — Brand social profile links

- `aeo brand social-links` — The brand's social profile links (sameAs URLs in structured data).
- `aeo brand social-links set --links https://…,https://…` — Replace the brand's social profile links. _(write)_

## audit show — One audit run

- `aeo audit show <auditId>` — One past site-audit run in full.

## channels requests — Integration requests

- `aeo channels requests` — Integration requests sent to the Aeolo team and their state.

## Not reachable on purpose

- Settings › Developer API tokens (list/create/revoke): Credential minting stays with the person: an agent that can mint tokens can hand out standing access.
- Settings › Account delete, Settings › Domains delete: Irreversible account/brand deletion stays a person's click.
- Settings › Billing checkout, top-up, customer portal, store billing sync: Purchases and payment-method changes are never made by the agent.
- Settings › Account dashboard language: Client-only preference of the viewer's browser.
- Article professional legal approval (legal-publication-review approvals): Approval by a qualified reviewer must be that person's own act; content approve states it does not grant it.
- Gallery file upload: Needs the file bytes from the viewer's device; the agent imports images by URL through image tools instead.
