# `pingcli pingone schemas attributes`
Schema Attributes

## Synopsis

Schema Attributes

```
pingcli pingone schemas attributes [flags]
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
| `pingcli pingone schemas attributes apply` | Create or update a schema attribute | [`cmd-pingcli-pingone-schemas-attributes-apply.md`](cmd-pingcli-pingone-schemas-attributes-apply.md) |
| `pingcli pingone schemas attributes create` | Create a new schema attribute | [`cmd-pingcli-pingone-schemas-attributes-create.md`](cmd-pingcli-pingone-schemas-attributes-create.md) |
| `pingcli pingone schemas attributes delete` | Delete a schema attribute | [`cmd-pingcli-pingone-schemas-attributes-delete.md`](cmd-pingcli-pingone-schemas-attributes-delete.md) |
| `pingcli pingone schemas attributes get` | Read a specific schema attribute | [`cmd-pingcli-pingone-schemas-attributes-get.md`](cmd-pingcli-pingone-schemas-attributes-get.md) |
| `pingcli pingone schemas attributes list` | List all schema attributes | [`cmd-pingcli-pingone-schemas-attributes-list.md`](cmd-pingcli-pingone-schemas-attributes-list.md) |
| `pingcli pingone schemas attributes replace` | Update a schema attribute | [`cmd-pingcli-pingone-schemas-attributes-replace.md`](cmd-pingcli-pingone-schemas-attributes-replace.md) |
| `pingcli pingone schemas attributes template` | Generate a schema attribute JSON template | [`cmd-pingcli-pingone-schemas-attributes-template.md`](cmd-pingcli-pingone-schemas-attributes-template.md) |

## Parent Command

- [`pingcli pingone schemas`](cmd-pingcli-pingone-schemas.md) — Schemas
