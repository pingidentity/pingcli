# `pingcli advanced-services pingfederate administrative-accounts reset-password`
Reset a PingFederate administrative account password

## Synopsis

Reset the password for a PingFederate administrative account. Supply the new password via --new-password, --from-file (decoded as a UserCredentials JSON body; only "newPassword" is read), or both — --new-password overrides the file's value. At least one must effectively supply a non-empty new password.

```
pingcli advanced-services pingfederate administrative-accounts reset-password [flags]
```

## Examples

```
# Reset password for an administrative account
  pingcli advanced-services pingfederate administrative-accounts reset-password --username <username> --new-password <new-password>

  # Reset password for an administrative account from a JSON file
  pingcli advanced-services pingfederate administrative-accounts reset-password --username <username> --from-file password.json
```

## Options

| Flag | Default | Description |
|------|---------|-------------|
| `-h, --help` | `` | help for reset-password |
| `-f, --from-file string` | `` | Path to a JSON file containing a UserCredentials body (the "newPassword" key is used; "currentPassword" is ignored) for the reset-password request, or "-" to read from stdin. --new-password overrides the file's newPassword if both are supplied. |
| `-u, --username string` | `` | The PingFederate administrative account username |
| `--new-password string` | `` | The new password for the PingFederate administrative account |


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

- [`pingcli advanced-services pingfederate administrative-accounts`](cmd-pingcli-advanced-services-pingfederate-administrative-accounts.md) — PingFederate Administrative Accounts
