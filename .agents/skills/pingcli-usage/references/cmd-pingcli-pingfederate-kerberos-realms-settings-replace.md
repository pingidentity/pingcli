# `pingcli pingfederate kerberos realms settings replace`
Update Kerberos Realms Settings

## Synopsis

Update (replace) the PingFederate Kerberos Realms Settings

```
pingcli pingfederate kerberos realms settings replace [flags]
```

## Examples

```
# Update Kerberos Realms Settings from a JSON file
  pingcli pingfederate kerberos realms settings replace --from-file settings.json

  # Update Kerberos Realms Settings from stdin
  pingcli pingfederate kerberos realms settings replace --from-file - < settings.json

  # Update using flags only, without --from-file
  pingcli pingfederate kerberos realms settings replace --kdc-retries 4 --kdc-timeout 30 --force-tcp --debug-log-output --key-set-retention-period-mins 610
```

## Options

| Flag | Default | Description |
|------|---------|-------------|
| `-h, --help` | `` | help for replace |
| `-f, --from-file string` | `` | Path to a JSON file containing the request body, or "-" to read from stdin. |
| `--debug-log-output` | `` | Enable debug logging for Kerberos Realms operations |
| `--force-tcp` | `` | Force Kerberos communication with the Key Distribution Center(s) to use TCP instead of UDP |
| `--kdc-retries string` | `` | The default number of retries for a Key Distribution Center connection attempt |
| `--kdc-timeout string` | `` | The default Key Distribution Center connection timeout, in seconds |
| `--key-set-retention-period-mins int64` | `` | Minutes a previous encryption key set is retained after a password change when retainPreviousKeysOnPasswordChange is enabled for a realm; PingFederate defaults to 610 when omitted |


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

- [`pingcli pingfederate kerberos realms settings`](cmd-pingcli-pingfederate-kerberos-realms-settings.md) — PingFederate Kerberos Realms Settings
