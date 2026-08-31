# `pingcli pingfederate session authentication-session-policies replace`
Replace an authentication session policy

## Synopsis

Replace a PingFederate authentication session policy

```
pingcli pingfederate session authentication-session-policies replace [flags]
```

## Examples

```
# Replace an authentication session policy from a JSON file
  pingcli pingfederate session authentication-session-policies replace --id <id> --from-file policy.json

  # Replace an authentication session policy from stdin
  pingcli pingfederate session authentication-session-policies replace --id <id> --from-file - < policy.json

  # Replace from a JSON file, overriding scalar body fields
  pingcli pingfederate session authentication-session-policies replace --id <id> --from-file policy.json --enable-sessions --user-device-type PRIVATE --persistent --idle-timeout-mins 60 --max-timeout-mins 480 --timeout-display-unit MINUTES --authn-context-sensitive
```

## Options

| Flag | Default | Description |
|------|---------|-------------|
| `-h, --help` | `` | help for replace |
| `-f, --from-file string` | `` | Path to a JSON file containing the request body, or "-" to read from stdin. |
| `--authn-context-sensitive` | `` | Whether the policy is sensitive to authentication context |
| `--enable-sessions` | `` | Whether authentication sessions are enabled for this source |
| `--id string` | `` | The PingFederate authentication session policy ID |
| `--idle-timeout-mins int64` | `` | Idle timeout in minutes for authentication sessions |
| `--max-timeout-mins int64` | `` | Maximum timeout in minutes for authentication sessions |
| `--persistent` | `` | Whether authentication sessions are persistent |
| `--timeout-display-unit string` | `` | Display unit for authentication session timeouts |
| `--user-device-type string` | `` | The user device type for authentication sessions |


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

- [`pingcli pingfederate session authentication-session-policies`](cmd-pingcli-pingfederate-session-authentication-session-policies.md) — PingFederate authentication session policies
