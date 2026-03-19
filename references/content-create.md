# Content Creation

Default draft directory: `.aeo/` in the working directory. Add to `.gitignore`.
File naming: `{AEOLO_DOMAIN_ID}_{slug}.md`. See Step 5 for path resolution.

---

## /aeo content write — Write a GEO-optimized article

Agent writes the article directly using the guidelines below. No Mastra workflow API call — the agent is the writer. After writing, save the draft and import via `/aeo content import`.

---

### Pre-flight: Article Brief

**outline으로 넘어가기 전 필수 관문.** 어떤 경로로 들어왔든 이 체크리스트가 채워져야 한다.

#### 트리거 경로 (에이전트가 자동 판단)

| Path | 트리거 | 이미 있는 것 | 추가로 필요한 것 |
|------|--------|------------|---------------|
| **Gap-driven** | `/aeo` visibility gap 또는 `/aeo content propose` 결과 | 브랜드 데이터, 타겟 프롬프트, gap context | Competitor facts |
| **Client-brief** | 유저가 고객 자료/미팅 노트/어필 포인트를 전달 | 앵글, 자료 (날 것 상태) | Claims 추출, competitor context |
| **Prompt-targeted** | "이 프롬프트에서 우리가 나왔으면 좋겠어" | 타겟 프롬프트 | Brand claims, competitors |
| **Content Refresh** | 기존 글의 recency 만료 (91일+) | 기존 글 구조 + 콘텐츠 | 최신 데이터, 업데이트된 competitor context |

Content Refresh는 "patch"가 아닌 **새 글 뽑기**다. `/aeo content get <id>`로 기존 글을 읽어 reference input으로 활용하되, 이후 플로우는 다른 Path와 동일하다. `dateModified`만 바꾸는 patch는 GEO recency signal로 인정되지 않는다.

#### Article Brief Checklist

```
## Article Brief

### Trigger
- 출발점: [gap / client brief / prompt target / content refresh]
- 관련 gap ID, 프롬프트, 또는 기존 content ID: [있으면 기입]

### Brand Ammunition
- [ ] Brand Claims (최소 3개): 주장 + proof point + source
- [ ] Competitor Context: 비교 대상 팩트 (articleType에 따라 필수/선택)

### 1st-Party Experience (진정성 소재)
- [ ] 자체 테스트/사용 경험: 누가, 몇 회, 어떤 조건에서? (구체적으로)
- [ ] 고객 후기/케이스: 실제 사용자 피드백이 있는가?
- [ ] 도메인 특화 맥락: 이 주제만의 고유한 팩트가 있는가? (일반론 금지)
- [ ] 포지셔닝: 브랜드가 쓴 글임을 솔직하게 밝힐 것인가?

> 위 항목이 비어 있으면 유저에게 요청한다. "써본 적 없는 사람이 쓴 글" 느낌을 피하려면 최소 1개 이상의 1st-party 소재가 필요하다.

### Article Direction
- [ ] Topic / H1 candidate
- [ ] articleType
- [ ] Target keywords
- [ ] Target engine(s)
```

Brand Ammunition이 비어 있으면 **[brand-ammunition.md](brand-ammunition.md)** 참고:
- Brand Claims → `/aeo domain brand` + 세션 내 문서에서 추출 후 Ammunition 포맷으로 변환
- Competitor Context → 웹 리서치 또는 유저에게 요청

> **Authority Sources는 Pre-flight 대상이 아니다.** 글 신뢰도를 위한 외부 제3자 데이터(통계, 연구 등)는 Step 3 (Research)에서 수집한다.

---

### Step 1 — Collect inputs

Ask the user (or infer from context):
- **Topic** (required)
- **Article type** — choose from format matrix below; `blog` is default
- **Target keywords** (required, 1–20)
- **Language** — `en` (default) | `ko` | `ja` | `zh` | `ar`
- **Word count** — default 1500

If brand profile isn't loaded yet, fetch it first (`/aeo domain brand`) — brand context shapes the entire article.

### ⚠️ Step 1.5 — Brand Tone & Voice (MANDATORY GATE)

**글을 한 글자도 쓰기 전에 반드시 완료해야 한다.**

Brand profile을 로드했으면 `brand_context`에서 **Tone & Voice 섹션을 찾는다.**

#### Tone & Voice 섹션이 있는 경우

아래 형식으로 추출하고 이후 모든 단계(Outline 헤딩 네이밍 → Writing 어조 → FAQ 문체)에 일관되게 적용한다:

