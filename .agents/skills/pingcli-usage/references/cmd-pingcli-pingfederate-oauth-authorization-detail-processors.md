# `pingcli pingfederate oauth authorization-detail-processors`
PingFederate OAuth authorization detail processors

## Synopsis

PingFederate OAuth authorization detail processors validate and process the authorization details of an OAuth request.

```
pingcli pingfederate oauth authorization-detail-processors [flags]
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
| `pingcli pingfederate oauth authorization-detail-processors apply` | Create or update an authorization detail processor | [`cmd-pingcli-pingfederate-oauth-authorization-detail-processors-apply.md`](cmd-pingcli-pingfederate-oauth-authorization-detail-processors-apply.md) |
| `pingcli pingfederate oauth authorization-detail-processors create` | Create a new authorization detail processor | [`cmd-pingcli-pingfederate-oauth-authorization-detail-processors-create.md`](cmd-pingcli-pingfederate-oauth-authorization-detail-processors-create.md) |
| `pingcli pingfederate oauth authorization-detail-processors delete` | Delete an authorization detail processor | [`cmd-pingcli-pingfederate-oauth-authorization-detail-processors-delete.md`](cmd-pingcli-pingfederate-oauth-authorization-detail-processors-delete.md) |
| `pingcli pingfederate oauth authorization-detail-processors get` | Read a specific authorization detail processor | [`cmd-pingcli-pingfederate-oauth-authorization-detail-processors-get.md`](cmd-pingcli-pingfederate-oauth-authorization-detail-processors-get.md) |
| `pingcli pingfederate oauth authorization-detail-processors list` | List all authorization detail processors | [`cmd-pingcli-pingfederate-oauth-authorization-detail-processors-list.md`](cmd-pingcli-pingfederate-oauth-authorization-detail-processors-list.md) |
| `pingcli pingfederate oauth authorization-detail-processors replace` | Update an authorization detail processor | [`cmd-pingcli-pingfederate-oauth-authorization-detail-processors-replace.md`](cmd-pingcli-pingfederate-oauth-authorization-detail-processors-replace.md) |
| `pingcli pingfederate oauth authorization-detail-processors template` | Generate an authorization detail processor JSON template | [`cmd-pingcli-pingfederate-oauth-authorization-detail-processors-template.md`](cmd-pingcli-pingfederate-oauth-authorization-detail-processors-template.md) |

## Parent Command

- [`pingcli pingfederate oauth`](cmd-pingcli-pingfederate-oauth.md) — Manage PingFederate OAuth resources
