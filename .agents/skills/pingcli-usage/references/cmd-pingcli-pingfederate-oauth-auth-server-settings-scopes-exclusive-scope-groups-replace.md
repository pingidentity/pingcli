# `pingcli pingfederate oauth auth-server-settings scopes exclusive-scope-groups replace`
Update an OAuth exclusive scope group

## Synopsis

Update (replace) a PingFederate OAuth exclusive scope group

```
pingcli pingfederate oauth auth-server-settings scopes exclusive-scope-groups replace [flags]
```

## Examples

```
# Update an OAuth exclusive scope group from a JSON file (--name is still required)
  pingcli pingfederate oauth auth-server-settings scopes exclusive-scope-groups replace --name <name> --from-file exclusive-scope-group.json

  # Update an OAuth exclusive scope group from stdin
  pingcli pingfederate oauth auth-server-settings scopes exclusive-scope-groups replace --name <name> --from-file - < exclusive-scope-group.json

  # Update an OAuth exclusive scope group from flags, without --from-file
  pingcli pingfederate oauth auth-server-settings scopes exclusive-scope-groups replace --name <name> --description "updated" --scopes scope_a,scope_b
```

## Options

| Flag | Default | Description |
|------|---------|-------------|
| `-h, --help` | `` | help for replace |
| `-f, --from-file string` | `` | Path to a JSON file containing the request body, or "-" to read from stdin. |
| `--description string` | `` | The description of the scope group |
| `--name string` | `` | The PingFederate exclusive scope group name |
| `--scopes []string` | `` | The set of scopes for this scope group; repeatable or comma-separated |


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

- [`pingcli pingfederate oauth auth-server-settings scopes exclusive-scope-groups`](cmd-pingcli-pingfederate-oauth-auth-server-settings-scopes-exclusive-scope-groups.md) — PingFederate OAuth exclusive scope groups
