# `pingcli pingfederate pingone-for-enterprise`
Manage the PingFederate connection to PingOne for Enterprise

## Synopsis

Manage the PingFederate connection to PingOne for Enterprise

```
pingcli pingfederate pingone-for-enterprise [flags]
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
| `pingcli pingfederate pingone-for-enterprise apply` | Update PingOne for Enterprise settings | [`cmd-pingcli-pingfederate-pingone-for-enterprise-apply.md`](cmd-pingcli-pingfederate-pingone-for-enterprise-apply.md) |
| `pingcli pingfederate pingone-for-enterprise disconnect` | Disconnect PingFederate from PingOne for Enterprise | [`cmd-pingcli-pingfederate-pingone-for-enterprise-disconnect.md`](cmd-pingcli-pingfederate-pingone-for-enterprise-disconnect.md) |
| `pingcli pingfederate pingone-for-enterprise get` | Read PingOne for Enterprise settings | [`cmd-pingcli-pingfederate-pingone-for-enterprise-get.md`](cmd-pingcli-pingfederate-pingone-for-enterprise-get.md) |
| `pingcli pingfederate pingone-for-enterprise key-pairs` | Manage PingOne for Enterprise key pairs | [`cmd-pingcli-pingfederate-pingone-for-enterprise-key-pairs.md`](cmd-pingcli-pingfederate-pingone-for-enterprise-key-pairs.md) |
| `pingcli pingfederate pingone-for-enterprise replace` | Update PingOne for Enterprise settings | [`cmd-pingcli-pingfederate-pingone-for-enterprise-replace.md`](cmd-pingcli-pingfederate-pingone-for-enterprise-replace.md) |
| `pingcli pingfederate pingone-for-enterprise template` | Generate a PingOne for Enterprise settings JSON template | [`cmd-pingcli-pingfederate-pingone-for-enterprise-template.md`](cmd-pingcli-pingfederate-pingone-for-enterprise-template.md) |
| `pingcli pingfederate pingone-for-enterprise update-identity-repository` | Update the PingOne for Enterprise identity repository | [`cmd-pingcli-pingfederate-pingone-for-enterprise-update-identity-repository.md`](cmd-pingcli-pingfederate-pingone-for-enterprise-update-identity-repository.md) |

## Parent Command

- [`pingcli pingfederate`](cmd-pingcli-pingfederate.md) — Administration tools for PingFederate deployed as software
