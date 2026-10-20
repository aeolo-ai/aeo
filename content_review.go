package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

var contentReviewUUID = regexp.MustCompile(`(?i)^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$`)

// Validate only the new review commands; the existing review workspace remains
// a read. The caller supplies the stable key, and the API owns staff admission.
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
			(name == "--job" && action == "review-result")
		if !allowed {
			return fmt.Errorf("unsupported flag %s", name)
		}
		if seen[name] {
			return fmt.Errorf("provide %s only once", name)
		}
		seen[name] = true
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
	return nil
}

func runContentReviewCommand(args []string, domainID string) {
	if err := validateContentReviewArgs(args[1], args[2:]); err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		os.Exit(1)
	}
	proxyCommand(args, domainID)
}
