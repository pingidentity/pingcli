# `pingcli pingone users user-population`
User Population

## Synopsis

User Population

```
pingcli pingone users user-population [flags]
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
| `pingcli pingone users user-population get` | Read user population | [`cmd-pingcli-pingone-users-user-population-get.md`](cmd-pingcli-pingone-users-user-population-get.md) |
| `pingcli pingone users user-population replace` | Update user population | [`cmd-pingcli-pingone-users-user-population-replace.md`](cmd-pingcli-pingone-users-user-population-replace.md) |
| `pingcli pingone users user-population template` | Generate a user population JSON template | [`cmd-pingcli-pingone-users-user-population-template.md`](cmd-pingcli-pingone-users-user-population-template.md) |

## Parent Command

- [`pingcli pingone users`](cmd-pingcli-pingone-users.md) — Users
