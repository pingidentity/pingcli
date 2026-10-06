### ENHANCEMENTS

- Added a notice when a newer Ping CLI release is available, shown after interactive commands and by `pingcli --version`; opt out by setting `skipUpdateCheck` in a configuration profile or the `PINGCLI_SKIP_UPDATE_CHECK` environment variable
- Added suggestions for the closest matching command when a subcommand is mistyped anywhere in the CLI
- Improved token exchange error messages by validating authorization server token responses and surfacing the server's error details, including under --debug
- PingOne: Made `--environment-id` optional on commands scoped to a PingOne environment, resolving it from `PINGCLI_PINGONE_ENVIRONMENT_ID` or the active profile's `service.pingOne.endpoint.environmentID` when omitted

### BUG FIXES

- Fixed login failures with "data passed to Set was too big" by storing oversized authentication sessions across multiple OS keychain entries instead of a single entry
- PingOne: Removed the non-functional `--client-credentials`, `--authorization-code`, and `--device-code` flags from `pingone auth login`, which were listed in help but never changed how authentication worked

