# `pingcli advanced-services pingfederate server-settings general-settings apply`
Update General Settings

## Synopsis

Idempotently update the General Settings. Apply is an alias for replace on this singleton resource.

```
pingcli advanced-services pingfederate server-settings general-settings apply [flags]
```

## Examples

```
# Update General Settings from a JSON file
  pingcli advanced-services pingfederate server-settings general-settings apply --from-file general-settings.json

  # Update General Settings from stdin
  pingcli advanced-services pingfederate server-settings general-settings apply --from-file - < general-settings.json

  # Update General Settings from flags, without --from-file
  pingcli advanced-services pingfederate server-settings general-settings apply --disable-automatic-connection-validation --datastore-validation-interval-secs 300
```

## Options

| Flag | Default | Description |
|------|---------|-------------|
| `-h, --help` | `` | help for apply |
| `-f, --from-file string` | `` | Path to a JSON file containing the request body, or "-" to read from stdin. |
| `--datastore-validation-interval-secs int64` | `` | Datastore connection-test result cache duration in seconds |
| `--disable-automatic-connection-validation` | `` | Whether to disable automatic connection validation |
| `--idp-connection-transaction-logging-override string` | `` | Transaction logging override for IdP connections |
| `--request-header-for-correlation-id string` | `` | HTTP request header used to retrieve the correlation ID |
| `--sp-connection-transaction-logging-override string` | `` | Transaction logging override for SP connections |


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

- [`pingcli advanced-services pingfederate server-settings general-settings`](cmd-pingcli-advanced-services-pingfederate-server-settings-general-settings.md) — PingFederate General Settings
