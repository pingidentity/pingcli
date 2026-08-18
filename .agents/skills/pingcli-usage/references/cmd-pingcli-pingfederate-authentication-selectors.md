# `pingcli pingfederate authentication-selectors`
PingFederate Authentication Selectors

## Synopsis

PingFederate authentication selectors provide a plugin capability for PingFederate to evaluate various conditions related to the requests.

```
pingcli pingfederate authentication-selectors [flags]
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
| `pingcli pingfederate authentication-selectors apply` | Create or update an authentication selector | [`cmd-pingcli-pingfederate-authentication-selectors-apply.md`](cmd-pingcli-pingfederate-authentication-selectors-apply.md) |
| `pingcli pingfederate authentication-selectors create` | Create a new authentication selector | [`cmd-pingcli-pingfederate-authentication-selectors-create.md`](cmd-pingcli-pingfederate-authentication-selectors-create.md) |
| `pingcli pingfederate authentication-selectors delete` | Delete an authentication selector | [`cmd-pingcli-pingfederate-authentication-selectors-delete.md`](cmd-pingcli-pingfederate-authentication-selectors-delete.md) |
| `pingcli pingfederate authentication-selectors descriptors` | PingFederate Authentication Selector Descriptors | [`cmd-pingcli-pingfederate-authentication-selectors-descriptors.md`](cmd-pingcli-pingfederate-authentication-selectors-descriptors.md) |
| `pingcli pingfederate authentication-selectors get` | Read a specific authentication selector | [`cmd-pingcli-pingfederate-authentication-selectors-get.md`](cmd-pingcli-pingfederate-authentication-selectors-get.md) |
| `pingcli pingfederate authentication-selectors list` | List all authentication selectors | [`cmd-pingcli-pingfederate-authentication-selectors-list.md`](cmd-pingcli-pingfederate-authentication-selectors-list.md) |
| `pingcli pingfederate authentication-selectors replace` | Update an authentication selector | [`cmd-pingcli-pingfederate-authentication-selectors-replace.md`](cmd-pingcli-pingfederate-authentication-selectors-replace.md) |
| `pingcli pingfederate authentication-selectors template` | Generate an authentication selector JSON template | [`cmd-pingcli-pingfederate-authentication-selectors-template.md`](cmd-pingcli-pingfederate-authentication-selectors-template.md) |

## Parent Command

- [`pingcli pingfederate`](cmd-pingcli-pingfederate.md) — Administration tools for PingFederate deployed as software
