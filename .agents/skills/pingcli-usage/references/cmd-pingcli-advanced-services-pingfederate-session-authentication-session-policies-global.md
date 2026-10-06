# `pingcli advanced-services pingfederate session authentication-session-policies global`
PingFederate global authentication session policy

## Synopsis

PingFederate Global Authentication Session Policy

```
pingcli advanced-services pingfederate session authentication-session-policies global [flags]
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
| `pingcli advanced-services pingfederate session authentication-session-policies global apply` | Apply the global authentication session policy | [`cmd-pingcli-advanced-services-pingfederate-session-authentication-session-policies-global-apply.md`](cmd-pingcli-advanced-services-pingfederate-session-authentication-session-policies-global-apply.md) |
| `pingcli advanced-services pingfederate session authentication-session-policies global get` | Read the global authentication session policy | [`cmd-pingcli-advanced-services-pingfederate-session-authentication-session-policies-global-get.md`](cmd-pingcli-advanced-services-pingfederate-session-authentication-session-policies-global-get.md) |
| `pingcli advanced-services pingfederate session authentication-session-policies global replace` | Replace the global authentication session policy | [`cmd-pingcli-advanced-services-pingfederate-session-authentication-session-policies-global-replace.md`](cmd-pingcli-advanced-services-pingfederate-session-authentication-session-policies-global-replace.md) |
| `pingcli advanced-services pingfederate session authentication-session-policies global template` | Generate a global authentication session policy JSON template | [`cmd-pingcli-advanced-services-pingfederate-session-authentication-session-policies-global-template.md`](cmd-pingcli-advanced-services-pingfederate-session-authentication-session-policies-global-template.md) |

## Parent Command

- [`pingcli advanced-services pingfederate session authentication-session-policies`](cmd-pingcli-advanced-services-pingfederate-session-authentication-session-policies.md) — PingFederate authentication session policies
