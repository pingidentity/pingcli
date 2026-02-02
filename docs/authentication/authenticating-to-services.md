# Authentication to Service

This guide covers authenticating Ping CLI to PingOne and clearing credentials when you are done.

## Quick start

1. Configure settings:

  ```bash
  pingcli config set service.pingone.regionCode=NA
  pingcli config set service.pingone.authentication.environmentID=<your-env-id>
  pingcli config set service.pingone.authentication.deviceCode.clientID=<your-client-id>
  ```

2. Login:

  ```bash
  pingcli login --device-code
  ```

3. Use authenticated commands:

  ```bash
  pingcli request get /environments
  ```

4. Logout when done:

  ```bash
  pingcli logout
  ```

## Prerequisites: Configure a PingOne Application

Before running `pingcli login`, configure a PingOne application for the grant type you intend to use. PingCLI supports:

- `client_credentials` (recommended for service/automation; legacy `worker` maps to this)
- `authorization_code` (interactive browser login)
- `device_code` (interactive terminal login on headless environments)

See the PingOne Platform API documentation:

- Application operations: <https://apidocs.pingidentity.com/pingone/platform/v1/api/#application-operations>

### Client credentials (Worker)

Configure your PingOne application to support `client_credentials`:

- Enable grant type: `client_credentials`
- Create Client ID and Client Secret

Collect for PingCLI:

- Environment ID
- Client ID
- Client Secret

Notes:

- Auth type `worker` is applied as `client_credentials` under the hood
- No refresh token is issued for `client_credentials`
- If a previous version of `pingcli` was used, `pingone.authentication.type` may be set to `worker`. `pingcli login` will interpret this as an intention to migrate away from the deprecated type and favor `client_credentials`. To use `device_code` or `authorization_code` instead, update your configuration or pass the appropriate login flag.

> Deprecation Notice: The `worker` authentication type is deprecated and will be removed in a future release. Use `client_credentials` instead.

### Authorization code

Configure your PingOne application to support `authorization_code`:

- Enable Response Type: `Code`
- Enable Grant Type: `Authorization Code`
- Select PKCE Enforcement: `OPTIONAL` (PKCE will be used by pingcli by default)
- Optionally enable Refresh Token
- Set redirect URI(s). PingCLI defaults to `http://127.0.0.1:7464/callback` with path `/callback` and port `7464` (customizable)

Collect for PingCLI:

- Environment ID
- Client ID
- Redirect URI path (e.g. `/callback`)
- Redirect URI port (e.g. `7464`)

### Device code

Configure your PingOne application to support device code:

- Enable grant type: `Device Authorization`
- Optionally enable Refresh Token

Collect for PingCLI:

- Environment ID
- Client ID

### Region selection

PingCLI prompts for your PingOne region and uses it to route API requests. Supported codes: `AP`, `AU`, `CA`, `EU`, `NA`, `SG`.

## Login (`pingcli login`)

Login using one of three supported OAuth2 flows. The CLI will securely store tokens for subsequent API calls.

### Usage

```bash
pingcli login [flags]
```

### Flags

#### Authentication Method (required - choose one)

- `-d, --device-code` - Use device code flow (recommended for interactive use)
- `-a, --auth-code` - Use authorization code flow (requires browser)
- `-c, --client-credentials` - Use client credentials flow (for automation)

#### Provider Selection

- `-p, --provider` - Target authentication provider (default: `pingone`)

#### Storage Options

- `--storage-type` - Auth token storage type (default: `secure_local`)
  - `secure_local` - Use OS keychain (default)
  - `file_system`  - Store tokens at `~/.pingcli/credentials`
  - `none`         - Do not persist tokens

### Authentication flows

### Interactive authentication

If you run `pingcli login` without specifying an authentication method flag (and no type is set in configuration), Ping CLI prompts you to select a method:

```bash
$ pingcli login
? Select authentication method:
  ▸ device_code (configured)
    authorization_code (configured)
    client_credentials (not configured)
```

#### Device Code Flow (`--device-code`)

```bash
pingcli login --device-code
```

Configuration:

```bash
pingcli config set service.pingone.authentication.environmentID=<env-id>
pingcli config set service.pingone.authentication.deviceCode.clientID=<client-id>
```

#### Authorization Code Flow (`--auth-code`)

```bash
pingcli login --auth-code
```

Configuration:

```bash
pingcli config set service.pingone.authentication.environmentID=<env-id>
pingcli config set service.pingone.authentication.authorizationCode.clientID=<client-id>
pingcli config set service.pingone.authentication.authorizationCode.redirectURIPath="/callback"
pingcli config set service.pingone.authentication.authorizationCode.redirectURIPort="7464"
```

#### Client Credentials Flow (`--client-credentials`)

```bash
pingcli login --client-credentials
```

Configuration:

```bash
pingcli config set service.pingone.authentication.environmentID=<env-id>
pingcli config set service.pingone.authentication.clientCredentials.clientID=<client-id>
pingcli config set service.pingone.authentication.clientCredentials.clientSecret=<client-secret>
```

## Token storage

Ping CLI offers a number of storage options:

- `secure_local`: OS credential stores (Keychain/Credential Manager)
- `file_system`: File storage at `~/.pingcli/credentials`
- `none`: Tokens are not stored

### Recommended: keychain with automatic fallback

Default behavior is `secure_local`. Ping CLI attempts to store credentials in the OS credential store.

If keychain storage fails (unavailable, permission denied, etc.) or `--storage-type=file_system` is selected, Ping CLI uses file storage at `~/.pingcli/credentials`.

