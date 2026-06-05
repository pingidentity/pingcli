# `pingcli pingone languages`
Languages

## Synopsis

Languages

```
pingcli pingone languages [flags]
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
| `pingcli pingone languages agreement-revisions` | Agreement Language Revisions | [`cmd-pingcli-pingone-languages-agreement-revisions.md`](cmd-pingcli-pingone-languages-agreement-revisions.md) |
| `pingcli pingone languages apply` | Create or update a language | [`cmd-pingcli-pingone-languages-apply.md`](cmd-pingcli-pingone-languages-apply.md) |
| `pingcli pingone languages create` | Create a new language | [`cmd-pingcli-pingone-languages-create.md`](cmd-pingcli-pingone-languages-create.md) |
| `pingcli pingone languages delete` | Delete a language | [`cmd-pingcli-pingone-languages-delete.md`](cmd-pingcli-pingone-languages-delete.md) |
| `pingcli pingone languages get` | Read a specific language | [`cmd-pingcli-pingone-languages-get.md`](cmd-pingcli-pingone-languages-get.md) |
| `pingcli pingone languages list` | List all languages | [`cmd-pingcli-pingone-languages-list.md`](cmd-pingcli-pingone-languages-list.md) |
| `pingcli pingone languages replace` | Update a language | [`cmd-pingcli-pingone-languages-replace.md`](cmd-pingcli-pingone-languages-replace.md) |
| `pingcli pingone languages template` | Generate a language JSON template | [`cmd-pingcli-pingone-languages-template.md`](cmd-pingcli-pingone-languages-template.md) |

## Parent Command

- [`pingcli pingone`](cmd-pingcli-pingone.md) — Administration tools for the PingOne platform.
