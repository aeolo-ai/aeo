package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

var contentReviewUUID = regexp.MustCompile(`(?i)^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$`)
var contentApprovalFingerprint = regexp.MustCompile(`^[0-9a-f]{32}$`)

// Accept global domain flags before the noun or verb as well as after them.
// Keep the proxy's noun/verb first without dropping any flags from validation.
// Other commands retain their existing dispatch behavior.
func normalizeContentApprovalArgs(args []string) []string {
	var command, domainFlags []string
	i := 0
	for i < len(args) && len(command) < 2 {
		if (len(command) == 0 && args[i] == "content") ||
			(len(command) == 1 && (args[i] == "approval" || args[i] == "approve")) {
			command = append(command, args[i])
			i++
			continue
		}
		name, _, inline := strings.Cut(args[i], "=")
		if name != "--domain" && name != "-d" {
			return args
		}
		domainFlags = append(domainFlags, args[i])
		i++
		if !inline && i < len(args) {
			domainFlags = append(domainFlags, args[i])
			i++
		}
	}
	if len(command) != 2 || len(domainFlags) == 0 {
		return args
	}
	return append(append(command, domainFlags...), args[i:]...)
}

// Review and publication approval use the existing command proxy. Validate the
// caller's exact version and explicit consent before credentials or HTTP work;
// the API owns access, approval receipts and status transitions.
func validateContentReviewArgs(action string, args []string) error {
	var ids []string
	seen := make(map[string]bool)
	for i := 0; i < len(args); i++ {
		arg := args[i]
		if !strings.HasPrefix(arg, "-") {
			ids = append(ids, arg)
			continue
		}
		name, value, inline := strings.Cut(arg, "=")
		if name == "-d" {
			name = "--domain"
		}
		allowed := name == "--domain" ||
			(name == "--key" && action == "review-run") ||
			(name == "--job" && action == "review-result") ||
			((name == "--fingerprint" || name == "--confirm") && action == "approve")
		if !allowed {
			return fmt.Errorf("unsupported flag %s", name)
		}
		if seen[name] {
			return fmt.Errorf("provide %s only once", name)
		}
		seen[name] = true
		if name == "--confirm" {
			if inline {
				return fmt.Errorf("--confirm must be a bare flag (no value)")
			}
			continue
		}
		if !inline {
			i++
			if i >= len(args) {
				return fmt.Errorf("%s requires a value", name)
			}
			value = args[i]
		}
		if strings.TrimSpace(value) == "" || strings.HasPrefix(value, "-") {
			return fmt.Errorf("%s requires a value", name)
		}
		if (name == "--key" || name == "--job") && !contentReviewUUID.MatchString(value) {
			return fmt.Errorf("%s must be a UUID", name)
		}
		if name == "--fingerprint" && !contentApprovalFingerprint.MatchString(value) {
			return fmt.Errorf("--fingerprint must be 32 lowercase hex characters")
		}
	}
	if len(ids) != 1 {
		return fmt.Errorf("exactly one content ID is required")
	}
	if !contentReviewUUID.MatchString(ids[0]) {
		return fmt.Errorf("content ID must be a UUID")
	}
	if action == "review-run" && !seen["--key"] {
		return fmt.Errorf("--key is required (UUID; reuse the same key when retrying)")
	}
	if action == "approve" {
		if !seen["--fingerprint"] {
			return fmt.Errorf("--fingerprint is required (read it with content approval <id>)")
		}
		if !seen["--confirm"] {
			return fmt.Errorf("--confirm is required (bare flag; approves this version without publishing)")
		}
	}
	return nil
}

func runContentReviewCommand(args []string, domainID string) {
	if err := validateContentReviewArgs(args[1], args[2:]); err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		os.Exit(1)
	}
	if args[1] == "approve" {
		args = canonicalContentApproveArgs(args)
	}
	proxyCommand(args, domainID)
}

// Only call after validation: duplicate flags must never be collapsed into an
// apparently valid approval. The server treats a token after a bare flag as its
// value, so --confirm must be last, with the article ID before any options.
func canonicalContentApproveArgs(args []string) []string {
	var id, fingerprint, domain string
	for i := 2; i < len(args); i++ {
		name, value, inline := strings.Cut(args[i], "=")
		switch name {
		case "--confirm":
			continue
		case "--fingerprint", "--domain", "-d":
			if !inline {
				i++
				value = args[i]
			}
			if name == "--fingerprint" {
				fingerprint = value
			} else {
				domain = value
			}
		default:
			id = args[i]
		}
	}
	out := []string{"content", "approve", id, "--fingerprint", fingerprint}
	if domain != "" {
		out = append(out, "--domain", domain)
	}
	return append(out, "--confirm")
}
