# `pingcli pingfederate ping-one-connections environments`
List the PingOne environments for a PingOne connection

## Synopsis

List the PingOne environments associated with a PingFederate PingOne connection

```
pingcli pingfederate ping-one-connections environments [flags]
```

## Examples

```
# List the PingOne environments for a connection
  pingcli pingfederate ping-one-connections environments --id <id>

  # Page through results and filter by name
  pingcli pingfederate ping-one-connections environments --id <id> --page 1 --number-per-page 10 --filter <expr>
```

## Options

| Flag | Default | Description |
|------|---------|-------------|
| `-h, --help` | `` | help for environments |
| `--filter string` | `` | Filter criteria limiting the environments returned. |
| `--id string` | `` | The ID of the PingOne connection |
| `--number-per-page int64` | `` | Number of environments to return per page. |
| `--page int64` | `` | Page number of environments to retrieve. |


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

- [`pingcli pingfederate ping-one-connections`](cmd-pingcli-pingfederate-ping-one-connections.md) — PingFederate PingOne connections
