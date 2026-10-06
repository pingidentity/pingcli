# `pingcli pingfederate oauth ciba-server-policy request-policies create`
Create a new CIBA request policy

## Synopsis

Create a new PingFederate CIBA request policy

```
pingcli pingfederate oauth ciba-server-policy request-policies create [flags]
```

## Examples

```
# Create a new CIBA request policy from a JSON file
  pingcli pingfederate oauth ciba-server-policy request-policies create --from-file request-policy.json

  # Create a new CIBA request policy from stdin
  pingcli pingfederate oauth ciba-server-policy request-policies create --from-file - < request-policy.json

  # Create from a JSON file, overriding the name and authenticator
  pingcli pingfederate oauth ciba-server-policy request-policies create --from-file request-policy.json --name "CIBA Policy" --authenticator-ref-id myAuthenticator
```

## Options

| Flag | Default | Description |
|------|---------|-------------|
| `-h, --help` | `` | help for create |
| `-f, --from-file string` | `` | Path to a JSON file containing the request body, or "-" to read from stdin. |
| `--allow-unsigned-login-hint-token` | `` | Allow an unsigned login hint token |
| `--authenticator-ref-id string` | `` | ID of the out-of-band authenticator this request policy uses |
| `--id string` | `` | The PingFederate CIBA request policy ID |
| `--name string` | `` | The CIBA request policy name |
| `--require-token-for-identity-hint` | `` | Require a token for the identity hint |
| `--transaction-lifetime int64` | `` | The transaction lifetime in seconds |
| `--user-code-pcv-ref-id string` | `` | ID of the password credential validator used to validate the user code |


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

- [`pingcli pingfederate oauth ciba-server-policy request-policies`](cmd-pingcli-pingfederate-oauth-ciba-server-policy-request-policies.md) — PingFederate OAuth CIBA Server Policy Request Policies
