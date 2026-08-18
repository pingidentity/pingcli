# `pingcli pingfederate key-pairs signing template`
Generate a signing key pair create, import, import-csr-response, or link JSON template

## Synopsis

Generate a JSON skeleton template for a signing key pair body. Use --template-type to select create (default), import, import-csr-response, or link.

```
pingcli pingfederate key-pairs signing template [flags]
```

## Examples

```
# Generate a create template (default) and save to a file
  pingcli pingfederate key-pairs signing template > body.json

  # Generate an import, import-csr-response, or link template
  pingcli pingfederate key-pairs signing template --template-type import > key-pair-file.json
  pingcli pingfederate key-pairs signing template --template-type import-csr-response > csr-response.json
  pingcli pingfederate key-pairs signing template --template-type link > key-pair-link.json

  # Edit the template, then run the matching action
  pingcli pingfederate key-pairs signing create --from-file body.json
  pingcli pingfederate key-pairs signing import --from-file key-pair-file.json
  pingcli pingfederate key-pairs signing import-csr-response --id <id> --from-file csr-response.json
  pingcli pingfederate key-pairs signing link --from-file key-pair-link.json
Use the JSON template as a starting point:
  1. Run the template command to generate the body skeleton.
  2. Edit the file, replacing placeholder values with real data.
  3. Feed the edited file back to the create action via --from-file.

Example workflow:
  pingcli pingfederate key-pairs signing template > body.json
  # edit body.json
  pingcli pingfederate key-pairs signing create --from-file body.json
```

## Options

| Flag | Default | Description |
|------|---------|-------------|
| `-h, --help` | `` | help for template |
| `-o, --output-file string` | `` | Write the JSON template to PATH instead of stdout. Overwrites any existing file. |
| `--template-type string` | `` | The variant of the template to generate. One of: create, import, import-csr-response, link. Defaults to "create". |


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
