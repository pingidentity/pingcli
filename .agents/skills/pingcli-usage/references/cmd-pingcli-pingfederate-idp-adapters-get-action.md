# `pingcli pingfederate idp adapters get-action`
Get an IDP adapter action

## Synopsis

Get a specific action for a PingFederate IDP adapter instance

```
pingcli pingfederate idp adapters get-action [flags]
```

## Examples

```
# Get a specific action for an IDP adapter
  pingcli pingfederate idp adapters get-action --id <id> --action-id <action-id>
```

## Options

| Flag | Default | Description |
|------|---------|-------------|
| `-h, --help` | `` | help for get-action |
| `--action-id string` | `` | The PingFederate IDP adapter action ID |
| `--id string` | `` | The PingFederate IDP Adapter ID |


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

- [`pingcli pingfederate idp adapters`](cmd-pingcli-pingfederate-idp-adapters.md) — PingFederate IDP Adapters
