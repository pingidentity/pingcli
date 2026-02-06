# Authenticating to Services

This document covers the available options for authenticating to services in `pingcli`: login flows, logout behavior, token storage, configuration, and practical token retrieval across environments.

## Interactive Login

First, ensure the [Prerequisites](#prerequisites-configure-a-pingone-application) are met.
Then, run `pingcli login` to interact with Ping CLI prompts to configure authentication parameters and launch a login flow:

```bash
$ pingcli login
No authentication methods configured. Let's set one up!
Use the arrow keys to navigate: ↓ ↑ → ←
? Select authorization grant type for this profile:
  ▸ device_code
    authorization_code
    client_credentials
```

## Prerequisites: Configure a PingOne Application

Before authenticating to a Ping service from Ping CLI, ensure the platform is prepared with a configured application. For the PingOne platform, Ping CLI supports:

- `client_credentials` (recommended for service/automation; legacy pingcli `worker` authentication maps to this)
- `authorization_code` (interactive browser login)
- `device_code` (interactive terminal login on headless environments)

For guidance on which authorization flow to use see [Authorization Flow Recommendations](#auhtorization-flow-recommendations).

For guidance on how to configure applications in PingOne see [PingOne Platform Application documentation](https://docs.pingidentity.com/pingone/applications/p1_applications_add_applications.html)

### Client credentials (Worker)

Configure your PingOne application to support `client_credentials`:

- Enable grant type: `client_credentials`
- Create Client ID and Client Secret

Collect for Ping CLI:

- Environment ID (the environment containing the application)
- Client ID
- Client Secret

Ping CLI notes:

- No refresh token is issued for `client_credentials`
- If a previous version of `pingcli` was used, the `pingone.authentication.type` may be set to `worker`. The `pingcli` login command will interpret this as an intention to migrate away from the deprecated type and favor `client_credentials`. If another authentication type is preferred (e.g. `device_code` or `authorization_code`) edit the configuration file or pass an appropriate flag to the `login` command.

> Deprecation Notice: The `worker` authentication type is deprecated and will be removed in a future release. Use `client_credentials` instead.

### Authorization code

Configure your PingOne application to support `authorization_code`:

- Enable Response Type: `Code`
- Enable Grant Type: `Authorization Code`
- Select PKCE Enforcement: `OPTIONAL` (PKCE will be used by pingcli by default)
- Optionally Enable Refresh Token
- Set redirect URI(s). By default, Ping CLI will receive the redirect at: `http://127.0.0.1:7464/callback`. `/callback` and port `7464` are customizable in CLI

Collect for Ping CLI:

- Environment ID
- Client ID
- Redirect URI path (e.g. `/callback`)
- Redirect URI port (e.g. `7464`)

### Device code

Configure your PingOne application to support `device_code`:

- Enable grant type: `Device Authorization`

Collect for Ping CLI:

- Environment ID
- Client ID

## Login Configuration

The command `pingcli login` with allow CLI to access a service using one of three supported flows and launch the interactive prompt if configuration is not set. By default, the CLI will securely store tokens for subsequent API calls.

Ping CLI reads settings from the active profile inside `$HOME/.pingcli/config.yaml` by default.

You can configure values either by:

- Using `pingcli config set` command
- Editing the YAML config file directly.
- Using `pingcli login` when no previous configuration exists

```bash
pingcli login [flags]
```

### Login Flags (required - choose one)

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

### Common Configuration

These settings are used by all PingOne grant types:

```bash
pingcli config set service.pingOne.regionCode=NA
pingcli config set service.pingOne.authentication.environmentID=<PINGONE_ENVIRONMENT_ID>
pingcli config set service.pingOne.authentication.type=<device_code|authorization_code|client_credentials>
```

### Device Code Flow (`--device-code`)

```bash
pingcli login --device-code
```

Configuration:

```bash
pingcli config set service.pingone.authentication.environmentID=<env-id>
pingcli config set service.pingone.authentication.deviceCode.clientID=<client-id>
```

### Authorization Code Flow (`--auth-code`)

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

### Client Credentials Flow (`--client-credentials`)

```bash
pingcli login --client-credentials
```

Configuration:

```bash
pingcli config set service.pingone.authentication.environmentID=<env-id>
pingcli config set service.pingone.authentication.clientCredentials.clientID=<client-id>
pingcli config set service.pingone.authentication.clientCredentials.clientSecret=<client-secret>
```

### Region selection

During interactive configuration Ping CLI prompts for your PingOne region and uses it to route API requests. Supported codes: `AP`, `AU`, `CA`, `EU`, `NA`, `SG`.

```bash
pingcli config set service.pingOne.regionCode=<region-code>
```

## Token Management

Upon successful authentication via `pingcli login` the CLI will optionally store, retrieve and refresh access tokens as needed for subsequent commands.

Ping CLI offers a number of storage options:

- `secure_local`: OS credential stores (Keychain/Credential Manager)
- `file_system`: File storage at `~/.pingcli/credentials`
- `none`: Tokens are not stored

### Recommended Token Storage

Default behavior is `secure_local`. Ping CLI attempts to store credentials in the OS credential store.

If keychain storage fails (unavailable, permission denied, etc.) or `--storage-type=file_system` is selected, Ping CLI uses file storage in the `~/.pingcli/credentials` directory.

## Logout Command

Clear stored authentication tokens from both keychain and file storage.

```bash
pingcli logout [flags]
```

### Logout Flags (optional)

- `-d, --device-code` - Clear only device code tokens
- `-a, --auth-code` - Clear only authorization code tokens
- `-c, --client-credentials` - Clear only client credentials tokens

If no flag is provided, clears tokens for all authentication methods.

### What gets cleared

### Tokens

- Access tokens
- Refresh tokens
- Token metadata (expiry)

### Storage locations

Logout clears tokens from both the OS credential store (keychain/credential manager) and file storage under `~/.pingcli/credentials`.

### Verification

After logout, verify tokens are cleared:

```bash
pingcli request get /environments
```

Expected response:

```text
Error: no valid authentication token found. Please run 'pingcli login --device-code' to authenticate
```

### Manual token removal

If logout fails, users can manually remove tokens from both storage locations.

#### Keychain/Credential store**

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

#### File storage

```bash
rm -rf ~/.pingcli/credentials
```

## Auhtorization Flow Recommendations

This section provides guidance on when to select each authorization type.

### `client_credentials`

Use when:

- Running in CI/CD or other automation without direct human access during authentication
- You want non-interactive auth (no browser, no prompts)

Notes:

- This is the best choice for “machine to machine” usage.
- Legacy `worker` (if present in old configs) maps to `client_credentials`.

### `authorization_code`

Use when:

- You’re working interactively and can complete a browser-based login on the same machine
- You need a user-context token
- You need to allow multiple users to use the same configuration.

Ping CLI requires a redirect URI (path + port) that must be allowed by the PingOne application.

### `device_code`

Use when:

- You’re working interactively but don’t want/need a local redirect listener
- You’re on a headless system (SSH, remote box) and can open a browser elsewhere to complete the verification
