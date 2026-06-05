# `pingcli agent-skills install`
Install an agent skill

## Synopsis

Install an agent skill to the output directory.
The skill directory and all its files are copied to <output-dir>/<skill-name>.
The output directory defaults to .claude/skills in the current working directory and is
created automatically if it does not exist.

Available skills:
  pingcli-usage

```
pingcli agent-skills install skill-name [flags]
```

## Examples

```
Install an agent skill to the default location (.claude/skills in the current directory).
    pingcli agent-skills install pingcli-usage

  Install an agent skill to a custom directory.
    pingcli agent-skills install pingcli-usage --output-dir /path/to/my/skills
```

## Options

| Flag | Default | Description |
|------|---------|-------------|
| `-h, --help` | `` | help for install |
| `--output-dir string` | `` | Directory to install the skill into. The skill will be placed at <output-dir>/<skill-name>. Defaults to .claude/skills in the current working directory. |


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


## Parent Command

- [`pingcli agent-skills`](cmd-pingcli-agent-skills.md) — Find and install agent skills for Ping CLI.
