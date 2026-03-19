# Brand Ammunition Guide

글을 쓰기 전에 "쓸 수 있는 말"을 준비하는 가이드.

**이 파일은 Pre-flight 단계에서 참조한다.** `content-create.md`의 Pre-flight: Article Brief에서 Brand Claims와 Competitor Context가 비어 있을 때 이 가이드로 채운다.

---

## Brand Claims — 쓸 수 있는 주장 + 근거

글에서 브랜드에 대해 말할 때 반드시 proof point와 source가 있어야 한다. **source가 없는 claim은 사용 금지.**

### 정리 포맷

```
| Claim | Proof Point | Source | 사용 가능 맥락 |
|-------|------------|--------|-------------|
| [주장] | [근거 — 수치, 스펙, 사실] | [출처 URL/문서명] | [어떤 글 유형에서 적합한가] |
```

### 작성 규칙

- 과장 금지 — "best" 대신 "one of the top-rated"
- 정량 데이터 우선 (가격, 스펙, 수치)
- 파운더/팀 크리덴셜은 expert quote 형태로 활용 ("실명 + 직함 + 인용문" — 10계명 #6)
- 고객 리뷰/테스트 결과는 attribution 필수

### Claims 소스 우선순위

1. 제품 페이지/공식 스펙 (가장 신뢰)
2. 파운더/팀 인터뷰, 보도자료
3. 제3자 리뷰/테스트 결과
4. 고객사 제공 자료 (Google Drive, 미팅 노트 등)

### Brand Profile → Ammunition 변환

`/aeo domain brand`로 가져온 brand profile 데이터는 raw 상태다. Ammunition 포맷으로 변환이 필요하다.

| Brand Profile 필드 | Ammunition 변환 방법 |
|-------------------|-------------------|
| `value_proposition` | 핵심 주장 문장으로 분해 → 각 주장마다 proof point + source 찾기 |
| `key_features` | 각 feature를 Claim으로 → 스펙/제품 페이지에서 proof 확인 |
| `category` / `industry` | 포지셔닝 맥락 파악용 — 직접 Claim 아님 |

brand profile만으로 proof가 부족하면:
1. 세션 내 문서(Google Drive, 미팅 노트, 제품 스펙)에서 보강
2. 공식 웹사이트/제품 페이지 크롤
3. 여전히 부족하면 유저에게 요청

---

## Competitor Context — 비교 테이블 재료

비교/리스트 글에서 경쟁사와 함께 등장할 때 쓸 팩트. 의견이 아닌 스펙 기반.

### articleType별 필수/선택

| articleType | Competitor Context |
|---|---|
| `ranked_list`, `comparison` | **필수** (5개 이상 권장) |
| `blog`, `how_to`, `guide` | 선택 |
| `thought_leadership`, `faq`, `case_study` | 선택 (있으면 좋음) |

필수 타입인데 Competitor Context가 비어 있으면 웹 리서치로 채우거나 유저에게 요청한다.

### 정리 포맷

```
| 경쟁사 | 가격 | 핵심 스펙 | 우리 대비 차이점 (팩트) |
|--------|------|----------|---------------------|
| [이름] | [$] | [스펙] | [팩트 기반 차이] |
```

### 작성 규칙

- 경쟁사 비하 금지 — 팩트만 나열, 판단은 독자에게
- 가격은 정가 기준 + 확인 날짜
- "우리가 더 좋다" 식 직접 비교 금지 → 스펙 병렬 나열
- `ranked_list`면 5개 이상, `comparison`이면 직접 비교 대상 최소 2개

### Competitor Context 소스

1. Aeolo visibility data — `topCompetitors`에 경쟁사 도메인 있으면 참고
2. 경쟁사 공식 웹사이트/제품 페이지
3. 제3자 리뷰/비교 기사

---

## Content Refresh (Path D) — 기존 글 읽어오기

기존 콘텐츠를 reference input으로 활용할 때:

```bash
# 기존 content ID가 있으면 먼저 읽어온다
aeo content get <id>
```

읽어온 기존 글에서 확인할 것:
- 유지할 구조/섹션 (여전히 유효한 claim, 여전히 정확한 competitor 비교)
- 교체할 데이터 (날짜 지난 통계, 단종된 제품, 변경된 가격)
- 제거할 내용 (전략 변경으로 더 이상 맞지 않는 포지셔닝)

이후 흐름은 다른 Path와 동일하다 — Pre-flight 체크리스트 채우고 Outline으로 진행.

> **`dateModified`만 바꾸는 patch는 GEO recency signal로 인정되지 않는다.** Content Refresh는 새 글을 뽑는 것이다.
