# Documentation Generation (Configuration Options & Command Reference)

This document explains how to generate all Ping CLI documentation artifacts, how the golden
tests validate output, and the available Makefile targets & direct `go run` equivalents.

## Overview

There are two primary documentation generators:

1. Configuration Options Reference (`tools/generate-options-docs`)
2. Command Reference Pages + Navigation (`tools/generate-command-docs`)

Both tools produce AsciiDoc that is ingested by the documentation portal. Golden tests
run automatically (via the Makefile targets) to ensure formatting changes are intentional.

## Configuration Options Documentation

Generate the configuration options reference (default output path is
`docs/dev-ux-portal-docs/general/cli-configuration-settings-reference.adoc`):

```shell
make generate-options-docs
```

Override the output using the `OUTPUT` variable (the argument you pass to `OUTPUT` is
forwarded directly to the generator):

```shell
make generate-options-docs OUTPUT='-o docs/options.md'
make generate-options-docs OUTPUT='-o docs/options.adoc'
```

When called through the Makefile without `OUTPUT`, AsciiDoc is written to the portal path.
When you invoke the generator directly without `-o`, output is written to stdout.

Direct invocation examples:

```shell
go run ./tools/generate-options-docs -o docs/options.md
go run ./tools/generate-options-docs -o docs/dev-ux-portal-docs/general/cli-configuration-settings-reference.adoc
go run ./tools/generate-options-docs -asciidoc > docs/dev-ux-portal-docs/general/cli-configuration-settings-reference.adoc
```

The AsciiDoc generator orders sections as: General, Service, Export, License, Request (as of September 2025 - new options may change this order).

Data types: If a Data Type cell shows "N/A", it means the option's data type hasn't been mapped in the generator yet. Review and add a mapping in `tools/generate-options-docs/docgen/docgen.go` (see `asciiDocDataType`). This is intentional to surface new or unknown types for review rather than silently defaulting.

## Command Reference Pages & Navigation

Generate a page for every command and subcommand plus a hierarchical navigation file
(`nav.adoc`) suitable for portal ingestion:

```shell
make generate-command-docs
```

The generator writes per-command `.adoc` files and `nav.adoc` into `docs/dev-ux-portal-docs`.
Each page includes AsciiDoc attributes:

```adoc
:created-date:
:revdate:
:resourceid:
```

These values appear immediately under the document title. `nav.adoc` is always regenerated;
manual edits will be overwritten.

Direct invocation:

```shell
go run ./tools/generate-command-docs -o ./docs/dev-ux-portal-docs
```

Override the date used in the page headers (for reproducible builds) with:

```shell
go run ./tools/generate-command-docs -date "March 23, 2025" -o ./docs/dev-ux-portal-docs
```

## Rebuilding All Documentation

Force a clean rebuild (removes `docs/dev-ux-portal-docs` then regenerates both sets):

```shell
make generate-all-docs
```

Sequence executed:

1. Remove existing `docs/dev-ux-portal-docs`
2. Generate configuration options reference
3. Generate all command pages + `nav.adoc`
4. Run golden tests (each generator target runs its own test suite)

Equivalent (manual) direct runs:

```shell
go run ./tools/generate-options-docs -o docs/dev-ux-portal-docs/general/cli-configuration-settings-reference.adoc
go run ./tools/generate-command-docs -o docs/dev-ux-portal-docs
```

## Golden Tests Integration

Golden tests live alongside each generator:

- `tools/generate-options-docs/docgen/docgen_test.go`
- `tools/generate-command-docs/main_test.go`

They compare current output against committed fixtures. Dynamic lines (`:created-date:` and
`:revdate:`) are stripped before comparison to keep goldens stable.

To intentionally update goldens (for formatting changes):

```shell
go test ./tools/generate-options-docs/docgen -run TestOptionsDocGeneration -update
go test ./tools/generate-command-docs -run TestCommandDocGeneration -update
```

Running the Makefile targets automatically executes the associated golden tests:

```shell
make generate-options-docs
make generate-command-docs
make generate-all-docs
```

## Makefile Targets Summary

| Target | Purpose |
|--------|---------|
| `generate-options-docs` | Generate configuration options reference (AsciiDoc by default) + run golden test |
| `generate-command-docs` | Generate per-command pages + navigation + run golden test |
| `generate-all-docs` | Clean and rebuild both sets (runs both golden tests) |

## Troubleshooting

| Issue | Resolution |
|-------|------------|
| Golden test fails after code change | Run with `-update` to refresh fixtures if changes are intentional |
| Navigation missing root command | Ensure `renderNav` includes the root (already implemented) |
| Dates cause diffs | They are normalized in tests; ensure you did not alter attribute names |
| "N/A" appears in Data Type column | The option type is not mapped in the generator. Update `asciiDocDataType` in `tools/generate-options-docs/docgen/docgen.go` (or introduce a new `options.Type` as appropriate). |

## See Also

Main project README: `../README.md`
