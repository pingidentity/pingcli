# `pingcli advanced-services pingfederate server-settings system-keys`
PingFederate Server Settings System Keys

## Synopsis

PingFederate Server Settings System Keys

```
pingcli advanced-services pingfederate server-settings system-keys [flags]
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
| `pingcli advanced-services pingfederate server-settings system-keys apply` | Update system keys | [`cmd-pingcli-advanced-services-pingfederate-server-settings-system-keys-apply.md`](cmd-pingcli-advanced-services-pingfederate-server-settings-system-keys-apply.md) |
| `pingcli advanced-services pingfederate server-settings system-keys get` | Read system keys | [`cmd-pingcli-advanced-services-pingfederate-server-settings-system-keys-get.md`](cmd-pingcli-advanced-services-pingfederate-server-settings-system-keys-get.md) |
| `pingcli advanced-services pingfederate server-settings system-keys replace` | Update system keys | [`cmd-pingcli-advanced-services-pingfederate-server-settings-system-keys-replace.md`](cmd-pingcli-advanced-services-pingfederate-server-settings-system-keys-replace.md) |
| `pingcli advanced-services pingfederate server-settings system-keys rotate` | Rotate system keys | [`cmd-pingcli-advanced-services-pingfederate-server-settings-system-keys-rotate.md`](cmd-pingcli-advanced-services-pingfederate-server-settings-system-keys-rotate.md) |
| `pingcli advanced-services pingfederate server-settings system-keys template` | Generate a system keys JSON template | [`cmd-pingcli-advanced-services-pingfederate-server-settings-system-keys-template.md`](cmd-pingcli-advanced-services-pingfederate-server-settings-system-keys-template.md) |

## Parent Command

- [`pingcli advanced-services pingfederate server-settings`](cmd-pingcli-advanced-services-pingfederate-server-settings.md) — PingFederate Server Settings
