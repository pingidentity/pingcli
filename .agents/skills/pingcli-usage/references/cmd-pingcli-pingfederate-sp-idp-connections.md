# `pingcli pingfederate sp idp-connections`
PingFederate SP IdP Connections

## Synopsis

PingFederate SP IdP Connections define the settings used to connect to an identity provider partner.

```
pingcli pingfederate sp idp-connections [flags]
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
| `pingcli pingfederate sp idp-connections apply` | Create or update an IdP connection | [`cmd-pingcli-pingfederate-sp-idp-connections-apply.md`](cmd-pingcli-pingfederate-sp-idp-connections-apply.md) |
| `pingcli pingfederate sp idp-connections certs` | PingFederate SP IdP connection certificates | [`cmd-pingcli-pingfederate-sp-idp-connections-certs.md`](cmd-pingcli-pingfederate-sp-idp-connections-certs.md) |
| `pingcli pingfederate sp idp-connections create` | Create a new IdP connection | [`cmd-pingcli-pingfederate-sp-idp-connections-create.md`](cmd-pingcli-pingfederate-sp-idp-connections-create.md) |
| `pingcli pingfederate sp idp-connections decryption-keys` | PingFederate SP IdP connection decryption keys | [`cmd-pingcli-pingfederate-sp-idp-connections-decryption-keys.md`](cmd-pingcli-pingfederate-sp-idp-connections-decryption-keys.md) |
| `pingcli pingfederate sp idp-connections delete` | Delete an IdP connection | [`cmd-pingcli-pingfederate-sp-idp-connections-delete.md`](cmd-pingcli-pingfederate-sp-idp-connections-delete.md) |
| `pingcli pingfederate sp idp-connections get` | Read a specific IdP connection | [`cmd-pingcli-pingfederate-sp-idp-connections-get.md`](cmd-pingcli-pingfederate-sp-idp-connections-get.md) |
| `pingcli pingfederate sp idp-connections list` | List all IdP connections | [`cmd-pingcli-pingfederate-sp-idp-connections-list.md`](cmd-pingcli-pingfederate-sp-idp-connections-list.md) |
| `pingcli pingfederate sp idp-connections replace` | Update an IdP connection | [`cmd-pingcli-pingfederate-sp-idp-connections-replace.md`](cmd-pingcli-pingfederate-sp-idp-connections-replace.md) |
| `pingcli pingfederate sp idp-connections signing-settings` | PingFederate SP IdP connection signing settings | [`cmd-pingcli-pingfederate-sp-idp-connections-signing-settings.md`](cmd-pingcli-pingfederate-sp-idp-connections-signing-settings.md) |
| `pingcli pingfederate sp idp-connections template` | Generate an IdP connection JSON template | [`cmd-pingcli-pingfederate-sp-idp-connections-template.md`](cmd-pingcli-pingfederate-sp-idp-connections-template.md) |

## Parent Command

- [`pingcli pingfederate sp`](cmd-pingcli-pingfederate-sp.md) — Manage PingFederate SP resources
