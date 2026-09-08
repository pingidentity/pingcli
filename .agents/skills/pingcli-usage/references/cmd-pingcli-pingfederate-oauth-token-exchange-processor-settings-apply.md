# `pingcli pingfederate oauth token-exchange processor settings apply`
Update token exchange processor settings

## Synopsis

Idempotently update the PingFederate OAuth 2.0 Token Exchange processor settings. Apply is an alias for replace on this singleton resource.

```
pingcli pingfederate oauth token-exchange processor settings apply [flags]
```

## Examples

```
# Update token exchange processor settings from a JSON file
  pingcli pingfederate oauth token-exchange processor settings apply --from-file settings.json

  # Update token exchange processor settings from stdin
  pingcli pingfederate oauth token-exchange processor settings apply --from-file - < settings.json

  # Update token exchange processor settings using flags only (no --from-file needed)
  pingcli pingfederate oauth token-exchange processor settings apply --default-processor-policy-ref-id <id>
```

## Options

| Flag | Default | Description |
|------|---------|-------------|
| `-h, --help` | `` | help for apply |
| `-f, --from-file string` | `` | Path to a JSON file containing the request body, or "-" to read from stdin. |
| `--default-processor-policy-ref-id string` | `` | ID of the token exchange processor policy to use as the default |


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

- [`pingcli pingfederate oauth token-exchange processor settings`](cmd-pingcli-pingfederate-oauth-token-exchange-processor-settings.md) — PingFederate OAuth 2.0 Token Exchange processor settings
