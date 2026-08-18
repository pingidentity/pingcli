# `pingcli pingone users create`
Create a new user

## Synopsis

Create a new user in a PingOne environment

```
pingcli pingone users create [flags]
```

## Examples

```
# Create a new user from a JSON file
  pingcli pingone users create --environment-id <env-id> --from-file user.json

  # Create a new user from stdin
  pingcli pingone users create --environment-id <env-id> --from-file - < user.json

  # Create a new user from a JSON file, overriding fields with flags
  pingcli pingone users create --environment-id <env-id> --from-file user.json --nickname "Jamie" --population-id <pop-id>
```

## Options

| Flag | Default | Description |
|------|---------|-------------|
| `-h, --help` | `` | help for create |
| `-e, --environment-id string` | `` | The PingOne environment ID |
| `-f, --from-file string` | `` | Path to a JSON file containing the request body, or "-" to read from stdin. |
| `--email string` | `` | The email address of the user |
| `--locale string` | `` | The locale tag for the user |
| `--mfa-enabled` | `` | Whether MFA is enabled for the user (create only — not writable on update) |
| `--mobile-phone string` | `` | The mobile phone number for the user |
| `--nickname string` | `` | The nickname for the user |
| `--population-id string` | `` | ID of the population to place the user in |
| `--preferred-language string` | `` | The preferred language for the user |
| `--primary-phone string` | `` | The primary phone number for the user |
| `--timezone string` | `` | The timezone for the user |
| `--title string` | `` | The title or honorific for the user |
| `--type string` | `` | The user type classification |
| `--username string` | `` | The username for the user account |


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

- [`pingcli pingone users`](cmd-pingcli-pingone-users.md) — Users
