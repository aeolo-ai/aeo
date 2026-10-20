package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

type noticeCache struct {
	CheckedAt time.Time `json:"checkedAt"`
	Latest    string    `json:"latest,omitempty"`
}

// Notices are best-effort and stderr-only, leaving command JSON intact.
func checkLatestVersion() {
	notifyUpdate(os.Getenv("AEO_NO_UPDATE_CHECK"), version,
		filepath.Join(filepath.Dir(configPath()), "update-check.json"), time.Now(), os.Stderr, fetchNoticeVersion)
}

func notifyUpdate(disabled, current, path string, now time.Time, out io.Writer, fetch func() string) {
	disabled = strings.TrimSpace(disabled)
	if disabled != "" && disabled != "0" && !strings.EqualFold(disabled, "false") {
		return
	}
	if latest := cachedNoticeVersion(path, now, fetch); newerNoticeVersion(latest, current) {
		fmt.Fprintf(out, "Update available: %s → %s. Run: aeo update\n", current, latest)
	}
}

// Only stable numeric release versions qualify; never suggest a downgrade.
func noticeVersion(v string) ([]uint64, bool) {
	parts := strings.Split(v, ".")
	if len(parts) != 3 {
		return nil, false
	}
	nums := make([]uint64, 3)
	for i, p := range parts {
		if p == "" {
			return nil, false
		}
		for _, r := range p {
			if r < '0' || r > '9' {
				return nil, false
			}
		}
		n, err := strconv.ParseUint(p, 10, 64)
		if err != nil {
			return nil, false
		}
		nums[i] = n
	}
	return nums, true
}
func newerNoticeVersion(latest, current string) bool {
	a, ok := noticeVersion(latest)
	if !ok {
		return false
	}
	b, ok := noticeVersion(current)
	if !ok {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return a[i] > b[i]
		}
	}
	return false
}

func cachedNoticeVersion(path string, now time.Time, fetch func() string) string {
	var cache noticeCache
	if raw, err := os.ReadFile(path); err == nil {
		_ = json.Unmarshal(raw, &cache)
	}
	if _, ok := noticeVersion(cache.Latest); !ok {
		cache.Latest = ""
	}
	if age := now.Sub(cache.CheckedAt); age >= 0 && age < 24*time.Hour {
		return cache.Latest
	}
	if latest := fetch(); latest != "" {
		if _, ok := noticeVersion(latest); ok {
			cache.Latest = latest
		}
	}
	// Cache failed checks too, so offline commands do not repeatedly wait.
	cache.CheckedAt = now
	if os.MkdirAll(filepath.Dir(path), 0700) == nil {
		if tmp, err := os.CreateTemp(filepath.Dir(path), ".update-check-*"); err == nil {
			defer os.Remove(tmp.Name())
			err = json.NewEncoder(tmp).Encode(cache)
			closeErr := tmp.Close()
			if err == nil && closeErr == nil {
				_ = os.Rename(tmp.Name(), path)
			}
		}
	}
	return cache.Latest
}

func fetchNoticeVersion() string {
	client := &http.Client{Timeout: 2 * time.Second}
	resp, err := client.Get("https://api.github.com/repos/aeolo-ai/aeo/releases/latest")
	if err != nil {
		return ""
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return ""
	}
	var release struct {
		TagName string `json:"tag_name"`
	}
	if json.NewDecoder(io.LimitReader(resp.Body, 65536)).Decode(&release) != nil {
		return ""
	}
	return strings.TrimPrefix(release.TagName, "v")
}
