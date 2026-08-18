# `pingcli pingone mfa fido2-policies replace`
Update a FIDO2 policy

## Synopsis

Update (replace) a FIDO2 policy in a PingOne environment with a full-body replacement

```
pingcli pingone mfa fido2-policies replace [flags]
```

## Examples

```
# Update a FIDO2 policy from a JSON file (--environment-id and --fido2-policy-id are still required)
  pingcli pingone mfa fido2-policies replace --environment-id <env-id> --fido2-policy-id <policy-id> --from-file fido2-policy.json

  # Update a FIDO2 policy from stdin
  pingcli pingone mfa fido2-policies replace --environment-id <env-id> --fido2-policy-id <policy-id> --from-file - < fido2-policy.json

  # Update a FIDO2 policy, overriding scalar fields with flags
  # ("backupEligibility", "mdsAuthenticatorsRequirements", "userDisplayNameAttributes",
  # and "userVerification" always come from the file - they cannot be set as flags)
  pingcli pingone mfa fido2-policies replace --environment-id <env-id> --fido2-policy-id <policy-id> --name "Default FIDO2 Policy" --from-file fido2-policy.json
```

## Options

| Flag | Default | Description |
|------|---------|-------------|
| `-h, --help` | `` | help for replace |
| `-e, --environment-id string` | `` | The PingOne environment ID |
| `-f, --from-file string` | `` | Path to a JSON file containing the request body, or "-" to read from stdin. |
| `--attestation-requirements string` | `` | The level of attestation required from authenticators |
| `--authenticator-attachment string` | `` | The allowed authenticator attachment modality |
| `--default` | `` | Whether this policy should serve as the environment default FIDO2 policy |
| `--description string` | `` | The description of the FIDO2 policy |
| `--device-display-name string` | `` | The name to display for the device in registration and authentication windows |
| `--discoverable-credentials string` | `` | The server-side discoverable credential preference |
| `--fido2-policy-id string` | `` | The FIDO2 policy ID |
| `--name string` | `` | The name to use for the FIDO2 policy |
| `--public-key-credential-hint []string` | `` | Ordered hints for preferred public-key credential types; repeatable or comma-separated |
| `--relying-party-id string` | `` | The ID of the relying party, typically a domain name |


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

- [`pingcli pingone mfa fido2-policies`](cmd-pingcli-pingone-mfa-fido2-policies.md) — FIDO2 Policies
