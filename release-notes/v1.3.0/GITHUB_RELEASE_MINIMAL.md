### BREAKING CHANGES

- Removed the --template output flag and its custom text/template rendering

### ENHANCEMENTS

- Add support for the `apply` command to all singleton resources as an alias for `replace`. Also add `apply` to the PingFederate data-stores and PingOne application-resource-permissions resources where it was missing.
- Renamed several nested subcommands to drop the redundant parent prefix (for example, `group group-nestings` is now `group nestings`, `verify verify-policies` is now `verify policies`, etc), while keeping the previous names available as aliases

- PingFederate: Added management commands for OAuth/OpenID Connect policies
- PingFederate: Added management commands for OAuth/OpenID Connect settings
- PingFederate: Added management commands for federation info
- PingFederate: Added management commands for OAuth client registration policies
- PingFederate: Added management commands for OAuth client settings
- PingFederate: Added management commands for session quotas
- PingFederate: Added management commands for session settings
- PingFederate: Added management commands for authentication policy fragments

- PingOne: Added `gateways gateway-credentials get` command to read a single gateway credential
- PingOne: Added a list command for resource application permissions
- PingOne: Added management commands for images
- PingOne: Failing paged export read operations now attach the HTTP response so `--debug` renders the response body and JSON error output reports the failing request's status and request ID
- PingOne: Added management commands for credential issuance rule usage counts and usage details
- PingOne: Added management commands for identity propagation plans
- PingOne: Added management commands for alert channels
- PingOne: Added management commands for integrations
- PingOne: Added notifications-settings management commands
- PingOne: Added management commands for email domains
- PingOne: Added pingone applications attribute-mappings management commands
- PingOne: Added management commands for reCAPTCHA v2 configuration
- PingOne: Added management commands for rate limit IP configurations

### BUG FIXES

- Fix some API error detail not being returned by the CLI for PingOne Authorize and PingOne Verify endpoints.

