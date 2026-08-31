### BREAKING CHANGES

- PingFederate: Renamed the `--id` flag to `--name` on OAuth common scope get, replace, and delete commands to match the underlying scope identifier

### ENHANCEMENTS

- PingFederate: Added a change-password action for administrative accounts to change the current account's own password
- PingFederate: Added a command to disconnect PingFederate from a connected PingOne for Enterprise account.
- PingFederate: Added a command to update the PingOne for Enterprise identity repository
- PingFederate: Added incoming proxy settings configuration-management commands
- PingFederate: Added management commands for authentication policies
- PingFederate: Added management commands for authentication session policies
- PingFederate: Added management commands for certificate revocation settings
- PingFederate: Added management commands for identity store provisioner descriptors
- PingFederate: Added management commands for IdP default URLs
- PingFederate: Added management commands for Kerberos realm settings
- PingFederate: Added management commands for Kerberos realms
- PingFederate: Added management commands for license
- PingFederate: Added management commands for license agreement
- PingFederate: Added management commands for metadata URLs
- PingFederate: Added management commands for notification publisher plugin descriptors
- PingFederate: Added management commands for notification publishers
- PingFederate: Added management commands for notification publishers and notification publisher settings
- PingFederate: Added management commands for OAuth CIBA server policy request policies and settings
- PingFederate: Added management commands for OAuth common scope groups
- PingFederate: Added management commands for OAuth exclusive scope groups
- PingFederate: Added management commands for OAuth exclusive scopes
- PingFederate: Added management commands for OCSP certificates
- PingFederate: Added management commands for PingOne for Enterprise key pairs
- PingFederate: Added management commands for redirect validation settings
- PingFederate: Added management commands for resource-owner-credentials-mappings
- PingFederate: Added management commands for system keys
- PingFederate: Added management commands for the application session policy
- PingFederate: Added management commands for the default authentication policy
- PingFederate: Added management commands for token exchange processor policies
- PingFederate: Added management commands for token processor descriptors
- PingFederate: Added management commands for token processors
- PingFederate: Added support for supplying the new password via --from-file to the administrative accounts change-password and reset-password actions
- PingFederate: Nested license agreement management under the license command

