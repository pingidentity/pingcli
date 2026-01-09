// Copyright © 2025 Ping Identity Corporation

package auth_internal

import (
	"context"
	"crypto/sha256"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/pingidentity/pingcli/internal/configuration/options"
	"github.com/pingidentity/pingcli/internal/customtypes"
	"github.com/pingidentity/pingcli/internal/errs"
	"github.com/pingidentity/pingcli/internal/profiles"
	"github.com/pingidentity/pingone-go-client/config"
	svcOAuth2 "github.com/pingidentity/pingone-go-client/oauth2"
	"golang.org/x/oauth2"
)

// Token storage keys for different authentication methods
const (
	deviceCodeTokenKey        = "device-code-token"
	authorizationCodeTokenKey = "authorization-code-token" // #nosec G101 -- This is a keychain identifier, not a credential
	clientCredentialsTokenKey = "client-credentials-token"
)

var (
	credentialsErrorPrefix = "failed to manage credentials"
)

// getTokenStorage returns the appropriate keychain storage instance for the given authentication method
func getTokenStorage(authMethod string) (*svcOAuth2.KeychainStorage, error) {
	return svcOAuth2.NewKeychainStorage("pingcli", authMethod)
}

// shouldUseKeychain checks if keychain storage should be used based on the storage type
// Returns true if storage type is secure_local (default), false for file_system/none
func shouldUseKeychain() bool {
	v, err := profiles.GetOptionValue(options.AuthStorageOption)
	if err != nil {
		return true // default to keychain
	}
	s := strings.TrimSpace(strings.ToLower(v))
	if s == "" {
		return true // default to keychain
	}
	// Back-compat: boolean handling (true => file_system, false => secure_local)
	if s == "true" {
		return false
	}
	if s == "false" {
		return true
	}
	switch s {
	case string(config.StorageTypeSecureLocal):
		return true
	case string(config.StorageTypeFileSystem), string(config.StorageTypeNone), string(config.StorageTypeSecureRemote):
		return false
	default:
		// Unrecognized: lean secure by not disabling keychain
		return true
	}
}

// getStorageType returns the appropriate storage type for SDK keychain operations
// SDK handles keychain storage, pingcli handles file storage separately
func getStorageType() config.StorageType {
	v, _ := profiles.GetOptionValue(options.AuthStorageOption)
	s := strings.TrimSpace(strings.ToLower(v))
	if s == "false" || s == string(config.StorageTypeSecureLocal) || s == "" {
		return config.StorageTypeSecureLocal
	}
	// For file_system/none/secure_remote, avoid SDK persistence (pingcli manages file persistence)
	return config.StorageTypeNone
}

// generateTokenKey generates a unique token key based on provider, environmentID, clientID, and grantType
// Format: token-<hash>_<service>_<grantType>_<profile>.json
// The hash is based on service:environmentID:clientID:grantType for uniqueness
// Service and profile name are added as suffixes to enable service-specific token management and cleanup
func generateTokenKey(providerName, profileName, environmentID, clientID, grantType string) string {
	if providerName == "" || environmentID == "" || clientID == "" || grantType == "" {
		return ""
	}

	// Hash service + environment + client + grant type for uniqueness
	hash := sha256.Sum256([]byte(fmt.Sprintf("%s:%s:%s:%s", providerName, environmentID, clientID, grantType)))

	// Add profile name as suffix (default to "default" if empty)
	if profileName == "" {
		profileName = "default"
	}

	return fmt.Sprintf("token-%x_%s_%s_%s", hash[:8], providerName, grantType, profileName)
}

// LoginResult contains the result of a login operation
type LoginResult struct {
	Token    *oauth2.Token
	NewAuth  bool
	Location config.StorageType
}

// SaveTokenForMethod saves an OAuth2 token to file storage using the specified authentication method key
// Note: SDK handles keychain storage separately with its own token key format
// Returns config.StorageType indicating where the token was saved
func SaveTokenForMethod(token *oauth2.Token, authMethod string) (config.StorageType, error) {
	if token == nil {
		return config.StorageTypeNone, ErrNilToken
	}

	// Check for "none" storage type - do not save anywhere
	v, _ := profiles.GetOptionValue(options.AuthStorageOption)
	if strings.TrimSpace(strings.ToLower(v)) == string(config.StorageTypeNone) {
		return config.StorageTypeNone, nil
	}

	// Avoid saving to keychain here: SDK handles keychain persistence via TokenSource.
	// When keychain is enabled, do NOT write a file. Only indicate keychain is in use.
	if shouldUseKeychain() {
		return config.StorageTypeSecureLocal, nil
	}

	// File-only mode: save only to file storage and error if unsuccessful.
	if err := saveTokenToFile(token, authMethod); err != nil {
		return config.StorageTypeFileSystem, err
	}

	return config.StorageTypeFileSystem, nil
}

