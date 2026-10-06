# `pingcli advanced-services pingfederate key-pairs oauth-openid-connect replace`
Update OAuth/OpenID Connect keys settings

## Synopsis

Update (replace) the PingFederate OAuth/OpenID Connect keys settings

```
pingcli advanced-services pingfederate key-pairs oauth-openid-connect replace [flags]
```

## Examples

```
# Update OAuth/OpenID Connect keys settings from a JSON file
  pingcli advanced-services pingfederate key-pairs oauth-openid-connect replace --from-file settings.json

  # Update OAuth/OpenID Connect keys settings from stdin
  pingcli advanced-services pingfederate key-pairs oauth-openid-connect replace --from-file - < settings.json

  # Update OAuth/OpenID Connect keys settings using flags only (no --from-file needed)
  pingcli advanced-services pingfederate key-pairs oauth-openid-connect replace --static-jwks-enabled=false
```

## Options

| Flag | Default | Description |
|------|---------|-------------|
| `-h, --help` | `` | help for replace |
| `-f, --from-file string` | `` | Path to a JSON file containing the request body, or "-" to read from stdin. |
| `--p256-active-cert-ref-id string` | `` | ID of the active P-256 certificate |
| `--p256-active-key-id string` | `` | ID of the active P-256 key |
| `--p256-decryption-active-cert-ref-id string` | `` | ID of the active P-256 decryption certificate |
| `--p256-decryption-active-key-id string` | `` | ID of the active P-256 decryption key |
| `--p256-decryption-previous-cert-ref-id string` | `` | ID of the previous P-256 decryption certificate |
| `--p256-decryption-previous-key-id string` | `` | ID of the previous P-256 decryption key |
| `--p256-decryption-publish-x5c-parameter` | `` | Whether to publish the P-256 decryption certificate chain |
| `--p256-previous-cert-ref-id string` | `` | ID of the previous P-256 certificate |
| `--p256-previous-key-id string` | `` | ID of the previous P-256 key |
| `--p256-publish-x5c-parameter` | `` | Whether to publish the P-256 certificate chain |
| `--p384-active-cert-ref-id string` | `` | ID of the active P-384 certificate |
| `--p384-active-key-id string` | `` | ID of the active P-384 key |
| `--p384-decryption-active-cert-ref-id string` | `` | ID of the active P-384 decryption certificate |
| `--p384-decryption-active-key-id string` | `` | ID of the active P-384 decryption key |
| `--p384-decryption-previous-cert-ref-id string` | `` | ID of the previous P-384 decryption certificate |
| `--p384-decryption-previous-key-id string` | `` | ID of the previous P-384 decryption key |
| `--p384-decryption-publish-x5c-parameter` | `` | Whether to publish the P-384 decryption certificate chain |
| `--p384-previous-cert-ref-id string` | `` | ID of the previous P-384 certificate |
| `--p384-previous-key-id string` | `` | ID of the previous P-384 key |
| `--p384-publish-x5c-parameter` | `` | Whether to publish the P-384 certificate chain |
| `--p521-active-cert-ref-id string` | `` | ID of the active P-521 certificate |
| `--p521-active-key-id string` | `` | ID of the active P-521 key |
| `--p521-decryption-active-cert-ref-id string` | `` | ID of the active P-521 decryption certificate |
| `--p521-decryption-active-key-id string` | `` | ID of the active P-521 decryption key |
| `--p521-decryption-previous-cert-ref-id string` | `` | ID of the previous P-521 decryption certificate |
| `--p521-decryption-previous-key-id string` | `` | ID of the previous P-521 decryption key |
| `--p521-decryption-publish-x5c-parameter` | `` | Whether to publish the P-521 decryption certificate chain |
| `--p521-previous-cert-ref-id string` | `` | ID of the previous P-521 certificate |
| `--p521-previous-key-id string` | `` | ID of the previous P-521 key |
| `--p521-publish-x5c-parameter` | `` | Whether to publish the P-521 certificate chain |
| `--publish-dynamic-key-x5cs` | `` | Whether to publish dynamic key certificate chains |
| `--rsa-active-cert-ref-id string` | `` | ID of the active RSA certificate |
| `--rsa-active-key-id string` | `` | ID of the active RSA key |
| `--rsa-decryption-active-cert-ref-id string` | `` | ID of the active RSA decryption certificate |
| `--rsa-decryption-active-key-id string` | `` | ID of the active RSA decryption key |
| `--rsa-decryption-previous-cert-ref-id string` | `` | ID of the previous RSA decryption certificate |
| `--rsa-decryption-previous-key-id string` | `` | ID of the previous RSA decryption key |
| `--rsa-decryption-publish-x5c-parameter` | `` | Whether to publish the RSA decryption certificate chain |
| `--rsa-previous-cert-ref-id string` | `` | ID of the previous RSA certificate |
| `--rsa-previous-key-id string` | `` | ID of the previous RSA key |
| `--rsa-publish-x5c-parameter` | `` | Whether to publish the RSA certificate chain |
| `--static-jwks-enabled` | `` | Whether static keys are enabled |


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

- [`pingcli advanced-services pingfederate key-pairs oauth-openid-connect`](cmd-pingcli-advanced-services-pingfederate-key-pairs-oauth-openid-connect.md) — PingFederate OAuth/OpenID Connect keys settings
