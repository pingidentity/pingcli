# `pingcli pingfederate authentication-policies default replace`
Update default authentication policy

## Synopsis

Update (replace) the PingFederate Default Authentication Policy

```
pingcli pingfederate authentication-policies default replace [flags]
```

## Examples

```
# Update the default authentication policy from a JSON file
  pingcli pingfederate authentication-policies default replace --from-file authentication-policies-default.json

  # Update the default authentication policy from a JSON file, overriding selected fields with flags
  pingcli pingfederate authentication-policies default replace --from-file authentication-policies-default.json --fail-if-no-selection --tracked-http-parameters paramA --tracked-http-parameters paramB

  # Update the default authentication policy from stdin
  pingcli pingfederate authentication-policies default replace --from-file - < authentication-policies-default.json
```

## Options

| Flag | Default | Description |
|------|---------|-------------|
| `-h, --help` | `` | help for replace |
| `-f, --from-file string` | `` | Path to a JSON file containing the request body, or "-" to read from stdin. |
| `--fail-if-no-selection` | `` | Fail if policy finds no authentication source. |
| `--tracked-http-parameters []string` | `` | HTTP parameters tracked by authentication policies; repeatable or comma-separated |


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

- [`pingcli pingfederate authentication-policies default`](cmd-pingcli-pingfederate-authentication-policies-default.md) — PingFederate Default Authentication Policy
