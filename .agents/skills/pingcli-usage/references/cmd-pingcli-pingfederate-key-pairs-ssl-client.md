# `pingcli pingfederate key-pairs ssl-client`
PingFederate SSL Client Key Pairs

## Synopsis

PingFederate SSL Client Key Pairs

```
pingcli pingfederate key-pairs ssl-client [flags]
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
| `pingcli pingfederate key-pairs ssl-client create` | Generate a new SSL client key pair | [`cmd-pingcli-pingfederate-key-pairs-ssl-client-create.md`](cmd-pingcli-pingfederate-key-pairs-ssl-client-create.md) |
| `pingcli pingfederate key-pairs ssl-client delete` | Delete an SSL client key pair | [`cmd-pingcli-pingfederate-key-pairs-ssl-client-delete.md`](cmd-pingcli-pingfederate-key-pairs-ssl-client-delete.md) |
| `pingcli pingfederate key-pairs ssl-client export-certificate` | Export an SSL client key pair certificate file | [`cmd-pingcli-pingfederate-key-pairs-ssl-client-export-certificate.md`](cmd-pingcli-pingfederate-key-pairs-ssl-client-export-certificate.md) |
| `pingcli pingfederate key-pairs ssl-client export-pem` | Export an SSL client key pair PEM file | [`cmd-pingcli-pingfederate-key-pairs-ssl-client-export-pem.md`](cmd-pingcli-pingfederate-key-pairs-ssl-client-export-pem.md) |
| `pingcli pingfederate key-pairs ssl-client export-pkcs12` | Export an SSL client key pair PKCS12 file | [`cmd-pingcli-pingfederate-key-pairs-ssl-client-export-pkcs12.md`](cmd-pingcli-pingfederate-key-pairs-ssl-client-export-pkcs12.md) |
| `pingcli pingfederate key-pairs ssl-client generate-csr` | Generate a CSR for an SSL client key pair | [`cmd-pingcli-pingfederate-key-pairs-ssl-client-generate-csr.md`](cmd-pingcli-pingfederate-key-pairs-ssl-client-generate-csr.md) |
| `pingcli pingfederate key-pairs ssl-client get` | Read a specific SSL client key pair | [`cmd-pingcli-pingfederate-key-pairs-ssl-client-get.md`](cmd-pingcli-pingfederate-key-pairs-ssl-client-get.md) |
| `pingcli pingfederate key-pairs ssl-client import` | Import an SSL client key pair | [`cmd-pingcli-pingfederate-key-pairs-ssl-client-import.md`](cmd-pingcli-pingfederate-key-pairs-ssl-client-import.md) |
| `pingcli pingfederate key-pairs ssl-client import-csr-response` | Import an SSL client key pair CSR response | [`cmd-pingcli-pingfederate-key-pairs-ssl-client-import-csr-response.md`](cmd-pingcli-pingfederate-key-pairs-ssl-client-import-csr-response.md) |
| `pingcli pingfederate key-pairs ssl-client link` | Link an SSL client key pair private key and certificate (HSM only) | [`cmd-pingcli-pingfederate-key-pairs-ssl-client-link.md`](cmd-pingcli-pingfederate-key-pairs-ssl-client-link.md) |
| `pingcli pingfederate key-pairs ssl-client list` | List all SSL client key pairs | [`cmd-pingcli-pingfederate-key-pairs-ssl-client-list.md`](cmd-pingcli-pingfederate-key-pairs-ssl-client-list.md) |
| `pingcli pingfederate key-pairs ssl-client template` | Generate an SSL client key pair create, import, import-csr-response, or link JSON template | [`cmd-pingcli-pingfederate-key-pairs-ssl-client-template.md`](cmd-pingcli-pingfederate-key-pairs-ssl-client-template.md) |

## Parent Command

- [`pingcli pingfederate key-pairs`](cmd-pingcli-pingfederate-key-pairs.md) — Manage PingFederate Key Pairs resources
