# `pingcli davinci applications flow-policies replace`
Update a DaVinci flow policy

## Synopsis

Update (replace) a DaVinci flow policy on a DaVinci application in a PingOne environment

```
pingcli davinci applications flow-policies replace [flags]
```

## Examples

```
# Update a DaVinci flow policy from a JSON file (--environment-id, --application-id, and --flow-policy-id are still required)
  pingcli davinci applications flow-policies replace --environment-id <env-id> --application-id <app-id> --flow-policy-id <fp-id> --from-file flow-policy.json

  # Update a DaVinci flow policy from stdin
  pingcli davinci applications flow-policies replace --environment-id <env-id> --application-id <app-id> --flow-policy-id <fp-id> --from-file - < flow-policy.json
```

## Options

| Flag | Default | Description |
|------|---------|-------------|
| `-h, --help` | `` | help for replace |
| `-a, --application-id string` | `` | The DaVinci application ID |
| `-e, --environment-id string` | `` | The PingOne environment ID |
| `-f, --from-file string` | `` | Path to a JSON file containing the request body, or "-" to read from stdin. |
| `--flow-policy-id string` | `` | The DaVinci flow policy ID |


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

- [`pingcli davinci applications flow-policies`](cmd-pingcli-davinci-applications-flow-policies.md) — DaVinci Flow Policies
