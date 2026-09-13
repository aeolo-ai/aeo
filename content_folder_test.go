package main

import "testing"

func TestWritingDestinationFlags(t *testing.T) {
	for _, command := range []string{"generate", "import"} {
		body := map[string]any{}
		applyWritingDestination(body, []string{"content", command, "--channel", "channel-id", "--folder", "/ingredients"})
		if body["targetChannelId"] != "channel-id" || body["hostedFolderPath"] != "/ingredients" {
			t.Fatalf("destination lost: %#v", body)
		}
	}
	body := map[string]any{}
	applyWritingDestination(body, []string{"content", "generate", "--folder-id=folder-id"})
	if body["hostedFolderId"] != "folder-id" {
		t.Fatalf("ID lost: %#v", body)
	}
	body = map[string]any{}
	applyWritingDestination(body, []string{"content", "generate"})
	if len(body) != 0 {
		t.Fatal("omitted selection must use server default")
	}
}
