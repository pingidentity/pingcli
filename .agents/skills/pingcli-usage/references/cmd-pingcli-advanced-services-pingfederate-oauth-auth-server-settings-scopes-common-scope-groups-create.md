# `pingcli advanced-services pingfederate oauth auth-server-settings scopes common-scope-groups create`
Create a new common scope group

## Synopsis

Create a new PingFederate OAuth common scope group

```
pingcli advanced-services pingfederate oauth auth-server-settings scopes common-scope-groups create [flags]
```

## Examples

```
# Create a new common scope group from a JSON file
  pingcli advanced-services pingfederate oauth auth-server-settings scopes common-scope-groups create --from-file common-scope-group.json

  # Create a new common scope group from stdin
  pingcli advanced-services pingfederate oauth auth-server-settings scopes common-scope-groups create --from-file - < common-scope-group.json

  # Create a new common scope group from flags only (no --from-file)
  pingcli advanced-services pingfederate oauth auth-server-settings scopes common-scope-groups create --name myGroup --description "example group" --scopes email,openid

  # Create from a JSON file, overriding the description and scopes
  pingcli advanced-services pingfederate oauth auth-server-settings scopes common-scope-groups create --from-file common-scope-group.json --description "override" --scopes email
```

## Options

| Flag | Default | Description |
|------|---------|-------------|
| `-h, --help` | `` | help for create |
| `-f, --from-file string` | `` | Path to a JSON file containing the request body, or "-" to read from stdin. |
| `--description string` | `` | Description of the common scope group |
| `--name string` | `` | The PingFederate OAuth common scope group name |
| `--scopes []string` | `` | Common scope names that belong to the group; repeatable or comma-separated |


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

- [`pingcli advanced-services pingfederate oauth auth-server-settings scopes common-scope-groups`](cmd-pingcli-advanced-services-pingfederate-oauth-auth-server-settings-scopes-common-scope-groups.md) — PingFederate OAuth common scope groups
