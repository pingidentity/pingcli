# `pingcli pingfederate notification-publishers`
PingFederate notification publishers

## Synopsis

PingFederate notification publisher plugin instances deliver email, SMS, or other notifications for events such as password resets, account provisioning, and certificate expiration.

```
pingcli pingfederate notification-publishers [flags]
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
| `pingcli pingfederate notification-publishers apply` | Create or update a notification publisher | [`cmd-pingcli-pingfederate-notification-publishers-apply.md`](cmd-pingcli-pingfederate-notification-publishers-apply.md) |
| `pingcli pingfederate notification-publishers create` | Create a new notification publisher | [`cmd-pingcli-pingfederate-notification-publishers-create.md`](cmd-pingcli-pingfederate-notification-publishers-create.md) |
| `pingcli pingfederate notification-publishers delete` | Delete a notification publisher | [`cmd-pingcli-pingfederate-notification-publishers-delete.md`](cmd-pingcli-pingfederate-notification-publishers-delete.md) |
| `pingcli pingfederate notification-publishers descriptors` | PingFederate Notification Publisher Descriptors | [`cmd-pingcli-pingfederate-notification-publishers-descriptors.md`](cmd-pingcli-pingfederate-notification-publishers-descriptors.md) |
| `pingcli pingfederate notification-publishers get` | Read a specific notification publisher | [`cmd-pingcli-pingfederate-notification-publishers-get.md`](cmd-pingcli-pingfederate-notification-publishers-get.md) |
| `pingcli pingfederate notification-publishers get-action` | Get a notification publisher action | [`cmd-pingcli-pingfederate-notification-publishers-get-action.md`](cmd-pingcli-pingfederate-notification-publishers-get-action.md) |
| `pingcli pingfederate notification-publishers invoke-action` | Invoke a notification publisher action | [`cmd-pingcli-pingfederate-notification-publishers-invoke-action.md`](cmd-pingcli-pingfederate-notification-publishers-invoke-action.md) |
| `pingcli pingfederate notification-publishers list` | List all notification publishers | [`cmd-pingcli-pingfederate-notification-publishers-list.md`](cmd-pingcli-pingfederate-notification-publishers-list.md) |
| `pingcli pingfederate notification-publishers list-actions` | List notification publisher actions | [`cmd-pingcli-pingfederate-notification-publishers-list-actions.md`](cmd-pingcli-pingfederate-notification-publishers-list-actions.md) |
| `pingcli pingfederate notification-publishers replace` | Update a notification publisher | [`cmd-pingcli-pingfederate-notification-publishers-replace.md`](cmd-pingcli-pingfederate-notification-publishers-replace.md) |
| `pingcli pingfederate notification-publishers settings` | PingFederate Notification Publishers Settings | [`cmd-pingcli-pingfederate-notification-publishers-settings.md`](cmd-pingcli-pingfederate-notification-publishers-settings.md) |
| `pingcli pingfederate notification-publishers template` | Generate a notification publisher JSON template | [`cmd-pingcli-pingfederate-notification-publishers-template.md`](cmd-pingcli-pingfederate-notification-publishers-template.md) |

## Parent Command

- [`pingcli pingfederate`](cmd-pingcli-pingfederate.md) — Administration tools for PingFederate deployed as software
