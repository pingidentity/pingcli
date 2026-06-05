# `pingcli config set`
Set stored configuration settings for the CLI.

## Synopsis

Set stored configuration settings for the CLI.

The `--profile` parameter can be used to set configuration settings for a specified custom configuration profile.
Where `--profile` is not specified, configuration settings will be set for the currently active profile.

```
pingcli config set [flags] key=value
```

## Examples

```
Set the color setting to true for the currently active profile.
    pingcli config set color=true

  Set the PingOne endpoint root domain setting to 'pingone.eu' for the profile named 'myprofile'.
    pingcli config set --profile myprofile service.pingOne.endpoint.rootDomain=pingone.eu

  Set the PingFederate basic authentication password with unmasked output
    pingcli config set --profile myprofile --unmask-values service.pingFederate.authentication.basicAuth.password=1234
```

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
