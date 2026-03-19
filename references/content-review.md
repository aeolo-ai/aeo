# Content Review

GEO 도메인 전문성을 기반으로 기존 콘텐츠를 리뷰한다.

---

## /aeo content review <id> — GEO 관점 콘텐츠 리뷰

### Flow

1. **컨텍스트 로드** — 아래 3개를 병렬로 가져온다:
   ```bash
   aeo content get <id> > /tmp/aeo_review_article.md &
   aeo domain brand > /tmp/aeo_review_brand.md &   # /aeo domain brand
   aeo domain audit > /tmp/aeo_review_audit.md &    # /aeo domain audit
   wait
   ```

2. **리뷰 수행** — 아래 체크리스트를 기준으로 글을 평가한다.

3. **리포트 출력** — 아래 포맷으로 결과를 보여준다.

4. **다음 행동 제안** — 수정이 필요하면 `/aeo content update <id>` 패치 제안.

---

### Review Checklist

#### 1. Structure & Quotability (구조)

| 항목 | 기준 | 참고 |
|------|------|------|
| **BLUF** | 첫 2~3문장에 핵심 답변이 있는가? | 10계명 #1 |
| **H1** | Question-based 타이틀인가? 실제 AI 쿼리와 매칭되는가? | 10계명 #2 |
| **H2/H3 hierarchy** | 각 섹션이 독립적으로 인용 가능한가? | 10계명 #3 |
| **Comparison tables** | 비교 데이터가 있다면 테이블로 구조화되었는가? | 10계명 #4 |
| **FAQ section** | 3~5개 FAQ가 글 하단에 있는가? | 10계명 #7 |
| **Schema hints** | 권장 schema 타입이 명시되어 있는가? | 10계명 #8 |

#### 2. Trust & Authority (신뢰)

| 항목 | 기준 | 참고 |
|------|------|------|
| **Inline citations** | `[Source Name](URL)` 형식 인라인 인용이 충분한가? (섹션당 1~2개) | 10계명 #5 |
| **Expert quotes** | 실명 + 직함 + 인용문 포맷인가? 가공 인용은 없는가? | 10계명 #6 |
| **Authority sources** | .edu/.gov/연구/통계 등 고권위 소스가 포함되어 있는가? | 10계명 #5 |
| **Internal + external links** | 자사 콘텐츠 링크 + 외부 권위 소스 링크 둘 다 있는가? | 10계명 #10 |

#### 3. Freshness (신선도)

| 항목 | 기준 | 참고 |
|------|------|------|
| **datePublished / dateModified** | 명시되어 있는가? | 10계명 #9 |
| **데이터 최신성** | 인용된 통계/데이터가 1년 이내인가? | 신선도 규칙 |
| **글 나이** | 발행일 기준 잔존 인용률 구간은? (0~30일 100% → 1년+ 18%) | 신선도 규칙 |

#### 4. Brand Integration (브랜드)

| 항목 | 기준 | 참고 |
|------|------|------|
| **멘션 밀도** | 글 전체의 15~25% 이내인가? | 브랜드 멘션 원칙 |
| **리스트 내 등장** | 브랜드가 단독 홍보가 아닌 리스트의 일부로 등장하는가? | 브랜드 멘션 원칙 |
| **경쟁사 동시 언급** | 자연스러움을 위해 경쟁사도 함께 언급되는가? | 브랜드 멘션 원칙 |
| **사실 기반** | 스펙, 가격, 리뷰 요약 등 검증 가능한 정보만 사용하는가? | 브랜드 멘션 원칙 |
| **톤 일치** | brand_context의 Tone & Voice와 일관되는가? | content-create Step 1.5 |

#### 5. Semantic Authenticity (진정성)

"AI가 쓴 글" 또는 "브랜드 광고"처럼 읽히는 시맨틱 문제를 잡는다. 구조/출처가 완벽해도 이 카테고리에서 실패하면 AI 엔진과 독자 모두 신뢰하지 않는다.

