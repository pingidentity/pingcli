# `pingcli pingone davinci applications flow-policies`
DaVinci Flow Policies

## Synopsis

DaVinci Flow Policies

```
pingcli pingone davinci applications flow-policies [flags]
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
| `pingcli pingone davinci applications flow-policies apply` | Create or update a DaVinci flow policy | [`cmd-pingcli-pingone-davinci-applications-flow-policies-apply.md`](cmd-pingcli-pingone-davinci-applications-flow-policies-apply.md) |
| `pingcli pingone davinci applications flow-policies create` | Create a new DaVinci flow policy | [`cmd-pingcli-pingone-davinci-applications-flow-policies-create.md`](cmd-pingcli-pingone-davinci-applications-flow-policies-create.md) |
| `pingcli pingone davinci applications flow-policies delete` | Delete a DaVinci flow policy | [`cmd-pingcli-pingone-davinci-applications-flow-policies-delete.md`](cmd-pingcli-pingone-davinci-applications-flow-policies-delete.md) |
| `pingcli pingone davinci applications flow-policies get` | Read a DaVinci flow policy | [`cmd-pingcli-pingone-davinci-applications-flow-policies-get.md`](cmd-pingcli-pingone-davinci-applications-flow-policies-get.md) |
| `pingcli pingone davinci applications flow-policies list` | List DaVinci flow policies | [`cmd-pingcli-pingone-davinci-applications-flow-policies-list.md`](cmd-pingcli-pingone-davinci-applications-flow-policies-list.md) |
| `pingcli pingone davinci applications flow-policies replace` | Update a DaVinci flow policy | [`cmd-pingcli-pingone-davinci-applications-flow-policies-replace.md`](cmd-pingcli-pingone-davinci-applications-flow-policies-replace.md) |
| `pingcli pingone davinci applications flow-policies template` | Generate a DaVinci flow policy JSON template | [`cmd-pingcli-pingone-davinci-applications-flow-policies-template.md`](cmd-pingcli-pingone-davinci-applications-flow-policies-template.md) |

## Parent Command

- [`pingcli pingone davinci applications`](cmd-pingcli-pingone-davinci-applications.md) — DaVinci Applications
