# `pingcli pingfederate administrative-accounts create`
Create a new administrative account

## Synopsis

Create a new PingFederate administrative account

```
pingcli pingfederate administrative-accounts create [flags]
```

## Examples

```
# Create a new administrative account from a JSON file
  pingcli pingfederate administrative-accounts create --from-file account.json

  # Create a new administrative account from stdin
  pingcli pingfederate administrative-accounts create --from-file - < account.json

  # Create a new administrative account, overriding scalar fields with flags
  pingcli pingfederate administrative-accounts create --username jdoe --description "Support engineer" --active --roles ADMINISTRATOR,CRYPTO_ADMINISTRATOR --from-file password.json
```

## Options

| Flag | Default | Description |
|------|---------|-------------|
| `-h, --help` | `` | help for create |
| `-f, --from-file string` | `` | Path to a JSON file containing the request body, or "-" to read from stdin. |
| `-u, --username string` | `` | The PingFederate administrative account username |
| `--active` | `` | Whether the account is active |
| `--auditor` | `` | Whether the account is an auditor |
| `--department string` | `` | Department name of the account user |
| `--description string` | `` | Description of the account |
| `--email-address string` | `` | Email address associated with the account |
| `--phone-number string` | `` | Phone number associated with the account |
| `--roles []string` | `` | Administrator roles; repeatable or comma-separated |


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

- [`pingcli pingfederate administrative-accounts`](cmd-pingcli-pingfederate-administrative-accounts.md) — PingFederate Administrative Accounts
