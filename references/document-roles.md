# Writing document roles

Generated from packages/types/src/writing-documents.ts (writing-documents-v1). Do not edit this projection.

## Brand Notes / 브랜드 참고용 메모

What is recorded about this brand, and what needs correcting?

Scoped brand/product/service facts and corrections, with source and confirmation state. A saved note is not automatically verified. Preserve uncertainty, negation, market/SKU/location limits and supersession history.

Keep elsewhere: Campaign priorities, voice instructions, file inventories and task logs belong elsewhere. Structured identity and catalog remain their own owners.

Store: `domains.brand_context`. Existing command: `aeo domain brand update --brand-context`.

## Content Strategy / 콘텐츠 방향

Whose question should we address first, within which approved scope?

Reader decisions, topic priorities, approved campaign/product/market scope and intended next action. Strategy selects work; it is not evidence for a claim and cannot force unsupported brand inclusion.

Keep elsewhere: Product corrections belong in Brand Notes; source locations in Source Guide; tone in voice settings. Preserve approval/hold restrictions until explicitly resolved.

Store: `content_strategy.manifest`. Existing command: `aeo strategy update --manifest`.

## Source Guide / 자료 안내

Where can the relevant evidence be found, and what does it cover?

Evidence locations, document contents, applicable product/market/period and access or public-use limits. A pointer or connected folder is not proof of access, reading, verification or tool permission.

Keep elsewhere: Verified or customer-provided product corrections belong in Brand Notes. Global no-fabrication rules belong in the system; link display in reference_policy; completed work in operational records.

Store: `domains.data_sources`. Existing command: `aeo config data-sources update --data-sources`.

## Editing protocol

Keep only content that changes what is written, selected or supported. Cover the necessary scope; do not pad or cut to meet an arbitrary length. If an existing passage already covers the point, refine it instead of adding a duplicate. Before consolidation, retain the exact original in a recoverable revision; never silently discard provenance, uncertainty, negation, applicability or unresolved approvals. Operational history may leave the active writing input only after it has been preserved in its proper owner. Before editing, read all three current documents and retain an exact snapshot. Classify each passage by its intended use, not by keywords. Separate voice, citation display, system rules and operational records; do not force them into these three documents. Mixed passages, conflicting scopes, uncertain targets and historical approvals need review, not deletion. Show source text, destination, reason and exact before/after. Preserve wording, dates, source links, uncertainty and approval/hold conditions. A classification is not factual verification or authorization to save. Re-read before saving; if any source changed, rebuild the proposal. Use existing commands within the user's authorization and read back every save. If a save returns DOCUMENT_REVIEW_REQUIRED, show the warning to the user; retry with --force and the returned --review-token only when the user explicitly acknowledges that warning. Force never bypasses access, limits, expiry or stale inputs. Writing agents consume these inputs; they do not reorganize or persist them while composing an article.

For internal aops work, record customer operations in aops; use authorized customer aeo commands for these brand documents. This reference adds no aops command or permission.
