# `pingcli config profiles`
Manage the configuration profiles.

## Synopsis

Manage the configuration profiles.

The Ping CLI supports the use of configuration profiles.
Configuration profiles can be used when connecting to multiple environments using the same Ping CLI instance, such as when managing multiple development or demonstration environments.

A pre-defined default profile will be used to store the configuration of the CLI.
Additional custom profiles can be created using the `pingcli config add-profile` command.
To use a custom profile, the `--profile` flag can be used on supported commands to specify the profile to use for that command.
To set a custom profile as the default, use the `pingcli config set-active-profile` command.

```
pingcli config profiles
```

## Inherited Options

| Flag | Default | Description |
|------|---------|-------------|
| `-C, --config string` | `` | The relative or full path to a custom Ping CLI configuration file. (default $HOME/.pingcli/config.yaml) |
| `-D, --detailed-exitcode` | `` | Enable detailed exit code output. (default false) 0 - pingcli command succeeded with no errors or warnings. 1 - pingcli command failed with errors. 2 - pingcli command succeeded with warnings. |
| `-O, --output-format string` | `` | Specify the console output format. (default text) Options are: json, ndjson, ndjson-wrapped, text. |
| `-P, --profile string` | `` | The name of a configuration profile to use. |
| `-U, --unmask-values` | `` | Unmask secret values. (default false) |
| `--debug` | `` | Enable debug output for error messages, including stack traces and transaction IDs. (default false) |
| `--log-file string` | `` | Write logs to a file at the given path. File logging is disabled when not set. |
| `--log-file-level string` | `` | Set the file log level. Options are: DEBUG, INFO, WARN, ERROR. (default DEBUG) |
| `--log-level string` | `` | Set the console log level. Options are: DEBUG, INFO, WARN, ERROR. (default WARN) |
| `--no-color` | `` | Disable text output in color. (default false) |
| `--query string` | `` | JMESPath expression to filter JSON output. Requires -O json, ndjson, or ndjson-wrapped. Example: --query 'data[?enabled].name' |


## Subcommands

| Command | Description | Reference |
|---------|-------------|----------|
| `pingcli config profiles create` | Add a new custom configuration profile. | [`cmd-pingcli-config-profiles-create.md`](cmd-pingcli-config-profiles-create.md) |
| `pingcli config profiles delete` | Delete a custom configuration profile. | [`cmd-pingcli-config-profiles-delete.md`](cmd-pingcli-config-profiles-delete.md) |
| `pingcli config profiles list` | List all custom configuration profiles. | [`cmd-pingcli-config-profiles-list.md`](cmd-pingcli-config-profiles-list.md) |
| `pingcli config profiles show` | View the stored configuration of a custom configuration profile. | [`cmd-pingcli-config-profiles-show.md`](cmd-pingcli-config-profiles-show.md) |
| `pingcli config profiles use` | Activate and set a custom configuration profile as the in-use profile. | [`cmd-pingcli-config-profiles-use.md`](cmd-pingcli-config-profiles-use.md) |

## Parent Command

- [`pingcli config`](cmd-pingcli-config.md) — Manage the CLI configuration.
