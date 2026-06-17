# `pingcli pingone groups group-role-assignments delete`
Delete a group role assignment

## Synopsis

Delete a group role assignment from a PingOne environment

```
pingcli pingone groups group-role-assignments delete [flags]
```

## Examples

```
# Delete a group role assignment
  pingcli pingone groups group-role-assignments delete --environment-id <env-id> --group-id <group-id> --role-assignment-id <role-assignment-id>
```

## Options

| Flag | Default | Description |
|------|---------|-------------|
| `-h, --help` | `` | help for delete |
| `-e, --environment-id string` | `` | The PingOne environment ID |
| `-g, --group-id string` | `` | The group ID |
| `-r, --role-assignment-id string` | `` | The role assignment ID |


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

- [`pingcli pingone groups group-role-assignments`](cmd-pingcli-pingone-groups-group-role-assignments.md) — Group Role Assignments
