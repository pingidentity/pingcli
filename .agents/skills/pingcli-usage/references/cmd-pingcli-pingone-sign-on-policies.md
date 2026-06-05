# `pingcli pingone sign-on-policies`
Sign-On Policies

## Synopsis

Sign-On Policies

```
pingcli pingone sign-on-policies [flags]
```

## Inherited Options

| Flag | Default | Description |
|------|---------|-------------|
| `-C, --config string` | `` | The relative or full path to a custom Ping CLI configuration file. (default $HOME/.pingcli/config.yaml) |
| `-D, --detailed-exitcode` | `` | Enable detailed exit code output. (default false) 0 - pingcli command succeeded with no errors or warnings. 1 - pingcli command failed with errors. 2 - pingcli command succeeded with warnings. |
| `-O, --output-format string` | `` | Specify the console output format. (default text) Options are: json, ndjson, ndjson-wrapped, text. |
| `-P, --profile string` | `` | The name of a configuration profile to use. |
| `--debug` | `` | Enable debug output for error messages, including stack traces and transaction IDs. (default false) |
| `--log-file string` | `` | Write logs to a file at the given path. File logging is disabled when not set. |
| `--log-file-level string` | `` | Set the file log level. Options are: DEBUG, INFO, WARN, ERROR. (default DEBUG) |
| `--log-level string` | `` | Set the console log level. Options are: DEBUG, INFO, WARN, ERROR. (default WARN) |
| `--no-color` | `` | Disable text output in color. (default false) |
| `--query string` | `` | JMESPath expression to filter JSON output. Requires -O json, ndjson, or ndjson-wrapped. Example: --query 'data[?enabled].name' |


## Subcommands

| Command | Description | Reference |
|---------|-------------|----------|
| `pingcli pingone sign-on-policies apply` | Create or update a sign-on policy | [`cmd-pingcli-pingone-sign-on-policies-apply.md`](cmd-pingcli-pingone-sign-on-policies-apply.md) |
| `pingcli pingone sign-on-policies create` | Create a new sign-on policy | [`cmd-pingcli-pingone-sign-on-policies-create.md`](cmd-pingcli-pingone-sign-on-policies-create.md) |
| `pingcli pingone sign-on-policies delete` | Delete a sign-on policy | [`cmd-pingcli-pingone-sign-on-policies-delete.md`](cmd-pingcli-pingone-sign-on-policies-delete.md) |
| `pingcli pingone sign-on-policies get` | Read a specific sign-on policy | [`cmd-pingcli-pingone-sign-on-policies-get.md`](cmd-pingcli-pingone-sign-on-policies-get.md) |
| `pingcli pingone sign-on-policies list` | List all sign-on policies | [`cmd-pingcli-pingone-sign-on-policies-list.md`](cmd-pingcli-pingone-sign-on-policies-list.md) |
| `pingcli pingone sign-on-policies replace` | Update a sign-on policy | [`cmd-pingcli-pingone-sign-on-policies-replace.md`](cmd-pingcli-pingone-sign-on-policies-replace.md) |
| `pingcli pingone sign-on-policies sop-actions` | Sign-On Policy Actions | [`cmd-pingcli-pingone-sign-on-policies-sop-actions.md`](cmd-pingcli-pingone-sign-on-policies-sop-actions.md) |
| `pingcli pingone sign-on-policies template` | Generate a sign-on policy JSON template | [`cmd-pingcli-pingone-sign-on-policies-template.md`](cmd-pingcli-pingone-sign-on-policies-template.md) |

## Parent Command

- [`pingcli pingone`](cmd-pingcli-pingone.md) — Administration tools for the PingOne platform.
