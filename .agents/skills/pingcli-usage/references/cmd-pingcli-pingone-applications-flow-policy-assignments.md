# `pingcli pingone applications flow-policy-assignments`
Flow Policy Assignments

## Synopsis

Flow Policy Assignments

```
pingcli pingone applications flow-policy-assignments [flags]
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
| `pingcli pingone applications flow-policy-assignments create` | Create a new flow policy assignment | [`cmd-pingcli-pingone-applications-flow-policy-assignments-create.md`](cmd-pingcli-pingone-applications-flow-policy-assignments-create.md) |
| `pingcli pingone applications flow-policy-assignments delete` | Delete a flow policy assignment | [`cmd-pingcli-pingone-applications-flow-policy-assignments-delete.md`](cmd-pingcli-pingone-applications-flow-policy-assignments-delete.md) |
| `pingcli pingone applications flow-policy-assignments get` | Read a specific flow policy assignment | [`cmd-pingcli-pingone-applications-flow-policy-assignments-get.md`](cmd-pingcli-pingone-applications-flow-policy-assignments-get.md) |
| `pingcli pingone applications flow-policy-assignments list` | List all flow policy assignments | [`cmd-pingcli-pingone-applications-flow-policy-assignments-list.md`](cmd-pingcli-pingone-applications-flow-policy-assignments-list.md) |
| `pingcli pingone applications flow-policy-assignments replace` | Update a flow policy assignment | [`cmd-pingcli-pingone-applications-flow-policy-assignments-replace.md`](cmd-pingcli-pingone-applications-flow-policy-assignments-replace.md) |
| `pingcli pingone applications flow-policy-assignments template` | Generate a flow policy assignment JSON template | [`cmd-pingcli-pingone-applications-flow-policy-assignments-template.md`](cmd-pingcli-pingone-applications-flow-policy-assignments-template.md) |

## Parent Command

- [`pingcli pingone applications`](cmd-pingcli-pingone-applications.md) — Applications
