# `pingcli pingfederate oauth authentication-policy-contract-mappings`
PingFederate OAuth Authentication Policy Contract Mappings

## Synopsis

PingFederate OAuth Authentication Policy Contract to Persistent Grant Mappings

```
pingcli pingfederate oauth authentication-policy-contract-mappings [flags]
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
| `pingcli pingfederate oauth authentication-policy-contract-mappings apply` | Create or update an authentication policy contract mapping | [`cmd-pingcli-pingfederate-oauth-authentication-policy-contract-mappings-apply.md`](cmd-pingcli-pingfederate-oauth-authentication-policy-contract-mappings-apply.md) |
| `pingcli pingfederate oauth authentication-policy-contract-mappings create` | Create a new authentication policy contract mapping | [`cmd-pingcli-pingfederate-oauth-authentication-policy-contract-mappings-create.md`](cmd-pingcli-pingfederate-oauth-authentication-policy-contract-mappings-create.md) |
| `pingcli pingfederate oauth authentication-policy-contract-mappings delete` | Delete an authentication policy contract mapping | [`cmd-pingcli-pingfederate-oauth-authentication-policy-contract-mappings-delete.md`](cmd-pingcli-pingfederate-oauth-authentication-policy-contract-mappings-delete.md) |
| `pingcli pingfederate oauth authentication-policy-contract-mappings get` | Read a specific authentication policy contract mapping | [`cmd-pingcli-pingfederate-oauth-authentication-policy-contract-mappings-get.md`](cmd-pingcli-pingfederate-oauth-authentication-policy-contract-mappings-get.md) |
| `pingcli pingfederate oauth authentication-policy-contract-mappings list` | List all authentication policy contract mappings | [`cmd-pingcli-pingfederate-oauth-authentication-policy-contract-mappings-list.md`](cmd-pingcli-pingfederate-oauth-authentication-policy-contract-mappings-list.md) |
| `pingcli pingfederate oauth authentication-policy-contract-mappings replace` | Update an authentication policy contract mapping | [`cmd-pingcli-pingfederate-oauth-authentication-policy-contract-mappings-replace.md`](cmd-pingcli-pingfederate-oauth-authentication-policy-contract-mappings-replace.md) |
| `pingcli pingfederate oauth authentication-policy-contract-mappings template` | Generate an authentication policy contract mapping JSON template | [`cmd-pingcli-pingfederate-oauth-authentication-policy-contract-mappings-template.md`](cmd-pingcli-pingfederate-oauth-authentication-policy-contract-mappings-template.md) |

## Parent Command

- [`pingcli pingfederate oauth`](cmd-pingcli-pingfederate-oauth.md) — Manage PingFederate OAuth resources