#### Where tokens are stored

- **OS credential stores** (when `--storage-type=secure_local`):
  - macOS: Keychain Services
  - Windows: Windows Credential Manager
  - Linux: Secret Service API (GNOME Keyring/KDE KWallet)
- **File storage** (when `--storage-type=file_system`, or as fallback): `~/.pingcli/credentials/`

#### Token retrieval

When keychain is enabled, Ping CLI attempts keychain first and falls back to file storage if keychain operations fail. When `--storage-type=file_system` is selected, Ping CLI uses file storage exclusively.

## Logout (`pingcli logout`)

Clear stored authentication tokens from both keychain and file storage.

### Usage

```bash
pingcli logout [flags]
```

### Flags

#### Authentication Method (optional)

- `-d, --device-code` - Clear only device code tokens
- `-a, --auth-code` - Clear only authorization code tokens
- `-c, --client-credentials` - Clear only client credentials tokens

If no flag is provided, clears tokens for all authentication methods.

## What gets cleared

### Tokens

- Access tokens
- Refresh tokens
- Token metadata (expiry)

### Storage locations

Logout clears tokens from both the OS credential store (keychain/credential manager) and file storage under `~/.pingcli/credentials`.

## Verification

After logout, verify tokens are cleared:

```bash
pingcli request get /environments
```

Expected response:

```
Error: no valid authentication token found. Please run 'pingcli login --device-code' to authenticate
```

## Manual token removal

If logout fails, manually remove tokens from both storage locations.

### Keychain/Credential store

macOS:

```bash
security delete-generic-password -s "pingcli" -a "<env-id>_<client-id>_device_code"
```

Windows:

```cmd
cmdkey /delete:LegacyGeneric:target=pingcli
```

Linux (GNOME):

```bash
secret-tool clear service pingcli
```

### File storage

```bash
rm -rf ~/.pingcli/credentials
```

## See also

- [Authentication README](#authentication-commands)

# Authentication Commands

## Authentication

Main authentication commands for managing CLI authentication with PingOne services.

### Available Commands
- [`pingcli login`](login.md) - Authenticate using OAuth2 flows
- [`pingcli logout`](logout.md) - Clear stored authentication tokens

### Interactive Authentication

When you run `pingcli login` without specifying an authentication method flag (or no type is set in the configuration), the CLI will prompt you to select from available methods:

```bash
$ pingcli login
? Select authentication method:
  ▸ device_code (configured)
   authorization_code (configured)
   client_credentials (not configured)
```

This interactive mode helps you choose the appropriate authentication flow for your use case without needing to remember the exact flag names. The status indicator shows whether each method has the required configuration settings:
- **(configured)** - All required settings (client ID, environment ID, etc.) are present in your config
- **(not configured)** - Missing one or more required configuration values

## Quick Start

1. **Configure authentication settings**:
  ```bash
  pingcli config set service.pingone.regionCode=NA
  pingcli config set service.pingone.authentication.deviceCode.clientID=<your-client-id>
  pingcli config set service.pingone.authentication.deviceCode.environmentID=<your-env-id>
  ```

2. **Authenticate**:
  ```bash
  pingcli login --device-code
  ```

3. **Use authenticated commands**:
  ```bash
  pingcli request get /environments
  ```

4. **Logout when done**:
  ```bash
  pingcli logout
  ```

## Technical Architecture

### Token Storage

pingcli uses a **dual storage system** to ensure tokens are accessible across different environments:

1. **Primary Storage**: Secure platform credential stores (via [`pingone-go-client`](https://github.com/pingidentity/pingone-go-client) SDK)
  - **macOS**: Keychain Services
  - **Windows**: Windows Credential Manager  
  - **Linux**: Secret Service API

2. **Secondary Storage**: File-based storage at `~/.pingcli/credentials/`
  - Automatically created and maintained
  - One file per grant type (e.g., `<env-id>_<client-id>_device_code.json`)
  - Provides compatibility with SSH sessions, containers, and CI/CD environments

### Storage Behavior

**Default: Dual Storage with Automatic Fallback**

By default (`--file-storage=false`), tokens are stored in **both** locations simultaneously:
- Keychain storage (primary) - for system-wide secure access
- File storage (backup) - for reliability and portability

```bash
# Default: Saves to both keychain and file
pingcli login --device-code
# Output: Successfully logged in using device_code. 
#         Credentials saved to keychain and file storage for profile 'default'.
```

**Fallback Protection:**
If keychain storage fails (unavailable, permission issues, etc.), the system automatically falls back to file storage only:
```bash
# Keychain unavailable - automatically uses file storage
pingcli login --device-code
# Output: Successfully logged in using device_code. 
#         Credentials saved to file storage for profile 'default'.
```

**File-Only Mode**

Use the `--file-storage` flag to explicitly skip keychain and use only file storage:

```bash
# Explicitly use file storage only (skip keychain entirely)
pingcli login --device-code --file-storage
# Output: Successfully logged in using device_code. 
#         Credentials saved to file storage for profile 'default'.
```

**When to use `--file-storage`:**
- SSH sessions where keychain access is unavailable
- Containers and Docker environments
- CI/CD pipelines
- Debugging keychain issues
- When you want to ensure file-only storage (no keychain attempts)

**Token Retrieval:**
- Default: Attempts keychain first, automatically falls back to file storage if keychain fails
- File-only mode (`--file-storage=true`): Uses file storage exclusively

## See Also
- [Authentication To Service](authentication-to-service.md)

