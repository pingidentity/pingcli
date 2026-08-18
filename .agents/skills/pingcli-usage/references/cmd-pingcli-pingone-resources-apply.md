# `pingcli pingone resources apply`
Create or update a resource

## Synopsis

Idempotently create or update a resource looked up by the "name" field in the JSON body within the supplied --environment-id. If no resource with the given name exists it is created; if exactly one exists it is updated; if more than one exists the command fails.

```
pingcli pingone resources apply [flags]
```

## Examples

```
# Create or update a resource (body supplies name and optional description)
  pingcli pingone resources apply --environment-id <env-id> --from-file resource.json

  # Read body from stdin
  pingcli pingone resources apply --environment-id <env-id> --from-file - < resource.json

  # Create or update from flags, without --from-file
  pingcli pingone resources apply --environment-id <env-id> --name "My Resource" --audience "https://api.example.com"
```

## Options

| Flag | Default | Description |
|------|---------|-------------|
| `-h, --help` | `` | help for apply |
| `-e, --environment-id string` | `` | The PingOne environment ID |
| `-f, --from-file string` | `` | Path to a JSON file containing the request body, or "-" to read from stdin. |
| `-r, --resource-id string` | `` | The resource ID |
| `--access-token-validity-seconds int64` | `` | The lifetime of access tokens issued for this resource, in seconds |
| `--audience string` | `` | The audience string included in access tokens issued for this resource |
| `--description string` | `` | The description of the resource |
| `--name string` | `` | The name of the resource |


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

- [`pingcli pingone resources`](cmd-pingcli-pingone-resources.md) — Resources
