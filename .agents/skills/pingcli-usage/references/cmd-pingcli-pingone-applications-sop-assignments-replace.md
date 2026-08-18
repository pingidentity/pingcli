# `pingcli pingone applications sop-assignments replace`
Update a sign-on policy assignment

## Synopsis

Update (replace) a sign-on policy assignment for an application in a PingOne environment

```
pingcli pingone applications sop-assignments replace [flags]
```

## Examples

```
# Update a sign-on policy assignment from flags
  pingcli pingone applications sop-assignments replace --environment-id <env-id> --application-id <app-id> --sop-assignment-id <assignment-id> --priority 1 --sign-on-policy-id <sop-id>

  # Update a sign-on policy assignment from a JSON file
  pingcli pingone applications sop-assignments replace --environment-id <env-id> --application-id <app-id> --sop-assignment-id <assignment-id> --from-file assignment.json

  # Update a sign-on policy assignment from stdin
  pingcli pingone applications sop-assignments replace --environment-id <env-id> --application-id <app-id> --sop-assignment-id <assignment-id> --from-file - < assignment.json

  # Update from a JSON file, overriding the priority from the file
  pingcli pingone applications sop-assignments replace --environment-id <env-id> --application-id <app-id> --sop-assignment-id <assignment-id> --from-file assignment.json --priority 2
```

## Options

| Flag | Default | Description |
|------|---------|-------------|
| `-h, --help` | `` | help for replace |
| `-a, --application-id string` | `` | The application ID |
| `-e, --environment-id string` | `` | The PingOne environment ID |
| `-f, --from-file string` | `` | Path to a JSON file containing the request body, or "-" to read from stdin. |
| `-s, --sop-assignment-id string` | `` | The sign-on policy assignment ID |
| `--priority int64` | `` | Priority of this sign-on policy assignment relative to other assignments on the application |
| `--sign-on-policy-id string` | `` | UUID of the sign-on policy to assign |


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

- [`pingcli pingone applications sop-assignments`](cmd-pingcli-pingone-applications-sop-assignments.md) — Sign-On Policy Assignments
