# `pingcli auth logout`

Clear stored authentication tokens from the system credential store.

## Synopsis

Logout removes all stored authentication tokens and clears cached API client connections. Use this command when switching between environments, ending sessions, or troubleshooting authentication issues.

## Usage
```bash
pingcli auth logout
```

## Flags
- `-h, --help` - Help for logout command

## What Gets Cleared

### Tokens
- Access tokens
- Refresh tokens
- Token metadata (expiry, creation time)

### Cache
- PingOne API client cache
- Cached authentication state

### Storage Locations
- **macOS**: Keychain Services
- **Windows**: Windows Credential Manager  
- **Linux**: Secret Service API (GNOME Keyring/KDE KWallet)

## Examples

### Basic Logout
```bash
pingcli auth logout
```
**Output:**
```
Successfully logged out. Credentials cleared from Keychain.
```

### Logout in Automation
```bash
#!/bin/bash
# CI/CD cleanup script
pingcli auth logout
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
Error: no authentication found in Keychain. Please run 'pingcli login --device-code' to authenticate
```

## Manual Token Removal

If logout fails, manually remove tokens:

### macOS
```bash
# Command line
security delete-generic-password -s "pingcli" -a "device-code-token"

# GUI: Keychain Access → search "pingcli" → delete entry
```

### Windows
```cmd
# Command line
cmdkey /delete:LegacyGeneric:target=pingcli

# GUI: Control Panel → Credential Manager → remove pingcli entry
```

### Linux
```bash
# GNOME
secret-tool clear service pingcli

# GUI: seahorse → search "pingcli" → delete
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
pingcli auth login --device-code

# Work with APIs
pingcli request get /environments

# End session (good practice)
pingcli auth logout
```

### CI/CD Integration
```yaml
# Example GitHub Actions
- name: Authenticate
  run: pingcli auth login --client-credentials

- name: Run commands
  run: |
    pingcli export --service pingone --format terraform
    
- name: Cleanup
  if: always()
  run: pingcli auth logout
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