# `pingcli pingone authorize application-roles`
Application Roles

## Synopsis

Application Roles

```
pingcli pingone authorize application-roles [flags]
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
| `pingcli pingone authorize application-roles apply` | Create or update an application role | [`cmd-pingcli-pingone-authorize-application-roles-apply.md`](cmd-pingcli-pingone-authorize-application-roles-apply.md) |
| `pingcli pingone authorize application-roles create` | Create a new application role | [`cmd-pingcli-pingone-authorize-application-roles-create.md`](cmd-pingcli-pingone-authorize-application-roles-create.md) |
| `pingcli pingone authorize application-roles delete` | Delete an application role | [`cmd-pingcli-pingone-authorize-application-roles-delete.md`](cmd-pingcli-pingone-authorize-application-roles-delete.md) |
| `pingcli pingone authorize application-roles get` | Read a specific application role | [`cmd-pingcli-pingone-authorize-application-roles-get.md`](cmd-pingcli-pingone-authorize-application-roles-get.md) |
| `pingcli pingone authorize application-roles list` | List all application roles | [`cmd-pingcli-pingone-authorize-application-roles-list.md`](cmd-pingcli-pingone-authorize-application-roles-list.md) |
| `pingcli pingone authorize application-roles replace` | Update an application role | [`cmd-pingcli-pingone-authorize-application-roles-replace.md`](cmd-pingcli-pingone-authorize-application-roles-replace.md) |
| `pingcli pingone authorize application-roles template` | Generate an application role JSON template | [`cmd-pingcli-pingone-authorize-application-roles-template.md`](cmd-pingcli-pingone-authorize-application-roles-template.md) |

## Parent Command

- [`pingcli pingone authorize`](cmd-pingcli-pingone-authorize.md) — Administration tools for the PingOne Authorize universal service.
