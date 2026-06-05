# `pingcli pingone applications application-grants`
Application Grants

## Synopsis

Application Grants

```
pingcli pingone applications application-grants [flags]
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
| `pingcli pingone applications application-grants create` | Create a new application grant | [`cmd-pingcli-pingone-applications-application-grants-create.md`](cmd-pingcli-pingone-applications-application-grants-create.md) |
| `pingcli pingone applications application-grants delete` | Delete an application grant | [`cmd-pingcli-pingone-applications-application-grants-delete.md`](cmd-pingcli-pingone-applications-application-grants-delete.md) |
| `pingcli pingone applications application-grants get` | Read a specific application grant | [`cmd-pingcli-pingone-applications-application-grants-get.md`](cmd-pingcli-pingone-applications-application-grants-get.md) |
| `pingcli pingone applications application-grants list` | List all application grants | [`cmd-pingcli-pingone-applications-application-grants-list.md`](cmd-pingcli-pingone-applications-application-grants-list.md) |
| `pingcli pingone applications application-grants replace` | Update an application grant | [`cmd-pingcli-pingone-applications-application-grants-replace.md`](cmd-pingcli-pingone-applications-application-grants-replace.md) |
| `pingcli pingone applications application-grants template` | Generate an application grant JSON template | [`cmd-pingcli-pingone-applications-application-grants-template.md`](cmd-pingcli-pingone-applications-application-grants-template.md) |

## Parent Command

- [`pingcli pingone applications`](cmd-pingcli-pingone-applications.md) — Applications
