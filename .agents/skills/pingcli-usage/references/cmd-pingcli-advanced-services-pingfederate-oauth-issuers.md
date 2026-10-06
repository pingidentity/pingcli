# `pingcli advanced-services pingfederate oauth issuers`
PingFederate OAuth virtual issuers

## Synopsis

PingFederate OAuth virtual issuers allow a single PingFederate server to issue OAuth/OIDC tokens under multiple distinct issuer identities.

```
pingcli advanced-services pingfederate oauth issuers [flags]
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
| `pingcli advanced-services pingfederate oauth issuers apply` | Create or update an OAuth virtual issuer | [`cmd-pingcli-advanced-services-pingfederate-oauth-issuers-apply.md`](cmd-pingcli-advanced-services-pingfederate-oauth-issuers-apply.md) |
| `pingcli advanced-services pingfederate oauth issuers create` | Create a new OAuth virtual issuer | [`cmd-pingcli-advanced-services-pingfederate-oauth-issuers-create.md`](cmd-pingcli-advanced-services-pingfederate-oauth-issuers-create.md) |
| `pingcli advanced-services pingfederate oauth issuers delete` | Delete an OAuth virtual issuer | [`cmd-pingcli-advanced-services-pingfederate-oauth-issuers-delete.md`](cmd-pingcli-advanced-services-pingfederate-oauth-issuers-delete.md) |
| `pingcli advanced-services pingfederate oauth issuers get` | Read a specific OAuth virtual issuer | [`cmd-pingcli-advanced-services-pingfederate-oauth-issuers-get.md`](cmd-pingcli-advanced-services-pingfederate-oauth-issuers-get.md) |
| `pingcli advanced-services pingfederate oauth issuers list` | List all OAuth virtual issuers | [`cmd-pingcli-advanced-services-pingfederate-oauth-issuers-list.md`](cmd-pingcli-advanced-services-pingfederate-oauth-issuers-list.md) |
| `pingcli advanced-services pingfederate oauth issuers replace` | Update an OAuth virtual issuer | [`cmd-pingcli-advanced-services-pingfederate-oauth-issuers-replace.md`](cmd-pingcli-advanced-services-pingfederate-oauth-issuers-replace.md) |
| `pingcli advanced-services pingfederate oauth issuers template` | Generate an OAuth virtual issuer JSON template | [`cmd-pingcli-advanced-services-pingfederate-oauth-issuers-template.md`](cmd-pingcli-advanced-services-pingfederate-oauth-issuers-template.md) |

## Parent Command

- [`pingcli advanced-services pingfederate oauth`](cmd-pingcli-advanced-services-pingfederate-oauth.md) — Manage PingFederate OAuth resources
