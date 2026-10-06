# `pingcli advanced-services pingfederate oauth resource-owner-credentials-mappings`
PingFederate OAuth Resource Owner Credentials Mappings

## Synopsis

PingFederate OAuth Resource Owner Credentials Mappings

```
pingcli advanced-services pingfederate oauth resource-owner-credentials-mappings [flags]
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
| `pingcli advanced-services pingfederate oauth resource-owner-credentials-mappings apply` | Create or update a resource owner credentials mapping | [`cmd-pingcli-advanced-services-pingfederate-oauth-resource-owner-credentials-mappings-apply.md`](cmd-pingcli-advanced-services-pingfederate-oauth-resource-owner-credentials-mappings-apply.md) |
| `pingcli advanced-services pingfederate oauth resource-owner-credentials-mappings create` | Create a new resource owner credentials mapping | [`cmd-pingcli-advanced-services-pingfederate-oauth-resource-owner-credentials-mappings-create.md`](cmd-pingcli-advanced-services-pingfederate-oauth-resource-owner-credentials-mappings-create.md) |
| `pingcli advanced-services pingfederate oauth resource-owner-credentials-mappings delete` | Delete a resource owner credentials mapping | [`cmd-pingcli-advanced-services-pingfederate-oauth-resource-owner-credentials-mappings-delete.md`](cmd-pingcli-advanced-services-pingfederate-oauth-resource-owner-credentials-mappings-delete.md) |
| `pingcli advanced-services pingfederate oauth resource-owner-credentials-mappings get` | Read a specific resource owner credentials mapping | [`cmd-pingcli-advanced-services-pingfederate-oauth-resource-owner-credentials-mappings-get.md`](cmd-pingcli-advanced-services-pingfederate-oauth-resource-owner-credentials-mappings-get.md) |
| `pingcli advanced-services pingfederate oauth resource-owner-credentials-mappings list` | List all resource owner credentials mappings | [`cmd-pingcli-advanced-services-pingfederate-oauth-resource-owner-credentials-mappings-list.md`](cmd-pingcli-advanced-services-pingfederate-oauth-resource-owner-credentials-mappings-list.md) |
| `pingcli advanced-services pingfederate oauth resource-owner-credentials-mappings replace` | Update a resource owner credentials mapping | [`cmd-pingcli-advanced-services-pingfederate-oauth-resource-owner-credentials-mappings-replace.md`](cmd-pingcli-advanced-services-pingfederate-oauth-resource-owner-credentials-mappings-replace.md) |
| `pingcli advanced-services pingfederate oauth resource-owner-credentials-mappings template` | Generate a resource owner credentials mapping JSON template | [`cmd-pingcli-advanced-services-pingfederate-oauth-resource-owner-credentials-mappings-template.md`](cmd-pingcli-advanced-services-pingfederate-oauth-resource-owner-credentials-mappings-template.md) |

## Parent Command

- [`pingcli advanced-services pingfederate oauth`](cmd-pingcli-advanced-services-pingfederate-oauth.md) — Manage PingFederate OAuth resources
