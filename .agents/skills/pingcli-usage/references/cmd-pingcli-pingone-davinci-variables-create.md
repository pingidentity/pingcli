# `pingcli pingone davinci variables create`
Create a new DaVinci variable

## Synopsis

Create a new DaVinci variable in a PingOne environment

```
pingcli pingone davinci variables create [flags]
```

## Examples

```
# Create a new DaVinci variable from flags
  pingcli pingone davinci variables create --environment-id <env-id> --name myVar --context company --data-type string --mutable=true

  # Create a new DaVinci variable from a JSON file
  pingcli pingone davinci variables create --environment-id <env-id> --from-file variable.json

  # Create a new DaVinci variable from stdin
  pingcli pingone davinci variables create --environment-id <env-id> --from-file - < variable.json
```

## Options

| Flag | Default | Description |
|------|---------|-------------|
| `-h, --help` | `` | help for create |
| `-e, --environment-id string` | `` | The PingOne environment ID |
| `-f, --from-file string` | `` | Path to a JSON file containing the request body, or "-" to read from stdin. |
| `--context string` | `` | The variable context scope (company, flow, flowInstance, or user) |
| `--data-type string` | `` | The variable data type (boolean, number, object, secret, or string) |
| `--display-name string` | `` | The human-readable display name for the variable |
| `--flow-id string` | `` | The DaVinci flow ID this variable is scoped to; applicable when --context is flow |
| `--max string` | `` | The maximum value constraint for numeric variables |
| `--min string` | `` | The minimum value constraint for numeric variables |
| `--mutable` | `` | Whether the variable value can be changed at runtime |
| `--name string` | `` | The variable name |


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

- [`pingcli pingone davinci variables`](cmd-pingcli-pingone-davinci-variables.md) — DaVinci Variables
