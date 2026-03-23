# aeo

CLI for [Aeolo](https://tryaeolo.com) — optimize your brand visibility in AI search engines (ChatGPT, Perplexity, Gemini, Grok).

## Install

```bash
curl -fsSL https://skills.tryaeolo.com | sh
```

Or with Homebrew:

```bash
brew install kithlabs/aeo/aeo
```

No runtime dependencies required — `aeo` is a single static binary.

## Update

```bash
# Auto-detects install method:
aeo update          # install.sh users → downloads latest
                    # Homebrew users → prints "brew upgrade aeo"

# Or explicitly:
brew upgrade aeo    # Homebrew
```

## Quick Start

```bash
# Authenticate
aeo auth login

# View your brand profile
aeo domain brand

# Check visibility across AI engines
aeo visibility check run

# List content
aeo content
```

## Commands

| Command | Description |
|---------|-------------|
| `aeo domain list` | List accessible domains |
| `aeo domain brand` | Show brand profile |
| `aeo domain audit` | Show latest audit report |
| `aeo visibility` | Show last visibility snapshot |
| `aeo visibility check run` | Trigger a new visibility check |
| `aeo strategy` | Show content strategy |
| `aeo content` | List content items |
| `aeo content propose` | Generate content proposals |
| `aeo metrics` | Article performance overview |
| `aeo prompts` | List tracked prompts |
| `aeo auth login` | Authenticate via browser |

Run `aeo --help` for full command reference.

## Development (submodule)

This repo is consumed as a git submodule in the [Aeolo monorepo](https://github.com/kithlabs/aeolo).

```bash
# After cloning the monorepo, init submodules
git submodule update --init --recursive

# Auto-update submodules on pull/checkout (recommended, one-time)
git config submodule.recurse true
```

### Releasing

Tags must be created **inside the submodule** (not the monorepo root):

```bash
cd packages/cli
git tag -a v0.X.Y -m "description"
git push origin v0.X.Y
```

This triggers GitHub Actions which:
1. GoReleaser → builds linux/darwin × amd64/arm64
2. Creates GitHub Release + uploads install.sh
3. Auto-updates Homebrew Formula (`kithlabs/homebrew-aeo`)

## For AI Agents

`aeo` is designed to be used by AI coding agents (Claude Code, Cursor, etc.) as a tool for GEO workflows. The CLI outputs structured JSON, making it easy for agents to parse and act on the data.

## License

[Apache License 2.0](LICENSE) — see [NOTICE](NOTICE) for attribution.
