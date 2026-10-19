# `pingcli pingone environments update-type`
Update environment type

## Synopsis

Update a PingOne environment's type. PRODUCTION and SANDBOX are the only values, and the change is one-way (no SANDBOX downgrade).

```
pingcli pingone environments update-type [flags]
```

## Examples

```
# Promote a sandbox environment to production
  pingcli pingone environments update-type --environment-id <environment-uuid> --type PRODUCTION
```

## Options

| Flag | Default | Description |
|------|---------|-------------|
| `-h, --help` | `` | help for update-type |
| `-e, --environment-id string` | `` | The PingOne environment ID |
| `--type string` | `` | The new environment type. Options are: PRODUCTION, SANDBOX. The change is one-way: a PRODUCTION environment cannot be downgraded to SANDBOX. |


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

- [`pingcli pingone environments`](cmd-pingcli-pingone-environments.md) — Environments
