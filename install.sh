#!/bin/sh
set -e

# aeo CLI installer
# Usage: curl -fsSL https://skills.tryaeolo.com | sh

REPO="kithlabs/aeo"
INSTALL_DIR="${AEO_INSTALL_DIR:-/usr/local/bin}"
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

# ── Download ─────────────────────────────────────────────────────────────────

TARBALL="aeo_${OS}_${ARCH}.tar.gz"
URL="https://github.com/${REPO}/releases/download/v${AEO_VERSION}/${TARBALL}"

echo "Installing aeo v${AEO_VERSION} (${OS}/${ARCH})..."

TMPDIR=$(mktemp -d)
trap 'rm -rf "$TMPDIR"' EXIT

curl -fsSL "$URL" -o "${TMPDIR}/${TARBALL}"
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
