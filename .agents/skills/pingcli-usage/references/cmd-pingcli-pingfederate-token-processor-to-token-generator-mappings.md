# `pingcli pingfederate token-processor-to-token-generator-mappings`
PingFederate Token Processor to Token Generator Mappings

## Synopsis

PingFederate Token Processor to Token Generator Mappings

```
pingcli pingfederate token-processor-to-token-generator-mappings [flags]
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
| `pingcli pingfederate token-processor-to-token-generator-mappings apply` | Create or update a token processor to token generator mapping | [`cmd-pingcli-pingfederate-token-processor-to-token-generator-mappings-apply.md`](cmd-pingcli-pingfederate-token-processor-to-token-generator-mappings-apply.md) |
| `pingcli pingfederate token-processor-to-token-generator-mappings create` | Create a new token processor to token generator mapping | [`cmd-pingcli-pingfederate-token-processor-to-token-generator-mappings-create.md`](cmd-pingcli-pingfederate-token-processor-to-token-generator-mappings-create.md) |
| `pingcli pingfederate token-processor-to-token-generator-mappings delete` | Delete a token processor to token generator mapping | [`cmd-pingcli-pingfederate-token-processor-to-token-generator-mappings-delete.md`](cmd-pingcli-pingfederate-token-processor-to-token-generator-mappings-delete.md) |
| `pingcli pingfederate token-processor-to-token-generator-mappings get` | Read a specific token processor to token generator mapping | [`cmd-pingcli-pingfederate-token-processor-to-token-generator-mappings-get.md`](cmd-pingcli-pingfederate-token-processor-to-token-generator-mappings-get.md) |
| `pingcli pingfederate token-processor-to-token-generator-mappings list` | List all token processor to token generator mappings | [`cmd-pingcli-pingfederate-token-processor-to-token-generator-mappings-list.md`](cmd-pingcli-pingfederate-token-processor-to-token-generator-mappings-list.md) |
| `pingcli pingfederate token-processor-to-token-generator-mappings replace` | Update a token processor to token generator mapping | [`cmd-pingcli-pingfederate-token-processor-to-token-generator-mappings-replace.md`](cmd-pingcli-pingfederate-token-processor-to-token-generator-mappings-replace.md) |
| `pingcli pingfederate token-processor-to-token-generator-mappings template` | Generate a token processor to token generator mapping JSON template | [`cmd-pingcli-pingfederate-token-processor-to-token-generator-mappings-template.md`](cmd-pingcli-pingfederate-token-processor-to-token-generator-mappings-template.md) |

## Parent Command

- [`pingcli pingfederate`](cmd-pingcli-pingfederate.md) — Administration tools for PingFederate deployed as software
