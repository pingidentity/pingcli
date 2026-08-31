# `pingcli pingfederate idp token-processors`
PingFederate Token Processors

## Synopsis

PingFederate Token Processors

```
pingcli pingfederate idp token-processors [flags]
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
| `pingcli pingfederate idp token-processors apply` | Create or update a token processor | [`cmd-pingcli-pingfederate-idp-token-processors-apply.md`](cmd-pingcli-pingfederate-idp-token-processors-apply.md) |
| `pingcli pingfederate idp token-processors create` | Create a new token processor | [`cmd-pingcli-pingfederate-idp-token-processors-create.md`](cmd-pingcli-pingfederate-idp-token-processors-create.md) |
| `pingcli pingfederate idp token-processors delete` | Delete a token processor | [`cmd-pingcli-pingfederate-idp-token-processors-delete.md`](cmd-pingcli-pingfederate-idp-token-processors-delete.md) |
| `pingcli pingfederate idp token-processors descriptors` | PingFederate Token Processor Descriptors | [`cmd-pingcli-pingfederate-idp-token-processors-descriptors.md`](cmd-pingcli-pingfederate-idp-token-processors-descriptors.md) |
| `pingcli pingfederate idp token-processors get` | Read a specific token processor | [`cmd-pingcli-pingfederate-idp-token-processors-get.md`](cmd-pingcli-pingfederate-idp-token-processors-get.md) |
| `pingcli pingfederate idp token-processors list` | List all token processors | [`cmd-pingcli-pingfederate-idp-token-processors-list.md`](cmd-pingcli-pingfederate-idp-token-processors-list.md) |
| `pingcli pingfederate idp token-processors replace` | Update a token processor | [`cmd-pingcli-pingfederate-idp-token-processors-replace.md`](cmd-pingcli-pingfederate-idp-token-processors-replace.md) |
| `pingcli pingfederate idp token-processors template` | Generate a token processor JSON template | [`cmd-pingcli-pingfederate-idp-token-processors-template.md`](cmd-pingcli-pingfederate-idp-token-processors-template.md) |

## Parent Command

- [`pingcli pingfederate idp`](cmd-pingcli-pingfederate-idp.md) — Manage PingFederate IdP resources
