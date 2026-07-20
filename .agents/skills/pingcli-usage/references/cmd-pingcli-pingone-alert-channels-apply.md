# `pingcli pingone alert-channels apply`
Create or update an alert channel

## Synopsis

Idempotently create or update an alert channel looked up by the "alertName" field in the JSON body within the supplied --environment-id. If no alert channel with the given alertName exists it is created; if exactly one exists it is updated; if more than one exists the command fails.

```
pingcli pingone alert-channels apply [flags]
```

## Examples

```
# Create or update an alert channel (body supplies alertName, channelType, and addresses)
  pingcli pingone alert-channels apply --environment-id <env-id> --from-file alert-channel.json

  # Read body from stdin
  pingcli pingone alert-channels apply --environment-id <env-id> --from-file - < alert-channel.json
```

## Options

| Flag | Default | Description |
|------|---------|-------------|
| `-h, --help` | `` | help for apply |
| `-a, --alert-channel-id string` | `` | The alert channel ID |
| `-e, --environment-id string` | `` | The PingOne environment ID |
| `-f, --from-file string` | `` | Path to a JSON file containing the request body, or "-" to read from stdin. |


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

- [`pingcli pingone alert-channels`](cmd-pingcli-pingone-alert-channels.md) — Alert Channels
