# `pingcli pingone notification-policies template`
Generate a notification policy JSON template

## Synopsis

Generate a JSON skeleton template for notification policy create or replace bodies

```
pingcli pingone notification-policies template [flags]
```

## Examples

```
# Generate a template and save to a file
  pingcli pingone notification-policies template > body.json

  # Edit the template, then create the notification policy
  pingcli pingone notification-policies create --environment-id <env-id> --from-file body.json
Use the JSON template as a starting point:
  1. Run the template command to generate the body skeleton.
  2. Edit the file, replacing placeholder values with real data.
  3. Feed the edited file back to the create or replace action via --from-file.

Example workflow:
  pingcli pingone notification-policies template > body.json
  # edit body.json
  pingcli pingone notification-policies create --from-file body.json
```

## Options

| Flag | Default | Description |
|------|---------|-------------|
| `-h, --help` | `` | help for template |
| `-o, --output-file string` | `` | Write the JSON template to PATH instead of stdout. Overwrites any existing file. |


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

- [`pingcli pingone notification-policies`](cmd-pingcli-pingone-notification-policies.md) — Notification Policies
