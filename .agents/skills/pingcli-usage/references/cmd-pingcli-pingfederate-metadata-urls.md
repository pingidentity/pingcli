# `pingcli pingfederate metadata-urls`
PingFederate metadata URLs

## Synopsis

PingFederate metadata URLs configure a URL from which SAML or WS-Federation metadata can be periodically retrieved, along with the certificate used to validate the metadata's signature.

```
pingcli pingfederate metadata-urls [flags]
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
| `pingcli pingfederate metadata-urls apply` | Create or update a metadata URL | [`cmd-pingcli-pingfederate-metadata-urls-apply.md`](cmd-pingcli-pingfederate-metadata-urls-apply.md) |
| `pingcli pingfederate metadata-urls create` | Create a new metadata URL | [`cmd-pingcli-pingfederate-metadata-urls-create.md`](cmd-pingcli-pingfederate-metadata-urls-create.md) |
| `pingcli pingfederate metadata-urls delete` | Delete a metadata URL | [`cmd-pingcli-pingfederate-metadata-urls-delete.md`](cmd-pingcli-pingfederate-metadata-urls-delete.md) |
| `pingcli pingfederate metadata-urls get` | Read a specific metadata URL | [`cmd-pingcli-pingfederate-metadata-urls-get.md`](cmd-pingcli-pingfederate-metadata-urls-get.md) |
| `pingcli pingfederate metadata-urls list` | List all metadata URLs | [`cmd-pingcli-pingfederate-metadata-urls-list.md`](cmd-pingcli-pingfederate-metadata-urls-list.md) |
| `pingcli pingfederate metadata-urls replace` | Update a metadata URL | [`cmd-pingcli-pingfederate-metadata-urls-replace.md`](cmd-pingcli-pingfederate-metadata-urls-replace.md) |
| `pingcli pingfederate metadata-urls template` | Generate a metadata URL JSON template | [`cmd-pingcli-pingfederate-metadata-urls-template.md`](cmd-pingcli-pingfederate-metadata-urls-template.md) |

## Parent Command

- [`pingcli pingfederate`](cmd-pingcli-pingfederate.md) — Administration tools for PingFederate deployed as software
