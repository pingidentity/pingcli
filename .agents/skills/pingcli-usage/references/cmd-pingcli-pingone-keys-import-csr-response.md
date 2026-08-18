# `pingcli pingone keys import-csr-response`
Import a CA-signed certificate response for a key

## Synopsis

Import a CA-signed certificate response for a PingOne key. The --file argument must be the CA-signed certificate in PEM or PKCS7 format. Use - to read from stdin.

```
pingcli pingone keys import-csr-response [flags]
```

## Examples

```
# Import a CA-signed certificate response for a key
  pingcli pingone keys import-csr-response --environment-id <env-id> --key-id <key-id> --file signed-cert.pem

  # Import from stdin
  pingcli pingone keys import-csr-response --environment-id <env-id> --key-id <key-id> --file -
```

## Options

| Flag | Default | Description |
|------|---------|-------------|
| `-h, --help` | `` | help for import-csr-response |
| `-e, --environment-id string` | `` | The PingOne environment ID |
| `-k, --key-id string` | `` | The key ID |
| `--file string` | `` | Path to the CA-signed certificate response file (PEM or PKCS7). Use - to read from stdin. |


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


## Parent Command

- [`pingcli pingone keys`](cmd-pingcli-pingone-keys.md) — Keys
