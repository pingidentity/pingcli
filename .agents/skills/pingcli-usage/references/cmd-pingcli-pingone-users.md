# `pingcli pingone users`
Users

## Synopsis

Users

```
pingcli pingone users [flags]
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
| `pingcli pingone users apply` | Create or update a user | [`cmd-pingcli-pingone-users-apply.md`](cmd-pingcli-pingone-users-apply.md) |
| `pingcli pingone users create` | Create a new user | [`cmd-pingcli-pingone-users-create.md`](cmd-pingcli-pingone-users-create.md) |
| `pingcli pingone users delete` | Delete a user | [`cmd-pingcli-pingone-users-delete.md`](cmd-pingcli-pingone-users-delete.md) |
| `pingcli pingone users get` | Read a specific user | [`cmd-pingcli-pingone-users-get.md`](cmd-pingcli-pingone-users-get.md) |
| `pingcli pingone users list` | List all users | [`cmd-pingcli-pingone-users-list.md`](cmd-pingcli-pingone-users-list.md) |
| `pingcli pingone users replace` | Update a user | [`cmd-pingcli-pingone-users-replace.md`](cmd-pingcli-pingone-users-replace.md) |
| `pingcli pingone users template` | Generate a user JSON template | [`cmd-pingcli-pingone-users-template.md`](cmd-pingcli-pingone-users-template.md) |
| `pingcli pingone users user-application-role-assignments` | User Application Role Assignments | [`cmd-pingcli-pingone-users-user-application-role-assignments.md`](cmd-pingcli-pingone-users-user-application-role-assignments.md) |
| `pingcli pingone users user-identity-provider` | User Identity Provider | [`cmd-pingcli-pingone-users-user-identity-provider.md`](cmd-pingcli-pingone-users-user-identity-provider.md) |
| `pingcli pingone users user-population` | User Population | [`cmd-pingcli-pingone-users-user-population.md`](cmd-pingcli-pingone-users-user-population.md) |
| `pingcli pingone users user-role-assignments` | User Role Assignments | [`cmd-pingcli-pingone-users-user-role-assignments.md`](cmd-pingcli-pingone-users-user-role-assignments.md) |

## Parent Command

- [`pingcli pingone`](cmd-pingcli-pingone.md) — Administration tools for the PingOne platform.
