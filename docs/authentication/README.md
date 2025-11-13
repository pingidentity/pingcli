# Authentication Commands

## Authentication

Main authentication commands for managing CLI authentication with PingOne services.

### Available Commands
- [`pingcli login`](login.md) - Authenticate using OAuth2 flows
- [`pingcli logout`](logout.md) - Clear stored authentication tokens

### Examples
```bash
# Interactive login - prompts for authentication method (if no type is configured)
pingcli login

# Login with specific authentication method
pingcli login --device-code
pingcli login --authorization-code
pingcli login --client-credentials

# Logout and clear tokens
pingcli logout
```

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
   - One file per authentication method (e.g., `<env-id>_<client-id>_device_code.json`)
   - Provides compatibility with SSH sessions, containers, and CI/CD environments

### Storage Behavior

**Default: Dual Storage with Automatic Fallback**

By default (`--file-storage=false`), tokens are stored in **both** locations simultaneously:
- Keychain storage (primary) - for system-wide secure access
- File storage (backup) - for reliability and portability

```bash
# Default: Saves to both keychain and file
pingcli login --device-code
# Output: Successfully logged in using device_code authentication. 
#         Credentials saved to keychain and file storage for profile 'default'.
```

**Fallback Protection:**
If keychain storage fails (unavailable, permission issues, etc.), the system automatically falls back to file storage only:
```bash
# Keychain unavailable - automatically uses file storage
pingcli login --device-code
# Output: Successfully logged in using device_code authentication. 
#         Credentials saved to file storage for profile 'default'.
```

**File-Only Mode**

Use the `--file-storage` flag to explicitly skip keychain and use only file storage:

```bash
# Explicitly use file storage only (skip keychain entirely)
pingcli login --device-code --file-storage
# Output: Successfully logged in using device_code authentication. 
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

### SDK Integration

Token storage leverages the SDK's `oauth2.KeychainStorage` implementation alongside local file storage:

```go
// Dual storage approach - saves to both locations
func SaveTokenForMethod(token *oauth2.Token, authMethod string) (StorageLocation, error) {
    location := StorageLocation{}
    
    // Try keychain storage
    if !fileStorageOnly() {
        keychainStorage := oauth2.NewKeychainStorage("pingcli", authMethod)
        if err := keychainStorage.SaveToken(token); err == nil {
            location.Keychain = true
        }
    }
    
    // Always save to file storage as backup
    if err := saveTokenToFile(token, authMethod); err == nil {
        location.File = true
    }
    
    return location, nil
}
```

This ensures consistent token management while providing maximum reliability across all environments.

## See Also
- [Authentication Overview](overview.md)
- [Login Command](login.md)
- [Logout Command](logout.md)