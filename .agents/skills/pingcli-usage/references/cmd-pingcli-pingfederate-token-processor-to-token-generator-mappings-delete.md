# `pingcli pingfederate token-processor-to-token-generator-mappings delete`
Delete a token processor to token generator mapping

## Synopsis

Delete a PingFederate token processor to token generator mapping

```
pingcli pingfederate token-processor-to-token-generator-mappings delete [flags]
```

## Examples

```
# Delete a token processor to token generator mapping
  pingcli pingfederate token-processor-to-token-generator-mappings delete --id <id>
```

## Options

| Flag | Default | Description |
|------|---------|-------------|
| `-h, --help` | `` | help for delete |
| `--id string` | `` | The ID of the token processor to token generator mapping |


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


## Parent Command

- [`pingcli pingfederate token-processor-to-token-generator-mappings`](cmd-pingcli-pingfederate-token-processor-to-token-generator-mappings.md) — PingFederate Token Processor to Token Generator Mappings
