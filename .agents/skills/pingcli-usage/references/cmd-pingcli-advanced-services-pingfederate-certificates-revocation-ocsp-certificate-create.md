# `pingcli advanced-services pingfederate certificates revocation ocsp-certificate create`
Import a new OCSP certificate

## Synopsis

Import a new PingFederate OCSP certificate

```
pingcli advanced-services pingfederate certificates revocation ocsp-certificate create [flags]
```

## Examples

```
# Import a new OCSP certificate from a JSON file
  pingcli advanced-services pingfederate certificates revocation ocsp-certificate create --from-file certificate.json

  # Import a new OCSP certificate from stdin
  pingcli advanced-services pingfederate certificates revocation ocsp-certificate create --from-file - < certificate.json

  # Import using --id, and --from-file for fileData
  pingcli advanced-services pingfederate certificates revocation ocsp-certificate create --id my-ocsp-cert --from-file certificate.json
```

## Options

| Flag | Default | Description |
|------|---------|-------------|
| `-h, --help` | `` | help for create |
| `-f, --from-file string` | `` | Path to a JSON file containing the request body, or "-" to read from stdin. |
| `--crypto-provider string` | `` | Cryptographic Provider name; only applicable if Hybrid HSM mode is enabled |
| `--id string` | `` | The persistent, unique ID of the PingFederate OCSP certificate |


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

- [`pingcli advanced-services pingfederate certificates revocation ocsp-certificate`](cmd-pingcli-advanced-services-pingfederate-certificates-revocation-ocsp-certificate.md) — PingFederate OCSP Certificates
