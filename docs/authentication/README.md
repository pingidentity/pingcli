# Authentication Commands

## `pingcli auth`

Main authentication command group for managing CLI authentication with PingOne services.

### Usage
```bash
pingcli auth [command]
```

### Available Commands
- [`login`](login.md) - Authenticate using OAuth2 flows
- [`logout`](logout.md) - Clear stored authentication tokens

### Flags
- `-h, --help` - Help for auth command

### Examples
```bash
# View auth help
pingcli auth --help

# Login with device code flow
pingcli auth login --device-code

# Logout and clear tokens
pingcli auth logout
```

## Quick Start

1. **Configure authentication settings**:
   ```bash
   pingcli config set service.pingone.regionCode=NA
   pingcli config set service.pingone.authentication.deviceCode.clientID=<your-client-id>
   pingcli config set service.pingone.authentication.deviceCode.environmentID=<your-env-id>
   ```

2. **Authenticate**:
   ```bash
   pingcli auth login --device-code
   ```

3. **Use authenticated commands**:
   ```bash
   pingcli request get /environments
   ```

4. **Logout when done**:
   ```bash
   pingcli auth logout
   ```

## Technical Architecture

### Token Storage

pingcli uses a **dual storage system** to ensure tokens are accessible across different environments:

1. **Primary Storage**: Secure platform credential stores (via [`pingone-go-client`](https://github.com/pingidentity/pingone-go-client) SDK)
   - **macOS**: Keychain Services
   - **Windows**: Windows Credential Manager  
   - **Linux**: Secret Service API

2. **Secondary Storage**: Encrypted file-based storage at `~/.pingcli/credentials/`
   - Used when keychain storage fails or is unavailable
   - Automatically created and maintained
   - One file per authentication method (e.g., `device-code-token.json`, `auth-code-token.json`)
   - Provides compatibility with SSH sessions, containers, and CI/CD environments

### Storage Behavior

By default, tokens are stored in **both** locations:
- Keychain storage (if available) - for system-wide secure access
- File storage (always) - for reliability and portability

#### Using the `--use-keychain` Flag

The `--use-keychain` flag controls token retrieval preference:

```bash
# Use keychain-stored token exclusively (fails if unavailable)
pingcli request --use-keychain get /environments

# Default: Try keychain first, fallback to file
pingcli request get /environments
```

**Behavior**:
- `--use-keychain=true`: Only attempts keychain retrieval. Fails if keychain token is missing or inaccessible.
- `--use-keychain=false` (default): Tries keychain first, automatically falls back to file storage if keychain fails.

**Use Cases**:
- `--use-keychain=true`: Enforce keychain security in trusted environments
- Default behavior: Maximum compatibility across environments (SSH, containers, CI/CD)

### SDK Integration

Token storage leverages the SDK's `oauth2.KeychainStorage` implementation alongside local file storage:

```go
// Dual storage approach
var keychainStorage = oauth2.NewKeychainStorage("pingcli", "device-code-token")
var fileStorage = auth.NewFileStorage("~/.pingcli/credentials/device-code-token.json")

// Tokens are stored in both locations
func SaveToken(token *oauth2.Token) error {
    keychainStorage.SaveToken(token)  // Best effort
    return fileStorage.SaveToken(token)  // Always succeeds
}
```

This ensures consistent token management while providing maximum reliability across all environments.

## See Also
- [Authentication Overview](overview.md)
- [Login Command](login.md)
- [Logout Command](logout.md)