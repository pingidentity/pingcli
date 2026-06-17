# `pingcli pingone davinci flows`
DaVinci Flows

## Synopsis

DaVinci Flows

```
pingcli pingone davinci flows [flags]
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
| `pingcli pingone davinci flows apply` | Create or update a DaVinci flow | [`cmd-pingcli-pingone-davinci-flows-apply.md`](cmd-pingcli-pingone-davinci-flows-apply.md) |
| `pingcli pingone davinci flows clone` | Clone a DaVinci flow | [`cmd-pingcli-pingone-davinci-flows-clone.md`](cmd-pingcli-pingone-davinci-flows-clone.md) |
| `pingcli pingone davinci flows create` | Create a new DaVinci flow | [`cmd-pingcli-pingone-davinci-flows-create.md`](cmd-pingcli-pingone-davinci-flows-create.md) |
| `pingcli pingone davinci flows delete` | Delete a DaVinci flow | [`cmd-pingcli-pingone-davinci-flows-delete.md`](cmd-pingcli-pingone-davinci-flows-delete.md) |
| `pingcli pingone davinci flows deploy` | Deploy a DaVinci flow | [`cmd-pingcli-pingone-davinci-flows-deploy.md`](cmd-pingcli-pingone-davinci-flows-deploy.md) |
| `pingcli pingone davinci flows get` | Read a DaVinci flow | [`cmd-pingcli-pingone-davinci-flows-get.md`](cmd-pingcli-pingone-davinci-flows-get.md) |
| `pingcli pingone davinci flows list` | List DaVinci flows | [`cmd-pingcli-pingone-davinci-flows-list.md`](cmd-pingcli-pingone-davinci-flows-list.md) |
| `pingcli pingone davinci flows replace` | Update a DaVinci flow | [`cmd-pingcli-pingone-davinci-flows-replace.md`](cmd-pingcli-pingone-davinci-flows-replace.md) |
| `pingcli pingone davinci flows template` | Generate a DaVinci flow JSON template | [`cmd-pingcli-pingone-davinci-flows-template.md`](cmd-pingcli-pingone-davinci-flows-template.md) |

## Parent Command

- [`pingcli pingone davinci`](cmd-pingcli-pingone-davinci.md) — Administration tools for the PingOne DaVinci universal service.
