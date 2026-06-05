# `pingcli pingone applications flow-policy-assignments delete`
Delete a flow policy assignment

## Synopsis

Delete a flow policy assignment from a PingOne environment

```
pingcli pingone applications flow-policy-assignments delete [flags]
```

## Examples

```
# Delete a flow policy assignment
  pingcli pingone applications flow-policy-assignments delete --environment-id <env-id> --application-id <app-id> --flow-policy-assignment-id <assignment-id>
```

## Options

| Flag | Default | Description |
|------|---------|-------------|
| `-h, --help` | `` | help for delete |
| `-a, --application-id string` | `` | The application ID |
| `-e, --environment-id string` | `` | The PingOne environment ID |
| `--flow-policy-assignment-id string` | `` | The flow policy assignment ID |


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

- [`pingcli pingone applications flow-policy-assignments`](cmd-pingcli-pingone-applications-flow-policy-assignments.md) — Flow Policy Assignments
