# `pingcli pingone custom-domains`
Custom Domains

## Synopsis

Custom Domains

```
pingcli pingone custom-domains [flags]
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
| `pingcli pingone custom-domains create` | Create a new custom domain | [`cmd-pingcli-pingone-custom-domains-create.md`](cmd-pingcli-pingone-custom-domains-create.md) |
| `pingcli pingone custom-domains delete` | Delete a custom domain | [`cmd-pingcli-pingone-custom-domains-delete.md`](cmd-pingcli-pingone-custom-domains-delete.md) |
| `pingcli pingone custom-domains get` | Read a specific custom domain | [`cmd-pingcli-pingone-custom-domains-get.md`](cmd-pingcli-pingone-custom-domains-get.md) |
| `pingcli pingone custom-domains import-certificate` | Import an SSL certificate for a custom domain | [`cmd-pingcli-pingone-custom-domains-import-certificate.md`](cmd-pingcli-pingone-custom-domains-import-certificate.md) |
| `pingcli pingone custom-domains list` | List all custom domains | [`cmd-pingcli-pingone-custom-domains-list.md`](cmd-pingcli-pingone-custom-domains-list.md) |
| `pingcli pingone custom-domains template` | Generate a custom domain JSON template | [`cmd-pingcli-pingone-custom-domains-template.md`](cmd-pingcli-pingone-custom-domains-template.md) |
| `pingcli pingone custom-domains verify` | Verify a custom domain | [`cmd-pingcli-pingone-custom-domains-verify.md`](cmd-pingcli-pingone-custom-domains-verify.md) |

## Parent Command

- [`pingcli pingone`](cmd-pingcli-pingone.md) — Administration tools for the PingOne platform.
