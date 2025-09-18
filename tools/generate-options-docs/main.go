package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"time"

	docgen "github.com/pingidentity/pingcli/tools/generate-options-docs/docgen"
)

// A tiny standalone tool (invoked via `make generate-options-docs`) to output
// documentation for configuration options in either Markdown or AsciiDoc.
func main() {
	outFile := flag.String("o", "", "Write output to file (extension .md or .adoc determines format unless flags override)")
	asAsciiDoc := flag.Bool("asciidoc", false, "Force AsciiDoc output (default: Markdown unless output file has .adoc/.asciidoc)")
	date := flag.String("date", time.Now().Format("January 2, 2006"), "Revision date to use if content changes (created-date preserved if file exists)")
	flag.Parse()

	useAscii := docgen.ShouldOutputAsciiDoc(*outFile, *asAsciiDoc)

	var content string
	if useAscii {
		// Pull existing created-date if file already present so we can preserve it.
		created := *date
		if *outFile != "" {
			if raw, err := os.ReadFile(*outFile); err == nil {
				if prevCreated := extractDateLine(string(raw), ":created-date:"); prevCreated != "" {
					created = prevCreated
				}
			}
		}
		content = docgen.GenerateAsciiDocWithDates(created, *date)
	} else {
		content = docgen.GenerateMarkdown()
	}

	if *outFile == "" {
		fmt.Print(content)
		return
	}

	// Conditional write: only update if non-date content changed.
	if oldRaw, err := os.ReadFile(*outFile); err == nil {
		if normalizeForCompare(string(oldRaw)) == normalizeForCompare(content) {
			// No meaningful change; avoid updating revision date line.
			return
		}
	}

	if err := os.WriteFile(*outFile, []byte(content), 0o600); err != nil {
		fmt.Fprintf(os.Stderr, "failed to write file: %v\n", err)
		os.Exit(1)
	}
}

// normalizeForCompare strips created / revision date lines for deterministic comparisons.
func normalizeForCompare(s string) string {
	out := make([]string, 0, 256)
	for _, line := range strings.Split(s, "\n") {
		if strings.HasPrefix(line, ":created-date:") || strings.HasPrefix(line, ":revdate:") {
			continue
		}
		out = append(out, line)
	}
	return strings.Join(out, "\n")
}

// extractDateLine returns the value of a date line matching the prefix.
func extractDateLine(content, prefix string) string {
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
