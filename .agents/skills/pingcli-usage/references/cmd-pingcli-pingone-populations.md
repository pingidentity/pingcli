# `pingcli pingone populations`
Populations

## Synopsis

Populations

```
pingcli pingone populations [flags]
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
| `pingcli pingone populations apply` | Create or update a population | [`cmd-pingcli-pingone-populations-apply.md`](cmd-pingcli-pingone-populations-apply.md) |
| `pingcli pingone populations create` | Create a new population | [`cmd-pingcli-pingone-populations-create.md`](cmd-pingcli-pingone-populations-create.md) |
| `pingcli pingone populations delete` | Delete a population | [`cmd-pingcli-pingone-populations-delete.md`](cmd-pingcli-pingone-populations-delete.md) |
| `pingcli pingone populations get` | Read a specific population | [`cmd-pingcli-pingone-populations-get.md`](cmd-pingcli-pingone-populations-get.md) |
| `pingcli pingone populations list` | List all populations | [`cmd-pingcli-pingone-populations-list.md`](cmd-pingcli-pingone-populations-list.md) |
| `pingcli pingone populations replace` | Update a population | [`cmd-pingcli-pingone-populations-replace.md`](cmd-pingcli-pingone-populations-replace.md) |
| `pingcli pingone populations template` | Generate a population JSON template | [`cmd-pingcli-pingone-populations-template.md`](cmd-pingcli-pingone-populations-template.md) |

## Parent Command

- [`pingcli pingone`](cmd-pingcli-pingone.md) — Administration tools for the PingOne platform.
