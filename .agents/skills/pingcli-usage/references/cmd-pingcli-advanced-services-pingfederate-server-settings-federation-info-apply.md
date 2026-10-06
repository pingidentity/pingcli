# `pingcli advanced-services pingfederate server-settings federation-info apply`
Update federation info

## Synopsis

Idempotently update the PingFederate Federation Info. Apply is an alias for replace on this singleton resource.

```
pingcli advanced-services pingfederate server-settings federation-info apply [flags]
```

## Examples

```
# Update federation info from a JSON file
  pingcli advanced-services pingfederate server-settings federation-info apply --from-file federation-info.json

  # Update federation info from stdin
  pingcli advanced-services pingfederate server-settings federation-info apply --from-file - < federation-info.json

  # Update federation info from flags, without --from-file
  pingcli advanced-services pingfederate server-settings federation-info apply --base-url https://fed.example.com --saml2-entity-id https://fed.example.com/saml
```

## Options

| Flag | Default | Description |
|------|---------|-------------|
| `-h, --help` | `` | help for apply |
| `-f, --from-file string` | `` | Path to a JSON file containing the request body, or "-" to read from stdin. |
| `--base-url string` | `` | Fully qualified host, port, and path on which PingFederate runs |
| `--saml1x-issuer-id string` | `` | SAML 1.x issuer ID for this federation server |
| `--saml1x-source-id string` | `` | SAML 1.x Source ID override (derived from the issuer ID if omitted) |
| `--saml2-entity-id string` | `` | SAML 2.0 entity ID for this federation server |
| `--wsfed-realm string` | `` | URI of the WS-Federation realm associated with this server |


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

- [`pingcli advanced-services pingfederate server-settings federation-info`](cmd-pingcli-advanced-services-pingfederate-server-settings-federation-info.md) — PingFederate Federation Info
