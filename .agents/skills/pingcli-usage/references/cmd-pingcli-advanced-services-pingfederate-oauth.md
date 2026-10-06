# `pingcli advanced-services pingfederate oauth`
Manage PingFederate OAuth resources

## Synopsis

Manage PingFederate OAuth resources

```
pingcli advanced-services pingfederate oauth
```

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


## Subcommands

| Command | Description | Reference |
|---------|-------------|----------|
| `pingcli advanced-services pingfederate oauth access-token-managers` | PingFederate OAuth access token managers | [`cmd-pingcli-advanced-services-pingfederate-oauth-access-token-managers.md`](cmd-pingcli-advanced-services-pingfederate-oauth-access-token-managers.md) |
| `pingcli advanced-services pingfederate oauth access-token-mappings` | PingFederate OAuth access token mappings | [`cmd-pingcli-advanced-services-pingfederate-oauth-access-token-mappings.md`](cmd-pingcli-advanced-services-pingfederate-oauth-access-token-mappings.md) |
| `pingcli advanced-services pingfederate oauth auth-server-settings` | PingFederate OAuth Authorization Server Settings | [`cmd-pingcli-advanced-services-pingfederate-oauth-auth-server-settings.md`](cmd-pingcli-advanced-services-pingfederate-oauth-auth-server-settings.md) |
| `pingcli advanced-services pingfederate oauth authentication-policy-contract-mappings` | PingFederate OAuth Authentication Policy Contract Mappings | [`cmd-pingcli-advanced-services-pingfederate-oauth-authentication-policy-contract-mappings.md`](cmd-pingcli-advanced-services-pingfederate-oauth-authentication-policy-contract-mappings.md) |
| `pingcli advanced-services pingfederate oauth authorization-detail-processors` | PingFederate OAuth authorization detail processors | [`cmd-pingcli-advanced-services-pingfederate-oauth-authorization-detail-processors.md`](cmd-pingcli-advanced-services-pingfederate-oauth-authorization-detail-processors.md) |
| `pingcli advanced-services pingfederate oauth authorization-detail-types` | PingFederate OAuth Authorization Detail Types | [`cmd-pingcli-advanced-services-pingfederate-oauth-authorization-detail-types.md`](cmd-pingcli-advanced-services-pingfederate-oauth-authorization-detail-types.md) |
| `pingcli advanced-services pingfederate oauth ciba-server-policy` | Manage PingFederate OAuth CIBA Server Policy resources | [`cmd-pingcli-advanced-services-pingfederate-oauth-ciba-server-policy.md`](cmd-pingcli-advanced-services-pingfederate-oauth-ciba-server-policy.md) |
| `pingcli advanced-services pingfederate oauth client-registration-policies` | PingFederate OAuth Client Registration Policies | [`cmd-pingcli-advanced-services-pingfederate-oauth-client-registration-policies.md`](cmd-pingcli-advanced-services-pingfederate-oauth-client-registration-policies.md) |
| `pingcli advanced-services pingfederate oauth client-settings` | PingFederate OAuth Client Settings | [`cmd-pingcli-advanced-services-pingfederate-oauth-client-settings.md`](cmd-pingcli-advanced-services-pingfederate-oauth-client-settings.md) |
| `pingcli advanced-services pingfederate oauth clients` | PingFederate OAuth Clients | [`cmd-pingcli-advanced-services-pingfederate-oauth-clients.md`](cmd-pingcli-advanced-services-pingfederate-oauth-clients.md) |
| `pingcli advanced-services pingfederate oauth idp-adapter-mappings` | PingFederate OAuth IdP Adapter Mappings | [`cmd-pingcli-advanced-services-pingfederate-oauth-idp-adapter-mappings.md`](cmd-pingcli-advanced-services-pingfederate-oauth-idp-adapter-mappings.md) |
| `pingcli advanced-services pingfederate oauth issuers` | PingFederate OAuth virtual issuers | [`cmd-pingcli-advanced-services-pingfederate-oauth-issuers.md`](cmd-pingcli-advanced-services-pingfederate-oauth-issuers.md) |
| `pingcli advanced-services pingfederate oauth oidc` | Manage PingFederate OAuth/OpenID Connect resources | [`cmd-pingcli-advanced-services-pingfederate-oauth-oidc.md`](cmd-pingcli-advanced-services-pingfederate-oauth-oidc.md) |
| `pingcli advanced-services pingfederate oauth out-of-band-auth-plugins` | PingFederate OAuth out-of-band authenticator plugin instances | [`cmd-pingcli-advanced-services-pingfederate-oauth-out-of-band-auth-plugins.md`](cmd-pingcli-advanced-services-pingfederate-oauth-out-of-band-auth-plugins.md) |
| `pingcli advanced-services pingfederate oauth processor-policy-mappings` | PingFederate OAuth processor policy mappings | [`cmd-pingcli-advanced-services-pingfederate-oauth-processor-policy-mappings.md`](cmd-pingcli-advanced-services-pingfederate-oauth-processor-policy-mappings.md) |
| `pingcli advanced-services pingfederate oauth resource-owner-credentials-mappings` | PingFederate OAuth Resource Owner Credentials Mappings | [`cmd-pingcli-advanced-services-pingfederate-oauth-resource-owner-credentials-mappings.md`](cmd-pingcli-advanced-services-pingfederate-oauth-resource-owner-credentials-mappings.md) |
| `pingcli advanced-services pingfederate oauth token-exchange` | Manage PingFederate OAuth Token Exchange resources | [`cmd-pingcli-advanced-services-pingfederate-oauth-token-exchange.md`](cmd-pingcli-advanced-services-pingfederate-oauth-token-exchange.md) |

## Parent Command

- [`pingcli advanced-services pingfederate`](cmd-pingcli-advanced-services-pingfederate.md) — Administration tools for PingFederate deployed as a private tenant service in PingOne Advanced Services
