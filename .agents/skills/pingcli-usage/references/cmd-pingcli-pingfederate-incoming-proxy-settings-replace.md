# `pingcli pingfederate incoming-proxy-settings replace`
Update incoming proxy settings

## Synopsis

Update (replace) the PingFederate incoming proxy settings

```
pingcli pingfederate incoming-proxy-settings replace [flags]
```

## Examples

```
# Update incoming proxy settings from a JSON file
  pingcli pingfederate incoming-proxy-settings replace --from-file settings.json

  # Update incoming proxy settings from stdin
  pingcli pingfederate incoming-proxy-settings replace --from-file - < settings.json

  # Update incoming proxy settings from flags, without --from-file
  pingcli pingfederate incoming-proxy-settings replace --forwarded-ip-address-header-name X-Forwarded-For --forwarded-ip-address-header-index FIRST --forwarded-host-header-name X-Forwarded-Host --forwarded-host-header-index FIRST
```

## Options

| Flag | Default | Description |
|------|---------|-------------|
| `-h, --help` | `` | help for replace |
| `-f, --from-file string` | `` | Path to a JSON file containing the request body, or "-" to read from stdin. |
| `--client-cert-chain-ssl-header-name string` | `` | Header name containing the client certificate chain |
| `--client-cert-header-encoding-format string` | `` | Encoding format of the client certificate header |
| `--client-cert-ssl-header-name string` | `` | Header name containing the client certificate |
| `--enable-client-cert-header-auth` | `` | Enable client certificate header authentication |
| `--forwarded-host-header-index string` | `` | Which forwarded hostname to use when multiple comma-separated values are present |
| `--forwarded-host-header-name string` | `` | Header name containing the hostname and port forwarded by the proxy |
| `--forwarded-ip-address-header-index string` | `` | Which forwarded client IP address to use when multiple comma-separated values are present |
| `--forwarded-ip-address-header-name string` | `` | Header name containing the client IP address forwarded by the proxy |
| `--proxy-terminates-https-conns` | `` | Treat connections to the reverse proxy as HTTPS connections |


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

- [`pingcli pingfederate incoming-proxy-settings`](cmd-pingcli-pingfederate-incoming-proxy-settings.md) — PingFederate incoming proxy settings
