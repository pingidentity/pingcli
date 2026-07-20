# `pingcli pingone authorize api-servers operations`
API Server Operations

## Synopsis

API Server Operations

```
pingcli pingone authorize api-servers operations [flags]
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
| `pingcli pingone authorize api-servers operations apply` | Create or update an API server operation | [`cmd-pingcli-pingone-authorize-api-servers-operations-apply.md`](cmd-pingcli-pingone-authorize-api-servers-operations-apply.md) |
| `pingcli pingone authorize api-servers operations create` | Create a new API server operation | [`cmd-pingcli-pingone-authorize-api-servers-operations-create.md`](cmd-pingcli-pingone-authorize-api-servers-operations-create.md) |
| `pingcli pingone authorize api-servers operations delete` | Delete an API server operation | [`cmd-pingcli-pingone-authorize-api-servers-operations-delete.md`](cmd-pingcli-pingone-authorize-api-servers-operations-delete.md) |
| `pingcli pingone authorize api-servers operations get` | Read a specific API server operation | [`cmd-pingcli-pingone-authorize-api-servers-operations-get.md`](cmd-pingcli-pingone-authorize-api-servers-operations-get.md) |
| `pingcli pingone authorize api-servers operations list` | List all API server operations | [`cmd-pingcli-pingone-authorize-api-servers-operations-list.md`](cmd-pingcli-pingone-authorize-api-servers-operations-list.md) |
| `pingcli pingone authorize api-servers operations replace` | Update an API server operation | [`cmd-pingcli-pingone-authorize-api-servers-operations-replace.md`](cmd-pingcli-pingone-authorize-api-servers-operations-replace.md) |
| `pingcli pingone authorize api-servers operations template` | Generate an API server operation JSON template | [`cmd-pingcli-pingone-authorize-api-servers-operations-template.md`](cmd-pingcli-pingone-authorize-api-servers-operations-template.md) |

## Parent Command

- [`pingcli pingone authorize api-servers`](cmd-pingcli-pingone-authorize-api-servers.md) — API Servers
