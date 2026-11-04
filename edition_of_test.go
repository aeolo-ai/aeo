package main

import "testing"

func TestApplyEditionOf(t *testing.T) {
	body := map[string]any{}
	applyEditionOf(body, []string{"content", "import", "--title", "x", "--edition-of", "00000000-0000-4000-8000-0000000000aa", "--language", "ja"})
	if body["editionOf"] != "00000000-0000-4000-8000-0000000000aa" {
		t.Fatalf("editionOf = %v", body["editionOf"])
	}

	body = map[string]any{}
	applyEditionOf(body, []string{"content", "import", "--edition-of=00000000-0000-4000-8000-0000000000bb"})
	if body["editionOf"] != "00000000-0000-4000-8000-0000000000bb" {
		t.Fatalf("editionOf (= form) = %v", body["editionOf"])
	}

	// Absent flag: the body is exactly what it was.
	body = map[string]any{}
	applyEditionOf(body, []string{"content", "import", "--title", "x"})
	if _, ok := body["editionOf"]; ok {
		t.Fatalf("editionOf set without the flag: %v", body)
	}
}
