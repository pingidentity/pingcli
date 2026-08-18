# `pingcli pingfederate connection-metadata convert`
Convert SAML connection metadata into a JSON connection

## Synopsis

Convert SAML connection metadata into a JSON PingFederate connection. --from-file must contain a JSON object with required fields: connectionType (IDP or SP), expectedProtocol (SAML20, SAML11, or SAML10), samlMetadata (the SAML metadata XML, base64-encoded). Optional fields: expectedEntityId, verificationCertificate (PEM or base-64 DER), templateConnection (a Connection object to seed the conversion).

```
pingcli pingfederate connection-metadata convert [flags]
```

## Examples

```
# Convert SAML metadata into a JSON connection
  pingcli pingfederate connection-metadata convert --from-file convert.json

  # Convert SAML metadata read from stdin
  pingcli pingfederate connection-metadata convert --from-file - < convert.json
```

## Options

| Flag | Default | Description |
|------|---------|-------------|
| `-h, --help` | `` | help for convert |
| `-f, --from-file string` | `` | Path to a JSON file containing the connection metadata convert request body, or "-" to read from stdin. |


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

- [`pingcli pingfederate connection-metadata`](cmd-pingcli-pingfederate-connection-metadata.md) — Manage PingFederate connection metadata
