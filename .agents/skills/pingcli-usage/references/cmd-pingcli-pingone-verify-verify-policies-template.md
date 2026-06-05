# `pingcli pingone verify verify-policies template`
Generate a verify policy JSON template

## Synopsis

Generate a JSON skeleton template for verify policy create or replace bodies.

The template includes all optional nested configuration blocks: governmentId,
facialComparison, liveness, email, phone, transaction, and voice. Remove the
blocks you do not need before passing the body to create or replace.

Note: the "voice" field is deprecated in the PingOne Verify API. Remove it from
your --from-file bodies unless you have an explicit requirement for voice
verification. The template includes it for completeness.

Server-set read-only fields (id, createdAt, updatedAt, environment, _links) are
excluded from the template. User-settable optional fields (description, default,
and all nested config blocks) are included with zero values.

```
pingcli pingone verify verify-policies template [flags]
```

## Examples

```
# Generate a template and save to a file
  pingcli pingone verify verify-policies template > body.json

  # Edit the template — fill in the required name and at least one verification
  # block with valid enum values, then pass the body to create or replace
Use the JSON template as a starting point:
  1. Run the template command to generate the body skeleton.
  2. Edit the file, replacing placeholder values with real data.
  3. Feed the edited file back to the create or replace action via --from-file.

Example workflow:
  pingcli pingone verify verify-policies template > body.json
  # edit body.json
  pingcli pingone verify verify-policies create --from-file body.json
```

## Options

| Flag | Default | Description |
|------|---------|-------------|
| `-h, --help` | `` | help for template |
| `-o, --output-file string` | `` | Write the JSON template to PATH instead of stdout. Overwrites any existing file. |


## Inherited Options

| Flag | Default | Description |
|------|---------|-------------|
| `-C, --config string` | `` | The relative or full path to a custom Ping CLI configuration file. (default $HOME/.pingcli/config.yaml) |
| `-D, --detailed-exitcode` | `` | Enable detailed exit code output. (default false) 0 - pingcli command succeeded with no errors or warnings. 1 - pingcli command failed with errors. 2 - pingcli command succeeded with warnings. |
| `-O, --output-format string` | `` | Specify the console output format. (default text) Options are: json, ndjson, ndjson-wrapped, text. |
| `-P, --profile string` | `` | The name of a configuration profile to use. |
| `--debug` | `` | Enable debug output for error messages, including stack traces and transaction IDs. (default false) |
| `--log-file string` | `` | Write logs to a file at the given path. File logging is disabled when not set. |
| `--log-file-level string` | `` | Set the file log level. Options are: DEBUG, INFO, WARN, ERROR. (default DEBUG) |
| `--log-level string` | `` | Set the console log level. Options are: DEBUG, INFO, WARN, ERROR. (default WARN) |
| `--no-color` | `` | Disable text output in color. (default false) |
| `--query string` | `` | JMESPath expression to filter JSON output. Requires -O json, ndjson, or ndjson-wrapped. Example: --query 'data[?enabled].name' |


## Parent Command

- [`pingcli pingone verify verify-policies`](cmd-pingcli-pingone-verify-verify-policies.md) — Verify Policies
