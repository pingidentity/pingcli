# `pingcli pingone custom-admin-roles`
Custom Admin Roles

## Synopsis

Custom Admin Roles

```
pingcli pingone custom-admin-roles [flags]
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
| `pingcli pingone custom-admin-roles apply` | Create or update a custom admin role | [`cmd-pingcli-pingone-custom-admin-roles-apply.md`](cmd-pingcli-pingone-custom-admin-roles-apply.md) |
| `pingcli pingone custom-admin-roles create` | Create a new custom admin role | [`cmd-pingcli-pingone-custom-admin-roles-create.md`](cmd-pingcli-pingone-custom-admin-roles-create.md) |
| `pingcli pingone custom-admin-roles delete` | Delete a custom admin role | [`cmd-pingcli-pingone-custom-admin-roles-delete.md`](cmd-pingcli-pingone-custom-admin-roles-delete.md) |
| `pingcli pingone custom-admin-roles get` | Read a specific custom admin role | [`cmd-pingcli-pingone-custom-admin-roles-get.md`](cmd-pingcli-pingone-custom-admin-roles-get.md) |
| `pingcli pingone custom-admin-roles list` | List all custom admin roles | [`cmd-pingcli-pingone-custom-admin-roles-list.md`](cmd-pingcli-pingone-custom-admin-roles-list.md) |
| `pingcli pingone custom-admin-roles replace` | Update a custom admin role | [`cmd-pingcli-pingone-custom-admin-roles-replace.md`](cmd-pingcli-pingone-custom-admin-roles-replace.md) |
| `pingcli pingone custom-admin-roles template` | Generate a custom admin role JSON template | [`cmd-pingcli-pingone-custom-admin-roles-template.md`](cmd-pingcli-pingone-custom-admin-roles-template.md) |

## Parent Command

- [`pingcli pingone`](cmd-pingcli-pingone.md) — Administration tools for the PingOne platform.
