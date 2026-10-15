package main

import (
	"reflect"
	"testing"
)

func TestCarriesOnly(t *testing.T) {
	cases := []struct {
		args []string
		want bool
	}{
		{[]string{"visibility", "check", "run", "--engines", "chatgpt"}, true},
		{[]string{"visibility", "check", "run", "--engines=chatgpt", "-d", "x"}, true},
		{[]string{"visibility", "check", "run", "--location", "JP"}, false},
		{[]string{"visibility", "check", "run", "--location=JP"}, false},
	}
	for _, c := range cases {
		if got := carriesOnly(c.args, "--engines", "--limit"); got != c.want {
			t.Errorf("carriesOnly(%v) = %v, want %v", c.args, got, c.want)
		}
	}
}

func TestWithPromptText(t *testing.T) {
	got := withPromptText([]string{"content", "generate", "--prompt-file", "brief.md", "--locales", "en,ja"}, "the brief")
	want := []string{"content", "generate", "--locales", "en,ja", "--prompt", "the brief"}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("withPromptText = %v, want %v", got, want)
	}
	got = withPromptText([]string{"content", "generate", "--prompt=x", "--article-length", "long"}, "x")
	want = []string{"content", "generate", "--article-length", "long", "--prompt", "x"}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("withPromptText inline = %v, want %v", got, want)
	}
}
