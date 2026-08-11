# Market map — the upstream source for prompts, topics, and writing direction

Built at onboarding time from two inputs, in this order:

1. **Qualification (brand snapshot)** — what the brand sells (`services.items`),
   what it is NOT (`identity.exclusions`), its buying situations (CEPs) and
   competitors. This GENERATES the question candidates (qualification × situation).
2. **Search terrain (DataForSEO)** — this only JUDGES: category-level demand
   (landing), commercial intent (CPC), and who holds the surface. 99.7% of the
   search-volume corpus is noun phrases; the questions a brand actually wins
   often have no row in it at all.

## Commands

```
aeo market-map [--market KR|US|JP|TW|HK|CN|GB|ES|MX]   # show (latest map when --market omitted)
aeo market-map run [--market US]                        # build/refresh — background job, ~2-3 min
aeo market-map poll <jobId>                             # poll the build job
aeo market-map populate [--topics "<name>,<name>"]      # rail → Topics + tracked prompts
```

`run` refuses (with the reason) only when `services.items` is empty — the map
cannot know what the brand sells. On one measured clinic this gate alone was
the difference between 6 and 12 usable proposals.

`identity.exclusions` is NOT gated: an old snapshot without it builds with an
empty qualification boundary (questions cannot be flagged out-of-qualification
until the snapshot is refreshed — worth doing: on one measured brand 23% of
proposals landed outside qualification without the boundary).

## Reading the map

- **시장 구조** — deduped journey-stage split (initial/browsing/review). Initial ≫
  browsing = open market (discovery content can work); a thick review layer =
  purchase-adjacent questions exist (price/effect/side-effects).
- **홈그라운드 어휘** — terrain terms the brand's own vocabulary reaches. Empty
  means the map missed the brand: distrust the rest and check the snapshot.
- **표면 점유** — who ranks on the landing terms, BY TYPE (platform / retail /
  editorial / medical / other). The type decides the campaign shape. This is
  Google-side data: treat it as a prediction the first visibility check scores.
- **토픽 레일** — populate units, priority-ordered by homeground overlap →
  CPC → demand (never demand alone). `⏸ 보류` topics contain self-named
  questions: they answer themselves and cannot move, so never track them.
- **fingerprint** — same fingerprint on a rerun = same market. A changed
  fingerprint is a finding (the market moved), not a bug.

## Populate flow

`aeo market-map populate` takes topic **names**, not question text — the server
re-reads the stored map and decides which questions are trackable, so `⏸ 보류`
topics and `violates` questions cannot be smuggled in. Omit `--topics` to take
every non-held topic (the dashboard's default).

It creates the **Topic first** and hangs the prompts under it. Topics are
**shared across markets** (`topics.market` stays NULL): the market axis already
lives on each prompt (`brand_prompts.region`), and a Topic per market would
multiply one business theme into five — measured 2026-08-11, 40 of 164 live
Topics already carried prompts from several markets.

Tracked prompts are what the plan sells, so a populate that would exceed the
pool takes the highest-priority ones and reports what it left out.

Then run a visibility check — the first check both scores the map's surface
prediction and starts filling the mindshare view the map's topics become the
axes of.

## Judgment fields on each prompt candidate

| Field | Meaning |
|---|---|
| `term / vol / cpc` | Most specific terrain landing: category demand + commercial intent |
| `★ (mine)` | Homeground — the brand's words overlap the market's here |
| `착지: 상황 고유` | Landed only on a generic head term — the question lives in a situation volume data cannot see (often the best instruments) |
| `violates` | Exclusion token in the question — not qualified, never populate |
