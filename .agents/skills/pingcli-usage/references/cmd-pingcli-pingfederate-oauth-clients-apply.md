# `pingcli pingfederate oauth clients apply`
Create or update an OAuth client

## Synopsis

Idempotently create or update a PingFederate OAuth client looked up by the "name" field in the JSON body. If no client with the given name exists it is created; if it exists it is updated.

```
pingcli pingfederate oauth clients apply [flags]
```

## Examples

```
# Create or update an OAuth client (body supplies clientId and other fields)
  pingcli pingfederate oauth clients apply --from-file oauth-client.json

  # Read body from stdin
  pingcli pingfederate oauth clients apply --from-file - < oauth-client.json

  # Create or update from a JSON file, overriding the description and redirect URIs
  pingcli pingfederate oauth clients apply --from-file oauth-client.json --description "Updated description" --redirect-uris https://example.com/callback
```

## Options

| Flag | Default | Description |
|------|---------|-------------|
| `-h, --help` | `` | help for apply |
| `-f, --from-file string` | `` | Path to a JSON file containing the request body, or "-" to read from stdin. |
| `--allow-authentication-api-init` | `` | Allow this client to initiate the Authentication API redirectless flow |
| `--authorization-detail-types []string` | `` | Authorization detail types available to this client; repeatable or comma-separated |
| `--bypass-activation-code-confirmation-override` | `` | Bypass the Activation Code Confirmation page for this client |
| `--bypass-approval-page` | `` | Skip the OAuth approval page for this client |
| `--ciba-delivery-mode string` | `` | CIBA token delivery mode for this client |
| `--ciba-notification-endpoint string` | `` | Endpoint the OP calls after a successful CIBA end-user authentication |
| `--ciba-polling-interval int64` | `` | Minimum interval, in seconds, the client must wait between CIBA token endpoint polling requests |
| `--ciba-request-object-signing-algorithm string` | `` | Signing algorithm required for the CIBA Request Object |
| `--ciba-require-signed-requests` | `` | Require signed requests for CIBA |
| `--ciba-user-code-supported` | `` | Whether this client supports a CIBA user code |
| `--client-id string` | `` | Unique client identifier presented to the resource server |
| `--client-secret-retention-period int64` | `` | Minutes secondary client secrets are retained after a secret change; 0 disables retention |
| `--client-secret-retention-period-type string` | `` | Overrides the global Client Secret Retention Period setting |
| `--default-access-token-manager-id string` | `` | ID of the access token manager used by default for this client |
| `--description string` | `` | Description of what the client application does |
| `--device-flow-setting-type string` | `` | Overrides the global Device Authorization Settings |
| `--device-polling-interval-override int64` | `` | Overrides the minimum interval, in seconds, the client must wait between device authorization polling requests |
| `--enable-cookieless-authentication-api` | `` | Allow the redirectless flow to function without relying on cookies |
| `--enabled` | `` | Whether the client is enabled |
| `--exclusive-scopes []string` | `` | Scopes exclusively available to this client; repeatable or comma-separated |
| `--grant-types []string` | `` | Grant types allowed for this client; repeatable or comma-separated |
| `--jwt-secured-authorization-response-mode-content-encryption-algorithm string` | `` | Content encryption algorithm used for the JWT Secured Authorization Response |
| `--jwt-secured-authorization-response-mode-encryption-algorithm string` | `` | Key management algorithm used to encrypt the JWT Secured Authorization Response |
| `--jwt-secured-authorization-response-mode-signing-algorithm string` | `` | Signing algorithm used for the JWT Secured Authorization Response |
| `--lockout-max-malicious-actions int64` | `` | Number of malicious actions allowed before this client is locked out |
| `--lockout-max-malicious-actions-type string` | `` | Overrides the global Max Malicious Actions Allowed setting |
| `--logo-url string` | `` | Location of the logo shown on OAuth grant and revocation pages |
| `--name string` | `` | Descriptive name shown to the user on the authorization page |
| `--offline-access-require-consent-prompt string` | `` | Overrides whether the offline_access scope requires the prompt parameter to contain consent |
| `--pending-authorization-timeout-override int64` | `` | Overrides the device code and user code timeout, in seconds |
| `--persistent-grant-expiration-time int64` | `` | Persistent grant expiration time; a negative value indicates the grant never expires |
| `--persistent-grant-expiration-time-unit string` | `` | Time unit for the persistent grant expiration time |
| `--persistent-grant-expiration-type string` | `` | Overrides the global Persistent Grant Lifetime setting |
| `--persistent-grant-idle-timeout int64` | `` | Persistent grant idle timeout; a negative value indicates no timeout |
| `--persistent-grant-idle-timeout-time-unit string` | `` | Time unit for the persistent grant idle timeout |
| `--persistent-grant-idle-timeout-type string` | `` | Overrides the global Persistent Grant Idle Timeout setting |
| `--persistent-grant-reuse-grant-types []string` | `` | Grant types the authorization server may reuse an existing grant for instead of creating a new one; repeatable or comma-separated |
| `--persistent-grant-reuse-type string` | `` | Overrides the Reuse Existing Persistent Access Grants for Grant Types setting |
| `--redirect-uris []string` | `` | Allowed OAuth redirect URIs for this client; repeatable or comma-separated |
| `--refresh-rolling string` | `` | Overrides the Roll Refresh Token Values setting for this client |
| `--refresh-token-rolling-grace-period int64` | `` | Grace period, in seconds, a rolled refresh token remains valid |
| `--refresh-token-rolling-grace-period-type string` | `` | Overrides the global Refresh Token Grace Period setting |
| `--refresh-token-rolling-interval int64` | `` | Minimum interval before a refresh token is rolled |
| `--refresh-token-rolling-interval-time-unit string` | `` | Time unit for the refresh token rolling interval |
| `--refresh-token-rolling-interval-type string` | `` | Overrides the Refresh Token Rolling Interval setting on the AS settings |
| `--request-object-signing-algorithm string` | `` | Signing algorithm required for the Request Object |
| `--request-policy-id string` | `` | ID of the OAuth request policy used by this client |
| `--require-dpop` | `` | Require Demonstrating Proof-of-Possession (DPoP) for this client |
| `--require-jwt-secured-authorization-response-mode` | `` | Require this client to use JWT Secured Authorization Response Mode (JARM) |
| `--require-offline-access-scope-to-issue-refresh-tokens string` | `` | Overrides whether the offline_access scope is required to issue a refresh token |
| `--require-proof-key-for-code-exchange` | `` | Require Proof Key for Code Exchange (PKCE) for this client |
| `--require-pushed-authorization-requests` | `` | Require this client to use pushed authorization requests |
| `--require-signed-requests` | `` | Require signed requests for this client |
| `--restrict-scopes` | `` | Restrict this client's access to a specific set of scopes |
| `--restrict-to-default-access-token-manager` | `` | Restrict this client to its default access token manager |
| `--restricted-response-types []string` | `` | Response types allowed for this client; repeatable or comma-separated |
| `--restricted-scopes []string` | `` | Scopes available to this client; repeatable or comma-separated |
| `--token-exchange-processor-policy-id string` | `` | ID of the token exchange processor policy used by this client |
| `--token-introspection-content-encryption-algorithm string` | `` | Content encryption algorithm used for the Token Introspection Response |
| `--token-introspection-encryption-algorithm string` | `` | Key management algorithm used to encrypt the Token Introspection Response |
| `--token-introspection-signing-algorithm string` | `` | Signing algorithm used for the Token Introspection Response |
| `--user-authorization-url-override string` | `` | Overrides the default authorization URL sent to the user for device authorization |
| `--validate-using-all-eligible-atms` | `` | Validate tokens issued by this client using all eligible access token managers |


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

- [`pingcli pingfederate oauth clients`](cmd-pingcli-pingfederate-oauth-clients.md) — PingFederate OAuth Clients
