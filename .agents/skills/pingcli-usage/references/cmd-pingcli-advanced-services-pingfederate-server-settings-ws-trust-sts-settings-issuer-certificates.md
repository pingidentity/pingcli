# `pingcli advanced-services pingfederate server-settings ws-trust-sts-settings issuer-certificates`
PingFederate WS-Trust STS Issuer Certificates

## Synopsis

PingFederate WS-Trust STS Issuer Certificates

```
pingcli advanced-services pingfederate server-settings ws-trust-sts-settings issuer-certificates [flags]
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
| `pingcli advanced-services pingfederate server-settings ws-trust-sts-settings issuer-certificates create` | Import a new WS-Trust STS issuer certificate | [`cmd-pingcli-advanced-services-pingfederate-server-settings-ws-trust-sts-settings-issuer-certificates-create.md`](cmd-pingcli-advanced-services-pingfederate-server-settings-ws-trust-sts-settings-issuer-certificates-create.md) |
| `pingcli advanced-services pingfederate server-settings ws-trust-sts-settings issuer-certificates delete` | Delete a WS-Trust STS issuer certificate | [`cmd-pingcli-advanced-services-pingfederate-server-settings-ws-trust-sts-settings-issuer-certificates-delete.md`](cmd-pingcli-advanced-services-pingfederate-server-settings-ws-trust-sts-settings-issuer-certificates-delete.md) |
| `pingcli advanced-services pingfederate server-settings ws-trust-sts-settings issuer-certificates get` | Read a specific WS-Trust STS issuer certificate | [`cmd-pingcli-advanced-services-pingfederate-server-settings-ws-trust-sts-settings-issuer-certificates-get.md`](cmd-pingcli-advanced-services-pingfederate-server-settings-ws-trust-sts-settings-issuer-certificates-get.md) |
| `pingcli advanced-services pingfederate server-settings ws-trust-sts-settings issuer-certificates list` | List all WS-Trust STS issuer certificates | [`cmd-pingcli-advanced-services-pingfederate-server-settings-ws-trust-sts-settings-issuer-certificates-list.md`](cmd-pingcli-advanced-services-pingfederate-server-settings-ws-trust-sts-settings-issuer-certificates-list.md) |
| `pingcli advanced-services pingfederate server-settings ws-trust-sts-settings issuer-certificates template` | Generate a WS-Trust STS issuer certificate JSON template | [`cmd-pingcli-advanced-services-pingfederate-server-settings-ws-trust-sts-settings-issuer-certificates-template.md`](cmd-pingcli-advanced-services-pingfederate-server-settings-ws-trust-sts-settings-issuer-certificates-template.md) |

## Parent Command

- [`pingcli advanced-services pingfederate server-settings ws-trust-sts-settings`](cmd-pingcli-advanced-services-pingfederate-server-settings-ws-trust-sts-settings.md) — PingFederate WS-Trust STS Settings
