# `pingcli pingone groups`
Groups

## Synopsis

Groups

```
pingcli pingone groups [flags]
```

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


## Subcommands

| Command | Description | Reference |
|---------|-------------|----------|
| `pingcli pingone groups apply` | Create or update a group | [`cmd-pingcli-pingone-groups-apply.md`](cmd-pingcli-pingone-groups-apply.md) |
| `pingcli pingone groups create` | Create a new group | [`cmd-pingcli-pingone-groups-create.md`](cmd-pingcli-pingone-groups-create.md) |
| `pingcli pingone groups delete` | Delete a group | [`cmd-pingcli-pingone-groups-delete.md`](cmd-pingcli-pingone-groups-delete.md) |
| `pingcli pingone groups get` | Read a specific group | [`cmd-pingcli-pingone-groups-get.md`](cmd-pingcli-pingone-groups-get.md) |
| `pingcli pingone groups group-nestings` | Group Nestings | [`cmd-pingcli-pingone-groups-group-nestings.md`](cmd-pingcli-pingone-groups-group-nestings.md) |
| `pingcli pingone groups group-role-assignments` | Group Role Assignments | [`cmd-pingcli-pingone-groups-group-role-assignments.md`](cmd-pingcli-pingone-groups-group-role-assignments.md) |
| `pingcli pingone groups list` | List all groups | [`cmd-pingcli-pingone-groups-list.md`](cmd-pingcli-pingone-groups-list.md) |
| `pingcli pingone groups replace` | Update a group | [`cmd-pingcli-pingone-groups-replace.md`](cmd-pingcli-pingone-groups-replace.md) |
| `pingcli pingone groups template` | Generate a group JSON template | [`cmd-pingcli-pingone-groups-template.md`](cmd-pingcli-pingone-groups-template.md) |

## Parent Command

- [`pingcli pingone`](cmd-pingcli-pingone.md) — Administration tools for the PingOne platform.
