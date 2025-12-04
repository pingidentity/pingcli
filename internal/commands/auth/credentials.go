// Copyright © 2025 Ping Identity Corporation

package auth_internal

import (
	"context"
	"crypto/sha256"
	"errors"
	"fmt"
	"strconv"
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

// shouldUseKeychain checks if keychain storage should be used based on the --file-storage flag
// Returns true if keychain should be used (default), false if file storage should be used
func shouldUseKeychain() bool {
	useFileStorage, err := profiles.GetOptionValue(options.AuthFileStorageOption)
	if err != nil {
		// If we can't get the value, default to true (use keychain)
		return true
	}

	if useFileStorage == "" {
		// If not set, default to true (use keychain)
		return true
	}

	// If --file-storage is true, we should NOT use keychain
	useFileStorageBool, err := strconv.ParseBool(useFileStorage)
	if err != nil {
		// If we can't parse the value, default to true (use keychain)
		return true
	}

	return !useFileStorageBool
}

// getStorageType returns the appropriate storage type for SDK keychain operations
// SDK handles keychain storage, pingcli handles file storage separately
func getStorageType() config.StorageType {
	if shouldUseKeychain() {
		return config.StorageTypeKeychain
	}
	// When keychain is disabled, SDK won't persist tokens - we handle file storage ourselves
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

// StorageLocation indicates where credentials were saved
type StorageLocation struct {
	Keychain bool
	File     bool
}

// LoginResult contains the result of a login operation
type LoginResult struct {
	Token    *oauth2.Token
	NewAuth  bool
	Location StorageLocation
}

// SaveTokenForMethod saves an OAuth2 token to file storage using the specified authentication method key
// Note: SDK handles keychain storage separately with its own token key format
// Returns StorageLocation indicating where the token was saved
func SaveTokenForMethod(token *oauth2.Token, authMethod string) (StorageLocation, error) {
	location := StorageLocation{}

	// Save to file storage
	// Note: SDK handles keychain storage separately with its own token key format
	if err := saveTokenToFile(token, authMethod); err != nil {
		// If it's a critical error (like nil token), fail immediately
		if errors.Is(err, ErrNilToken) {
			return location, err
		}

		return location, err
	}

	location.File = true

	return location, nil
}

// LoadTokenForMethod loads an OAuth2 token from the keychain using the specified authentication method key
// Falls back to file storage if keychain operations fail or if --use-keychain=false
func LoadTokenForMethod(authMethod string) (*oauth2.Token, error) {
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
		// Normalize auth type to snake_case format
		switch authType {
		case "worker":
			authType = "client_credentials"
		case "clientCredentials":
			authType = "client_credentials"
		case "deviceCode":
			authType = "device_code"
		case "authorization_code":
			authType = "authorization_code"
		}

		// Try to get configuration for the configured auth method
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

	// No authentication type configured
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

	// Normalize auth type to snake_case format
	switch authType {
	case "worker":
		authType = "client_credentials"
	case "clientCredentials":
		authType = "client_credentials"
	case "deviceCode":
		authType = "device_code"
	case "authorization_code":
		authType = "authorization_code"
	}

	// Try to get configuration for the configured auth method
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
// Returns StorageLocation indicating what was cleared
func ClearTokenForMethod(authMethod string) (StorageLocation, error) {
	var errList []error
	location := StorageLocation{}

	// Clear from keychain
	storage, err := getTokenStorage(authMethod)
	if err == nil {
		if err := storage.ClearToken(); err != nil {
			errList = append(errList, &errs.PingCLIError{
				Prefix: credentialsErrorPrefix,
				Err:    err,
			})
		} else {
			location.Keychain = true
		}
	}

	// Also clear from file storage
	if err := clearTokenFromFile(authMethod); err != nil {
		errList = append(errList, &errs.PingCLIError{
			Prefix: credentialsErrorPrefix,
			Err:    err,
		})
	} else {
		location.File = true
	}

	// Also clear all token files for this grant type and current profile
	// This handles cases where the user changed their configuration
	// Determine grant type from auth method (authMethod is the token key)
	// We need to parse it to get the grant type
	// For now, just try all grant types - inefficient but safe
	profileName, err := profiles.GetOptionValue(options.RootActiveProfileOption)
	if err != nil {
		profileName = "default"
	}

	providerName, err := profiles.GetOptionValue(options.AuthProviderOption)
	if err != nil || providerName == "" {
		providerName = "pingone" // Default to pingone
	}

	// Try all grant types to make sure we clean up
	grantTypes := []string{
		string(svcOAuth2.GrantTypeDeviceCode),
		string(svcOAuth2.GrantTypeClientCredentials),
		string(svcOAuth2.GrantTypeAuthorizationCode),
	}

	for _, grantType := range grantTypes {
		if err := clearAllTokenFilesForGrantType(providerName, grantType, profileName); err != nil {
			// Don't fail the whole operation if cleanup fails
			errList = append(errList, fmt.Errorf("failed to clear all %s tokens: %w", grantType, err))
		}
	}

	return location, errors.Join(errList...)
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

	// Get client ID for token key generation
	clientID := ""
	if cfg.Auth.DeviceCode != nil && cfg.Auth.DeviceCode.DeviceCodeClientID != nil {
		clientID = *cfg.Auth.DeviceCode.DeviceCodeClientID
	}

	// Get environment ID for token key generation
	environmentID := ""
	if cfg.Endpoint.EnvironmentID != nil {
		environmentID = *cfg.Endpoint.EnvironmentID
	}

	// Set grant type to device code
	cfg = cfg.WithGrantType(svcOAuth2.GrantTypeDeviceCode)

	// Generate unique token key based on provider, profile and configuration
	tokenKey := generateTokenKey(providerName, profileName, environmentID, clientID, string(svcOAuth2.GrantTypeDeviceCode))
	if tokenKey == "" {
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
	_ = clearAllTokenFilesForGrantType(providerName, string(svcOAuth2.GrantTypeDeviceCode), profileName)

	// Save token using our own storage logic (handles both file and keychain based on flags)
	location, err := SaveTokenForMethod(token, tokenKey)
	if err != nil {
		return nil, &errs.PingCLIError{
			Prefix: credentialsErrorPrefix,
			Err:    err,
		}
	}

	// SDK handles keychain storage separately - mark if keychain is enabled
	if shouldUseKeychain() {
		location.Keychain = true
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

	// Get device code scopes (optional)
	scopes, err := profiles.GetOptionValue(options.PingOneAuthenticationDeviceCodeScopesOption)
	if err != nil {
		return nil, &errs.PingCLIError{
			Prefix: credentialsErrorPrefix,
			Err:    err,
		}
	}

	// Configure device code settings
	cfg = cfg.WithDeviceCodeClientID(clientID)

	scopesList := parseScopesList(scopes)
	cfg = cfg.WithDeviceCodeScopes(scopesList)

	// Configure storage options based on --file-storage flag
	cfg = cfg.WithStorageType(getStorageType()).WithStorageName("pingcli")

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

	// Get client ID for token key generation
	clientID := ""
	if cfg.Auth.AuthorizationCode != nil && cfg.Auth.AuthorizationCode.AuthorizationCodeClientID != nil {
		clientID = *cfg.Auth.AuthorizationCode.AuthorizationCodeClientID
	}

	// Get environment ID for token key generation
	environmentID := ""
	if cfg.Endpoint.EnvironmentID != nil {
		environmentID = *cfg.Endpoint.EnvironmentID
	}

	// Set grant type to authorization code
	cfg = cfg.WithGrantType(svcOAuth2.GrantTypeAuthorizationCode)

	// Generate unique token key based on provider, profile and configuration
	tokenKey := generateTokenKey(providerName, profileName, environmentID, clientID, string(svcOAuth2.GrantTypeAuthorizationCode))
	if tokenKey == "" {
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
	_ = clearAllTokenFilesForGrantType(providerName, string(svcOAuth2.GrantTypeAuthorizationCode), profileName)

	// Save token using our own storage logic (handles both file and keychain based on flags)
	location, err := SaveTokenForMethod(token, tokenKey)
	if err != nil {
		return nil, &errs.PingCLIError{
			Prefix: credentialsErrorPrefix,
			Err:    err,
		}
	}

	// SDK handles keychain storage separately - mark if keychain is enabled
	if shouldUseKeychain() {
		location.Keychain = true
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

	// Get auth code scopes (optional)
	scopes, err := profiles.GetOptionValue(options.PingOneAuthenticationAuthorizationCodeScopesOption)
	if err != nil {
		return nil, &errs.PingCLIError{
			Prefix: credentialsErrorPrefix,
			Err:    err,
		}
	}

	// Configure auth code settings
	cfg = cfg.WithAuthorizationCodeClientID(clientID).
		WithAuthorizationCodeRedirectURI(redirectURI)

	scopesList := parseScopesList(scopes)
	if len(scopesList) > 0 {
		cfg = cfg.WithAuthorizationCodeScopes(scopesList)
	}

	// Configure storage options based on --file-storage flag
	cfg = cfg.WithStorageType(getStorageType()).
		WithStorageName("pingcli")

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

	// Get configuration values for token key generation
	clientID := ""
	if cfg.Auth.ClientCredentials != nil {
		if cfg.Auth.ClientCredentials.ClientCredentialsClientID != nil {
			clientID = *cfg.Auth.ClientCredentials.ClientCredentialsClientID
		}
	}
	environmentID := ""
	if cfg.Endpoint.EnvironmentID != nil {
		environmentID = *cfg.Endpoint.EnvironmentID
	}

	// Set grant type to client credentials
	cfg = cfg.WithGrantType(svcOAuth2.GrantTypeClientCredentials)

	// Generate unique token key based on provider, profile and configuration
	tokenKey := generateTokenKey(providerName, profileName, environmentID, clientID, string(svcOAuth2.GrantTypeClientCredentials))
	if tokenKey == "" {
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

	// SDK handles keychain storage separately - mark if keychain is enabled
	if shouldUseKeychain() {
		location.Keychain = true
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

	// Get client credentials scopes (optional)
	scopes, err := profiles.GetOptionValue(options.PingOneAuthenticationClientCredentialsScopesOption)
	if err != nil {
		return nil, &errs.PingCLIError{
			Prefix: credentialsErrorPrefix,
			Err:    err,
		}
	}

	// Configure client credentials settings
	cfg = cfg.WithClientCredentialsClientID(clientID).
		WithClientCredentialsClientSecret(clientSecret)

	scopesList := parseScopesList(scopes)
	if len(scopesList) > 0 {
		cfg = cfg.WithClientCredentialsScopes(scopesList)
	}

	// Configure storage options based on --file-storage flag
	cfg = cfg.WithStorageType(getStorageType()).
		WithStorageName("pingcli")

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

	// Configure worker settings
	cfg = cfg.WithClientCredentialsClientID(clientID).
		WithClientCredentialsClientSecret(clientSecret)
	// Configure storage options based on --file-storage flag
	cfg = cfg.WithStorageType(getStorageType()).
		WithStorageName("pingcli")

	// Apply region configuration
	cfg, err = applyRegionConfiguration(cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}
