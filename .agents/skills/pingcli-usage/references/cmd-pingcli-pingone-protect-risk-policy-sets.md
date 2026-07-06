# `pingcli pingone protect risk-policy-sets`
Risk Policy Sets

## Synopsis

Risk Policy Sets

```
pingcli pingone protect risk-policy-sets [flags]
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
| `pingcli pingone protect risk-policy-sets apply` | Create or update a risk policy set | [`cmd-pingcli-pingone-protect-risk-policy-sets-apply.md`](cmd-pingcli-pingone-protect-risk-policy-sets-apply.md) |
| `pingcli pingone protect risk-policy-sets create` | Create a new risk policy set | [`cmd-pingcli-pingone-protect-risk-policy-sets-create.md`](cmd-pingcli-pingone-protect-risk-policy-sets-create.md) |
| `pingcli pingone protect risk-policy-sets delete` | Delete a risk policy set | [`cmd-pingcli-pingone-protect-risk-policy-sets-delete.md`](cmd-pingcli-pingone-protect-risk-policy-sets-delete.md) |
| `pingcli pingone protect risk-policy-sets get` | Read a specific risk policy set | [`cmd-pingcli-pingone-protect-risk-policy-sets-get.md`](cmd-pingcli-pingone-protect-risk-policy-sets-get.md) |
| `pingcli pingone protect risk-policy-sets list` | List all risk policy sets | [`cmd-pingcli-pingone-protect-risk-policy-sets-list.md`](cmd-pingcli-pingone-protect-risk-policy-sets-list.md) |
| `pingcli pingone protect risk-policy-sets replace` | Replace a risk policy set | [`cmd-pingcli-pingone-protect-risk-policy-sets-replace.md`](cmd-pingcli-pingone-protect-risk-policy-sets-replace.md) |
| `pingcli pingone protect risk-policy-sets template` | Generate a risk policy set JSON template | [`cmd-pingcli-pingone-protect-risk-policy-sets-template.md`](cmd-pingcli-pingone-protect-risk-policy-sets-template.md) |

## Parent Command

- [`pingcli pingone protect`](cmd-pingcli-pingone-protect.md) — Administration tools for the PingOne Protect universal service.
