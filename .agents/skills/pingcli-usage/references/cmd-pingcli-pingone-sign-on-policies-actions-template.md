# `pingcli pingone sign-on-policies actions template`
Generate a sign-on policy action JSON template

## Synopsis

Generate a JSON skeleton template for sign-on policy action create or replace bodies.

Select the desired type with --template-type (one of: login,
multiFactorAuthentication, identifierFirst, identityProvider, agreement,
progressiveProfiling, pingIdWinLoginPasswordless, pingIdAuthentication).
Defaults to login when omitted.

```
pingcli pingone sign-on-policies actions template [flags]
```

## Examples

```
# Generate a login action template and save to a file
  pingcli pingone sign-on-policies actions template --template-type login > body.json

  # Edit the template — fill in the required fields, then pass the body to
  # the create or replace command
Use the JSON template as a starting point:
  1. Run the template command to generate the body skeleton.
  2. Edit the file, replacing placeholder values with real data.
  3. Feed the edited file back to the create or replace action via --from-file.

Example workflow:
  pingcli pingone sign-on-policies actions template > body.json
  # edit body.json
  pingcli pingone sign-on-policies actions create --from-file body.json
```

## Options

| Flag | Default | Description |
|------|---------|-------------|
| `-h, --help` | `` | help for template |
| `-o, --output-file string` | `` | Write the JSON template to PATH instead of stdout. Overwrites any existing file. |
| `--template-type string` | `` | The variant of the template to generate. One of: login, multiFactorAuthentication, identifierFirst, identityProvider, agreement, progressiveProfiling, pingIdWinLoginPasswordless, pingIdAuthentication. Defaults to "login". |


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

- [`pingcli pingone sign-on-policies actions`](cmd-pingcli-pingone-sign-on-policies-actions.md) — Sign-On Policy Actions
