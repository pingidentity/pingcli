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

pingcli uses the [`pingone-go-client`](https://github.com/pingidentity/pingone-go-client) SDK for secure token storage across platforms:

- **macOS**: Keychain Services
- **Windows**: Windows Credential Manager  
- **Linux**: Secret Service API

The SDK provides a consistent `TokenStorage` interface that handles:
- Token serialization and encryption
- Cross-platform credential store access
- Automatic token cleanup on logout

### SDK Integration

Authentication tokens are managed through the SDK's `oauth2.KeychainStorage` implementation:

```go
// Global token storage instance
var tokenStorage = oauth2.NewKeychainStorage("pingcli", "device-code-token")

// Token operations delegate to SDK
func SaveToken(token *oauth2.Token) error {
    return tokenStorage.SaveToken(token)
}
```

This ensures consistent token management across all applications using the `pingone-go-client` SDK.

## See Also
- [Authentication Overview](overview.md)
- [Login Command](login.md)
- [Logout Command](logout.md)