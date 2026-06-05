# `pingcli pingone languages agreement-revisions`
Agreement Language Revisions

## Synopsis

Agreement Language Revisions

```
pingcli pingone languages agreement-revisions [flags]
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
| `pingcli pingone languages agreement-revisions create` | Create a new agreement language revision | [`cmd-pingcli-pingone-languages-agreement-revisions-create.md`](cmd-pingcli-pingone-languages-agreement-revisions-create.md) |
| `pingcli pingone languages agreement-revisions delete` | Delete an agreement language revision | [`cmd-pingcli-pingone-languages-agreement-revisions-delete.md`](cmd-pingcli-pingone-languages-agreement-revisions-delete.md) |
| `pingcli pingone languages agreement-revisions get` | Read a specific agreement language revision | [`cmd-pingcli-pingone-languages-agreement-revisions-get.md`](cmd-pingcli-pingone-languages-agreement-revisions-get.md) |
| `pingcli pingone languages agreement-revisions list` | List all agreement language revisions | [`cmd-pingcli-pingone-languages-agreement-revisions-list.md`](cmd-pingcli-pingone-languages-agreement-revisions-list.md) |
| `pingcli pingone languages agreement-revisions template` | Generate an agreement language revision JSON template | [`cmd-pingcli-pingone-languages-agreement-revisions-template.md`](cmd-pingcli-pingone-languages-agreement-revisions-template.md) |

## Parent Command

- [`pingcli pingone languages`](cmd-pingcli-pingone-languages.md) — Languages
