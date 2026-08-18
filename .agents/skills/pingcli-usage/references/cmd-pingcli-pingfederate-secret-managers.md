# `pingcli pingfederate secret-managers`
PingFederate Secret Managers

## Synopsis

PingFederate Secret Managers provide integrations with external credential vaults for secure secret retrieval.

```
pingcli pingfederate secret-managers [flags]
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
| `pingcli pingfederate secret-managers apply` | Create or update a secret manager | [`cmd-pingcli-pingfederate-secret-managers-apply.md`](cmd-pingcli-pingfederate-secret-managers-apply.md) |
| `pingcli pingfederate secret-managers create` | Create a new secret manager | [`cmd-pingcli-pingfederate-secret-managers-create.md`](cmd-pingcli-pingfederate-secret-managers-create.md) |
| `pingcli pingfederate secret-managers delete` | Delete a secret manager | [`cmd-pingcli-pingfederate-secret-managers-delete.md`](cmd-pingcli-pingfederate-secret-managers-delete.md) |
| `pingcli pingfederate secret-managers descriptors` | PingFederate Secret Manager Descriptors | [`cmd-pingcli-pingfederate-secret-managers-descriptors.md`](cmd-pingcli-pingfederate-secret-managers-descriptors.md) |
| `pingcli pingfederate secret-managers get` | Read a specific secret manager | [`cmd-pingcli-pingfederate-secret-managers-get.md`](cmd-pingcli-pingfederate-secret-managers-get.md) |
| `pingcli pingfederate secret-managers get-action` | Get a secret manager action | [`cmd-pingcli-pingfederate-secret-managers-get-action.md`](cmd-pingcli-pingfederate-secret-managers-get-action.md) |
| `pingcli pingfederate secret-managers invoke-action` | Invoke a secret manager action | [`cmd-pingcli-pingfederate-secret-managers-invoke-action.md`](cmd-pingcli-pingfederate-secret-managers-invoke-action.md) |
| `pingcli pingfederate secret-managers list` | List all secret managers | [`cmd-pingcli-pingfederate-secret-managers-list.md`](cmd-pingcli-pingfederate-secret-managers-list.md) |
| `pingcli pingfederate secret-managers list-actions` | List secret manager actions | [`cmd-pingcli-pingfederate-secret-managers-list-actions.md`](cmd-pingcli-pingfederate-secret-managers-list-actions.md) |
| `pingcli pingfederate secret-managers replace` | Update a secret manager | [`cmd-pingcli-pingfederate-secret-managers-replace.md`](cmd-pingcli-pingfederate-secret-managers-replace.md) |
| `pingcli pingfederate secret-managers template` | Generate a secret manager JSON template | [`cmd-pingcli-pingfederate-secret-managers-template.md`](cmd-pingcli-pingfederate-secret-managers-template.md) |

## Parent Command

- [`pingcli pingfederate`](cmd-pingcli-pingfederate.md) — Administration tools for PingFederate deployed as software