| 항목 | 기준 | Red Flag 예시 |
|------|------|--------------|
| **포지셔닝 정직성** | 글의 정체가 솔직한가? 브랜드가 쓴 글이면서 독립 리뷰를 가장하지 않는가? | 자사 제품을 #1으로 뽑으면서 "Editorial Team"이 저자 → 광고 위장 |
| **자사 제품 편향** | 자사 단점을 경쟁사 단점만큼 솔직하게 다뤘는가? 경쟁사만 까고 자사는 봐주지 않는가? | 경쟁사: "only 40min water resistance" vs 자사: "water resistance not independently rated" (쿠션) |
| **독립 전문가 목소리** | 창업자/내부 인용 외에 독립 전문가(dermatologist, 연구원 등)가 있는가? | 인용이 전부 창업자 → PR 보도자료 느낌 |
| **경험 구체성** | 테스트/경험 주장이 구체적인가? 누가, 몇 회, 어떤 조건에서? | "tested through months of sessions" (구체성 0) vs "tested by 3 players across 12 sessions on outdoor courts in 30°C+" |
| **테스팅 방법론** | 비교/리뷰 글이라면 평가 기준과 방법이 명시되어 있는가? | Wirecutter: 별도 methodology 섹션, 시기/장소/인원 명시. 우리: 없음 |
| **도메인 특화 맥락** | 해당 주제만의 고유한 맥락이 녹아있는가? 일반론에 그치지 않는가? | padel 글인데 padel 코트 특성(유리벽 반사광, 경기 시간)이 없음 → 아무 스포츠나 끼워넣은 느낌 |
| **1st-party 데이터** | 자체 테스트, 고객 후기, 사용 데이터 등 1st-party 경험이 포함되어 있는가? | 전부 외부 출처만 → "써본 적 없는 사람이 쓴 글" |
| **저자 E-E-A-T** | 실제 저자 이름 + bio + credentials가 있는가? "Editorial Team" 같은 익명이 아닌가? | "HAESKN Editorial Team" → 누군지 알 수 없음, AI/유령 필진 의심 |

**핵심 원칙**: 구조와 출처가 완벽해도, 진짜 경험과 솔직함이 없으면 "AI 글 티"가 난다. 이 카테고리는 글이 인간의 실제 경험에서 나온 것인지를 판단한다.

#### 6. Engine Fit (엔진 적합성)

brand profile의 타겟 엔진 또는 visibility gap 데이터 기준:

| 엔진 | 체크 포인트 |
|------|-----------|
| **ChatGPT** | 실용적/conversational 톤, ~2800 words, how-to 구조 |
| **Gemini** | Schema 강화, 구조화 데이터, YouTube 연계 고려 |
| **Perplexity** | 니치 전문성, 90일 이내 최신 데이터, 팩트 밀도 |
| **Grok** | 실시간 트렌드, 커뮤니티 반응 반영 |

---

### Report Format

```
## GEO Content Review — "{article title}"

### Summary
- **Overall**: ✅ Good / ⚠️ Needs Work / ❌ Major Issues
- **Word count**: {n} words
- **Article type**: {type}

### Scores

| Category | Score | Notes |
|----------|-------|-------|
| Structure & Quotability | ✅ / ⚠️ / ❌ | ... |
| Trust & Authority | ✅ / ⚠️ / ❌ | ... |
| Freshness | ✅ / ⚠️ / ❌ | ... |
| Brand Integration | ✅ / ⚠️ / ❌ | ... |
| Semantic Authenticity | ✅ / ⚠️ / ❌ | ... |
| Engine Fit | ✅ / ⚠️ / ❌ | ... |

### Issues Found
1. **[Category]** — {구체적 문제} → {권장 수정}
2. ...

### Recommended Patches
> 수정이 필요한 경우, `/aeo content update <id>` 에 바로 쓸 수 있는 patch 제안을 포함한다.
> 유저 확인 후 적용.
```

---

### Notes

- 리뷰는 **읽기 전용** — CUD Rule 대상이 아니다. 수정은 유저 확인 후 `/aeo content update`로.
- 외부에서 작성된 글(로컬 파일)도 리뷰 가능 — `<id>` 대신 파일 경로를 받으면 파일을 읽어서 동일한 체크리스트를 적용한다.
- brand profile이 없는 상태에서도 리뷰 가능하지만, Brand Integration 카테고리는 skip하고 그 사실을 명시한다.
