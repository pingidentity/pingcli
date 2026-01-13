# `pingcli login`

Authenticate the CLI with PingOne using OAuth2 flows.

## Prerequisites: Configure a PingOne Application

Before running `pingcli login`, configure a PingOne application for the grant type you intend to use. PingCLI supports:

- client_credentials (recommended for service/automation; legacy `worker` maps to this)
- authorization_code (interactive browser login)
- device_code (interactive terminal login on headless environments)

See the PingOne Platform API documentation to manage applications:

- Application operations: <https://apidocs.pingidentity.com/pingone/platform/v1/api/#application-operations>

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

> Deprecation Notice: The `worker` authentication type is deprecated and will be removed in a future release. Use `client_credentials` instead.

### Authorization code

Configure your PingOne application to support `authorization_code`:

- Enable Response Type: `Code`
- Enable Grant Type: `Authorization Code`
- Select PKCE Enforcement: `OPTIONAL` (PKCE will be used by pingcli by default)
- Optionally Enable Refresh Token
- Set redirect URI(s). PingCLI defaults to `http://127.0.0.1:7464/callback` with path `/callback` and port `7464` (customizable in CLI)

Collect for PingCLI:

- Environment ID
- Client ID
- Redirect URI path (e.g. `/callback`)
- Redirect URI port (e.g. `7464`)

### Device code

Configure your PingOne application to support device code:

- Enable grant type: `Device Authorization`
- Optionally Enable Refresh Token

Collect for PingCLI:

- Environment ID
- Client ID

### Region selection

PingCLI prompts for your PingOne region and uses it to route API requests. Supported codes: `AP`, `AU`, `CA`, `EU`, `NA`, `SG`.

## Synopsis

Login using one of three supported OAuth2 flows. The CLI will securely store tokens for subsequent API calls. 

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

- `--storage` - Auth token storage type (default: secure_local)
  - `secure_local`  - Use OS keychain (default)
  - `file_system`   - Store tokens on a file in ~/.pingcli/credentials
  - `none`          - Do not persist tokens

### Global Flags

- `-h, --help` - Help for login command

## Authentication Flows

### Device Code Flow (`-d, --device-code`)

**Recommended for interactive development and environments without internet browser access.**

```bash
pingcli login --device-code
```

**Requirements:**

- Device Code client application configured in PingOne (see [Prerequisites](#prerequisites-configure-a-pingone-application))
- Interactive terminal access

**Configuration:**

```bash
pingcli config set service.pingone.authentication.environmentID=<env-id>
pingcli config set service.pingone.authentication.deviceCode.clientID=<client-id>
```

**Flow:**

1. CLI displays device code and verification URL
2. User visits URL in browser and enters code
3. User authenticates in browser
4. CLI receives and stores tokens. CLI will use access token that has roles associated with authenticated user.

### Authorization Code Flow (`-a, --auth-code`)

**Requires browser on same machine. Recommended for interactive development.**

```bash
pingcli login --auth-code
```

**Requirements:**

- Authorization Code client application configured in PingOne (see [Prerequisites](#prerequisites-configure-a-pingone-application))
- Interactive terminal access
- Browser access on local machine

**Configuration:**

```bash
pingcli config set service.pingone.authentication.environmentID=<env-id>
pingcli config set service.pingone.authentication.authorizationCode.clientID=<client-id>
pingcli config set service.pingone.authentication.authorizationCode.redirectURIPath="/callback"
pingcli config set service.pingone.authentication.authorizationCode.redirectURIPort="7464"
```

**Flow:**

1. CLI opens browser to PingOne authorization URL
2. User authenticates in browser, and authorizes client.
3. Browser redirects to local callback server
4. CLI receives authorization code and exchanges for tokens. CLI will use access token that has roles associated with authenticated user.

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
pingcli config set service.pingone.authentication.environmentID=<env-id>
pingcli config set service.pingone.authentication.clientCredentials.clientID=<client-id>
pingcli config set service.pingone.authentication.clientCredentials.clientSecret=<client-secret>
```

**Flow:**

1. CLI sends client credentials directly to token endpoint
2. Receives access token (no refresh token)
3. Stores token for API calls. CLI will use access token that has roles associated with client application.

## Token Storage

Ping CLI offers a number of storage options:

- `secure_local`: OS credential stores (Keychain/Credential Manager)
- `file_system`: File storage at `~/.pingcli/credentials`
- `none`: Tokens are not stored

### Storage Behavior

#### Recommended - Keychain Storage

**Default Behavior**

This describes the recommended behavior for Ping CLI token storage, which is enabled by default.

`login.storage.type` is set to `secure_local` for a profile when `pingcli login` is run. With this option Ping CLI will first look to store credentials in the OS credential store.

```bash
pingcli login --device-code
# Output: Successfully logged in using device_code. 
#         Credentials saved to keychain for profile 'default'.
```

**Automatic Fallback:**

If `secure_local` fails (unavailable, permission denied, etc.) or `login.storage.type` is set to `file_system` Ping CLI uses file storage at `~/.pingcli/credentials`:

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

#### Alternatives

##### File System Storage

Use `--storage=file_system` flag or `pingcli config set login.storage.type="file_system"` to explicitly skip keychain:

```bash
pingcli login --device-code --storage=file_system
# Output: Successfully logged in using device_code. 
#         Credentials saved to file storage for profile 'default'.
```

**When to use `--storage=file_system`:**

- SSH sessions where keychain is unavailable
- Systems without keychain support
- When you want to guarantee file-only storage

##### No Storage

Use `--storage=none` flag or `pingcli config set login.storage.type="none"` to explicitly skip token storage. This means, a new authentication will be run for each Ping CLI command instantiation.

```bash
pingcli login --device-code --storage=none
# Output: Successfully logged in using device_code. 
```

**When to use `--storage=none`:**

- CI/CD pipelines where human interaction is unavailable and authorization relies on client credentials.
- When you want to guarantee tokens are not stored on the host machine.

## Examples

### Interactive Development

```bash
# Configure device code settings
pingcli config set service.pingone.regionCode=NA
pingcli config set service.pingone.authentication.environmentID=abcd1234-ac12-ab12-ab12-abcdef123456
pingcli config set service.pingone.authentication.deviceCode.clientID=abcd1234-ac12-ab12-ab12-abcdef123456

# Login (--provider defaults to pingone)
pingcli login --device-code

# Explicitly specify provider
pingcli login --device-code --provider pingone
```

### CI/CD Pipeline

```bash
# Set via environment variables
export PINGCLI_PINGONE_CLIENT_CREDENTIALS_CLIENT_ID="$CI_CLIENT_ID"
export PINGCLI_PINGONE_CLIENT_CREDENTIALS_CLIENT_SECRET="$CI_CLIENT_SECRET"
export PINGCLI_PINGONE_ENVIRONMENT_ID="$CI_ENV_ID"

# Login with file-only storage (skip keychain)
pingcli login --client-credentials --file-storage
```

## See Also

- [Authentication Overview](overview.md)
- [Logout Command](logout.md)
- [Configuration Guide](../tool-configuration/configuration-key.md)
