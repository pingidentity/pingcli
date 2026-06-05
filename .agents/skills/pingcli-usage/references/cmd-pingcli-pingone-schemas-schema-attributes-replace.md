# `pingcli pingone schemas schema-attributes replace`
Update a schema attribute

## Synopsis

Update (replace) a schema attribute in a PingOne environment schema

```
pingcli pingone schemas schema-attributes replace [flags]
```

## Examples

```
# Update a schema attribute from a JSON file (path flags are still required)
  pingcli pingone schemas schema-attributes replace --environment-id <env-id> --schema-id <schema-id> --attribute-id <attribute-id> --from-file schema-attribute.json

  # Update a schema attribute from stdin
  pingcli pingone schemas schema-attributes replace --environment-id <env-id> --schema-id <schema-id> --attribute-id <attribute-id> --from-file - < schema-attribute.json
```

## Options

| Flag | Default | Description |
|------|---------|-------------|
| `-h, --help` | `` | help for replace |
| `-e, --environment-id string` | `` | The PingOne environment ID |
| `-f, --from-file string` | `` | Path to a JSON file containing the request body, or "-" to read from stdin. |
| `-i, --attribute-id string` | `` | The schema attribute ID |
| `-s, --schema-id string` | `` | The PingOne schema ID |


## Inherited Options

| Flag | Default | Description |
|------|---------|-------------|
| `-C, --config string` | `` | The relative or full path to a custom Ping CLI configuration file. (default $HOME/.pingcli/config.yaml) |
| `-D, --detailed-exitcode` | `` | Enable detailed exit code output. (default false) 0 - pingcli command succeeded with no errors or warnings. 1 - pingcli command failed with errors. 2 - pingcli command succeeded with warnings. |
| `-O, --output-format string` | `` | Specify the console output format. (default text) Options are: json, ndjson, ndjson-wrapped, text. |
| `-P, --profile string` | `` | The name of a configuration profile to use. |
| `--debug` | `` | Enable debug output for error messages, including stack traces and transaction IDs. (default false) |
| `--log-file string` | `` | Write logs to a file at the given path. File logging is disabled when not set. |
| `--log-file-level string` | `` | Set the file log level. Options are: DEBUG, INFO, WARN, ERROR. (default DEBUG) |
| `--log-level string` | `` | Set the console log level. Options are: DEBUG, INFO, WARN, ERROR. (default WARN) |
| `--no-color` | `` | Disable text output in color. (default false) |
| `--query string` | `` | JMESPath expression to filter JSON output. Requires -O json, ndjson, or ndjson-wrapped. Example: --query 'data[?enabled].name' |


## Parent Command

- [`pingcli pingone schemas schema-attributes`](cmd-pingcli-pingone-schemas-schema-attributes.md) — Schema Attributes
