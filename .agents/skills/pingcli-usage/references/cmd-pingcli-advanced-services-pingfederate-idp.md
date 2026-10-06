# `pingcli advanced-services pingfederate idp`
Manage PingFederate IdP resources

## Synopsis

Manage PingFederate Identity Provider (IdP) resources

```
pingcli advanced-services pingfederate idp
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
| `pingcli advanced-services pingfederate idp adapters` | PingFederate IDP Adapters | [`cmd-pingcli-advanced-services-pingfederate-idp-adapters.md`](cmd-pingcli-advanced-services-pingfederate-idp-adapters.md) |
| `pingcli advanced-services pingfederate idp connectors` | Manage PingFederate IdP Connector resources | [`cmd-pingcli-advanced-services-pingfederate-idp-connectors.md`](cmd-pingcli-advanced-services-pingfederate-idp-connectors.md) |
| `pingcli advanced-services pingfederate idp default-urls` | PingFederate IdP Default URLs | [`cmd-pingcli-advanced-services-pingfederate-idp-default-urls.md`](cmd-pingcli-advanced-services-pingfederate-idp-default-urls.md) |
| `pingcli advanced-services pingfederate idp sp-connections` | PingFederate SP Connections | [`cmd-pingcli-advanced-services-pingfederate-idp-sp-connections.md`](cmd-pingcli-advanced-services-pingfederate-idp-sp-connections.md) |
| `pingcli advanced-services pingfederate idp sts-request-parameters-contracts` | PingFederate STS request parameters contracts | [`cmd-pingcli-advanced-services-pingfederate-idp-sts-request-parameters-contracts.md`](cmd-pingcli-advanced-services-pingfederate-idp-sts-request-parameters-contracts.md) |
| `pingcli advanced-services pingfederate idp token-processors` | PingFederate Token Processors | [`cmd-pingcli-advanced-services-pingfederate-idp-token-processors.md`](cmd-pingcli-advanced-services-pingfederate-idp-token-processors.md) |

## Parent Command

- [`pingcli advanced-services pingfederate`](cmd-pingcli-advanced-services-pingfederate.md) — Administration tools for PingFederate deployed as a private tenant service in PingOne Advanced Services
