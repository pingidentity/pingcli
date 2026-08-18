# `pingcli pingone keys apply`
Create or update a key

## Synopsis

Idempotently create or update a key looked up by the "name" field in the JSON body within the supplied --environment-id. If no key with the given name exists it is created; if exactly one exists it is updated; if more than one exists the command fails.

```
pingcli pingone keys apply [flags]
```

## Examples

```
# Create or update a key (body supplies name and required fields)
  pingcli pingone keys apply --environment-id <env-id> --from-file key.json

  # Read body from stdin
  pingcli pingone keys apply --environment-id <env-id> --from-file - < key.json

  # Create or update a key from flags, without --from-file
  pingcli pingone keys apply --environment-id <env-id> --algorithm RSA --key-length 2048 --name "My Key" --signature-algorithm SHA256withRSA --subject-dn "CN=example.com" --usage-type SIGNING --validity-period 365
```

## Options

| Flag | Default | Description |
|------|---------|-------------|
| `-h, --help` | `` | help for apply |
| `-e, --environment-id string` | `` | The PingOne environment ID |
| `-f, --from-file string` | `` | Path to a JSON file containing the request body, or "-" to read from stdin. |
| `-k, --key-id string` | `` | The key ID |
| `--algorithm string` | `` | The key algorithm (e.g. RSA, EC) |
| `--custom-crl string` | `` | A custom Certificate Revocation List endpoint URL, used for certificates of type ISSUANCE |
| `--default` | `` | Whether this is the default key for the environment |
| `--issuer-dn string` | `` | The distinguished name of the certificate issuer |
| `--key-length int64` | `` | The key length in bits (e.g. 2048, 4096 for RSA; 256, 384 for EC) |
| `--name string` | `` | The display name for the key |
| `--signature-algorithm string` | `` | The signature algorithm (e.g. SHA256withRSA, SHA256withECDSA) |
| `--subject-dn string` | `` | The distinguished name for the generated certificate (e.g. CN=example.com) |
| `--usage-type string` | `` | The intended key usage (e.g. SIGNING, ENCRYPTION, SIGNING_AND_ENCRYPTION) |
| `--validity-period int64` | `` | The number of days the key is valid |


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
