# `pingcli config list-keys`
List all configuration keys.

## Synopsis

View the complete list of available configuration options. These attributes can be saved via the set and unset config subcommands or stored in a profile within the config file.
For details on the configuration options visit: https://github.com/pingidentity/pingcli/blob/main/docs/tool-configuration/configuration-key.md

```
pingcli config list-keys [flags]
```

## Examples

```
List all configuration keys stored in the CLI configuration file.
  pingcli config list-keys
	pingcli config list-keys --yaml
```

## Options

| Flag | Default | Description |
|------|---------|-------------|
| `-h, --help` | `` | help for list-keys |
| `-y, --yaml` | `` | Output configuration keys in YAML format. (default false) |
| `--template string` | `` | A Go text/template string. When provided, the command output is rendered through the template instead of the default format. The template receives the command's structured response data. Example: --template '{{.Name}}' |


## Inherited Options

| Flag | Default | Description |
|------|---------|-------------|
| `-C, --config string` | `` | The relative or full path to a custom Ping CLI configuration file. (default $HOME/.pingcli/config.yaml) |
| `-D, --detailed-exitcode` | `` | Enable detailed exit code output. (default false) 0 - pingcli command succeeded with no errors or warnings. 1 - pingcli command failed with errors. 2 - pingcli command succeeded with warnings. |
| `-O, --output-format string` | `` | Specify the console output format. (default text) Options are: json, ndjson, ndjson-typed, ndjson-wrapped, text. |
| `-P, --profile string` | `` | The name of a configuration profile to use. |
| `-U, --unmask-values` | `` | Unmask secret values. (default false) |
| `--debug` | `` | Enable debug output for error messages, including stack traces and transaction IDs. (default false) |
| `--log-file string` | `` | Write logs to a file at the given path. File logging is disabled when not set. |
| `--log-file-level string` | `` | Set the file log level. Options are: DEBUG, INFO, WARN, ERROR. (default DEBUG) |
| `--log-level string` | `` | Set the console log level. Options are: DEBUG, INFO, WARN, ERROR. (default WARN) |
| `--no-color` | `` | Disable text output in color. (default false) |
| `--query string` | `` | JMESPath expression to filter JSON output. Requires -O json, ndjson, ndjson-typed, or ndjson-wrapped. Example: --query 'data[?enabled].name' |


## Parent Command

- [`pingcli config`](cmd-pingcli-config.md) — Manage the CLI configuration.
