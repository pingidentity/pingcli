# `pingcli pingfederate sp idp-connections signing-settings apply`
Update IdP connection signing settings

## Synopsis

Idempotently update the signing settings singleton for a PingFederate SP IdP connection. Apply is an alias for replace on this singleton resource.

```
pingcli pingfederate sp idp-connections signing-settings apply [flags]
```

## Examples

```
# Replace the signing settings from a JSON file
  pingcli pingfederate sp idp-connections signing-settings apply --connection-id <connection-id> --from-file signing-settings.json

  # Replace the signing settings from stdin
  pingcli pingfederate sp idp-connections signing-settings apply --connection-id <connection-id> --from-file - < signing-settings.json
```

## Options

| Flag | Default | Description |
|------|---------|-------------|
| `-h, --help` | `` | help for apply |
| `-f, --from-file string` | `` | Path to a JSON file containing the request body, or "-" to read from stdin. |
| `--connection-id string` | `` | The ID of the parent PingFederate SP IdP connection |


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

- [`pingcli pingfederate sp idp-connections signing-settings`](cmd-pingcli-pingfederate-sp-idp-connections-signing-settings.md) — PingFederate SP IdP connection signing settings
