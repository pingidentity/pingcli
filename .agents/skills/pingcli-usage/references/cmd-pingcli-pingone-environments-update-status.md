# `pingcli pingone environments update-status`
Update environment status

## Synopsis

Update a PingOne environment's status. Applies to PRODUCTION environments only. --status DELETE_PENDING soft-deletes the environment; --status ACTIVE (re)activates it and requires --license-id.

```
pingcli pingone environments update-status [flags]
```

## Examples

```
# Soft-delete an environment (status becomes DELETE_PENDING)
  pingcli pingone environments update-status --environment-id <environment-uuid> --status DELETE_PENDING

  # Restore a soft-deleted environment and assign a license
  pingcli pingone environments update-status --environment-id <environment-uuid> --status ACTIVE --license-id <license-uuid>
```

## Options

| Flag | Default | Description |
|------|---------|-------------|
| `-h, --help` | `` | help for update-status |
| `-e, --environment-id string` | `` | The PingOne environment ID |
| `--license-id string` | `` | The ID of the PingOne license to assign when --status is ACTIVE. Required for ACTIVE; ignored for DELETE_PENDING. Discover licenses with the pingcli pingone licenses list command. |
| `--status string` | `` | The new environment status. Options are: ACTIVE, DELETE_PENDING. Applies to PRODUCTION environments only. --status ACTIVE requires --license-id. |


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
