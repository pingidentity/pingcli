# `pingcli advanced-services pingfederate key-pairs signing rotation-settings replace`
Update signing key pair rotation settings

## Synopsis

Update (replace) the rotationSettings singleton for a PingFederate signing key pair

```
pingcli advanced-services pingfederate key-pairs signing rotation-settings replace [flags]
```

## Examples

```
# Update rotation settings from a JSON file
  pingcli advanced-services pingfederate key-pairs signing rotation-settings replace --signing-key-pair-id <id> --from-file rotation-settings.json

  # Update rotation settings from stdin
  pingcli advanced-services pingfederate key-pairs signing rotation-settings replace --signing-key-pair-id <id> --from-file - < rotation-settings.json

  # Update rotation settings from flags, without --from-file
  pingcli advanced-services pingfederate key-pairs signing rotation-settings replace --signing-key-pair-id <id> --creation-buffer-days 5 --activation-buffer-days 2 --valid-days 365 --key-algorithm RSA --key-size 2048 --signature-algorithm SHA256withRSA
```

## Options

| Flag | Default | Description |
|------|---------|-------------|
| `-h, --help` | `` | help for replace |
| `-f, --from-file string` | `` | Path to a JSON file containing the request body, or "-" to read from stdin. |
| `--activation-buffer-days int64` | `` | Buffer days before key pair expiration for activation of the new key pair |
| `--creation-buffer-days int64` | `` | Buffer days before key pair expiration for creation of a new key pair |
| `--key-algorithm string` | `` | Key algorithm for the new key pair |
| `--key-size int64` | `` | Key size in bits for the new key pair |
| `--signature-algorithm string` | `` | Signature algorithm for the new key pair |
| `--signing-key-pair-id string` | `` | The persistent, unique ID of the parent PingFederate signing key pair |
| `--valid-days int64` | `` | Valid days for the new key pair |


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

- [`pingcli advanced-services pingfederate key-pairs signing rotation-settings`](cmd-pingcli-advanced-services-pingfederate-key-pairs-signing-rotation-settings.md) — PingFederate signing key pair rotation settings
