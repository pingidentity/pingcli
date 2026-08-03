# `pingcli pingfederate idp sp-connections`
PingFederate SP Connections

## Synopsis

PingFederate SP Connections provide browser-based SSO, attribute query, and WS-Trust configuration for service provider partners.

```
pingcli pingfederate idp sp-connections [flags]
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
| `pingcli pingfederate idp sp-connections apply` | Create or update an SP connection | [`cmd-pingcli-pingfederate-idp-sp-connections-apply.md`](cmd-pingcli-pingfederate-idp-sp-connections-apply.md) |
| `pingcli pingfederate idp sp-connections certs` | PingFederate SP Connection certs | [`cmd-pingcli-pingfederate-idp-sp-connections-certs.md`](cmd-pingcli-pingfederate-idp-sp-connections-certs.md) |
| `pingcli pingfederate idp sp-connections create` | Create a new SP connection | [`cmd-pingcli-pingfederate-idp-sp-connections-create.md`](cmd-pingcli-pingfederate-idp-sp-connections-create.md) |
| `pingcli pingfederate idp sp-connections decryption-keys` | PingFederate SP Connection decryption keys | [`cmd-pingcli-pingfederate-idp-sp-connections-decryption-keys.md`](cmd-pingcli-pingfederate-idp-sp-connections-decryption-keys.md) |
| `pingcli pingfederate idp sp-connections delete` | Delete an SP connection | [`cmd-pingcli-pingfederate-idp-sp-connections-delete.md`](cmd-pingcli-pingfederate-idp-sp-connections-delete.md) |
| `pingcli pingfederate idp sp-connections get` | Read a specific SP connection | [`cmd-pingcli-pingfederate-idp-sp-connections-get.md`](cmd-pingcli-pingfederate-idp-sp-connections-get.md) |
| `pingcli pingfederate idp sp-connections list` | List all SP connections | [`cmd-pingcli-pingfederate-idp-sp-connections-list.md`](cmd-pingcli-pingfederate-idp-sp-connections-list.md) |
| `pingcli pingfederate idp sp-connections replace` | Update an SP connection | [`cmd-pingcli-pingfederate-idp-sp-connections-replace.md`](cmd-pingcli-pingfederate-idp-sp-connections-replace.md) |
| `pingcli pingfederate idp sp-connections signing-settings` | PingFederate SP Connection signing settings | [`cmd-pingcli-pingfederate-idp-sp-connections-signing-settings.md`](cmd-pingcli-pingfederate-idp-sp-connections-signing-settings.md) |
| `pingcli pingfederate idp sp-connections template` | Generate an SP connection JSON template | [`cmd-pingcli-pingfederate-idp-sp-connections-template.md`](cmd-pingcli-pingfederate-idp-sp-connections-template.md) |

## Parent Command

- [`pingcli pingfederate idp`](cmd-pingcli-pingfederate-idp.md) — Manage PingFederate IdP resources
