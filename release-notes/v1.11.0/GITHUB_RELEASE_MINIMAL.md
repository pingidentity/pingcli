### FEATURES

- Reserved `-i` as the consistent shorthand for each resource's own identifier in CRUD commands. Existing shorthand flags for self-identifiers are maintained for backwards-compatibility, but have been marked as deprecated.

### ENHANCEMENTS

- PingOne: Added a warning to `pingone auth login` when it reuses a still-valid session that has no refresh token, with the steps to enable refresh tokens
- Restored shell autocompletion for the `--profile` and `--output-format` global flags, the `--http-method` flag of `pingcli <product> request`, and the profile-name arguments of `config profiles use`, `config profiles delete`, and `config profiles show`; shell completion for `use` and `delete` omits the active profile

### BUG FIXES

- Fixed `--profile` with a profile that does not exist running guided setup prompts or reporting that authentication is not configured, so commands now report that the profile does not exist and how to create it
- Fixed `pingcli config get`, `config set`, and `config unset` ignoring the `PINGCLI_PROFILE` environment variable, so with it exported `config set` and `config unset` now modify that profile instead of the active profile
- Fixed authentication sessions being stored under the active profile instead of the profile selected with `--profile` or `PINGCLI_PROFILE`, so login, logout, and later commands now use the same session
- Fixed guided authentication setup saving its answers to the active profile instead of the profile selected with `--profile` or `PINGCLI_PROFILE`, and it now stops before prompting when that profile does not exist
- Fixed the `api` command sending an empty request body when the file passed to `--data` does not exist; it now reports an error before sending the request
- PingOne: Fixed `pingone auth refresh` reporting an expired session for failures that are not expiry, so a missing refresh token now explains how to enable refresh tokens and other failures, such as network errors or incorrect client credentials, show their underlying cause
- When the API returned a successful response with an empty body, a small subset of commands crashed; they now display an empty result

### NOTES

- Removed the unused `telemetry.tls.insecureSkipVerify` configuration option (environment variable `OTEL_EXPORTER_OTLP_INSECURE`); the setting was registered but never read by the CLI

