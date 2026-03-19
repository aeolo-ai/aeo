# GEO Strategy — Domain Knowledge

이 파일은 에이전트가 GEO 데이터를 읽고 콘텐츠 전략을 판단하기 위한 도메인 지식이다.
`/aeo` 로드 후, 데이터 해석과 콘텐츠 방향 결정 전에 반드시 참고한다.

---

## GEO란 무엇인가

SEO가 검색 결과 "순위"에 집중했다면, GEO는 ChatGPT, Claude, Perplexity, Gemini 같은 AI 엔진에 **인용되고, 추천받고, 신뢰받는 것**이 목표. 근본적으로 다른 패러다임.

| 기존 SEO | GEO |
|----------|-----|
| 페이지 순위 최적화 | 인용·추천 최적화 |
| 볼륨 우선 | 일관성 우선 |
| 한번 만들고 방치 | 30-60일 주기 지속 업데이트 |
| DA(도메인 권위) 중심 | 서드파티 멘션 빈도 + 구조화된 콘텐츠 |
| 하나의 검색엔진 | 4+ AI 엔진 동시 최적화 |

**GEO는 속도와 구조적 명확성에 보상한다. 둘 다 브랜드가 통제할 수 있는 변수다.**

---

## Aeolo 4단계 파이프라인

```
Step 1: Brand    → "이 도메인이 뭐 하는 곳이야?"
Step 2: Prompts  → "AI 엔진에 뭘 물어봐야 해?"
Step 3: Visibility → "AI가 이 브랜드를 언급하고 있어?"
Step 4: Content  → "빠진 게 뭐고, 뭘 써야 해?"
```

에이전트는 Step 4 — 갭을 콘텐츠로 메우는 실행자다. 먼저 Step 1~3 데이터를 읽고 판단한다.

---

## Visibility 데이터 읽는 법

### 갭 = "이 프롬프트에서 이 엔진이 우리 브랜드를 인용하지 않았다"

visibility report에서 확인할 것:
1. **어느 엔진**에서 빠졌나? → 엔진별 전략이 달라짐
2. **어느 stage** 프롬프트에서 빠졌나? → 콘텐츠 타입과 앵글 결정
3. **경쟁사는 나오나?** → 나오면 비교 콘텐츠 기회; 아무도 없으면 foundational 콘텐츠 기회
4. **몇 개 엔진이 동시에 빠졌나?** → 많을수록 높은 우선순위

### Stage별 의미와 우선순위

| Stage | 의미 | 쿼리 예시 | GEO 우선순위 | 이유 |
|-------|------|----------|-------------|------|
| `comparison` | 선택지 비교 — 구매 직전 | "best CRM for startups", "HubSpot vs Salesforce" | **1순위** | 브랜드명 없는 고의향 쿼리. AI 추천이 구매에 직접 영향 |
| `use-case` | 구체적 상황 기반 | "CRM for 10-person B2B team" | **2순위** | long-tail, 경쟁 낮고 conversion 높음 |
| `foundational` | 개념/필요성 탐색 | "what is CRM", "why use CRM" | **3순위** | 브랜드 인지 구축. 장기적 효과 |
| `implementation` | 구매 후/직전 검증 | "HubSpot pricing", "how to use [BrandName]" | **패스** | 쿼리에 브랜드명이 이미 포함 → GEO 대상 아님 |

### 엔진 우선순위

특별한 지시가 없으면: **ChatGPT > Gemini > Perplexity > Grok**

단, 두 가지 예외:
- 유저가 특정 엔진을 지정하면 그것을 따른다
- visibility check를 일부 엔진만 했다면 체크된 엔진 중에서만 판단한다

---

## 갭 → 콘텐츠 결정 프레임워크

> **Gap은 콘텐츠 트리거 중 하나일 뿐이다.** 고객 브리프, 특정 프롬프트 타겟, 기존 콘텐츠 리프레시(같은 주제로 새 글 뽑기) 등 다양한 경로가 있으며, 어떤 경로든 `content-create.md`의 Pre-flight → Outline 흐름으로 합류한다.

### 1단계: 갭 클러스터링

여러 갭을 하나의 글로 커버할 수 있는지 의미적으로 묶는다. 판단 기준:
- **같은 의도를 다른 방식으로 묻는가?** → 같은 클러스터
  - "best sunscreen stick for runners" + "top SPF sticks for outdoor sports" → 하나의 listicle
