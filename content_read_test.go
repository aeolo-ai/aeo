package main

import "testing"

// A read that pulls the whole article costs the caller thousands of characters
// of input window for a paragraph it will not use. These pin that the narrowing
// flags actually reach the server — a dropped flag is invisible, because the
// response is still a valid (just far larger) article.
func TestBuildContentGetPathNarrowsTheRead(t *testing.T) {
	tests := []struct {
		name string
		args []string
		want string
	}{
		{
			name: "default is the whole body",
			args: []string{"content", "get", "abc"},
			want: "/content/abc",
		},
		{
			name: "head",
			args: []string{"content", "get", "abc", "--head"},
			want: "/content/abc?head=true",
		},
		{
			name: "block index",
			args: []string{"content", "get", "abc", "--blocks"},
			want: "/content/abc?blocks=true",
		},
		{
			name: "single block",
			args: []string{"content", "get", "abc", "--block", "b3"},
			want: "/content/abc?block=b3",
		},
		{
			name: "single block equals form",
			args: []string{"content", "get", "abc", "--block=b12"},
			want: "/content/abc?block=b12",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := buildContentGetPath("abc", tt.args)
			if got != tt.want {
				t.Fatalf("got %q, want %q", got, tt.want)
			}
		})
	}
}

func TestBuildContentGetPathEscapesTheBlockId(t *testing.T) {
	got := buildContentGetPath("abc", []string{"content", "get", "abc", "--block", "b 3&x"})
	want := "/content/abc?block=b+3%26x"
	if got != want {
		t.Fatalf("got %q, want %q", got, want)
	}
}

// The prompt exists to stop an agent rewriting a customer's published words
// without a person seeing the change. It must not fire where nobody can answer:
// a piped or CI invocation would hang forever instead of protecting anything.
func TestConfirmPatchSkipsPromptWhenExplicitlyApproved(t *testing.T) {
	for _, flag := range []string{"--yes", "-y"} {
		t.Run(flag, func(t *testing.T) {
			if !confirmPatch("before", "after", []string{"content", "update", "abc", flag}) {
				t.Fatal("expected an explicit approval to skip the prompt")
			}
		})
	}
}

func TestTruncateForPromptCollapsesAndCaps(t *testing.T) {
	if got := truncateForPrompt("  여러   줄\n한 문장  "); got != "여러 줄 한 문장" {
		t.Fatalf("got %q", got)
	}

	long := make([]byte, 600)
	for i := range long {
		long[i] = 'a'
	}
	got := truncateForPrompt(string(long))
	if len([]rune(got)) != 241 {
		t.Fatalf("expected 240 chars plus an ellipsis, got %d", len([]rune(got)))
	}
}
