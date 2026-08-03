# `pingcli pingfederate key-pairs signing`
PingFederate Signing Key Pairs

## Synopsis

PingFederate Signing Key Pairs

```
pingcli pingfederate key-pairs signing [flags]
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
| `pingcli pingfederate key-pairs signing create` | Generate a new signing key pair | [`cmd-pingcli-pingfederate-key-pairs-signing-create.md`](cmd-pingcli-pingfederate-key-pairs-signing-create.md) |
| `pingcli pingfederate key-pairs signing delete` | Delete a signing key pair | [`cmd-pingcli-pingfederate-key-pairs-signing-delete.md`](cmd-pingcli-pingfederate-key-pairs-signing-delete.md) |
| `pingcli pingfederate key-pairs signing export-certificate` | Export a signing key pair certificate file | [`cmd-pingcli-pingfederate-key-pairs-signing-export-certificate.md`](cmd-pingcli-pingfederate-key-pairs-signing-export-certificate.md) |
| `pingcli pingfederate key-pairs signing export-pem` | Export a signing key pair PEM file | [`cmd-pingcli-pingfederate-key-pairs-signing-export-pem.md`](cmd-pingcli-pingfederate-key-pairs-signing-export-pem.md) |
| `pingcli pingfederate key-pairs signing export-pkcs12` | Export a signing key pair PKCS12 file | [`cmd-pingcli-pingfederate-key-pairs-signing-export-pkcs12.md`](cmd-pingcli-pingfederate-key-pairs-signing-export-pkcs12.md) |
| `pingcli pingfederate key-pairs signing generate-csr` | Generate a CSR for a signing key pair | [`cmd-pingcli-pingfederate-key-pairs-signing-generate-csr.md`](cmd-pingcli-pingfederate-key-pairs-signing-generate-csr.md) |
| `pingcli pingfederate key-pairs signing get` | Read a specific signing key pair | [`cmd-pingcli-pingfederate-key-pairs-signing-get.md`](cmd-pingcli-pingfederate-key-pairs-signing-get.md) |
| `pingcli pingfederate key-pairs signing import` | Import a signing key pair | [`cmd-pingcli-pingfederate-key-pairs-signing-import.md`](cmd-pingcli-pingfederate-key-pairs-signing-import.md) |
| `pingcli pingfederate key-pairs signing import-csr-response` | Import a signing key pair CSR response | [`cmd-pingcli-pingfederate-key-pairs-signing-import-csr-response.md`](cmd-pingcli-pingfederate-key-pairs-signing-import-csr-response.md) |
| `pingcli pingfederate key-pairs signing link` | Link a signing key pair private key and certificate (HSM only) | [`cmd-pingcli-pingfederate-key-pairs-signing-link.md`](cmd-pingcli-pingfederate-key-pairs-signing-link.md) |
| `pingcli pingfederate key-pairs signing list` | List all signing key pairs | [`cmd-pingcli-pingfederate-key-pairs-signing-list.md`](cmd-pingcli-pingfederate-key-pairs-signing-list.md) |
| `pingcli pingfederate key-pairs signing rotation-settings` | PingFederate signing key pair rotation settings | [`cmd-pingcli-pingfederate-key-pairs-signing-rotation-settings.md`](cmd-pingcli-pingfederate-key-pairs-signing-rotation-settings.md) |
| `pingcli pingfederate key-pairs signing template` | Generate a signing key pair JSON template | [`cmd-pingcli-pingfederate-key-pairs-signing-template.md`](cmd-pingcli-pingfederate-key-pairs-signing-template.md) |

## Parent Command

- [`pingcli pingfederate key-pairs`](cmd-pingcli-pingfederate-key-pairs.md) — Manage PingFederate Key Pairs resources