- **같은 지식 영역인가?** → 허브 아티클 기회
  - "what is GEO" + "how does AI citation work" + "why AI search matters" → foundational 허브
- **엔진별로 같은 프롬프트에서 빠졌나?** → 하나의 글로 멀티 엔진 커버 가능

### 2단계: articleType 결정

갭의 stage + 쿼리 패턴으로 결정:

| 갭 패턴 | 추천 articleType |
|---------|-----------------|
| comparison stage + "best X", "top X" | `ranked_list` ← **항상 첫 번째로 고려** |
| comparison stage + "X vs Y" | `comparison` |
| use-case stage + "best X for [situation]" | `ranked_list` 또는 `how_to` |
| foundational stage + "what is X" | `guide` |
| foundational stage + "how to X" | `how_to` |
| 여러 stage 동시 | `faq` (허브 역할) |

**Listicle(ranked_list)이 기본값**: AI 인용의 53%가 리스티클에서 발생. 포맷이 확실하지 않으면 ranked_list로.

### 3단계: 엔진별 전략 반영

콘텐츠를 쓸 때 갭이 있는 엔진의 특성에 맞춘다:

| 엔진 | 주요 인용 소스 | 톤 & 전략 |
|------|-------------|----------|
| **ChatGPT** | Wikipedia, Global news sites, Blogs | 실용적, conversational, how-to. 백과사전식 어조 선호. 평균 인용 페이지 **2,800 words** |
| **Gemini** | **YouTube (대부분 카테고리 1위)**, Blogs, News sites | Schema 강화, 구조화 데이터. YouTube 콘텐츠 병행 고려 |
| **Perplexity** | Blog/editorial, News, Expert reviews | 니치 전문성 + **90일 이내** 최신 콘텐츠 우선 선호 |
| **Grok** | X(Twitter), 실시간 뉴스 | 실시간 트렌드, 커뮤니티 반응 반영 |

여러 엔진에 동시 갭이 있으면 ChatGPT + Gemini 조합("실용적 + 구조화")을 기본으로.

---

## Audit 데이터 읽는 법

Audit score는 "AI 엔진이 이 사이트를 얼마나 잘 읽고 신뢰할 수 있나"의 척도.

### 항목별 의미와 액션

| Audit 항목 | 빠졌을 때 의미 | 즉각 액션 |
|-----------|-------------|---------|
| **Schema (FAQ, HowTo, Article)** | AI가 콘텐츠 타입을 구분 못함 | FAQ/HowTo JSON-LD 추가 |
| **H1 비어있음** | AI가 페이지 핵심 주제 파악 불가 | H1에 키워드 포함 텍스트 추가 |
| **TL;DR/BLUF 없음** | AI가 "특정 답변"을 인용할 수 없음 | 모든 글 상단에 2~3문장 요약 추가 |
| **datePublished/Modified 없음** | AI가 신선도 판단 불가 → citation 급감 | `<time datetime="">` + OG meta 추가 |
| **내부 링크 부족** | AI가 관련 콘텐츠 연결 못함 | hub-and-spoke 구조 + 글 간 크로스 링크 |
| **리스티클/비교 콘텐츠 없음** | AI 인용 53%를 놓치고 있음 | ranked_list + comparison 글 우선 제작 |
| **저자 바이라인 없음** | AI 신뢰도 판단 어려움 | 이름 + 직함 + 자격 추가 |

### Audit 우선순위 판단

- **HIGH (즉시)**: Schema 없음, H1 비어있음, TL;DR 없음, 날짜 구조화 안됨
- **HIGH (콘텐츠)**: 리스티클 없음, Hub-and-Spoke 없음, 내부 링크 부족
- **MED**: 저자 바이라인 없음, 외부 링크 부족

Audit 문제 = 글을 아무리 잘 써도 AI가 읽지 못하는 상태. **글 쓰기 전에 HIGH 항목부터 짚어준다.**

### 기술 크롤러 접근성 (AI 크롤러는 Googlebot과 다르다)

Schema/콘텐츠 이전에 AI가 사이트를 읽을 수 있는지 확인한다:

| 항목 | 문제 시 의미 | 확인 방법 |
|------|-----------|---------|
| **SSR(서버 사이드 렌더링)** | AI 크롤러는 JS 실행 못하는 경우 많음 → 콘텐츠 미인식 | HTML 소스에 본문 텍스트 있는지 확인 |
| **robots.txt AI 크롤러 허용** | GPTBot, anthropic-ai, PerplexityBot 등이 차단되면 invisible | robots.txt에 `Allow: /` 명시 |
| **페이지 로드 2초 이내** | AI 크롤러는 Google만큼 기다리지 않음 | Core Web Vitals 확인 |

