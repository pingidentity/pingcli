# `pingcli advanced-services pingfederate idp default-urls replace`
Update IdP default URLs

## Synopsis

Update (replace) the PingFederate IdP Default URLs

```
pingcli advanced-services pingfederate idp default-urls replace [flags]
```

## Examples

```
# Update IdP default URLs from a JSON file
  pingcli advanced-services pingfederate idp default-urls replace --from-file settings.json

  # Update IdP default URLs from stdin
  pingcli advanced-services pingfederate idp default-urls replace --from-file - < settings.json

  # Update IdP default URLs using flags only (no --from-file needed)
  pingcli advanced-services pingfederate idp default-urls replace --idp-error-msg "An error occurred" --idp-slo-success-url https://example.com/slo-success --confirm-idp-slo
```

## Options

| Flag | Default | Description |
|------|---------|-------------|
| `-h, --help` | `` | help for replace |
| `-f, --from-file string` | `` | Path to a JSON file containing the request body, or "-" to read from stdin. |
| `--confirm-idp-slo` | `` | Prompt user to confirm Single Logout (SLO) |
| `--idp-error-msg string` | `` | The error text displayed in a user's browser when an SSO operation fails |
| `--idp-slo-success-url string` | `` | The default URL to send the user to when Single Logout has succeeded |


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

- [`pingcli advanced-services pingfederate idp default-urls`](cmd-pingcli-advanced-services-pingfederate-idp-default-urls.md) — PingFederate IdP Default URLs
