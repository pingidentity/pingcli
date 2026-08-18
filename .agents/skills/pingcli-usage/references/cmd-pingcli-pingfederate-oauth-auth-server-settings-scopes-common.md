# `pingcli pingfederate oauth auth-server-settings scopes common`
PingFederate OAuth common scopes

## Synopsis

PingFederate OAuth common scopes allow a single PingFederate server to scope OAuth/OIDC tokens under multiple distinct common-scope identities.

```
pingcli pingfederate oauth auth-server-settings scopes common [flags]
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
| `pingcli pingfederate oauth auth-server-settings scopes common apply` | Create or update an OAuth common scope | [`cmd-pingcli-pingfederate-oauth-auth-server-settings-scopes-common-apply.md`](cmd-pingcli-pingfederate-oauth-auth-server-settings-scopes-common-apply.md) |
| `pingcli pingfederate oauth auth-server-settings scopes common create` | Create a new OAuth common scope | [`cmd-pingcli-pingfederate-oauth-auth-server-settings-scopes-common-create.md`](cmd-pingcli-pingfederate-oauth-auth-server-settings-scopes-common-create.md) |
| `pingcli pingfederate oauth auth-server-settings scopes common delete` | Delete an OAuth common scope | [`cmd-pingcli-pingfederate-oauth-auth-server-settings-scopes-common-delete.md`](cmd-pingcli-pingfederate-oauth-auth-server-settings-scopes-common-delete.md) |
| `pingcli pingfederate oauth auth-server-settings scopes common get` | Read a specific OAuth common scope | [`cmd-pingcli-pingfederate-oauth-auth-server-settings-scopes-common-get.md`](cmd-pingcli-pingfederate-oauth-auth-server-settings-scopes-common-get.md) |
| `pingcli pingfederate oauth auth-server-settings scopes common list` | List all OAuth common scopes | [`cmd-pingcli-pingfederate-oauth-auth-server-settings-scopes-common-list.md`](cmd-pingcli-pingfederate-oauth-auth-server-settings-scopes-common-list.md) |
| `pingcli pingfederate oauth auth-server-settings scopes common replace` | Update an OAuth common scope | [`cmd-pingcli-pingfederate-oauth-auth-server-settings-scopes-common-replace.md`](cmd-pingcli-pingfederate-oauth-auth-server-settings-scopes-common-replace.md) |
| `pingcli pingfederate oauth auth-server-settings scopes common template` | Generate an OAuth common scope JSON template | [`cmd-pingcli-pingfederate-oauth-auth-server-settings-scopes-common-template.md`](cmd-pingcli-pingfederate-oauth-auth-server-settings-scopes-common-template.md) |

## Parent Command

- [`pingcli pingfederate oauth auth-server-settings scopes`](cmd-pingcli-pingfederate-oauth-auth-server-settings-scopes.md) — Manage OAuth Authorization Server Settings scope collections
