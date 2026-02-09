// Copyright © 2026 Ping Identity Corporation

// Package docutil provides small helpers for documentation generators to share.
package docutil

import (
	"bufio"
	"strings"
)

// NormalizeForCompare returns a version of the input with volatile header lines removed.
// Currently strips :created-date: and :revdate: lines so generators can perform
// stable comparisons and avoid rewriting files when only the date changes.
func NormalizeForCompare(s string) string {
	var b strings.Builder
	scanner := bufio.NewScanner(strings.NewReader(s))
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, ":created-date:") || strings.HasPrefix(line, ":revdate:") {
			continue
		}
		b.WriteString(line)
		b.WriteByte('\n')
	}
	// Trim trailing newline for consistency with previous implementation.
	return strings.TrimSuffix(b.String(), "\n")
}

// ExtractDateLine returns the value of a date-like header line matching the given prefix.
// Example: prefix ":created-date:" matches a line like ":created-date: March 23, 2026" and returns "March 23, 2026".
func ExtractDateLine(content, prefix string) string {
	for _, line := range strings.Split(content, "\n") {
		if strings.HasPrefix(line, prefix) {
			parts := strings.SplitN(line, ": ", 2)
			if len(parts) == 2 {
				return parts[1]
			}
		}
	}

	return ""
}
