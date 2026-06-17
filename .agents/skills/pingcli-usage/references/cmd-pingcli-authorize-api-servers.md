# `pingcli authorize api-servers`
API Servers

## Synopsis

API Servers

```
pingcli authorize api-servers [flags]
```

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


## Subcommands

| Command | Description | Reference |
|---------|-------------|----------|
| `pingcli authorize api-servers api-server-operations` | API Server Operations | [`cmd-pingcli-authorize-api-servers-api-server-operations.md`](cmd-pingcli-authorize-api-servers-api-server-operations.md) |
| `pingcli authorize api-servers apply` | Create or update an API server | [`cmd-pingcli-authorize-api-servers-apply.md`](cmd-pingcli-authorize-api-servers-apply.md) |
| `pingcli authorize api-servers create` | Create a new API server | [`cmd-pingcli-authorize-api-servers-create.md`](cmd-pingcli-authorize-api-servers-create.md) |
| `pingcli authorize api-servers delete` | Delete an API server | [`cmd-pingcli-authorize-api-servers-delete.md`](cmd-pingcli-authorize-api-servers-delete.md) |
| `pingcli authorize api-servers get` | Read a specific API server | [`cmd-pingcli-authorize-api-servers-get.md`](cmd-pingcli-authorize-api-servers-get.md) |
| `pingcli authorize api-servers list` | List all API servers | [`cmd-pingcli-authorize-api-servers-list.md`](cmd-pingcli-authorize-api-servers-list.md) |
| `pingcli authorize api-servers replace` | Update an API server | [`cmd-pingcli-authorize-api-servers-replace.md`](cmd-pingcli-authorize-api-servers-replace.md) |
| `pingcli authorize api-servers template` | Generate an API server JSON template | [`cmd-pingcli-authorize-api-servers-template.md`](cmd-pingcli-authorize-api-servers-template.md) |

## Parent Command

- [`pingcli authorize`](cmd-pingcli-authorize.md) — Administration tools for the PingOne Authorize universal service.
