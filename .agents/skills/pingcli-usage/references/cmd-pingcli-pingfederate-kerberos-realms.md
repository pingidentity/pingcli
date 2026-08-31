# `pingcli pingfederate kerberos realms`
PingFederate Kerberos Realms

## Synopsis

PingFederate Kerberos Realms control how PingFederate connects to an Active Directory/Kerberos Realm to validate Kerberos tickets.

```
pingcli pingfederate kerberos realms [flags]
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
| `pingcli pingfederate kerberos realms apply` | Create or update a Kerberos realm | [`cmd-pingcli-pingfederate-kerberos-realms-apply.md`](cmd-pingcli-pingfederate-kerberos-realms-apply.md) |
| `pingcli pingfederate kerberos realms create` | Create a new Kerberos realm | [`cmd-pingcli-pingfederate-kerberos-realms-create.md`](cmd-pingcli-pingfederate-kerberos-realms-create.md) |
| `pingcli pingfederate kerberos realms delete` | Delete a Kerberos realm | [`cmd-pingcli-pingfederate-kerberos-realms-delete.md`](cmd-pingcli-pingfederate-kerberos-realms-delete.md) |
| `pingcli pingfederate kerberos realms get` | Read a specific Kerberos realm | [`cmd-pingcli-pingfederate-kerberos-realms-get.md`](cmd-pingcli-pingfederate-kerberos-realms-get.md) |
| `pingcli pingfederate kerberos realms list` | List all Kerberos realms | [`cmd-pingcli-pingfederate-kerberos-realms-list.md`](cmd-pingcli-pingfederate-kerberos-realms-list.md) |
| `pingcli pingfederate kerberos realms replace` | Update a Kerberos realm | [`cmd-pingcli-pingfederate-kerberos-realms-replace.md`](cmd-pingcli-pingfederate-kerberos-realms-replace.md) |
| `pingcli pingfederate kerberos realms settings` | PingFederate Kerberos Realms Settings | [`cmd-pingcli-pingfederate-kerberos-realms-settings.md`](cmd-pingcli-pingfederate-kerberos-realms-settings.md) |
| `pingcli pingfederate kerberos realms template` | Generate a Kerberos realm JSON template | [`cmd-pingcli-pingfederate-kerberos-realms-template.md`](cmd-pingcli-pingfederate-kerberos-realms-template.md) |

## Parent Command

- [`pingcli pingfederate kerberos`](cmd-pingcli-pingfederate-kerberos.md) — Manage PingFederate Kerberos resources
