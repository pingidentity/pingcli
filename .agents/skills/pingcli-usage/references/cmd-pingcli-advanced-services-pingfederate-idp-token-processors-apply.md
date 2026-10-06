# `pingcli advanced-services pingfederate idp token-processors apply`
Create or update a token processor

## Synopsis

Idempotently create or update a PingFederate token processor looked up by the "name" field in the JSON body. If no token processor with the given name exists it is created; if it exists it is updated.

```
pingcli advanced-services pingfederate idp token-processors apply [flags]
```

## Examples

```
# Create or update a token processor (body supplies name and other fields)
  pingcli advanced-services pingfederate idp token-processors apply --from-file processor.json

  # Read body from stdin
  pingcli advanced-services pingfederate idp token-processors apply --from-file - < processor.json

  # Create or update from a JSON file, overriding the parent reference
  pingcli advanced-services pingfederate idp token-processors apply --from-file processor.json --parent-ref-id my-parent-processor
```

## Options

| Flag | Default | Description |
|------|---------|-------------|
| `-h, --help` | `` | help for apply |
| `-f, --from-file string` | `` | Path to a JSON file containing the request body, or "-" to read from stdin. |
| `--id string` | `` | The PingFederate Token Processor ID |
| `--name string` | `` | Plugin instance display name |
| `--parent-ref-id string` | `` | ID of a parent token processor instance to inherit configuration from |
| `--plugin-descriptor-ref-id string` | `` | ID of the token processor plugin type descriptor |


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

- [`pingcli advanced-services pingfederate idp token-processors`](cmd-pingcli-advanced-services-pingfederate-idp-token-processors.md) — PingFederate Token Processors
