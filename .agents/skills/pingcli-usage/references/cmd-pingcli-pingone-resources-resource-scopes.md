# `pingcli pingone resources resource-scopes`
Resource Scopes

## Synopsis

Resource Scopes

```
pingcli pingone resources resource-scopes [flags]
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
| `pingcli pingone resources resource-scopes apply` | Create or update a resource scope | [`cmd-pingcli-pingone-resources-resource-scopes-apply.md`](cmd-pingcli-pingone-resources-resource-scopes-apply.md) |
| `pingcli pingone resources resource-scopes create` | Create a new resource scope | [`cmd-pingcli-pingone-resources-resource-scopes-create.md`](cmd-pingcli-pingone-resources-resource-scopes-create.md) |
| `pingcli pingone resources resource-scopes delete` | Delete a resource scope | [`cmd-pingcli-pingone-resources-resource-scopes-delete.md`](cmd-pingcli-pingone-resources-resource-scopes-delete.md) |
| `pingcli pingone resources resource-scopes get` | Read a specific resource scope | [`cmd-pingcli-pingone-resources-resource-scopes-get.md`](cmd-pingcli-pingone-resources-resource-scopes-get.md) |
| `pingcli pingone resources resource-scopes list` | List all resource scopes | [`cmd-pingcli-pingone-resources-resource-scopes-list.md`](cmd-pingcli-pingone-resources-resource-scopes-list.md) |
| `pingcli pingone resources resource-scopes replace` | Update a resource scope | [`cmd-pingcli-pingone-resources-resource-scopes-replace.md`](cmd-pingcli-pingone-resources-resource-scopes-replace.md) |
| `pingcli pingone resources resource-scopes template` | Generate a resource scope JSON template | [`cmd-pingcli-pingone-resources-resource-scopes-template.md`](cmd-pingcli-pingone-resources-resource-scopes-template.md) |

## Parent Command

- [`pingcli pingone resources`](cmd-pingcli-pingone-resources.md) — Resources
