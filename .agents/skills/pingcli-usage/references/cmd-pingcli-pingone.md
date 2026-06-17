# `pingcli pingone`
Administration tools for the PingOne platform.

## Synopsis

Administration tools for the PingOne platform and universal services.
		
When multiple products are configured in the CLI, the platform command can be used to manage one or more products collectively.

The --profile command switch can be used to specify the profile of Ping products to be managed.

```
pingcli pingone
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
| `pingcli pingone active-identity-counts` | Active identity counts | [`cmd-pingcli-pingone-active-identity-counts.md`](cmd-pingcli-pingone-active-identity-counts.md) |
| `pingcli pingone agreements` | Agreements | [`cmd-pingcli-pingone-agreements.md`](cmd-pingcli-pingone-agreements.md) |
| `pingcli pingone api` | Send a custom REST API request to the management API of PingOne. | [`cmd-pingcli-pingone-api.md`](cmd-pingcli-pingone-api.md) |
| `pingcli pingone applications` | Applications | [`cmd-pingcli-pingone-applications.md`](cmd-pingcli-pingone-applications.md) |
| `pingcli pingone auth` | Authenticate Ping CLI to the PingOne management APIs. | [`cmd-pingcli-pingone-auth.md`](cmd-pingcli-pingone-auth.md) |
| `pingcli pingone authorize` | Administration tools for the PingOne Authorize universal service. | [`cmd-pingcli-pingone-authorize.md`](cmd-pingcli-pingone-authorize.md) |
| `pingcli pingone credentials` | Administration tools for the PingOne Credentials universal service. | [`cmd-pingcli-pingone-credentials.md`](cmd-pingcli-pingone-credentials.md) |
| `pingcli pingone custom-admin-roles` | Custom Admin Roles | [`cmd-pingcli-pingone-custom-admin-roles.md`](cmd-pingcli-pingone-custom-admin-roles.md) |
| `pingcli pingone davinci` | Administration tools for the PingOne DaVinci universal service. | [`cmd-pingcli-pingone-davinci.md`](cmd-pingcli-pingone-davinci.md) |
| `pingcli pingone environments` | Environments | [`cmd-pingcli-pingone-environments.md`](cmd-pingcli-pingone-environments.md) |
| `pingcli pingone gateways` | Gateways | [`cmd-pingcli-pingone-gateways.md`](cmd-pingcli-pingone-gateways.md) |
| `pingcli pingone groups` | Groups | [`cmd-pingcli-pingone-groups.md`](cmd-pingcli-pingone-groups.md) |
| `pingcli pingone identity-providers` | Identity Providers | [`cmd-pingcli-pingone-identity-providers.md`](cmd-pingcli-pingone-identity-providers.md) |
| `pingcli pingone init` | Initialize Ping CLI for the PingOne management APIs. | [`cmd-pingcli-pingone-init.md`](cmd-pingcli-pingone-init.md) |
| `pingcli pingone languages` | Languages | [`cmd-pingcli-pingone-languages.md`](cmd-pingcli-pingone-languages.md) |
| `pingcli pingone licenses` | PingOne Licenses | [`cmd-pingcli-pingone-licenses.md`](cmd-pingcli-pingone-licenses.md) |
| `pingcli pingone mfa` | Administration tools for the PingOne MFA universal service. | [`cmd-pingcli-pingone-mfa.md`](cmd-pingcli-pingone-mfa.md) |
| `pingcli pingone notification-policies` | Notification Policies | [`cmd-pingcli-pingone-notification-policies.md`](cmd-pingcli-pingone-notification-policies.md) |
| `pingcli pingone notification-templates` | Notification Templates | [`cmd-pingcli-pingone-notification-templates.md`](cmd-pingcli-pingone-notification-templates.md) |
| `pingcli pingone password-policies` | Password Policies | [`cmd-pingcli-pingone-password-policies.md`](cmd-pingcli-pingone-password-policies.md) |
| `pingcli pingone populations` | Populations | [`cmd-pingcli-pingone-populations.md`](cmd-pingcli-pingone-populations.md) |
| `pingcli pingone protect` | Administration tools for the PingOne Protect universal service. | [`cmd-pingcli-pingone-protect.md`](cmd-pingcli-pingone-protect.md) |
| `pingcli pingone resources` | Resources | [`cmd-pingcli-pingone-resources.md`](cmd-pingcli-pingone-resources.md) |
| `pingcli pingone roles` | PingOne built-in admin roles | [`cmd-pingcli-pingone-roles.md`](cmd-pingcli-pingone-roles.md) |
| `pingcli pingone schemas` | Schemas | [`cmd-pingcli-pingone-schemas.md`](cmd-pingcli-pingone-schemas.md) |
| `pingcli pingone sign-on-policies` | Sign-On Policies | [`cmd-pingcli-pingone-sign-on-policies.md`](cmd-pingcli-pingone-sign-on-policies.md) |
| `pingcli pingone total-identities` | Total identity counts | [`cmd-pingcli-pingone-total-identities.md`](cmd-pingcli-pingone-total-identities.md) |
| `pingcli pingone users` | Users | [`cmd-pingcli-pingone-users.md`](cmd-pingcli-pingone-users.md) |
| `pingcli pingone verify` | Administration tools for the PingOne Verify universal service. | [`cmd-pingcli-pingone-verify.md`](cmd-pingcli-pingone-verify.md) |
| `pingcli pingone webhooks` | Webhooks | [`cmd-pingcli-pingone-webhooks.md`](cmd-pingcli-pingone-webhooks.md) |

## Parent Command

- [`pingcli`](cmd-pingcli.md) — A CLI tool for managing the configuration of Ping Identity products.
