### BREAKING CHANGES

- Remove `--pingone-*` authentication flag options in favor of configuration file and environment variable settings

### ENHANCEMENTS

- Config migration now removes the deprecated `export.*` configuration subtree, which is no longer persisted to the configuration file (the export command's settings are flag/environment-only).
- Add more detail to response metadata - HTTP status, request ID, duration, resource, action, effect. Ensure meta fields are populated on error.
- Include HTTP status and request ID in response metadata.
- PingOne: Added management commands for custom admin roles.
- PingOne: Added management commands for active identity counts.
- Added `ndjson-typed` output format.
- PingOne: Auto-login and auto-refresh for `authorization_code` and `device_code` grants.

### BUG FIXES

- Both `PINGCLI_PINGONE_*` and bare `PINGONE_*` environment variable names are now fully supported (e.g. both `PINGCLI_PINGONE_ENVIRONMENT_ID` and `PINGONE_ENVIRONMENT_ID` resolve), fixing a silent breaking change for pipelines using the pre-v1 names. `PINGCLI_LOGIN_STORAGE_TYPE` is honoured alongside `PINGCLI_AUTH_STORAGE_TYPE`.
- Config migration now preserves PingOne authentication values written by older (v0.8.x) Ping CLI versions: the legacy `service.pingone.*` worker layout (worker client credentials, `regioncode`, `type: worker`) is migrated to the current `service.pingOne.*` schema instead of being silently dropped, with `regioncode` translated to the equivalent `endpoint.rootDomain`. The PingOne environment ID and authorization-code `redirectUri` path/port are also relocated to their canonical keys rather than being lost.
- Config migration now correctly relocates PingFederate `clientCredentialsAuth` credentials (client ID, client secret, token URL, scopes) to the canonical `oauth.clientCredentials.*` keys and removes the legacy subtree, including for configs written all-lowercase by older (v0.8.x) Ping CLI versions, which were previously left untouched.
- Fix `ndjson` output printing JSON bytes directly
- Suppress Configuration Management Commands header in `--help` output when command has no children
- Fix `oauth authentication type is not yet implemented` error when enabling client credentials authentication for PingFederate

### DOCUMENTATION

- Update PingOne Roles (built-in admin roles) description to distinguish from other role types
