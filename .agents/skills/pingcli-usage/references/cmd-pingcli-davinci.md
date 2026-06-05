# `pingcli davinci`
Administration tools for the PingOne DaVinci universal service.

## Synopsis

Administration tools for the PingOne DaVinci universal service.

```
pingcli davinci
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
| `pingcli davinci applications` | DaVinci Applications | [`cmd-pingcli-davinci-applications.md`](cmd-pingcli-davinci-applications.md) |
| `pingcli davinci connector-instances` | DaVinci Connector Instances | [`cmd-pingcli-davinci-connector-instances.md`](cmd-pingcli-davinci-connector-instances.md) |
| `pingcli davinci connectors` | DaVinci Connector Catalog | [`cmd-pingcli-davinci-connectors.md`](cmd-pingcli-davinci-connectors.md) |
| `pingcli davinci flows` | DaVinci Flows | [`cmd-pingcli-davinci-flows.md`](cmd-pingcli-davinci-flows.md) |
| `pingcli davinci variables` | DaVinci Variables | [`cmd-pingcli-davinci-variables.md`](cmd-pingcli-davinci-variables.md) |

## Parent Command

- [`pingcli`](cmd-pingcli.md) — A CLI tool for managing the configuration of Ping Identity products.
