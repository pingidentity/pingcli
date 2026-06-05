# `pingcli davinci variables`
DaVinci Variables

## Synopsis

DaVinci Variables

```
pingcli davinci variables [flags]
```

## Inherited Options

| Flag | Default | Description |
|------|---------|-------------|
| `-C, --config string` | `` | The relative or full path to a custom Ping CLI configuration file. (default $HOME/.pingcli/config.yaml) |
| `-D, --detailed-exitcode` | `` | Enable detailed exit code output. (default false) 0 - pingcli command succeeded with no errors or warnings. 1 - pingcli command failed with errors. 2 - pingcli command succeeded with warnings. |
| `-O, --output-format string` | `` | Specify the console output format. (default text) Options are: json, ndjson, ndjson-wrapped, text. |
| `-P, --profile string` | `` | The name of a configuration profile to use. |
| `--debug` | `` | Enable debug output for error messages, including stack traces and transaction IDs. (default false) |
| `--log-file string` | `` | Write logs to a file at the given path. File logging is disabled when not set. |
| `--log-file-level string` | `` | Set the file log level. Options are: DEBUG, INFO, WARN, ERROR. (default DEBUG) |
| `--log-level string` | `` | Set the console log level. Options are: DEBUG, INFO, WARN, ERROR. (default WARN) |
| `--no-color` | `` | Disable text output in color. (default false) |
| `--query string` | `` | JMESPath expression to filter JSON output. Requires -O json, ndjson, or ndjson-wrapped. Example: --query 'data[?enabled].name' |


## Subcommands

| Command | Description | Reference |
|---------|-------------|----------|
| `pingcli davinci variables apply` | Create or update a DaVinci variable | [`cmd-pingcli-davinci-variables-apply.md`](cmd-pingcli-davinci-variables-apply.md) |
| `pingcli davinci variables create` | Create a new DaVinci variable | [`cmd-pingcli-davinci-variables-create.md`](cmd-pingcli-davinci-variables-create.md) |
| `pingcli davinci variables delete` | Delete a DaVinci variable | [`cmd-pingcli-davinci-variables-delete.md`](cmd-pingcli-davinci-variables-delete.md) |
| `pingcli davinci variables get` | Read a DaVinci variable | [`cmd-pingcli-davinci-variables-get.md`](cmd-pingcli-davinci-variables-get.md) |
| `pingcli davinci variables list` | List DaVinci variables | [`cmd-pingcli-davinci-variables-list.md`](cmd-pingcli-davinci-variables-list.md) |
| `pingcli davinci variables replace` | Update a DaVinci variable | [`cmd-pingcli-davinci-variables-replace.md`](cmd-pingcli-davinci-variables-replace.md) |
| `pingcli davinci variables template` | Generate a DaVinci variable JSON template | [`cmd-pingcli-davinci-variables-template.md`](cmd-pingcli-davinci-variables-template.md) |

## Parent Command

- [`pingcli davinci`](cmd-pingcli-davinci.md) — Administration tools for the PingOne DaVinci universal service.
