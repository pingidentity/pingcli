# `pingcli pingfederate local-identity identity-profiles`
PingFederate Local Identity Profiles

## Synopsis

PingFederate Local Identity Profiles

```
pingcli pingfederate local-identity identity-profiles [flags]
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
| `pingcli pingfederate local-identity identity-profiles apply` | Create or update a local identity profile | [`cmd-pingcli-pingfederate-local-identity-identity-profiles-apply.md`](cmd-pingcli-pingfederate-local-identity-identity-profiles-apply.md) |
| `pingcli pingfederate local-identity identity-profiles create` | Create a new local identity profile | [`cmd-pingcli-pingfederate-local-identity-identity-profiles-create.md`](cmd-pingcli-pingfederate-local-identity-identity-profiles-create.md) |
| `pingcli pingfederate local-identity identity-profiles delete` | Delete a local identity profile | [`cmd-pingcli-pingfederate-local-identity-identity-profiles-delete.md`](cmd-pingcli-pingfederate-local-identity-identity-profiles-delete.md) |
| `pingcli pingfederate local-identity identity-profiles get` | Read a specific local identity profile | [`cmd-pingcli-pingfederate-local-identity-identity-profiles-get.md`](cmd-pingcli-pingfederate-local-identity-identity-profiles-get.md) |
| `pingcli pingfederate local-identity identity-profiles list` | List all local identity profiles | [`cmd-pingcli-pingfederate-local-identity-identity-profiles-list.md`](cmd-pingcli-pingfederate-local-identity-identity-profiles-list.md) |
| `pingcli pingfederate local-identity identity-profiles replace` | Update a local identity profile | [`cmd-pingcli-pingfederate-local-identity-identity-profiles-replace.md`](cmd-pingcli-pingfederate-local-identity-identity-profiles-replace.md) |
| `pingcli pingfederate local-identity identity-profiles template` | Generate a local identity profile JSON template | [`cmd-pingcli-pingfederate-local-identity-identity-profiles-template.md`](cmd-pingcli-pingfederate-local-identity-identity-profiles-template.md) |

## Parent Command

- [`pingcli pingfederate local-identity`](cmd-pingcli-pingfederate-local-identity.md) — Manage PingFederate Local Identity resources
