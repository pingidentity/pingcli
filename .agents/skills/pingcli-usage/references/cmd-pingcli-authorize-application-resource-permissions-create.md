# `pingcli authorize application-resource-permissions create`
Create a new application resource permission

## Synopsis

Create a new application resource permission in a PingOne environment

```
pingcli authorize application-resource-permissions create [flags]
```

## Examples

```
# Create a new application resource permission from a JSON file
  pingcli authorize application-resource-permissions create --environment-id <env-id> --application-resource-id <app-resource-id> --from-file permission.json

  # Create a new application resource permission from stdin
  pingcli authorize application-resource-permissions create --environment-id <env-id> --application-resource-id <app-resource-id> --from-file - < permission.json
```

## Options

| Flag | Default | Description |
|------|---------|-------------|
| `-h, --help` | `` | help for create |
| `-a, --application-resource-id string` | `` | The parent application resource ID |
| `-e, --environment-id string` | `` | The PingOne environment ID |
| `-f, --from-file string` | `` | Path to a JSON file containing the request body, or "-" to read from stdin. |


## Inherited Options

| Flag | Default | Description |
|------|---------|-------------|
| `-C, --config string` | `` | The relative or full path to a custom Ping CLI configuration file. (default $HOME/.pingcli/config.yaml) |
| `-D, --detailed-exitcode` | `` | Enable detailed exit code output. (default false) 0 - pingcli command succeeded with no errors or warnings. 1 - pingcli command failed with errors. 2 - pingcli command succeeded with warnings. |
| `-O, --output-format string` | `` | Specify the console output format. (default text) Options are: json, ndjson, ndjson-wrapped, text. |
| `-P, --profile string` | `` | The name of a configuration profile to use. |
| `--debug` | `` | Enable debug output for error messages, including stack traces and transaction IDs. (default false) |
| `--log-file string` | `` | Write logs to a file at the given path. File logging is disabled when not set. |
| `--log-file-level string` | `` | Set the file log level. Options are: DEBUG, INFO, WARN, ERROR. (default DEBUG) |
| `--log-level string` | `` | Set the console log level. Options are: DEBUG, INFO, WARN, ERROR. (default WARN) |
| `--no-color` | `` | Disable text output in color. (default false) |
| `--query string` | `` | JMESPath expression to filter JSON output. Requires -O json, ndjson, or ndjson-wrapped. Example: --query 'data[?enabled].name' |


## Parent Command

- [`pingcli authorize application-resource-permissions`](cmd-pingcli-authorize-application-resource-permissions.md) — Application Resource Permissions
