# `pingcli pingfederate session settings`
PingFederate Session Settings

## Synopsis

PingFederate Session Settings

```
pingcli pingfederate session settings [flags]
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
| `pingcli pingfederate session settings apply` | Update session settings | [`cmd-pingcli-pingfederate-session-settings-apply.md`](cmd-pingcli-pingfederate-session-settings-apply.md) |
| `pingcli pingfederate session settings get` | Read session settings | [`cmd-pingcli-pingfederate-session-settings-get.md`](cmd-pingcli-pingfederate-session-settings-get.md) |
| `pingcli pingfederate session settings replace` | Update session settings | [`cmd-pingcli-pingfederate-session-settings-replace.md`](cmd-pingcli-pingfederate-session-settings-replace.md) |
| `pingcli pingfederate session settings template` | Generate a session settings JSON template | [`cmd-pingcli-pingfederate-session-settings-template.md`](cmd-pingcli-pingfederate-session-settings-template.md) |

## Parent Command

- [`pingcli pingfederate session`](cmd-pingcli-pingfederate-session.md) — Manage PingFederate Session resources
