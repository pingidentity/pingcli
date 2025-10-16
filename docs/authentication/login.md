# `pingcli auth login`

Authenticate the CLI with PingOne using OAuth2 flows.

## Synopsis

Login using one of three supported OAuth2 authentication flows. The CLI will securely store tokens for subsequent API calls.

## Usage
```bash
pingcli auth login [flags]
```

## Flags

### Authentication Method (required - choose one)
- `-d, --device-code` - Use device code flow (recommended for interactive use)
- `-a, --auth-code` - Use authorization code flow (requires browser)
- `-c, --client-credentials` - Use client credentials flow (for automation)

### Global Flags
- `-h, --help` - Help for login command

## Authentication Flows

### Device Code Flow (`-d, --device-code`)
**Recommended for interactive development**

```bash
pingcli auth login --device-code
```

**Requirements:**
- Device code client application configured in PingOne
- Interactive terminal access

**Configuration:**
```bash
pingcli config set service.pingone.authentication.deviceCode.clientID=<client-id>
pingcli config set service.pingone.authentication.deviceCode.environmentID=<env-id>
pingcli config set service.pingone.authentication.deviceCode.scopes="openid,profile"  # optional
```

**Flow:**
1. CLI displays device code and verification URL
2. User visits URL in browser and enters code
3. User authenticates in browser
4. CLI receives and stores tokens

### Authorization Code Flow (`-a, --auth-code`)
**Requires browser on same machine**

```bash
pingcli auth login --auth-code
```

**Requirements:**
- Web application configured in PingOne with redirect URI
- Browser access on local machine

**Configuration:**
```bash
pingcli config set service.pingone.authentication.authCode.clientID=<client-id>
pingcli config set service.pingone.authentication.authCode.environmentID=<env-id>
pingcli config set service.pingone.authentication.authCode.redirectURI=http://localhost:8080/callback
pingcli config set service.pingone.authentication.authCode.scopes="openid,profile"  # optional
```

**Flow:**
1. CLI opens browser to PingOne authorization URL
2. User authenticates in browser
3. Browser redirects to local callback server
4. CLI receives authorization code and exchanges for tokens

### Client Credentials Flow (`-c, --client-credentials`)
**For automation and CI/CD**

```bash
pingcli auth login --client-credentials
```

**Requirements:**
- Worker application configured in PingOne
- Client secret securely managed

**Configuration:**
```bash
pingcli config set service.pingone.authentication.clientCredentials.clientID=<client-id>
pingcli config set service.pingone.authentication.clientCredentials.clientSecret=<client-secret>
pingcli config set service.pingone.authentication.clientCredentials.environmentID=<env-id>
pingcli config set service.pingone.authentication.clientCredentials.scopes="p1:read:*"  # optional
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

# Login
pingcli auth login --device-code
```

### CI/CD Pipeline
```bash
# Set via environment variables (recommended)
export PINGCLI_SERVICE_PINGONE_AUTHENTICATION_CLIENTCREDENTIALS_CLIENTID="$CI_CLIENT_ID"
export PINGCLI_SERVICE_PINGONE_AUTHENTICATION_CLIENTCREDENTIALS_CLIENTSECRET="$CI_CLIENT_SECRET"
export PINGCLI_SERVICE_PINGONE_AUTHENTICATION_CLIENTCREDENTIALS_ENVIRONMENTID="$CI_ENV_ID"

# Login
pingcli auth login --client-credentials
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
**Solution:** Check configuration and try again. Use `pingcli auth logout` to clear any corrupted tokens.

## Security Notes

- Tokens are stored securely in OS credential store
- Device code and auth code flows provide refresh tokens for automatic renewal
- Client credentials flow requires secure secret management
- Use `pingcli auth logout` to clear tokens when switching environments

## See Also
- [Authentication Overview](overview.md)
- [Logout Command](logout.md)
- [Configuration Guide](../tool-configuration/configuration-key.md)