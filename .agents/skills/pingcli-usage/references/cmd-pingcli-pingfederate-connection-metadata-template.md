# `pingcli pingfederate connection-metadata template`
Generate a connection metadata convert or export JSON template

## Synopsis

Generate a JSON skeleton template for the connection metadata convert or export body. Use --template-type to select convert (default) or export.

```
pingcli pingfederate connection-metadata template [flags]
```

## Examples

```
# Generate a convert template (default) and save to a file
  pingcli pingfederate connection-metadata template > convert-body.json

  # Generate an export template
  pingcli pingfederate connection-metadata template --template-type export > export-body.json

  # Then run the matching action
  pingcli pingfederate connection-metadata convert --from-file convert-body.json
  pingcli pingfederate connection-metadata export --from-file export-body.json
```

## Options

| Flag | Default | Description |
|------|---------|-------------|
| `-h, --help` | `` | help for template |
| `-o, --output-file string` | `` | Write the JSON template to PATH instead of stdout. Overwrites any existing file. |
| `--template-type string` | `` | The variant of the template to generate. One of: convert, export. Defaults to "convert". |


## Inherited Options

| Flag | Default | Description |
|------|---------|-------------|
| `-C, --config string` | `` | The relative or full path to a custom Ping CLI configuration file. (default $HOME/.pingcli/config.yaml) |
| `-D, --detailed-exitcode` | `` | Enable detailed exit code output. (default false) 0 - pingcli command succeeded with no errors or warnings. 1 - pingcli command failed with errors. 2 - pingcli command succeeded with warnings. |
| `-O, --output-format string` | `` | Specify the console output format. (default text) Options are: json, ndjson, ndjson-typed, ndjson-wrapped, text. |
| `-P, --profile string` | `` | The name of a configuration profile to use. |
| `--debug` | `` | Enable debug output for error messages, including stack traces and transaction IDs. (default false) |
| `--log-file string` | `` | Write logs to a file at the given path. File logging is disabled when not set. |
| `--log-file-level string` | `` | Set the file log level. Options are: DEBUG, INFO, WARN, ERROR. (default DEBUG) |
| `--log-level string` | `` | Set the console log level. Options are: DEBUG, INFO, WARN, ERROR. (default WARN) |
| `--no-color` | `` | Disable text output in color. (default false) |
| `--query string` | `` | JMESPath expression to filter JSON output. Requires -O json, ndjson, ndjson-typed, or ndjson-wrapped. Example: --query 'data[?enabled].name' |


## Parent Command

- [`pingcli pingfederate connection-metadata`](cmd-pingcli-pingfederate-connection-metadata.md) — Manage PingFederate connection metadata
