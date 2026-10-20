package main

import (
	"encoding/json"
	"fmt"
	"os"
	"regexp"
)

// agent ask | approve | result: talk to the Aeolo agent from the terminal.
//
// The server owns these (COMMAND_REGISTRY agent.ask/approve/result, cliOnly):
// the same turn engine as MCP aeo_ask, with its plan gate (Starter and up),
// monthly budget and threads. The CLI only adds one thing: a turn that is
// still working after one request is followed with `agent result` until it
// finishes, so one command prints one finished answer. --no-follow returns
// the turnId at once.

// agentFollowMax bounds the follow loop: each read waits up to ~40s server-side.
const agentFollowMax = 15

var agentWorkingTurn = regexp.MustCompile(`status: working · turnId: (turn_[A-Za-z0-9-]+)`)

// agentWorkingTurnID returns the turnId of a reply that is still working, or "".
func agentWorkingTurnID(out string) string {
	if m := agentWorkingTurn.FindStringSubmatch(out); m != nil {
		return m[1]
	}
	return ""
}

// agentArgv drops the CLI-only --no-follow flag and reports whether to follow.
func agentArgv(args []string) ([]string, bool) {
	argv := make([]string, 0, len(args))
	follow := true
	for _, arg := range args {
		if arg == "--no-follow" {
			follow = false
			continue
		}
		argv = append(argv, arg)
	}
	return argv, follow
}

// followAgentTurn keeps reading a working turn through exec until it finishes,
// fails, or the loop bound is reached (then the last reply, with its turnId,
// is returned so the user can run `aeo agent result` later).
func followAgentTurn(out string, exec func([]string) (string, error), onWait func()) (string, error) {
	for i := 0; i < agentFollowMax; i++ {
		turnID := agentWorkingTurnID(out)
		if turnID == "" {
			return out, nil
		}
		if onWait != nil {
			onWait()
		}
		next, err := exec([]string{"agent", "result", turnID})
		if err != nil {
			return out, err
		}
		out = next
	}
	return out, nil
}

func runAgentConversation(args []string, domainID string) {
	argv, follow := agentArgv(args)
	exec := func(a []string) (string, error) {
		body, err := jsonArgv(a)
		if err != nil {
			return "", err
		}
		return callConnector("/execute", "POST", body, domainID)
	}
	out, err := exec(argv)
	if err == nil && follow {
		out, err = followAgentTurn(out, exec, func() {
			fmt.Fprintln(os.Stderr, "… the Aeolo agent is still working")
		})
	}
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err)
		os.Exit(1)
	}
	fmt.Println(out)
}

// jsonArgv is the /execute body proxyCommand sends: argv only, the domain is
// filled in by callConnectorOnce.
func jsonArgv(argv []string) ([]byte, error) {
	return json.Marshal(map[string]interface{}{"argv": argv})
}
