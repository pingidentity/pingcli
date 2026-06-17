# `pingcli pingone agreements agreement-languages`
Agreement Languages

## Synopsis

Agreement Languages

```
pingcli pingone agreements agreement-languages [flags]
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
| `pingcli pingone agreements agreement-languages apply` | Create or update an agreement language | [`cmd-pingcli-pingone-agreements-agreement-languages-apply.md`](cmd-pingcli-pingone-agreements-agreement-languages-apply.md) |
| `pingcli pingone agreements agreement-languages create` | Create a new agreement language | [`cmd-pingcli-pingone-agreements-agreement-languages-create.md`](cmd-pingcli-pingone-agreements-agreement-languages-create.md) |
| `pingcli pingone agreements agreement-languages delete` | Delete an agreement language | [`cmd-pingcli-pingone-agreements-agreement-languages-delete.md`](cmd-pingcli-pingone-agreements-agreement-languages-delete.md) |
| `pingcli pingone agreements agreement-languages get` | Read a specific agreement language | [`cmd-pingcli-pingone-agreements-agreement-languages-get.md`](cmd-pingcli-pingone-agreements-agreement-languages-get.md) |
| `pingcli pingone agreements agreement-languages list` | List all agreement languages | [`cmd-pingcli-pingone-agreements-agreement-languages-list.md`](cmd-pingcli-pingone-agreements-agreement-languages-list.md) |
| `pingcli pingone agreements agreement-languages replace` | Update an agreement language | [`cmd-pingcli-pingone-agreements-agreement-languages-replace.md`](cmd-pingcli-pingone-agreements-agreement-languages-replace.md) |
| `pingcli pingone agreements agreement-languages template` | Generate an agreement language JSON template | [`cmd-pingcli-pingone-agreements-agreement-languages-template.md`](cmd-pingcli-pingone-agreements-agreement-languages-template.md) |

## Parent Command

- [`pingcli pingone agreements`](cmd-pingcli-pingone-agreements.md) — Agreements
