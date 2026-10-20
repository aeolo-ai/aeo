package main

import (
	"bytes"
	"os"
	"path/filepath"
	"testing"
	"time"
)

func TestUpdateNoticeVersions(t *testing.T) {
	for _, tc := range []struct {
		latest, current string
		want            bool
	}{
		{"2.10.0", "2.9.9", true}, {"2.3.46", "2.3.46", false},
		{"2.3.45", "2.3.46", false}, {"garbage", "2.3.45", false},
		{"2.3.46-rc1", "2.3.45", false}, {"2.3.46", "dev", false},
	} {
		if got := newerNoticeVersion(tc.latest, tc.current); got != tc.want {
			t.Errorf("%s > %s = %v", tc.latest, tc.current, got)
		}
	}
}

func TestUpdateNoticeCache(t *testing.T) {
	for _, fetched := range []string{"2.3.46", ""} {
		t.Run("fetch_"+fetched, func(t *testing.T) {
			path := filepath.Join(t.TempDir(), "cache.json")
			now := time.Now()
			calls := 0
			fetch := func() string { calls++; return fetched }
			for i := 0; i < 2; i++ {
				if got := cachedNoticeVersion(path, now, fetch); got != fetched {
					t.Fatalf("got %q", got)
				}
			}
			if calls != 1 {
				t.Fatalf("fetches = %d", calls)
			}
			cachedNoticeVersion(path, now.Add(25*time.Hour), fetch)
			if calls != 2 {
				t.Fatalf("expired fetches = %d", calls)
			}
		})
	}
}

func TestUpdateNoticeDisabledAndOutput(t *testing.T) {
	path := filepath.Join(t.TempDir(), "cache.json")
	fetch := func() string { return "2.3.46" }
	var out bytes.Buffer
	notifyUpdate("1", "2.3.45", path, time.Now(), &out, func() string { t.Fatal("disabled fetched"); return "" })
	if out.Len() != 0 {
		t.Fatal(out.String())
	}
	if _, err := os.Stat(path); !os.IsNotExist(err) {
		t.Fatal("disabled created cache")
	}
	notifyUpdate("", "2.3.45", path, time.Now(), &out, fetch)
	if got := out.String(); got != "Update available: 2.3.45 → 2.3.46. Run: aeo update\n" {
		t.Fatal(got)
	}
}

func TestUpdateNoticeAfterOrdinaryCommand(t *testing.T) {
	t.Setenv("HOME", t.TempDir())
	t.Setenv("AEO_NO_UPDATE_CHECK", "0")
	cachePath := filepath.Join(filepath.Dir(configPath()), "update-check.json")
	cachedNoticeVersion(cachePath, time.Now(), func() string { return "999.0.0" })
	oldArgs, oldOut, oldErr := os.Args, os.Stdout, os.Stderr
	defer func() { os.Args, os.Stdout, os.Stderr = oldArgs, oldOut, oldErr }()
	out, err := os.CreateTemp(t.TempDir(), "stdout")
	if err != nil {
		t.Fatal(err)
	}
	defer out.Close()
	errOut, err := os.CreateTemp(t.TempDir(), "stderr")
	if err != nil {
		t.Fatal(err)
	}
	defer errOut.Close()
	os.Stdout, os.Stderr = out, errOut
	os.Args = []string{"aeo", "content", "--help"}
	main()
	raw, _ := os.ReadFile(errOut.Name())
	if len(raw) != 0 {
		t.Fatalf("help notice: %s", raw)
	}
	// An ordinary command returning locally still runs the shared completion hook.
	os.Args = []string{"aeo", "content"}
	main()
	raw, _ = os.ReadFile(errOut.Name())
	if !bytes.Contains(raw, []byte("Run: aeo update")) {
		t.Fatalf("missing notice: %s", raw)
	}
	raw, _ = os.ReadFile(out.Name())
	if bytes.Contains(raw, []byte("Update available:")) {
		t.Fatal("notice polluted stdout")
	}
}
