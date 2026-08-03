# `pingcli pingfederate key-pairs ssl-client export-certificate`
Export an SSL client key pair certificate file

## Synopsis

Export the PEM-encoded certificate file contents of a PingFederate SSL client key pair

```
pingcli pingfederate key-pairs ssl-client export-certificate [flags]
```

## Examples

```
# Export the PEM certificate for an SSL client key pair
  pingcli pingfederate key-pairs ssl-client export-certificate --id <id>

  # Export the PEM certificate to a file
  pingcli pingfederate key-pairs ssl-client export-certificate --id <id> --output-file certificate.pem
```

## Options

| Flag | Default | Description |
|------|---------|-------------|
| `-h, --help` | `` | help for export-certificate |
| `-o, --output-file string` | `` | Optional path to write the raw result payload to, instead of only printing it to the terminal. Overwrites any existing file. |
| `--id string` | `` | The persistent, unique ID of the PingFederate SSL client key pair |


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

- [`pingcli pingfederate key-pairs ssl-client`](cmd-pingcli-pingfederate-key-pairs-ssl-client.md) — PingFederate SSL Client Key Pairs
