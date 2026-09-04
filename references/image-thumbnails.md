# Covers & Thumbnails — generate a sweep, then curate

A blog cover is produced by **generating candidates and picking one**, not by describing a picture and hoping. The brand's own visual style guide steers the look; a catalog product, when there is one, is carried in as a *subject* the model must reproduce rather than invent.

> **Cost**: every rendered candidate is 5 credits (`image generate`, `image swap` alike). `--sweep 4` is 20 credits. Failed renders are refunded. Tier gate: `content-create` (Starter+). `image search`, `image upload`, `image gallery`, and `products` are open to any member with edit access.
>
> **Size gate, every path**: images must be **≤25 megapixels** (Shopify's article-attachment limit). Enforced at `image upload`, `content update --thumbnail-url`, and on generated output. Oversized returns `IMAGE_TOO_LARGE` (HTTP 422) with the dimensions. Resize with `sips -Z 4800 <in> --out <out>`.

---

## Pick the path

| You have | Use |
|---|---|
| An article and a brand style guide | `image generate --brand-style --sweep N` |
| …and the article is about a catalog product | add `--product <id> --subject-policy ambient` |
| An article and nothing else to say about the shot | `image generate --recommend <contentId>` — the server writes the scene from the article's own visual brief |
| A finished image already (own photoshoot, Drive, downloaded) | `image upload --file <path> --content <id>` |
| A public image URL | `content update <id> --thumbnail-url <url>` |
| A clean standalone product cut **and** a scene with a hand holding a same-shaped object | `image swap` (see below — narrow, and it fails loudly when the packshot is a box) |

---

## Step 0 — read the visual style guide first

```bash
aeo strategy visual
```

**This is the only channel brand knowledge has into image generation.** When it is empty, `--brand-style` has nothing to apply and generations come out as stock photography — the warm-wood-bench, generic-wellness look that fits any brand and belongs to none. Filling it is not optional polish; it is the difference between a brand's own frame and an anonymous one.

If it is empty or thin, fill it before generating (full flag reference: [strategy.md](strategy.md)):

```bash
aeo strategy visual update \
  --description "high-key studio on brand-yellow seamless; bare skin; hands applying a clear square patch. No packaging, no box, no logo, no text, lettering or numerals." \
  --keywords "high-key,seamless backdrop,brand yellow,skin close-up" \
  --add-images "https://…/0122_shoot.jpg,https://…/spb01_lifestyle.jpg" \
  --definitions '{"https://…/0122_shoot.jpg":"IDEAL REFERENCE: lighting and backdrop colour — copy the high-key falloff, not the model"}'
```

Board rules that come from measured failures:

- **Board images should be the brand's real photographs** (`products.media` → `media->'images'`; shoot-numbered filenames like `0122_*`, `*_Lifestyle_*` are usually the real thing). Generated images fed back in make the house style self-replicate and drift.
- **No box packshots, no text-overlay graphics on the board.** Packaging and lettering leak into everything generated afterwards.
- **Write a `--definitions` note per image.** Without one the model receives an anonymous attachment and guesses what it was supposed to copy.
- **Put the ban list in `--description`**: no packaging, no box, no logo, no text/lettering/numerals. Small label text is the one thing image models reliably smear.

---

## `/aeo image generate` — the default path

```bash
aeo image generate \
  --prompt "close-up of a hand pressing a clear patch onto a forearm, high-key" \
  --brand-style \
  --product <product_id> \
  --subject-policy ambient \
  --sweep 4 \
  --aspect 16:9
```

**Conditioning axes** — these are what separate a cover from a stock photo:

| Flag | What it does |
|---|---|
| `--prompt` | The scene. Required unless `--recommend` is set. |
| `--brand-style` | Applies the brand's stored taste: style-guide description, keywords, mood board, brand colors and typography. |
| `--product <id[,id]>` | The catalog product **is the subject**. Turns on product-fidelity grammar (product photo pool, legend, edit-imperative "keep this exact product unchanged" clause) — and it does so independently of `--brand-style`, because fidelity is not a matter of taste. IDs come from `aeo products`. |
| `--subject-policy feature\|ambient` | `feature` (default): exactly one product, scene built around it. `ambient`: the product **may** appear — small, never the centerpiece, omittable — but wherever it shows it is preserved exactly. `ambient` is the cover policy; `feature` is the product-hero policy. |
| `--subject-ref url,url` | Reproduce these **exactly**. The WHAT slot for brands with no catalog row: a clinic's treatment room, a service's people. Max 4 — more conflicting angles measured *less* faithful, not more. |
| `--style-ref url,url` | Lend **look only**. Outranks the stored board for this generation. Max 8. |
| `--ref url,url` | Role-free attachments, directed by your prompt alone. Max 4. |
| `--sweep N` | 1–8 candidates. Renders are a physical dice roll (a product can come out floating); N candidates then a pick is the operating shape, not a luxury. 5 credits each. |
| `--model` | `nano-banana-pro` (default) or `gpt-image-2`. |
| `--aspect` / `--resolution` | Default `16:9`. Vertical surfaces want `4:5` **natively** — do not crop a 16:9 master, a crop is not a vertical composition. |

Reference URLs for `--subject-ref` / `--style-ref` / `--ref` come from `aeo image gallery` (past outputs and uploads) or any public URL.

### `--recommend` — let the server write the scene

```bash
aeo image generate --recommend <contentId> --sweep 4
```

The server composes the scene from that article's stored visual brief with the same builder the automatic cover sweep uses, picks the article's products, and forces brand style on. A `--prompt` alongside it **layers on** the composed scene rather than replacing it. Use this when the article is the whole brief and you have nothing to add.

### Then poll, look, and pin

```bash
aeo image poll <jobId> <jobId> …          # result URLs appear when status is completed
aeo content thumbnail <contentId> --url <imageUrl>
```

**Curate by eye, not by faith.** Download the candidate URLs and actually look at them before pinning. What to reject:

- The product is a different shape, has a different cap, a mirrored wordmark, or shows up twice.
- Any lettering the brand did not put there (props leak text — socks, labels, signage).
- A wide environment shot: thumbnails render small, so a clear gesture beats a lovely room.

Series consistency comes from **same grammar, different body part** (heel / forearm / wrist). Reuse the same prompt verbatim across articles and the list ends up showing the same photo twice.

---

## `/aeo image swap` — one product into one scene

```bash
aeo image search "kitchen counter natural light" --per-page 12
aeo image swap --content <id> --product <id> --reference "https://images.pexels.com/…"
```

Async: returns a job ID, poll with `aeo image poll`. Persists to `content_history.thumbnail_url` on success; `--no-persist` previews without committing.

> 🚨 **Swap reproduces the product's `og_image` verbatim.** That is the whole instruction the prompt gives the model, so it is only as good as that one photo. **Open the packshot before you swap.** If it is a retail box, the cover becomes a photo of a hand holding a box — with the label text smeared. Measured twice on the same brand: a packshot showing loose units on their own swapped cleanly; the box-forward SKU produced a box ad. `--product` reads `og_image_url` only and cannot be pointed at a gallery image.

When it does apply, the reference scene decides the result. Filter candidates in this order:

1. **A hand holding/applying an object** — the prompt's strong path is "replace what the hands are holding". Scenes with no held object fall back to "place the product at the visual focus" and routinely float.
2. **Same shape category** — stick → stick/marker, tube → tube, jar → jar. Cross-form-factor swaps morph the silhouette.
3. **Matching body part / setting** — a face-serum article needs a face scene, not legs.

If no first-page result clears all three, refine the query rather than settling; every attempt costs credits.

**With no clean standalone product cut, do not use swap.** `image generate --product <id> --sweep 4` with "no packaging, no box, no labels, no text" in the prompt is the answer.

---

## `/aeo products` — the catalog

```bash
aeo products                 # id | title | price | has_image | source | added
aeo product add --pdp "https://shop.example.com/products/sku-123"
aeo products discover        # crawl the sitemap for candidate PDPs
aeo products rescan          # re-scrape up to 30 PDPs to refresh media/title/image/price
```

`has_image` reports whether `og_image_url` is set. A product without one cannot be a swap source (generation with `--product` still works off the wider `media` pool).

---

## `/aeo image upload` — bring your own image

```bash
aeo image upload --file ./cover.jpg --content <article_id>
```

- `--file` (required) — local path; the CLI base64-encodes it, the server validates and stores.
- `--content` (optional) — pin it as that article's `thumbnail_url` in the same call.
- `--mime-type` (optional) — auto-detected from the extension.

No provider call, no credits, no product needed. Uploads also become style references for later generations, so a brand's real photographs are worth uploading even when they are not the cover.

To drop a thumbnail: `aeo content update <id> --clear-thumbnail`.

---

## `/aeo image gallery` — what has already been made

```bash
aeo image gallery --limit 20
aeo image gallery delete <assetId>
```

Lists the domain's generated and uploaded assets. Use it to find URLs for `--subject-ref` / `--style-ref`, and to check whether a cover already exists before paying for a new sweep.

---

## Failure modes

| Error contains | What to do |
|---|---|
| `INSUFFICIENT_CREDITS` | The sweep costs 5 credits per candidate. Lower `--sweep` or top up. |
| `Product has no og_image_url` | Add the image via the dashboard product editor, `products rescan`, or pick another product. |
| `Image generation provider quota exceeded` | Upstream KIE/Gemini quota — wait a minute and retry. |
| `Image generation provider rejected the API key` | Server-side config — `aeo diag report` it. |
| `IMAGE_FETCH_FAILED` | A reference URL is unreachable or not an image. **`http://` image URLs are rejected outright** — use the `https://` form. |
| `NO_IMAGE_IN_RESPONSE` | The model refused. Try a different reference or reword the scene. |
| `IMAGE_TOO_LARGE` (422) | Over 25 MP. Resize and re-upload. |
| Jobs sit at `processing` | KIE renders asynchronously; poll again before assuming failure. |

---

## CUD Rule

`image generate`, `image swap`, `image upload`, `strategy visual update`, `content thumbnail`, and `product add` all write. Show the article title, the chosen product, the sweep count and the credit cost, and ask "Proceed?" before running. See SKILL.md → "CUD Rule".

---

## Programmatic equivalents (direct REST callers)

- `GET  /v2/connector/domains/:domainId/products`
- `POST /v2/connector/domains/:domainId/products` — `{ pdpUrl }`
- `GET  /v2/connector/image/search?q=…&perPage=…&page=…`
- `POST /v2/connector/domains/:domainId/image/upload` — `{ base64, mimeType, contentId? }`
- `POST /v2/connector/domains/:domainId/image/swap` — `{ contentId, productId, referenceUrl, persist? }` (async → `{ jobId }`)
- `POST /v2/connector/domains/:domainId/image/generate` — `{ prompt, model?, aspectRatio?, resolution?, count?, applyBrandStyle?, referenceUrls?, subjectReferenceUrls?, styleReferenceUrls?, styleProductIds?, subjectPolicy?, recommendFromContentId? }` (async → `{ jobs, taskIds }`)
- `POST /v2/connector/domains/:domainId/video-generation/status` — `{ ids }` (mode-agnostic poll; backs `image poll`)

All return `text/markdown` on success, JSON `{ code, message }` on error.
