# `pingcli advanced-services pingfederate authentication-policies policy`
Manage an authentication policy tree

## Synopsis

Manage a PingFederate authentication policy tree.

```
pingcli advanced-services pingfederate authentication-policies policy [flags]
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
| `pingcli advanced-services pingfederate authentication-policies policy create` | Create an Authentication Policy | [`cmd-pingcli-advanced-services-pingfederate-authentication-policies-policy-create.md`](cmd-pingcli-advanced-services-pingfederate-authentication-policies-policy-create.md) |
| `pingcli advanced-services pingfederate authentication-policies policy delete` | Delete an Authentication Policy | [`cmd-pingcli-advanced-services-pingfederate-authentication-policies-policy-delete.md`](cmd-pingcli-advanced-services-pingfederate-authentication-policies-policy-delete.md) |
| `pingcli advanced-services pingfederate authentication-policies policy get` | Read Authentication Policy | [`cmd-pingcli-advanced-services-pingfederate-authentication-policies-policy-get.md`](cmd-pingcli-advanced-services-pingfederate-authentication-policies-policy-get.md) |
| `pingcli advanced-services pingfederate authentication-policies policy move` | Move an authentication policy to a location within the policy tree | [`cmd-pingcli-advanced-services-pingfederate-authentication-policies-policy-move.md`](cmd-pingcli-advanced-services-pingfederate-authentication-policies-policy-move.md) |
| `pingcli advanced-services pingfederate authentication-policies policy replace` | Update Authentication Policy | [`cmd-pingcli-advanced-services-pingfederate-authentication-policies-policy-replace.md`](cmd-pingcli-advanced-services-pingfederate-authentication-policies-policy-replace.md) |
| `pingcli advanced-services pingfederate authentication-policies policy template` | Generate an Authentication Policy JSON template | [`cmd-pingcli-advanced-services-pingfederate-authentication-policies-policy-template.md`](cmd-pingcli-advanced-services-pingfederate-authentication-policies-policy-template.md) |

## Parent Command

- [`pingcli advanced-services pingfederate authentication-policies`](cmd-pingcli-advanced-services-pingfederate-authentication-policies.md) — Manage PingFederate Authentication Policies resources
