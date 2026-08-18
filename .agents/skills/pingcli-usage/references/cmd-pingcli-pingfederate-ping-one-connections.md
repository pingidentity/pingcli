# `pingcli pingfederate ping-one-connections`
PingFederate PingOne connections

## Synopsis

PingFederate PingOne Connections

```
pingcli pingfederate ping-one-connections [flags]
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
| `pingcli pingfederate ping-one-connections apply` | Create or update a PingOne connection | [`cmd-pingcli-pingfederate-ping-one-connections-apply.md`](cmd-pingcli-pingfederate-ping-one-connections-apply.md) |
| `pingcli pingfederate ping-one-connections create` | Create a new PingOne connection | [`cmd-pingcli-pingfederate-ping-one-connections-create.md`](cmd-pingcli-pingfederate-ping-one-connections-create.md) |
| `pingcli pingfederate ping-one-connections credential-status` | Read a PingOne connection credential status | [`cmd-pingcli-pingfederate-ping-one-connections-credential-status.md`](cmd-pingcli-pingfederate-ping-one-connections-credential-status.md) |
| `pingcli pingfederate ping-one-connections delete` | Delete a PingOne connection | [`cmd-pingcli-pingfederate-ping-one-connections-delete.md`](cmd-pingcli-pingfederate-ping-one-connections-delete.md) |
| `pingcli pingfederate ping-one-connections environments` | List the PingOne environments for a PingOne connection | [`cmd-pingcli-pingfederate-ping-one-connections-environments.md`](cmd-pingcli-pingfederate-ping-one-connections-environments.md) |
| `pingcli pingfederate ping-one-connections get` | Read a specific PingOne connection | [`cmd-pingcli-pingfederate-ping-one-connections-get.md`](cmd-pingcli-pingfederate-ping-one-connections-get.md) |
| `pingcli pingfederate ping-one-connections list` | List all PingOne connections | [`cmd-pingcli-pingfederate-ping-one-connections-list.md`](cmd-pingcli-pingfederate-ping-one-connections-list.md) |
| `pingcli pingfederate ping-one-connections replace` | Update a PingOne connection | [`cmd-pingcli-pingfederate-ping-one-connections-replace.md`](cmd-pingcli-pingfederate-ping-one-connections-replace.md) |
| `pingcli pingfederate ping-one-connections service-associations` | Read a PingOne connection's service associations | [`cmd-pingcli-pingfederate-ping-one-connections-service-associations.md`](cmd-pingcli-pingfederate-ping-one-connections-service-associations.md) |
| `pingcli pingfederate ping-one-connections template` | Generate a PingOne connection JSON template | [`cmd-pingcli-pingfederate-ping-one-connections-template.md`](cmd-pingcli-pingfederate-ping-one-connections-template.md) |
| `pingcli pingfederate ping-one-connections usage` | Read a PingOne connection's resource usage | [`cmd-pingcli-pingfederate-ping-one-connections-usage.md`](cmd-pingcli-pingfederate-ping-one-connections-usage.md) |

## Parent Command

- [`pingcli pingfederate`](cmd-pingcli-pingfederate.md) — Administration tools for PingFederate deployed as software
