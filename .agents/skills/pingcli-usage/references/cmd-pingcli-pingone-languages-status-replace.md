# `pingcli pingone languages status replace`
Update a language localization status

## Synopsis

Update (replace) a language localization status in a PingOne environment

```
pingcli pingone languages status replace [flags]
```

## Examples

```
# Update a language localization status from a JSON file
  pingcli pingone languages status replace --environment-id <env-id> --language-id <language-id> --localization-status-id <status-id> --from-file status.json

  # Update a language localization status from stdin
  pingcli pingone languages status replace --environment-id <env-id> --language-id <language-id> --localization-status-id <status-id> --from-file - < status.json

  # Update from a JSON file, overriding the localization-complete state
  pingcli pingone languages status replace --environment-id <env-id> --language-id <language-id> --localization-status-id <status-id> --from-file status.json --localization-complete=false
```

## Options

| Flag | Default | Description |
|------|---------|-------------|
| `-h, --help` | `` | help for replace |
| `-e, --environment-id string` | `` | The PingOne environment ID |
| `-f, --from-file string` | `` | Path to a JSON file containing the request body, or "-" to read from stdin. |
| `-l, --language-id string` | `` | The language ID |
| `-s, --localization-status-id string` | `` | The language localization status ID |
| `--localization-complete` | `` | Whether localization for this service is complete |
| `--service string` | `` | The PingOne service whose localization status is being recorded |
| `--status-details string` | `` | Additional detail about the localization status |


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

- [`pingcli pingone languages status`](cmd-pingcli-pingone-languages-status.md) — Language Localization Statuses
