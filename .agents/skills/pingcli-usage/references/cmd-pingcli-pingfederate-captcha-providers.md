# `pingcli pingfederate captcha-providers`
PingFederate CAPTCHA Providers

## Synopsis

PingFederate CAPTCHA Providers

```
pingcli pingfederate captcha-providers [flags]
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
| `pingcli pingfederate captcha-providers apply` | Create or update a CAPTCHA provider | [`cmd-pingcli-pingfederate-captcha-providers-apply.md`](cmd-pingcli-pingfederate-captcha-providers-apply.md) |
| `pingcli pingfederate captcha-providers create` | Create a new CAPTCHA provider | [`cmd-pingcli-pingfederate-captcha-providers-create.md`](cmd-pingcli-pingfederate-captcha-providers-create.md) |
| `pingcli pingfederate captcha-providers delete` | Delete a CAPTCHA provider | [`cmd-pingcli-pingfederate-captcha-providers-delete.md`](cmd-pingcli-pingfederate-captcha-providers-delete.md) |
| `pingcli pingfederate captcha-providers descriptors` | PingFederate CAPTCHA Provider Descriptors | [`cmd-pingcli-pingfederate-captcha-providers-descriptors.md`](cmd-pingcli-pingfederate-captcha-providers-descriptors.md) |
| `pingcli pingfederate captcha-providers get` | Read a specific CAPTCHA provider | [`cmd-pingcli-pingfederate-captcha-providers-get.md`](cmd-pingcli-pingfederate-captcha-providers-get.md) |
| `pingcli pingfederate captcha-providers list` | List all CAPTCHA providers | [`cmd-pingcli-pingfederate-captcha-providers-list.md`](cmd-pingcli-pingfederate-captcha-providers-list.md) |
| `pingcli pingfederate captcha-providers replace` | Update a CAPTCHA provider | [`cmd-pingcli-pingfederate-captcha-providers-replace.md`](cmd-pingcli-pingfederate-captcha-providers-replace.md) |
| `pingcli pingfederate captcha-providers settings` | PingFederate CAPTCHA Providers Settings | [`cmd-pingcli-pingfederate-captcha-providers-settings.md`](cmd-pingcli-pingfederate-captcha-providers-settings.md) |
| `pingcli pingfederate captcha-providers template` | Generate a CAPTCHA provider JSON template | [`cmd-pingcli-pingfederate-captcha-providers-template.md`](cmd-pingcli-pingfederate-captcha-providers-template.md) |

## Parent Command

- [`pingcli pingfederate`](cmd-pingcli-pingfederate.md) — Administration tools for PingFederate deployed as software
