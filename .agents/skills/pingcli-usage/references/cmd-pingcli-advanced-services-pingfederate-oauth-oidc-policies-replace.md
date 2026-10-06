# `pingcli advanced-services pingfederate oauth oidc policies replace`
Update an OAuth/OpenID Connect policy

## Synopsis

Update (replace) a PingFederate OAuth/OpenID Connect policy

```
pingcli advanced-services pingfederate oauth oidc policies replace [flags]
```

## Examples

```
# Update an OAuth/OpenID Connect policy from a JSON file (--id is still required)
  pingcli advanced-services pingfederate oauth oidc policies replace --id <id> --from-file policy.json

  # Update an OAuth/OpenID Connect policy from stdin
  pingcli advanced-services pingfederate oauth oidc policies replace --id <id> --from-file - < policy.json

  # Update from a JSON file, overriding the ID token lifetime
  pingcli advanced-services pingfederate oauth oidc policies replace --id <id> --from-file policy.json --id-token-lifetime 10
```

## Options

| Flag | Default | Description |
|------|---------|-------------|
| `-h, --help` | `` | help for replace |
| `-f, --from-file string` | `` | Path to a JSON file containing the request body, or "-" to read from stdin. |
| `--access-token-manager-id string` | `` | ID of the access token manager associated with this policy |
| `--allow-id-token-introspection` | `` | Allow the introspection endpoint to validate an ID token |
| `--id string` | `` | The PingFederate OAuth/OpenID Connect Policy ID |
| `--id-token-lifetime int64` | `` | ID Token lifetime, in minutes |
| `--id-token-typ-header-value string` | `` | ID Token Type (typ) header value |
| `--include-shash-in-id-token` | `` | Include the State Hash in the ID token |
| `--include-sri-in-id-token` | `` | Include a Session Reference Identifier in the ID token |
| `--include-user-info-in-id-token` | `` | Always include User Info in the ID token |
| `--include-x5t-in-id-token` | `` | Include the X.509 thumbprint header in the ID token |
| `--name string` | `` | Display name for the OIDC policy |
| `--reissue-id-token-in-hybrid-flow` | `` | Return a new ID Token during hybrid-flow token requests |
| `--return-id-token-on-refresh-grant` | `` | Return an ID Token when the refresh grant is requested |
| `--return-id-token-on-token-exchange-grant` | `` | Return an ID Token when token exchange is requested |


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

- [`pingcli advanced-services pingfederate oauth oidc policies`](cmd-pingcli-advanced-services-pingfederate-oauth-oidc-policies.md) — PingFederate OAuth/OpenID Connect Policies
