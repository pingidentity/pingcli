# `pingcli pingfederate authentication-api applications apply`
Create or update an authentication API application

## Synopsis

Idempotently create or update a PingFederate authentication API application looked up by the "name" field in the JSON body. If no application with the given name exists it is created; if it exists it is updated.

```
pingcli pingfederate authentication-api applications apply [flags]
```

## Examples

```
# Create or update an authentication API application (body supplies name and other fields)
  pingcli pingfederate authentication-api applications apply --from-file application.json

  # Read body from stdin
  pingcli pingfederate authentication-api applications apply --from-file - < application.json

  # Create or update an authentication API application with flags
  pingcli pingfederate authentication-api applications apply --url https://example.com/callback --from-file application.json
```

## Options

| Flag | Default | Description |
|------|---------|-------------|
| `-h, --help` | `` | help for apply |
| `-f, --from-file string` | `` | Path to a JSON file containing the request body, or "-" to read from stdin. |
| `--additional-allowed-origins []string` | `` | Additional allowed CORS origin URLs beyond the domain in the redirect URL; repeatable or comma-separated |
| `--client-for-redirectless-mode-ref-id string` | `` | ID of the OAuth client to use in redirectless mode |
| `--description string` | `` | Description of the application |
| `--id string` | `` | The PingFederate Authentication API Application ID |
| `--name string` | `` | Authentication API application name |
| `--url string` | `` | Redirect URL for the application |


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

- [`pingcli pingfederate authentication-api applications`](cmd-pingcli-pingfederate-authentication-api-applications.md) — PingFederate Authentication API Applications
