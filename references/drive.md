---
title: Google Drive
description: Read files from a connected Google Drive folder
commands:
  - aeo drive
  - aeo drive list
  - aeo drive read
---

# Google Drive

You can browse and read files from a connected Google Drive folder.

## Prerequisites

You must connect a Google Drive folder from the Dashboard:
1. Dashboard → Integrations → Google Drive
2. Invite the SA email (`geoclaw@tryaeolo.iam.gserviceaccount.com`) to the folder as a **Viewer**
3. Enter the folder ID → Confirm connection

## Commands

### `aeo drive` / `aeo drive list`

Lists the files in the connected folder.

```
aeo drive
```

Output: Table of file names, types, sizes, and IDs. Folders are shown with type `folder`.

### `aeo drive list --folder <folder_id>`

Browse a subfolder by its ID. Use this when `aeo drive` shows folders in the list.

```
aeo drive list --folder 1abc2def3ghi
```

### `aeo drive read <file_id>`

Reads the contents of a specific file. Returns an error if the target is a folder (use `drive list --folder` instead).

```
aeo drive read 1abc2def3ghi
```

### File Type Handling

| Type | Handling | Cost |
|------|----------|------|
| Google Docs | Converted to text | Near 0 |
| Google Sheets | Converted to CSV | Near 0 |
| Text files (txt, json, md, csv) | Read directly | Near 0 |
| **PDF** | Server-side text extraction (pdf-parse) | Low |
| **XLSX / XLS / XLSM** | All sheets → CSV, capped at 200 rows per sheet | Low |
| **DOCX** | Raw text extraction (mammoth) | Low |
| **Images** (png, jpg, etc.) | base64 returned (under 5MB) | Medium |
| Other binary (.doc, .pptx, .key, zip, etc.) | Metadata only (name, type, size) | 0 |

**5MB transfer limit**: Binary files (PDF, XLSX, DOCX, images) exceeding 5MB return metadata only. To work around, split or export to a smaller format before uploading.

**XLSX row cap**: Each sheet is truncated at 200 rows with a `[… truncated, N more row(s)]` marker. For larger datasets, export the relevant range as a separate sheet/file.

**Not supported**: Legacy `.doc` (use `.docx`), `.pptx` (export slides as PDF), Apple iWork (`.pages`, `.numbers`, `.key`).

## Security

- **Read-only**: Cannot modify or delete files (SA = viewer, no write commands in CLI)
- **Folder-scoped**: Only accessible to folders the SA has been invited to
- **SA key location**: Stored only on the API server (not in agent containers)
- **Proxy architecture**: CLI → Connector API → API Server (SA key) → Google Drive
