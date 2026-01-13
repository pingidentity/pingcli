# `pingcli login`

Authenticate the CLI with PingOne using OAuth2 flows.

## Prerequisites: Configure a PingOne Application

Before running `pingcli login`, configure a PingOne application for the grant type you intend to use. PingCLI supports:

- client_credentials (recommended for service/automation; legacy `worker` maps to this)
- authorization_code (interactive browser login)
- device_code (interactive terminal login on headless environments)

See the PingOne Platform API documentation to manage applications:
- Application operations: https://apidocs.pingidentity.com/pingone/platform/v1/api/#application-operations

### Client credentials (Worker)

Configure your PingOne application to support `client_credentials`:
- Enable grant type: `client_credentials`
- Create Client ID and Client Secret

Collect for PingCLI:
- Environment ID (the environment containing the application)
- Client ID
- Client Secret

PingCLI notes:
- Auth type `worker` is applied as `client_credentials` under the hood
- No refresh token is issued for `client_credentials`

### Authorization code

Configure your PingOne application to support `authorization_code`:
- Enable grant type: `authorization_code`
- Set redirect URI(s). PingCLI defaults to `http://localhost:<port><path>` with path `/callback` and port `7464` (customizable in CLI)
- Create Client ID

Collect for PingCLI:
- Environment ID
- Client ID
- Redirect URI path (e.g., `/callback`)
- Redirect URI port (e.g., `7464`)

### Device code

Configure your PingOne application to support device code:
- Enable grant type: `urn:ietf:params:oauth:grant-type:device_code`
- Create Client ID

Collect for PingCLI:
- Environment ID
- Client ID

### Region selection

PingCLI prompts for your PingOne region and uses it to route API requests. Supported codes: `AP`, `AU`, `CA`, `EU`, `NA`, `SG`.

## Synopsis

Login using one of three supported OAuth2 authentication flows. The CLI will securely store tokens for subsequent API calls.

## Usage
```bash
pingcli login [flags]
```

## Flags

### Authentication Method (required - choose one)
- `-d, --device-code` - Use device code flow (recommended for interactive use)
- `-a, --auth-code` - Use authorization code flow (requires browser)
- `-c, --client-credentials` - Use client credentials flow (for automation)

### Provider Selection
- `-p, --provider` - Target authentication provider (default: `pingone`)
  - Currently only `pingone` is supported
  - Future versions will support multiple providers

### Storage Options
- `--file-storage` - Use only file storage (skip keychain).

### Global Flags
- `-h, --help` - Help for login command

## Authentication Flows

### Device Code Flow (`-d, --device-code`)
**Recommended for interactive development**

```bash
pingcli login --device-code
```

**Requirements:**
- Device code client application configured in PingOne
- Interactive terminal access

**Configuration:**
```bash
pingcli config set service.pingone.authentication.deviceCode.clientID=<client-id>
pingcli config set service.pingone.authentication.deviceCode.environmentID=<env-id>
```

**Flow:**
1. CLI displays device code and verification URL
2. User visits URL in browser and enters code
3. User authenticates in browser
4. CLI receives and stores tokens

### Authorization Code Flow (`-a, --auth-code`)
**Requires browser on same machine**

```bash
pingcli login --auth-code
```

**Requirements:**
- Web application configured in PingOne with redirect URI
- Browser access on local machine

**Configuration:**
```bash
pingcli config set service.pingone.authentication.authCode.clientID=<client-id>
pingcli config set service.pingone.authentication.authCode.environmentID=<env-id>
pingcli config set service.pingone.authentication.authCode.redirectURI=http://localhost:8080/callback
```

**Flow:**
1. CLI opens browser to PingOne authorization URL
2. User authenticates in browser
3. Browser redirects to local callback server
4. CLI receives authorization code and exchanges for tokens

### Client Credentials Flow (`-c, --client-credentials`)
**For automation and CI/CD**

```bash
pingcli login --client-credentials
```

**Requirements:**
- Worker application configured in PingOne
- Client secret securely managed

