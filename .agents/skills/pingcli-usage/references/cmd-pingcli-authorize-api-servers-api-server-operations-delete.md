# `pingcli authorize api-servers api-server-operations delete`
Delete an API server operation

## Synopsis

Delete an API server operation from a PingOne environment

```
pingcli authorize api-servers api-server-operations delete [flags]
```

## Examples

```
# Delete an API server operation
  pingcli authorize api-servers api-server-operations delete --environment-id <env-id> --api-server-id <server-id> --api-server-operation-id <operation-id>
```

## Options

| Flag | Default | Description |
|------|---------|-------------|
| `-h, --help` | `` | help for delete |
| `-a, --api-server-id string` | `` | The API server ID |
| `-e, --environment-id string` | `` | The PingOne environment ID |
| `-o, --api-server-operation-id string` | `` | The API server operation ID |


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

- [`pingcli authorize api-servers api-server-operations`](cmd-pingcli-authorize-api-servers-api-server-operations.md) — API Server Operations
