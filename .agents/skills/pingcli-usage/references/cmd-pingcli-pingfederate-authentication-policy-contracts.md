# `pingcli pingfederate authentication-policy-contracts`
PingFederate Authentication Policy Contracts

## Synopsis

PingFederate Authentication Policy Contracts

```
pingcli pingfederate authentication-policy-contracts [flags]
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
| `pingcli pingfederate authentication-policy-contracts apply` | Create or update an authentication policy contract | [`cmd-pingcli-pingfederate-authentication-policy-contracts-apply.md`](cmd-pingcli-pingfederate-authentication-policy-contracts-apply.md) |
| `pingcli pingfederate authentication-policy-contracts create` | Create a new authentication policy contract | [`cmd-pingcli-pingfederate-authentication-policy-contracts-create.md`](cmd-pingcli-pingfederate-authentication-policy-contracts-create.md) |
| `pingcli pingfederate authentication-policy-contracts delete` | Delete an authentication policy contract | [`cmd-pingcli-pingfederate-authentication-policy-contracts-delete.md`](cmd-pingcli-pingfederate-authentication-policy-contracts-delete.md) |
| `pingcli pingfederate authentication-policy-contracts get` | Read a specific authentication policy contract | [`cmd-pingcli-pingfederate-authentication-policy-contracts-get.md`](cmd-pingcli-pingfederate-authentication-policy-contracts-get.md) |
| `pingcli pingfederate authentication-policy-contracts list` | List all authentication policy contracts | [`cmd-pingcli-pingfederate-authentication-policy-contracts-list.md`](cmd-pingcli-pingfederate-authentication-policy-contracts-list.md) |
| `pingcli pingfederate authentication-policy-contracts replace` | Update an authentication policy contract | [`cmd-pingcli-pingfederate-authentication-policy-contracts-replace.md`](cmd-pingcli-pingfederate-authentication-policy-contracts-replace.md) |
| `pingcli pingfederate authentication-policy-contracts template` | Generate an authentication policy contract JSON template | [`cmd-pingcli-pingfederate-authentication-policy-contracts-template.md`](cmd-pingcli-pingfederate-authentication-policy-contracts-template.md) |

## Parent Command

- [`pingcli pingfederate`](cmd-pingcli-pingfederate.md) — Administration tools for PingFederate deployed as software
