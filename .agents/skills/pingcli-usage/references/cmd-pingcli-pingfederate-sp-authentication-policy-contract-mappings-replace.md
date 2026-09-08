# `pingcli pingfederate sp authentication-policy-contract-mappings replace`
Update an authentication policy contract mapping

## Synopsis

Update (replace) a PingFederate Authentication Policy Contract to SP Adapter mapping

```
pingcli pingfederate sp authentication-policy-contract-mappings replace [flags]
```

## Examples

```
# Update an authentication policy contract mapping from a JSON file (--id is still required)
  pingcli pingfederate sp authentication-policy-contract-mappings replace --id <id> --from-file apc-to-sp-adapter-mapping.json

  # Update an authentication policy contract mapping from stdin
  pingcli pingfederate sp authentication-policy-contract-mappings replace --id <id> --from-file - < apc-to-sp-adapter-mapping.json

  # Update using --from-file for the attribute contract fulfillment, overriding the source and target
  pingcli pingfederate sp authentication-policy-contract-mappings replace --id <id> --from-file apc-to-sp-adapter-mapping.json --source-id <apc-id> --target-id <sp-adapter-id> --default-target-resource https://sp.example.com/resource
```

## Options

| Flag | Default | Description |
|------|---------|-------------|
| `-h, --help` | `` | help for replace |
| `-f, --from-file string` | `` | Path to a JSON file containing the request body, or "-" to read from stdin. |
| `--default-target-resource string` | `` | Default target URL for this APC-to-adapter mapping configuration |
| `--id string` | `` | The ID of the APC-to-SP Adapter mapping |
| `--license-connection-group-assignment string` | `` | The license connection group |
| `--source-id string` | `` | The id of the Authentication Policy Contract |
| `--target-id string` | `` | The id of the SP Adapter |


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

- [`pingcli pingfederate sp authentication-policy-contract-mappings`](cmd-pingcli-pingfederate-sp-authentication-policy-contract-mappings.md) — PingFederate Authentication Policy Contract Mappings
