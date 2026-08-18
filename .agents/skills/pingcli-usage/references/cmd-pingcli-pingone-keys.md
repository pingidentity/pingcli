# `pingcli pingone keys`
Keys

## Synopsis

Keys

```
pingcli pingone keys [flags]
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
| `pingcli pingone keys applications` | Key Applications | [`cmd-pingcli-pingone-keys-applications.md`](cmd-pingcli-pingone-keys-applications.md) |
| `pingcli pingone keys apply` | Create or update a key | [`cmd-pingcli-pingone-keys-apply.md`](cmd-pingcli-pingone-keys-apply.md) |
| `pingcli pingone keys create` | Create a new key | [`cmd-pingcli-pingone-keys-create.md`](cmd-pingcli-pingone-keys-create.md) |
| `pingcli pingone keys create-with-pkcs12` | Create a key pair from a PKCS12 file | [`cmd-pingcli-pingone-keys-create-with-pkcs12.md`](cmd-pingcli-pingone-keys-create-with-pkcs12.md) |
| `pingcli pingone keys delete` | Delete a key | [`cmd-pingcli-pingone-keys-delete.md`](cmd-pingcli-pingone-keys-delete.md) |
| `pingcli pingone keys get` | Read a specific key | [`cmd-pingcli-pingone-keys-get.md`](cmd-pingcli-pingone-keys-get.md) |
| `pingcli pingone keys get-csr` | Export the CSR for a key | [`cmd-pingcli-pingone-keys-get-csr.md`](cmd-pingcli-pingone-keys-get-csr.md) |
| `pingcli pingone keys import-csr-response` | Import a CA-signed certificate response for a key | [`cmd-pingcli-pingone-keys-import-csr-response.md`](cmd-pingcli-pingone-keys-import-csr-response.md) |
| `pingcli pingone keys list` | List all keys | [`cmd-pingcli-pingone-keys-list.md`](cmd-pingcli-pingone-keys-list.md) |
| `pingcli pingone keys replace` | Update a key | [`cmd-pingcli-pingone-keys-replace.md`](cmd-pingcli-pingone-keys-replace.md) |
| `pingcli pingone keys template` | Generate a key JSON template | [`cmd-pingcli-pingone-keys-template.md`](cmd-pingcli-pingone-keys-template.md) |

## Parent Command

- [`pingcli pingone`](cmd-pingcli-pingone.md) — Administration tools for the PingOne platform.
