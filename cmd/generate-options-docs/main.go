package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/pingidentity/pingcli/internal/configuration/options/docgen"
)

// A tiny standalone tool (invoked via `make generate-options-docs`) to output
// markdown documentation for all configuration options.
//
// Usage:
//
//	go run ./cmd/generate-options-docs > options.md
//
// or via Makefile target:
//
//	make generate-options-docs (writes to stdout unless -o is specified)
func main() {
	outFile := flag.String("o", "", "If set, write markdown output to the provided file path instead of stdout")
	flag.Parse()

	md := docgen.Markdown()

	if *outFile == "" {
		fmt.Print(md)
		return
	}

	if err := os.WriteFile(*outFile, []byte(md), 0o644); err != nil {
		fmt.Fprintf(os.Stderr, "failed to write file: %v\n", err)
		os.Exit(1)
	}
}
