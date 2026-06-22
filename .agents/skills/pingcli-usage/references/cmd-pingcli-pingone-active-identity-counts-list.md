# `pingcli pingone active-identity-counts list`
List active identity counts

## Synopsis

List all active identity counts in a PingOne environment, organized by date

```
pingcli pingone active-identity-counts list [flags]
```

## Examples

```
# List active identity counts from a start date (daily, default)
  pingcli pingone active-identity-counts list --environment-id <env-id> --start-date 2024-01-01T00:00:00Z

  # List with monthly aggregation
  pingcli pingone active-identity-counts list --environment-id <env-id> --start-date 2024-01-01T00:00:00Z --sample-period MONTH

  # List up to 100 daily records
  pingcli pingone active-identity-counts list --environment-id <env-id> --start-date 2024-01-01T00:00:00Z --limit 100
```

## Options

| Flag | Default | Description |
|------|---------|-------------|
| `-h, --help` | `` | help for list |
| `-e, --environment-id string` | `` | The PingOne environment ID |
| `--limit int64` | `` | Maximum number of results to return (1-100 for DAY, 1-24 for MONTH). Defaults to 10 if not set. |
| `--order string` | `` | Sort order for results. Valid values: ASC, DESC. |
| `--sample-period string` | `` | Sampling period for active identity count aggregation. Valid values: DAY, MONTH. |
| `--start-date string` | `` | Start date for the query in RFC3339 format (e.g. 2024-01-01T00:00:00Z). Must be within the past 2 years. Required. |
| `--template string` | `` | A Go text/template string. When provided, the command output is rendered through the template instead of the default format. The template receives the command's structured response data. Example: --template '{{.Name}}' |


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

- [`pingcli pingone active-identity-counts`](cmd-pingcli-pingone-active-identity-counts.md) — Active identity counts
