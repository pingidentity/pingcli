# `pingcli advanced-services pingfederate oauth auth-server-settings scopes exclusive-scope-groups`
PingFederate OAuth exclusive scope groups

## Synopsis

PingFederate OAuth exclusive scope groups define a named group of static OAuth scopes that is unavailable to clients by default; administrators must explicitly authorize individual clients to use the group.

```
pingcli advanced-services pingfederate oauth auth-server-settings scopes exclusive-scope-groups [flags]
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
| `pingcli advanced-services pingfederate oauth auth-server-settings scopes exclusive-scope-groups apply` | Create or update an OAuth exclusive scope group | [`cmd-pingcli-advanced-services-pingfederate-oauth-auth-server-settings-scopes-exclusive-scope-groups-apply.md`](cmd-pingcli-advanced-services-pingfederate-oauth-auth-server-settings-scopes-exclusive-scope-groups-apply.md) |
| `pingcli advanced-services pingfederate oauth auth-server-settings scopes exclusive-scope-groups create` | Create a new OAuth exclusive scope group | [`cmd-pingcli-advanced-services-pingfederate-oauth-auth-server-settings-scopes-exclusive-scope-groups-create.md`](cmd-pingcli-advanced-services-pingfederate-oauth-auth-server-settings-scopes-exclusive-scope-groups-create.md) |
| `pingcli advanced-services pingfederate oauth auth-server-settings scopes exclusive-scope-groups delete` | Delete an OAuth exclusive scope group | [`cmd-pingcli-advanced-services-pingfederate-oauth-auth-server-settings-scopes-exclusive-scope-groups-delete.md`](cmd-pingcli-advanced-services-pingfederate-oauth-auth-server-settings-scopes-exclusive-scope-groups-delete.md) |
| `pingcli advanced-services pingfederate oauth auth-server-settings scopes exclusive-scope-groups get` | Read a specific OAuth exclusive scope group | [`cmd-pingcli-advanced-services-pingfederate-oauth-auth-server-settings-scopes-exclusive-scope-groups-get.md`](cmd-pingcli-advanced-services-pingfederate-oauth-auth-server-settings-scopes-exclusive-scope-groups-get.md) |
| `pingcli advanced-services pingfederate oauth auth-server-settings scopes exclusive-scope-groups list` | List all OAuth exclusive scope groups | [`cmd-pingcli-advanced-services-pingfederate-oauth-auth-server-settings-scopes-exclusive-scope-groups-list.md`](cmd-pingcli-advanced-services-pingfederate-oauth-auth-server-settings-scopes-exclusive-scope-groups-list.md) |
| `pingcli advanced-services pingfederate oauth auth-server-settings scopes exclusive-scope-groups replace` | Update an OAuth exclusive scope group | [`cmd-pingcli-advanced-services-pingfederate-oauth-auth-server-settings-scopes-exclusive-scope-groups-replace.md`](cmd-pingcli-advanced-services-pingfederate-oauth-auth-server-settings-scopes-exclusive-scope-groups-replace.md) |
| `pingcli advanced-services pingfederate oauth auth-server-settings scopes exclusive-scope-groups template` | Generate an OAuth exclusive scope group JSON template | [`cmd-pingcli-advanced-services-pingfederate-oauth-auth-server-settings-scopes-exclusive-scope-groups-template.md`](cmd-pingcli-advanced-services-pingfederate-oauth-auth-server-settings-scopes-exclusive-scope-groups-template.md) |

## Parent Command

- [`pingcli advanced-services pingfederate oauth auth-server-settings scopes`](cmd-pingcli-advanced-services-pingfederate-oauth-auth-server-settings-scopes.md) — Manage OAuth Authorization Server Settings scope collections
