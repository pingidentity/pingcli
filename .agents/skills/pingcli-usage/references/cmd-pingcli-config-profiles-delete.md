# `pingcli config profiles delete`
Delete a custom configuration profile.

## Synopsis

Delete an existing custom configuration profile from the CLI.
		
The profile to delete will be removed from the CLI configuration file.

```
pingcli config profiles delete [flags] [profile-name]
```

## Examples

```
Delete a configuration profile by selecting from the available profiles.
    pingcli config profiles delete

  Delete a configuration profile by specifying the name of an existing configured profile.
    pingcli config profiles delete myprofile
	
  Delete a configuration profile by auto-accepting the deletion.
    pingcli config profiles delete --yes myprofile
```

## Options

| Flag | Default | Description |
|------|---------|-------------|
| `-h, --help` | `` | help for delete |
| `-y, --yes` | `` | Auto-accept the profile deletion confirmation prompt. (default false) |


## Inherited Options

| Flag | Default | Description |
|------|---------|-------------|
| `-C, --config string` | `` | The relative or full path to a custom Ping CLI configuration file. (default $HOME/.pingcli/config.yaml) |
| `-D, --detailed-exitcode` | `` | Enable detailed exit code output. (default false) 0 - pingcli command succeeded with no errors or warnings. 1 - pingcli command failed with errors. 2 - pingcli command succeeded with warnings. |
| `-O, --output-format string` | `` | Specify the console output format. (default text) Options are: json, ndjson, ndjson-wrapped, text. |
| `-P, --profile string` | `` | The name of a configuration profile to use. |
| `-U, --unmask-values` | `` | Unmask secret values. (default false) |
| `--debug` | `` | Enable debug output for error messages, including stack traces and transaction IDs. (default false) |
| `--log-file string` | `` | Write logs to a file at the given path. File logging is disabled when not set. |
| `--log-file-level string` | `` | Set the file log level. Options are: DEBUG, INFO, WARN, ERROR. (default DEBUG) |
| `--log-level string` | `` | Set the console log level. Options are: DEBUG, INFO, WARN, ERROR. (default WARN) |
| `--no-color` | `` | Disable text output in color. (default false) |
| `--query string` | `` | JMESPath expression to filter JSON output. Requires -O json, ndjson, or ndjson-wrapped. Example: --query 'data[?enabled].name' |


## Parent Command

- [`pingcli config profiles`](cmd-pingcli-config-profiles.md) — Manage the configuration profiles.
