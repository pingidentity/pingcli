# `pingcli pingfederate server-settings audit-log-settings replace`
Update audit log settings

## Synopsis

Update (replace) the PingFederate Audit Log Settings

```
pingcli pingfederate server-settings audit-log-settings replace [flags]
```

## Examples

```
# Update audit log settings from a JSON file
  pingcli pingfederate server-settings audit-log-settings replace --from-file audit-log-settings.json

  # Update audit log settings from stdin
  pingcli pingfederate server-settings audit-log-settings replace --from-file - < audit-log-settings.json

  # Update audit log settings from flags, without --from-file
  pingcli pingfederate server-settings audit-log-settings replace --failure-mode BLOCK --threshold 80 --interval 300
```

## Options

| Flag | Default | Description |
|------|---------|-------------|
| `-h, --help` | `` | help for replace |
| `-f, --from-file string` | `` | Path to a JSON file containing the request body, or "-" to read from stdin. |
| `--emails-to-notify []string` | `` | Email addresses notified when the audit logging failure threshold is hit; repeatable or comma-separated |
| `--failure-mode string` | `` | What happens to transactions when the failure threshold is hit |
| `--interval int64` | `` | Interval in seconds over which the failure rate is calculated |
| `--notification-publisher-id string` | `` | ID of the notification publisher used for audit log notifications |
| `--threshold int64` | `` | Percent of failed auditing attempts within the interval that triggers a failure state |
| `--track-audit-log-failures` | `` | Whether PingFederate tracks audit log failures to enable failure notifications |


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


## Parent Command

- [`pingcli pingfederate server-settings audit-log-settings`](cmd-pingcli-pingfederate-server-settings-audit-log-settings.md) — PingFederate Audit Log Settings
