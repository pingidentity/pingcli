# `pingcli advanced-services pingfederate oauth ciba-server-policy request-policies`
PingFederate OAuth CIBA Server Policy Request Policies

## Synopsis

PingFederate OAuth CIBA Server Policy Request Policies control how CIBA authentication requests are processed.

```
pingcli advanced-services pingfederate oauth ciba-server-policy request-policies [flags]
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
| `pingcli advanced-services pingfederate oauth ciba-server-policy request-policies apply` | Create or update a CIBA request policy | [`cmd-pingcli-advanced-services-pingfederate-oauth-ciba-server-policy-request-policies-apply.md`](cmd-pingcli-advanced-services-pingfederate-oauth-ciba-server-policy-request-policies-apply.md) |
| `pingcli advanced-services pingfederate oauth ciba-server-policy request-policies create` | Create a new CIBA request policy | [`cmd-pingcli-advanced-services-pingfederate-oauth-ciba-server-policy-request-policies-create.md`](cmd-pingcli-advanced-services-pingfederate-oauth-ciba-server-policy-request-policies-create.md) |
| `pingcli advanced-services pingfederate oauth ciba-server-policy request-policies delete` | Delete a CIBA request policy | [`cmd-pingcli-advanced-services-pingfederate-oauth-ciba-server-policy-request-policies-delete.md`](cmd-pingcli-advanced-services-pingfederate-oauth-ciba-server-policy-request-policies-delete.md) |
| `pingcli advanced-services pingfederate oauth ciba-server-policy request-policies get` | Read a specific CIBA request policy | [`cmd-pingcli-advanced-services-pingfederate-oauth-ciba-server-policy-request-policies-get.md`](cmd-pingcli-advanced-services-pingfederate-oauth-ciba-server-policy-request-policies-get.md) |
| `pingcli advanced-services pingfederate oauth ciba-server-policy request-policies list` | List all CIBA request policies | [`cmd-pingcli-advanced-services-pingfederate-oauth-ciba-server-policy-request-policies-list.md`](cmd-pingcli-advanced-services-pingfederate-oauth-ciba-server-policy-request-policies-list.md) |
| `pingcli advanced-services pingfederate oauth ciba-server-policy request-policies replace` | Update a CIBA request policy | [`cmd-pingcli-advanced-services-pingfederate-oauth-ciba-server-policy-request-policies-replace.md`](cmd-pingcli-advanced-services-pingfederate-oauth-ciba-server-policy-request-policies-replace.md) |
| `pingcli advanced-services pingfederate oauth ciba-server-policy request-policies template` | Generate a CIBA request policy JSON template | [`cmd-pingcli-advanced-services-pingfederate-oauth-ciba-server-policy-request-policies-template.md`](cmd-pingcli-advanced-services-pingfederate-oauth-ciba-server-policy-request-policies-template.md) |

## Parent Command

- [`pingcli advanced-services pingfederate oauth ciba-server-policy`](cmd-pingcli-advanced-services-pingfederate-oauth-ciba-server-policy.md) — Manage PingFederate OAuth CIBA Server Policy resources
