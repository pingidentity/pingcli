# `pingcli davinci connector-instances`
DaVinci Connector Instances

## Synopsis

DaVinci Connector Instances

```
pingcli davinci connector-instances [flags]
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
| `pingcli davinci connector-instances apply` | Create or update a DaVinci connector instance | [`cmd-pingcli-davinci-connector-instances-apply.md`](cmd-pingcli-davinci-connector-instances-apply.md) |
| `pingcli davinci connector-instances create` | Create a new DaVinci connector instance | [`cmd-pingcli-davinci-connector-instances-create.md`](cmd-pingcli-davinci-connector-instances-create.md) |
| `pingcli davinci connector-instances delete` | Delete a DaVinci connector instance | [`cmd-pingcli-davinci-connector-instances-delete.md`](cmd-pingcli-davinci-connector-instances-delete.md) |
| `pingcli davinci connector-instances get` | Read a DaVinci connector instance | [`cmd-pingcli-davinci-connector-instances-get.md`](cmd-pingcli-davinci-connector-instances-get.md) |
| `pingcli davinci connector-instances list` | List DaVinci connector instances | [`cmd-pingcli-davinci-connector-instances-list.md`](cmd-pingcli-davinci-connector-instances-list.md) |
| `pingcli davinci connector-instances replace` | Update a DaVinci connector instance | [`cmd-pingcli-davinci-connector-instances-replace.md`](cmd-pingcli-davinci-connector-instances-replace.md) |
| `pingcli davinci connector-instances template` | Generate a DaVinci connector instance JSON template | [`cmd-pingcli-davinci-connector-instances-template.md`](cmd-pingcli-davinci-connector-instances-template.md) |

## Parent Command

- [`pingcli davinci`](cmd-pingcli-davinci.md) — Administration tools for the PingOne DaVinci universal service.
