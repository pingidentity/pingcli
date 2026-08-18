# `pingcli pingone notifications-settings apply`
Update notifications settings

## Synopsis

Idempotently update the notifications settings configuration for a PingOne environment. Apply is an alias for replace on this singleton resource.

```
pingcli pingone notifications-settings apply [flags]
```

## Examples

```
# Update the notifications settings configuration from a JSON file
  pingcli pingone notifications-settings apply --environment-id <env-id> --from-file notifications-settings.json

  # Update the notifications settings configuration from stdin
  pingcli pingone notifications-settings apply --environment-id <env-id> --from-file - < notifications-settings.json

  # Update from a JSON file, overriding the SMS providers fallback chain
  pingcli pingone notifications-settings apply --environment-id <env-id> --from-file notifications-settings.json --sms-providers-fallback-chain PINGONE_TWILIO
```

## Options

| Flag | Default | Description |
|------|---------|-------------|
| `-h, --help` | `` | help for apply |
| `-e, --environment-id string` | `` | The PingOne environment ID |
| `-f, --from-file string` | `` | Path to a JSON file containing the request body, or "-" to read from stdin. |
| `--sms-providers-fallback-chain []string` | `` | Ordered list of phone delivery settings IDs or PINGONE_TWILIO; repeatable or comma-separated |


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

- [`pingcli pingone notifications-settings`](cmd-pingcli-pingone-notifications-settings.md) — Notifications Settings
