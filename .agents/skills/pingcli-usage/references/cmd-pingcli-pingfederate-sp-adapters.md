# `pingcli pingfederate sp adapters`
PingFederate SP adapters

## Synopsis

PingFederate SP adapters map user session data into an application-specific token for a service provider.

```
pingcli pingfederate sp adapters [flags]
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
| `pingcli pingfederate sp adapters apply` | Create or update an SP adapter | [`cmd-pingcli-pingfederate-sp-adapters-apply.md`](cmd-pingcli-pingfederate-sp-adapters-apply.md) |
| `pingcli pingfederate sp adapters create` | Create a new SP adapter | [`cmd-pingcli-pingfederate-sp-adapters-create.md`](cmd-pingcli-pingfederate-sp-adapters-create.md) |
| `pingcli pingfederate sp adapters delete` | Delete an SP adapter | [`cmd-pingcli-pingfederate-sp-adapters-delete.md`](cmd-pingcli-pingfederate-sp-adapters-delete.md) |
| `pingcli pingfederate sp adapters get` | Read a specific SP adapter | [`cmd-pingcli-pingfederate-sp-adapters-get.md`](cmd-pingcli-pingfederate-sp-adapters-get.md) |
| `pingcli pingfederate sp adapters list` | List all SP adapters | [`cmd-pingcli-pingfederate-sp-adapters-list.md`](cmd-pingcli-pingfederate-sp-adapters-list.md) |
| `pingcli pingfederate sp adapters replace` | Update an SP adapter | [`cmd-pingcli-pingfederate-sp-adapters-replace.md`](cmd-pingcli-pingfederate-sp-adapters-replace.md) |
| `pingcli pingfederate sp adapters template` | Generate an SP adapter JSON template | [`cmd-pingcli-pingfederate-sp-adapters-template.md`](cmd-pingcli-pingfederate-sp-adapters-template.md) |

## Parent Command

- [`pingcli pingfederate sp`](cmd-pingcli-pingfederate-sp.md) — Manage PingFederate SP resources
