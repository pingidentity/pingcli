# `pingcli pingone keys template`
Generate a key JSON template

## Synopsis

Generate a JSON skeleton template for key create or replace bodies

```
pingcli pingone keys template [flags]
```

## Examples

```
# Generate a create template (default)
  pingcli pingone keys template > body.json

  # Generate a replace template (only mutable fields)
  pingcli pingone keys template --template-type replace > replace.json

  # Edit the template, then create or replace
  pingcli pingone keys create --environment-id <env-id> --from-file body.json
  pingcli pingone keys replace --environment-id <env-id> --key-id <key-id> --from-file replace.json
Use the JSON template as a starting point:
  1. Run the template command to generate the body skeleton.
  2. Edit the file, replacing placeholder values with real data.
  3. Feed the edited file back to the create or replace action via --from-file.

Example workflow:
  pingcli pingone keys template > body.json
  # edit body.json
  pingcli pingone keys create --from-file body.json
```

## Options

| Flag | Default | Description |
|------|---------|-------------|
| `-h, --help` | `` | help for template |
| `-o, --output-file string` | `` | Write the JSON template to PATH instead of stdout. Overwrites any existing file. |
| `--template-type string` | `` | The variant of the template to generate. One of: create, replace. Defaults to "create". |


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
