# `pingcli pingone keys replace`
Update a key

## Synopsis

Update (replace) a key in a PingOne environment. The replace body accepts only the updatable fields: default, usageType, and optionally issuerDN.

```
pingcli pingone keys replace [flags]
```

## Examples

```
# Update a key from a JSON file (--environment-id and --key-id are still required)
  pingcli pingone keys replace --environment-id <env-id> --key-id <key-id> --from-file key-update.json

  # Update a key from stdin
  pingcli pingone keys replace --environment-id <env-id> --key-id <key-id> --from-file - < key-update.json

  # Update a key from flags, without --from-file
  pingcli pingone keys replace --environment-id <env-id> --key-id <key-id> --default=true --usage-type SIGNING
```

## Options

| Flag | Default | Description |
|------|---------|-------------|
| `-h, --help` | `` | help for replace |
| `-e, --environment-id string` | `` | The PingOne environment ID |
| `-f, --from-file string` | `` | Path to a JSON file containing the request body, or "-" to read from stdin. |
| `-k, --key-id string` | `` | The key ID |
| `--default` | `` | Whether this is the default key for the environment |
| `--issuer-dn string` | `` | The distinguished name of the certificate issuer |
| `--usage-type string` | `` | The intended key usage (e.g. SIGNING, ENCRYPTION, SIGNING_AND_ENCRYPTION) |


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

- [`pingcli pingone keys`](cmd-pingcli-pingone-keys.md) — Keys
