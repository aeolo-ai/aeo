# Image Thumbnails — Pexels reference + Nano Banana Pro product swap

Generate a 16:9 OG thumbnail by compositing a brand product onto a Pexels reference scene. The pipeline is one shot per call (~25s, ~$0.13) and persists to `content_history.thumbnail_url` so the next `aeo content deploy` ships the image to Shopify.

> **Budget**: per-user monthly cap of **$10 USD** (UTC calendar month). The `aeo image swap` response shows remaining budget after each call. Failed swaps don't charge.
>
> **Tier gate**: `image swap` requires `content-create` (Starter+). `image search`, `image upload`, and `products` are open to all members with edit access.

> **Skip the composition?** Two no-Gemini paths:
> 1. Already-public URL: `aeo content update <id> --thumbnail-url <url>` pins it directly.
> 2. Local file: `aeo image upload --file <path> --content <id>` uploads to the thumbnail bucket and pins atomically.
>
> To remove an existing thumbnail, use `--clear-thumbnail`.

> **All paths share one gate**: images must be **≤25 megapixels** (Shopify article attachment limit). Validation runs at every ingress — `image swap` output, `image upload`, and `content update --thumbnail-url`. Oversized images return `IMAGE_TOO_LARGE` (HTTP 422) with the exact dimensions in the message. Resize with `sips -Z 4800 <input> --out <output>` or similar.

---

## Workflow

1. **Pick a content item** — `/aeo content` to find the article ID you want a thumbnail for.
2. **Pick a product** — `/aeo products` to find the product ID (`og_image_url` must be set; check the "Has Image" column).
3. **Find a reference scene** — `/aeo image search <query>` to search Pexels. Copy a URL from the result table.
4. **Generate** — `/aeo image swap --content <id> --product <id> --reference <url>`.

The result is auto-persisted to `content_history.thumbnail_url`. Pass `--no-persist` if you just want to preview the URL without committing.

---

## /aeo products — List the product catalog

```bash
aeo products
```

Returns a table of `id | title | price | has_image | source | added`. The `has_image` column tells you if the product has an `og_image_url` — products without an image can't be used as a swap source.

If you don't see your product, add it from its PDP URL:

```bash
aeo product add --pdp "https://shop.example.com/products/sku-123"
```

The scraper pulls title, og:image, and price from JSON-LD or meta tags. If the scrape misses fields, fix them via the dashboard's product editor (no CLI patch yet).

---

## /aeo image search — Pexels reference scenes

```bash
aeo image search "skincare bathroom morning routine" --per-page 12 --page 1
```

Returns a table of `id | alt | photographer | url`. Pick the `url` that best matches the scene you want the product to live in.

**Query tips**:
- Be concrete about the setting and mood ("kitchen counter natural light", "gym bag overhead shot").
- Avoid product nouns in the query — the product gets composited in. Search for the *scene*, not the product.
- Landscape orientation is requested by default (OG thumbnails are 1200×630).

### Picking the reference (do not skip)

The swap prompt has a clear strong path: **"if hands hold something, replace what they're holding with the brand product, at the same position/scale/orientation."** Matching this path is the single biggest lever on output quality — references that show "human + hand + object near body" hand the model an obvious target to substitute. Fallback paths (no object in scene → place product at visual focus) routinely produce floating/awkward composites.

Apply these three filters in order when you have a list of candidates:

1. **Hand-holding-an-object scene** — pick scenes where hands are visibly holding, applying, or pouring an object onto a body part or surface. Skip "object alone on table," "person posing without product," "lifestyle shot with no application gesture."
2. **Form-factor compatible** — the held object should be roughly the SAME SHAPE CATEGORY as the brand product (stick → another stick/lipstick/marker; tube → tube/bottle; jar → jar/container). Cross-form-factor swaps (stick ↔ spray bottle, jar ↔ pipette) routinely cause the model to morph the silhouette and break product fidelity. Read the article + product to know the product form before searching.
3. **Article topic context** — the scene's body part / setting should match what the article is about. A face-skincare article needs a face-application scene; a haircare article needs hair/scalp; a body-lotion article needs limbs. Topic mismatch (article about facial sunscreen, ref shows sunscreen on legs) makes the thumbnail feel off-brand even when the swap itself is clean.

Worked example — article: "How to choose mineral or chemical sunscreen for sensitive face skin," product: sunscreen stick:

