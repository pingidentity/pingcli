# `pingcli advanced-services pingfederate auth login`
Log in to allow Ping CLI to administer PingFederate (Advanced Services)

## Synopsis

Log in to allow Ping CLI to administer PingFederate (Advanced Services)

```
pingcli advanced-services pingfederate auth login [flags]
```

## Examples

```
# Log in to PingFederate (Advanced Services) using the configured authentication method.
    pingcli advanced-services pingfederate auth login
```

## Options

| Flag | Default | Description |
|------|---------|-------------|
| `-h, --help` | `` | help for login |
| `--as-pingfederate-admin-api-path string` | `` | The PingFederate admin API URL path for the Advanced Services private tenant. (default /pf-admin-api/v1) |
| `--as-pingfederate-ca-certificate-pem-files []string` | `` | Relative or full paths to PEM-encoded certificate files for the Advanced Services PingFederate server. (default []) Accepts a comma-separated string to delimit multiple PEM files. |
| `--as-pingfederate-https-host string` | `` | The PingFederate HTTPS host for the Advanced Services private tenant. Example: 'https://pingfederate-admin.bxretail.org' |
| `--as-pingfederate-insecure-trust-all-tls` | `` | Trust any certificate when connecting to the PingFederate Advanced Services admin API. (default false) This is insecure and shouldn't be enabled outside of testing. |
| `--as-pingfederate-software-version string` | `` | The PingFederate software version for the Advanced Services private tenant. (default 13.0) |
| `--as-pingfederate-x-bypass-external-validation-header` | `` | Bypass connection tests when configuring PingFederate in Advanced Services (the X-BypassExternalValidation header). (default false) |


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
| `--storage-type string` | `` | Auth token storage (default: secure_local)   secure_local  - Use OS keychain (default)   file_system   - Store tokens in ~/.pingcli/credentials   none          - Do not persist tokens |


## Parent Command

- [`pingcli advanced-services pingfederate auth`](cmd-pingcli-advanced-services-pingfederate-auth.md) — Authenticate Ping CLI to the PingFederate (Advanced Services) management APIs.
