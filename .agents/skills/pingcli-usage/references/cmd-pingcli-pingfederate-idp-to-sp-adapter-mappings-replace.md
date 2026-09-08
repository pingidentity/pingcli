# `pingcli pingfederate idp-to-sp-adapter-mappings replace`
Update an IdP-to-SP Adapter mapping

## Synopsis

Update (replace) a PingFederate IdP-to-SP Adapter mapping

```
pingcli pingfederate idp-to-sp-adapter-mappings replace [flags]
```

## Examples

```
# Update an IdP-to-SP Adapter mapping from a JSON file (--id is still required)
  pingcli pingfederate idp-to-sp-adapter-mappings replace --id <id> --from-file idp-to-sp-adapter-mapping.json

  # Update an IdP-to-SP Adapter mapping from stdin
  pingcli pingfederate idp-to-sp-adapter-mappings replace --id <id> --from-file - < idp-to-sp-adapter-mapping.json

  # Update using flags for the scalar fields, and --from-file for the attribute contract fulfillment
  pingcli pingfederate idp-to-sp-adapter-mappings replace --id <id> --from-file idp-to-sp-adapter-mapping.json --source-id <idp-adapter-id> --target-id <sp-adapter-id> --default-target-resource <url> --license-connection-group-assignment <group> --application-name <name> --application-icon-url <url>
```

## Options

| Flag | Default | Description |
|------|---------|-------------|
| `-h, --help` | `` | help for replace |
| `-f, --from-file string` | `` | Path to a JSON file containing the request body, or "-" to read from stdin. |
| `--application-icon-url string` | `` | The application icon URL |
| `--application-name string` | `` | The application name |
| `--default-target-resource string` | `` | The default target URL for this adapter-to-adapter mapping configuration |
| `--id string` | `` | The ID of the IdP-to-SP Adapter mapping |
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

- [`pingcli pingfederate idp-to-sp-adapter-mappings`](cmd-pingcli-pingfederate-idp-to-sp-adapter-mappings.md) — PingFederate IdP-to-SP Adapter Mappings
