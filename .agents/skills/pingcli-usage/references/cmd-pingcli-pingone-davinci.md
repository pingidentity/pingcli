# `pingcli pingone davinci`
Administration tools for the PingOne DaVinci universal service.

## Synopsis

Administration tools for the PingOne DaVinci universal service.

```
pingcli pingone davinci
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
| `pingcli pingone davinci applications` | DaVinci Applications | [`cmd-pingcli-pingone-davinci-applications.md`](cmd-pingcli-pingone-davinci-applications.md) |
| `pingcli pingone davinci connector-instances` | DaVinci Connector Instances | [`cmd-pingcli-pingone-davinci-connector-instances.md`](cmd-pingcli-pingone-davinci-connector-instances.md) |
| `pingcli pingone davinci connectors` | DaVinci Connector Catalog | [`cmd-pingcli-pingone-davinci-connectors.md`](cmd-pingcli-pingone-davinci-connectors.md) |
| `pingcli pingone davinci flows` | DaVinci Flows | [`cmd-pingcli-pingone-davinci-flows.md`](cmd-pingcli-pingone-davinci-flows.md) |
| `pingcli pingone davinci variables` | DaVinci Variables | [`cmd-pingcli-pingone-davinci-variables.md`](cmd-pingcli-pingone-davinci-variables.md) |

## Parent Command

- [`pingcli pingone`](cmd-pingcli-pingone.md) — Administration tools for the PingOne platform.