// LoadTokenForMethod loads an OAuth2 token from the keychain using the specified authentication method key
// Falls back to file storage if keychain operations fail or if --use-keychain=false
func LoadTokenForMethod(authMethod string) (*oauth2.Token, error) {
	// Check for "none" storage type - do not load from anywhere
	v, _ := profiles.GetOptionValue(options.AuthStorageOption)
	if strings.TrimSpace(strings.ToLower(v)) == string(config.StorageTypeNone) {
		return nil, nil
	}

	// Check if user disabled keychain
	if !shouldUseKeychain() {
		// Directly load from file storage
		return loadTokenFromFile(authMethod)
	}

	storage, err := getTokenStorage(authMethod)
	if err != nil {
		return nil, &errs.PingCLIError{
			Prefix: credentialsErrorPrefix,
			Err:    err,
		}
	}

	// Try keychain storage first
	token, err := storage.LoadToken()
	if err == nil {
		return token, nil // Success!
	}

	// Keychain failed, try file fallback
	token, fileErr := loadTokenFromFile(authMethod)
	if fileErr == nil {
		return token, nil // Success with fallback!
	}

	// Both failed (err and fileErr are non-nil)
	return nil, &errs.PingCLIError{
		Prefix: credentialsErrorPrefix,
		Err:    errors.Join(err, fileErr),
	}
}

// LoadToken attempts to load an OAuth2 token from the keychain, trying configured auth methods first
func LoadToken() (*oauth2.Token, error) {
	// First, try to load using configuration-based keys from the active profile
	authType, err := profiles.GetOptionValue(options.PingOneAuthenticationTypeOption)
	if err == nil && authType != "" {
		// Normalize auth type to snake_case format and handle camelCase aliases
		switch authType {
		case "clientCredentials":
			authType = "client_credentials"
		case "deviceCode":
			authType = "device_code"
		case "authorizationCode":
			authType = "authorization_code"
		case "authorization_code":
			authType = "authorization_code"
		}

		// Try to get configuration for the configured grant type
		var cfg *config.Configuration
		var grantType svcOAuth2.GrantType
		switch authType {
		case "device_code":
			cfg, err = GetDeviceCodeConfiguration()
			if err != nil {
				return nil, &errs.PingCLIError{
					Prefix: credentialsErrorPrefix,
					Err:    err,
				}
			}
			grantType = svcOAuth2.GrantTypeDeviceCode
		case "authorization_code":
			cfg, err = GetAuthorizationCodeConfiguration()
			if err != nil {
				return nil, &errs.PingCLIError{
					Prefix: credentialsErrorPrefix,
					Err:    err,
				}
			}
			grantType = svcOAuth2.GrantTypeAuthorizationCode
		case "client_credentials":
			cfg, err = GetClientCredentialsConfiguration()
			if err != nil {
				return nil, &errs.PingCLIError{
					Prefix: credentialsErrorPrefix,
					Err:    err,
				}
			}
			grantType = svcOAuth2.GrantTypeClientCredentials
		case "worker":
			cfg, err = GetWorkerConfiguration()
			if err != nil {
				return nil, &errs.PingCLIError{
					Prefix: credentialsErrorPrefix,
					Err:    err,
				}
			}
			grantType = svcOAuth2.GrantTypeClientCredentials
		}

		if cfg != nil {
			// Set the grant type before generating the token key
			cfg = cfg.WithGrantType(grantType)
			tokenKey, err := GetAuthMethodKeyFromConfig(cfg)
			if err != nil {
				return nil, &errs.PingCLIError{
					Prefix: credentialsErrorPrefix,
					Err:    err,
				}
			}

			token, err := LoadTokenForMethod(tokenKey)
			if err != nil {
				return nil, &errs.PingCLIError{
					Prefix: credentialsErrorPrefix,
					Err:    err,
				}
			}

			return token, nil
		}
	}

	// No authorization grant type configured
	return nil, &errs.PingCLIError{
		Prefix: credentialsErrorPrefix,
		Err:    ErrUnsupportedAuthType,
	}
}

