# `pingcli advanced-services pingfederate authentication-policies fragments`
PingFederate Authentication Policy Fragments

## Synopsis

PingFederate Authentication Policy Fragments

```
pingcli advanced-services pingfederate authentication-policies fragments [flags]
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
| `pingcli advanced-services pingfederate authentication-policies fragments apply` | Create or update an authentication policy fragment | [`cmd-pingcli-advanced-services-pingfederate-authentication-policies-fragments-apply.md`](cmd-pingcli-advanced-services-pingfederate-authentication-policies-fragments-apply.md) |
| `pingcli advanced-services pingfederate authentication-policies fragments create` | Create a new authentication policy fragment | [`cmd-pingcli-advanced-services-pingfederate-authentication-policies-fragments-create.md`](cmd-pingcli-advanced-services-pingfederate-authentication-policies-fragments-create.md) |
| `pingcli advanced-services pingfederate authentication-policies fragments delete` | Delete an authentication policy fragment | [`cmd-pingcli-advanced-services-pingfederate-authentication-policies-fragments-delete.md`](cmd-pingcli-advanced-services-pingfederate-authentication-policies-fragments-delete.md) |
| `pingcli advanced-services pingfederate authentication-policies fragments get` | Read a specific authentication policy fragment | [`cmd-pingcli-advanced-services-pingfederate-authentication-policies-fragments-get.md`](cmd-pingcli-advanced-services-pingfederate-authentication-policies-fragments-get.md) |
| `pingcli advanced-services pingfederate authentication-policies fragments list` | List all authentication policy fragments | [`cmd-pingcli-advanced-services-pingfederate-authentication-policies-fragments-list.md`](cmd-pingcli-advanced-services-pingfederate-authentication-policies-fragments-list.md) |
| `pingcli advanced-services pingfederate authentication-policies fragments replace` | Update an authentication policy fragment | [`cmd-pingcli-advanced-services-pingfederate-authentication-policies-fragments-replace.md`](cmd-pingcli-advanced-services-pingfederate-authentication-policies-fragments-replace.md) |
| `pingcli advanced-services pingfederate authentication-policies fragments template` | Generate an authentication policy fragment JSON template | [`cmd-pingcli-advanced-services-pingfederate-authentication-policies-fragments-template.md`](cmd-pingcli-advanced-services-pingfederate-authentication-policies-fragments-template.md) |

## Parent Command

- [`pingcli advanced-services pingfederate authentication-policies`](cmd-pingcli-advanced-services-pingfederate-authentication-policies.md) — Manage PingFederate Authentication Policies resources
