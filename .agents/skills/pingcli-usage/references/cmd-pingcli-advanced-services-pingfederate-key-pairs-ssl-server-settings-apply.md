# `pingcli advanced-services pingfederate key-pairs ssl-server settings apply`
Update SSL server settings

## Synopsis

Idempotently update the PingFederate SSL Server Settings. Apply is an alias for replace on this singleton resource.

```
pingcli advanced-services pingfederate key-pairs ssl-server settings apply [flags]
```

## Examples

```
# Update SSL server settings from a JSON file
  pingcli advanced-services pingfederate key-pairs ssl-server settings apply --from-file settings.json

  # Update SSL server settings from stdin
  pingcli advanced-services pingfederate key-pairs ssl-server settings apply --from-file - < settings.json

  # Update from a JSON file, overriding the runtime server certificate reference
  pingcli advanced-services pingfederate key-pairs ssl-server settings apply --from-file settings.json --runtime-server-cert-ref-id <id>
```

## Options

| Flag | Default | Description |
|------|---------|-------------|
| `-h, --help` | `` | help for apply |
| `-f, --from-file string` | `` | Path to a JSON file containing the request body, or "-" to read from stdin. |
| `--admin-console-cert-ref-id string` | `` | ID of the administrative console SSL certificate key pair |
| `--runtime-server-cert-ref-id string` | `` | ID of the runtime server SSL certificate key pair |


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

- [`pingcli advanced-services pingfederate key-pairs ssl-server settings`](cmd-pingcli-advanced-services-pingfederate-key-pairs-ssl-server-settings.md) — PingFederate SSL Server Settings
