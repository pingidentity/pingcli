# `pingcli pingfederate password-credential-validators`
PingFederate Password Credential Validators

## Synopsis

PingFederate Password Credential Validators

```
pingcli pingfederate password-credential-validators [flags]
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
| `pingcli pingfederate password-credential-validators apply` | Create or update a password credential validator | [`cmd-pingcli-pingfederate-password-credential-validators-apply.md`](cmd-pingcli-pingfederate-password-credential-validators-apply.md) |
| `pingcli pingfederate password-credential-validators create` | Create a new password credential validator | [`cmd-pingcli-pingfederate-password-credential-validators-create.md`](cmd-pingcli-pingfederate-password-credential-validators-create.md) |
| `pingcli pingfederate password-credential-validators delete` | Delete a password credential validator | [`cmd-pingcli-pingfederate-password-credential-validators-delete.md`](cmd-pingcli-pingfederate-password-credential-validators-delete.md) |
| `pingcli pingfederate password-credential-validators descriptors` | PingFederate Password Credential Validator Descriptors | [`cmd-pingcli-pingfederate-password-credential-validators-descriptors.md`](cmd-pingcli-pingfederate-password-credential-validators-descriptors.md) |
| `pingcli pingfederate password-credential-validators get` | Read a specific password credential validator | [`cmd-pingcli-pingfederate-password-credential-validators-get.md`](cmd-pingcli-pingfederate-password-credential-validators-get.md) |
| `pingcli pingfederate password-credential-validators list` | List all password credential validators | [`cmd-pingcli-pingfederate-password-credential-validators-list.md`](cmd-pingcli-pingfederate-password-credential-validators-list.md) |
| `pingcli pingfederate password-credential-validators replace` | Update a password credential validator | [`cmd-pingcli-pingfederate-password-credential-validators-replace.md`](cmd-pingcli-pingfederate-password-credential-validators-replace.md) |
| `pingcli pingfederate password-credential-validators template` | Generate a password credential validator JSON template | [`cmd-pingcli-pingfederate-password-credential-validators-template.md`](cmd-pingcli-pingfederate-password-credential-validators-template.md) |

## Parent Command

- [`pingcli pingfederate`](cmd-pingcli-pingfederate.md) — Administration tools for PingFederate deployed as software
