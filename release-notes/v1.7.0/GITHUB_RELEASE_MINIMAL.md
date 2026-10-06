### BREAKING CHANGES

- PingFederate: Added a required `--from-file` flag to `invoke-action` for IDP adapters, data stores, and secret managers to support actions that require parameters
- The `apply` command no longer accepts ID flags such as `--id` (they are removed from help text, and passing one returns an error). The ID is assigned automatically when the resource already exists

### ENHANCEMENTS

- PingFederate: Added commands for out-of-band authenticator plugin descriptors and instance actions, including action invocation
- PingFederate: Added management commands for access token manager descriptors
- PingFederate: Added management commands for authentication policy contract to SP adapter mappings
- PingFederate: Added management commands for cluster status
- PingFederate: Added management commands for IdP adapter descriptors
- PingFederate: Added management commands for IdP STS request parameters contracts
- PingFederate: Added management commands for IdP-to-SP adapter mappings
- PingFederate: Added management commands for OAuth authorization detail processor descriptors
- PingFederate: Added management commands for OAuth client registration policy descriptors
- PingFederate: Added management commands for OAuth token exchange generator settings
- PingFederate: Added management commands for outbound provisioning settings
- PingFederate: Added management commands for PingOne for Enterprise settings
- PingFederate: Added management commands for protocol metadata lifetime settings
- PingFederate: Added management commands for protocol metadata signing settings
- PingFederate: Added management commands for server settings notifications
- PingFederate: Added management commands for service authentication
- PingFederate: Added management commands for SP adapter descriptors and adapter actions
- PingFederate: Added management commands for SP token generator descriptors
- PingFederate: Added management commands for token exchange generator groups
- PingFederate: Added management commands for token exchange processor policy to token generator mappings
- PingFederate: Added management commands for token exchange processor settings
- PingFederate: Added management commands for token generators
- PingFederate: Added management commands for WS-Trust STS Settings and issuer certificates
- PingFederate: Added support for moving authentication policies within the policy tree
- PingFederate: Added the revoke-secondary-secrets action to OAuth clients management commands

### BUG FIXES

- Fixed the CLI always rejecting empty request bodies

