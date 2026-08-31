# `pingcli pingfederate sp idp-connections apply`
Create or update an IdP connection

## Synopsis

Idempotently create or update a PingFederate SP IdP connection looked up by the "name" field in the JSON body. If no connection with the given name exists it is created; if it exists it is updated.

```
pingcli pingfederate sp idp-connections apply [flags]
```

## Examples

```
# Create or update an IdP connection (body supplies name and other fields)
  pingcli pingfederate sp idp-connections apply --from-file idp-connection.json

  # Read body from stdin
  pingcli pingfederate sp idp-connections apply --from-file - < idp-connection.json

  # Create or update using flags to override fields in the JSON body
  pingcli pingfederate sp idp-connections apply --entity-id https://idp.example.com --active --from-file config.json
```

## Options

| Flag | Default | Description |
|------|---------|-------------|
| `-h, --help` | `` | help for apply |
| `-f, --from-file string` | `` | Path to a JSON file containing the request body, or "-" to read from stdin. |
| `--active` | `` | Whether the connection is active |
| `--base-url string` | `` | The partner federation deployment base URL |
| `--default-virtual-entity-id string` | `` | The default alternate entity ID for this connection |
| `--entity-id string` | `` | The partner entity ID or issuer value |
| `--error-page-msg-id string` | `` | The identifier for the user-facing error page message |
| `--id string` | `` | The PingFederate SP IdP connection ID |
| `--license-connection-group string` | `` | The license connection group assigned to this connection |
| `--logging-mode string` | `` | The transaction logging level for this connection |
| `--name string` | `` | The connection display name |
| `--type string` | `` | The connection type |
| `--virtual-entity-ids []string` | `` | Alternate entity IDs for this connection; repeat or separate values with commas |


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

- [`pingcli pingfederate sp idp-connections`](cmd-pingcli-pingfederate-sp-idp-connections.md) — PingFederate SP IdP Connections