```
## Extracted Tone Profile
- Formality: [formal / balanced / casual]
- Voice characteristics: [추출된 형용사들]
- Phrases to use: [예시 문장/표현]
- Phrases to avoid: [금지 표현]
- Other notes: [기타 스타일 지시]
```

추출 후 유저에게 한 줄로 요약 확인: _"이 글은 [특성] 톤으로 작성됩니다. 맞나요?"_

#### Tone & Voice 섹션이 없는 경우

**작업을 멈추고 유저에게 직접 물어본다. 절대 임의로 톤을 정하지 않는다.**

```
Brand context에 Tone & Voice 정보가 없습니다.
글의 어조를 정하기 위해 몇 가지 여쭤볼게요:

1. 전체적인 톤 — formal(전문적/격식) / casual(친근/대화체) / authoritative(권위있는) 중 어느 쪽에 가깝나요?
2. 피해야 할 표현이나 어조가 있나요? (예: 과장된 마케팅 언어, 지나친 기술 용어 등)
3. 브랜드를 잘 표현한다고 생각하는 문장이나 카피 예시가 있으면 공유해 주세요.
```

답변을 받으면 → `/aeo brand update`로 `brand_context`의 Tone & Voice 섹션에 즉시 저장한 뒤 진행한다.
이렇게 해야 다음 글부터는 같은 질문을 반복하지 않는다.

### Step 2 — Outline

Build a structured outline:
- H1: question-based title matching actual AI search queries (see 10 commandments below)
- H2/H3: hierarchy where each section is independently quotable
- Note which sections will include comparison tables, expert quotes, or authority citations
- Identify 3–5 FAQ questions to cover at the end

Show the outline to the user and confirm before writing.

### Step 3 — Research

Gather supporting material:
- If you have documents/PDFs/URLs already read in this session → use that context directly
- Otherwise, search for: recent statistics, authority sources (.edu/.gov), expert quotes with name + title
- Collect `{ name, url, description }` for each source — these become inline citations

### Step 4 — Write

Write the full article following the **GEO Writing Instructions** below. Key rules:
- **Step 1.5에서 추출한 Tone Profile을 전 섹션에 일관되게 적용한다** — 헤딩 네이밍, 문장 길이, 어조 모두 포함
- **Title ≠ H1 (HARD RULE)** — 두 가지 제목을 반드시 분리한다:
  - **Title** (`/aeo content import`의 title 파라미터): SEO 최적화 긴 버전 (question-based, 키워드 포함)
    예: `"How to Reapply Sunscreen During a Long Run Without Stopping — A Stick Sunscreen Guide for Runners"`
  - **H1** (본문 첫 `#`): 짧고 punchy한 독자용 헤딩
    예: `"Reapply Sunscreen Mid-Run: The Runner's Guide"`
  - 본문은 반드시 `# {H1}` 으로 시작하고, title 문구를 본문 어디에서도 반복하지 않는다
  - H1 뒤 첫 문단은 바로 BLUF 답변으로 직행한다
  - ❌ Title과 H1이 같은 문구
  - ❌ H1 문구를 인트로에서 다시 반복
  - ✅ `# Do You Need Sunscreen for Padel?` → `Glass walls on a padel court block wind, not UV. Here's what dermatologists say...`
- BLUF in first 2–3 sentences
- Inline citations as `[Source Name](URL)` throughout
- Brand mentions at 15–25% density, always as part of a list (never solo promo)
- FAQ section at the end (3–5 questions)

**Output format:**
- **본문**: 순수 markdown만 (H1부터 FAQ까지)
- **메타데이터**: 본문과 분리하여 별도 표시

### Step 4.5 — Semantic Authenticity Check (MANDATORY GATE)

**글을 저장하기 전에 반드시 수행한다.** [content-review.md](content-review.md)의 "Semantic Authenticity" 체크리스트를 기준으로 자가 점검:

| 체크 항목 | 통과 기준 |
|-----------|----------|
| 포지셔닝 정직성 | 브랜드가 쓴 글이면서 독립 리뷰를 가장하지 않는가? |
| 자사 제품 편향 | 자사 단점도 경쟁사 수준으로 솔직하게 다뤘는가? |
| 독립 전문가 목소리 | 창업자/내부 인용 외 독립 전문가가 최소 1명 있는가? |
| 경험 구체성 | 테스트/경험 주장에 누가·몇 회·어떤 조건이 명시되었는가? |
| 도메인 특화 맥락 | 이 주제만의 고유 팩트가 녹아있는가? (일반론에 그치지 않는가?) |
| 1st-party 데이터 | Pre-flight에서 수집한 1st-party 소재가 실제로 글에 반영되었는가? |

