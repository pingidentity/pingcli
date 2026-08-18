# `pingcli pingfederate captcha-providers replace`
Update a CAPTCHA provider

## Synopsis

Update (replace) a PingFederate CAPTCHA provider

```
pingcli pingfederate captcha-providers replace [flags]
```

## Examples

```
# Update a CAPTCHA provider from a JSON file (--id is still required)
  pingcli pingfederate captcha-providers replace --id <id> --from-file provider.json

  # Update a CAPTCHA provider from stdin
  pingcli pingfederate captcha-providers replace --id <id> --from-file - < provider.json

  # Update using flags for identity fields, and --from-file for configuration
  pingcli pingfederate captcha-providers replace --id <id> --name "My Provider" --plugin-descriptor-ref-id com.example.Plugin --from-file config.json
```

## Options

| Flag | Default | Description |
|------|---------|-------------|
| `-h, --help` | `` | help for replace |
| `-f, --from-file string` | `` | Path to a JSON file containing the request body, or "-" to read from stdin. |
| `--id string` | `` | The PingFederate CAPTCHA Provider ID |
| `--name string` | `` | Display name of the CAPTCHA provider |
| `--parent-ref-id string` | `` | ID of a parent CAPTCHA provider instance to inherit configuration from |
| `--plugin-descriptor-ref-id string` | `` | ID of the CAPTCHA provider plugin type descriptor |


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


## Parent Command

- [`pingcli pingfederate captcha-providers`](cmd-pingcli-pingfederate-captcha-providers.md) — PingFederate CAPTCHA Providers
