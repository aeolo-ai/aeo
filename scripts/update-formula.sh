#!/usr/bin/env bash
# Generate Homebrew Formula from release checksums
# Usage: ./scripts/update-formula.sh <version> <checksums_file> <output_file>
set -euo pipefail

VERSION="$1"
CHECKSUMS="$2"
OUTPUT="$3"
BASE="https://github.com/kithlabs/aeo/releases/download/v${VERSION}"

sha() { grep "$1" "$CHECKSUMS" | awk '{print $1}'; }

cat > "$OUTPUT" <<FORMULA
class Aeo < Formula
  desc "GEO CLI for AI search engine visibility"
  homepage "https://github.com/kithlabs/aeo"
  license "Apache-2.0"

  on_macos do
    if Hardware::CPU.arm?
      url "${BASE}/aeo_darwin_arm64.tar.gz"
      sha256 "$(sha aeo_darwin_arm64.tar.gz)"
    else
      url "${BASE}/aeo_darwin_amd64.tar.gz"
      sha256 "$(sha aeo_darwin_amd64.tar.gz)"
    end
  end

  on_linux do
    if Hardware::CPU.arm?
      url "${BASE}/aeo_linux_arm64.tar.gz"
      sha256 "$(sha aeo_linux_arm64.tar.gz)"
    else
      url "${BASE}/aeo_linux_amd64.tar.gz"
      sha256 "$(sha aeo_linux_amd64.tar.gz)"
    end
  end

  def install
    bin.install "aeo"
  end

  test do
    assert_match "aeo", shell_output("#{bin}/aeo --version")
  end
end
FORMULA
