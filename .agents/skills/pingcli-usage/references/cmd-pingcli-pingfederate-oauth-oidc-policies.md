# `pingcli pingfederate oauth oidc policies`
PingFederate OAuth/OpenID Connect Policies

## Synopsis

PingFederate OAuth/OpenID Connect Policies

```
pingcli pingfederate oauth oidc policies [flags]
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
| `pingcli pingfederate oauth oidc policies apply` | Create or update an OAuth/OpenID Connect policy | [`cmd-pingcli-pingfederate-oauth-oidc-policies-apply.md`](cmd-pingcli-pingfederate-oauth-oidc-policies-apply.md) |
| `pingcli pingfederate oauth oidc policies create` | Create a new OAuth/OpenID Connect policy | [`cmd-pingcli-pingfederate-oauth-oidc-policies-create.md`](cmd-pingcli-pingfederate-oauth-oidc-policies-create.md) |
| `pingcli pingfederate oauth oidc policies delete` | Delete an OAuth/OpenID Connect policy | [`cmd-pingcli-pingfederate-oauth-oidc-policies-delete.md`](cmd-pingcli-pingfederate-oauth-oidc-policies-delete.md) |
| `pingcli pingfederate oauth oidc policies get` | Read a specific OAuth/OpenID Connect policy | [`cmd-pingcli-pingfederate-oauth-oidc-policies-get.md`](cmd-pingcli-pingfederate-oauth-oidc-policies-get.md) |
| `pingcli pingfederate oauth oidc policies list` | List all OAuth/OpenID Connect policies | [`cmd-pingcli-pingfederate-oauth-oidc-policies-list.md`](cmd-pingcli-pingfederate-oauth-oidc-policies-list.md) |
| `pingcli pingfederate oauth oidc policies replace` | Update an OAuth/OpenID Connect policy | [`cmd-pingcli-pingfederate-oauth-oidc-policies-replace.md`](cmd-pingcli-pingfederate-oauth-oidc-policies-replace.md) |
| `pingcli pingfederate oauth oidc policies template` | Generate an OAuth/OpenID Connect policy JSON template | [`cmd-pingcli-pingfederate-oauth-oidc-policies-template.md`](cmd-pingcli-pingfederate-oauth-oidc-policies-template.md) |

## Parent Command

- [`pingcli pingfederate oauth oidc`](cmd-pingcli-pingfederate-oauth-oidc.md) — Manage PingFederate OAuth/OpenID Connect resources
