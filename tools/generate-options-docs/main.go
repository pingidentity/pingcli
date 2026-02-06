// Copyright © 2026 Ping Identity Corporation

package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/pingidentity/pingcli/tools/docutil"
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
				if prevCreated := docutil.ExtractDateLine(string(raw), ":created-date:"); prevCreated != "" {
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
		if docutil.NormalizeForCompare(string(oldRaw)) == docutil.NormalizeForCompare(content) {
			// No meaningful change; avoid updating revision date line.
			return
		}
	}

	if err := os.WriteFile(*outFile, []byte(content), 0o600); err != nil {
		fmt.Fprintf(os.Stderr, "failed to write file: %v\n", err)
		os.Exit(1)
	}
}
