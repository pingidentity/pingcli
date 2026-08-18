# `pingcli pingfederate oauth out-of-band-auth-plugins`
PingFederate OAuth out-of-band authenticator plugin instances

## Synopsis

PingFederate OAuth out-of-band authenticator plugin instances deliver authentication challenges to end users outside the browser, such as via a mobile push notification.

```
pingcli pingfederate oauth out-of-band-auth-plugins [flags]
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
| `pingcli pingfederate oauth out-of-band-auth-plugins apply` | Create or update an out-of-band authenticator plugin instance | [`cmd-pingcli-pingfederate-oauth-out-of-band-auth-plugins-apply.md`](cmd-pingcli-pingfederate-oauth-out-of-band-auth-plugins-apply.md) |
| `pingcli pingfederate oauth out-of-band-auth-plugins create` | Create a new out-of-band authenticator plugin instance | [`cmd-pingcli-pingfederate-oauth-out-of-band-auth-plugins-create.md`](cmd-pingcli-pingfederate-oauth-out-of-band-auth-plugins-create.md) |
| `pingcli pingfederate oauth out-of-band-auth-plugins delete` | Delete an out-of-band authenticator plugin instance | [`cmd-pingcli-pingfederate-oauth-out-of-band-auth-plugins-delete.md`](cmd-pingcli-pingfederate-oauth-out-of-band-auth-plugins-delete.md) |
| `pingcli pingfederate oauth out-of-band-auth-plugins get` | Read a specific out-of-band authenticator plugin instance | [`cmd-pingcli-pingfederate-oauth-out-of-band-auth-plugins-get.md`](cmd-pingcli-pingfederate-oauth-out-of-band-auth-plugins-get.md) |
| `pingcli pingfederate oauth out-of-band-auth-plugins list` | List all out-of-band authenticator plugin instances | [`cmd-pingcli-pingfederate-oauth-out-of-band-auth-plugins-list.md`](cmd-pingcli-pingfederate-oauth-out-of-band-auth-plugins-list.md) |
| `pingcli pingfederate oauth out-of-band-auth-plugins replace` | Update an out-of-band authenticator plugin instance | [`cmd-pingcli-pingfederate-oauth-out-of-band-auth-plugins-replace.md`](cmd-pingcli-pingfederate-oauth-out-of-band-auth-plugins-replace.md) |
| `pingcli pingfederate oauth out-of-band-auth-plugins template` | Generate an out-of-band authenticator plugin instance JSON template | [`cmd-pingcli-pingfederate-oauth-out-of-band-auth-plugins-template.md`](cmd-pingcli-pingfederate-oauth-out-of-band-auth-plugins-template.md) |

## Parent Command

- [`pingcli pingfederate oauth`](cmd-pingcli-pingfederate-oauth.md) — Manage PingFederate OAuth resources
