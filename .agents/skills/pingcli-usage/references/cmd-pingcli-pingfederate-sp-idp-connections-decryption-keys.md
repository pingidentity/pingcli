# `pingcli pingfederate sp idp-connections decryption-keys`
PingFederate SP IdP connection decryption keys

## Synopsis

Decryption keys used to decrypt message content received from the partner on a PingFederate SP IdP connection

```
pingcli pingfederate sp idp-connections decryption-keys [flags]
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
| `pingcli pingfederate sp idp-connections decryption-keys apply` | Update IdP connection decryption keys | [`cmd-pingcli-pingfederate-sp-idp-connections-decryption-keys-apply.md`](cmd-pingcli-pingfederate-sp-idp-connections-decryption-keys-apply.md) |
| `pingcli pingfederate sp idp-connections decryption-keys get` | Read IdP connection decryption keys | [`cmd-pingcli-pingfederate-sp-idp-connections-decryption-keys-get.md`](cmd-pingcli-pingfederate-sp-idp-connections-decryption-keys-get.md) |
| `pingcli pingfederate sp idp-connections decryption-keys replace` | Update IdP connection decryption keys | [`cmd-pingcli-pingfederate-sp-idp-connections-decryption-keys-replace.md`](cmd-pingcli-pingfederate-sp-idp-connections-decryption-keys-replace.md) |
| `pingcli pingfederate sp idp-connections decryption-keys template` | Generate an IdP connection decryption keys JSON template | [`cmd-pingcli-pingfederate-sp-idp-connections-decryption-keys-template.md`](cmd-pingcli-pingfederate-sp-idp-connections-decryption-keys-template.md) |

## Parent Command

- [`pingcli pingfederate sp idp-connections`](cmd-pingcli-pingfederate-sp-idp-connections.md) — PingFederate SP IdP Connections
