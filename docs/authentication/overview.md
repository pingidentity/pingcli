# Authentication

## Overview

The CLI uses OAuth2 flows to authenticate with PingOne services and stores tokens securely in the operating system's credential store.

## Authentication Methods

```bash
# Device Code Flow (recommended for interactive use)
pingcli auth login --device-code

# Authorization Code Flow (web browser required)
pingcli auth login --auth-code

# Client Credentials Flow (for automation/CI)
pingcli auth login --client-credentials
```

## Token Storage

Tokens are stored securely using OS-native credential stores:

| OS | Storage Location | Access Method |
|---|---|---|
| **macOS** | Keychain Services | Keychain Access app |
| **Windows** | Credential Manager | Control Panel → Credential Manager |
| **Linux** | Secret Service API | GNOME Keyring / KDE KWallet |

**Service Name**: `pingcli`  
**Account Name**: `device-code-token`

## Token Management

### View Token Status
```bash
# Check if authenticated (does not show token)
pingcli auth status  # (if implemented)
```

### Logout (Clear Tokens)
```bash
# Remove all stored tokens
pingcli auth logout
```

### Manual Token Access

**⚠️ Security Warning**: Tokens provide full access to your PingOne environment. Handle with care.

<details>
<summary>macOS - View stored token</summary>

```bash
# Command line
security find-generic-password -s "pingcli" -a "device-code-token" -w

# GUI: Keychain Access app → search "pingcli" → Show password
```
</details>

<details>
<summary>Windows - View stored token</summary>

```cmd
# Command line
cmdkey /list | findstr pingcli

# GUI: Control Panel → User Accounts → Credential Manager → Windows Credentials
```
</details>

<details>
<summary>Linux - View stored token</summary>

```bash
# GNOME
seahorse  # Search for "pingcli"

# KDE  
kwalletmanager

# Command line (depends on backend)
secret-tool lookup service pingcli account device-code-token
```
</details>

## Token Lifecycle

- **Automatic Refresh**: Expired tokens are automatically refreshed using stored refresh tokens
- **Re-authentication**: If refresh fails, run `pingcli auth login` again
- **Expiration**: Tokens typically expire after 1 hour (configurable by PingOne admin)

## Security Best Practices

### Development
- Use device code flow for interactive development
- Logout when switching between environments
- Never commit tokens to version control

### CI/CD
- Use client credentials flow for automation
- Store credentials as secure environment variables
- Use dedicated service accounts with minimal permissions
- Rotate credentials regularly

### Troubleshooting
- Use `pingcli auth logout` to clear corrupted tokens
- Ensure system keyring services are running on Linux
- On headless systems, consider using client credentials with environment variables

## Configuration

Authentication settings are managed via:
```bash
# Device code settings
pingcli config set service.pingone.authentication.deviceCode.clientID=<client-id>
pingcli config set service.pingone.authentication.deviceCode.environmentID=<env-id>

# Client credentials settings  
pingcli config set service.pingone.authentication.clientCredentials.clientID=<client-id>
pingcli config set service.pingone.authentication.clientCredentials.clientSecret=<secret>
pingcli config set service.pingone.authentication.clientCredentials.environmentID=<env-id>

# Regional settings
pingcli config set service.pingone.regionCode=<NA|EU|CA|AP|AU|SG>
```

## Troubleshooting

| Issue | Solution |
|---|---|
| "No authentication found" | Run `pingcli auth login` |
| "Failed to get valid token" | Run `pingcli auth logout` then `pingcli auth login` |
| Linux keyring errors | Ensure desktop session and keyring service running |
| Permission denied | Check OS credential store permissions |