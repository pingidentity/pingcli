# `pingcli pingfederate oauth idp-adapter-mappings`
PingFederate OAuth IdP Adapter Mappings

## Synopsis

PingFederate OAuth IdP Adapter Mappings

```
pingcli pingfederate oauth idp-adapter-mappings [flags]
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
| `pingcli pingfederate oauth idp-adapter-mappings apply` | Create or update an IdP adapter mapping | [`cmd-pingcli-pingfederate-oauth-idp-adapter-mappings-apply.md`](cmd-pingcli-pingfederate-oauth-idp-adapter-mappings-apply.md) |
| `pingcli pingfederate oauth idp-adapter-mappings create` | Create a new IdP adapter mapping | [`cmd-pingcli-pingfederate-oauth-idp-adapter-mappings-create.md`](cmd-pingcli-pingfederate-oauth-idp-adapter-mappings-create.md) |
| `pingcli pingfederate oauth idp-adapter-mappings delete` | Delete an IdP adapter mapping | [`cmd-pingcli-pingfederate-oauth-idp-adapter-mappings-delete.md`](cmd-pingcli-pingfederate-oauth-idp-adapter-mappings-delete.md) |
| `pingcli pingfederate oauth idp-adapter-mappings get` | Read a specific IdP adapter mapping | [`cmd-pingcli-pingfederate-oauth-idp-adapter-mappings-get.md`](cmd-pingcli-pingfederate-oauth-idp-adapter-mappings-get.md) |
| `pingcli pingfederate oauth idp-adapter-mappings list` | List all IdP adapter mappings | [`cmd-pingcli-pingfederate-oauth-idp-adapter-mappings-list.md`](cmd-pingcli-pingfederate-oauth-idp-adapter-mappings-list.md) |
| `pingcli pingfederate oauth idp-adapter-mappings replace` | Update an IdP adapter mapping | [`cmd-pingcli-pingfederate-oauth-idp-adapter-mappings-replace.md`](cmd-pingcli-pingfederate-oauth-idp-adapter-mappings-replace.md) |
| `pingcli pingfederate oauth idp-adapter-mappings template` | Generate an IdP adapter mapping JSON template | [`cmd-pingcli-pingfederate-oauth-idp-adapter-mappings-template.md`](cmd-pingcli-pingfederate-oauth-idp-adapter-mappings-template.md) |

## Parent Command

- [`pingcli pingfederate oauth`](cmd-pingcli-pingfederate-oauth.md) — Manage PingFederate OAuth resources
