# `pingcli advanced-services pingfederate sp adapters`
PingFederate SP adapters

## Synopsis

PingFederate SP adapters map user session data into an application-specific token for a service provider.

```
pingcli advanced-services pingfederate sp adapters [flags]
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
| `pingcli advanced-services pingfederate sp adapters apply` | Create or update an SP adapter | [`cmd-pingcli-advanced-services-pingfederate-sp-adapters-apply.md`](cmd-pingcli-advanced-services-pingfederate-sp-adapters-apply.md) |
| `pingcli advanced-services pingfederate sp adapters create` | Create a new SP adapter | [`cmd-pingcli-advanced-services-pingfederate-sp-adapters-create.md`](cmd-pingcli-advanced-services-pingfederate-sp-adapters-create.md) |
| `pingcli advanced-services pingfederate sp adapters delete` | Delete an SP adapter | [`cmd-pingcli-advanced-services-pingfederate-sp-adapters-delete.md`](cmd-pingcli-advanced-services-pingfederate-sp-adapters-delete.md) |
| `pingcli advanced-services pingfederate sp adapters descriptors` | PingFederate SP Adapter Descriptors | [`cmd-pingcli-advanced-services-pingfederate-sp-adapters-descriptors.md`](cmd-pingcli-advanced-services-pingfederate-sp-adapters-descriptors.md) |
| `pingcli advanced-services pingfederate sp adapters get` | Read a specific SP adapter | [`cmd-pingcli-advanced-services-pingfederate-sp-adapters-get.md`](cmd-pingcli-advanced-services-pingfederate-sp-adapters-get.md) |
| `pingcli advanced-services pingfederate sp adapters get-action` | Get an SP adapter action | [`cmd-pingcli-advanced-services-pingfederate-sp-adapters-get-action.md`](cmd-pingcli-advanced-services-pingfederate-sp-adapters-get-action.md) |
| `pingcli advanced-services pingfederate sp adapters invoke-action` | Invoke an SP adapter action | [`cmd-pingcli-advanced-services-pingfederate-sp-adapters-invoke-action.md`](cmd-pingcli-advanced-services-pingfederate-sp-adapters-invoke-action.md) |
| `pingcli advanced-services pingfederate sp adapters list` | List all SP adapters | [`cmd-pingcli-advanced-services-pingfederate-sp-adapters-list.md`](cmd-pingcli-advanced-services-pingfederate-sp-adapters-list.md) |
| `pingcli advanced-services pingfederate sp adapters list-actions` | List SP adapter actions | [`cmd-pingcli-advanced-services-pingfederate-sp-adapters-list-actions.md`](cmd-pingcli-advanced-services-pingfederate-sp-adapters-list-actions.md) |
| `pingcli advanced-services pingfederate sp adapters replace` | Update an SP adapter | [`cmd-pingcli-advanced-services-pingfederate-sp-adapters-replace.md`](cmd-pingcli-advanced-services-pingfederate-sp-adapters-replace.md) |
| `pingcli advanced-services pingfederate sp adapters template` | Generate an SP adapter JSON template | [`cmd-pingcli-advanced-services-pingfederate-sp-adapters-template.md`](cmd-pingcli-advanced-services-pingfederate-sp-adapters-template.md) |

## Parent Command

- [`pingcli advanced-services pingfederate sp`](cmd-pingcli-advanced-services-pingfederate-sp.md) — Manage PingFederate SP resources
