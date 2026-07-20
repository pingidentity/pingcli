# `pingcli pingone alert-channels`
Alert Channels

## Synopsis

Alert Channels

```
pingcli pingone alert-channels [flags]
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
| `pingcli pingone alert-channels apply` | Create or update an alert channel | [`cmd-pingcli-pingone-alert-channels-apply.md`](cmd-pingcli-pingone-alert-channels-apply.md) |
| `pingcli pingone alert-channels create` | Create a new alert channel | [`cmd-pingcli-pingone-alert-channels-create.md`](cmd-pingcli-pingone-alert-channels-create.md) |
| `pingcli pingone alert-channels delete` | Delete an alert channel | [`cmd-pingcli-pingone-alert-channels-delete.md`](cmd-pingcli-pingone-alert-channels-delete.md) |
| `pingcli pingone alert-channels list` | List all alert channels | [`cmd-pingcli-pingone-alert-channels-list.md`](cmd-pingcli-pingone-alert-channels-list.md) |
| `pingcli pingone alert-channels replace` | Update an alert channel | [`cmd-pingcli-pingone-alert-channels-replace.md`](cmd-pingcli-pingone-alert-channels-replace.md) |
| `pingcli pingone alert-channels template` | Generate an alert channel JSON template | [`cmd-pingcli-pingone-alert-channels-template.md`](cmd-pingcli-pingone-alert-channels-template.md) |

## Parent Command

- [`pingcli pingone`](cmd-pingcli-pingone.md) — Administration tools for the PingOne platform.
