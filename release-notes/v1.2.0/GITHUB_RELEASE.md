### BREAKING CHANGES

[bc390a6b](https://github.com/ping-internal/ppds-pingcli/commit/bc390a6b) PingFederate: Removed the `--pingfederate-*` and `--as-pingfederate-*` authentication CLI flags (such as `--pingfederate-client-id`); configure PingFederate authentication via environment variables, the configuration file, or the init wizard instead [#514](https://github.com/ping-internal/ppds-pingcli/pull/514)

### SECURITY

[96c71892](https://github.com/ping-internal/ppds-pingcli/commit/96c71892) PingOne and PingFederate: Refresh the access token when authentication credentials change so a cached token issued for previous credentials is no longer reused [#505](https://github.com/ping-internal/ppds-pingcli/pull/505)

### FEATURES

[d07114f2](https://github.com/ping-internal/ppds-pingcli/commit/d07114f2) Added `credential-issuer-profile` resource to the PingOne Credentials connector, supporting `get`, `replace`, and `template` commands [#482](https://github.com/ping-internal/ppds-pingcli/pull/482)

### ENHANCEMENTS

[66b4ad99](https://github.com/ping-internal/ppds-pingcli/commit/66b4ad99) PingOne: Added management commands for FIDO2 policies [#464](https://github.com/ping-internal/ppds-pingcli/pull/464)
[ac0dd16d](https://github.com/ping-internal/ppds-pingcli/commit/ac0dd16d) PingOne: Added management commands for admin config [#470](https://github.com/ping-internal/ppds-pingcli/pull/470)
[078246b9](https://github.com/ping-internal/ppds-pingcli/commit/078246b9) PingFederate: Added management commands for administrative accounts [#474](https://github.com/ping-internal/ppds-pingcli/pull/474)
[078246b9](https://github.com/ping-internal/ppds-pingcli/commit/078246b9) PingFederate: Added management commands for authentication API settings [#481](https://github.com/ping-internal/ppds-pingcli/pull/481)
[9ff90b47](https://github.com/ping-internal/ppds-pingcli/commit/9ff90b47) PingFederate: Added management commands for authentication API applications [#483](https://github.com/ping-internal/ppds-pingcli/pull/483)
[0fa797d2](https://github.com/ping-internal/ppds-pingcli/commit/0fa797d2) PingOne Protect: Added management commands for risk policy sets [#485](https://github.com/ping-internal/ppds-pingcli/pull/485)
[6b205685](https://github.com/ping-internal/ppds-pingcli/commit/6b205685) PingOne Authorize: Added management commands for decision endpoints [#486](https://github.com/ping-internal/ppds-pingcli/pull/486)
[f714a694](https://github.com/ping-internal/ppds-pingcli/commit/f714a694) PingFederate: Added management commands for server settings [#488](https://github.com/ping-internal/ppds-pingcli/pull/488)
[ff3f8c9f](https://github.com/ping-internal/ppds-pingcli/commit/ff3f8c9f) PingFederate: Added management commands for data stores [#489](https://github.com/ping-internal/ppds-pingcli/pull/489)
[f765d6f5](https://github.com/ping-internal/ppds-pingcli/commit/f765d6f5) PingFederate: Added management command for version [#490](https://github.com/ping-internal/ppds-pingcli/pull/490)
[d8cacfb1](https://github.com/ping-internal/ppds-pingcli/commit/d8cacfb1) PingFederate: Added management commands for virtual host names [#493](https://github.com/ping-internal/ppds-pingcli/pull/493)
[136e5ec5](https://github.com/ping-internal/ppds-pingcli/commit/136e5ec5) PingOne: Added management commands for custom domains [#495](https://github.com/ping-internal/ppds-pingcli/pull/495)
[e4237f38](https://github.com/ping-internal/ppds-pingcli/commit/e4237f38) PingFederate: Added management commands for IdP adapters [#500](https://github.com/ping-internal/ppds-pingcli/pull/500)
[8aa0f4d0](https://github.com/ping-internal/ppds-pingcli/commit/8aa0f4d0) PingFederate: Added management commands for password credential validators [#503](https://github.com/ping-internal/ppds-pingcli/pull/503)
[14a6f403](https://github.com/ping-internal/ppds-pingcli/commit/14a6f403) PingOne: Added management commands for user digital wallets [#506](https://github.com/ping-internal/ppds-pingcli/pull/506)
[4c2dbdb5](https://github.com/ping-internal/ppds-pingcli/commit/4c2dbdb5) PingOne: Added management commands for credential issuance rules [#507](https://github.com/ping-internal/ppds-pingcli/pull/507)
[06815d5a](https://github.com/ping-internal/ppds-pingcli/commit/06815d5a) PingFederate: Added management commands for local identity profiles [#508](https://github.com/ping-internal/ppds-pingcli/pull/508)
[89ad97f9](https://github.com/ping-internal/ppds-pingcli/commit/89ad97f9) PingFederate: Added management commands for OAuth clients [#512](https://github.com/ping-internal/ppds-pingcli/pull/512)
[1059f716](https://github.com/ping-internal/ppds-pingcli/commit/1059f716) PingFederate: Added management commands for extended properties [#516](https://github.com/ping-internal/ppds-pingcli/pull/516)

### BUG FIXES

[1926b669](https://github.com/ping-internal/ppds-pingcli/commit/1926b669) Fixed template command example text for replace-only resources to show `replace --from-file` instead of `create --from-file` [#491](https://github.com/ping-internal/ppds-pingcli/pull/491)
[8e8f3c52](https://github.com/ping-internal/ppds-pingcli/commit/8e8f3c52) auth: Added a connector section header to each `auth status` block so statuses for multiple services are clearly distinguishable [#498](https://github.com/ping-internal/ppds-pingcli/pull/498)
[40888e9a](https://github.com/ping-internal/ppds-pingcli/commit/40888e9a) DaVinci: Fixed template for application resources to correctly omit server-set read-only fields (`apiKey.value`, `oauth.clientSecret`) that were previously included in the generated template skeleton [#502](https://github.com/ping-internal/ppds-pingcli/pull/502)
[06f1f780](https://github.com/ping-internal/ppds-pingcli/commit/06f1f780) PingOne: Fixed the `--template` flag on `pingone api` so template field traversal works against JSON response data [#515](https://github.com/ping-internal/ppds-pingcli/pull/515)

### DEPRECATIONS

[bc390a6b](https://github.com/ping-internal/ppds-pingcli/commit/bc390a6b) PingFederate: The `accessTokenAuth` authentication type is deprecated and will be removed in a future release; use the `basicAuth` or `oauth` authentication type instead [#514](https://github.com/ping-internal/ppds-pingcli/pull/514)

### DOCUMENTATION

[cf9ceca7](https://github.com/ping-internal/ppds-pingcli/commit/cf9ceca7) Updated output schema reference generator with editorial fixes and subfolder restructure [#480](https://github.com/ping-internal/ppds-pingcli/pull/480)

