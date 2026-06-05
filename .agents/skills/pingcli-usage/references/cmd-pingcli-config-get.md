# `pingcli config get`
Read stored configuration settings for the CLI.

## Synopsis

Read stored configuration settings for the CLI.

The `--profile` parameter can be used to read configuration settings for a specified custom configuration profile.
Where `--profile` is not specified, configuration settings will be read for the currently active profile.

```
pingcli config get [flags] key
```

## Examples

```
Read all the configuration settings for the PingOne service in the active (or default) profile.
    pingcli config get pingone

  Read the noColor setting for the profile named 'myprofile'.
    pingcli config get --profile myprofile noColor

  Read the worker ID used to authenticate to the PingOne service management API.
    pingcli config get service.pingOne.authentication.worker.environmentID
	
  Read the unmasked value for the request access token.
    pingcli config get --unmask-values request.accessToken
```

## Options

| Flag | Default | Description |
|------|---------|-------------|
| `-h, --help` | `` | help for get |
| `--template string` | `` | A Go text/template string. When provided, the command output is rendered through the template instead of the default format. The template receives the command's structured response data. Example: --template '{{.Name}}' |


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

- [`pingcli config`](cmd-pingcli-config.md) — Manage the CLI configuration.
