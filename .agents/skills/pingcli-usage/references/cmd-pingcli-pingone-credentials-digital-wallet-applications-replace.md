# `pingcli pingone credentials digital-wallet-applications replace`
Replace a digital wallet application

## Synopsis

Replace (PUT) a digital wallet application in a PingOne environment

```
pingcli pingone credentials digital-wallet-applications replace [flags]
```

## Examples

```
# Replace a digital wallet application from a JSON file
  pingcli pingone credentials digital-wallet-applications replace --environment-id <env-id> --digital-wallet-application-id <dwa-id> --from-file digital-wallet-application.json

  # Replace a digital wallet application from stdin
  pingcli pingone credentials digital-wallet-applications replace --environment-id <env-id> --digital-wallet-application-id <dwa-id> --from-file - < digital-wallet-application.json
```

## Options

| Flag | Default | Description |
|------|---------|-------------|
| `-h, --help` | `` | help for replace |
| `-d, --digital-wallet-application-id string` | `` | The digital wallet application ID |
| `-e, --environment-id string` | `` | The PingOne environment ID |
| `-f, --from-file string` | `` | Path to a JSON file containing the request body, or "-" to read from stdin. |


## Inherited Options

| Flag | Default | Description |
|------|---------|-------------|
| `-C, --config string` | `` | The relative or full path to a custom Ping CLI configuration file. (default $HOME/.pingcli/config.yaml) |
| `-D, --detailed-exitcode` | `` | Enable detailed exit code output. (default false) 0 - pingcli command succeeded with no errors or warnings. 1 - pingcli command failed with errors. 2 - pingcli command succeeded with warnings. |
| `-O, --output-format string` | `` | Specify the console output format. (default text) Options are: json, ndjson, ndjson-wrapped, text. |
| `-P, --profile string` | `` | The name of a configuration profile to use. |
| `--debug` | `` | Enable debug output for error messages, including stack traces and transaction IDs. (default false) |
| `--log-file string` | `` | Write logs to a file at the given path. File logging is disabled when not set. |
| `--log-file-level string` | `` | Set the file log level. Options are: DEBUG, INFO, WARN, ERROR. (default DEBUG) |
| `--log-level string` | `` | Set the console log level. Options are: DEBUG, INFO, WARN, ERROR. (default WARN) |
| `--no-color` | `` | Disable text output in color. (default false) |
| `--query string` | `` | JMESPath expression to filter JSON output. Requires -O json, ndjson, or ndjson-wrapped. Example: --query 'data[?enabled].name' |


## Parent Command

- [`pingcli pingone credentials digital-wallet-applications`](cmd-pingcli-pingone-credentials-digital-wallet-applications.md) — Digital Wallet Applications
