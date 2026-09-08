# `pingcli pingfederate idp sts-request-parameters-contracts create`
Create a new STS request parameters contract

## Synopsis

Create a new PingFederate STS request parameters contract

```
pingcli pingfederate idp sts-request-parameters-contracts create [flags]
```

## Examples

```
# Create a contract from a JSON file
  pingcli pingfederate idp sts-request-parameters-contracts create --from-file contract.json

  # Read the body from stdin
  pingcli pingfederate idp sts-request-parameters-contracts create --from-file - < contract.json

  # Create a contract using flags
  pingcli pingfederate idp sts-request-parameters-contracts create --id contract-id --name "Example contract" --parameters "param-a,param-b"
```

## Options

| Flag | Default | Description |
|------|---------|-------------|
| `-h, --help` | `` | help for create |
| `-f, --from-file string` | `` | Path to a JSON file containing the request body, or "-" to read from stdin. |
| `--id string` | `` | The STS request parameters contract ID |
| `--name string` | `` | The STS request parameters contract name |
| `--parameters []string` | `` | Request parameter names or values; repeatable or comma-separated |


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

- [`pingcli pingfederate idp sts-request-parameters-contracts`](cmd-pingcli-pingfederate-idp-sts-request-parameters-contracts.md) — PingFederate STS request parameters contracts
