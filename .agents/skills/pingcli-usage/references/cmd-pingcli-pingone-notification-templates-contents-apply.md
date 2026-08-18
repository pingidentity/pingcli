# `pingcli pingone notification-templates contents apply`
Create or update a notification template content

## Synopsis

Idempotently create or update a notification template content looked up by the synthesised name (locale/deliveryMethod[/variant]) within the supplied --environment-id and --template-name. If no matching content exists it is created; if exactly one exists it is updated; if more than one exists the command fails.

```
pingcli pingone notification-templates contents apply [flags]
```

## Examples

```
# Create or update a notification template content
  pingcli pingone notification-templates contents apply --environment-id <env-id> --template-name <template-name> --from-file content.json

  # Read body from stdin
  pingcli pingone notification-templates contents apply --environment-id <env-id> --template-name <template-name> --from-file - < content.json

  # Create or update from a JSON file, overriding the locale and default state
  pingcli pingone notification-templates contents apply --environment-id <env-id> --template-name <template-name> --content-id <content-id> --from-file content.json --locale fr --default=false
```

## Options

| Flag | Default | Description |
|------|---------|-------------|
| `-h, --help` | `` | help for apply |
| `-c, --content-id string` | `` | The notification template content ID |
| `-e, --environment-id string` | `` | The PingOne environment ID |
| `-f, --from-file string` | `` | Path to a JSON file containing the request body, or "-" to read from stdin. |
| `-t, --template-name string` | `` | The notification template name |
| `--default` | `` | Specifies whether the template is a predefined default template |
| `--locale string` | `` | The locale for this content variant |
| `--variant string` | `` | Holds the unique user-defined name for each content variant that uses the same template + deliveryMethod + locale combination |


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


## Parent Command

- [`pingcli pingone notification-templates contents`](cmd-pingcli-pingone-notification-templates-contents.md) — Notification Template Contents
