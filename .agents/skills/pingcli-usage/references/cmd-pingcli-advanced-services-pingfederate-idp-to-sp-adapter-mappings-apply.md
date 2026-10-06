# `pingcli advanced-services pingfederate idp-to-sp-adapter-mappings apply`
Create or update an IdP-to-SP Adapter mapping

## Synopsis

Idempotently create or update a PingFederate IdP-to-SP Adapter mapping looked up by the "sourceId|targetId" composite key in the JSON body. If no mapping for that adapter pair exists it is created; if one exists it is updated.

```
pingcli advanced-services pingfederate idp-to-sp-adapter-mappings apply [flags]
```

## Examples

```
# Create or update an IdP-to-SP Adapter mapping (body supplies sourceId, targetId, and other fields)
  pingcli advanced-services pingfederate idp-to-sp-adapter-mappings apply --from-file idp-to-sp-adapter-mapping.json

  # Read body from stdin
  pingcli advanced-services pingfederate idp-to-sp-adapter-mappings apply --from-file - < idp-to-sp-adapter-mapping.json

  # Create or update from a JSON file, overriding scalar fields
  pingcli advanced-services pingfederate idp-to-sp-adapter-mappings apply --from-file idp-to-sp-adapter-mapping.json --source-id <idp-adapter-id> --target-id <sp-adapter-id> --default-target-resource <url> --license-connection-group-assignment <group> --application-name <name> --application-icon-url <url>
```

## Options

| Flag | Default | Description |
|------|---------|-------------|
| `-h, --help` | `` | help for apply |
| `-f, --from-file string` | `` | Path to a JSON file containing the request body, or "-" to read from stdin. |
| `--application-icon-url string` | `` | The application icon URL |
| `--application-name string` | `` | The application name |
| `--default-target-resource string` | `` | The default target URL for this adapter-to-adapter mapping configuration |
| `--license-connection-group-assignment string` | `` | The license connection group |
| `--source-id string` | `` | The ID of the IdP Adapter |
| `--target-id string` | `` | The ID of the SP Adapter |


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

- [`pingcli advanced-services pingfederate idp-to-sp-adapter-mappings`](cmd-pingcli-advanced-services-pingfederate-idp-to-sp-adapter-mappings.md) — PingFederate IdP-to-SP Adapter Mappings
