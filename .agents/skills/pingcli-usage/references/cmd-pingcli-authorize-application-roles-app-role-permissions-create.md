# `pingcli authorize application-roles app-role-permissions create`
Create a new application role permission

## Synopsis

Create a new application role permission in a PingOne Authorize environment

```
pingcli authorize application-roles app-role-permissions create [flags]
```

## Examples

```
# Create a new application role permission from a JSON file
  pingcli authorize application-roles app-role-permissions create --environment-id <env-id> --application-role-id <role-id> --from-file permission.json

  # Create a new application role permission from stdin
  pingcli authorize application-roles app-role-permissions create --environment-id <env-id> --application-role-id <role-id> --from-file - < permission.json
```

## Options

| Flag | Default | Description |
|------|---------|-------------|
| `-h, --help` | `` | help for create |
| `-e, --environment-id string` | `` | The PingOne environment ID |
| `-f, --from-file string` | `` | Path to a JSON file containing the request body, or "-" to read from stdin. |
| `-r, --application-role-id string` | `` | The application role ID |


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

- [`pingcli authorize application-roles app-role-permissions`](cmd-pingcli-authorize-application-roles-app-role-permissions.md) — Application Role Permissions
