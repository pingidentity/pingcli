# `pingcli protect risk-predictors replace`
Replace a risk predictor

## Synopsis

Replace (PUT) a risk predictor in a PingOne environment.

The body supplied via --from-file must use the RiskPredictorBody wrapper shape:
exactly one nested variant block — "adversaryInTheMiddle", "anonymousNetwork",
"bot", "composite", "custom", "device", "emailReputation", "geovelocity",
"ipReputation", "trafficAnomaly", "userLocationAnomaly", "userRiskBehavior",
or "velocity" — containing all fields including name.

Note: compactName is immutable after creation but must be included in replace
bodies with the original value. The SDK always serialises it and the API
rejects an empty compactName string.

Output from get/list uses the SDK flat polymorphic shape where "type" and all
variant fields are at the top level. Because the input and output shapes differ,
piping a get result directly into replace requires a manual reshape: move all
fields under the appropriate nested block and remove server-set read-only fields
(id, _links, createdAt, updatedAt, environment, licensed, deletable).

```
pingcli protect risk-predictors replace [flags]
```

## Examples

```
# Replace a risk predictor from a JSON file (RiskPredictorBody wrapper shape: one nested variant block)
  # Note: include compactName with the original value — the SDK serialises it and the API rejects an empty value.
  # Note: get/list output uses the SDK flat shape; reshape before piping into replace.
  pingcli protect risk-predictors replace --environment-id <env-id> --risk-predictor-id <predictor-id> --from-file risk-predictor.json

  # Replace a risk predictor from stdin
  pingcli protect risk-predictors replace --environment-id <env-id> --risk-predictor-id <predictor-id> --from-file - < risk-predictor.json
```

## Options

| Flag | Default | Description |
|------|---------|-------------|
| `-h, --help` | `` | help for replace |
| `-e, --environment-id string` | `` | The PingOne environment ID |
| `-f, --from-file string` | `` | Path to a JSON file containing the request body, or "-" to read from stdin. |
| `-r, --risk-predictor-id string` | `` | The risk predictor ID |


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

- [`pingcli protect risk-predictors`](cmd-pingcli-protect-risk-predictors.md) — Risk Predictors
