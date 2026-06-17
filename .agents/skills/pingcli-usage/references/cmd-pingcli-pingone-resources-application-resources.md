# `pingcli pingone resources application-resources`
Application Resources

## Synopsis

Manage PingOne Authorize application resources — custom protected resources and parents of application resource permissions

```
pingcli pingone resources application-resources [flags]
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
| `pingcli pingone resources application-resources apply` | Create or update an application resource | [`cmd-pingcli-pingone-resources-application-resources-apply.md`](cmd-pingcli-pingone-resources-application-resources-apply.md) |
| `pingcli pingone resources application-resources create` | Create a new application resource | [`cmd-pingcli-pingone-resources-application-resources-create.md`](cmd-pingcli-pingone-resources-application-resources-create.md) |
| `pingcli pingone resources application-resources delete` | Delete an application resource | [`cmd-pingcli-pingone-resources-application-resources-delete.md`](cmd-pingcli-pingone-resources-application-resources-delete.md) |
| `pingcli pingone resources application-resources get` | Read a specific application resource | [`cmd-pingcli-pingone-resources-application-resources-get.md`](cmd-pingcli-pingone-resources-application-resources-get.md) |
| `pingcli pingone resources application-resources list` | List all application resources | [`cmd-pingcli-pingone-resources-application-resources-list.md`](cmd-pingcli-pingone-resources-application-resources-list.md) |
| `pingcli pingone resources application-resources replace` | Update an application resource | [`cmd-pingcli-pingone-resources-application-resources-replace.md`](cmd-pingcli-pingone-resources-application-resources-replace.md) |
| `pingcli pingone resources application-resources template` | Generate an application resource JSON template | [`cmd-pingcli-pingone-resources-application-resources-template.md`](cmd-pingcli-pingone-resources-application-resources-template.md) |

## Parent Command

- [`pingcli pingone resources`](cmd-pingcli-pingone-resources.md) — Resources
