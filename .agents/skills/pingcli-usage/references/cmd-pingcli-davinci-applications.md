# `pingcli davinci applications`
DaVinci Applications

## Synopsis

DaVinci Applications

```
pingcli davinci applications [flags]
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
| `pingcli davinci applications apply` | Create or update a DaVinci application | [`cmd-pingcli-davinci-applications-apply.md`](cmd-pingcli-davinci-applications-apply.md) |
| `pingcli davinci applications create` | Create a new DaVinci application | [`cmd-pingcli-davinci-applications-create.md`](cmd-pingcli-davinci-applications-create.md) |
| `pingcli davinci applications delete` | Delete a DaVinci application | [`cmd-pingcli-davinci-applications-delete.md`](cmd-pingcli-davinci-applications-delete.md) |
| `pingcli davinci applications flow-policies` | DaVinci Flow Policies | [`cmd-pingcli-davinci-applications-flow-policies.md`](cmd-pingcli-davinci-applications-flow-policies.md) |
| `pingcli davinci applications get` | Read a DaVinci application | [`cmd-pingcli-davinci-applications-get.md`](cmd-pingcli-davinci-applications-get.md) |
| `pingcli davinci applications list` | List DaVinci applications | [`cmd-pingcli-davinci-applications-list.md`](cmd-pingcli-davinci-applications-list.md) |
| `pingcli davinci applications replace` | Update a DaVinci application | [`cmd-pingcli-davinci-applications-replace.md`](cmd-pingcli-davinci-applications-replace.md) |
| `pingcli davinci applications rotate-key` | Rotate a DaVinci application API key | [`cmd-pingcli-davinci-applications-rotate-key.md`](cmd-pingcli-davinci-applications-rotate-key.md) |
| `pingcli davinci applications rotate-secret` | Rotate a DaVinci application OAuth secret | [`cmd-pingcli-davinci-applications-rotate-secret.md`](cmd-pingcli-davinci-applications-rotate-secret.md) |
| `pingcli davinci applications template` | Generate a DaVinci application JSON template | [`cmd-pingcli-davinci-applications-template.md`](cmd-pingcli-davinci-applications-template.md) |

## Parent Command

- [`pingcli davinci`](cmd-pingcli-davinci.md) — Administration tools for the PingOne DaVinci universal service.
