# `pingcli pingone webhooks`
Webhooks

## Synopsis

Webhooks

```
pingcli pingone webhooks [flags]
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
| `pingcli pingone webhooks apply` | Create or update a webhook | [`cmd-pingcli-pingone-webhooks-apply.md`](cmd-pingcli-pingone-webhooks-apply.md) |
| `pingcli pingone webhooks create` | Create a new webhook | [`cmd-pingcli-pingone-webhooks-create.md`](cmd-pingcli-pingone-webhooks-create.md) |
| `pingcli pingone webhooks delete` | Delete a webhook | [`cmd-pingcli-pingone-webhooks-delete.md`](cmd-pingcli-pingone-webhooks-delete.md) |
| `pingcli pingone webhooks get` | Read a specific webhook | [`cmd-pingcli-pingone-webhooks-get.md`](cmd-pingcli-pingone-webhooks-get.md) |
| `pingcli pingone webhooks list` | List all webhooks | [`cmd-pingcli-pingone-webhooks-list.md`](cmd-pingcli-pingone-webhooks-list.md) |
| `pingcli pingone webhooks replace` | Update a webhook | [`cmd-pingcli-pingone-webhooks-replace.md`](cmd-pingcli-pingone-webhooks-replace.md) |
| `pingcli pingone webhooks template` | Generate a webhook JSON template | [`cmd-pingcli-pingone-webhooks-template.md`](cmd-pingcli-pingone-webhooks-template.md) |

## Parent Command

- [`pingcli pingone`](cmd-pingcli-pingone.md) — Administration tools for the PingOne platform.
