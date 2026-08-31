# `pingcli pingfederate session authentication-session-policies global replace`
Replace the global authentication session policy

## Synopsis

Replace the PingFederate global authentication session policy

```
pingcli pingfederate session authentication-session-policies global replace [flags]
```

## Examples

```
# Replace the global authentication session policy from a JSON file
  pingcli pingfederate session authentication-session-policies global replace --from-file policy.json

  # Replace the global policy from stdin
  pingcli pingfederate session authentication-session-policies global replace --from-file - < policy.json

  # Replace the global policy from flags without --from-file
  pingcli pingfederate session authentication-session-policies global replace --enable-sessions --persistent-sessions --hash-unique-user-key-attribute --idle-timeout-mins 60 --idle-timeout-display-unit MINUTES --max-timeout-mins 480 --max-timeout-display-unit MINUTES
```

## Options

| Flag | Default | Description |
|------|---------|-------------|
| `-h, --help` | `` | help for replace |
| `-f, --from-file string` | `` | Path to a JSON file containing the request body, or "-" to read from stdin. |
| `--enable-sessions` | `` | Whether PingFederate authentication sessions are enabled (required when --from-file is omitted) |
| `--hash-unique-user-key-attribute` | `` | Whether the unique user key attribute is hashed |
| `--idle-timeout-display-unit string` | `` | Display unit for the idle authentication session timeout |
| `--idle-timeout-mins int64` | `` | Idle timeout in minutes for authentication sessions |
| `--max-timeout-display-unit string` | `` | Display unit for the maximum authentication session timeout |
| `--max-timeout-mins int64` | `` | Maximum timeout in minutes for authentication sessions |
| `--persistent-sessions` | `` | Whether authentication sessions persist across requests |


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

- [`pingcli pingfederate session authentication-session-policies global`](cmd-pingcli-pingfederate-session-authentication-session-policies-global.md) — PingFederate global authentication session policy
