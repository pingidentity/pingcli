# `pingcli mfa mfa-device-policies`
MFA Device Policies

## Synopsis

MFA Device Policies

```
pingcli mfa mfa-device-policies [flags]
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
| `pingcli mfa mfa-device-policies apply` | Create or update an MFA device policy | [`cmd-pingcli-mfa-mfa-device-policies-apply.md`](cmd-pingcli-mfa-mfa-device-policies-apply.md) |
| `pingcli mfa mfa-device-policies create` | Create a new MFA device policy | [`cmd-pingcli-mfa-mfa-device-policies-create.md`](cmd-pingcli-mfa-mfa-device-policies-create.md) |
| `pingcli mfa mfa-device-policies delete` | Delete an MFA device policy | [`cmd-pingcli-mfa-mfa-device-policies-delete.md`](cmd-pingcli-mfa-mfa-device-policies-delete.md) |
| `pingcli mfa mfa-device-policies get` | Read a specific MFA device policy | [`cmd-pingcli-mfa-mfa-device-policies-get.md`](cmd-pingcli-mfa-mfa-device-policies-get.md) |
| `pingcli mfa mfa-device-policies list` | List all MFA device policies | [`cmd-pingcli-mfa-mfa-device-policies-list.md`](cmd-pingcli-mfa-mfa-device-policies-list.md) |
| `pingcli mfa mfa-device-policies replace` | Update an MFA device policy | [`cmd-pingcli-mfa-mfa-device-policies-replace.md`](cmd-pingcli-mfa-mfa-device-policies-replace.md) |
| `pingcli mfa mfa-device-policies template` | Generate an MFA device policy JSON template | [`cmd-pingcli-mfa-mfa-device-policies-template.md`](cmd-pingcli-mfa-mfa-device-policies-template.md) |

## Parent Command

- [`pingcli mfa`](cmd-pingcli-mfa.md) — Administration tools for the PingOne MFA universal service.
