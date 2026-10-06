# `pingcli advanced-services pingfederate oauth auth-server-settings scopes common-scope-groups`
PingFederate OAuth common scope groups

## Synopsis

PingFederate OAuth common scope groups define named groups of common scopes under the OAuth Authorization Server Settings.

```
pingcli advanced-services pingfederate oauth auth-server-settings scopes common-scope-groups [flags]
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
| `pingcli advanced-services pingfederate oauth auth-server-settings scopes common-scope-groups apply` | Create or update a common scope group | [`cmd-pingcli-advanced-services-pingfederate-oauth-auth-server-settings-scopes-common-scope-groups-apply.md`](cmd-pingcli-advanced-services-pingfederate-oauth-auth-server-settings-scopes-common-scope-groups-apply.md) |
| `pingcli advanced-services pingfederate oauth auth-server-settings scopes common-scope-groups create` | Create a new common scope group | [`cmd-pingcli-advanced-services-pingfederate-oauth-auth-server-settings-scopes-common-scope-groups-create.md`](cmd-pingcli-advanced-services-pingfederate-oauth-auth-server-settings-scopes-common-scope-groups-create.md) |
| `pingcli advanced-services pingfederate oauth auth-server-settings scopes common-scope-groups delete` | Delete a common scope group | [`cmd-pingcli-advanced-services-pingfederate-oauth-auth-server-settings-scopes-common-scope-groups-delete.md`](cmd-pingcli-advanced-services-pingfederate-oauth-auth-server-settings-scopes-common-scope-groups-delete.md) |
| `pingcli advanced-services pingfederate oauth auth-server-settings scopes common-scope-groups get` | Read a specific common scope group | [`cmd-pingcli-advanced-services-pingfederate-oauth-auth-server-settings-scopes-common-scope-groups-get.md`](cmd-pingcli-advanced-services-pingfederate-oauth-auth-server-settings-scopes-common-scope-groups-get.md) |
| `pingcli advanced-services pingfederate oauth auth-server-settings scopes common-scope-groups list` | List all common scope groups | [`cmd-pingcli-advanced-services-pingfederate-oauth-auth-server-settings-scopes-common-scope-groups-list.md`](cmd-pingcli-advanced-services-pingfederate-oauth-auth-server-settings-scopes-common-scope-groups-list.md) |
| `pingcli advanced-services pingfederate oauth auth-server-settings scopes common-scope-groups replace` | Update a common scope group | [`cmd-pingcli-advanced-services-pingfederate-oauth-auth-server-settings-scopes-common-scope-groups-replace.md`](cmd-pingcli-advanced-services-pingfederate-oauth-auth-server-settings-scopes-common-scope-groups-replace.md) |
| `pingcli advanced-services pingfederate oauth auth-server-settings scopes common-scope-groups template` | Generate a common scope group JSON template | [`cmd-pingcli-advanced-services-pingfederate-oauth-auth-server-settings-scopes-common-scope-groups-template.md`](cmd-pingcli-advanced-services-pingfederate-oauth-auth-server-settings-scopes-common-scope-groups-template.md) |

## Parent Command

- [`pingcli advanced-services pingfederate oauth auth-server-settings scopes`](cmd-pingcli-advanced-services-pingfederate-oauth-auth-server-settings-scopes.md) — Manage OAuth Authorization Server Settings scope collections
