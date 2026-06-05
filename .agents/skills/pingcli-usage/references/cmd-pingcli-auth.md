# `pingcli auth`
Authenticate (or refresh authentication) for multiple connected products and services at once.

## Synopsis

Authenticate (or refresh authentication) for multiple connected products and services at once.

```
pingcli auth
```

## Examples

```
Authenticate (or refresh authentication) across all connected products and services.
    pingcli auth
```

## Options

| Flag | Default | Description |
|------|---------|-------------|
| `-h, --help` | `` | help for auth |
| `--storage-type string` | `` | Auth token storage (default: secure_local)   secure_local  - Use OS keychain (default)   file_system   - Store tokens in ~/.pingcli/credentials   none          - Do not persist tokens |


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
| `pingcli auth login` | Log in to all enabled Ping services at once. | [`cmd-pingcli-auth-login.md`](cmd-pingcli-auth-login.md) |
| `pingcli auth logout` | Log out of all enabled Ping services at once. | [`cmd-pingcli-auth-logout.md`](cmd-pingcli-auth-logout.md) |
| `pingcli auth status` | Show authentication status for all enabled Ping services. | [`cmd-pingcli-auth-status.md`](cmd-pingcli-auth-status.md) |

## Parent Command

- [`pingcli`](cmd-pingcli.md) — A CLI tool for managing the configuration of Ping Identity products.
