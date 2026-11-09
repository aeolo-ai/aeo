# Dashboard screens

<!-- Generated from packages/core/src/connector/screen-commands.ts. Edit that table, then run:
     pnpm --filter @geo/api exec tsx scripts/sync-screen-commands-doc.ts -->

Each command below reads or changes one dashboard screen through the same route the screen calls, so it returns what the screen shows and enforces the screen's rules (owner-only actions stay owner-only). `aeo <group>` alone prints that group's commands. Writes marked _(write)_, _(execute)_ or _(external)_ change something: show the user what will change and get their yes first.

## members — Members screen: team, roles, invitations

- `aeo members list` — Team members of this domain with role and join date. Same data and rules as the dashboard screen.
- `aeo members invitations` — Pending invitations of this domain (email, role, expiry). Same data and rules as the dashboard screen.
- `aeo members teams` — What the signed-in person belongs to: the team they own (members, pending invitations), teams they joined, and domains they were invited to alone. Each owner has one team, and its members reach all the owner's domains.
- `aeo members team-name --name <name>` — Rename the team you own. The default name is "{Owner}'s Team". _(write)_
- `aeo members team-invite --email <email> --role editor|viewer` — Invite a person to the team you own (sends the invitation mail). They reach ALL your domains, including future ones; to share one domain use `members invite`. Confirm first; the response lists them in team_domains. _(external)_
- `aeo members team-invite-cancel <invitationId>` — Cancel a pending invitation to your team. _(write)_
- `aeo members team-role <userId> --role editor|viewer` — Change a team member's role on all your domains. _(write)_
- `aeo members team-remove <userId>` — Remove a person from your team and all your domains. Confirm first. _(write)_
- `aeo members team-leave --owner-id <ownerId>` — Leave a team you joined (owner_id from `members teams`). You lose all its domains. Confirm first. _(write)_
- `aeo members invite --email <email> --role editor|viewer|owner` — Invite to this domain only; for every domain of the team use `members team-invite`. Invite a person by email (sends the invitation mail). Owner/editor rules as on the Members screen. Confirm the address and role with the user first. _(external)_
- `aeo members invite-link --email <email> --role editor|viewer` — Create an invite link addressed to one email, without sending mail. This domain only. _(write)_
- `aeo members invite-cancel <invitationId>` — Cancel a pending invitation. _(write)_
- `aeo members role <memberId> --role owner|editor|viewer` — Change a member's role (owners only; the last owner cannot be demoted). memberId is the user_id from `members list`. _(write)_
- `aeo members remove <memberId>` — A team member is removed with `members team-remove`. Remove a member from this domain, or leave it yourself (your own user_id; memberId is the user_id from `members list`). Irreversible for that member's access; confirm first. _(write)_

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
- `aeo planning allocation set --input-json '{"ratios":{...},"expectedRevision":N,"expectedDefaultRevision":N}' [--channel-id <id>] [--query settingsOnly=true]` — Save the planning allocation (send the revisions you read; both are required for a channel override). _(write)_
- `aeo planning allocation reset --channel-id <id> --input-json '{"expectedRevision":N}'` — Drop a channel's allocation override back to the domain default. _(write)_
- `aeo planning settings [--channel-id <id>]` — The composer's planning settings (method, language, channel defaults).
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
- `aeo competitors tracked` — Tracked competitors as the comparison picker lists them.

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
- `aeo content comments agent <id> <threadId> --block-html '<p>…</p>'` — Ask the editor's AI to answer a comment thread with a proposed patch (chat-metered). _(execute)_

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

## writing — Writing inputs, web research scope, per-site review language and regulatory review, planning overview

- `aeo writing review-language --channel-id <id>` — The language reviewers read one blog destination's drafts in (the brand default unless that site overrides it).
- `aeo writing review-language set --channel-id <id> --language <locale>` — Override the review language for one blog destination (the brand default is brand settings set writing_review_language). _(write)_
- `aeo writing planning-overview [--channel-id <id>]` — The planning overview the agent panel shows: what the next articles are built from, for the brand or one site.
- `aeo writing regulatory --channel-id <id>` — One site's regulatory review setting (the brand default unless that site overrides it).
- `aeo writing regulatory set --channel-id <id> --input-json '{"vertical":"medical|cosmetics|legal","review":"off|annotate|annotate_fix","targetMarket":"KR","notes":"…"}'` — Override the regulatory review for one site. _(write)_
- `aeo writing inputs` — Which inputs writing uses (brand notes, strategy, sources, web search…) and their state.
- `aeo writing inputs set <node> --input-json '{...}'` — Turn one writing input on/off or change it, as the screen's switchboard. _(write)_
- `aeo writing web-search` — Web research scope for writing (allowed/blocked sources).
- `aeo writing web-search set --input-json '{...}'` — Save the web research scope. Read `writing web-search` first. _(write)_
- `aeo writing models` — Which AI vendor writes each writing stage.
- `aeo writing documents` — The three writing documents (notes, strategy, policy) and their snapshot hash.
- `aeo writing documents save --input-json '{"action":"review|save","kind":"notes|strategy|policy","value":"…","snapshotHash":"…","reviewToken":"…"}'` — Review or save one writing document with the screen's stale-check. _(write)_

