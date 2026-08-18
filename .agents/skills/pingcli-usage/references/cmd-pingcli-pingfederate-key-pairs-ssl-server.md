# `pingcli pingfederate key-pairs ssl-server`
PingFederate SSL Server Key Pairs

## Synopsis

PingFederate SSL Server Key Pairs

```
pingcli pingfederate key-pairs ssl-server [flags]
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
| `pingcli pingfederate key-pairs ssl-server create` | Generate a new SSL server key pair | [`cmd-pingcli-pingfederate-key-pairs-ssl-server-create.md`](cmd-pingcli-pingfederate-key-pairs-ssl-server-create.md) |
| `pingcli pingfederate key-pairs ssl-server delete` | Delete an SSL server key pair | [`cmd-pingcli-pingfederate-key-pairs-ssl-server-delete.md`](cmd-pingcli-pingfederate-key-pairs-ssl-server-delete.md) |
| `pingcli pingfederate key-pairs ssl-server export-certificate` | Export an SSL server key pair certificate file | [`cmd-pingcli-pingfederate-key-pairs-ssl-server-export-certificate.md`](cmd-pingcli-pingfederate-key-pairs-ssl-server-export-certificate.md) |
| `pingcli pingfederate key-pairs ssl-server export-pem` | Export an SSL server key pair PEM file | [`cmd-pingcli-pingfederate-key-pairs-ssl-server-export-pem.md`](cmd-pingcli-pingfederate-key-pairs-ssl-server-export-pem.md) |
| `pingcli pingfederate key-pairs ssl-server export-pkcs12` | Export an SSL server key pair PKCS12 file | [`cmd-pingcli-pingfederate-key-pairs-ssl-server-export-pkcs12.md`](cmd-pingcli-pingfederate-key-pairs-ssl-server-export-pkcs12.md) |
| `pingcli pingfederate key-pairs ssl-server generate-csr` | Generate a CSR for an SSL server key pair | [`cmd-pingcli-pingfederate-key-pairs-ssl-server-generate-csr.md`](cmd-pingcli-pingfederate-key-pairs-ssl-server-generate-csr.md) |
| `pingcli pingfederate key-pairs ssl-server get` | Read a specific SSL server key pair | [`cmd-pingcli-pingfederate-key-pairs-ssl-server-get.md`](cmd-pingcli-pingfederate-key-pairs-ssl-server-get.md) |
| `pingcli pingfederate key-pairs ssl-server import` | Import an SSL server key pair | [`cmd-pingcli-pingfederate-key-pairs-ssl-server-import.md`](cmd-pingcli-pingfederate-key-pairs-ssl-server-import.md) |
| `pingcli pingfederate key-pairs ssl-server import-csr-response` | Import an SSL server key pair CSR response | [`cmd-pingcli-pingfederate-key-pairs-ssl-server-import-csr-response.md`](cmd-pingcli-pingfederate-key-pairs-ssl-server-import-csr-response.md) |
| `pingcli pingfederate key-pairs ssl-server link` | Link an SSL server key pair private key and certificate (HSM only) | [`cmd-pingcli-pingfederate-key-pairs-ssl-server-link.md`](cmd-pingcli-pingfederate-key-pairs-ssl-server-link.md) |
| `pingcli pingfederate key-pairs ssl-server list` | List all SSL server key pairs | [`cmd-pingcli-pingfederate-key-pairs-ssl-server-list.md`](cmd-pingcli-pingfederate-key-pairs-ssl-server-list.md) |
| `pingcli pingfederate key-pairs ssl-server settings` | PingFederate SSL Server Settings | [`cmd-pingcli-pingfederate-key-pairs-ssl-server-settings.md`](cmd-pingcli-pingfederate-key-pairs-ssl-server-settings.md) |
| `pingcli pingfederate key-pairs ssl-server template` | Generate an SSL server key pair create, import, import-csr-response, or link JSON template | [`cmd-pingcli-pingfederate-key-pairs-ssl-server-template.md`](cmd-pingcli-pingfederate-key-pairs-ssl-server-template.md) |

## Parent Command

- [`pingcli pingfederate key-pairs`](cmd-pingcli-pingfederate-key-pairs.md) — Manage PingFederate Key Pairs resources
