# `pingcli pingone notification-policies apply`
Create or update a notification policy

## Synopsis

Idempotently create or update a notification policy looked up by the "name" field in the JSON body within the supplied --environment-id. If no notification policy with the given name exists it is created; if exactly one exists it is updated; if more than one exists the command fails.

```
pingcli pingone notification-policies apply [flags]
```

## Examples

```
# Create or update a notification policy (body supplies name and required fields)
  pingcli pingone notification-policies apply --environment-id <env-id> --from-file notification-policy.json

  # Read body from stdin
  pingcli pingone notification-policies apply --environment-id <env-id> --from-file - < notification-policy.json
```

## Options

| Flag | Default | Description |
|------|---------|-------------|
| `-h, --help` | `` | help for apply |
| `-e, --environment-id string` | `` | The PingOne environment ID |
| `-f, --from-file string` | `` | Path to a JSON file containing the request body, or "-" to read from stdin. |
| `-n, --notification-policy-id string` | `` | The notification policy ID |


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

- [`pingcli pingone notification-policies`](cmd-pingcli-pingone-notification-policies.md) — Notification Policies