// GetValidTokenSource returns a valid OAuth2 token source for the configured authentication method
func GetValidTokenSource(ctx context.Context) (oauth2.TokenSource, error) {
	// First, try to load using configuration-based keys from the active profile
	authType, err := profiles.GetOptionValue(options.PingOneAuthenticationTypeOption)
	if err != nil || authType == "" {
		return nil, &errs.PingCLIError{
			Prefix: credentialsErrorPrefix,
			Err:    ErrUnsupportedAuthType,
		}
	}

	// Normalize auth type to snake_case format and handle camelCase aliases
	switch authType {
	case "clientCredentials":
		authType = "client_credentials"
	case "deviceCode":
		authType = "device_code"
	case "authorizationCode":
		authType = "authorization_code"
	case "authorization_code":
		authType = "authorization_code"
	}

	// Try to get configuration for the configured grant type
	var cfg *config.Configuration
	var grantType svcOAuth2.GrantType
	switch authType {
	case "device_code":
		cfg, err = GetDeviceCodeConfiguration()
		if err != nil {
			return nil, &errs.PingCLIError{
				Prefix: credentialsErrorPrefix,
				Err:    err,
			}
		}
		grantType = svcOAuth2.GrantTypeDeviceCode
	case "authorization_code":
		cfg, err = GetAuthorizationCodeConfiguration()
		if err != nil {
			return nil, &errs.PingCLIError{
				Prefix: credentialsErrorPrefix,
				Err:    err,
			}
		}
		grantType = svcOAuth2.GrantTypeAuthorizationCode
	case "client_credentials":
		cfg, err = GetClientCredentialsConfiguration()
		if err != nil {
			return nil, &errs.PingCLIError{
				Prefix: credentialsErrorPrefix,
				Err:    err,
			}
		}
		grantType = svcOAuth2.GrantTypeClientCredentials
	case "worker":
		cfg, err = GetWorkerConfiguration()
		if err != nil {
			return nil, &errs.PingCLIError{
				Prefix: credentialsErrorPrefix,
				Err:    err,
			}
		}
		grantType = svcOAuth2.GrantTypeClientCredentials
	default:
		return nil, &errs.PingCLIError{
			Prefix: credentialsErrorPrefix,
			Err:    fmt.Errorf("%w: %s", ErrUnsupportedAuthType, authType),
		}
	}

	if cfg != nil {
		// Set the grant type before getting the token source
		cfg = cfg.WithGrantType(grantType)

		// If using file storage, try to seed refresh from existing file token before new login
		if !shouldUseKeychain() {
			tokenKey, err := GetAuthMethodKeyFromConfig(cfg)
			if err == nil && tokenKey != "" {
				if existingToken, ferr := loadTokenFromFile(tokenKey); ferr == nil && existingToken != nil && existingToken.RefreshToken != "" {
					// Build minimal oauth2.Config for refresh using SDK endpoints
					endpoints, eerr := cfg.AuthEndpoints()
					if eerr == nil {
						var oauthCfg *oauth2.Config
						switch grantType {
						case svcOAuth2.GrantTypeDeviceCode:
							// Device Code: use client ID and optional scopes
							if cfg.Auth.DeviceCode != nil && cfg.Auth.DeviceCode.DeviceCodeClientID != nil {
								var scopes []string
								if cfg.Auth.DeviceCode.DeviceCodeScopes != nil {
									scopes = *cfg.Auth.DeviceCode.DeviceCodeScopes
								}
								oauthCfg = &oauth2.Config{ClientID: *cfg.Auth.DeviceCode.DeviceCodeClientID, Endpoint: endpoints.Endpoint, Scopes: scopes}
							}
						case svcOAuth2.GrantTypeAuthorizationCode:
							// Auth Code: use client ID and optional scopes
							if cfg.Auth.AuthorizationCode != nil && cfg.Auth.AuthorizationCode.AuthorizationCodeClientID != nil {
								var scopes []string
								if cfg.Auth.AuthorizationCode.AuthorizationCodeScopes != nil {
									scopes = *cfg.Auth.AuthorizationCode.AuthorizationCodeScopes
								}
								oauthCfg = &oauth2.Config{ClientID: *cfg.Auth.AuthorizationCode.AuthorizationCodeClientID, Endpoint: endpoints.Endpoint, Scopes: scopes}
							}
						default:
							// client_credentials typically lacks refresh; fall through
						}

						if oauthCfg != nil {
							baseTS := oauthCfg.TokenSource(ctx, existingToken)

							return oauth2.ReuseTokenSource(nil, baseTS), nil
						}
					}
				}
			}
		}

		// Fallback: use SDK TokenSource (may perform new auth)
		tokenSource, err := cfg.TokenSource(ctx)
		if err != nil {
			return nil, &errs.PingCLIError{
				Prefix: credentialsErrorPrefix,
				Err:    err,
			}
		}

		return tokenSource, nil
	}

	return nil, &errs.PingCLIError{
		Prefix: credentialsErrorPrefix,
		Err:    ErrUnsupportedAuthType,
	}
}

// ClearToken removes all cached tokens from the keychain for all authentication methods.
// This clears tokens from ALL grant types, not just the currently configured one,
// to handle cases where users switch between authentication methods
func ClearToken() error {
	var errs []error

	// Clear configuration-based tokens for all auth methods
	// Also clear any old tokens from previous configurations with different client IDs
	authMethods := []struct {
		name      string
		getConfig func() (*config.Configuration, error)
		grantType svcOAuth2.GrantType
	}{
		{"client_credentials", GetClientCredentialsConfiguration, svcOAuth2.GrantTypeClientCredentials},
		{"device_code", GetDeviceCodeConfiguration, svcOAuth2.GrantTypeDeviceCode},
		{"authorization_code", GetAuthorizationCodeConfiguration, svcOAuth2.GrantTypeAuthorizationCode},
	}

	for _, method := range authMethods {
		// Try to clear token with current configuration (if it exists)
		cfg, err := method.getConfig()
		if err == nil && cfg != nil {
			// Set the grant type before generating the token key
			cfg = cfg.WithGrantType(method.grantType)
			tokenKey, err := GetAuthMethodKeyFromConfig(cfg)
			if err == nil {
				// Clear from keychain using current config
				keychainStorage, err := svcOAuth2.NewKeychainStorage("pingcli", tokenKey)
				if err == nil {
					if err := keychainStorage.ClearToken(); err != nil {
						errs = append(errs, err)
					}
				}
				// Clear from file storage using current config
				if err := clearTokenFromFile(tokenKey); err != nil {
					errs = append(errs, err)
				}
			}
		}

		// Always clear all token files for this grant type and current profile (handles old configurations)
		// This is important even if the user isn't currently using this grant type
		grantTypeStr := string(method.grantType)
		profileName, err := profiles.GetOptionValue(options.RootActiveProfileOption)
		if err != nil {
			profileName = "default" // Fallback to default if we can't get profile name
		}
		providerName, err := profiles.GetOptionValue(options.AuthProviderOption)
		if err != nil || providerName == "" {
			providerName = "pingone" // Default to pingone
		}
		if err := clearAllTokenFilesForGrantType(providerName, grantTypeStr, profileName); err != nil {
			errs = append(errs, err)
		}
	}

	// Also clear tokens using simple string keys for backward compatibility
	methods := []string{deviceCodeTokenKey, authorizationCodeTokenKey, clientCredentialsTokenKey}

	for _, method := range methods {
		storage, err := getTokenStorage(method)
		if err == nil {
			if err := storage.ClearToken(); err != nil {
				errs = append(errs, err)
			}
		}
		// Also clear from file storage
		if err := clearTokenFromFile(method); err != nil {
			errs = append(errs, err)
		}
	}

	return errors.Join(errs...)
}

