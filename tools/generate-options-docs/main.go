package main

import (
	"flag"
	"fmt"
	"os"

	docgen "github.com/pingidentity/pingcli/tools/generate-options-docs/docgen"
)

// A tiny standalone tool (invoked via `make generate-options-docs`) to output
// documentation for configuration options in either Markdown or AsciiDoc.
func main() {
	outFile := flag.String("o", "", "Write output to file (extension .md or .adoc determines format unless flags override)")
	asAsciiDoc := flag.Bool("asciidoc", false, "Force AsciiDoc output (default: Markdown unless output file has .adoc/.asciidoc)")
	flag.Parse()

	useAscii := docgen.ShouldOutputAsciiDoc(*outFile, *asAsciiDoc)

	var content string
	if useAscii {
		content = docgen.GenerateAsciiDoc()
	} else {
		content = docgen.GenerateMarkdown()
	}

	if *outFile == "" {
		fmt.Print(content)

		return
	}

	if err := os.WriteFile(*outFile, []byte(content), 0o600); err != nil {
		fmt.Fprintf(os.Stderr, "failed to write file: %v\n", err)
		os.Exit(1)
	}
}
