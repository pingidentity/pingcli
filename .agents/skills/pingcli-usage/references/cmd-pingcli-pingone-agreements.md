# `pingcli pingone agreements`
Agreements

## Synopsis

Agreements

```
pingcli pingone agreements [flags]
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
| `pingcli pingone agreements agreement-languages` | Agreement Languages | [`cmd-pingcli-pingone-agreements-agreement-languages.md`](cmd-pingcli-pingone-agreements-agreement-languages.md) |
| `pingcli pingone agreements apply` | Create or update an agreement | [`cmd-pingcli-pingone-agreements-apply.md`](cmd-pingcli-pingone-agreements-apply.md) |
| `pingcli pingone agreements create` | Create a new agreement | [`cmd-pingcli-pingone-agreements-create.md`](cmd-pingcli-pingone-agreements-create.md) |
| `pingcli pingone agreements delete` | Delete an agreement | [`cmd-pingcli-pingone-agreements-delete.md`](cmd-pingcli-pingone-agreements-delete.md) |
| `pingcli pingone agreements get` | Read a specific agreement | [`cmd-pingcli-pingone-agreements-get.md`](cmd-pingcli-pingone-agreements-get.md) |
| `pingcli pingone agreements list` | List all agreements | [`cmd-pingcli-pingone-agreements-list.md`](cmd-pingcli-pingone-agreements-list.md) |
| `pingcli pingone agreements replace` | Update an agreement | [`cmd-pingcli-pingone-agreements-replace.md`](cmd-pingcli-pingone-agreements-replace.md) |
| `pingcli pingone agreements template` | Generate an agreement JSON template | [`cmd-pingcli-pingone-agreements-template.md`](cmd-pingcli-pingone-agreements-template.md) |

## Parent Command

- [`pingcli pingone`](cmd-pingcli-pingone.md) — Administration tools for the PingOne platform.