## brand settings — Regulatory profile, measured targets, store link UTM, OG template

- `aeo brand settings` — Brand screen settings brand update does not show: regulatory_profile, prompt_targets (measured region × languages), store_link_params, og_template, writing_review_language (the language reviewers read drafts in) and the curated brand visual assets.
- `aeo brand settings set --input-json '{"regulatory_profile":{…}|"prompt_targets":[…]|"store_link_params":{…}|"og_template":{…}|"writing_review_language":"ko"}'` — Save one or more of those Brand screen settings. Only these keys are accepted; everything else stays with brand update. _(write)_

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
- `aeo index status [--query channelId=…&statuses=0]` — The Indexing screen's scope status: connection, sitemap linked in robots, on-read sitemap reconcile.

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
- `aeo channels shopify detect --url <url>` — The myshopify host behind a store URL.
- `aeo channels shopify status <channelId>` — A Shopify destination's connection: connected, shop and target blog.
- `aeo channels shopify access <channelId> --shop-domain <shop>.myshopify.com` — State of a Shopify app-setup request and, once ready, the install link.
- `aeo channels shopify access request <channelId> --shop-domain <shop>.myshopify.com` — Ask the Aeolo team to set up the store-specific Shopify app (the connect modal's request step). _(external)_

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

## report layouts — Dashboard screen

- `aeo report layouts` — Saved report layout presets for the selected brand. Element order and widths only; data follows the selected report period.

## report layout-create — Dashboard screen

- `aeo report layout-create --input-json <json>` — Save a report layout preset: {name,spec:{version:2,blocks:[{id,type:"section",section,width:"full"|"half"}]}}. Requires editor access. _(write)_

## report layout-update — Dashboard screen

- `aeo report layout-update <layoutId> --input-json <json>` — Update a report layout preset. Read report layouts first; send name, spec, and expectedUpdatedAt from the stored preset. Conflicting edits are rejected. _(write)_

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

- `aeo prompts research <promptId>` — A prompt's topic and market context for its region and language.

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

## channels custom-site — Custom site verification tag, verify and check

- `aeo channels custom-site tag <channelId>` — The verification meta tag to paste into a custom site's <head>.
- `aeo channels custom-site check <channelId>` — Check whether a custom site serves our newest article.
- `aeo channels custom-site verify <channelId> --url <articleBaseUrl>` — Verify a custom site's meta tag and connect it (the Sites screen Verify button). The first verify returns the channel-bound feed key once; set it as a server-only AEOLO_KEY. _(write)_

## channels feed-key — Custom site feed key read and rotate

- `aeo channels feed-key <channelId>` — A verified custom site's feed key state (prefix, created) and its feed URL.
- `aeo channels feed-key rotate <channelId>` — Issue a new channel-bound feed key and revoke the old one. The site stops reading the feed until AEOLO_KEY is replaced. _(write)_

## channels wordpress — WordPress site info

- `aeo channels wordpress site <channelId>` — A WordPress site's languages, default edition and translation linking (live probe).
- `aeo channels wordpress categories <channelId>` — The WordPress site's categories, to pick a publish target.
- `aeo channels wordpress category set <channelId> --input-json '{"category_id":N|null,"category_name":"…"}'` — Set or clear the WordPress category. _(write)_

## channels cafe24 — Cafe24 boards, language routes and store SEO

- `aeo channels cafe24 boards <channelId> [--shop-no 1]` — The Cafe24 shop's boards, to pick a publish target.
- `aeo channels cafe24 routes set <channelId> --input-json '{"routes":[{"locale":"ko","language_code":"ko","shop_no":1,"board_no":N,"domain":"…"}],"link_sitemap":true}'` — Route each language to a Cafe24 shop board. _(external)_
- `aeo channels cafe24 seo <channelId>` — The Cafe24 store's SEO settings (title, meta, robots, llms.txt, sitemap).
- `aeo channels cafe24 seo set <channelId> --input-json '{"shop_no":1,"settings":{…}}'` — Save the Cafe24 store's SEO settings (live store). _(external)_
- `aeo channels cafe24 seo restore <channelId> --shop-no N` — Restore the Cafe24 store's original SEO settings from the saved backup. _(external)_
- `aeo channels cafe24 seo link-sitemap <channelId>` — Add our sitemap line to each Cafe24 shop's robots.txt. _(external)_
- `aeo channels cafe24 shops <channelId>` — The Cafe24 mall's shops (languages) for routing and SEO.
- `aeo channels cafe24 board set <channelId> --input-json '{"board_no":N,"board_name":"…"}'` — Set the Cafe24 board, updating the default shop's language routes too. _(write)_

## brand social-links — Brand social profile links

- `aeo brand social-links` — The brand's social profile links (sameAs URLs in structured data).
- `aeo brand social-links set --links https://…,https://…` — Replace the brand's social profile links. _(write)_

## audit show — One audit run

- `aeo audit show <auditId>` — One past site-audit run in full.

## channels requests — Integration requests

- `aeo channels requests` — Integration requests sent to the Aeolo team and their state.

## visibility sources — Cited sources with filters

- `aeo visibility sources [--check-id <id>] [--engines chatgpt,gemini] [--locales KR,US] [--langs ko,en]` — Cited sources with the Sources lens filters (engine, market, language) for any check.

## visibility competitor-options — Comparison picker brands

- `aeo visibility competitor-options --check-id <id>` — The brands a check can be compared against (the comparison picker); their ids feed visibility matrices --competitor-ids.

## visibility matrices — Head-to-head matrices

- `aeo visibility matrices <checkId> --competitor-ids a,b [--compare-to <checkId>] [--engines …] [--locales …] [--langs …]` — Head-to-head engine/market/topic matrices for chosen competitors on any check, with run-over-run deltas against --compare-to.

## channels detect — Dashboard screen

- `aeo channels detect --url <url>` — What platform/stack a site runs (the add-site form's detection).

## channels all — Dashboard screen

- `aeo channels all` — Destinations exactly as the Sites screen lists them (all columns; managed sites only where enabled).

## visibility quote — Dashboard screen

- `aeo visibility quote [--engines chatgpt,gemini]` — The free credit quote before a visibility check: prompts, engines, credits.

## visibility active-job — Dashboard screen

- `aeo visibility active-job` — The running visibility check, if any: phase, progress, estimate.

## visibility export — Dashboard screen

- `aeo visibility export --check-ids a,b` — The raw AI answers of checks as the CSV the screen downloads.

## visibility change — Dashboard screen

- `aeo visibility change <checkId> [--query compareTo=…&brands=a,b&engines=…&locales=…&langs=…]` — Comparable change between two checks with the screen's brand/engine/market/language filters. Any other screen parameter goes in --query 'a=1&b=2'.

## visibility dashboard — Dashboard screen

- `aeo visibility dashboard [--query view=summary|prompts|report&month=YYYY-MM&timeZone=…]` — The AI search dashboard payload in any of its views and periods. Any other screen parameter goes in --query 'a=1&b=2'.

## visibility checks — Dashboard screen

- `aeo visibility checks [--query limit=10&includeResults=true&view=report|fanout&rows=packed]` — Check history as the screen loads it. Any other screen parameter goes in --query 'a=1&b=2'.

## visibility latest — Dashboard screen

- `aeo visibility latest [--query view=summary&fresh=1]` — The latest check the screen shows (active pointer) and any running job.

## visibility run-screen — Dashboard screen

- `aeo visibility run-screen [--engines chatgpt,gemini]` — Run a visibility check exactly as the screen's button (no engines = the automation's engines; refused while an experiment runs). Credit-metered. _(execute)_

## visibility experiment — Dashboard screen

- `aeo visibility experiment --input-json '{"promptIds":[…],"engines":[…],"mode":"experiment","requestId":"uuid","repeat":N}'` — Run a prompt experiment or targeted check with every option the screen sends. Credit-metered. _(execute)_

## traffic ga-events — Dashboard screen

- `aeo traffic ga-events --query propertyId=…&startDate=…&endDate=…[&siteHost=…]` — The GA4 event catalog (every event name and count) the event picker shows.

## traffic ga-event — Dashboard screen

- `aeo traffic ga-event --query propertyId=…&startDate=…&endDate=…&eventName=…&scope=all|aeolo_campaign|article_context` — One GA4 event's report with the screen's scope.

## traffic ga-overview — Dashboard screen

- `aeo traffic ga-overview --query propertyId=…&startDate=…&endDate=…` — Live GA4 overview: totals, sources, top landing pages, over time.

## traffic report — Dashboard screen

- `aeo traffic report --query startDate=YYYY-MM-DD&endDate=YYYY-MM-DD[&view=ai-growth]` — The full traffic report the Traffic screen builds (GA4 + Search Console).

## traffic page-ga — Dashboard screen

- `aeo traffic page-ga --query propertyId=…&pageUrl=…&startDate=…&endDate=…` — GA4 stats for any page URL and dates.

## traffic page-gsc — Dashboard screen

- `aeo traffic page-gsc --query siteUrl=…&pageUrl=…&startDate=…&endDate=…` — Search Console stats for any page URL and dates.

## traffic history — Dashboard screen

- `aeo traffic history [--query days=30&sourceId=…]` — Managed-site traffic history for a source.

## integrations google — Dashboard screen

- `aeo integrations google channel <channelId>` — Google connection and bindings for one site (inherited or its own).
- `aeo integrations google properties-for <channelId>` — GA4 properties selectable for one site.
- `aeo integrations google sites-for <channelId>` — Search Console sites selectable for one site.

## site articles — Dashboard screen

- `aeo site articles [--query limit=…&cursor=…&q=…&counts=1]` — What the managed site actually serves, with folder counts.

## site template — Dashboard screen

- `aeo site template example` — The managed site's rendered example page (inspect view).
- `aeo site template write --input-json '{"html":"…","expected_revision":N,"note":"…"}'` — Write the whole managed-site template (site template edit patches parts). _(write)_

## pages published — Dashboard screen

- `aeo pages published` — Non-article pages Aeolo published (traffic attribution list).

## content search — Dashboard screen

- `aeo content search [--query status=…&article_type=…&search=…&date_from=…&date_to=…&deploy_target=…&group=family&sort_by=…&limit=…&offset=…]` — The Articles list with every filter, sort and page the screen uses. Any other screen parameter goes in --query 'a=1&b=2'.

## content detail — Dashboard screen

- `aeo content detail <id>` — The whole article row as the editor loads it, with its source documents.

## content performance — Dashboard screen

- `aeo content performance <id> [--days 30]` — An article's performance panel: citations, traffic, Search Console, reading.

## content jobs — Dashboard screen

- `aeo content jobs all [--query media=blog&active=false]` — Writing jobs as the screen lists them, with events and config.

## content patch — Dashboard screen

- `aeo content patch <id> --input-json '{"patches":[{"find":"…","replace":"…"}],"content_format":"markdown|html","expected_updated_at":"…","apply_to_locale_family":true,"schema_types":[…],"cta":{…},"thumbnail_crop":{…}}'` — Edit an article with every field the editor saves. _(write)_

## content begin-edit — Dashboard screen

- `aeo content begin-edit <id> --expected-updated-at ISO` — Open a review revision of an approved article without touching the live version. _(write)_

## content draft — Dashboard screen

- `aeo content draft [--title …] [--article-type …] [--language …]` — Create a blank draft (the New article button). _(write)_

## content approve-deploy — Dashboard screen

- `aeo content approve-deploy <id>` — Approve and deploy an article in one step, as the Articles screen. _(external)_

## content publish-at — Dashboard screen

- `aeo content publish-at <id> --input-json '{"channel_id":"…"|"target":"hosted"|"published_url":"…","platform":"…"}'` — Publish to a destination, or mark the article published at an external URL. _(external)_

## image — Dashboard screen

- `aeo image generate-full --input-json '{"prompt":"…","styleProductId":"…","moodReferenceJobId":"…","productUseCards":[…],"linkTarget":{"targetType":"blog_article","targetId":"…","role":"thumbnail_candidate","autoSelect":true}}'` — Generate images with every option the composer sends, linked to an article. Credit-metered. _(execute)_
- `aeo image link --input-json '{"outputId":"…","targetType":"blog_article","targetId":"…","role":"selected_thumbnail"}'` — Link a gallery image to an article in a role (thumbnail candidate, selected, related). _(write)_
- `aeo image gallery browse [--query media=visual_asset&limit=60&offset=…&createdAfter=…&createdBefore=…]` — The gallery with the screen's paging and date filters.

## publish — Dashboard screen

- `aeo publish thumbnail <contentId> --platform shopify|wordpress` — Push only the cover image to the live post (body and title untouched). _(external)_
- `aeo publish batch --platform shopify|wordpress|cafe24 --input-json '{"channel_id":"…","content_ids":[…]}'` — Deploy several articles to one destination at once. _(external)_
- `aeo publish shopify <contentId> --channel-id <id>` — Deploy to Shopify as the screen (also hosted-published articles). _(external)_

## strategy — Dashboard screen

- `aeo strategy visual full` — The whole visual style with cover settings, saved looks and the brand-snapshot preview.
- `aeo strategy visual save --input-json '{…coverTitleRequired, coverTitleRenderer, coverAspectRatio, coverLogoOverlay, coverTemplate, moodReferenceJobId, savedLooks…}'` — Save the visual style with every field the screen edits. _(write)_

## prompts add-full — Dashboard screen

- `aeo prompts add-full --input-json '{"prompts":[{"stage":"…","query_form":"…","canonical":"…","prompt":"…","topic_id":"…","regions":[…]}]}'` — Add prompts with topic and regions, as the screen. _(write)_

## prompts generate-full — Dashboard screen

- `aeo prompts generate-full --input-json '{"stages":[…],"languages":[…],"bizContext":"…","instruction":"…"}'` — Generate prompts with stages and business context, as the screen. _(execute)_

## prompts delete-many — Dashboard screen

- `aeo prompts delete-many --ids a,b,c` — Delete several prompts at once. _(write)_

## channels create — Dashboard screen

- `aeo channels create --input-json '{"url":"…","type":"…","label":"…","detected_stack":{…}}'` — Add a destination as the Sites screen (reuses an unserved row, stamps the stack, auto-verifies custom sites). _(write)_

## audit run-screen — Dashboard screen

- `aeo audit run-screen [--channel-id <id>] [--max-pages N]` — Run a site audit as the screen (no channel = the primary site). Credit-metered. _(execute)_

## audit latest — Dashboard screen

- `aeo audit latest [--channel-id <id>]` — The latest audit for one site, as the audit tab.

## domain list-full — Dashboard screen

- `aeo domain list-full` — Brands with the cap-lock flag, as the brand switcher.

## domain entitlements — Dashboard screen

- `aeo domain entitlements` — What this brand may use: engines, features, who is billed.

## domain capacity — Dashboard screen

- `aeo domain capacity` — How many brands this account may add (used, limit, remaining).

## domain create — Dashboard screen

- `aeo domain create --host <host> [--name …]` — Create a brand without starting a crawl (the Settings add-brand form). _(write)_

## domain claim — Dashboard screen

- `aeo domain claim <draftId> --input-json '{"preview":{…}}'` — Claim an onboarding draft with the edited preview. _(write)_

## account me — Dashboard screen

- `aeo account me` — The signed-in user: id, role, avatar, created.

## account subscription-full — Dashboard screen

- `aeo account subscription-full` — The whole subscription: status, period, usage, billing source.

## account ledger-full — Dashboard screen

- `aeo account ledger-full [--query start=…&end=…&limit=500&offset=…&actor=…]` — The credit ledger with totals and summary, by member and dates.

## brand products — Dashboard screen

- `aeo brand products catalog [--query fresh=1]` — The full product catalog with sync status.
- `aeo brand products bulk --pdp-urls a,b` — Add up to 50 products by PDP URL at once. _(external)_

## report snapshot-full — Dashboard screen

- `aeo report snapshot-full --input-json '{"reportType":"…","locale":"…","htmlContent":"…","accessMode":"…","expiresInDays":30,"rotate":true,"ogMeta":{…}}'` — Create or rotate a shared report link with every option the screen sends. _(external)_

## posts update — Edit a social post

- `aeo posts update <postId> --input-json '{"title":"…","body":"…","status":"…","target":"…","postType":"…"}'` — Edit a social post as its detail panel saves it. _(write)_

## agent card — Agent context card

- `aeo agent card` — The agent context card: this week's published and awaiting-approval counts, connected sources and their health, and the inbox.

## agent voice — Agent language and tone

- `aeo agent voice` — How the agent speaks to this user: language and tone.
- `aeo agent voice set --lang ko|en [--tone "…"]` — Change the agent's language and tone for this user (also re-languages existing agent rooms). _(write)_

## Not reachable on purpose

- Settings › Developer API tokens (list/create/revoke): Credential minting stays with the person: an agent that can mint tokens can hand out standing access.
- Settings › Account delete, Settings › Domains delete: Irreversible account/brand deletion stays a person's click.
- Settings › Billing checkout, credit packs, customer portal, store billing sync: Purchases and payment-method changes are never made by the agent.
- Settings › Account dashboard language: Client-only preference of the viewer's browser.
- Article professional legal approval (legal-publication-review approvals): Approval by a qualified reviewer must be that person's own act; content approve states it does not grant it.
- Gallery file upload: Needs the file bytes from the viewer's device; the agent pins an image by URL with media set-thumbnail instead.
