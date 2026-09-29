package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

const pagesUsage = `aeo pages <command> --domain <domain> --channel-id <channel>

Active Custom domain members: read and comment. Owners/editors: create and edit.
Private page changes do not publish. All operations use the shared Pages service.

  access | channels | list [--trash true]
  get | elements | inspect <pageId> [--revision N]
  create | preview                         --input-file request.json
  update | edit | editions | review <id>    --input-file request.json
  trash | restore <id>                     --input-file request.json
  builds list | get <jobId> | edition <pageId>
  builds analyze | generate <id> | apply <id> | translate <id>
  research list | get <id> | start          --product-id <id>
  comments list | create | update | agent | apply
                                            --page-id <id> [--thread-id <id>]

Mutation bodies: --input-file <JSON file> or --input-json '<JSON>'.
Use elements to obtain a revision-bound DOM selection before edit/comment create.
Use get to obtain expectedRevision before update/review/trash/restore.
Comment mutations require version; generating a proposal does not apply it.
See the aeo skill references/pages.md for exact request examples.
`

func pageInputArgs(args []string) ([]string, error) {
	out := make([]string, 0, len(args))
	fileSeen, jsonSeen := false, false
	for i := 0; i < len(args); i++ {
		arg := args[i]
		if arg == "--input-json" || strings.HasPrefix(arg, "--input-json=") {
			jsonSeen = true
		}
		if arg == "--input-file" || strings.HasPrefix(arg, "--input-file=") {
			if fileSeen {
				return nil, fmt.Errorf("provide --input-file only once")
			}
			fileSeen = true
			path := strings.TrimPrefix(arg, "--input-file=")
			if arg == "--input-file" {
				i++
				if i >= len(args) {
					return nil, fmt.Errorf("--input-file requires a path")
				}
				path = args[i]
			}
			data, err := os.ReadFile(path)
			if err != nil {
				return nil, err
			}
			if len(data) > 2_000_000 || !json.Valid(data) {
				return nil, fmt.Errorf("input must be valid JSON under 2 MB")
			}
			out = append(out, "--input-json", string(data))
			continue
		}
		out = append(out, arg)
	}
	if fileSeen && jsonSeen {
		return nil, fmt.Errorf("use either --input-file or --input-json")
	}
	return out, nil
}
func runPagesCommand(args []string, domainID string) {
	if wantsHelp(args) {
		fmt.Print(pagesUsage)
		return
	}
	converted, err := pageInputArgs(args)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		os.Exit(1)
	}
	proxyCommand(converted, domainID)
}