| Candidate | Form factor | Topic | Verdict |
|---|---|---|---|
| Woman with face mask, no product visible | — | face ✓ | weak (no held object → fallback path) |
| Hand applying lotion from tube to LEGS | tube ≠ stick | legs ≠ face | reject |
| Person with sunscreen and camera in desert | object on display | no application | reject |
| **Close-up of hand applying cream to cheek** | hand+object→face ✓ | face ✓ | **pick** |

If none of the first-page results match all three filters, refine the query (`--page 2`, or re-run with a more specific scene). Don't settle for a weak fit — every swap costs ~$0.13.

---

## /aeo image upload — Upload a local image to the thumbnail bucket

```bash
aeo image upload --file ./tex1.jpg --content <article_id>
```

**Flags**:
- `--file` (required) — local image path. CLI reads + base64-encodes; server validates and uploads.
- `--content` (optional) — pin the uploaded URL as the article's `thumbnail_url` in the same call. Skip to just get a URL back.
- `--mime-type` (optional) — override MIME. Auto-detected from file extension (`.jpg`, `.png`, `.webp`).

**What it does**: validates ≤25 MP → uploads to `domain-thumbnails/{domain_id}/{content_id or "standalone"}-{timestamp}.{ext}` → returns the public URL. No Gemini cost. No product needed.

**Use when**:
- You have a brand asset (own photoshoot, Drive file) and want to skip the swap pipeline.
- You already downloaded a Pexels result locally and want to pin it directly.
- Quick fix for a >25 MP thumbnail: resize, then re-upload with this command.

**Response shape**:

```
# Image Uploaded

- **URL:** https://.../domain-thumbnails/<domain_id>/<content_id>-<ts>.jpg
- **Dimensions:** 4800×3200 (15.4 MP)
- **Size:** 781 KB
- **Format:** jpeg
- **Pinned to:** <article title>   ← only if --content was passed
```

---

## /aeo image swap — Generate the thumbnail

```bash
aeo image swap \
  --content <content_id> \
  --product <product_id> \
  --reference "https://images.pexels.com/photos/.../landscape.jpg"
```

**Flags**:
- `--content` (required) — content_history ID
- `--product` (required) — products table ID (must have og_image_url)
- `--reference` (required) — reference scene URL (from `aeo image search` or a direct URL)
- `--no-persist` — return the generated URL without updating `content_history`. Useful for previewing before commit.

**What the model does** (`gemini-3-pro-image-preview`):
- Takes the reference scene as image #1, the product packshot as image #2.
- Replaces any held/displayed object in the scene with the brand product.
- Preserves the scene's framing, lighting, and color temperature.
- Preserves the product's form factor (stick stays stick, tube stays tube), label text, and packaging colors verbatim from the packshot.
- Output: a 16:9 landscape image, uploaded to the `domain-thumbnails` Supabase bucket.

**Response shape**:

```
# Thumbnail Generated — <article title>

- **Thumbnail:** https://.../domain-thumbnails/<domain_id>/<content_id>-<ts>.png
- **Product:** <product title>
- **Reference:** <pexels url>
- **Persisted:** yes — saved to content_history

**Monthly budget:** $X.XX / $10.00 spent · $Y.YY remaining (resets <iso ts>)
```

---

## Failure modes & retries

| Error message contains | What to do |
|------|------|
| `Monthly thumbnail budget reached` | Wait until next UTC month, or ask user to increase cap (server-side change). |
| `Product has no og_image_url` | Add `og_image_url` via the dashboard product editor, or pick a different product. |
| `Image generation provider quota exceeded` | Upstream Gemini quota — wait a minute and retry. |
| `Image generation provider rejected the API key` | Server-side config issue — `aeo report` it. |
| `IMAGE_FETCH_FAILED` | The reference URL is unreachable or not an image. Pick a different Pexels result. |
| `NO_IMAGE_IN_RESPONSE` | Model refused (rare for product swaps). Try a different reference scene. |

The handler only charges the budget *after* a successful swap, so failed calls are free.

---

## CUD Rule

`aeo image swap` and `aeo product add` are write operations — confirm with the user before calling. Show the chosen content title, product title, and reference URL preview, and ask "Proceed?" before running. See SKILL.md → "CUD Rule".

---

## Programmatic equivalents (for direct REST callers)

The CLI commands resolve to:

- `GET  /v1/connector/domains/:domainId/products`
- `POST /v1/connector/domains/:domainId/products` — body `{ pdpUrl }`
- `GET  /v1/connector/image/search?q=...&perPage=...&page=...`
- `POST /v1/connector/domains/:domainId/image/swap` — body `{ contentId, productId, referenceUrl, persist? }`

All return `text/markdown` on success, JSON `{ code, message }` on error.
