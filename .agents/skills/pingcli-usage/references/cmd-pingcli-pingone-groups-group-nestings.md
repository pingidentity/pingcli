# `pingcli pingone groups group-nestings`
Group Nestings

## Synopsis

Group Nestings

```
pingcli pingone groups group-nestings [flags]
```

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


## Subcommands

| Command | Description | Reference |
|---------|-------------|----------|
| `pingcli pingone groups group-nestings create` | Create a new group nesting | [`cmd-pingcli-pingone-groups-group-nestings-create.md`](cmd-pingcli-pingone-groups-group-nestings-create.md) |
| `pingcli pingone groups group-nestings delete` | Delete a group nesting | [`cmd-pingcli-pingone-groups-group-nestings-delete.md`](cmd-pingcli-pingone-groups-group-nestings-delete.md) |
| `pingcli pingone groups group-nestings get` | Read a specific group nesting | [`cmd-pingcli-pingone-groups-group-nestings-get.md`](cmd-pingcli-pingone-groups-group-nestings-get.md) |
| `pingcli pingone groups group-nestings list` | List all group nestings | [`cmd-pingcli-pingone-groups-group-nestings-list.md`](cmd-pingcli-pingone-groups-group-nestings-list.md) |
| `pingcli pingone groups group-nestings template` | Generate a group nesting JSON template | [`cmd-pingcli-pingone-groups-group-nestings-template.md`](cmd-pingcli-pingone-groups-group-nestings-template.md) |

## Parent Command

- [`pingcli pingone groups`](cmd-pingcli-pingone-groups.md) — Groups
