# `pingcli pingone authorize permissions`
Application Resource Permissions

## Synopsis

Manage PingOne Authorize application resource permissions — actions that can be performed on an application resource

```
pingcli pingone authorize permissions [flags]
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
| `pingcli pingone authorize permissions apply` | Create or update an application resource permission | [`cmd-pingcli-pingone-authorize-permissions-apply.md`](cmd-pingcli-pingone-authorize-permissions-apply.md) |
| `pingcli pingone authorize permissions create` | Create a new application resource permission | [`cmd-pingcli-pingone-authorize-permissions-create.md`](cmd-pingcli-pingone-authorize-permissions-create.md) |
| `pingcli pingone authorize permissions delete` | Delete an application resource permission | [`cmd-pingcli-pingone-authorize-permissions-delete.md`](cmd-pingcli-pingone-authorize-permissions-delete.md) |
| `pingcli pingone authorize permissions get` | Read a specific application resource permission | [`cmd-pingcli-pingone-authorize-permissions-get.md`](cmd-pingcli-pingone-authorize-permissions-get.md) |
| `pingcli pingone authorize permissions list` | List all application resource permissions | [`cmd-pingcli-pingone-authorize-permissions-list.md`](cmd-pingcli-pingone-authorize-permissions-list.md) |
| `pingcli pingone authorize permissions replace` | Update an application resource permission | [`cmd-pingcli-pingone-authorize-permissions-replace.md`](cmd-pingcli-pingone-authorize-permissions-replace.md) |
| `pingcli pingone authorize permissions template` | Generate an application resource permission JSON template | [`cmd-pingcli-pingone-authorize-permissions-template.md`](cmd-pingcli-pingone-authorize-permissions-template.md) |

## Parent Command

- [`pingcli pingone authorize`](cmd-pingcli-pingone-authorize.md) — Administration tools for the PingOne Authorize universal service.
