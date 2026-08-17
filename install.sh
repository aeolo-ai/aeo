#!/bin/sh
set -e

# aeo CLI installer
# Usage: curl -fsSL https://skills.aeolo.io | sh

REPO="aeolo-ai/aeo"
BINARY="aeo"

# ── Detect OS & Arch ─────────────────────────────────────────────────────────

OS="$(uname -s)"
ARCH="$(uname -m)"

case "$OS" in
  Linux)   OS="linux" ;;
  Darwin)  OS="darwin" ;;
  *)       echo "Error: unsupported OS: $OS"; exit 1 ;;
esac

case "$ARCH" in
  x86_64|amd64)  ARCH="amd64" ;;
  arm64|aarch64) ARCH="arm64" ;;
  *)             echo "Error: unsupported architecture: $ARCH"; exit 1 ;;
esac

# ── Resolve latest version ───────────────────────────────────────────────────
# Follow the `/releases/latest` redirect instead of hitting the JSON API —
# unauthenticated api.github.com is rate-limited to 60 req/hr/IP and shared
# households / offices burn through that easily, producing a 403. The HTML
# endpoint has no quota.

if [ -z "$AEO_VERSION" ]; then
  LATEST_URL=$(curl -fsSLI -o /dev/null -w '%{url_effective}' \
    "https://github.com/${REPO}/releases/latest")
  AEO_VERSION=$(printf '%s' "$LATEST_URL" | sed -E 's|.*/tag/v?([^/]+)$|\1|')
  if [ -z "$AEO_VERSION" ] || [ "$AEO_VERSION" = "$LATEST_URL" ]; then
    echo "Error: could not determine latest version"
    exit 1
  fi
fi

# ── Preserve existing install location / package manager ─────────────────────

if [ -n "$AEO_INSTALL_DIR" ]; then
  INSTALL_DIR="$AEO_INSTALL_DIR"
else
  EXISTING_AEO=$(command -v "$BINARY" 2>/dev/null || true)
  if [ -n "$EXISTING_AEO" ]; then
    LINK_TARGET=""
    if [ -L "$EXISTING_AEO" ]; then
      LINK_TARGET=$(readlink "$EXISTING_AEO" 2>/dev/null || true)
    fi

    case "$LINK_TARGET" in
      *Cellar/aeo/*)
        if ! command -v brew >/dev/null 2>&1; then
          echo "Error: existing aeo is managed by Homebrew, but brew is not on PATH"
          exit 1
        fi
        echo "Updating Homebrew-managed aeo to v${AEO_VERSION}..."
        brew update
        brew upgrade aeolo-ai/aeo/aeo
        INSTALLED_VERSION=$("$EXISTING_AEO" --version | awk '{print $2}' | sed 's/^v//')
        if [ "$INSTALLED_VERSION" != "$AEO_VERSION" ]; then
          echo "Error: Homebrew left aeo at v${INSTALLED_VERSION}; expected v${AEO_VERSION}"
          exit 1
        fi
        echo ""
        echo "✓ aeo v${AEO_VERSION} updated through Homebrew"
        exit 0
        ;;
    esac

    INSTALL_DIR=$(dirname "$EXISTING_AEO")
  else
    INSTALL_DIR="/usr/local/bin"
  fi
fi

# ── Download ─────────────────────────────────────────────────────────────────

TARBALL="aeo_${OS}_${ARCH}.tar.gz"
CHECKSUMS="checksums.txt"
BASE_URL="https://github.com/${REPO}/releases/download/v${AEO_VERSION}"

echo "Installing aeo v${AEO_VERSION} (${OS}/${ARCH})..."

TMPDIR=$(mktemp -d)
trap 'rm -rf "$TMPDIR"' EXIT

curl -fsSL "${BASE_URL}/${TARBALL}" -o "${TMPDIR}/${TARBALL}"
curl -fsSL "${BASE_URL}/${CHECKSUMS}" -o "${TMPDIR}/${CHECKSUMS}"

# ── Verify checksum ───────────────────────────────────────────────────────────
# GoReleaser publishes checksums.txt alongside every release archive — verify
# the downloaded tarball against it instead of extracting it unchecked.

EXPECTED_SHA=$(awk -v f="$TARBALL" '$2 == f { print $1 }' "${TMPDIR}/${CHECKSUMS}")
if [ -z "$EXPECTED_SHA" ]; then
  echo "Error: no checksum entry for ${TARBALL} in checksums.txt"
  exit 1
fi

if command -v sha256sum >/dev/null 2>&1; then
  ACTUAL_SHA=$(sha256sum "${TMPDIR}/${TARBALL}" | awk '{print $1}')
elif command -v shasum >/dev/null 2>&1; then
  ACTUAL_SHA=$(shasum -a 256 "${TMPDIR}/${TARBALL}" | awk '{print $1}')
else
  echo "Error: no sha256 tool found (need sha256sum or shasum)"
  exit 1
fi

if [ "$EXPECTED_SHA" != "$ACTUAL_SHA" ]; then
  echo "Error: checksum mismatch for ${TARBALL}"
  echo "  expected: ${EXPECTED_SHA}"
  echo "  actual:   ${ACTUAL_SHA}"
  exit 1
fi

tar xzf "${TMPDIR}/${TARBALL}" -C "$TMPDIR"

# ── Install ──────────────────────────────────────────────────────────────────

if [ -w "$INSTALL_DIR" ]; then
  mv "${TMPDIR}/${BINARY}" "${INSTALL_DIR}/${BINARY}"
else
  echo "Need sudo to install to ${INSTALL_DIR}"
  sudo mv "${TMPDIR}/${BINARY}" "${INSTALL_DIR}/${BINARY}"
fi

chmod +x "${INSTALL_DIR}/${BINARY}"

echo ""
echo "✓ aeo v${AEO_VERSION} installed to ${INSTALL_DIR}/${BINARY}"
echo ""
echo "Get started:"
echo "  aeo auth login"
echo ""
