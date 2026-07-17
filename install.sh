#!/bin/sh
set -e

# aeo CLI installer
# Usage: curl -fsSL https://skills.tryaeolo.com | sh

REPO="aeolo-ai/aeo"
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
  LATEST_URL=$(curl -fsSLI --connect-timeout 10 --max-time 60 -o /dev/null -w '%{url_effective}' \
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

curl -fsSL --connect-timeout 10 --max-time 300 "$URL" -o "${TMPDIR}/${TARBALL}"

# ── Verify checksum ──────────────────────────────────────────────────────────
# Fetch the release checksums.txt and verify the tarball BEFORE extracting, so a
# tampered download (MITM, compromised asset) can't be unpacked and executed.
# Fail closed: abort if checksums are missing, unmatched, or don't verify.

CHECKSUMS_URL="https://github.com/${REPO}/releases/download/v${AEO_VERSION}/checksums.txt"
if ! curl -fsSL --connect-timeout 10 --max-time 60 "$CHECKSUMS_URL" -o "${TMPDIR}/checksums.txt"; then
  echo "Error: could not download checksums.txt from ${CHECKSUMS_URL}"
  exit 1
fi

EXPECTED=$(grep " ${TARBALL}\$" "${TMPDIR}/checksums.txt" | awk '{print $1}')
if [ -z "$EXPECTED" ]; then
  echo "Error: no checksum listed for ${TARBALL}"
  exit 1
fi

if command -v sha256sum >/dev/null 2>&1; then
  ACTUAL=$(sha256sum "${TMPDIR}/${TARBALL}" | awk '{print $1}')
elif command -v shasum >/dev/null 2>&1; then
  ACTUAL=$(shasum -a 256 "${TMPDIR}/${TARBALL}" | awk '{print $1}')
else
  echo "Error: no sha256 tool (sha256sum/shasum) available to verify download"
  exit 1
fi

if [ "$EXPECTED" != "$ACTUAL" ]; then
  echo "Error: checksum mismatch for ${TARBALL}"
  echo "  expected: $EXPECTED"
  echo "  actual:   $ACTUAL"
  exit 1
fi
echo "✓ Checksum verified"

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