**Configuration:**
```bash
pingcli config set service.pingone.authentication.clientCredentials.clientID=<client-id>
pingcli config set service.pingone.authentication.clientCredentials.clientSecret=<client-secret>
pingcli config set service.pingone.authentication.clientCredentials.environmentID=<env-id>
```

**Flow:**
1. CLI sends client credentials directly to token endpoint
2. Receives access token (no refresh token)
3. Stores token for API calls

## Examples

### Interactive Development
```bash
# Configure device code settings
pingcli config set service.pingone.regionCode=NA
pingcli config set service.pingone.authentication.deviceCode.clientID=abc123
pingcli config set service.pingone.authentication.deviceCode.environmentID=env456

# Login (--provider defaults to pingone)
pingcli login --device-code

# Explicitly specify provider
pingcli login --device-code --provider pingone
```

### CI/CD Pipeline
```bash
# Set via environment variables
export PINGCLI_SERVICE_PINGONE_AUTHENTICATION_CLIENTCREDENTIALS_CLIENTID="$CI_CLIENT_ID"
export PINGCLI_SERVICE_PINGONE_AUTHENTICATION_CLIENTCREDENTIALS_CLIENTSECRET="$CI_CLIENT_SECRET"
export PINGCLI_SERVICE_PINGONE_AUTHENTICATION_CLIENTCREDENTIALS_ENVIRONMENTID="$CI_ENV_ID"

# Login with file-only storage (skip keychain)
pingcli login --client-credentials --file-storage
```

## Error Handling

### Common Errors

**No authentication method specified:**
```
Error: please specify an authentication method: --auth-code, --client-credentials, or --device-code
```
**Solution:** Add one of the required flags.

**Multiple authentication methods:**
```
Error: please specify only one authentication method
```
**Solution:** Use only one authentication flag.

**Missing configuration:**
```
Error: device code client ID is not configured. Please run 'pingcli config set service.pingone.authentication.deviceCode.clientID=<your-client-id>'
```
**Solution:** Configure required settings before authentication.

**Authentication failed:**
```
Error: failed to get valid token (may need to re-authenticate)
```
**Solution:** Check configuration and try again. Use `pingcli logout` to clear any corrupted tokens.

## Token Storage

pingcli uses a **dual storage system** for maximum reliability:

1. **Primary**: OS credential stores (Keychain/Credential Manager/Secret Service)
2. **Secondary**: Encrypted file storage at `~/.pingcli/credentials`

### Storage Behavior

**Default:**
Tokens are automatically stored in **both** locations:
```bash
pingcli login --device-code
# Output: Successfully logged in using device_code. 
#         Credentials saved to keychain and file storage for profile 'default'.
```

**Automatic Fallback:**
If keychain fails (unavailable, permission denied, etc.), automatically falls back to file storage:
```bash
# Keychain unavailable - uses file storage instead
pingcli login --device-code
# Output: Successfully logged in using device_code. 
#         Credentials saved to file storage for profile 'default'.
```

**Benefits:**
- Keychain provides system-wide secure access when available
- File storage ensures tokens are never lost
- Automatic fallback handles all edge cases
- Zero user intervention required

**File-Only Mode:**
Use `--file-storage` flag to explicitly skip keychain:
```bash
pingcli login --device-code --file-storage
# Output: Successfully logged in using device_code. 
#         Credentials saved to file storage for profile 'default'.
```

**When to use `--file-storage`:**
- SSH sessions where keychain is unavailable
- Docker containers
- CI/CD pipelines
- Systems without keychain support
- When you want to guarantee file-only storage

### Token Retrieval

When loading tokens, pingcli automatically:
1. Tries keychain first (unless `--file-storage` was used during login)
2. Falls back to file storage if keychain fails
3. Returns error only if both fail

This ensures maximum reliability across all environments.

## Security Notes

- Tokens are stored in both OS credential store and encrypted file by default
- Use `--file-storage` flag in environments without keychain access
- Device code and auth code flows provide refresh tokens for automatic renewal
- Client credentials flow requires secure secret management
- Use `pingcli logout` to clear tokens from both locations when switching environments

## See Also
- [Authentication Overview](overview.md)
- [Logout Command](logout.md)
- [Configuration Guide](../tool-configuration/configuration-key.md)