package main

import "testing"

func TestBuildMetricsTrafficPathSupportsDashboardRanges(t *testing.T) {
	tests := []struct {
		name string
		args []string
		want string
	}{
		{name: "default", args: []string{"traffic"}, want: "/metrics/traffic"},
		{name: "thirty days", args: []string{"traffic", "--days", "30"}, want: "/metrics/traffic?days=30"},
		{name: "two months", args: []string{"traffic", "--days", "60"}, want: "/metrics/traffic?days=60"},
		{name: "ninety days", args: []string{"traffic", "--days", "90"}, want: "/metrics/traffic?days=90"},
		{name: "six months", args: []string{"traffic", "--days", "180"}, want: "/metrics/traffic?days=180"},
		{name: "one year equals form", args: []string{"traffic", "--days=365"}, want: "/metrics/traffic?days=365"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := buildMetricsTrafficPath(tt.args)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if got != tt.want {
				t.Fatalf("got %q, want %q", got, tt.want)
			}
		})
	}
}

func TestBuildMetricsTrafficPathRejectsUnsupportedRanges(t *testing.T) {
	for _, value := range []string{"0", "7", "14", "28", "181", "366", "30days", "all"} {
		t.Run(value, func(t *testing.T) {
			if _, err := buildMetricsTrafficPath([]string{"traffic", "--days", value}); err == nil {
				t.Fatal("expected an error")
			}
		})
	}
}

func TestBuildMetricsTrafficPathPassesLimitAndQuery(t *testing.T) {
	tests := []struct {
		name string
		args []string
		want string
	}{
		{name: "limit only", args: []string{"traffic", "--limit", "500"}, want: "/metrics/traffic?limit=500"},
		{name: "query only", args: []string{"traffic", "--query", "울쎄라,ulthera"}, want: "/metrics/traffic?query=%EC%9A%B8%EC%8E%84%EB%9D%BC%2Culthera"},
		{name: "all three", args: []string{"traffic", "--days=365", "--limit=1000", "--query", "ウルセラ"}, want: "/metrics/traffic?days=365&limit=1000&query=%E3%82%A6%E3%83%AB%E3%82%BB%E3%83%A9"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := buildMetricsTrafficPath(tt.args)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if got != tt.want {
				t.Fatalf("got %q, want %q", got, tt.want)
			}
		})
	}
}

func TestBuildMetricsTrafficPathRejectsBadLimitAndQuery(t *testing.T) {
	for _, args := range [][]string{
		{"traffic", "--limit", "0"},
		{"traffic", "--limit", "1001"},
		{"traffic", "--limit", "all"},
		{"traffic", "--query", "  "},
	} {
		if _, err := buildMetricsTrafficPath(args); err == nil {
			t.Fatalf("expected an error for %v", args)
		}
	}
}
