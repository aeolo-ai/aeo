---
title: Google Drive
description: Read files from a connected Google Drive folder
commands:
  - aeo drive
  - aeo drive list
  - aeo drive read
---

# Google Drive

연결된 Google Drive 폴더의 파일을 조회하고 읽을 수 있습니다.

## 사전 조건

Dashboard에서 Google Drive 폴더를 연결해야 합니다:
1. Dashboard → Integrations → Google Drive
2. SA 이메일(`geoclaw@tryaeolo.iam.gserviceaccount.com`)을 폴더에 **뷰어**로 초대
3. 폴더 ID 입력 → 연결 확인

## Commands

### `aeo drive` / `aeo drive list`

연결된 폴더의 파일 목록을 조회합니다.

```
aeo drive
```

출력: 파일명, 타입, 크기, ID 테이블

### `aeo drive read <file_id>`

특정 파일의 내용을 읽습니다.

```
aeo drive read 1abc2def3ghi
```

### 파일 타입별 처리

| 타입 | 처리 | 비용 |
|------|------|------|
| Google Docs | 텍스트로 변환 | 거의 0 |
| Google Sheets | CSV로 변환 | 거의 0 |
| 텍스트 파일 (txt, json, md) | 직접 읽기 | 거의 0 |
| **PDF** | **서버에서 텍스트 추출** (pdf-parse) | 낮음 |
| **이미지** (png, jpg 등) | **base64 반환** (5MB 이하) | 중간 |
| 기타 바이너리 | 메타 정보만 (파일명, 타입, 크기) | 0 |

**5MB 제한**: PDF/이미지 등 바이너리 파일은 5MB 초과 시 메타 정보만 반환됩니다.

## 보안

- **읽기 전용**: 파일 수정/삭제 불가 (SA = viewer, CLI에 write 커맨드 없음)
- **폴더 범위**: SA가 초대된 폴더만 접근 가능
- **SA 키 위치**: API 서버에만 보관 (에이전트 컨테이너에 없음)
- **프록시 구조**: CLI → Connector API → API Server (SA key) → Google Drive
