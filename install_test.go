package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
)

const fakeTarballContent = "fake-tarball-bytes-for-test\n"

func fakeTarballChecksum() string {
	sum := sha256.Sum256([]byte(fakeTarballContent))
	return hex.EncodeToString(sum[:])
}

// writeFakeCurl generates a fake `curl` that serves fixture files instead of
// hitting the network: tarballFixture for the release archive URL,
// checksumsFixture for checksums.txt. It mirrors install.sh's real
// `curl -fsSL <url> -o <path>` invocation order (URL before -o).
func writeFakeCurl(t *testing.T, path, tarballFixture, checksumsFixture string) {
	t.Helper()
	script := fmt.Sprintf(`#!/bin/sh
URL=""
while [ "$#" -gt 0 ]; do
  case "$1" in
    -o)
      shift
      case "$URL" in
        *checksums.txt) cp %q "$1" ;;
        *) cp %q "$1" ;;
      esac
      exit 0
      ;;
    http*) URL="$1" ;;
  esac
  shift
done
exit 1
`, checksumsFixture, tarballFixture)
	writeExecutable(t, path, script)
}

func TestInstallerReusesActiveAEODirectory(t *testing.T) {
	root := t.TempDir()
	activeDir := filepath.Join(root, "active")
	fallbackDir := filepath.Join(root, "fallback")
	toolDir := filepath.Join(root, "tools")
	for _, dir := range []string{activeDir, fallbackDir, toolDir} {
		if err := os.Mkdir(dir, 0o755); err != nil {
			t.Fatal(err)
		}
	}

	writeExecutable(t, filepath.Join(activeDir, "aeo"), "#!/bin/sh\necho 'aeo 2.3.4 (native)'\n")
	writeExecutable(t, filepath.Join(toolDir, "uname"), `#!/bin/sh
case "$1" in
  -s) echo Darwin ;;
  -m) echo arm64 ;;
esac
`)
	tarballFixture := filepath.Join(root, "fixture-tarball")
	if err := os.WriteFile(tarballFixture, []byte(fakeTarballContent), 0o644); err != nil {
		t.Fatal(err)
	}
	checksumsFixture := filepath.Join(root, "fixture-checksums.txt")
	checksumsContent := fmt.Sprintf("%s  aeo_darwin_arm64.tar.gz\n", fakeTarballChecksum())
	if err := os.WriteFile(checksumsFixture, []byte(checksumsContent), 0o644); err != nil {
		t.Fatal(err)
	}
	writeFakeCurl(t, filepath.Join(toolDir, "curl"), tarballFixture, checksumsFixture)
	writeExecutable(t, filepath.Join(toolDir, "tar"), `#!/bin/sh
while [ "$#" -gt 0 ]; do
  if [ "$1" = "-C" ]; then
    shift
    cat > "$1/aeo" <<'EOF'
#!/bin/sh
echo 'aeo 2.3.5 (native)'
EOF
    chmod +x "$1/aeo"
    exit 0
  fi
  shift
done
exit 1
`)

	installer, err := os.ReadFile("install.sh")
	if err != nil {
		t.Fatal(err)
	}
	safeInstaller := strings.Replace(
		string(installer),
		"/usr/local/bin",
		fallbackDir,
		1,
	)
	installerPath := filepath.Join(root, "install.sh")
	if err := os.WriteFile(installerPath, []byte(safeInstaller), 0o755); err != nil {
		t.Fatal(err)
	}

	cmd := exec.Command("sh", installerPath)
	cmd.Env = []string{
		"AEO_VERSION=2.3.5",
		"PATH=" + strings.Join([]string{toolDir, activeDir, "/usr/bin", "/bin"}, string(os.PathListSeparator)),
	}
	if output, err := cmd.CombinedOutput(); err != nil {
		t.Fatalf("installer failed: %v\n%s", err, output)
	}

	output, err := exec.Command(filepath.Join(activeDir, "aeo"), "--version").CombinedOutput()
	if err != nil {
		t.Fatal(err)
	}
	if got := strings.TrimSpace(string(output)); got != "aeo 2.3.5 (native)" {
		t.Fatalf("active aeo was not replaced: %q", got)
	}
	if _, err := os.Stat(filepath.Join(fallbackDir, "aeo")); !os.IsNotExist(err) {
		t.Fatalf("installer unexpectedly wrote to fallback directory: %v", err)
	}
}

