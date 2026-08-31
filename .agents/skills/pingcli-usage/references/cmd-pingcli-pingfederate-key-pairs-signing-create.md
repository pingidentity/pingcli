# `pingcli pingfederate key-pairs signing create`
Generate a new signing key pair

## Synopsis

Generate a new PingFederate signing key pair

```
pingcli pingfederate key-pairs signing create [flags]
```

## Examples

```
# Generate a new signing key pair from flags
  pingcli pingfederate key-pairs signing create --id my-signing-key --common-name example.com --subject-alternative-names example.com,www.example.com --organization "Example Corp" --organization-unit Engineering --city Denver --state Colorado --country US --valid-days 365 --key-algorithm RSA --key-size 2048 --signature-algorithm SHA256withRSA

  # Generate a new signing key pair from a JSON file
  pingcli pingfederate key-pairs signing create --from-file key-pair.json

  # Generate a new signing key pair from stdin
  pingcli pingfederate key-pairs signing create --from-file - < key-pair.json
```

## Options

| Flag | Default | Description |
|------|---------|-------------|
| `-h, --help` | `` | help for create |
| `-f, --from-file string` | `` | Path to a JSON file containing the request body, or "-" to read from stdin. |
| `--city string` | `` | City for the key pair subject |
| `--common-name string` | `` | Common name for the key pair subject |
| `--country string` | `` | Country for the key pair subject |
| `--crypto-provider string` | `` | Cryptographic provider used in Hybrid HSM mode |
| `--id string` | `` | The persistent, unique ID of the PingFederate signing key pair |
| `--key-algorithm string` | `` | Key generation algorithm |
| `--key-size int64` | `` | Key size in bits |
| `--organization string` | `` | Organization for the key pair subject |
| `--organization-unit string` | `` | Organization unit for the key pair subject |
| `--signature-algorithm string` | `` | Signature algorithm |
| `--state string` | `` | State for the key pair subject |
| `--subject-alternative-names []string` | `` | Subject alternative names for the key pair |
| `--valid-days int64` | `` | Number of days the key pair will be valid |


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

- [`pingcli pingfederate key-pairs signing`](cmd-pingcli-pingfederate-key-pairs-signing.md) — PingFederate Signing Key Pairs
