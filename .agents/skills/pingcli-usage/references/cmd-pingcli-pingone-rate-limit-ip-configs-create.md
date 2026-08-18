# `pingcli pingone rate-limit-ip-configs create`
Create a new rate limit IP configuration

## Synopsis

Create a new rate limit IP configuration in a PingOne environment

```
pingcli pingone rate-limit-ip-configs create [flags]
```

## Examples

```
# Create a new rate limit IP configuration from a JSON file
  pingcli pingone rate-limit-ip-configs create --environment-id <env-id> --from-file rate-limit-ip-config.json

  # Create a new rate limit IP configuration from stdin
  pingcli pingone rate-limit-ip-configs create --environment-id <env-id> --from-file - < rate-limit-ip-config.json

  # Create a new rate limit IP configuration from flags, without --from-file
  pingcli pingone rate-limit-ip-configs create --environment-id <env-id> --type WHITELIST --value 203.0.113.0/24
```

## Options

| Flag | Default | Description |
|------|---------|-------------|
| `-h, --help` | `` | help for create |
| `-e, --environment-id string` | `` | The PingOne environment ID |
| `-f, --from-file string` | `` | Path to a JSON file containing the request body, or "-" to read from stdin. |
| `--type string` | `` | The rate limit configuration type (e.g. WHITELIST) |
| `--value string` | `` | The IP address or CIDR block to exclude from rate limiting |


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

- [`pingcli pingone rate-limit-ip-configs`](cmd-pingcli-pingone-rate-limit-ip-configs.md) — Rate Limit IP Configurations
