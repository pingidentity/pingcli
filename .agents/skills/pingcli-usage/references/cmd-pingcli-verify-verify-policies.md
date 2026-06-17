# `pingcli verify verify-policies`
Verify Policies

## Synopsis

Verify Policies

```
pingcli verify verify-policies [flags]
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
| `pingcli verify verify-policies apply` | Create or update a verify policy | [`cmd-pingcli-verify-verify-policies-apply.md`](cmd-pingcli-verify-verify-policies-apply.md) |
| `pingcli verify verify-policies create` | Create a new verify policy | [`cmd-pingcli-verify-verify-policies-create.md`](cmd-pingcli-verify-verify-policies-create.md) |
| `pingcli verify verify-policies delete` | Delete a verify policy | [`cmd-pingcli-verify-verify-policies-delete.md`](cmd-pingcli-verify-verify-policies-delete.md) |
| `pingcli verify verify-policies get` | Read a specific verify policy | [`cmd-pingcli-verify-verify-policies-get.md`](cmd-pingcli-verify-verify-policies-get.md) |
| `pingcli verify verify-policies list` | List all verify policies | [`cmd-pingcli-verify-verify-policies-list.md`](cmd-pingcli-verify-verify-policies-list.md) |
| `pingcli verify verify-policies replace` | Replace a verify policy | [`cmd-pingcli-verify-verify-policies-replace.md`](cmd-pingcli-verify-verify-policies-replace.md) |
| `pingcli verify verify-policies template` | Generate a verify policy JSON template | [`cmd-pingcli-verify-verify-policies-template.md`](cmd-pingcli-verify-verify-policies-template.md) |

## Parent Command

- [`pingcli verify`](cmd-pingcli-verify.md) — Administration tools for the PingOne Verify universal service.
