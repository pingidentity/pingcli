# `pingcli logout`

Clear stored authentication tokens from both keychain and file storage.

## Synopsis

Logout removes all stored authentication tokens from both the OS credential store and file storage. Use this command when switching between environments, ending sessions, or troubleshooting authentication issues.

## Usage
```bash
pingcli logout [flags]
```

## Flags

### Authentication Method (optional)
- `-d, --device-code` - Clear only device code tokens
- `-a, --auth-code` - Clear only authorization code tokens  
- `-c, --client-credentials` - Clear only client credentials tokens

If no flag is provided, clears tokens for **all** authentication methods.

### Global Flags
- `-h, --help` - Help for logout command

## What Gets Cleared

### Tokens
- Access tokens
- Refresh tokens
- Token metadata (expiry, creation time)

### Storage Locations (Both Cleared)
1. **OS Credential Stores:**
   - **macOS**: Keychain Services
   - **Windows**: Windows Credential Manager  
   - **Linux**: Secret Service API (GNOME Keyring/KDE KWallet)

2. **File Storage:**
   - `~/.pingcli/credentials/<env-id>_<client-id>_<method>.json` - token files

### Cache
- PingOne API client cache
- Cached authentication state

## Examples

### Clear All Tokens (Default)
```bash
pingcli logout
```
**Output:**
```
Successfully logged out from all methods. All credentials cleared from storage for profile 'default'.
```

### Clear Specific Authentication Method
```bash
# Clear only device code tokens
pingcli logout --device-code
```
**Output:**
```
Successfully logged out from device_code. Credentials cleared from keychain and file storage for profile 'default'.
```

```bash
# Clear only client credentials tokens
pingcli logout --client-credentials
```
**Output:**
```
Successfully logged out from client_credentials. Credentials cleared from keychain and file storage for profile 'default'.
```

### Logout in Automation
```bash
#!/bin/bash
# CI/CD cleanup script
pingcli logout --client-credentials
echo "Authentication cleanup complete"
```

## When to Use Logout

### Development
- **Environment switching**: Before authenticating to a different PingOne environment
- **End of session**: When finished working with sensitive data
- **Troubleshooting**: To clear corrupted or invalid tokens

### CI/CD
- **Pipeline cleanup**: At the end of automated workflows
- **Security practice**: Ensure no tokens persist in build environments
- **Error recovery**: Clear state when authentication fails

### Security
- **Shared machines**: Always logout on shared development machines
- **Before rotation**: Clear old tokens before updating credentials
- **Incident response**: Immediately revoke access if credentials compromised

## Verification

After logout, verify tokens are cleared:

```bash
# This should prompt for authentication
pingcli request get /environments
```

Expected response:
```
Error: no valid authentication token found. Please run 'pingcli login --device-code' to authenticate
```

## Manual Token Removal

If logout fails, manually remove tokens from both storage locations:

### Keychain/Credential Store

**macOS:**
```bash
# Command line (replace with your specific key)
security delete-generic-password -s "pingcli" -a "<env-id>_<client-id>_device_code"

# GUI: Keychain Access → search "pingcli" → delete entry
```

**Windows:**
```cmd
# Command line
cmdkey /delete:LegacyGeneric:target=pingcli

# GUI: Control Panel → Credential Manager → remove pingcli entry
```

**Linux:**
```bash
# GNOME
secret-tool clear service pingcli

# GUI: seahorse → search "pingcli" → delete
```

### File Storage

**All Platforms:**
```bash
# Remove all token files
rm -rf ~/.pingcli/credentials

# Or remove specific grant type
rm ~/.pingcli/credentials/<env-id>_<client-id>_device_code.json
```

## Troubleshooting

### Permission Errors
**Error:** `Failed to remove token from keychain`
**Solution:** Ensure proper OS permissions for credential store access

### Token Not Found
**Warning:** Token not found during logout
**Result:** Normal - indicates already logged out or no previous authentication

### Cache Issues
**Problem:** Still authenticated after logout
**Solution:** Restart CLI or clear application cache manually

## Best Practices

### Development Workflow
```bash
# Start session
pingcli login --device-code

# Work with APIs
pingcli request get /environments

# End session (good practice)
pingcli logout
```

### CI/CD Integration
```yaml
# Example GitHub Actions
- name: Authenticate
  run: pingcli login --client-credentials

- name: Run commands
  run: |
    pingcli export --service pingone --format terraform
    
- name: Cleanup
  if: always()
  run: pingcli logout
```

### Security Checklist
- ✅ Logout after each session on shared machines
- ✅ Include logout in automation cleanup steps  
- ✅ Verify logout success before leaving sessions
- ✅ Use logout for troubleshooting auth issues

## Return Codes

| Code | Meaning |
|------|---------|
| `0` | Success - tokens cleared |
| `1` | Error - logout failed |

## See Also
- [Authentication Overview](overview.md)
- [Login Command](login.md)
- [Security Best Practices](overview.md#security-best-practices)