# `pingcli config`
Manage the CLI configuration.

## Synopsis

Manage the configuration of the CLI, including Ping product connection parameters.

The Ping CLI supports the use of configuration profiles.
Configuration profiles can be used when connecting to multiple environments using the same Ping CLI instance, such as when managing multiple development or demonstration environments.

A pre-defined default profile will be used to store the configuration of the CLI.
Additional custom profiles can be created using the `pingcli config profiles add` command.
To use a custom profile, the `--profile` flag can be used on supported commands to specify the profile to use for that command.
To set a custom profile as the default, use the `pingcli config profiles set-active` command.

```
pingcli config
```

## Options

| Flag | Default | Description |
|------|---------|-------------|
| `-h, --help` | `` | help for config |
| `-U, --unmask-values` | `` | Unmask secret values. (default false) |


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
| `pingcli config get` | Read stored configuration settings for the CLI. | [`cmd-pingcli-config-get.md`](cmd-pingcli-config-get.md) |
| `pingcli config list-keys` | List all configuration keys. | [`cmd-pingcli-config-list-keys.md`](cmd-pingcli-config-list-keys.md) |
| `pingcli config profiles` | Manage the configuration profiles. | [`cmd-pingcli-config-profiles.md`](cmd-pingcli-config-profiles.md) |
| `pingcli config set` | Set stored configuration settings for the CLI. | [`cmd-pingcli-config-set.md`](cmd-pingcli-config-set.md) |
| `pingcli config unset` | Unset stored configuration settings for the CLI. | [`cmd-pingcli-config-unset.md`](cmd-pingcli-config-unset.md) |

## Parent Command

- [`pingcli`](cmd-pingcli.md) — A CLI tool for managing the configuration of Ping Identity products.
