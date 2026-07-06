# `pingcli pingfederate oauth clients`
PingFederate OAuth Clients

## Synopsis

PingFederate OAuth Clients

```
pingcli pingfederate oauth clients [flags]
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
| `pingcli pingfederate oauth clients apply` | Create or update an OAuth client | [`cmd-pingcli-pingfederate-oauth-clients-apply.md`](cmd-pingcli-pingfederate-oauth-clients-apply.md) |
| `pingcli pingfederate oauth clients create` | Create a new OAuth client | [`cmd-pingcli-pingfederate-oauth-clients-create.md`](cmd-pingcli-pingfederate-oauth-clients-create.md) |
| `pingcli pingfederate oauth clients delete` | Delete an OAuth client | [`cmd-pingcli-pingfederate-oauth-clients-delete.md`](cmd-pingcli-pingfederate-oauth-clients-delete.md) |
| `pingcli pingfederate oauth clients get` | Read a specific OAuth client | [`cmd-pingcli-pingfederate-oauth-clients-get.md`](cmd-pingcli-pingfederate-oauth-clients-get.md) |
| `pingcli pingfederate oauth clients get-secret` | Get an OAuth client secret | [`cmd-pingcli-pingfederate-oauth-clients-get-secret.md`](cmd-pingcli-pingfederate-oauth-clients-get-secret.md) |
| `pingcli pingfederate oauth clients list` | List all OAuth clients | [`cmd-pingcli-pingfederate-oauth-clients-list.md`](cmd-pingcli-pingfederate-oauth-clients-list.md) |
| `pingcli pingfederate oauth clients replace` | Update an OAuth client | [`cmd-pingcli-pingfederate-oauth-clients-replace.md`](cmd-pingcli-pingfederate-oauth-clients-replace.md) |
| `pingcli pingfederate oauth clients template` | Generate an OAuth client JSON template | [`cmd-pingcli-pingfederate-oauth-clients-template.md`](cmd-pingcli-pingfederate-oauth-clients-template.md) |
| `pingcli pingfederate oauth clients update-secret` | Update an OAuth client secret | [`cmd-pingcli-pingfederate-oauth-clients-update-secret.md`](cmd-pingcli-pingfederate-oauth-clients-update-secret.md) |

## Parent Command

- [`pingcli pingfederate oauth`](cmd-pingcli-pingfederate-oauth.md) — Manage PingFederate OAuth resources
