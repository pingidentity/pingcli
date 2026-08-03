# `pingcli pingone davinci flows versions`
DaVinci Flow Versions

## Synopsis

DaVinci Flow Versions

```
pingcli pingone davinci flows versions [flags]
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
| `pingcli pingone davinci flows versions delete` | Delete a DaVinci flow version | [`cmd-pingcli-pingone-davinci-flows-versions-delete.md`](cmd-pingcli-pingone-davinci-flows-versions-delete.md) |
| `pingcli pingone davinci flows versions details` | Get DaVinci flow version details | [`cmd-pingcli-pingone-davinci-flows-versions-details.md`](cmd-pingcli-pingone-davinci-flows-versions-details.md) |
| `pingcli pingone davinci flows versions get` | Read a DaVinci flow version | [`cmd-pingcli-pingone-davinci-flows-versions-get.md`](cmd-pingcli-pingone-davinci-flows-versions-get.md) |
| `pingcli pingone davinci flows versions list` | List DaVinci flow versions | [`cmd-pingcli-pingone-davinci-flows-versions-list.md`](cmd-pingcli-pingone-davinci-flows-versions-list.md) |
| `pingcli pingone davinci flows versions set-alias` | Set the alias for a DaVinci flow version | [`cmd-pingcli-pingone-davinci-flows-versions-set-alias.md`](cmd-pingcli-pingone-davinci-flows-versions-set-alias.md) |

## Parent Command

- [`pingcli pingone davinci flows`](cmd-pingcli-pingone-davinci-flows.md) — DaVinci Flows
