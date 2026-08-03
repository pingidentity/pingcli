# `pingcli pingfederate bulk export`
Read the PingFederate bulk configuration export

## Synopsis

Read the PingFederate bulk configuration export

```
pingcli pingfederate bulk export [flags]
```

## Examples

```
# Read the PingFederate bulk configuration export to a file
  pingcli pingfederate bulk export --output-file export.json

  # Read the bulk configuration export, including external resource references
  pingcli pingfederate bulk export --output-file export.json --include-external-resources
```

## Options

| Flag | Default | Description |
|------|---------|-------------|
| `-h, --help` | `` | help for export |
| `-o, --output-file string` | `` | Path to write the bulk configuration export JSON to. Overwrites any existing file. |
| `--include-external-resources` | `` | Whether to include external resource references in the bulk configuration export |


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

- [`pingcli pingfederate bulk`](cmd-pingcli-pingfederate-bulk.md) — Manage PingFederate Bulk resources
