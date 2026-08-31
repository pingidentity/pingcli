# `pingcli pingfederate oauth auth-server-settings scopes exclusive`
PingFederate OAuth exclusive scopes

## Synopsis

PingFederate OAuth exclusive scopes allow a single PingFederate server to scope OAuth/OIDC tokens under multiple distinct exclusive-scope identities.

```
pingcli pingfederate oauth auth-server-settings scopes exclusive [flags]
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
| `pingcli pingfederate oauth auth-server-settings scopes exclusive apply` | Create or update an OAuth exclusive scope | [`cmd-pingcli-pingfederate-oauth-auth-server-settings-scopes-exclusive-apply.md`](cmd-pingcli-pingfederate-oauth-auth-server-settings-scopes-exclusive-apply.md) |
| `pingcli pingfederate oauth auth-server-settings scopes exclusive create` | Create a new OAuth exclusive scope | [`cmd-pingcli-pingfederate-oauth-auth-server-settings-scopes-exclusive-create.md`](cmd-pingcli-pingfederate-oauth-auth-server-settings-scopes-exclusive-create.md) |
| `pingcli pingfederate oauth auth-server-settings scopes exclusive delete` | Delete an OAuth exclusive scope | [`cmd-pingcli-pingfederate-oauth-auth-server-settings-scopes-exclusive-delete.md`](cmd-pingcli-pingfederate-oauth-auth-server-settings-scopes-exclusive-delete.md) |
| `pingcli pingfederate oauth auth-server-settings scopes exclusive get` | Read a specific OAuth exclusive scope | [`cmd-pingcli-pingfederate-oauth-auth-server-settings-scopes-exclusive-get.md`](cmd-pingcli-pingfederate-oauth-auth-server-settings-scopes-exclusive-get.md) |
| `pingcli pingfederate oauth auth-server-settings scopes exclusive list` | List all OAuth exclusive scopes | [`cmd-pingcli-pingfederate-oauth-auth-server-settings-scopes-exclusive-list.md`](cmd-pingcli-pingfederate-oauth-auth-server-settings-scopes-exclusive-list.md) |
| `pingcli pingfederate oauth auth-server-settings scopes exclusive replace` | Update an OAuth exclusive scope | [`cmd-pingcli-pingfederate-oauth-auth-server-settings-scopes-exclusive-replace.md`](cmd-pingcli-pingfederate-oauth-auth-server-settings-scopes-exclusive-replace.md) |
| `pingcli pingfederate oauth auth-server-settings scopes exclusive template` | Generate an OAuth exclusive scope JSON template | [`cmd-pingcli-pingfederate-oauth-auth-server-settings-scopes-exclusive-template.md`](cmd-pingcli-pingfederate-oauth-auth-server-settings-scopes-exclusive-template.md) |

## Parent Command

- [`pingcli pingfederate oauth auth-server-settings scopes`](cmd-pingcli-pingfederate-oauth-auth-server-settings-scopes.md) — Manage OAuth Authorization Server Settings scope collections
