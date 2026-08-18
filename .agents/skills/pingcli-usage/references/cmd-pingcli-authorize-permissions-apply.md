# `pingcli authorize permissions apply`
Create or update an application resource permission

## Synopsis

Idempotently create or update an application resource permission looked up by the "action" field in the JSON body within the supplied --environment-id and --application-resource-id. If no permission with the given action exists it is created; if exactly one exists it is updated; if more than one exists the command fails.

```
pingcli authorize permissions apply [flags]
```

## Examples

```
# Create or update an application resource permission from flags
  pingcli authorize permissions apply --environment-id <env-id> --application-resource-id <app-resource-id> --action read --description "Read access"

  # Create or update an application resource permission (body supplies action and optional description)
  pingcli authorize permissions apply --environment-id <env-id> --application-resource-id <app-resource-id> --from-file permission.json

  # Read body from stdin
  pingcli authorize permissions apply --environment-id <env-id> --application-resource-id <app-resource-id> --from-file - < permission.json

  # Create or update from flags, without --from-file
  pingcli authorize permissions apply --environment-id <env-id> --application-resource-id <app-resource-id> --action read --description "Read access"
```

## Options

| Flag | Default | Description |
|------|---------|-------------|
| `-h, --help` | `` | help for apply |
| `-a, --application-resource-id string` | `` | The parent application resource ID |
| `-e, --environment-id string` | `` | The PingOne environment ID |
| `-f, --from-file string` | `` | Path to a JSON file containing the request body, or "-" to read from stdin. |
| `--action string` | `` | The action string identifying this permission (e.g. read, write, execute) |
| `--application-resource-permission-id string` | `` | The application resource permission ID |
| `--description string` | `` | The description of the application resource permission |


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

- [`pingcli authorize permissions`](cmd-pingcli-authorize-permissions.md) — Application Resource Permissions
