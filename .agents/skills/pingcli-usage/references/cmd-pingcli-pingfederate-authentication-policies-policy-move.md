# `pingcli pingfederate authentication-policies policy move`
Move an authentication policy to a location within the policy tree

## Synopsis

Move an authentication policy to a location within the policy tree. Location must be START, END, BEFORE, or AFTER; --move-to-id is required for BEFORE and AFTER.

```
pingcli pingfederate authentication-policies policy move [flags]
```

## Examples

```
# Move an authentication policy to the start of the list
  pingcli pingfederate authentication-policies policy move --id policy-id --location START

  # Move an authentication policy before another policy
  pingcli pingfederate authentication-policies policy move --id policy-id --location BEFORE --move-to-id target-policy-id
```

## Options

| Flag | Default | Description |
|------|---------|-------------|
| `-h, --help` | `` | help for move |
| `--id string` | `` | The PingFederate authentication policy ID |
| `--location string` | `` | Position for the policy: START, END, BEFORE, or AFTER |
| `--move-to-id string` | `` | Target policy ID when location is BEFORE or AFTER |


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

- [`pingcli pingfederate authentication-policies policy`](cmd-pingcli-pingfederate-authentication-policies-policy.md) — Manage an authentication policy tree
