# `pingcli credentials types`
Credential Types

## Synopsis

Credential Types

```
pingcli credentials types [flags]
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
| `pingcli credentials types apply` | Create or update a credential type | [`cmd-pingcli-credentials-types-apply.md`](cmd-pingcli-credentials-types-apply.md) |
| `pingcli credentials types create` | Create a new credential type | [`cmd-pingcli-credentials-types-create.md`](cmd-pingcli-credentials-types-create.md) |
| `pingcli credentials types delete` | Delete a credential type | [`cmd-pingcli-credentials-types-delete.md`](cmd-pingcli-credentials-types-delete.md) |
| `pingcli credentials types get` | Read a specific credential type | [`cmd-pingcli-credentials-types-get.md`](cmd-pingcli-credentials-types-get.md) |
| `pingcli credentials types issuance-rules` | Credential Issuance Rules | [`cmd-pingcli-credentials-types-issuance-rules.md`](cmd-pingcli-credentials-types-issuance-rules.md) |
| `pingcli credentials types list` | List all credential types | [`cmd-pingcli-credentials-types-list.md`](cmd-pingcli-credentials-types-list.md) |
| `pingcli credentials types replace` | Replace a credential type | [`cmd-pingcli-credentials-types-replace.md`](cmd-pingcli-credentials-types-replace.md) |
| `pingcli credentials types template` | Generate a credential type JSON template | [`cmd-pingcli-credentials-types-template.md`](cmd-pingcli-credentials-types-template.md) |
| `pingcli credentials types versions` | Credential Type Versions | [`cmd-pingcli-credentials-types-versions.md`](cmd-pingcli-credentials-types-versions.md) |

## Parent Command

- [`pingcli credentials`](cmd-pingcli-credentials.md) — Administration tools for the PingOne Credentials universal service.
