# `pingcli pingone applications`
Applications

## Synopsis

Applications

```
pingcli pingone applications [flags]
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
| `pingcli pingone applications application-grants` | Application Grants | [`cmd-pingcli-pingone-applications-application-grants.md`](cmd-pingcli-pingone-applications-application-grants.md) |
| `pingcli pingone applications application-role-assignments` | Application Role Assignments | [`cmd-pingcli-pingone-applications-application-role-assignments.md`](cmd-pingcli-pingone-applications-application-role-assignments.md) |
| `pingcli pingone applications application-secrets` | Application Secrets | [`cmd-pingcli-pingone-applications-application-secrets.md`](cmd-pingcli-pingone-applications-application-secrets.md) |
| `pingcli pingone applications apply` | Create or update an application | [`cmd-pingcli-pingone-applications-apply.md`](cmd-pingcli-pingone-applications-apply.md) |
| `pingcli pingone applications create` | Create a new application | [`cmd-pingcli-pingone-applications-create.md`](cmd-pingcli-pingone-applications-create.md) |
| `pingcli pingone applications delete` | Delete an application | [`cmd-pingcli-pingone-applications-delete.md`](cmd-pingcli-pingone-applications-delete.md) |
| `pingcli pingone applications flow-policy-assignments` | Flow Policy Assignments | [`cmd-pingcli-pingone-applications-flow-policy-assignments.md`](cmd-pingcli-pingone-applications-flow-policy-assignments.md) |
| `pingcli pingone applications get` | Read a specific application | [`cmd-pingcli-pingone-applications-get.md`](cmd-pingcli-pingone-applications-get.md) |
| `pingcli pingone applications list` | List all applications | [`cmd-pingcli-pingone-applications-list.md`](cmd-pingcli-pingone-applications-list.md) |
| `pingcli pingone applications replace` | Replace an application | [`cmd-pingcli-pingone-applications-replace.md`](cmd-pingcli-pingone-applications-replace.md) |
| `pingcli pingone applications sop-assignments` | Sign-On Policy Assignments | [`cmd-pingcli-pingone-applications-sop-assignments.md`](cmd-pingcli-pingone-applications-sop-assignments.md) |
| `pingcli pingone applications template` | Generate an application JSON template | [`cmd-pingcli-pingone-applications-template.md`](cmd-pingcli-pingone-applications-template.md) |

## Parent Command

- [`pingcli pingone`](cmd-pingcli-pingone.md) — Administration tools for the PingOne platform.
