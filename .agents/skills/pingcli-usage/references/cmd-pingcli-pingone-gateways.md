# `pingcli pingone gateways`
Gateways

## Synopsis

Gateways

```
pingcli pingone gateways [flags]
```

## Inherited Options

| Flag | Default | Description |
|------|---------|-------------|
| `-C, --config string` | `` | The relative or full path to a custom Ping CLI configuration file. (default $HOME/.pingcli/config.yaml) |
| `-D, --detailed-exitcode` | `` | Enable detailed exit code output. (default false) 0 - pingcli command succeeded with no errors or warnings. 1 - pingcli command failed with errors. 2 - pingcli command succeeded with warnings. |
| `-O, --output-format string` | `` | Specify the console output format. (default text) Options are: json, ndjson, ndjson-wrapped, text. |
| `-P, --profile string` | `` | The name of a configuration profile to use. |
| `--debug` | `` | Enable debug output for error messages, including stack traces and transaction IDs. (default false) |
| `--log-file string` | `` | Write logs to a file at the given path. File logging is disabled when not set. |
| `--log-file-level string` | `` | Set the file log level. Options are: DEBUG, INFO, WARN, ERROR. (default DEBUG) |
| `--log-level string` | `` | Set the console log level. Options are: DEBUG, INFO, WARN, ERROR. (default WARN) |
| `--no-color` | `` | Disable text output in color. (default false) |
| `--query string` | `` | JMESPath expression to filter JSON output. Requires -O json, ndjson, or ndjson-wrapped. Example: --query 'data[?enabled].name' |


## Subcommands

| Command | Description | Reference |
|---------|-------------|----------|
| `pingcli pingone gateways apply` | Create or update a gateway | [`cmd-pingcli-pingone-gateways-apply.md`](cmd-pingcli-pingone-gateways-apply.md) |
| `pingcli pingone gateways create` | Create a new gateway | [`cmd-pingcli-pingone-gateways-create.md`](cmd-pingcli-pingone-gateways-create.md) |
| `pingcli pingone gateways delete` | Delete a gateway | [`cmd-pingcli-pingone-gateways-delete.md`](cmd-pingcli-pingone-gateways-delete.md) |
| `pingcli pingone gateways gateway-credentials` | Gateway credentials | [`cmd-pingcli-pingone-gateways-gateway-credentials.md`](cmd-pingcli-pingone-gateways-gateway-credentials.md) |
| `pingcli pingone gateways gateway-instances` | Gateway instances | [`cmd-pingcli-pingone-gateways-gateway-instances.md`](cmd-pingcli-pingone-gateways-gateway-instances.md) |
| `pingcli pingone gateways get` | Read a specific gateway | [`cmd-pingcli-pingone-gateways-get.md`](cmd-pingcli-pingone-gateways-get.md) |
| `pingcli pingone gateways list` | List all gateways | [`cmd-pingcli-pingone-gateways-list.md`](cmd-pingcli-pingone-gateways-list.md) |
| `pingcli pingone gateways replace` | Replace a gateway | [`cmd-pingcli-pingone-gateways-replace.md`](cmd-pingcli-pingone-gateways-replace.md) |
| `pingcli pingone gateways template` | Generate a gateway JSON template | [`cmd-pingcli-pingone-gateways-template.md`](cmd-pingcli-pingone-gateways-template.md) |

## Parent Command

- [`pingcli pingone`](cmd-pingcli-pingone.md) — Administration tools for the PingOne platform.
