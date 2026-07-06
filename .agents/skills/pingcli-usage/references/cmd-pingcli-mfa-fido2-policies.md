# `pingcli mfa fido2-policies`
FIDO2 Policies

## Synopsis

FIDO2 Policies

```
pingcli mfa fido2-policies [flags]
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
| `pingcli mfa fido2-policies apply` | Create or update a FIDO2 policy | [`cmd-pingcli-mfa-fido2-policies-apply.md`](cmd-pingcli-mfa-fido2-policies-apply.md) |
| `pingcli mfa fido2-policies create` | Create a new FIDO2 policy | [`cmd-pingcli-mfa-fido2-policies-create.md`](cmd-pingcli-mfa-fido2-policies-create.md) |
| `pingcli mfa fido2-policies delete` | Delete a FIDO2 policy | [`cmd-pingcli-mfa-fido2-policies-delete.md`](cmd-pingcli-mfa-fido2-policies-delete.md) |
| `pingcli mfa fido2-policies get` | Read a specific FIDO2 policy | [`cmd-pingcli-mfa-fido2-policies-get.md`](cmd-pingcli-mfa-fido2-policies-get.md) |
| `pingcli mfa fido2-policies list` | List all FIDO2 policies | [`cmd-pingcli-mfa-fido2-policies-list.md`](cmd-pingcli-mfa-fido2-policies-list.md) |
| `pingcli mfa fido2-policies replace` | Update a FIDO2 policy | [`cmd-pingcli-mfa-fido2-policies-replace.md`](cmd-pingcli-mfa-fido2-policies-replace.md) |
| `pingcli mfa fido2-policies template` | Generate a FIDO2 policy JSON template | [`cmd-pingcli-mfa-fido2-policies-template.md`](cmd-pingcli-mfa-fido2-policies-template.md) |

## Parent Command

- [`pingcli mfa`](cmd-pingcli-mfa.md) — Administration tools for the PingOne MFA universal service.
