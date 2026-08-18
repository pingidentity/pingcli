# `pingcli pingone keys get`
Read a specific key

## Synopsis

Read a specific key in a PingOne environment. Use --format to retrieve the public certificate in PEM or PKCS7 format instead of the default JSON metadata. Use --output-file to write a PEM/PKCS7 response to a file instead of stdout.

```
pingcli pingone keys get [flags]
```

## Examples

```
# Read a specific key as JSON (default)
  pingcli pingone keys get --environment-id <env-id> --key-id <key-id>

  # Read the public certificate in PEM format
  pingcli pingone keys get --environment-id <env-id> --key-id <key-id> --format PEM

  # Read the public certificate in PKCS7 format
  pingcli pingone keys get --environment-id <env-id> --key-id <key-id> --format PKCS7

  # Write the PKCS7 certificate to a file instead of stdout
  pingcli pingone keys get --environment-id <env-id> --key-id <key-id> --format PKCS7 --output-file cert.p7b
```

## Options

| Flag | Default | Description |
|------|---------|-------------|
| `-h, --help` | `` | help for get |
| `-e, --environment-id string` | `` | The PingOne environment ID |
| `-k, --key-id string` | `` | The key ID |
| `-o, --output-file string` | `` | Write the response to PATH instead of stdout. Only valid with --format PEM or PKCS7. Overwrites any existing file. |
| `--format string` | `JSON` | Response format. Valid values: JSON, PEM, PKCS7. Defaults to JSON. |


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
