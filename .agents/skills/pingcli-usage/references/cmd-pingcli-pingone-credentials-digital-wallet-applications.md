# `pingcli pingone credentials digital-wallet-applications`
Digital Wallet Applications

## Synopsis

Digital Wallet Applications

```
pingcli pingone credentials digital-wallet-applications [flags]
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
| `pingcli pingone credentials digital-wallet-applications apply` | Create or update a digital wallet application | [`cmd-pingcli-pingone-credentials-digital-wallet-applications-apply.md`](cmd-pingcli-pingone-credentials-digital-wallet-applications-apply.md) |
| `pingcli pingone credentials digital-wallet-applications create` | Create a new digital wallet application | [`cmd-pingcli-pingone-credentials-digital-wallet-applications-create.md`](cmd-pingcli-pingone-credentials-digital-wallet-applications-create.md) |
| `pingcli pingone credentials digital-wallet-applications delete` | Delete a digital wallet application | [`cmd-pingcli-pingone-credentials-digital-wallet-applications-delete.md`](cmd-pingcli-pingone-credentials-digital-wallet-applications-delete.md) |
| `pingcli pingone credentials digital-wallet-applications get` | Read a specific digital wallet application | [`cmd-pingcli-pingone-credentials-digital-wallet-applications-get.md`](cmd-pingcli-pingone-credentials-digital-wallet-applications-get.md) |
| `pingcli pingone credentials digital-wallet-applications list` | List all digital wallet applications | [`cmd-pingcli-pingone-credentials-digital-wallet-applications-list.md`](cmd-pingcli-pingone-credentials-digital-wallet-applications-list.md) |
| `pingcli pingone credentials digital-wallet-applications replace` | Replace a digital wallet application | [`cmd-pingcli-pingone-credentials-digital-wallet-applications-replace.md`](cmd-pingcli-pingone-credentials-digital-wallet-applications-replace.md) |
| `pingcli pingone credentials digital-wallet-applications template` | Generate a digital wallet application JSON template | [`cmd-pingcli-pingone-credentials-digital-wallet-applications-template.md`](cmd-pingcli-pingone-credentials-digital-wallet-applications-template.md) |

## Parent Command

- [`pingcli pingone credentials`](cmd-pingcli-pingone-credentials.md) — Administration tools for the PingOne Credentials universal service.
