# `pingcli pingfederate key-pairs oauth-openid-connect`
PingFederate OAuth/OpenID Connect keys settings

## Synopsis

PingFederate OAuth/OpenID Connect keys settings

```
pingcli pingfederate key-pairs oauth-openid-connect [flags]
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
| `pingcli pingfederate key-pairs oauth-openid-connect additional-key-sets` | PingFederate OAuth/OpenID Connect additional key sets | [`cmd-pingcli-pingfederate-key-pairs-oauth-openid-connect-additional-key-sets.md`](cmd-pingcli-pingfederate-key-pairs-oauth-openid-connect-additional-key-sets.md) |
| `pingcli pingfederate key-pairs oauth-openid-connect apply` | Update OAuth/OpenID Connect keys settings | [`cmd-pingcli-pingfederate-key-pairs-oauth-openid-connect-apply.md`](cmd-pingcli-pingfederate-key-pairs-oauth-openid-connect-apply.md) |
| `pingcli pingfederate key-pairs oauth-openid-connect get` | Read OAuth/OpenID Connect keys settings | [`cmd-pingcli-pingfederate-key-pairs-oauth-openid-connect-get.md`](cmd-pingcli-pingfederate-key-pairs-oauth-openid-connect-get.md) |
| `pingcli pingfederate key-pairs oauth-openid-connect replace` | Update OAuth/OpenID Connect keys settings | [`cmd-pingcli-pingfederate-key-pairs-oauth-openid-connect-replace.md`](cmd-pingcli-pingfederate-key-pairs-oauth-openid-connect-replace.md) |
| `pingcli pingfederate key-pairs oauth-openid-connect template` | Generate an OAuth/OpenID Connect keys settings JSON template | [`cmd-pingcli-pingfederate-key-pairs-oauth-openid-connect-template.md`](cmd-pingcli-pingfederate-key-pairs-oauth-openid-connect-template.md) |

## Parent Command

- [`pingcli pingfederate key-pairs`](cmd-pingcli-pingfederate-key-pairs.md) — Manage PingFederate Key Pairs resources
