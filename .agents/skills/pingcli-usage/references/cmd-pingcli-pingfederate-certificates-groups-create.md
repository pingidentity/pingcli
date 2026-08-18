# `pingcli pingfederate certificates groups create`
Import a new certificate into a certificate group

## Synopsis

Import a new certificate into a PingFederate certificate group

```
pingcli pingfederate certificates groups create [flags]
```

## Examples

```
# Import a new certificate into a certificate group from a JSON file
  pingcli pingfederate certificates groups create --group-name <group-name> --from-file cert.json

  # Import a new certificate into a certificate group from stdin
  pingcli pingfederate certificates groups create --group-name <group-name> --from-file - < cert.json

  # Import using --id, and --from-file for fileData
  pingcli pingfederate certificates groups create --group-name <group-name> --id my-cert --from-file cert.json
```

## Options

| Flag | Default | Description |
|------|---------|-------------|
| `-h, --help` | `` | help for create |
| `-f, --from-file string` | `` | Path to a JSON file containing the request body, or "-" to read from stdin. |
| `-g, --group-name string` | `` | The name of the certificate group |
| `-i, --id string` | `` | The certificate ID within the group |
| `--crypto-provider string` | `` | Cryptographic Provider name; only applicable if Hybrid HSM mode is enabled |


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

- [`pingcli pingfederate certificates groups`](cmd-pingcli-pingfederate-certificates-groups.md) — PingFederate certificate group certificates
