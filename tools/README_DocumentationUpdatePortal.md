# Documentation Portal Update Instructions

This document explains how to use the generated documentation artifacts to update the Ping CLI documentation portal. The goal for future use is automation, but for now this is a manual process.

Follow these steps after generating the documentation artifacts as described in
`README_DocumentGeneration.md`.

## Overview

There are three primary document types produced by the generators:

1. nav.adoc (command hierarchy navigation snippet)
2. Command reference pages (one per command/subcommand)
3. Configuration options reference (single page)

The nav.adoc is a code snippet to be inserted into the portal's `nav.adoc` file.
The command reference pages and configuration options reference are full AsciiDoc documents suitable for direct inclusion in the portal and can be simply copied over to the appropriate locations.

## Process

### nav.adoc

Copy the contents of `docs/dev-ux-portal-docs/nav.adoc` and paste it into the
portal's `nav.adoc` file, replacing the content that starts with `* Command Reference`.
This section will need to be replaced each time the command docs are regenerated as new subcommands are added.

The target file location in the portal repo is:`asciidoc/modules/ROOT/nav.adoc`

### Command Reference Pages

Copy the remaining `.adoc` files from `docs/dev-ux-portal-docs/` to the folder
`asciidoc/modules/command_reference\pages` of the documentation portal repository.
You can overwrite existing files and add any new ones.

### Configuration Options Reference

Copy the file `docs/dev-ux-portal-docs/general/cli-configuration-settings-reference.adoc`
to the folder `asciidoc/modules/ROOT/pages/general` of the documentation portal repository, overwriting the existing file of the same name.

From this point, you can follow the process for building, reviewing, and publishing the portal documentation as outlined internally.
