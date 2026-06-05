# `pingcli pingone protect risk-predictors`
Risk Predictors

## Synopsis

Risk Predictors

```
pingcli pingone protect risk-predictors [flags]
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
| `pingcli pingone protect risk-predictors apply` | Create or update a risk predictor | [`cmd-pingcli-pingone-protect-risk-predictors-apply.md`](cmd-pingcli-pingone-protect-risk-predictors-apply.md) |
| `pingcli pingone protect risk-predictors create` | Create a new risk predictor | [`cmd-pingcli-pingone-protect-risk-predictors-create.md`](cmd-pingcli-pingone-protect-risk-predictors-create.md) |
| `pingcli pingone protect risk-predictors delete` | Delete a risk predictor | [`cmd-pingcli-pingone-protect-risk-predictors-delete.md`](cmd-pingcli-pingone-protect-risk-predictors-delete.md) |
| `pingcli pingone protect risk-predictors get` | Read a specific risk predictor | [`cmd-pingcli-pingone-protect-risk-predictors-get.md`](cmd-pingcli-pingone-protect-risk-predictors-get.md) |
| `pingcli pingone protect risk-predictors list` | List all risk predictors | [`cmd-pingcli-pingone-protect-risk-predictors-list.md`](cmd-pingcli-pingone-protect-risk-predictors-list.md) |
| `pingcli pingone protect risk-predictors replace` | Replace a risk predictor | [`cmd-pingcli-pingone-protect-risk-predictors-replace.md`](cmd-pingcli-pingone-protect-risk-predictors-replace.md) |
| `pingcli pingone protect risk-predictors template` | Generate a risk predictor JSON template | [`cmd-pingcli-pingone-protect-risk-predictors-template.md`](cmd-pingcli-pingone-protect-risk-predictors-template.md) |

## Parent Command

- [`pingcli pingone protect`](cmd-pingcli-pingone-protect.md) — Administration tools for the PingOne Protect universal service.
