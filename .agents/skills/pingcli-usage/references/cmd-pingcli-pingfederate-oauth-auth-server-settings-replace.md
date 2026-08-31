# `pingcli pingfederate oauth auth-server-settings replace`
Update oauth authorization server settings

## Synopsis

Update (replace) the PingFederate OAuth Authorization Server Settings

```
pingcli pingfederate oauth auth-server-settings replace [flags]
```

## Examples

```
# Update oauth authorization server settings from a JSON file
  pingcli pingfederate oauth auth-server-settings replace --from-file settings.json

  # Update oauth authorization server settings from stdin
  pingcli pingfederate oauth auth-server-settings replace --from-file - < settings.json

  # Update oauth authorization server settings using flags only (no --from-file needed)
  pingcli pingfederate oauth auth-server-settings replace --authorization-code-timeout 60 --authorization-code-entropy 30 --refresh-token-length 42 --refresh-rolling-interval 1 --include-issuer-in-authorization-response=false
```

## Options

| Flag | Default | Description |
|------|---------|-------------|
| `-h, --help` | `` | help for replace |
| `-f, --from-file string` | `` | Path to a JSON file containing the request body, or "-" to read from stdin. |
| `--activation-code-check-mode string` | `` | Activation code check mode |
| `--admin-web-service-pcv-ref-id string` | `` | ID of the admin web service PCV |
| `--allow-unidentified-client-extension-grants` | `` | Whether unidentified clients may use extension grants |
| `--allow-unidentified-client-ro-creds` | `` | Whether unidentified clients may use resource owner credentials |
| `--allowed-origins []string` | `` | Allowed origins |
| `--approved-authorization-detail-attribute string` | `` | Approved authorization detail attribute |
| `--approved-scopes-attribute string` | `` | Approved scopes attribute |
| `--atm-id-for-oauth-grant-management string` | `` | Access Token Manager ID for grant management |
| `--authorization-code-entropy int64` | `` | Authorization code entropy |
| `--authorization-code-timeout int64` | `` | Authorization code timeout |
| `--bypass-activation-code-confirmation` | `` | Whether activation code confirmation is bypassed |
| `--bypass-authorization-for-approved-consents` | `` | Whether authorization is bypassed for approved consents |
| `--bypass-authorization-for-approved-grants` | `` | Whether authorization is bypassed for approved grants |
| `--client-secret-retention-period int64` | `` | Client secret retention period |
| `--consent-lifetime-days int64` | `` | Consent lifetime in days |
| `--default-scope-description string` | `` | Default scope description |
| `--device-polling-interval int64` | `` | Device polling interval |
| `--disallow-plain-pkce` | `` | Whether plain PKCE is disallowed |
| `--dpop-proof-enforce-replay-prevention` | `` | Whether DPoP replay prevention is enforced |
| `--dpop-proof-lifetime-seconds int64` | `` | DPoP proof lifetime |
| `--dpop-proof-require-nonce` | `` | Whether DPoP proofs require a nonce |
| `--enable-cookieless-user-authorization-authentication-api` | `` | Whether cookieless user authorization authentication API is enabled |
| `--include-issuer-in-authorization-response` | `` | Whether the issuer is included in authorization responses |
| `--jwt-secured-authorization-response-mode-lifetime int64` | `` | JWT secured authorization response mode lifetime |
| `--offline-access-require-consent-prompt` | `` | Whether offline access requires consent |
| `--par-reference-length int64` | `` | Pushed authorization request reference length |
| `--par-reference-timeout int64` | `` | Pushed authorization request reference timeout |
| `--par-status string` | `` | Pushed authorization request status |
| `--pending-authorization-timeout int64` | `` | Pending authorization timeout |
| `--persistent-grant-idle-timeout int64` | `` | Persistent grant idle timeout |
| `--persistent-grant-idle-timeout-time-unit string` | `` | Persistent grant idle timeout unit |
| `--persistent-grant-lifetime int64` | `` | Persistent grant lifetime |
| `--persistent-grant-lifetime-unit string` | `` | Persistent grant lifetime unit |
| `--persistent-grant-reuse-grant-types []string` | `` | Grant types reusable by persistent grants |
| `--refresh-rolling-interval int64` | `` | Refresh token rolling interval |
| `--refresh-rolling-interval-time-unit string` | `` | Refresh token rolling interval unit |
| `--refresh-token-length int64` | `` | Refresh token length |
| `--refresh-token-rolling-grace-period int64` | `` | Refresh token rolling grace period |
| `--registered-authorization-path string` | `` | Registered authorization path |
| `--require-offline-access-scope-to-issue-refresh-tokens` | `` | Whether offline access scope is required for refresh tokens |
| `--return-id-token-on-open-id-with-device-authz-grant` | `` | Whether an ID token is returned for device authorization |
| `--roll-refresh-token-values` | `` | Whether refresh token values roll |
| `--scope-for-oauth-grant-management string` | `` | OAuth scope for grant management |
| `--token-endpoint-base-url string` | `` | Token endpoint base URL |
| `--track-user-sessions-for-logout` | `` | Whether user sessions are tracked for logout |
| `--user-authorization-consent-adapter string` | `` | User authorization consent adapter |
| `--user-authorization-consent-page-setting string` | `` | User authorization consent page setting |
| `--user-authorization-url string` | `` | User authorization URL |


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

- [`pingcli pingfederate oauth auth-server-settings`](cmd-pingcli-pingfederate-oauth-auth-server-settings.md) — PingFederate OAuth Authorization Server Settings
