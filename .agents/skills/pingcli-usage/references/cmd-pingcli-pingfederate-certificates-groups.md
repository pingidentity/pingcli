# `pingcli pingfederate certificates groups`
PingFederate certificate group certificates

## Synopsis

PingFederate certificates within a named certificate group

```
pingcli pingfederate certificates groups [flags]
```

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


## Subcommands

| Command | Description | Reference |
|---------|-------------|----------|
| `pingcli pingfederate certificates groups create` | Import a new certificate into a certificate group | [`cmd-pingcli-pingfederate-certificates-groups-create.md`](cmd-pingcli-pingfederate-certificates-groups-create.md) |
| `pingcli pingfederate certificates groups delete` | Delete a certificate from a certificate group | [`cmd-pingcli-pingfederate-certificates-groups-delete.md`](cmd-pingcli-pingfederate-certificates-groups-delete.md) |
| `pingcli pingfederate certificates groups get` | Read a specific certificate from a certificate group | [`cmd-pingcli-pingfederate-certificates-groups-get.md`](cmd-pingcli-pingfederate-certificates-groups-get.md) |
| `pingcli pingfederate certificates groups list` | List all certificates in a certificate group | [`cmd-pingcli-pingfederate-certificates-groups-list.md`](cmd-pingcli-pingfederate-certificates-groups-list.md) |
| `pingcli pingfederate certificates groups template` | Generate a certificate group JSON template | [`cmd-pingcli-pingfederate-certificates-groups-template.md`](cmd-pingcli-pingfederate-certificates-groups-template.md) |

## Parent Command

- [`pingcli pingfederate certificates`](cmd-pingcli-pingfederate-certificates.md) — Manage PingFederate Certificates resources
