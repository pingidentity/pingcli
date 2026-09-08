# `pingcli pingfederate server-settings ws-trust-sts-settings apply`
Update WS-Trust STS Settings

## Synopsis

Idempotently update the WS-Trust STS Settings. Apply is an alias for replace on this singleton resource.

```
pingcli pingfederate server-settings ws-trust-sts-settings apply [flags]
```

## Examples

```
# Update WS-Trust STS Settings from a JSON file
  pingcli pingfederate server-settings ws-trust-sts-settings apply --from-file ws-trust-sts-settings.json

  # Update WS-Trust STS Settings from stdin
  pingcli pingfederate server-settings ws-trust-sts-settings apply --from-file - < ws-trust-sts-settings.json

  # Update WS-Trust STS Settings from flags, without --from-file
  pingcli pingfederate server-settings ws-trust-sts-settings apply --basic-authn-enabled=false --subject-dns "CN=sts-client" --client-cert-authn-enabled --restrict-by-subject-dn --restrict-by-issuer-cert=false
```

## Options

| Flag | Default | Description |
|------|---------|-------------|
| `-h, --help` | `` | help for apply |
| `-f, --from-file string` | `` | Path to a JSON file containing the request body, or "-" to read from stdin. |
| `--basic-authn-enabled` | `` | Require the use of HTTP Basic Authentication to access WS-Trust STS endpoints. Requires users be populated. |
| `--client-cert-authn-enabled` | `` | Require the use of Client Cert Authentication to access WS-Trust STS endpoints. Requires either restrictBySubjectDn and/or restrictByIssuerCert be enabled. |
| `--restrict-by-issuer-cert` | `` | Restrict access by issuer certificate. Ignored if clientCertAuthnEnabled is disabled. |
| `--restrict-by-subject-dn` | `` | Restrict access by Subject DN. Ignored if clientCertAuthnEnabled is disabled. |
| `--subject-dns []string` | `` | List of Subject DNs for certificates that are allowed to authenticate to WS-Trust STS endpoints. Required if restrictBySubjectDn is enabled. |


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

- [`pingcli pingfederate server-settings ws-trust-sts-settings`](cmd-pingcli-pingfederate-server-settings-ws-trust-sts-settings.md) — PingFederate WS-Trust STS Settings