이 3가지 중 하나라도 막혀있으면 다른 모든 최적화가 무의미하다. 기술 이슈가 의심되면 유저에게 먼저 짚어준다.

---

## 브랜드 컨텍스트 활용

`/aeo domain brand`로 로드한 데이터에서:

- **optimization_type** 확인:
  - `brand` → 브랜드명이 AI에 인용되는 것이 목표. 서드파티 listicle에서 자연스럽게 포함
  - `product` → 특정 제품이 추천/비교에 포함. 제품 비교 콘텐츠 우선
  - `campaign` → 시즌/메시지 노출. 트렌드 연계 콘텐츠

- **competitors** 활용: 비교 콘텐츠에서 경쟁사와 함께 등장시켜 자연스러움 확보
- **key_features + value_proposition**: 브랜드 멘션 시 사실 기반 설명의 소스
- **brand_context**: 시장 포지셔닝, 타겟 오디언스, 핵심 내러티브 — 글의 앵글 결정에 사용

---

## 유저에게 수집해야 할 추가 맥락

`/aeo` 데이터만으로 모자랄 수 있는 것들. 글 쓰기 전에 확인한다:

**제품/서비스 관련:**
- 최신 스펙, 가격, 출시일 (API 데이터에 없을 수 있음)
- 실제 고객 사용 후기 또는 케이스 스터디
- 경쟁사 대비 우위 포인트 (사실 기반)

**콘텐츠 전략 관련:**
- 타겟 엔진 지정이 있는가? (없으면 ChatGPT+Gemini 기본)
- 타겟 언어/시장 지정이 있는가?
- 기존에 발행한 관련 글이 있는가? (내부 링크 연결용)
- 퍼블리시 채널: 자사 블로그인가, Shopify인가, 외부 미디어인가?

**리서치 관련:**
- 외부 문서(PDF, Google Drive, 웹 리서치)를 이미 가져왔는가? → `externalResearch`로 활용
- 특정 통계나 데이터 소스를 써야 하는가?

→ 이것들이 명확하면 research step을 빠르게 진행할 수 있고 글의 신뢰도가 높아진다.

---

## 고급 GEO 개념 (판단 시 참고)

- **Semantic Neighbourhood**: AI가 브랜드를 어떤 개념·브랜드와 함께 연결짓는지. 원하는 키워드와 함께 자주 등장하면 연상 강화
- **Hub-and-Spoke**: 허브(장문 가이드) 1개 + 스포크(특화 글) N개. AI가 사이트 구조를 이해하고 관련 콘텐츠를 연결
- **Citation drift**: AI 인용은 월 40-60% 변동. 단일 스냅샷보다 추이가 중요. 30-60일 주기 콘텐츠 리프레시 권장
- **Trust Spine**: AI가 인용할 때 참조하는 5~10개 핵심 고권위 소스. 이게 쌓여야 visibility가 안정됨

### GEO 추천 전략 유형 (참고)

브랜드의 현재 상황에 따라 어떤 갭을 먼저 메울지 앵글이 달라질 수 있다:

| 전략 | 핵심 질문 | 적합한 상황 |
|------|----------|-----------|
| `product-discovery` | "What's the best X?" / "X vs Y?" | 비교·추천 쿼리에서 빠진 경우. 가장 일반적 |
| `thought-leadership` | "How do I do X?" | 전문가·권위 소스로 인용받고 싶을 때 |
| `trust-reviews` | "Should I trust X?" | 신뢰 장벽이 높거나 리뷰가 부족할 때 |
| `local-authority` | "Best X in [location]?" | 지역 기반 쿼리에서 visibility 확보 필요 시 |
| `brand-awareness` | "What is X?" | AI가 브랜드 존재 자체를 모를 때 — 다른 전략의 전제 |

`optimization_type`(brand/product/campaign)과 함께 참고해서 콘텐츠 앵글 결정에 활용한다.

### 크로스 포스팅 흐름 (참고)

글 발행 후 AI 모델의 신뢰도와 멘션 빈도를 높이는 배포 순서:

**블로그 (canonical)** → LinkedIn 아티클 → Medium (canonical tag 포함) → Substack → Reddit (관련 서브레딧, 진정성 있는 참여)

모든 채널 간 상호 링크 필수. AI는 링크 그래프를 적극 추적한다.