// ClearTokenForMethod removes the cached token for a specific authentication method
// Clears from both keychain and file storage
// Returns config.StorageType indicating what was cleared
func ClearTokenForMethod(authMethod string) (config.StorageType, error) {
	var errList []error
	clearedType := config.StorageTypeNone

	// Clear from keychain
	storage, err := getTokenStorage(authMethod)
	if err == nil {
		if err := storage.ClearToken(); err != nil {
			errList = append(errList, &errs.PingCLIError{
				Prefix: credentialsErrorPrefix,
				Err:    err,
			})
		} else {
			clearedType = config.StorageTypeSecureLocal
		}
	}

	// Also clear from file storage
	if err := clearTokenFromFile(authMethod); err != nil {
		errList = append(errList, &errs.PingCLIError{
			Prefix: credentialsErrorPrefix,
			Err:    err,
		})
	} else {
		if clearedType == config.StorageTypeNone {
			clearedType = config.StorageTypeFileSystem
		}
	}

	return clearedType, errors.Join(errList...)
}

// PerformDeviceCodeLogin performs device code authentication, returning the result
func PerformDeviceCodeLogin(ctx context.Context) (*LoginResult, error) {
	cfg, err := GetDeviceCodeConfiguration()
	if err != nil {
		return nil, &errs.PingCLIError{
			Prefix: credentialsErrorPrefix,
			Err:    err,
		}
	}

	// Get profile name for token key generation
	profileName, err := profiles.GetOptionValue(options.RootActiveProfileOption)
	if err != nil {
		profileName = "default" // Fallback to default if we can't get profile name
	}

	// Get service name for token key generation
	providerName, err := profiles.GetOptionValue(options.AuthProviderOption)
	if err != nil || strings.TrimSpace(providerName) == "" {
		providerName = customtypes.ENUM_AUTH_PROVIDER_PINGONE // Default to pingone
	}

	// Client ID and environment ID no longer needed for manual key generation

	// Set grant type to device code
	cfg = cfg.WithGrantType(svcOAuth2.GrantTypeDeviceCode)

	// Use SDK-consistent token key generation to avoid mismatches
	tokenKey, err := GetAuthMethodKeyFromConfig(cfg)
	if err != nil || tokenKey == "" {
		// Fallback to simple key if generation fails
		tokenKey = deviceCodeTokenKey
	}

	// Check if we have a valid cached token before calling TokenSource
	// Store the existing token's expiry to compare later
	var existingTokenExpiry *time.Time

	// First try SDK keychain storage if enabled
	if shouldUseKeychain() {
		keychainStorage, err := svcOAuth2.NewKeychainStorage("pingcli", tokenKey)
		if err == nil {
			if existingToken, err := keychainStorage.LoadToken(); err == nil && existingToken != nil && existingToken.Valid() {
				existingTokenExpiry = &existingToken.Expiry
			}
		}
	}

	// If not found in keychain, check file storage
	if existingTokenExpiry == nil {
		if existingToken, err := loadTokenFromFile(tokenKey); err == nil && existingToken != nil && existingToken.Valid() {
			existingTokenExpiry = &existingToken.Expiry
		}
	}

	// If using file storage and we have a refresh token, seed refresh via oauth2.ReuseTokenSource
	var tokenSource oauth2.TokenSource
	if !shouldUseKeychain() {
		if existingToken, err := loadTokenFromFile(tokenKey); err == nil && existingToken != nil && existingToken.RefreshToken != "" {
			endpoints, eerr := cfg.AuthEndpoints()
			if eerr == nil && cfg.Auth.DeviceCode != nil && cfg.Auth.DeviceCode.DeviceCodeClientID != nil {
				var scopes []string
				if cfg.Auth.DeviceCode.DeviceCodeScopes != nil {
					scopes = *cfg.Auth.DeviceCode.DeviceCodeScopes
				}
				oauthCfg := &oauth2.Config{ClientID: *cfg.Auth.DeviceCode.DeviceCodeClientID, Endpoint: endpoints.Endpoint, Scopes: scopes}
				baseTS := oauthCfg.TokenSource(ctx, existingToken)
				tokenSource = oauth2.ReuseTokenSource(nil, baseTS)
			}
		}
	}
	// Fallback to SDK token source if we didn't create a seeded one
	if tokenSource == nil {
		var tsErr error
		tokenSource, tsErr = cfg.TokenSource(ctx)
		if tsErr != nil {
			return nil, &errs.PingCLIError{
				Prefix: credentialsErrorPrefix,
				Err:    tsErr,
			}
		}
	}
	if err != nil {
		return nil, &errs.PingCLIError{
			Prefix: credentialsErrorPrefix,
			Err:    err,
		}
	}

	// Get token (SDK will return cached token if valid, or perform new authentication)
	token, err := tokenSource.Token()
	if err != nil {
		return nil, &errs.PingCLIError{
			Prefix: credentialsErrorPrefix,
			Err:    err,
		}
	}

	// Clean up old token files for this grant type and profile (in case configuration changed)
	// Ignore errors from cleanup - we still want to save the new token
	_ = clearAllTokenFilesForGrantType(providerName, string(svcOAuth2.GrantTypeDeviceCode), profileName)

	// Save token using our own storage logic (handles both file and keychain based on flags)
	location, err := SaveTokenForMethod(token, tokenKey)
	if err != nil {
		return nil, &errs.PingCLIError{
			Prefix: credentialsErrorPrefix,
			Err:    err,
		}
	}

	// Determine if this was new authentication
	// If we had an existing token with the same expiry, it's cached
	// If expiry is different, new auth was performed
	isNewAuth := existingTokenExpiry == nil || !token.Expiry.Equal(*existingTokenExpiry)

	// NewAuth indicates whether new authentication was performed
	return &LoginResult{
		Token:    token,
		NewAuth:  isNewAuth,
		Location: location,
	}, nil
}