**⚠️ 2개 이상 실패 시** 유저에게 결과를 보여주고 수정 여부를 확인한다. 1개 실패는 경고만 표시하고 진행 가능.

### Step 5 — Save draft

Determine the slug from the title (kebab-case, max 60 chars).

**저장 경로 결정 방법** — 환경마다 `pwd`가 다를 수 있으므로(Cowork VM, 로컬 등) 아래 순서로 결정한다:

1. 유저가 저장 경로를 명시했으면 그대로 사용
2. 현재 세션에서 이미 사용한 경로가 있으면 동일하게 사용
3. 없으면 유저에게 확인: "초안을 어디에 저장할까요? (예: `~/Documents/aeo-drafts/`, 현재 작업 폴더 등)"

파일명: `{AEOLO_DOMAIN_ID}_{slug}.md`

Tell the user the exact file path and that they can review/edit before importing.
Then suggest running `/aeo content import <path>`.

---

## /aeo content import [path] — Push draft to content history

Push a completed draft to the Aeolo content history for review and publishing.

### When to use

- After writing an article (via `/aeo content write` or manually) — import to Aeolo pipeline
- Any externally-written draft you want to register

### Flow

1. Read the draft file at `[path]` (default: the path used in the current session's content write step)
2. Extract or confirm: `title`, `targetKeywords`, `articleType`, `language`
3. Show the user a summary and confirm
4. Call the Connector API directly (this is an agent-only action, no CLI command)
5. On success: "Imported → View in Aeolo dashboard → Content Queue"

**API endpoint:** `POST /v1/connector/domains/:domainId/articles`

Payload fields:

| Field | Type | Required | Default |
|-------|------|----------|---------|
| `title` | string (max 500) | ✅ | — |
| `content` | string (min 100, markdown) | ✅ | — |
| `articleType` | enum | — | `blog` |
| `targetKeywords` | string[] (1–20) | ✅ | — |
| `language` | enum | — | `en` |
| `rationale` | string (max 1000) | ✅ | — |
| `metaDescription` | string (max 320) | — | — |
| `sources` | `{name, url, description?}`[] | — | — |
| `visibilityGapKeywords` | string[] | — | — |
| `schemaTypes` | string[] | — | auto from articleType |

Import flow에서 글 상단의 `<!-- schema: ... -->` 주석이 있으면 파싱해서 `schemaTypes` 필드에 매핑한다. 없으면 articleType 기반 자동 매핑.

Response: `{ "success": true, "data": { "queueItemId": "uuid", "status": "review" } }`

---

## GEO Writing Instructions — 에이전트 글쓰기 가이드라인

### 포맷 선택 매트릭스

확신이 없으면 **ranked_list를 기본으로** — AI 인용의 53%가 리스티클에서 발생한다.

| 프롬프트 패턴 | articleType | AI 인용 비율 | 분량 가이드 |
|---|---|---|---|
| "best X", "top X", "X recommendations" | `ranked_list` | 32% | 2,000~4,000 words |
| "X vs Y", "compare X and Y" | `comparison` | 18% | 1,200~2,500 words |
| "how to X", "step by step X" | `how_to` | 15% | 1,500~3,000 words |
| "what is X", "why X", "X explained" | `guide` | — | 800~1,500 words |
| 복합 질문, 여러 질문 동시 | `faq` | 11% | 1,500~2,500 words |
| 업계 트렌드, 전문가 시각 | `thought_leadership` | — | 1,500~3,000 words |
| 고객 사례, 도입 결과 | `case_study` | — | 1,200~2,000 words |

### GEO 글쓰기 10계명

1. **BLUF (Bottom Line Up Front)** — 첫 2~3문장에 핵심 답변 배치. AI는 전체 글이 아닌 "특정 답변"을 인용한다. 인트로에서 빙빙 돌지 않는다.
2. **Title ≠ H1** — Title은 SEO용 긴 버전 (question-based, keyword-rich). 본문 `#`은 짧고 punchy한 독자용 헤딩. 둘은 반드시 다른 문구. 예: Title `"What's the Best SPF Stick for Outdoor Sports in 2026?"` → H1 `"Best SPF Sticks for Outdoor Sports"`
3. **Logical H2/H3 hierarchy** — 시맨틱 HTML5 구조. 각 섹션이 독립적으로 인용 가능해야 한다. H2 하나만 떼어내도 의미가 통해야 함.
4. **Comparison tables** — 비교 데이터는 반드시 markdown/HTML table. AI가 구조화된 데이터를 비구조화 텍스트보다 선호한다.
5. **Authority signals** — 외부 권위 소스 인용 필수. 통계, 연구, .edu/.gov 소스. 인라인 인용 포맷: `[Source Name](URL)`. 출처 없는 주장 금지.
6. **Expert quotes (attributed)** — "실명 + 직함 + 인용문" 포맷. AI가 신뢰도 판단에 사용. 가공의 인용 금지.
7. **FAQ section** — 글 하단에 3~5개 FAQ. 본문에서 다루지 못한 관련 질문 커버.
8. **Schema markup 힌트** — 글 최상단(H1 앞)에 HTML 주석으로 schema 타입 명시. 포맷: `<!-- schema: Type1, Type2 -->`

   | articleType | schema 힌트 |
   |---|---|
   | ranked_list | `<!-- schema: Article, ItemList -->` |
   | comparison | `<!-- schema: Article, ItemList -->` |
   | how_to | `<!-- schema: HowTo, Article -->` |
   | faq | `<!-- schema: FAQPage, Article -->` |
   | guide | `<!-- schema: Article -->` |
   | blog | `<!-- schema: Article, BlogPosting -->` |
   | thought_leadership | `<!-- schema: Article, BlogPosting -->` |
   | case_study | `<!-- schema: Article -->` |
9. **Freshness signals** — 글 상단에 `datePublished`, `dateModified` 명시. 30일 이내 콘텐츠 인용률 100% → 1년 후 18%로 급감.
10. **Internal + external links** — 자사 콘텐츠 상호 링크 + 외부 권위 소스 링크. AI는 링크 그래프를 적극 추적한다.

### 플랫폼별 톤 가이드

지정 없으면 ChatGPT + Gemini 양쪽에 먹히는 "실용적 + 구조화" 조합을 기본으로.

| 엔진 | 주요 인용 소스 | 톤 & 구조 |
|------|-------------|----------|
| **ChatGPT** | Wikipedia, Global news, Blogs | 실용적, conversational, how-to 중심. 백과사전식 어조. 평균 **2,800 words** 기준. |
| **Claude** | 학술 콘텐츠 | 5,000자+ 장문. 기술 문서 톤. 학술 인용, 구조적 논증. 깊이 우선. |
| **Perplexity** | Blog/editorial, News, Expert reviews | 니치 전문성 + **90일 이내** 최신 콘텐츠. 팩트 밀도 높게. |
| **Gemini** | **YouTube (카테고리 1위)**, Blogs, News | Schema 강화. 구조화 데이터. YouTube 콘텐츠 병행 고려. |

### 콘텐츠 신선도 규칙

| 글 나이 | 잔존 인용률 | 에이전트 행동 |
|---------|-----------|-------------|
| 0~30일 | 100% | 새 글 발행 최적기 |
| 31~90일 | 73% | 데이터/통계 업데이트 권장 |
| 91~180일 | 51% | 리프레시 필요 — 새 데이터 + 구조 개선 |
| 181~365일 | 34% | 대폭 리라이트 |
| 1년+ | 18% | 새 글 작성 권장 |

기존 콘텐츠 수정 시 반드시 `dateModified` 갱신. 가능하면 최신 통계/데이터로 교체.

### 브랜드 멘션 원칙

- 브랜드는 리스트의 **일부**로 등장 — 단독 홍보글 절대 금지
- 경쟁사도 함께 언급하여 자연스러움 유지
- 사실 기반 정보만 (스펙, 가격, 리뷰 요약)
- 멘션 밀도: 글 전체의 **15~25%** — 과도하면 AI 신뢰도 하락
- 브랜드 VP와 key features를 brand profile (`/aeo domain brand`)에서 가져와서 사용

### 메타데이터 생성

글 작성 완료 시 아래 메타데이터를 함께 생성 (import payload에 사용):

- `title` — SEO 최적화 긴 버전 (question-based, keyword-rich). **H1과 반드시 다르게** (Step 4 HARD RULE 참조)
- `metaDescription` — 150자 이내 (BLUF 기반 요약)
- `targetKeywords` — 1~20개
- `articleType` — 포맷 매트릭스 기반
- `estimatedRefreshDate` — 발행일 + 60일

**⚠️ 중요: 메타데이터는 본문에 포함하지 않는다.**
- 초안 저장 시: 순수 markdown 본문만 파일에 저장
- 메타데이터는 에이전트가 별도로 추적하여 import 시 API payload에만 사용
- 유저에게 보여줄 때는 본문과 분리된 형태로 제시 (예: "이 글을 작성했습니다 + 별도로 메타데이터 표시")
