# `pingcli pingfederate oauth token-exchange token-generator-mappings create`
Create a new token generator mapping

## Synopsis

Create a new PingFederate token exchange processor policy to token generator mapping

```
pingcli pingfederate oauth token-exchange token-generator-mappings create [flags]
```

## Examples

```
# Create a new token generator mapping from a JSON file
  pingcli pingfederate oauth token-exchange token-generator-mappings create --from-file token-generator-mapping.json

  # Create a new token generator mapping from stdin
  pingcli pingfederate oauth token-exchange token-generator-mappings create --from-file - < token-generator-mapping.json

  # Create using --from-file for the attribute contract fulfillment, overriding the source and target
  pingcli pingfederate oauth token-exchange token-generator-mappings create --from-file token-generator-mapping.json --source-id <processor-policy-id> --target-id <generator-id>
```

## Options

| Flag | Default | Description |
|------|---------|-------------|
| `-h, --help` | `` | help for create |
| `-f, --from-file string` | `` | Path to a JSON file containing the request body, or "-" to read from stdin. |
| `--license-connection-group-assignment string` | `` | The license connection group |
| `--source-id string` | `` | The ID of the Token Exchange Processor policy |
| `--target-id string` | `` | The ID of the Token Generator |


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

- [`pingcli pingfederate oauth token-exchange token-generator-mappings`](cmd-pingcli-pingfederate-oauth-token-exchange-token-generator-mappings.md) — PingFederate Token Exchange Processor policy to Token Generator Mappings
