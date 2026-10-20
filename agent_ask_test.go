package main

import (
	"errors"
	"reflect"
	"strings"
	"testing"
)

func TestAgentWorkingTurnID(t *testing.T) {
	working := "The Aeolo agent is still working on this.\n\n---\nstatus: working · turnId: turn_ab-12 · threadId: th_1\n"
	if got := agentWorkingTurnID(working); got != "turn_ab-12" {
		t.Fatalf("got %q", got)
	}
	done := "Write the comparison first.\n\n---\nstatus: completed · threadId: th_1"
	if got := agentWorkingTurnID(done); got != "" {
		t.Fatalf("completed reply should have no working turn, got %q", got)
	}
}

func TestAgentArgvDropsNoFollow(t *testing.T) {
	argv, follow := agentArgv([]string{"agent", "ask", "hi", "--no-follow", "--thread", "th_1"})
	if follow {
		t.Fatal("--no-follow should turn following off")
	}
	if want := []string{"agent", "ask", "hi", "--thread", "th_1"}; !reflect.DeepEqual(argv, want) {
		t.Fatalf("argv %v, want %v", argv, want)
	}
	if _, follow := agentArgv([]string{"agent", "ask", "hi"}); !follow {
		t.Fatal("following is the default")
	}
}

func TestFollowAgentTurnReadsUntilDone(t *testing.T) {
	var calls [][]string
	replies := []string{
		"…\n---\nstatus: working · turnId: turn_1 · threadId: th_1",
		"Done.\n---\nstatus: completed · threadId: th_1",
	}
	exec := func(a []string) (string, error) {
		calls = append(calls, a)
		r := replies[0]
		replies = replies[1:]
		return r, nil
	}
	out, err := followAgentTurn("start\n---\nstatus: working · turnId: turn_1 · threadId: th_1", exec, nil)
	if err != nil || !strings.HasPrefix(out, "Done.") {
		t.Fatalf("out %q err %v", out, err)
	}
	if len(calls) != 2 || !reflect.DeepEqual(calls[0], []string{"agent", "result", "turn_1"}) {
		t.Fatalf("calls %v", calls)
	}
}

func TestFollowAgentTurnStopsAtTheBound(t *testing.T) {
	n := 0
	working := "…\n---\nstatus: working · turnId: turn_1 · threadId: th_1"
	out, _ := followAgentTurn(working, func([]string) (string, error) { n++; return working, nil }, nil)
	if n != agentFollowMax || agentWorkingTurnID(out) != "turn_1" {
		t.Fatalf("n=%d out=%q", n, out)
	}
}

func TestFollowAgentTurnKeepsTheLastReplyOnError(t *testing.T) {
	working := "…\n---\nstatus: working · turnId: turn_1 · threadId: th_1"
	out, err := followAgentTurn(working, func([]string) (string, error) { return "", errors.New("network") }, nil)
	if err == nil || out != working {
		t.Fatalf("out %q err %v", out, err)
	}
}
