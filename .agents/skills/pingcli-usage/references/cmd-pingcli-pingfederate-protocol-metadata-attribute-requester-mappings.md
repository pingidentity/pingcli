# `pingcli pingfederate protocol-metadata attribute-requester-mappings`
PingFederate Protocol Metadata Attribute Requester Mappings

## Synopsis

PingFederate Protocol Metadata Attribute Requester Mappings

```
pingcli pingfederate protocol-metadata attribute-requester-mappings [flags]
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
| `pingcli pingfederate protocol-metadata attribute-requester-mappings apply` | Update attribute requester mappings | [`cmd-pingcli-pingfederate-protocol-metadata-attribute-requester-mappings-apply.md`](cmd-pingcli-pingfederate-protocol-metadata-attribute-requester-mappings-apply.md) |
| `pingcli pingfederate protocol-metadata attribute-requester-mappings get` | Read attribute requester mappings | [`cmd-pingcli-pingfederate-protocol-metadata-attribute-requester-mappings-get.md`](cmd-pingcli-pingfederate-protocol-metadata-attribute-requester-mappings-get.md) |
| `pingcli pingfederate protocol-metadata attribute-requester-mappings replace` | Update attribute requester mappings | [`cmd-pingcli-pingfederate-protocol-metadata-attribute-requester-mappings-replace.md`](cmd-pingcli-pingfederate-protocol-metadata-attribute-requester-mappings-replace.md) |
| `pingcli pingfederate protocol-metadata attribute-requester-mappings template` | Generate an attribute requester mappings JSON template | [`cmd-pingcli-pingfederate-protocol-metadata-attribute-requester-mappings-template.md`](cmd-pingcli-pingfederate-protocol-metadata-attribute-requester-mappings-template.md) |

## Parent Command

- [`pingcli pingfederate protocol-metadata`](cmd-pingcli-pingfederate-protocol-metadata.md) — Manage PingFederate Protocol Metadata resources
