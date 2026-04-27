# GSC Indexing — Browser Automation

Google Search Console에서 "색인 생성 요청"을 자동화한다. CLI가 아닌 **브라우저 자동화**(Claude in Chrome)로 동작한다.

## Prerequisites

이 커맨드는 다른 aeo 커맨드와 달리 **브라우저 환경**이 필요하다. 실행 전 아래 3가지를 순서대로 확인하고, 하나라도 빠지면 유저에게 안내 후 중단한다.

### 1. Claude in Chrome 연결 확인

`tabs_context_mcp`를 호출한다.

- **성공** → 브라우저 연결됨, 다음 단계로
- **실패 / tool not found** → 브라우저 자동화를 사용할 수 없는 환경이다. 아래 메시지를 보여주고 중단:

```
GSC 색인 요청은 브라우저 자동화가 필요합니다.

필요한 것:
1. Chrome 브라우저 + "Claude in Chrome" 확장 프로그램 설치
2. Claude Code 데스크톱 앱 또는 브라우저 MCP가 연결된 환경

현재 환경에서는 이 기능을 사용할 수 없습니다.
수동으로 GSC에서 색인 요청을 진행해주세요:
→ https://search.google.com/search-console
```

### 2. GSC 탭 확인

`tabs_context_mcp` 결과에서 `search.google.com/search-console` URL을 포함하는 탭을 찾는다.

| 상태 | 처리 |
|------|------|
| GSC 드릴다운 페이지 열려있음 | Phase 1로 진행 |
| GSC 열려있지만 다른 페이지 | 드릴다운으로 네비게이트 |
| GSC 탭 없음 | 유저에게 도메인 물어본 후 새 탭 생성 |

### 3. GSC 로그인 확인

네비게이트 후 페이지가 로그인을 요구하거나 에러를 보여주면:

```
GSC에 로그인되어 있지 않습니다.

Chrome에서 직접 Google Search Console에 로그인한 후 다시 시도해주세요:
→ https://search.google.com/search-console
```

---

## Workflow

### Phase 1: URL 목록 추출

GSC 드릴다운 페이지 URL 패턴:
```
https://search.google.com/search-console/index/drilldown?resource_id=sc-domain%3A<domain>&item_key=CAMYFiAC
```

Known `item_key` 값:
- `CAMYFiAC` = "발견됨 - 현재 색인이 생성되지 않음" (Discovered - currently not indexed)

**JavaScript로 URL 추출** (접근성 트리는 긴 URL을 잘라먹으므로 DOM에서 직접 추출):

```javascript
// javascript_tool로 실행
const allEls = document.querySelectorAll('*');
const urls = [];
for (const el of allEls) {
  if (el.children.length === 0 || el.childNodes[0]?.nodeType === 3) {
    const text = el.textContent.trim();
    if (text.startsWith('https://') && text.length > 30 && !text.includes('sitemap') && !urls.includes(text)) {
      urls.push(text);
    }
  }
}
JSON.stringify(urls);
```

**페이지네이션 확인**: "총 N행 중 1~10"이 보이면 "페이지당 행 수"를 100+로 변경하거나 모든 페이지를 순회한 후 추출.

**URL 목록을 유저에게 확인받은 후** Phase 2로 진행한다 (CUD Rule 적용).

### Phase 2: 개별 색인 요청

각 URL에 대해 아래 사이클 반복:

#### Step 1 — 검색창에 URL 입력

```
read_page(filter: "interactive")  # combobox ref 찾기 (라벨: "<domain>에 있는 모든 URL 검사")
form_input(ref: "<combobox_ref>", value: "<full_url>")
key(Return)
wait(5s)
```

`form_input` 후 드롭다운이 나타나고 네비게이트하지 않으면 → 드롭다운 항목 클릭 (검색창 바로 아래, ~y:98).

#### Step 2 — "색인 생성 요청" 클릭

```
screenshot()   # 페이지 로드 확인
read_page(filter: "interactive")  # "색인 생성 요청" 버튼 ref 찾기
click("<button_ref>")
```

페이지 상태별 처리:

| 상태 | 표시 | 행동 |
|------|------|------|
| URL이 Google에 등록되어 있지 않음 | 미등록 | "색인 생성 요청" 클릭 |
| URL이 Google에 등록되어 있음 | 이미 색인됨 (녹색) | "색인 생성 요청" 클릭 (재요청) |
| Google에는 아직 알려지지 않은 URL | 완전 미등록 | "색인 생성 요청" 클릭 |

#### Step 3 — 색인 테스트 대기

"실제 URL의 색인을 생성할 수 있는지 테스트 중" 다이얼로그가 나타난다. 1-2분 소요.

```
wait(10s) × 6-7회
screenshot()  # 결과 확인
```

결과:

| 결과 | 기록 |
|------|------|
| 색인 생성 요청됨 (녹색 배너) | 성공 |
| 색인 생성 요청이 거부됨 (빨간/주황 배너) | 거부됨 |
| 다시 요청 표시 | 성공 (재요청) |

#### Step 4 — 닫기 후 다음 URL

```
click("닫기")
# → Step 1로 돌아가 다음 URL 처리
```

### Phase 3: 결과 리포트

```markdown
## GSC 색인 요청 결과 — {domain}

| # | URL (slug) | 상태 | 결과 |
|---|-----------|------|------|
| 1 | /beyond-the-beach... | 미등록 | 성공 |
| 2 | /some-article... | 이미 색인됨 | 성공 |
| 3 | /problem-url... | 미등록 | 거부됨 |

**합계**: N개 성공, N개 거부됨, N개 이미 색인됨
```

거부된 URL이 있으면 점검 사항 안내:
- 페이지 존재 여부 (404?)
- sitemap 포함 여부
- robots.txt 차단 여부

---

## Troubleshooting

| 문제 | 해결 |
|------|------|
| Chrome 확장 연결 끊김 | 5초 대기 후 재시도 — 보통 자동 복구됨 |
| 검색창 입력 후 네비게이트 안 됨 | combobox 클릭 → form_input → Return 순서로 재시도 |
| "색인 생성 요청" 버튼 클릭 안 됨 | read_page(filter: "interactive")로 정확한 ref 찾아서 ref로 클릭 |
| 페이지당 10개만 표시 | "페이지당 행 수" 드롭다운에서 100+ 선택 후 재추출 |

## Rate Limits

Google Search Console 일일 색인 요청 할당량: 보통 ~200건/일. 제한 에러 발생 시 중단하고 유저에게 안내.
