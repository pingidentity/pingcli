# `pingcli advanced-services pingfederate`
Administration tools for PingFederate deployed as a private tenant service in PingOne Advanced Services

## Synopsis

Administration tools for PingFederate deployed as a private tenant service in PingOne Advanced Services.
		
When multiple products are configured in the CLI, the platform command can be used to manage one or more products collectively.

The --profile command switch can be used to specify the profile of Ping products to be managed.

```
pingcli advanced-services pingfederate
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
| `pingcli advanced-services pingfederate administrative-accounts` | PingFederate Administrative Accounts | [`cmd-pingcli-advanced-services-pingfederate-administrative-accounts.md`](cmd-pingcli-advanced-services-pingfederate-administrative-accounts.md) |
| `pingcli advanced-services pingfederate administrative-api` | Manage PingFederate Administrative API resources | [`cmd-pingcli-advanced-services-pingfederate-administrative-api.md`](cmd-pingcli-advanced-services-pingfederate-administrative-api.md) |
| `pingcli advanced-services pingfederate api` | Send a custom REST API request to the management API of PingFederate (Advanced Services). | [`cmd-pingcli-advanced-services-pingfederate-api.md`](cmd-pingcli-advanced-services-pingfederate-api.md) |
| `pingcli advanced-services pingfederate auth` | Authenticate Ping CLI to the PingFederate (Advanced Services) management APIs. | [`cmd-pingcli-advanced-services-pingfederate-auth.md`](cmd-pingcli-advanced-services-pingfederate-auth.md) |
| `pingcli advanced-services pingfederate authentication-api` | Manage PingFederate Authentication API resources | [`cmd-pingcli-advanced-services-pingfederate-authentication-api.md`](cmd-pingcli-advanced-services-pingfederate-authentication-api.md) |
| `pingcli advanced-services pingfederate authentication-policies` | Manage PingFederate Authentication Policies resources | [`cmd-pingcli-advanced-services-pingfederate-authentication-policies.md`](cmd-pingcli-advanced-services-pingfederate-authentication-policies.md) |
| `pingcli advanced-services pingfederate authentication-policy-contracts` | PingFederate Authentication Policy Contracts | [`cmd-pingcli-advanced-services-pingfederate-authentication-policy-contracts.md`](cmd-pingcli-advanced-services-pingfederate-authentication-policy-contracts.md) |
| `pingcli advanced-services pingfederate authentication-selectors` | PingFederate Authentication Selectors | [`cmd-pingcli-advanced-services-pingfederate-authentication-selectors.md`](cmd-pingcli-advanced-services-pingfederate-authentication-selectors.md) |
| `pingcli advanced-services pingfederate bulk` | Manage PingFederate Bulk resources | [`cmd-pingcli-advanced-services-pingfederate-bulk.md`](cmd-pingcli-advanced-services-pingfederate-bulk.md) |
| `pingcli advanced-services pingfederate captcha-providers` | PingFederate CAPTCHA Providers | [`cmd-pingcli-advanced-services-pingfederate-captcha-providers.md`](cmd-pingcli-advanced-services-pingfederate-captcha-providers.md) |
| `pingcli advanced-services pingfederate certificates` | Manage PingFederate Certificates resources | [`cmd-pingcli-advanced-services-pingfederate-certificates.md`](cmd-pingcli-advanced-services-pingfederate-certificates.md) |
| `pingcli advanced-services pingfederate cluster` | Manage PingFederate Cluster resources | [`cmd-pingcli-advanced-services-pingfederate-cluster.md`](cmd-pingcli-advanced-services-pingfederate-cluster.md) |
| `pingcli advanced-services pingfederate collect-support-data` | Manage PingFederate Collect Support Data (CSD) resources | [`cmd-pingcli-advanced-services-pingfederate-collect-support-data.md`](cmd-pingcli-advanced-services-pingfederate-collect-support-data.md) |
| `pingcli advanced-services pingfederate config-archive` | Manage the PingFederate configuration archive | [`cmd-pingcli-advanced-services-pingfederate-config-archive.md`](cmd-pingcli-advanced-services-pingfederate-config-archive.md) |
| `pingcli advanced-services pingfederate config-store-settings` | PingFederate Configuration Store Settings | [`cmd-pingcli-advanced-services-pingfederate-config-store-settings.md`](cmd-pingcli-advanced-services-pingfederate-config-store-settings.md) |
| `pingcli advanced-services pingfederate configuration-encryption-keys` | PingFederate Configuration Encryption Keys | [`cmd-pingcli-advanced-services-pingfederate-configuration-encryption-keys.md`](cmd-pingcli-advanced-services-pingfederate-configuration-encryption-keys.md) |
| `pingcli advanced-services pingfederate connection-metadata` | Manage PingFederate connection metadata | [`cmd-pingcli-advanced-services-pingfederate-connection-metadata.md`](cmd-pingcli-advanced-services-pingfederate-connection-metadata.md) |
| `pingcli advanced-services pingfederate data-stores` | PingFederate Data Stores | [`cmd-pingcli-advanced-services-pingfederate-data-stores.md`](cmd-pingcli-advanced-services-pingfederate-data-stores.md) |
| `pingcli advanced-services pingfederate extended-properties` | PingFederate Extended Properties | [`cmd-pingcli-advanced-services-pingfederate-extended-properties.md`](cmd-pingcli-advanced-services-pingfederate-extended-properties.md) |
| `pingcli advanced-services pingfederate identity-store-provisioners` | PingFederate Identity Store Provisioners | [`cmd-pingcli-advanced-services-pingfederate-identity-store-provisioners.md`](cmd-pingcli-advanced-services-pingfederate-identity-store-provisioners.md) |
| `pingcli advanced-services pingfederate idp` | Manage PingFederate IdP resources | [`cmd-pingcli-advanced-services-pingfederate-idp.md`](cmd-pingcli-advanced-services-pingfederate-idp.md) |
| `pingcli advanced-services pingfederate idp-to-sp-adapter-mappings` | PingFederate IdP-to-SP Adapter Mappings | [`cmd-pingcli-advanced-services-pingfederate-idp-to-sp-adapter-mappings.md`](cmd-pingcli-advanced-services-pingfederate-idp-to-sp-adapter-mappings.md) |
| `pingcli advanced-services pingfederate incoming-proxy-settings` | PingFederate incoming proxy settings | [`cmd-pingcli-advanced-services-pingfederate-incoming-proxy-settings.md`](cmd-pingcli-advanced-services-pingfederate-incoming-proxy-settings.md) |
| `pingcli advanced-services pingfederate init` | Initialize Ping CLI for the PingFederate (Advanced Services) management APIs. | [`cmd-pingcli-advanced-services-pingfederate-init.md`](cmd-pingcli-advanced-services-pingfederate-init.md) |
| `pingcli advanced-services pingfederate kerberos` | Manage PingFederate Kerberos resources | [`cmd-pingcli-advanced-services-pingfederate-kerberos.md`](cmd-pingcli-advanced-services-pingfederate-kerberos.md) |
| `pingcli advanced-services pingfederate key-pairs` | Manage PingFederate Key Pairs resources | [`cmd-pingcli-advanced-services-pingfederate-key-pairs.md`](cmd-pingcli-advanced-services-pingfederate-key-pairs.md) |
| `pingcli advanced-services pingfederate license` | PingFederate license | [`cmd-pingcli-advanced-services-pingfederate-license.md`](cmd-pingcli-advanced-services-pingfederate-license.md) |
| `pingcli advanced-services pingfederate local-identity` | Manage PingFederate Local Identity resources | [`cmd-pingcli-advanced-services-pingfederate-local-identity.md`](cmd-pingcli-advanced-services-pingfederate-local-identity.md) |
| `pingcli advanced-services pingfederate metadata-urls` | PingFederate metadata URLs | [`cmd-pingcli-advanced-services-pingfederate-metadata-urls.md`](cmd-pingcli-advanced-services-pingfederate-metadata-urls.md) |
| `pingcli advanced-services pingfederate notification-publishers` | PingFederate notification publishers | [`cmd-pingcli-advanced-services-pingfederate-notification-publishers.md`](cmd-pingcli-advanced-services-pingfederate-notification-publishers.md) |
| `pingcli advanced-services pingfederate oauth` | Manage PingFederate OAuth resources | [`cmd-pingcli-advanced-services-pingfederate-oauth.md`](cmd-pingcli-advanced-services-pingfederate-oauth.md) |
| `pingcli advanced-services pingfederate password-credential-validators` | PingFederate Password Credential Validators | [`cmd-pingcli-advanced-services-pingfederate-password-credential-validators.md`](cmd-pingcli-advanced-services-pingfederate-password-credential-validators.md) |
| `pingcli advanced-services pingfederate ping-one-connections` | PingFederate PingOne connections | [`cmd-pingcli-advanced-services-pingfederate-ping-one-connections.md`](cmd-pingcli-advanced-services-pingfederate-ping-one-connections.md) |
| `pingcli advanced-services pingfederate pingone-for-enterprise` | Manage the PingFederate connection to PingOne for Enterprise | [`cmd-pingcli-advanced-services-pingfederate-pingone-for-enterprise.md`](cmd-pingcli-advanced-services-pingfederate-pingone-for-enterprise.md) |
| `pingcli advanced-services pingfederate protocol-metadata` | Manage PingFederate Protocol Metadata resources | [`cmd-pingcli-advanced-services-pingfederate-protocol-metadata.md`](cmd-pingcli-advanced-services-pingfederate-protocol-metadata.md) |
| `pingcli advanced-services pingfederate redirect-validation` | PingFederate Redirect Validation Settings | [`cmd-pingcli-advanced-services-pingfederate-redirect-validation.md`](cmd-pingcli-advanced-services-pingfederate-redirect-validation.md) |
| `pingcli advanced-services pingfederate secret-managers` | PingFederate Secret Managers | [`cmd-pingcli-advanced-services-pingfederate-secret-managers.md`](cmd-pingcli-advanced-services-pingfederate-secret-managers.md) |
| `pingcli advanced-services pingfederate server-settings` | PingFederate Server Settings | [`cmd-pingcli-advanced-services-pingfederate-server-settings.md`](cmd-pingcli-advanced-services-pingfederate-server-settings.md) |
| `pingcli advanced-services pingfederate service-authentication` | PingFederate Service Authentication | [`cmd-pingcli-advanced-services-pingfederate-service-authentication.md`](cmd-pingcli-advanced-services-pingfederate-service-authentication.md) |
| `pingcli advanced-services pingfederate session` | Manage PingFederate Session resources | [`cmd-pingcli-advanced-services-pingfederate-session.md`](cmd-pingcli-advanced-services-pingfederate-session.md) |
| `pingcli advanced-services pingfederate sp` | Manage PingFederate SP resources | [`cmd-pingcli-advanced-services-pingfederate-sp.md`](cmd-pingcli-advanced-services-pingfederate-sp.md) |
| `pingcli advanced-services pingfederate token-processor-to-token-generator-mappings` | PingFederate Token Processor to Token Generator Mappings | [`cmd-pingcli-advanced-services-pingfederate-token-processor-to-token-generator-mappings.md`](cmd-pingcli-advanced-services-pingfederate-token-processor-to-token-generator-mappings.md) |
| `pingcli advanced-services pingfederate version` | PingFederate Version | [`cmd-pingcli-advanced-services-pingfederate-version.md`](cmd-pingcli-advanced-services-pingfederate-version.md) |
| `pingcli advanced-services pingfederate virtual-host-names` | PingFederate Virtual Host Names | [`cmd-pingcli-advanced-services-pingfederate-virtual-host-names.md`](cmd-pingcli-advanced-services-pingfederate-virtual-host-names.md) |

## Parent Command

- [`pingcli advanced-services`](cmd-pingcli-advanced-services.md) — Administration tools for your PingOne Advanced Services tenant.
