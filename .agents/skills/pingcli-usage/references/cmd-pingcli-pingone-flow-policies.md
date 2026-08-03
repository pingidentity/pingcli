# `pingcli pingone flow-policies`
Flow policies

## Synopsis

View flow policies in a PingOne environment

```
pingcli pingone flow-policies [flags]
```

## Examples

```
# List all flow policies in an environment
  pingcli pingone flow-policies --environment-id <env-id>

  # Read a specific flow policy
  pingcli pingone flow-policies --environment-id <env-id> --flow-policy-id <flow-policy-id>
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
| `pingcli pingone flow-policies get` | Read a specific flow policy | [`cmd-pingcli-pingone-flow-policies-get.md`](cmd-pingcli-pingone-flow-policies-get.md) |
| `pingcli pingone flow-policies list` | List all flow policies | [`cmd-pingcli-pingone-flow-policies-list.md`](cmd-pingcli-pingone-flow-policies-list.md) |

## Parent Command

- [`pingcli pingone`](cmd-pingcli-pingone.md) — Administration tools for the PingOne platform.
