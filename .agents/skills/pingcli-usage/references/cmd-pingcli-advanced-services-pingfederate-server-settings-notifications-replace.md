# `pingcli advanced-services pingfederate server-settings notifications replace`
Update notification settings

## Synopsis

Update (replace) the PingFederate Notification Settings

```
pingcli advanced-services pingfederate server-settings notifications replace [flags]
```

## Examples

```
# Update notification settings from a JSON file
  pingcli advanced-services pingfederate server-settings notifications replace --from-file settings.json

  # Update notification settings from stdin
  pingcli advanced-services pingfederate server-settings notifications replace --from-file - < settings.json

  # Update using the flag, without --from-file
  pingcli advanced-services pingfederate server-settings notifications replace --notify-admin-user-password-changes
```

## Options

| Flag | Default | Description |
|------|---------|-------------|
| `-h, --help` | `` | help for replace |
| `-f, --from-file string` | `` | Path to a JSON file containing the request body, or "-" to read from stdin. |
| `--account-changes-notification-publisher-ref-id string` | `` | ID of the notification publisher instance used to notify administrators of account changes |
| `--expired-certificate-administrative-console-warning-days int64` | `` | Number of days past the certificate expiry date the administrative console warning ends |
| `--expiring-certificate-administrative-console-warning-days int64` | `` | Number of days prior to the certificate expiry date the administrative console warning starts |
| `--notify-admin-user-password-changes` | `` | Whether admin users are notified through email when their account is changed |


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

- [`pingcli advanced-services pingfederate server-settings notifications`](cmd-pingcli-advanced-services-pingfederate-server-settings-notifications.md) — PingFederate Notification Settings
