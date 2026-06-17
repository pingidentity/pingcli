# `pingcli pingone password-policies`
Password Policies

## Synopsis

Password Policies

```
pingcli pingone password-policies [flags]
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
| `pingcli pingone password-policies apply` | Create or update a password policy | [`cmd-pingcli-pingone-password-policies-apply.md`](cmd-pingcli-pingone-password-policies-apply.md) |
| `pingcli pingone password-policies create` | Create a new password policy | [`cmd-pingcli-pingone-password-policies-create.md`](cmd-pingcli-pingone-password-policies-create.md) |
| `pingcli pingone password-policies delete` | Delete a password policy | [`cmd-pingcli-pingone-password-policies-delete.md`](cmd-pingcli-pingone-password-policies-delete.md) |
| `pingcli pingone password-policies get` | Read a specific password policy | [`cmd-pingcli-pingone-password-policies-get.md`](cmd-pingcli-pingone-password-policies-get.md) |
| `pingcli pingone password-policies list` | List all password policies | [`cmd-pingcli-pingone-password-policies-list.md`](cmd-pingcli-pingone-password-policies-list.md) |
| `pingcli pingone password-policies replace` | Update a password policy | [`cmd-pingcli-pingone-password-policies-replace.md`](cmd-pingcli-pingone-password-policies-replace.md) |
| `pingcli pingone password-policies template` | Generate a password policy JSON template | [`cmd-pingcli-pingone-password-policies-template.md`](cmd-pingcli-pingone-password-policies-template.md) |

## Parent Command

- [`pingcli pingone`](cmd-pingcli-pingone.md) — Administration tools for the PingOne platform.
