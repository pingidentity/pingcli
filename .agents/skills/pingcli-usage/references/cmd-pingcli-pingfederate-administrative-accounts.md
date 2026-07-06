# `pingcli pingfederate administrative-accounts`
PingFederate Administrative Accounts

## Synopsis

PingFederate Administrative Accounts

```
pingcli pingfederate administrative-accounts [flags]
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
| `pingcli pingfederate administrative-accounts apply` | Create or update an administrative account | [`cmd-pingcli-pingfederate-administrative-accounts-apply.md`](cmd-pingcli-pingfederate-administrative-accounts-apply.md) |
| `pingcli pingfederate administrative-accounts create` | Create a new administrative account | [`cmd-pingcli-pingfederate-administrative-accounts-create.md`](cmd-pingcli-pingfederate-administrative-accounts-create.md) |
| `pingcli pingfederate administrative-accounts delete` | Delete an administrative account | [`cmd-pingcli-pingfederate-administrative-accounts-delete.md`](cmd-pingcli-pingfederate-administrative-accounts-delete.md) |
| `pingcli pingfederate administrative-accounts get` | Read a specific administrative account | [`cmd-pingcli-pingfederate-administrative-accounts-get.md`](cmd-pingcli-pingfederate-administrative-accounts-get.md) |
| `pingcli pingfederate administrative-accounts list` | List all administrative accounts | [`cmd-pingcli-pingfederate-administrative-accounts-list.md`](cmd-pingcli-pingfederate-administrative-accounts-list.md) |
| `pingcli pingfederate administrative-accounts replace` | Update an administrative account | [`cmd-pingcli-pingfederate-administrative-accounts-replace.md`](cmd-pingcli-pingfederate-administrative-accounts-replace.md) |
| `pingcli pingfederate administrative-accounts reset-password` | Reset a PingFederate administrative account password | [`cmd-pingcli-pingfederate-administrative-accounts-reset-password.md`](cmd-pingcli-pingfederate-administrative-accounts-reset-password.md) |
| `pingcli pingfederate administrative-accounts template` | Generate an administrative account JSON template | [`cmd-pingcli-pingfederate-administrative-accounts-template.md`](cmd-pingcli-pingfederate-administrative-accounts-template.md) |

## Parent Command

- [`pingcli pingfederate`](cmd-pingcli-pingfederate.md) — Administration tools for PingFederate deployed as software
