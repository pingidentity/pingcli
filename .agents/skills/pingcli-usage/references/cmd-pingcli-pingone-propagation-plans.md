# `pingcli pingone propagation-plans`
Identity Propagation Plans

## Synopsis

Identity Propagation Plans

```
pingcli pingone propagation-plans [flags]
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
| `pingcli pingone propagation-plans apply` | Create or update an identity propagation plan | [`cmd-pingcli-pingone-propagation-plans-apply.md`](cmd-pingcli-pingone-propagation-plans-apply.md) |
| `pingcli pingone propagation-plans create` | Create a new identity propagation plan | [`cmd-pingcli-pingone-propagation-plans-create.md`](cmd-pingcli-pingone-propagation-plans-create.md) |
| `pingcli pingone propagation-plans delete` | Delete an identity propagation plan | [`cmd-pingcli-pingone-propagation-plans-delete.md`](cmd-pingcli-pingone-propagation-plans-delete.md) |
| `pingcli pingone propagation-plans get` | Read a specific identity propagation plan | [`cmd-pingcli-pingone-propagation-plans-get.md`](cmd-pingcli-pingone-propagation-plans-get.md) |
| `pingcli pingone propagation-plans list` | List all identity propagation plans | [`cmd-pingcli-pingone-propagation-plans-list.md`](cmd-pingcli-pingone-propagation-plans-list.md) |
| `pingcli pingone propagation-plans replace` | Update an identity propagation plan | [`cmd-pingcli-pingone-propagation-plans-replace.md`](cmd-pingcli-pingone-propagation-plans-replace.md) |
| `pingcli pingone propagation-plans template` | Generate an identity propagation plan JSON template | [`cmd-pingcli-pingone-propagation-plans-template.md`](cmd-pingcli-pingone-propagation-plans-template.md) |

## Parent Command

- [`pingcli pingone`](cmd-pingcli-pingone.md) — Administration tools for the PingOne platform.
