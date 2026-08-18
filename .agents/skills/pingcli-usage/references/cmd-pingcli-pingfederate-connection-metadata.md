# `pingcli pingfederate connection-metadata`
Manage PingFederate connection metadata

## Synopsis

Manage PingFederate connection metadata operations

```
pingcli pingfederate connection-metadata
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
| `pingcli pingfederate connection-metadata convert` | Convert SAML connection metadata into a JSON connection | [`cmd-pingcli-pingfederate-connection-metadata-convert.md`](cmd-pingcli-pingfederate-connection-metadata-convert.md) |
| `pingcli pingfederate connection-metadata export` | Export a connection's SAML metadata | [`cmd-pingcli-pingfederate-connection-metadata-export.md`](cmd-pingcli-pingfederate-connection-metadata-export.md) |
| `pingcli pingfederate connection-metadata template` | Generate a connection metadata convert or export JSON template | [`cmd-pingcli-pingfederate-connection-metadata-template.md`](cmd-pingcli-pingfederate-connection-metadata-template.md) |

## Parent Command

- [`pingcli pingfederate`](cmd-pingcli-pingfederate.md) — Administration tools for PingFederate deployed as software
