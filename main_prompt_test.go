package main

import (
	"encoding/json"
	"testing"
)

func TestPromptUpdateBodyUsesSinglePromptField(t *testing.T) {
	body := buildPromptJSON(map[string]string{
		"prompt": "new measured text",
		"stage":  "comparison",
	}, nil)

	var got map[string]any
	if err := json.Unmarshal(body, &got); err != nil {
		t.Fatal(err)
	}
	if got["prompt"] != "new measured text" {
		t.Fatalf("prompt = %v", got["prompt"])
	}
	if _, exists := got["canonical"]; exists {
		t.Fatal("prompt update must not write canonical")
	}
	if _, exists := got["localized_prompt"]; exists {
		t.Fatal("prompt update must not write localized_prompt")
	}
}

func TestPromptBatchNormalizesLegacyAliasesToPrompt(t *testing.T) {
	body, err := buildPromptsBatchJSON(
		`[{"localized_prompt":"legacy localized"},{"canonical":"legacy canonical"}]`,
		map[string]string{"language": "en"},
		nil,
	)
	if err != nil {
		t.Fatal(err)
	}

	var got struct {
		Items []map[string]any `json:"items"`
	}
	if err := json.Unmarshal(body, &got); err != nil {
		t.Fatal(err)
	}
	for i, item := range got.Items {
		if _, ok := item["prompt"]; !ok {
			t.Fatalf("items[%d] has no prompt: %#v", i, item)
		}
		if _, ok := item["canonical"]; ok {
			t.Fatalf("items[%d] leaked canonical: %#v", i, item)
		}
		if _, ok := item["localized_prompt"]; ok {
			t.Fatalf("items[%d] leaked localized_prompt: %#v", i, item)
		}
	}
}
