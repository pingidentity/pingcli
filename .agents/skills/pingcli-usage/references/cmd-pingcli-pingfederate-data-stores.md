# `pingcli pingfederate data-stores`
PingFederate Data Stores

## Synopsis

PingFederate Data Stores provide connections to external data sources used for authentication, user lookup, and attribute retrieval.

```
pingcli pingfederate data-stores [flags]
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
| `pingcli pingfederate data-stores apply` | Create or update a data store | [`cmd-pingcli-pingfederate-data-stores-apply.md`](cmd-pingcli-pingfederate-data-stores-apply.md) |
| `pingcli pingfederate data-stores create` | Create a new data store | [`cmd-pingcli-pingfederate-data-stores-create.md`](cmd-pingcli-pingfederate-data-stores-create.md) |
| `pingcli pingfederate data-stores delete` | Delete a data store | [`cmd-pingcli-pingfederate-data-stores-delete.md`](cmd-pingcli-pingfederate-data-stores-delete.md) |
| `pingcli pingfederate data-stores get` | Read a specific data store | [`cmd-pingcli-pingfederate-data-stores-get.md`](cmd-pingcli-pingfederate-data-stores-get.md) |
| `pingcli pingfederate data-stores get-action` | Get a data store action | [`cmd-pingcli-pingfederate-data-stores-get-action.md`](cmd-pingcli-pingfederate-data-stores-get-action.md) |
| `pingcli pingfederate data-stores invoke-action` | Invoke a data store action | [`cmd-pingcli-pingfederate-data-stores-invoke-action.md`](cmd-pingcli-pingfederate-data-stores-invoke-action.md) |
| `pingcli pingfederate data-stores list` | List all data stores | [`cmd-pingcli-pingfederate-data-stores-list.md`](cmd-pingcli-pingfederate-data-stores-list.md) |
| `pingcli pingfederate data-stores list-actions` | List data store actions | [`cmd-pingcli-pingfederate-data-stores-list-actions.md`](cmd-pingcli-pingfederate-data-stores-list-actions.md) |
| `pingcli pingfederate data-stores replace` | Update a data store | [`cmd-pingcli-pingfederate-data-stores-replace.md`](cmd-pingcli-pingfederate-data-stores-replace.md) |
| `pingcli pingfederate data-stores template` | Generate a data store JSON template (LDAP skeleton; adjust type and fields for other subtypes) | [`cmd-pingcli-pingfederate-data-stores-template.md`](cmd-pingcli-pingfederate-data-stores-template.md) |

## Parent Command

- [`pingcli pingfederate`](cmd-pingcli-pingfederate.md) — Administration tools for PingFederate deployed as software
