package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
)

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
	writeExecutable(t, filepath.Join(toolDir, "curl"), `#!/bin/sh
while [ "$#" -gt 0 ]; do
  if [ "$1" = "-o" ]; then
    shift
    : > "$1"
    exit 0
  fi
  shift
done
exit 1
`)
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
