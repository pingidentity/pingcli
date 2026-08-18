# `pingcli pingfederate oauth access-token-mappings`
PingFederate OAuth access token mappings

## Synopsis

PingFederate OAuth access token mappings

```
pingcli pingfederate oauth access-token-mappings [flags]
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
| `pingcli pingfederate oauth access-token-mappings create` | Create a new access token mapping | [`cmd-pingcli-pingfederate-oauth-access-token-mappings-create.md`](cmd-pingcli-pingfederate-oauth-access-token-mappings-create.md) |
| `pingcli pingfederate oauth access-token-mappings delete` | Delete an access token mapping | [`cmd-pingcli-pingfederate-oauth-access-token-mappings-delete.md`](cmd-pingcli-pingfederate-oauth-access-token-mappings-delete.md) |
| `pingcli pingfederate oauth access-token-mappings get` | Read a specific access token mapping | [`cmd-pingcli-pingfederate-oauth-access-token-mappings-get.md`](cmd-pingcli-pingfederate-oauth-access-token-mappings-get.md) |
| `pingcli pingfederate oauth access-token-mappings list` | List all access token mappings | [`cmd-pingcli-pingfederate-oauth-access-token-mappings-list.md`](cmd-pingcli-pingfederate-oauth-access-token-mappings-list.md) |
| `pingcli pingfederate oauth access-token-mappings replace` | Update an access token mapping | [`cmd-pingcli-pingfederate-oauth-access-token-mappings-replace.md`](cmd-pingcli-pingfederate-oauth-access-token-mappings-replace.md) |
| `pingcli pingfederate oauth access-token-mappings template` | Generate an access token mapping JSON template | [`cmd-pingcli-pingfederate-oauth-access-token-mappings-template.md`](cmd-pingcli-pingfederate-oauth-access-token-mappings-template.md) |

## Parent Command

- [`pingcli pingfederate oauth`](cmd-pingcli-pingfederate-oauth.md) — Manage PingFederate OAuth resources
