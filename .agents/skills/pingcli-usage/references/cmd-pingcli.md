# `pingcli`
A CLI tool for managing the configuration of Ping Identity products.

## Synopsis

A CLI tool for managing the configuration of Ping Identity products.

```
pingcli
```

## Options

| Flag | Default | Description |
|------|---------|-------------|
| `-h, --help` | `` | help for pingcli |
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
| `pingcli agent-skills` | Find and install agent skills for Ping CLI. | [`cmd-pingcli-agent-skills.md`](cmd-pingcli-agent-skills.md) |
| `pingcli auth` | Authenticate (or refresh authentication) for multiple connected products and services at once. | [`cmd-pingcli-auth.md`](cmd-pingcli-auth.md) |
| `pingcli authorize` | Administration tools for the PingOne Authorize universal service. | [`cmd-pingcli-authorize.md`](cmd-pingcli-authorize.md) |
| `pingcli completion` | Prints shell completion scripts | [`cmd-pingcli-completion.md`](cmd-pingcli-completion.md) |
| `pingcli config` | Manage the CLI configuration. | [`cmd-pingcli-config.md`](cmd-pingcli-config.md) |
| `pingcli credentials` | Administration tools for the PingOne Credentials universal service. | [`cmd-pingcli-credentials.md`](cmd-pingcli-credentials.md) |
| `pingcli davinci` | Administration tools for the PingOne DaVinci universal service. | [`cmd-pingcli-davinci.md`](cmd-pingcli-davinci.md) |
| `pingcli feedback` | Help us improve the CLI. Report issues or send us feedback on using the CLI tool. | [`cmd-pingcli-feedback.md`](cmd-pingcli-feedback.md) |
| `pingcli init` | Initialize Ping CLI with a guided setup wizard. | [`cmd-pingcli-init.md`](cmd-pingcli-init.md) |
| `pingcli mfa` | Administration tools for the PingOne MFA universal service. | [`cmd-pingcli-mfa.md`](cmd-pingcli-mfa.md) |
| `pingcli pingfederate` | Administration tools for PingFederate deployed as software | [`cmd-pingcli-pingfederate.md`](cmd-pingcli-pingfederate.md) |
| `pingcli pingone` | Administration tools for the PingOne platform. | [`cmd-pingcli-pingone.md`](cmd-pingcli-pingone.md) |
| `pingcli protect` | Administration tools for the PingOne Protect universal service. | [`cmd-pingcli-protect.md`](cmd-pingcli-protect.md) |
| `pingcli verify` | Administration tools for the PingOne Verify universal service. | [`cmd-pingcli-verify.md`](cmd-pingcli-verify.md) |
