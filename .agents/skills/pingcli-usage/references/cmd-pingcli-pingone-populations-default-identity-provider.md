# `pingcli pingone populations default-identity-provider`
Population Default Identity Provider

## Synopsis

Population Default Identity Provider

```
pingcli pingone populations default-identity-provider [flags]
```

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


## Subcommands

| Command | Description | Reference |
|---------|-------------|----------|
| `pingcli pingone populations default-identity-provider apply` | Update population default identity provider | [`cmd-pingcli-pingone-populations-default-identity-provider-apply.md`](cmd-pingcli-pingone-populations-default-identity-provider-apply.md) |
| `pingcli pingone populations default-identity-provider get` | Read population default identity provider | [`cmd-pingcli-pingone-populations-default-identity-provider-get.md`](cmd-pingcli-pingone-populations-default-identity-provider-get.md) |
| `pingcli pingone populations default-identity-provider replace` | Update population default identity provider | [`cmd-pingcli-pingone-populations-default-identity-provider-replace.md`](cmd-pingcli-pingone-populations-default-identity-provider-replace.md) |
| `pingcli pingone populations default-identity-provider template` | Generate a population default identity provider JSON template | [`cmd-pingcli-pingone-populations-default-identity-provider-template.md`](cmd-pingcli-pingone-populations-default-identity-provider-template.md) |

## Parent Command

- [`pingcli pingone populations`](cmd-pingcli-pingone-populations.md) — Populations
