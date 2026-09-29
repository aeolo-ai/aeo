# Atlas continent classification (staff preview)

Classify official product/service evidence into a versioned market taxonomy. Do not generate market names. A continent is the market scope proposed from actual offerings for the current research; it is not a full-industry TAM measurement or an automatically approved scope.

```bash
aeo atlas taxonomy --domain <domain-id>
# Shopify: use structured feed fields; PDP URLs are evidence references only.
aeo atlas sources --feed-url 'https://brand.example/products.json?limit=250' --urls 'https://brand.example/products/product-a' --domain <domain-id>
# No product feed (e.g. a clinic): fetch official service-page content only when needed.
aeo atlas sources --urls 'https://brand.example/service-b' --domain <domain-id>
aeo atlas continents --input-json '<frozen JSON from sources>' --domain <domain-id>
```

These commands require staff access and membership/access to the selected domain. The existing CLI generic connector proxy supports them. Availability requires the matching API deployment; a local preview is not a deployed release.

- `taxonomy` returns fixed IDs, display names, definitions, version and hash. This pilot vocabulary is curated and incomplete.
- `sources --feed-url` reads Shopify product ID, title, type, vendor and description directly from JSON, selecting the requested PDP handles. It does not fetch PDP HTML. Variants share one product evidence entry; this does not deduplicate different product IDs into families. The feed URL is one bounded page; absent products remain explicit failures. Use existing catalog data directly in `continents` when already available. Only sites without a feed need `sources --urls` HTML extraction. Both modes enforce the exact official host and SSRF guards. It returns `{brand:{name,host},sources:[{id,url,title,text,observedAt,kind?,retrievedFrom?}],coverage}` plus collection metadata/failures. No crawl or full-catalog coverage is implied. Inspect the returned text before classification.
- `continents` accepts frozen evidence (at most 25 KB), rechecks the domain host, and makes one Jev Choice call without caching, retries or a generation fallback. The server funds the provider call during the staff pilot; no customer credits or customer records are changed.
- The response includes `continent` (fixed ID/name or null), `status` (`proposed`/`needs_review`), per-source assignments, probabilities/margins, evidence URLs, input/taxonomy/policy hashes and limitations. Raw provider cost/usage is not exposed.
- Low probability, low margin, an unsupported primary category or no matching taxonomy entry stays unclassified. Pilot thresholds are not calibrated accuracy. Do not turn a repeated answer into a correctness claim.
- Source text is operator-supplied evidence; classification does not recrawl or certify it. Missing sources and collection failures must remain visible. A proposed classification never confirms the user's research scope.

For repeatability evaluation, freeze the taxonomy and exact input before execution. Save every output/error from repeated uncached calls. Compare category IDs separately from source-grounded correctness; report both.
