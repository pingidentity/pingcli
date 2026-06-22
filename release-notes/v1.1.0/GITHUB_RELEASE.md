### ENHANCEMENTS

- Config migration now removes the deprecated `export.*` configuration subtree, which is no longer persisted to the configuration file (the export command's settings are flag/environment-only).

### BUG FIXES

- Both `PINGCLI_PINGONE_*` and bare `PINGONE_*` environment variable names are now fully supported (e.g. both `PINGCLI_PINGONE_ENVIRONMENT_ID` and `PINGONE_ENVIRONMENT_ID` resolve), fixing a silent breaking change for pipelines using the pre-v1 names. `PINGCLI_LOGIN_STORAGE_TYPE` is honoured alongside `PINGCLI_AUTH_STORAGE_TYPE`.
- Config migration now preserves PingOne authentication values written by older (v0.8.x) Ping CLI versions: the legacy `service.pingone.*` worker layout (worker client credentials, `regioncode`, `type: worker`) is migrated to the current `service.pingOne.*` schema instead of being silently dropped, with `regioncode` translated to the equivalent `endpoint.rootDomain`. The PingOne environment ID and authorization-code `redirectUri` path/port are also relocated to their canonical keys rather than being lost.
- Config migration now correctly relocates PingFederate `clientCredentialsAuth` credentials (client ID, client secret, token URL, scopes) to the canonical `oauth.clientCredentials.*` keys and removes the legacy subtree, including for configs written all-lowercase by older (v0.8.x) Ping CLI versions, which were previously left untouched.

## COMPLETE CHANGES
- feat(connector): thread SDK response metadata into envelope meta.api
- docs(pingone): clarify roles command describes built-in admin roles
- feat(pingone): add CustomAdminRole CRUD operations
- feat(output): add -O ndjson-typed stream format and deprecate ndjson-wrapped
- feat(output): add duration_ms, resource, action, effect to envelope meta
- fix(output): populate resource/action/effect/duration_ms on error-path envelopes
- fix(connector): response-meta boundary — meta.api on errors, rich envelopes
- feat(pingone): auto-login and silent refresh for human auth grants
- feat(pingone): add ActiveIdentityCounts read-only resource
- fix(output): stamp duration_ms on ndjson-typed summary line
- feat(config): remove --pingone-* auth flags; full env-var support
- fix(output): treat json.RawMessage as atomic in all emit paths
- fix(configmgmt): suppress empty Configuration Management Commands header
- fix(pingfederate): unblock oauth client_credentials login
- fix(config): v0.8.x migration fixes, PingOne env var aliases, and auth-init cleanup