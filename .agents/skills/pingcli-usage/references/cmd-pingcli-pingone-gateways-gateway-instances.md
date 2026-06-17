# `pingcli pingone gateways gateway-instances`
Gateway instances

## Synopsis

Gateway instances for a PingOne gateway

```
pingcli pingone gateways gateway-instances [flags]
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
| `pingcli pingone gateways gateway-instances get` | Read a specific gateway instance | [`cmd-pingcli-pingone-gateways-gateway-instances-get.md`](cmd-pingcli-pingone-gateways-gateway-instances-get.md) |
| `pingcli pingone gateways gateway-instances list` | List all gateway instances | [`cmd-pingcli-pingone-gateways-gateway-instances-list.md`](cmd-pingcli-pingone-gateways-gateway-instances-list.md) |

## Parent Command

- [`pingcli pingone gateways`](cmd-pingcli-pingone-gateways.md) — Gateways
