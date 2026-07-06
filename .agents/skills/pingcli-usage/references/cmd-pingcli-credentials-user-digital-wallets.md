# `pingcli credentials user-digital-wallets`
User Digital Wallets

## Synopsis

User Digital Wallets

```
pingcli credentials user-digital-wallets [flags]
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
| `pingcli credentials user-digital-wallets create` | Create a new user digital wallet | [`cmd-pingcli-credentials-user-digital-wallets-create.md`](cmd-pingcli-credentials-user-digital-wallets-create.md) |
| `pingcli credentials user-digital-wallets delete` | Delete a user digital wallet | [`cmd-pingcli-credentials-user-digital-wallets-delete.md`](cmd-pingcli-credentials-user-digital-wallets-delete.md) |
| `pingcli credentials user-digital-wallets get` | Read a specific user digital wallet | [`cmd-pingcli-credentials-user-digital-wallets-get.md`](cmd-pingcli-credentials-user-digital-wallets-get.md) |
| `pingcli credentials user-digital-wallets list` | List all user digital wallets | [`cmd-pingcli-credentials-user-digital-wallets-list.md`](cmd-pingcli-credentials-user-digital-wallets-list.md) |
| `pingcli credentials user-digital-wallets provisioned-credentials` | List provisioned credentials for a user digital wallet | [`cmd-pingcli-credentials-user-digital-wallets-provisioned-credentials.md`](cmd-pingcli-credentials-user-digital-wallets-provisioned-credentials.md) |
| `pingcli credentials user-digital-wallets replace` | Replace a user digital wallet | [`cmd-pingcli-credentials-user-digital-wallets-replace.md`](cmd-pingcli-credentials-user-digital-wallets-replace.md) |
| `pingcli credentials user-digital-wallets template` | Generate a user digital wallet JSON template | [`cmd-pingcli-credentials-user-digital-wallets-template.md`](cmd-pingcli-credentials-user-digital-wallets-template.md) |

## Parent Command

- [`pingcli credentials`](cmd-pingcli-credentials.md) — Administration tools for the PingOne Credentials universal service.
