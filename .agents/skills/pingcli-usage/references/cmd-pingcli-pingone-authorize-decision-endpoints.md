# `pingcli pingone authorize decision-endpoints`
Decision Endpoints

## Synopsis

Decision Endpoints

```
pingcli pingone authorize decision-endpoints [flags]
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
| `pingcli pingone authorize decision-endpoints apply` | Create or update a decision endpoint | [`cmd-pingcli-pingone-authorize-decision-endpoints-apply.md`](cmd-pingcli-pingone-authorize-decision-endpoints-apply.md) |
| `pingcli pingone authorize decision-endpoints create` | Create a new decision endpoint | [`cmd-pingcli-pingone-authorize-decision-endpoints-create.md`](cmd-pingcli-pingone-authorize-decision-endpoints-create.md) |
| `pingcli pingone authorize decision-endpoints delete` | Delete a decision endpoint | [`cmd-pingcli-pingone-authorize-decision-endpoints-delete.md`](cmd-pingcli-pingone-authorize-decision-endpoints-delete.md) |
| `pingcli pingone authorize decision-endpoints get` | Read a specific decision endpoint | [`cmd-pingcli-pingone-authorize-decision-endpoints-get.md`](cmd-pingcli-pingone-authorize-decision-endpoints-get.md) |
| `pingcli pingone authorize decision-endpoints list` | List all decision endpoints | [`cmd-pingcli-pingone-authorize-decision-endpoints-list.md`](cmd-pingcli-pingone-authorize-decision-endpoints-list.md) |
| `pingcli pingone authorize decision-endpoints replace` | Update a decision endpoint | [`cmd-pingcli-pingone-authorize-decision-endpoints-replace.md`](cmd-pingcli-pingone-authorize-decision-endpoints-replace.md) |
| `pingcli pingone authorize decision-endpoints template` | Generate a decision endpoint JSON template | [`cmd-pingcli-pingone-authorize-decision-endpoints-template.md`](cmd-pingcli-pingone-authorize-decision-endpoints-template.md) |

## Parent Command

- [`pingcli pingone authorize`](cmd-pingcli-pingone-authorize.md) — Administration tools for the PingOne Authorize universal service.
