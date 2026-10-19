# `pingcli advanced-services pingfederate oauth cimd-policies`
PingFederate OAuth CIMD policies

## Synopsis

PingFederate OAuth CIMD (Client ID Metadata Document) policies

```
pingcli advanced-services pingfederate oauth cimd-policies [flags]
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
| `pingcli advanced-services pingfederate oauth cimd-policies apply` | Create or update a CIMD policy | [`cmd-pingcli-advanced-services-pingfederate-oauth-cimd-policies-apply.md`](cmd-pingcli-advanced-services-pingfederate-oauth-cimd-policies-apply.md) |
| `pingcli advanced-services pingfederate oauth cimd-policies create` | Create a new CIMD policy | [`cmd-pingcli-advanced-services-pingfederate-oauth-cimd-policies-create.md`](cmd-pingcli-advanced-services-pingfederate-oauth-cimd-policies-create.md) |
| `pingcli advanced-services pingfederate oauth cimd-policies delete` | Delete a CIMD policy | [`cmd-pingcli-advanced-services-pingfederate-oauth-cimd-policies-delete.md`](cmd-pingcli-advanced-services-pingfederate-oauth-cimd-policies-delete.md) |
| `pingcli advanced-services pingfederate oauth cimd-policies get` | Read a specific CIMD policy | [`cmd-pingcli-advanced-services-pingfederate-oauth-cimd-policies-get.md`](cmd-pingcli-advanced-services-pingfederate-oauth-cimd-policies-get.md) |
| `pingcli advanced-services pingfederate oauth cimd-policies list` | List all CIMD policies | [`cmd-pingcli-advanced-services-pingfederate-oauth-cimd-policies-list.md`](cmd-pingcli-advanced-services-pingfederate-oauth-cimd-policies-list.md) |
| `pingcli advanced-services pingfederate oauth cimd-policies replace` | Update a CIMD policy | [`cmd-pingcli-advanced-services-pingfederate-oauth-cimd-policies-replace.md`](cmd-pingcli-advanced-services-pingfederate-oauth-cimd-policies-replace.md) |
| `pingcli advanced-services pingfederate oauth cimd-policies template` | Generate a CIMD policy JSON template | [`cmd-pingcli-advanced-services-pingfederate-oauth-cimd-policies-template.md`](cmd-pingcli-advanced-services-pingfederate-oauth-cimd-policies-template.md) |

## Parent Command

- [`pingcli advanced-services pingfederate oauth`](cmd-pingcli-advanced-services-pingfederate-oauth.md) — Manage PingFederate OAuth resources
