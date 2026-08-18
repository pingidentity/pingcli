# `pingcli pingfederate key-pairs oauth-openid-connect additional-key-sets`
PingFederate OAuth/OpenID Connect additional key sets

## Synopsis

PingFederate OAuth/OpenID Connect additional key sets

```
pingcli pingfederate key-pairs oauth-openid-connect additional-key-sets [flags]
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
| `pingcli pingfederate key-pairs oauth-openid-connect additional-key-sets apply` | Create or update an OAuth/OpenID Connect additional key set | [`cmd-pingcli-pingfederate-key-pairs-oauth-openid-connect-additional-key-sets-apply.md`](cmd-pingcli-pingfederate-key-pairs-oauth-openid-connect-additional-key-sets-apply.md) |
| `pingcli pingfederate key-pairs oauth-openid-connect additional-key-sets create` | Create a new OAuth/OpenID Connect additional key set | [`cmd-pingcli-pingfederate-key-pairs-oauth-openid-connect-additional-key-sets-create.md`](cmd-pingcli-pingfederate-key-pairs-oauth-openid-connect-additional-key-sets-create.md) |
| `pingcli pingfederate key-pairs oauth-openid-connect additional-key-sets delete` | Delete an OAuth/OpenID Connect additional key set | [`cmd-pingcli-pingfederate-key-pairs-oauth-openid-connect-additional-key-sets-delete.md`](cmd-pingcli-pingfederate-key-pairs-oauth-openid-connect-additional-key-sets-delete.md) |
| `pingcli pingfederate key-pairs oauth-openid-connect additional-key-sets get` | Read a specific OAuth/OpenID Connect additional key set | [`cmd-pingcli-pingfederate-key-pairs-oauth-openid-connect-additional-key-sets-get.md`](cmd-pingcli-pingfederate-key-pairs-oauth-openid-connect-additional-key-sets-get.md) |
| `pingcli pingfederate key-pairs oauth-openid-connect additional-key-sets list` | List all OAuth/OpenID Connect additional key sets | [`cmd-pingcli-pingfederate-key-pairs-oauth-openid-connect-additional-key-sets-list.md`](cmd-pingcli-pingfederate-key-pairs-oauth-openid-connect-additional-key-sets-list.md) |
| `pingcli pingfederate key-pairs oauth-openid-connect additional-key-sets replace` | Update an OAuth/OpenID Connect additional key set | [`cmd-pingcli-pingfederate-key-pairs-oauth-openid-connect-additional-key-sets-replace.md`](cmd-pingcli-pingfederate-key-pairs-oauth-openid-connect-additional-key-sets-replace.md) |
| `pingcli pingfederate key-pairs oauth-openid-connect additional-key-sets template` | Generate an OAuth/OpenID Connect additional key set JSON template | [`cmd-pingcli-pingfederate-key-pairs-oauth-openid-connect-additional-key-sets-template.md`](cmd-pingcli-pingfederate-key-pairs-oauth-openid-connect-additional-key-sets-template.md) |

## Parent Command

- [`pingcli pingfederate key-pairs oauth-openid-connect`](cmd-pingcli-pingfederate-key-pairs-oauth-openid-connect.md) — PingFederate OAuth/OpenID Connect keys settings
