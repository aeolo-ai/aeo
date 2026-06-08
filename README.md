<div align="center">

# aeo

**Get your brand cited by AI search — from the terminal.**
**AI 검색이 인용하는 브랜드를, 터미널에서.**

[![Release](https://img.shields.io/github/v/release/kithlabs/aeo?color=black)](https://github.com/kithlabs/aeo/releases)
[![License](https://img.shields.io/badge/license-Apache--2.0-blue.svg)](LICENSE)
[![Homebrew](https://img.shields.io/badge/brew-kithlabs%2Faeo-orange)](https://github.com/kithlabs/homebrew-aeo)

</div>

`aeo` is the official command-line interface for [Aeolo](https://tryaeolo.com), a Generative Engine Optimization platform. Track how your brand shows up in **ChatGPT, Perplexity, Gemini, and Grok**, manage the strategy + content that fixes it, and deploy directly to your channels — all without leaving your terminal.

`aeo` 는 [Aeolo](https://tryaeolo.com) 의 공식 CLI 다. **ChatGPT, Perplexity, Gemini, Grok** 이 브랜드를 어떻게 인용하는지 추적하고, 가시성을 끌어올릴 전략·콘텐츠를 관리하고, 채널로 바로 배포하는 — 이 모든 걸 터미널에서.

> Built to be driven by AI coding agents (Claude Code, Cursor, Codex). One install, and your agent gets a full GEO toolchain. → [Use with Claude Code](#use-with-claude-code)

---

## Install

```bash
# Homebrew
brew install kithlabs/aeo/aeo

# Or one-liner
curl -fsSL https://skills.tryaeolo.com | sh
```

Single static Go binary. No runtime dependencies. macOS + Linux, amd64 + arm64.

To upgrade:

```bash
aeo update          # auto-detects Homebrew vs install.sh
brew upgrade aeo    # explicit Homebrew path
```

## 60-second tour

```bash
aeo auth login                   # browser-based OAuth → API key in ~/.config/aeo/
aeo domain list                  # pick a domain
aeo domain switch <id>

aeo visibility                   # last snapshot across 4 AI engines
aeo content                      # what's drafted, scheduled, deployed
aeo billing subscription         # subscription tier + production credit balance
aeo metrics traffic --days 30    # GA4 + Search Console for the same domain
```

Sample output:

```
$ aeo visibility

# tryaeolo.com — Visibility Snapshot

| Engine     | Mentioned | Rate |
|------------|-----------|------|
| chatgpt    |   9/12    |  75% |
| perplexity |   7/12    |  58% |
| gemini     |   5/12    |  42% |
| grok       |   4/12    |  33% |

## Visibility Gaps (Not Mentioned)
- best GEO platform for Shopify
- how to optimize for ChatGPT citations
- ...

_Full report: https://tryaeolo.com/report/.../visibility_
```

## What you can do

| Area | Commands |
|------|----------|
| **Brand & domain** | `aeo domain list / brand / brand update / audit / setup` |
| **Visibility** | `aeo visibility` · `aeo visibility check run / poll` |
| **Site audit** | `aeo audit run / poll` |
| **Content lifecycle** | `aeo content list / get / write / jobs / update / preview / deploy / redeploy / import` |
| **Strategy** | `aeo strategy` · `aeo strategy update` |
| **Prompts** | `aeo prompts list / add / update / delete` |
| **Channels** | `aeo channel add / connect / disconnect` (Shopify, LinkedIn, Threads, Reddit) |
| **Channel posts** | `aeo post list / import / approve / publish` |
| **Analysis** | `aeo reference analyze` · `aeo video analyze` |
| **Metrics** | `aeo metrics overview / article <id> / traffic` |
| **Drive integration** | `aeo drive list / read <fileId>` (read-only Google Drive) |
| **Account & billing** | `aeo whoami` · `aeo billing subscription / credits / ledger` · `aeo auth login / status / logout` |
| **Send feedback** | `aeo feedback "msg"` or `aeo feedback` (opens `$EDITOR`) |

Production actions reserve and capture Aeolo credits server-side. Failed background jobs are refunded by the worker finalizers. Current costs: visibility checks cost 1 credit per prompt × engine; site audit starts at 3 credits per 5 pages; writing, reference analysis, video analysis, channel voice analysis, and image swap cost 5 credits each.

Run `aeo --help` for the complete reference, or `aeo <command> --help` for detail on any verb.

## Use with Claude Code

`aeo` ships a [Claude Code skill](SKILL.md) that turns your agent into a full GEO co-pilot. Drop the skill into your Claude Code skills directory (`~/.claude/skills/aeo/`) — see [SKILL.md](SKILL.md) for the trigger phrases — then ask in plain language:

```
> 우리 브랜드 가시성 어때?
> Show me where ChatGPT misses tryaeolo.com.
> Write a 1500-word article about GEO best practices and deploy it to Shopify.
```

The agent picks the right `aeo` commands, parses results, and chains them into multi-step workflows. Same pattern works in Cursor, Codex CLI, and any agent that can call shell commands — `aeo` outputs clean Markdown / JSON specifically so it can be consumed by an LLM downstream.

## How auth works

`aeo auth login` opens a browser device-flow against [tryaeolo.com](https://tryaeolo.com). On success, an API key is written to `~/.config/aeo/config.json` and used as a Bearer token on all subsequent calls. No telemetry, no background processes — every command is a single HTTPS request to the Aeolo API.

To use a non-default API base (self-hosting, staging):

```bash
aeo auth login --api-base https://api.example.com
# or via env
AEOLO_API_BASE=... AEOLO_API_KEY=... aeo whoami
```

## Development (monorepo)

This repo is consumed as a git submodule in the [Aeolo monorepo](https://github.com/kithlabs/aeolo). Standalone development is supported — clone this repo directly and `go build .`.

```bash
git submodule update --init --recursive
git config submodule.recurse true   # auto-update submodules on pull
```

### Releasing

Tags must be created **inside the submodule** (not the monorepo root):

```bash
cd packages/cli
git tag -a v1.X.Y -m "description"
git push origin v1.X.Y
```

This triggers GitHub Actions, which:
1. **GoReleaser** — builds linux/darwin × amd64/arm64
2. Creates GitHub Release + uploads `install.sh`
3. Auto-updates the [Homebrew tap](https://github.com/kithlabs/homebrew-aeo)

## Contributing

Bug reports and feature ideas are welcome — open an [issue](https://github.com/kithlabs/aeo/issues), or send feedback inline:

```bash
aeo feedback "your message here"
```

PRs welcome for bug fixes and small improvements. For larger changes, please open an issue first to discuss direction.

## License

[Apache License 2.0](LICENSE) — see [NOTICE](NOTICE) for attribution.
