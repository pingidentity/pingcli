# `pingcli pingfederate idp adapters replace`
Update an IDP adapter

## Synopsis

Update (replace) a PingFederate IDP adapter

```
pingcli pingfederate idp adapters replace [flags]
```

## Examples

```
# Update an IDP adapter from a JSON file (--id is still required)
  pingcli pingfederate idp adapters replace --id <id> --from-file idp-adapter.json

  # Update an IDP adapter from stdin
  pingcli pingfederate idp adapters replace --id <id> --from-file - < idp-adapter.json

  # Update using flags for identity fields and --from-file for full configuration
  pingcli pingfederate idp adapters replace --id <id> --name "My IDP Adapter" --authn-ctx-class-ref <authn-ctx-class-ref> --from-file config.json
```

## Options

| Flag | Default | Description |
|------|---------|-------------|
| `-h, --help` | `` | help for replace |
| `-f, --from-file string` | `` | Path to a JSON file containing the request body, or "-" to read from stdin. |
| `--authn-ctx-class-ref string` | `` | The fixed value indicating how the user was authenticated |
| `--id string` | `` | The PingFederate IDP Adapter ID |
| `--name string` | `` | The plugin instance name for the IDP adapter |


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

- [`pingcli pingfederate idp adapters`](cmd-pingcli-pingfederate-idp-adapters.md) — PingFederate IDP Adapters
