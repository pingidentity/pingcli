# `pingcli advanced-services pingfederate idp sp-connections replace`
Update an SP connection

## Synopsis

Update (replace) a PingFederate SP connection

```
pingcli advanced-services pingfederate idp sp-connections replace [flags]
```

## Examples

```
# Update an SP connection from a JSON file (--id is still required)
  pingcli advanced-services pingfederate idp sp-connections replace --id <id> --from-file sp-connection.json

  # Update an SP connection from stdin
  pingcli advanced-services pingfederate idp sp-connections replace --id <id> --from-file - < sp-connection.json

  # Update using flags for identity fields and --from-file for full configuration
  pingcli advanced-services pingfederate idp sp-connections replace --id <id> --entity-id https://sp.example.com --name "My SP Connection" --active --from-file config.json
```

## Options

| Flag | Default | Description |
|------|---------|-------------|
| `-h, --help` | `` | help for replace |
| `-f, --from-file string` | `` | Path to a JSON file containing the request body, or "-" to read from stdin. |
| `--active` | `` | Whether the connection is active |
| `--application-icon-url string` | `` | URL of the application icon |
| `--application-name string` | `` | The application name |
| `--base-url string` | `` | The partner federation deployment base URL |
| `--connection-target-type string` | `` | The connection target type (bulk import/export usage only) |
| `--default-virtual-entity-id string` | `` | The default alternate entity ID for this connection |
| `--entity-id string` | `` | The partner entity ID or issuer value |
| `--id string` | `` | The PingFederate SP connection ID |
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

- [`pingcli advanced-services pingfederate idp sp-connections`](cmd-pingcli-advanced-services-pingfederate-idp-sp-connections.md) — PingFederate SP Connections
