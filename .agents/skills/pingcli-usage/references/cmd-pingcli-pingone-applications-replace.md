# `pingcli pingone applications replace`
Replace an application

## Synopsis

Replace (PUT) an application in a PingOne environment.

The body supplied via --from-file must use the ApplicationBody wrapper shape:
exactly one nested protocol block — "oidc", "saml", "wsfed", or "externalLink"
— containing all application fields (including name and enabled).

Output from get/list uses the SDK flat polymorphic shape where "protocol",
"type", and all variant fields are at the top level. Because the input and output
shapes differ, piping a get result directly into replace requires a manual
reshape: move all fields under the appropriate nested block and remove server-set
read-only fields (id, _links, createdAt, updatedAt, environment).

```
pingcli pingone applications replace [flags]
```

## Examples

```
# Replace an application from a JSON file (ApplicationBody wrapper shape: one nested protocol block)
  # Note: get/list output uses the SDK flat shape; reshape before piping into replace.
  pingcli pingone applications replace --environment-id <env-id> --application-id <app-id> --from-file application.json

  # Replace an application from stdin
  pingcli pingone applications replace --environment-id <env-id> --application-id <app-id> --from-file - < application.json
```

## Options

| Flag | Default | Description |
|------|---------|-------------|
| `-h, --help` | `` | help for replace |
| `-a, --application-id string` | `` | The application ID |
| `-e, --environment-id string` | `` | The PingOne environment ID |
| `-f, --from-file string` | `` | Path to a JSON file containing the request body, or "-" to read from stdin. |


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

- [`pingcli pingone applications`](cmd-pingcli-pingone-applications.md) — Applications
