# Plugin Authentication

Plugins allow you to seamlessly interact with Ping Identity services by leveraging the authentication context of the `pingcli` host. This guide explains how to use the authentication interface within your plugin.

## Overview

The `pingcli` host manages the complexities of OAuth2 and OIDC authentication (tokens, refresh flows, storage). Plugins simply request a token from the host when needed. This design ensures that plugins remain lightweight and secure, as they do not need to manage sensitive credentials directly.

## The Authentication Interface

Plugins receive an `auth` object (implementing `grpc.Authentication`) in their `Run` method.

**Interface Definition:**
```go
type Authentication interface {
    GetToken() (string, error)
}
```

## detailed Usage

To use the authentication token in your plugin:

1.  **Call GetToken**: Invoke the method on the provided `auth` instance.
2.  **Handle Errors**: Always check for errors. The host may return an error if the user is not logged in, the session has expired and cannot be refreshed, or if no profile is active.
3.  **Use the Token**: The returned string is a raw access token (JWT or opaque). Typically, you will use this as a Bearer token in the `Authorization` header of your HTTP requests.

**Example:**
```go
func (c *MyCommand) Run(args []string, logger grpc.Logger, auth grpc.Authentication) error {
    // 1. Request the token
    token, err := auth.GetToken()
    if err != nil {
        logger.UserError("Failed to get access token. Please run 'pingcli auth login' and try again.", nil)
        return err
    }

    // 2. Use the token to call an API
    client := &http.Client{}
    req, _ := http.NewRequest("GET", "https://api.pingone.com/v1/environments", nil)
    req.Header.Add("Authorization", "Bearer " + token)
    
    // ... execute request
    return nil
}
```

## How It Works Under the Hood

When `auth.GetToken()` is called:

1.  The `pingcli` host identifies the currently active profile (selected via `pingcli config`).
2.  It retrieves the cached access token for that profile.
3.  If the token is expired, the host automatically attempts to refresh it using the stored refresh token.
4.  If the token is valid (or successfully refreshed), it is returned to the plugin.
5.  If authentication fails (e.g., refresh token expired), an error is returned.

This abstraction allows your plugin to work with any environment or region configuration without needing specific logic for each authentication method.