// GetDeviceCodeConfiguration builds a device code authentication configuration from the CLI profile options
func GetDeviceCodeConfiguration() (*config.Configuration, error) {
	cfg := config.NewConfiguration()

	// Get device code client ID
	clientID, err := profiles.GetOptionValue(options.PingOneAuthenticationDeviceCodeClientIDOption)
	if err != nil {
		return nil, &errs.PingCLIError{
			Prefix: credentialsErrorPrefix,
			Err:    err,
		}
	}
	if clientID == "" {
		return nil, &errs.PingCLIError{
			Prefix: credentialsErrorPrefix,
			Err:    ErrDeviceCodeClientIDNotConfigured,
		}
	}

	// Configure device code settings
	cfg = cfg.WithDeviceCodeClientID(clientID)

	// This is the default scope. Additional scopes can be appended by the user later if needed.
	scopeDefaults := []string{"openid"}
	cfg = cfg.WithDeviceCodeScopes(scopeDefaults)

	// Configure storage options based on --file-storage flag
	cfg = cfg.WithStorageType(getStorageType()).WithStorageName("pingcli")

	// Provide optional suffix so SDK keychain entries align with file names
	profileName, _ := profiles.GetOptionValue(options.RootActiveProfileOption)
	if strings.TrimSpace(profileName) == "" {
		profileName = "default"
	}
	providerName, _ := profiles.GetOptionValue(options.AuthProviderOption)
	if strings.TrimSpace(providerName) == "" {
		providerName = customtypes.ENUM_AUTH_PROVIDER_PINGONE
	}
	cfg = cfg.WithStorageOptionalSuffix(fmt.Sprintf("_%s_%s_%s", providerName, string(svcOAuth2.GrantTypeDeviceCode), profileName))

	// Apply Environment ID for consistent token key generation and endpoints
	environmentID, err := profiles.GetOptionValue(options.PingOneAuthenticationAPIEnvironmentIDOption)
	if err != nil {
		return nil, &errs.PingCLIError{
			Prefix: credentialsErrorPrefix,
			Err:    err,
		}
	}
	if strings.TrimSpace(environmentID) == "" {
		// Fallback: deprecated worker environment ID
		workerEnvID, wErr := profiles.GetOptionValue(options.PingOneAuthenticationWorkerEnvironmentIDOption)
		if wErr == nil && strings.TrimSpace(workerEnvID) != "" {
			environmentID = workerEnvID
		}
	}
	if strings.TrimSpace(environmentID) != "" {
		cfg = cfg.WithEnvironmentID(environmentID)
	}

	// Apply region configuration
	cfg, err = applyRegionConfiguration(cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}

func PerformAuthorizationCodeLogin(ctx context.Context) (*LoginResult, error) {
	cfg, err := GetAuthorizationCodeConfiguration()
	if err != nil {
		return nil, &errs.PingCLIError{
			Prefix: credentialsErrorPrefix,
			Err:    err,
		}
	}

	// Get profile name for token key generation
	profileName, err := profiles.GetOptionValue(options.RootActiveProfileOption)
	if err != nil {
		profileName = "default" // Fallback to default if we can't get profile name
	}

	// Get service name for token key generation
	providerName, err := profiles.GetOptionValue(options.AuthProviderOption)
	if err != nil || strings.TrimSpace(providerName) == "" {
		providerName = customtypes.ENUM_AUTH_PROVIDER_PINGONE // Default to pingone
	}

	// Client ID and environment ID no longer needed for manual key generation

	// Set grant type to authorization code
	cfg = cfg.WithGrantType(svcOAuth2.GrantTypeAuthorizationCode)

	// Use SDK-consistent token key generation to avoid mismatches
	tokenKey, err := GetAuthMethodKeyFromConfig(cfg)
	if err != nil || tokenKey == "" {
		// Fallback to simple key if generation fails
		tokenKey = authorizationCodeTokenKey
	}

	// Check if we have a valid cached token before calling TokenSource
	// Store the existing token's expiry to compare later
	var existingTokenExpiry *time.Time

	// First try SDK keychain storage if enabled
	if shouldUseKeychain() {
		keychainStorage, err := svcOAuth2.NewKeychainStorage("pingcli", tokenKey)
		if err == nil {
			if existingToken, err := keychainStorage.LoadToken(); err == nil && existingToken != nil && existingToken.Valid() {
				existingTokenExpiry = &existingToken.Expiry
			}
		}
	}

	// If not found in keychain, check file storage
	if existingTokenExpiry == nil {
		if existingToken, err := loadTokenFromFile(tokenKey); err == nil && existingToken != nil && existingToken.Valid() {
			existingTokenExpiry = &existingToken.Expiry
		}
	}

	// If using file storage and we have a refresh token, seed refresh via oauth2.ReuseTokenSource
	var tokenSource oauth2.TokenSource
	if !shouldUseKeychain() {
		if existingToken, err := loadTokenFromFile(tokenKey); err == nil && existingToken != nil && existingToken.RefreshToken != "" {
			endpoints, eerr := cfg.AuthEndpoints()
			if eerr == nil && cfg.Auth.AuthorizationCode != nil && cfg.Auth.AuthorizationCode.AuthorizationCodeClientID != nil {
				var scopes []string
				if cfg.Auth.AuthorizationCode.AuthorizationCodeScopes != nil {
					scopes = *cfg.Auth.AuthorizationCode.AuthorizationCodeScopes
				}
				oauthCfg := &oauth2.Config{ClientID: *cfg.Auth.AuthorizationCode.AuthorizationCodeClientID, Endpoint: endpoints.Endpoint, Scopes: scopes}
				baseTS := oauthCfg.TokenSource(ctx, existingToken)
				tokenSource = oauth2.ReuseTokenSource(nil, baseTS)
			}
		}
	}
	// Fallback to SDK token source if we didn't create a seeded one
	if tokenSource == nil {
		var tsErr error
		tokenSource, tsErr = cfg.TokenSource(ctx)
		if tsErr != nil {
			return nil, &errs.PingCLIError{
				Prefix: credentialsErrorPrefix,
				Err:    tsErr,
			}
		}
	}
	if err != nil {
		return nil, &errs.PingCLIError{
			Prefix: credentialsErrorPrefix,
			Err:    err,
		}
	}

	// Get token (SDK will return cached token if valid, or perform new authentication)
	token, err := tokenSource.Token()
	if err != nil {
		return nil, &errs.PingCLIError{
			Prefix: credentialsErrorPrefix,
			Err:    err,
		}
	}

	// Clean up old token files for this grant type and profile (in case configuration changed)
	// Ignore errors from cleanup - we still want to save the new token
	_ = clearAllTokenFilesForGrantType(providerName, string(svcOAuth2.GrantTypeAuthorizationCode), profileName)

	// Save token using our own storage logic (handles both file and keychain based on flags)
	location, err := SaveTokenForMethod(token, tokenKey)
	if err != nil {
		return nil, &errs.PingCLIError{
			Prefix: credentialsErrorPrefix,
			Err:    err,
		}
	}

	// Determine if this was new authentication
	// If we had an existing token with the same expiry, it's cached
	// If expiry is different, new auth was performed
	isNewAuth := existingTokenExpiry == nil || !token.Expiry.Equal(*existingTokenExpiry)

	// NewAuth indicates whether new authentication was performed
	return &LoginResult{
		Token:    token,
		NewAuth:  isNewAuth,
		Location: location,
	}, nil
}

// GetAuthorizationCodeConfiguration builds an authorization code authentication configuration from the CLI profile options
func GetAuthorizationCodeConfiguration() (*config.Configuration, error) {
	cfg := config.NewConfiguration()

	// Get authorization code client ID
	clientID, err := profiles.GetOptionValue(options.PingOneAuthenticationAuthorizationCodeClientIDOption)
	if err != nil {
		return nil, &errs.PingCLIError{
			Prefix: credentialsErrorPrefix,
			Err:    err,
		}
	}
	if clientID == "" {
		return nil, &errs.PingCLIError{
			Prefix: credentialsErrorPrefix,
			Err:    ErrAuthorizationCodeClientIDNotConfigured,
		}
	}

	// Get authorization code redirect URI path
	redirectURIPath, err := profiles.GetOptionValue(options.PingOneAuthenticationAuthorizationCodeRedirectURIPathOption)
	if err != nil {
		return nil, &errs.PingCLIError{
			Prefix: credentialsErrorPrefix,
			Err:    err,
		}
	}
	if redirectURIPath == "" {
		return nil, &errs.PingCLIError{
			Prefix: credentialsErrorPrefix,
			Err:    ErrAuthorizationCodeRedirectURIPathNotConfigured,
		}
	}

	// Get authorization code redirect URI port
	redirectURIPort, err := profiles.GetOptionValue(options.PingOneAuthenticationAuthorizationCodeRedirectURIPortOption)
	if err != nil {
		return nil, &errs.PingCLIError{
			Prefix: credentialsErrorPrefix,
			Err:    err,
		}
	}
	if redirectURIPort == "" {
		return nil, &errs.PingCLIError{
			Prefix: credentialsErrorPrefix,
			Err:    ErrAuthorizationCodeRedirectURIPortNotConfigured,
		}
	}

	redirectURI := config.AuthorizationCodeRedirectURI{
		Port: redirectURIPort,
		Path: redirectURIPath,
	}

	// Configure auth code settings
	cfg = cfg.WithAuthorizationCodeClientID(clientID).
		WithAuthorizationCodeRedirectURI(redirectURI)

	// This is the default scope. Additional scopes can be appended by the user later if needed.
	scopeDefaults := []string{"openid"}
	cfg = cfg.WithAuthorizationCodeScopes(scopeDefaults)

	// Configure storage options based on --file-storage flag
	cfg = cfg.WithStorageType(getStorageType()).
		WithStorageName("pingcli")

	// Provide optional suffix so SDK keychain entries align with file names
	profileName, _ := profiles.GetOptionValue(options.RootActiveProfileOption)
	if strings.TrimSpace(profileName) == "" {
		profileName = "default"
	}
	providerName, _ := profiles.GetOptionValue(options.AuthProviderOption)
	if strings.TrimSpace(providerName) == "" {
		providerName = customtypes.ENUM_AUTH_PROVIDER_PINGONE
	}
	cfg = cfg.WithStorageOptionalSuffix(fmt.Sprintf("_%s_%s_%s", providerName, string(svcOAuth2.GrantTypeAuthorizationCode), profileName))

	// Apply Environment ID for consistent token key generation and endpoints
	environmentID, err := profiles.GetOptionValue(options.PingOneAuthenticationAPIEnvironmentIDOption)
	if err != nil {
		return nil, &errs.PingCLIError{
			Prefix: credentialsErrorPrefix,
			Err:    err,
		}
	}
	if strings.TrimSpace(environmentID) == "" {
		// Fallback: deprecated worker environment ID
		workerEnvID, wErr := profiles.GetOptionValue(options.PingOneAuthenticationWorkerEnvironmentIDOption)
		if wErr == nil && strings.TrimSpace(workerEnvID) != "" {
			environmentID = workerEnvID
		}
	}
	if strings.TrimSpace(environmentID) != "" {
		cfg = cfg.WithEnvironmentID(environmentID)
	}

	// Apply region configuration
	cfg, err = applyRegionConfiguration(cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}

func PerformClientCredentialsLogin(ctx context.Context) (*LoginResult, error) {
	cfg, err := GetClientCredentialsConfiguration()
	if err != nil {
		return nil, &errs.PingCLIError{
			Prefix: credentialsErrorPrefix,
			Err:    err,
		}
	}

	// Get profile name for token key generation
	profileName, err := profiles.GetOptionValue(options.RootActiveProfileOption)
	if err != nil {
		profileName = "default" // Fallback to default if we can't get profile name
	}

	// Get service name for token key generation
	providerName, err := profiles.GetOptionValue(options.AuthProviderOption)
	if err != nil || strings.TrimSpace(providerName) == "" {
		providerName = customtypes.ENUM_AUTH_PROVIDER_PINGONE // Default to pingone
	}

	// Client ID and environment ID no longer needed for manual key generation

	// Set grant type to client credentials
	cfg = cfg.WithGrantType(svcOAuth2.GrantTypeClientCredentials)

	// Use SDK-consistent token key generation to avoid mismatches
	tokenKey, err := GetAuthMethodKeyFromConfig(cfg)
	if err != nil || tokenKey == "" {
		// Fallback to simple key if generation fails
		tokenKey = clientCredentialsTokenKey
	}

	// Check if we have a valid cached token before calling TokenSource
	// Store the existing token's expiry to compare later
	var existingTokenExpiry *time.Time

	// First try SDK keychain storage if enabled
	if shouldUseKeychain() {
		keychainStorage, err := svcOAuth2.NewKeychainStorage("pingcli", tokenKey)
		if err == nil {
			if existingToken, err := keychainStorage.LoadToken(); err == nil && existingToken != nil && existingToken.Valid() {
				existingTokenExpiry = &existingToken.Expiry
			}
		}
	}

	// If not found in keychain, check file storage
	if existingTokenExpiry == nil {
		if existingToken, err := loadTokenFromFile(tokenKey); err == nil && existingToken != nil && existingToken.Valid() {
			existingTokenExpiry = &existingToken.Expiry
		}
	}

	// Get token source - SDK handles keychain storage based on configuration
	tokenSource, err := cfg.TokenSource(ctx)
	if err != nil {
		return nil, &errs.PingCLIError{
			Prefix: credentialsErrorPrefix,
			Err:    err,
		}
	}

	// Get token (SDK will return cached token if valid, or perform new authentication)
	token, err := tokenSource.Token()
	if err != nil {
		return nil, &errs.PingCLIError{
			Prefix: credentialsErrorPrefix,
			Err:    err,
		}
	}

	// Clean up old token files for this grant type and profile (in case configuration changed)
	// Ignore errors from cleanup - we still want to save the new token
	_ = clearAllTokenFilesForGrantType(providerName, string(svcOAuth2.GrantTypeClientCredentials), profileName)

	// Save token using our own storage logic (handles both file and keychain based on flags)
	location, err := SaveTokenForMethod(token, tokenKey)
	if err != nil {
		return nil, &errs.PingCLIError{
			Prefix: credentialsErrorPrefix,
			Err:    err,
		}
	}

	// Determine if this was new authentication
	// If we had an existing token with the same expiry, it's cached
	// If expiry is different, new auth was performed
	isNewAuth := existingTokenExpiry == nil || !token.Expiry.Equal(*existingTokenExpiry)

	// NewAuth indicates whether new authentication was performed
	return &LoginResult{
		Token:    token,
		NewAuth:  isNewAuth,
		Location: location,
	}, nil
}

// GetClientCredentialsConfiguration builds a client credentials authentication configuration from the CLI profile options
func GetClientCredentialsConfiguration() (*config.Configuration, error) {
	cfg := config.NewConfiguration()

	// Get client credentials client ID
	clientID, err := profiles.GetOptionValue(options.PingOneAuthenticationClientCredentialsClientIDOption)
	if err != nil {
		return nil, &errs.PingCLIError{
			Prefix: credentialsErrorPrefix,
			Err:    err,
		}
	}
	if clientID == "" {
		return nil, &errs.PingCLIError{
			Prefix: credentialsErrorPrefix,
			Err:    ErrClientCredentialsClientIDNotConfigured,
		}
	}

	// Get client credentials client secret
	clientSecret, err := profiles.GetOptionValue(options.PingOneAuthenticationClientCredentialsClientSecretOption)
	if err != nil {
		return nil, &errs.PingCLIError{
			Prefix: credentialsErrorPrefix,
			Err:    err,
		}
	}
	if clientSecret == "" {
		return nil, &errs.PingCLIError{
			Prefix: credentialsErrorPrefix,
			Err:    ErrClientCredentialsClientSecretNotConfigured,
		}
	}

	// Configure client credentials settings
	cfg = cfg.WithClientCredentialsClientID(clientID).
		WithClientCredentialsClientSecret(clientSecret)

	// This is the default scope. Additional scopes can be appended by the user later if needed.
	scopeDefaults := []string{"openid"}
	cfg = cfg.WithClientCredentialsScopes(scopeDefaults)

	// Configure storage options based on --file-storage flag
	cfg = cfg.WithStorageType(getStorageType()).
		WithStorageName("pingcli")

	// Provide optional suffix so SDK keychain entries align with file names
	profileName, _ := profiles.GetOptionValue(options.RootActiveProfileOption)
	if strings.TrimSpace(profileName) == "" {
		profileName = "default"
	}
	providerName, _ := profiles.GetOptionValue(options.AuthProviderOption)
	if strings.TrimSpace(providerName) == "" {
		providerName = customtypes.ENUM_AUTH_PROVIDER_PINGONE
	}
	cfg = cfg.WithStorageOptionalSuffix(fmt.Sprintf("_%s_%s_%s", providerName, string(svcOAuth2.GrantTypeClientCredentials), profileName))

	// Apply Environment ID for consistent token key generation and endpoints
	environmentID, err := profiles.GetOptionValue(options.PingOneAuthenticationAPIEnvironmentIDOption)
	if err != nil {
		return nil, &errs.PingCLIError{
			Prefix: credentialsErrorPrefix,
			Err:    err,
		}
	}
	if strings.TrimSpace(environmentID) != "" {
		cfg = cfg.WithEnvironmentID(environmentID)
	}

	// Apply region configuration
	cfg, err = applyRegionConfiguration(cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}

// GetWorkerConfiguration builds a worker authentication configuration from the CLI profile options
func GetWorkerConfiguration() (*config.Configuration, error) {
	cfg := config.NewConfiguration()

	// Get worker client ID
	clientID, err := profiles.GetOptionValue(options.PingOneAuthenticationWorkerClientIDOption)
	if err != nil {
		return nil, &errs.PingCLIError{
			Prefix: credentialsErrorPrefix,
			Err:    err,
		}
	}
	if clientID == "" {
		return nil, &errs.PingCLIError{
			Prefix: credentialsErrorPrefix,
			Err:    ErrWorkerClientIDNotConfigured,
		}
	}

	// Get worker client secret
	clientSecret, err := profiles.GetOptionValue(options.PingOneAuthenticationWorkerClientSecretOption)
	if err != nil {
		return nil, &errs.PingCLIError{
			Prefix: credentialsErrorPrefix,
			Err:    err,
		}
	}
	if clientSecret == "" {
		return nil, &errs.PingCLIError{
			Prefix: credentialsErrorPrefix,
			Err:    ErrWorkerClientSecretNotConfigured,
		}
	}

	// Configure worker settings (client_credentials under the hood)
	cfg = cfg.WithClientCredentialsClientID(clientID).
		WithClientCredentialsClientSecret(clientSecret)
	// Align default scopes with client credentials flow
	scopeDefaults := []string{"openid"}
	cfg = cfg.WithClientCredentialsScopes(scopeDefaults)
	// Configure storage options based on --file-storage flag
	cfg = cfg.WithStorageType(getStorageType()).
		WithStorageName("pingcli")

	// Provide optional suffix so SDK keychain entries align with file names
	profileName, _ := profiles.GetOptionValue(options.RootActiveProfileOption)
	if strings.TrimSpace(profileName) == "" {
		profileName = "default"
	}
	providerName, _ := profiles.GetOptionValue(options.AuthProviderOption)
	if strings.TrimSpace(providerName) == "" {
		providerName = customtypes.ENUM_AUTH_PROVIDER_PINGONE
	}
	cfg = cfg.WithStorageOptionalSuffix(fmt.Sprintf("_%s_%s_%s", providerName, string(svcOAuth2.GrantTypeClientCredentials), profileName))

	// Apply Environment ID for consistent token key generation and endpoints
	environmentID, err := profiles.GetOptionValue(options.PingOneAuthenticationAPIEnvironmentIDOption)
	if err != nil {
		return nil, &errs.PingCLIError{
			Prefix: credentialsErrorPrefix,
			Err:    err,
		}
	}
	if strings.TrimSpace(environmentID) != "" {
		cfg = cfg.WithEnvironmentID(environmentID)
	}

	// Apply region configuration
	cfg, err = applyRegionConfiguration(cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}
