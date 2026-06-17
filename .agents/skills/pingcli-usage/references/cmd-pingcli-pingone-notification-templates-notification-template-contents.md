# `pingcli pingone notification-templates notification-template-contents`
Notification Template Contents

## Synopsis

Notification Template Contents

```
pingcli pingone notification-templates notification-template-contents [flags]
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
| `pingcli pingone notification-templates notification-template-contents apply` | Create or update a notification template content | [`cmd-pingcli-pingone-notification-templates-notification-template-contents-apply.md`](cmd-pingcli-pingone-notification-templates-notification-template-contents-apply.md) |
| `pingcli pingone notification-templates notification-template-contents create` | Create a new notification template content | [`cmd-pingcli-pingone-notification-templates-notification-template-contents-create.md`](cmd-pingcli-pingone-notification-templates-notification-template-contents-create.md) |
| `pingcli pingone notification-templates notification-template-contents delete` | Delete a notification template content | [`cmd-pingcli-pingone-notification-templates-notification-template-contents-delete.md`](cmd-pingcli-pingone-notification-templates-notification-template-contents-delete.md) |
| `pingcli pingone notification-templates notification-template-contents get` | Read a specific notification template content | [`cmd-pingcli-pingone-notification-templates-notification-template-contents-get.md`](cmd-pingcli-pingone-notification-templates-notification-template-contents-get.md) |
| `pingcli pingone notification-templates notification-template-contents list` | List all notification template contents | [`cmd-pingcli-pingone-notification-templates-notification-template-contents-list.md`](cmd-pingcli-pingone-notification-templates-notification-template-contents-list.md) |
| `pingcli pingone notification-templates notification-template-contents replace` | Replace a notification template content | [`cmd-pingcli-pingone-notification-templates-notification-template-contents-replace.md`](cmd-pingcli-pingone-notification-templates-notification-template-contents-replace.md) |
| `pingcli pingone notification-templates notification-template-contents template` | Generate a notification template content JSON template | [`cmd-pingcli-pingone-notification-templates-notification-template-contents-template.md`](cmd-pingcli-pingone-notification-templates-notification-template-contents-template.md) |

## Parent Command

- [`pingcli pingone notification-templates`](cmd-pingcli-pingone-notification-templates.md) — Notification Templates
