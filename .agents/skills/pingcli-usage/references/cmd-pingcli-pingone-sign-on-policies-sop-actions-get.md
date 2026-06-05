# `pingcli pingone sign-on-policies sop-actions get`
Read a specific sign-on policy action

## Synopsis

Read a specific sign-on policy action in a PingOne sign-on policy

```
pingcli pingone sign-on-policies sop-actions get [flags]
```

## Examples

```
# Read a specific sign-on policy action
  pingcli pingone sign-on-policies sop-actions get --environment-id <env-id> --sign-on-policy-id <sop-id> --sign-on-policy-action-id <action-id>
```

## Options

| Flag | Default | Description |
|------|---------|-------------|
| `-h, --help` | `` | help for get |
| `-a, --sign-on-policy-action-id string` | `` | The sign-on policy action ID |
| `-e, --environment-id string` | `` | The PingOne environment ID |
| `-s, --sign-on-policy-id string` | `` | The sign-on policy ID |
| `--template string` | `` | A Go text/template string. When provided, the command output is rendered through the template instead of the default format. The template receives the command's structured response data. Example: --template '{{.Name}}' |


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


## Parent Command

- [`pingcli pingone sign-on-policies sop-actions`](cmd-pingcli-pingone-sign-on-policies-sop-actions.md) — Sign-On Policy Actions
