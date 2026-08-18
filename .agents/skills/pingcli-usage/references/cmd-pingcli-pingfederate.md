# `pingcli pingfederate`
Administration tools for PingFederate deployed as software

## Synopsis

Administration tools for PingFederate deployed as software.
		
When multiple products are configured in the CLI, the platform command can be used to manage one or more products collectively.

The --profile command switch can be used to specify the profile of Ping products to be managed.

```
pingcli pingfederate
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
| `pingcli pingfederate administrative-accounts` | PingFederate Administrative Accounts | [`cmd-pingcli-pingfederate-administrative-accounts.md`](cmd-pingcli-pingfederate-administrative-accounts.md) |
| `pingcli pingfederate administrative-api` | Manage PingFederate Administrative API resources | [`cmd-pingcli-pingfederate-administrative-api.md`](cmd-pingcli-pingfederate-administrative-api.md) |
| `pingcli pingfederate api` | Send a custom REST API request to the management API of PingFederate. | [`cmd-pingcli-pingfederate-api.md`](cmd-pingcli-pingfederate-api.md) |
| `pingcli pingfederate auth` | Authenticate Ping CLI to the PingFederate management APIs. | [`cmd-pingcli-pingfederate-auth.md`](cmd-pingcli-pingfederate-auth.md) |
| `pingcli pingfederate authentication-api` | Manage PingFederate Authentication API resources | [`cmd-pingcli-pingfederate-authentication-api.md`](cmd-pingcli-pingfederate-authentication-api.md) |
| `pingcli pingfederate authentication-policies` | Manage PingFederate Authentication Policies resources | [`cmd-pingcli-pingfederate-authentication-policies.md`](cmd-pingcli-pingfederate-authentication-policies.md) |
| `pingcli pingfederate authentication-policy-contracts` | PingFederate Authentication Policy Contracts | [`cmd-pingcli-pingfederate-authentication-policy-contracts.md`](cmd-pingcli-pingfederate-authentication-policy-contracts.md) |
| `pingcli pingfederate authentication-selectors` | PingFederate Authentication Selectors | [`cmd-pingcli-pingfederate-authentication-selectors.md`](cmd-pingcli-pingfederate-authentication-selectors.md) |
| `pingcli pingfederate bulk` | Manage PingFederate Bulk resources | [`cmd-pingcli-pingfederate-bulk.md`](cmd-pingcli-pingfederate-bulk.md) |
| `pingcli pingfederate captcha-providers` | PingFederate CAPTCHA Providers | [`cmd-pingcli-pingfederate-captcha-providers.md`](cmd-pingcli-pingfederate-captcha-providers.md) |
| `pingcli pingfederate certificates` | Manage PingFederate Certificates resources | [`cmd-pingcli-pingfederate-certificates.md`](cmd-pingcli-pingfederate-certificates.md) |
| `pingcli pingfederate cluster` | Manage PingFederate Cluster resources | [`cmd-pingcli-pingfederate-cluster.md`](cmd-pingcli-pingfederate-cluster.md) |
| `pingcli pingfederate collect-support-data` | Manage PingFederate Collect Support Data (CSD) resources | [`cmd-pingcli-pingfederate-collect-support-data.md`](cmd-pingcli-pingfederate-collect-support-data.md) |
| `pingcli pingfederate config-archive` | Manage the PingFederate configuration archive | [`cmd-pingcli-pingfederate-config-archive.md`](cmd-pingcli-pingfederate-config-archive.md) |
| `pingcli pingfederate config-store-settings` | PingFederate Configuration Store Settings | [`cmd-pingcli-pingfederate-config-store-settings.md`](cmd-pingcli-pingfederate-config-store-settings.md) |
| `pingcli pingfederate configuration-encryption-keys` | PingFederate Configuration Encryption Keys | [`cmd-pingcli-pingfederate-configuration-encryption-keys.md`](cmd-pingcli-pingfederate-configuration-encryption-keys.md) |
| `pingcli pingfederate connection-metadata` | Manage PingFederate connection metadata | [`cmd-pingcli-pingfederate-connection-metadata.md`](cmd-pingcli-pingfederate-connection-metadata.md) |
| `pingcli pingfederate data-stores` | PingFederate Data Stores | [`cmd-pingcli-pingfederate-data-stores.md`](cmd-pingcli-pingfederate-data-stores.md) |
| `pingcli pingfederate extended-properties` | PingFederate Extended Properties | [`cmd-pingcli-pingfederate-extended-properties.md`](cmd-pingcli-pingfederate-extended-properties.md) |
| `pingcli pingfederate identity-store-provisioners` | PingFederate Identity Store Provisioners | [`cmd-pingcli-pingfederate-identity-store-provisioners.md`](cmd-pingcli-pingfederate-identity-store-provisioners.md) |
| `pingcli pingfederate idp` | Manage PingFederate IdP resources | [`cmd-pingcli-pingfederate-idp.md`](cmd-pingcli-pingfederate-idp.md) |
| `pingcli pingfederate init` | Initialize Ping CLI for the PingFederate management APIs. | [`cmd-pingcli-pingfederate-init.md`](cmd-pingcli-pingfederate-init.md) |
| `pingcli pingfederate key-pairs` | Manage PingFederate Key Pairs resources | [`cmd-pingcli-pingfederate-key-pairs.md`](cmd-pingcli-pingfederate-key-pairs.md) |
| `pingcli pingfederate local-identity` | Manage PingFederate Local Identity resources | [`cmd-pingcli-pingfederate-local-identity.md`](cmd-pingcli-pingfederate-local-identity.md) |
| `pingcli pingfederate oauth` | Manage PingFederate OAuth resources | [`cmd-pingcli-pingfederate-oauth.md`](cmd-pingcli-pingfederate-oauth.md) |
| `pingcli pingfederate password-credential-validators` | PingFederate Password Credential Validators | [`cmd-pingcli-pingfederate-password-credential-validators.md`](cmd-pingcli-pingfederate-password-credential-validators.md) |
| `pingcli pingfederate ping-one-connections` | PingFederate PingOne connections | [`cmd-pingcli-pingfederate-ping-one-connections.md`](cmd-pingcli-pingfederate-ping-one-connections.md) |
| `pingcli pingfederate protocol-metadata` | Manage PingFederate Protocol Metadata resources | [`cmd-pingcli-pingfederate-protocol-metadata.md`](cmd-pingcli-pingfederate-protocol-metadata.md) |
| `pingcli pingfederate secret-managers` | PingFederate Secret Managers | [`cmd-pingcli-pingfederate-secret-managers.md`](cmd-pingcli-pingfederate-secret-managers.md) |
| `pingcli pingfederate server-settings` | PingFederate Server Settings | [`cmd-pingcli-pingfederate-server-settings.md`](cmd-pingcli-pingfederate-server-settings.md) |
| `pingcli pingfederate session` | Manage PingFederate Session resources | [`cmd-pingcli-pingfederate-session.md`](cmd-pingcli-pingfederate-session.md) |
| `pingcli pingfederate sp` | Manage PingFederate SP resources | [`cmd-pingcli-pingfederate-sp.md`](cmd-pingcli-pingfederate-sp.md) |
| `pingcli pingfederate token-processor-to-token-generator-mappings` | PingFederate Token Processor to Token Generator Mappings | [`cmd-pingcli-pingfederate-token-processor-to-token-generator-mappings.md`](cmd-pingcli-pingfederate-token-processor-to-token-generator-mappings.md) |
| `pingcli pingfederate version` | PingFederate Version | [`cmd-pingcli-pingfederate-version.md`](cmd-pingcli-pingfederate-version.md) |
| `pingcli pingfederate virtual-host-names` | PingFederate Virtual Host Names | [`cmd-pingcli-pingfederate-virtual-host-names.md`](cmd-pingcli-pingfederate-virtual-host-names.md) |

## Parent Command

- [`pingcli`](cmd-pingcli.md) — A CLI tool for managing the configuration of Ping Identity products.
