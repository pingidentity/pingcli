# `pingcli pingone email-domains trusted-email-addresses`
Trusted Email Addresses

## Synopsis

Trusted Email Addresses

```
pingcli pingone email-domains trusted-email-addresses [flags]
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
| `pingcli pingone email-domains trusted-email-addresses create` | Create a new trusted email address | [`cmd-pingcli-pingone-email-domains-trusted-email-addresses-create.md`](cmd-pingcli-pingone-email-domains-trusted-email-addresses-create.md) |
| `pingcli pingone email-domains trusted-email-addresses delete` | Delete a trusted email address | [`cmd-pingcli-pingone-email-domains-trusted-email-addresses-delete.md`](cmd-pingcli-pingone-email-domains-trusted-email-addresses-delete.md) |
| `pingcli pingone email-domains trusted-email-addresses get` | Read a specific trusted email address | [`cmd-pingcli-pingone-email-domains-trusted-email-addresses-get.md`](cmd-pingcli-pingone-email-domains-trusted-email-addresses-get.md) |
| `pingcli pingone email-domains trusted-email-addresses list` | List all trusted email addresses | [`cmd-pingcli-pingone-email-domains-trusted-email-addresses-list.md`](cmd-pingcli-pingone-email-domains-trusted-email-addresses-list.md) |
| `pingcli pingone email-domains trusted-email-addresses resend-verification-code` | Resend the verification code to a trusted email address | [`cmd-pingcli-pingone-email-domains-trusted-email-addresses-resend-verification-code.md`](cmd-pingcli-pingone-email-domains-trusted-email-addresses-resend-verification-code.md) |
| `pingcli pingone email-domains trusted-email-addresses template` | Generate a trusted email address JSON template | [`cmd-pingcli-pingone-email-domains-trusted-email-addresses-template.md`](cmd-pingcli-pingone-email-domains-trusted-email-addresses-template.md) |

## Parent Command

- [`pingcli pingone email-domains`](cmd-pingcli-pingone-email-domains.md) — Email Domains
