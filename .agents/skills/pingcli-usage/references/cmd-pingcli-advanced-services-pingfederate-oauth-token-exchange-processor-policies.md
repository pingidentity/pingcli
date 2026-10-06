# `pingcli advanced-services pingfederate oauth token-exchange processor policies`
PingFederate OAuth 2.0 Token Exchange processor policies

## Synopsis

PingFederate OAuth 2.0 Token Exchange processor policies

```
pingcli advanced-services pingfederate oauth token-exchange processor policies [flags]
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
| `pingcli advanced-services pingfederate oauth token-exchange processor policies apply` | Create or update a token exchange processor policy | [`cmd-pingcli-advanced-services-pingfederate-oauth-token-exchange-processor-policies-apply.md`](cmd-pingcli-advanced-services-pingfederate-oauth-token-exchange-processor-policies-apply.md) |
| `pingcli advanced-services pingfederate oauth token-exchange processor policies create` | Create a new token exchange processor policy | [`cmd-pingcli-advanced-services-pingfederate-oauth-token-exchange-processor-policies-create.md`](cmd-pingcli-advanced-services-pingfederate-oauth-token-exchange-processor-policies-create.md) |
| `pingcli advanced-services pingfederate oauth token-exchange processor policies delete` | Delete a token exchange processor policy | [`cmd-pingcli-advanced-services-pingfederate-oauth-token-exchange-processor-policies-delete.md`](cmd-pingcli-advanced-services-pingfederate-oauth-token-exchange-processor-policies-delete.md) |
| `pingcli advanced-services pingfederate oauth token-exchange processor policies get` | Read a specific token exchange processor policy | [`cmd-pingcli-advanced-services-pingfederate-oauth-token-exchange-processor-policies-get.md`](cmd-pingcli-advanced-services-pingfederate-oauth-token-exchange-processor-policies-get.md) |
| `pingcli advanced-services pingfederate oauth token-exchange processor policies list` | List all token exchange processor policies | [`cmd-pingcli-advanced-services-pingfederate-oauth-token-exchange-processor-policies-list.md`](cmd-pingcli-advanced-services-pingfederate-oauth-token-exchange-processor-policies-list.md) |
| `pingcli advanced-services pingfederate oauth token-exchange processor policies replace` | Update a token exchange processor policy | [`cmd-pingcli-advanced-services-pingfederate-oauth-token-exchange-processor-policies-replace.md`](cmd-pingcli-advanced-services-pingfederate-oauth-token-exchange-processor-policies-replace.md) |
| `pingcli advanced-services pingfederate oauth token-exchange processor policies template` | Generate a token exchange processor policy JSON template | [`cmd-pingcli-advanced-services-pingfederate-oauth-token-exchange-processor-policies-template.md`](cmd-pingcli-advanced-services-pingfederate-oauth-token-exchange-processor-policies-template.md) |

## Parent Command

- [`pingcli advanced-services pingfederate oauth token-exchange processor`](cmd-pingcli-advanced-services-pingfederate-oauth-token-exchange-processor.md) — Manage PingFederate OAuth Token Exchange processor resources
