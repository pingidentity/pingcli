# `pingcli advanced-services pingfederate kerberos realms replace`
Update a Kerberos realm

## Synopsis

Update (replace) a PingFederate Kerberos realm

```
pingcli advanced-services pingfederate kerberos realms replace [flags]
```

## Examples

```
# Update a Kerberos realm from a JSON file (--id is still required)
  pingcli advanced-services pingfederate kerberos realms replace --id <id> --from-file realm.json

  # Update a Kerberos realm from stdin
  pingcli advanced-services pingfederate kerberos realms replace --id <id> --from-file - < realm.json

  # Replace from a JSON file, overriding the Kerberos username
  pingcli advanced-services pingfederate kerberos realms replace --id <id> --from-file realm.json --kerberos-username admin
```

## Options

| Flag | Default | Description |
|------|---------|-------------|
| `-h, --help` | `` | help for replace |
| `-f, --from-file string` | `` | Path to a JSON file containing the request body, or "-" to read from stdin. |
| `--connection-type string` | `` | Controls how PingFederate connects to the Active Directory/Kerberos Realm. One of DIRECT, LDAP_GATEWAY, or LOCAL_VALIDATION. Defaults to DIRECT |
| `--id string` | `` | The persistent, unique ID for the Kerberos Realm |
| `--kerberos-realm-name string` | `` | The Domain/Realm name used for display in UI screens |
| `--kerberos-username string` | `` | The Domain/Realm username. Only required when connection-type is DIRECT or LOCAL_VALIDATION |
| `--key-distribution-centers []string` | `` | The Domain Controller/Key Distribution Center host names; repeatable or comma-separated. Only applicable when connection-type is DIRECT |
| `--ldap-gateway-data-store-ref-id string` | `` | ID of the Data Store (PING_ONE_LDAP_GATEWAY type) used when connection-type is LDAP_GATEWAY |
| `--retain-previous-keys-on-password-change` | `` | Determines whether the previous encryption keys are retained when the password is updated, allowing existing Kerberos tickets to continue to be validated. Defaults to false |
| `--suppress-domain-name-concatenation` | `` | Controls whether the KDC hostnames and the realm name are concatenated in the auto-generated krb5.conf file. Defaults to false. Only applicable when connection-type is DIRECT |


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

- [`pingcli advanced-services pingfederate kerberos realms`](cmd-pingcli-advanced-services-pingfederate-kerberos-realms.md) — PingFederate Kerberos Realms
