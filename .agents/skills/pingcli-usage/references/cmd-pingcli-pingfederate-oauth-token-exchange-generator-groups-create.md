# `pingcli pingfederate oauth token-exchange generator groups create`
Create a new token exchange generator group

## Synopsis

Create a new PingFederate Token Exchange generator group

```
pingcli pingfederate oauth token-exchange generator groups create [flags]
```

## Examples

```
# Create a new token exchange generator group from a JSON file
  pingcli pingfederate oauth token-exchange generator groups create --from-file group.json

  # Create a new token exchange generator group from stdin
  pingcli pingfederate oauth token-exchange generator groups create --from-file - < group.json

  # Create using flags for identity fields, and --from-file for the generator mappings
  pingcli pingfederate oauth token-exchange generator groups create --name "My Token Exchange Generator Group" --from-file group.json
```

## Options

| Flag | Default | Description |
|------|---------|-------------|
| `-h, --help` | `` | help for create |
| `-f, --from-file string` | `` | Path to a JSON file containing the request body, or "-" to read from stdin. |
| `--id string` | `` | The PingFederate Token Exchange generator group ID |
| `--name string` | `` | The Token Exchange generator group name |
| `--resource-uris []string` | `` | The list of resource URIs which map to this Token Exchange generator group |


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

- [`pingcli pingfederate oauth token-exchange generator groups`](cmd-pingcli-pingfederate-oauth-token-exchange-generator-groups.md) — PingFederate OAuth 2.0 Token Exchange generator groups
