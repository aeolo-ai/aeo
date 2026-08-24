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
