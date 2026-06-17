# `pingcli pingone identity-providers`
Identity Providers

## Synopsis

Identity Providers

```
pingcli pingone identity-providers [flags]
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
| `pingcli pingone identity-providers apply` | Create or update an identity provider | [`cmd-pingcli-pingone-identity-providers-apply.md`](cmd-pingcli-pingone-identity-providers-apply.md) |
| `pingcli pingone identity-providers create` | Create a new identity provider | [`cmd-pingcli-pingone-identity-providers-create.md`](cmd-pingcli-pingone-identity-providers-create.md) |
| `pingcli pingone identity-providers delete` | Delete an identity provider | [`cmd-pingcli-pingone-identity-providers-delete.md`](cmd-pingcli-pingone-identity-providers-delete.md) |
| `pingcli pingone identity-providers get` | Read a specific identity provider | [`cmd-pingcli-pingone-identity-providers-get.md`](cmd-pingcli-pingone-identity-providers-get.md) |
| `pingcli pingone identity-providers idp-attributes` | Identity Provider Attributes | [`cmd-pingcli-pingone-identity-providers-idp-attributes.md`](cmd-pingcli-pingone-identity-providers-idp-attributes.md) |
| `pingcli pingone identity-providers list` | List all identity providers | [`cmd-pingcli-pingone-identity-providers-list.md`](cmd-pingcli-pingone-identity-providers-list.md) |
| `pingcli pingone identity-providers replace` | Replace an identity provider | [`cmd-pingcli-pingone-identity-providers-replace.md`](cmd-pingcli-pingone-identity-providers-replace.md) |
| `pingcli pingone identity-providers template` | Generate an identity provider JSON template | [`cmd-pingcli-pingone-identity-providers-template.md`](cmd-pingcli-pingone-identity-providers-template.md) |

## Parent Command

- [`pingcli pingone`](cmd-pingcli-pingone.md) — Administration tools for the PingOne platform.
