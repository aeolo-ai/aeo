package main

import (
	"os"
	"path/filepath"
	"reflect"
	"testing"
)

func TestPageInputFilePreservesHTMLAndRevision(t *testing.T) {
	payload := `{"expectedRevision":6,"bodyHtml":"<h1>Title</h1>\n<p>$5 and quotes</p>"}`
	path := filepath.Join(t.TempDir(), "page.json")
	if err := os.WriteFile(path, []byte(payload), 0600); err != nil {
		t.Fatal(err)
	}
	got, err := pageInputArgs([]string{"pages", "update", "page-id", "--input-file", path, "--channel-id", "channel"})
	if err != nil || !reflect.DeepEqual(got, []string{"pages", "update", "page-id", "--input-json", payload, "--channel-id", "channel"}) {
		t.Fatalf("%v %v", got, err)
	}
	if _, err = pageInputArgs([]string{"pages", "create", "--input-file", path, "--input-json", "{}"}); err == nil {
		t.Fatal("ambiguous input accepted")
	}
}
