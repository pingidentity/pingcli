# `pingcli pingfederate secret-managers create`
Create a new secret manager

## Synopsis

Create a new PingFederate secret manager

```
pingcli pingfederate secret-managers create [flags]
```

## Examples

```
# Create a new secret manager from a JSON file
  pingcli pingfederate secret-managers create --from-file secret-manager.json

  # Create a new secret manager from stdin
  pingcli pingfederate secret-managers create --from-file - < secret-manager.json

  # Create from a JSON file, overriding the name and plugin descriptor reference
  pingcli pingfederate secret-managers create --from-file secret-manager.json --name "My Secret Manager" --plugin-descriptor-ref-id com.pingidentity.pf.secretmanagers.cyberark.CyberArkCredentialProvider
```

## Options

| Flag | Default | Description |
|------|---------|-------------|
| `-h, --help` | `` | help for create |
| `-f, --from-file string` | `` | Path to a JSON file containing the request body, or "-" to read from stdin. |
| `--id string` | `` | The PingFederate Secret Manager ID |
| `--name string` | `` | Secret manager instance name |
| `--parent-ref-id string` | `` | ID of a parent secret manager instance to inherit configuration from |
| `--plugin-descriptor-ref-id string` | `` | ID of the secret manager plugin type descriptor |


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

- [`pingcli pingfederate secret-managers`](cmd-pingcli-pingfederate-secret-managers.md) — PingFederate Secret Managers