func TestInstallerRejectsChecksumMismatch(t *testing.T) {
	root := t.TempDir()
	fallbackDir := filepath.Join(root, "fallback")
	toolDir := filepath.Join(root, "tools")
	for _, dir := range []string{fallbackDir, toolDir} {
		if err := os.Mkdir(dir, 0o755); err != nil {
			t.Fatal(err)
		}
	}

	writeExecutable(t, filepath.Join(toolDir, "uname"), `#!/bin/sh
case "$1" in
  -s) echo Darwin ;;
  -m) echo arm64 ;;
esac
`)
	tarballFixture := filepath.Join(root, "fixture-tarball")
	if err := os.WriteFile(tarballFixture, []byte(fakeTarballContent), 0o644); err != nil {
		t.Fatal(err)
	}
	checksumsFixture := filepath.Join(root, "fixture-checksums.txt")
	// Deliberately wrong hash — simulates a tampered/corrupted download.
	checksumsContent := "0000000000000000000000000000000000000000000000000000000000000000  aeo_darwin_arm64.tar.gz\n"
	if err := os.WriteFile(checksumsFixture, []byte(checksumsContent), 0o644); err != nil {
		t.Fatal(err)
	}
	writeFakeCurl(t, filepath.Join(toolDir, "curl"), tarballFixture, checksumsFixture)
	// No `tar` on PATH: if the installer wrongly proceeds past the checksum
	// check, the failure mode should still be "never extracted", not a
	// coincidentally-successful extraction.

	installer, err := os.ReadFile("install.sh")
	if err != nil {
		t.Fatal(err)
	}
	safeInstaller := strings.Replace(string(installer), "/usr/local/bin", fallbackDir, 1)
	installerPath := filepath.Join(root, "install.sh")
	if err := os.WriteFile(installerPath, []byte(safeInstaller), 0o755); err != nil {
		t.Fatal(err)
	}

	cmd := exec.Command("sh", installerPath)
	cmd.Env = []string{
		"AEO_VERSION=2.3.5",
		"PATH=" + strings.Join([]string{toolDir, "/usr/bin", "/bin"}, string(os.PathListSeparator)),
	}
	output, err := cmd.CombinedOutput()
	if err == nil {
		t.Fatalf("installer unexpectedly succeeded with a mismatched checksum:\n%s", output)
	}
	if !strings.Contains(string(output), "checksum mismatch") {
		t.Fatalf("expected a checksum mismatch error, got:\n%s", output)
	}
	if _, err := os.Stat(filepath.Join(fallbackDir, "aeo")); !os.IsNotExist(err) {
		t.Fatalf("installer wrote a binary despite the checksum mismatch: %v", err)
	}
}

func TestInstallerPreservesHomebrewInstall(t *testing.T) {
	root := t.TempDir()
	binDir := filepath.Join(root, "bin")
	cellarDir := filepath.Join(root, "Cellar", "aeo", "2.3.4", "bin")
	toolDir := filepath.Join(root, "tools")
	fallbackDir := filepath.Join(root, "fallback")
	for _, dir := range []string{binDir, cellarDir, toolDir, fallbackDir} {
		if err := os.MkdirAll(dir, 0o755); err != nil {
			t.Fatal(err)
		}
	}

	cellarBinary := filepath.Join(cellarDir, "aeo")
	writeExecutable(t, cellarBinary, "#!/bin/sh\necho 'aeo 2.3.4 (native)'\n")
	activeBinary := filepath.Join(binDir, "aeo")
	if err := os.Symlink(filepath.Join("..", "Cellar", "aeo", "2.3.4", "bin", "aeo"), activeBinary); err != nil {
		t.Fatal(err)
	}
	writeExecutable(t, filepath.Join(toolDir, "uname"), `#!/bin/sh
case "$1" in
  -s) echo Darwin ;;
  -m) echo arm64 ;;
esac
`)
	brewLog := filepath.Join(root, "brew.log")
	writeExecutable(t, filepath.Join(toolDir, "brew"), fmt.Sprintf(`#!/bin/sh
echo "$*" >> %q
if [ "$1" = "upgrade" ]; then
  cat > %q <<'EOF'
#!/bin/sh
echo 'aeo 2.3.5 (native)'
EOF
  chmod +x %q
fi
`, brewLog, cellarBinary, cellarBinary))
	writeExecutable(t, filepath.Join(toolDir, "curl"), "#!/bin/sh\nexit 99\n")

	installer, err := os.ReadFile("install.sh")
	if err != nil {
		t.Fatal(err)
	}
	safeInstaller := strings.Replace(string(installer), "/usr/local/bin", fallbackDir, 1)
	installerPath := filepath.Join(root, "install.sh")
	if err := os.WriteFile(installerPath, []byte(safeInstaller), 0o755); err != nil {
		t.Fatal(err)
	}

	cmd := exec.Command("sh", installerPath)
	cmd.Env = []string{
		"AEO_VERSION=2.3.5",
		"PATH=" + strings.Join([]string{toolDir, binDir, "/usr/bin", "/bin"}, string(os.PathListSeparator)),
	}
	if output, err := cmd.CombinedOutput(); err != nil {
		t.Fatalf("installer failed: %v\n%s", err, output)
	}

	if info, err := os.Lstat(activeBinary); err != nil || info.Mode()&os.ModeSymlink == 0 {
		t.Fatalf("Homebrew symlink was not preserved: info=%v err=%v", info, err)
	}
	output, err := exec.Command(activeBinary, "--version").CombinedOutput()
	if err != nil {
		t.Fatal(err)
	}
	if got := strings.TrimSpace(string(output)); got != "aeo 2.3.5 (native)" {
		t.Fatalf("Homebrew aeo was not upgraded: %q", got)
	}
	logged, err := os.ReadFile(brewLog)
	if err != nil {
		t.Fatal(err)
	}
	if got := string(logged); !strings.Contains(got, "update\n") || !strings.Contains(got, "upgrade aeolo-ai/aeo/aeo\n") {
		t.Fatalf("unexpected brew calls:\n%s", got)
	}
}

func writeExecutable(t *testing.T, path, body string) {
	t.Helper()
	if err := os.WriteFile(path, []byte(body), 0o755); err != nil {
		t.Fatal(err)
	}
}
