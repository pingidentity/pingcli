# `pingcli pingone credentials credential-types`
Credential Types

## Synopsis

Credential Types

```
pingcli pingone credentials credential-types [flags]
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
| `pingcli pingone credentials credential-types apply` | Create or update a credential type | [`cmd-pingcli-pingone-credentials-credential-types-apply.md`](cmd-pingcli-pingone-credentials-credential-types-apply.md) |
| `pingcli pingone credentials credential-types create` | Create a new credential type | [`cmd-pingcli-pingone-credentials-credential-types-create.md`](cmd-pingcli-pingone-credentials-credential-types-create.md) |
| `pingcli pingone credentials credential-types delete` | Delete a credential type | [`cmd-pingcli-pingone-credentials-credential-types-delete.md`](cmd-pingcli-pingone-credentials-credential-types-delete.md) |
| `pingcli pingone credentials credential-types get` | Read a specific credential type | [`cmd-pingcli-pingone-credentials-credential-types-get.md`](cmd-pingcli-pingone-credentials-credential-types-get.md) |
| `pingcli pingone credentials credential-types list` | List all credential types | [`cmd-pingcli-pingone-credentials-credential-types-list.md`](cmd-pingcli-pingone-credentials-credential-types-list.md) |
| `pingcli pingone credentials credential-types replace` | Replace a credential type | [`cmd-pingcli-pingone-credentials-credential-types-replace.md`](cmd-pingcli-pingone-credentials-credential-types-replace.md) |
| `pingcli pingone credentials credential-types template` | Generate a credential type JSON template | [`cmd-pingcli-pingone-credentials-credential-types-template.md`](cmd-pingcli-pingone-credentials-credential-types-template.md) |

## Parent Command

- [`pingcli pingone credentials`](cmd-pingcli-pingone-credentials.md) — Administration tools for the PingOne Credentials universal service.
