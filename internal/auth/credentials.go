// Copyright © 2025 Ping Identity Corporation

package auth_internal

import (
	"context"
	"errors"
	"fmt"

	"github.com/pingidentity/pingcli/internal/configuration/options"
	"github.com/pingidentity/pingcli/internal/errs"
	"github.com/pingidentity/pingcli/internal/profiles"
	"github.com/pingidentity/pingone-go-client/config"
	svcOAuth2 "github.com/pingidentity/pingone-go-client/oauth2"
	"golang.org/x/oauth2"
)

// Token storage keys for different authentication methods
const (
	deviceCodeTokenKey        = "device-code-token"
	authCodeTokenKey          = "auth-code-token" // #nosec G101 -- This is a keychain identifier, not a credential
	clientCredentialsTokenKey = "client-credentials-token"
)

// getTokenStorage returns the appropriate keychain storage instance for the given authentication method
func getTokenStorage(authMethod string) *svcOAuth2.KeychainStorage {
	return svcOAuth2.NewKeychainStorage("pingcli", authMethod)
}

// shouldUseKeychain checks if keychain storage should be used based on the --files-option flag
func shouldUseKeychain() bool {
	useKeychain, err := profiles.GetOptionValue(options.AuthFileStorageOption)
	if err != nil {
		// If we can't get the value, default to true (use keychain)
		return true
	}

	if useKeychain == "" {
		// If not set, default to true
		return false
	}

	// Parse the string value to bool
	return useKeychain == "true"
}

// SaveTokenForMethod saves an OAuth2 token to the keychain using the specified authentication method key
// Falls back to file storage if keychain operations fail or if --use-keychain=false
func SaveTokenForMethod(token *oauth2.Token, authMethod string) error {
	// Check if user disabled keychain
	if !shouldUseKeychain() {
		// Directly save to file storage
		return saveTokenToFile(token, authMethod)
	}

	storage := getTokenStorage(authMethod)

	// Try keychain storage first
	err := storage.SaveToken(token)
	if err != nil {
		// Fallback to file storage if keychain fails
		if fileErr := saveTokenToFile(token, authMethod); fileErr != nil {
			return fmt.Errorf("failed to save token to keychain (%w) and file storage (%w)", err, fileErr)
		}
		// Token saved to file successfully
		return nil
	}

	return nil
}

// LoadTokenForMethod loads an OAuth2 token from the keychain using the specified authentication method key
// Falls back to file storage if keychain operations fail or if --use-keychain=false
func LoadTokenForMethod(authMethod string) (*oauth2.Token, error) {
	// Check if user disabled keychain
	if !shouldUseKeychain() {
		// Directly load from file storage
		return loadTokenFromFile(authMethod)
	}

	storage := getTokenStorage(authMethod)

	// Try keychain storage first
	token, err := storage.LoadToken()
	if err != nil {
		// Fallback to file storage if keychain fails
		token, fileErr := loadTokenFromFile(authMethod)
		if fileErr != nil {
			return nil, fmt.Errorf("failed to load token from keychain (%w) and file storage (%w)", err, fileErr)
		}

		return token, nil
	}

	return token, nil
}

// SaveToken saves an OAuth2 token using device code storage for backward compatibility with older versions
func SaveToken(token *oauth2.Token) error {
	return SaveTokenForMethod(token, deviceCodeTokenKey)
}

// LoadToken attempts to load an OAuth2 token from the keychain, trying configured auth methods first,
// then falling back to legacy storage for backward compatibility
func LoadToken() (*oauth2.Token, error) {
	// First, try to load using configuration-based keys from the active profile
	authType, err := profiles.GetOptionValue(options.PingOneAuthenticationTypeOption)
	if err == nil && authType != "" {
		// Normalize "worker" to "client_credentials" for token loading
		if authType == "worker" {
			authType = "client_credentials"
		}

		// Try to get configuration for the configured auth method
		var cfg *config.Configuration
		var grantType svcOAuth2.GrantType
		switch authType {
		case "device_code":
			cfg, _ = GetDeviceCodeConfiguration()
			grantType = svcOAuth2.GrantTypeDeviceCode
		case "auth_code":
			cfg, _ = GetAuthCodeConfiguration()
			grantType = svcOAuth2.GrantTypeAuthCode
		case "client_credentials":
			cfg, _ = GetClientCredentialsConfiguration()
			grantType = svcOAuth2.GrantTypeClientCredentials
		}

		if cfg != nil {
			// Set the grant type before generating the token key
			cfg = cfg.WithGrantType(grantType)
			tokenKey, err := GetAuthMethodKeyFromConfig(cfg)
			if err == nil {
				keychainStorage := svcOAuth2.NewKeychainStorage("pingcli", tokenKey)
				token, err := keychainStorage.LoadToken()
				if err == nil && token != nil {
					return token, nil
				}
			}
		}
	}

	// Also try all configured auth methods in case the type doesn't match
	authMethods := []struct {
		name      string
		getConfig func() (*config.Configuration, error)
		grantType svcOAuth2.GrantType
	}{
		{"client_credentials", GetClientCredentialsConfiguration, svcOAuth2.GrantTypeClientCredentials},
		{"device_code", GetDeviceCodeConfiguration, svcOAuth2.GrantTypeDeviceCode},
		{"auth_code", GetAuthCodeConfiguration, svcOAuth2.GrantTypeAuthCode},
	}

	for _, method := range authMethods {
		cfg, err := method.getConfig()
		if err == nil && cfg != nil {
			// Set the grant type before generating the token key
			cfg = cfg.WithGrantType(method.grantType)
			tokenKey, err := GetAuthMethodKeyFromConfig(cfg)
			if err == nil {
				keychainStorage := svcOAuth2.NewKeychainStorage("pingcli", tokenKey)
				token, err := keychainStorage.LoadToken()
				if err == nil && token != nil {
					return token, nil
				}
			}
		}
	}

	// Fall back to legacy token loading for backward compatibility
	methods := []string{deviceCodeTokenKey, authCodeTokenKey, clientCredentialsTokenKey}

	for _, method := range methods {
		token, err := LoadTokenForMethod(method)
		if err == nil && token != nil {
			return token, nil
		}
	}

	return nil, ErrNoTokenFound
}

// GetTokenSource returns an OAuth2 token source from the cached token in the keychain
func GetTokenSource(ctx context.Context) (oauth2.TokenSource, error) {
	// Try to load cached token
	cachedToken, err := LoadToken()
	if err != nil {
		return nil, fmt.Errorf("failed to load cached token: %w", err)
	}

	if cachedToken == nil {
		return nil, ErrNoCachedToken
	}

	// Return a simple static token source - SDK handles complex refresh during login
	return oauth2.StaticTokenSource(cachedToken), nil
}

// ClearToken removes all cached tokens from the keychain for all authentication methods,
// including both configuration-based and legacy token storage, and file storage
func ClearToken() error {
	var errs []error

	// Clear configuration-based tokens for all auth methods
	authMethods := []struct {
		name      string
		getConfig func() (*config.Configuration, error)
		grantType svcOAuth2.GrantType
	}{
		{"client_credentials", GetClientCredentialsConfiguration, svcOAuth2.GrantTypeClientCredentials},
		{"device_code", GetDeviceCodeConfiguration, svcOAuth2.GrantTypeDeviceCode},
		{"auth_code", GetAuthCodeConfiguration, svcOAuth2.GrantTypeAuthCode},
	}

	for _, method := range authMethods {
		cfg, err := method.getConfig()
		if err == nil && cfg != nil {
			// Set the grant type before generating the token key
			cfg = cfg.WithGrantType(method.grantType)
			tokenKey, err := GetAuthMethodKeyFromConfig(cfg)
			if err == nil {
				keychainStorage := svcOAuth2.NewKeychainStorage("pingcli", tokenKey)
				if err := keychainStorage.ClearToken(); err != nil {
					errs = append(errs, err)
				}
				// Also clear from file storage
				if err := clearTokenFromFile(tokenKey); err != nil {
					errs = append(errs, err)
				}
			}
		}
	}

	// Also clear legacy tokens from all auth methods for backward compatibility
	methods := []string{deviceCodeTokenKey, authCodeTokenKey, clientCredentialsTokenKey}

	for _, method := range methods {
		storage := getTokenStorage(method)
		if err := storage.ClearToken(); err != nil {
			errs = append(errs, err)
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
func ClearTokenForMethod(authMethod string) error {
	var errs []error

	// Clear from keychain
	storage := getTokenStorage(authMethod)
	if err := storage.ClearToken(); err != nil {
		errs = append(errs, fmt.Errorf("keychain clear failed: %w", err))
	}

	// Also clear from file storage
	if err := clearTokenFromFile(authMethod); err != nil {
		errs = append(errs, fmt.Errorf("file clear failed: %w", err))
	}

	return errors.Join(errs...)
}

// PerformDeviceCodeLogin performs device code authentication, returning the token, whether it's new authentication,
// and any error encountered during the process
func PerformDeviceCodeLogin(ctx context.Context) (*oauth2.Token, bool, error) {
	cfg, err := GetDeviceCodeConfiguration()
	if err != nil {
		return nil, false, fmt.Errorf("failed to get device code configuration: %w", err)
	}

	// Set grant type to device code
	cfg = cfg.WithGrantType(svcOAuth2.GrantTypeDeviceCode)

	// Load any existing cached token before calling SDK (checks both keychain and file storage)
	tokenKey, err := GetAuthMethodKeyFromConfig(cfg)
	var existingToken *oauth2.Token
	if err == nil {
		// Use LoadTokenForMethod which handles both keychain and file storage
		existingToken, _ = LoadTokenForMethod(tokenKey)

		// If we have a valid token, return it without performing new authentication
		if existingToken != nil && existingToken.Valid() {
			return existingToken, false, nil
		}
	}

	// Get token source to perform authentication or use cached token
	tokenSource, err := cfg.TokenSource(ctx)
	if err != nil {
		return nil, false, fmt.Errorf("failed to get token source: %w", err)
	}

	// Get token (SDK will return cached token if valid, or perform new authentication)
	token, err := tokenSource.Token()
	if err != nil {
		return nil, false, fmt.Errorf("failed to get token: %w", err)
	}

	// Save token using our storage method (respects --use-keychain flag)
	if tokenKey != "" {
		_ = SaveTokenForMethod(token, tokenKey)
	}

	// Determine if this was new authentication by comparing with what we loaded
	newAuth := existingToken == nil || existingToken.AccessToken != token.AccessToken

	return token, newAuth, nil
}

// GetDeviceCodeConfiguration builds a device code authentication configuration from the CLI profile options
func GetDeviceCodeConfiguration() (*config.Configuration, error) {
	cfg := config.NewConfiguration()

	// Get device code client ID
	clientID, err := profiles.GetOptionValue(options.PingOneAuthenticationDeviceCodeClientIDOption)
	if err != nil {
		return nil, fmt.Errorf("failed to get device code client ID: %w", err)
	}
	if clientID == "" {
		return nil, &errs.PingCLIError{
			Prefix: "failed to get device code configuration",
			Err:    ErrDeviceCodeClientIDNotConfigured,
		}
	}

	// Get device code environment ID
	environmentID, err := profiles.GetOptionValue(options.PingOneAuthenticationDeviceCodeEnvironmentIDOption)
	if err != nil {
		return nil, fmt.Errorf("failed to get device code environment ID: %w", err)
	}
	if environmentID == "" {
		return nil, &errs.PingCLIError{
			Prefix: "failed to get device code configuration",
			Err:    ErrDeviceCodeEnvironmentIDNotConfigured,
		}
	}

	// Get device code scopes (optional)
	scopes, err := profiles.GetOptionValue(options.PingOneAuthenticationDeviceCodeScopesOption)
	if err != nil {
		return nil, fmt.Errorf("failed to get device code scopes: %w", err)
	}

	// Configure device code settings
	cfg = cfg.WithDeviceCodeClientID(clientID).
		WithDeviceCodeEnvironmentID(environmentID)

	scopesList := parseScopesList(scopes)
	cfg = cfg.WithDeviceCodeScopes(scopesList)

	// Apply region configuration
	return applyRegionConfiguration(cfg)
}

func PerformAuthCodeLogin(ctx context.Context) (*oauth2.Token, bool, error) {
	cfg, err := GetAuthCodeConfiguration()
	if err != nil {
		return nil, false, fmt.Errorf("failed to get auth code configuration: %w", err)
	}

	// Set grant type to auth code
	cfg = cfg.WithGrantType(svcOAuth2.GrantTypeAuthCode)

	// Load any existing cached token before calling SDK (checks both keychain and file storage)
	tokenKey, err := GetAuthMethodKeyFromConfig(cfg)
	var existingToken *oauth2.Token
	if err == nil {
		// Use LoadTokenForMethod which handles both keychain and file storage
		existingToken, _ = LoadTokenForMethod(tokenKey)

		// If we have a valid token, return it without performing new authentication
		if existingToken != nil && existingToken.Valid() {
			return existingToken, false, nil
		}
	}

	// Get token source to perform authentication or use cached token
	tokenSource, err := cfg.TokenSource(ctx)
	if err != nil {
		return nil, false, fmt.Errorf("failed to get token source: %w", err)
	}

	// Get token (SDK will return cached token if valid, or perform new authentication)
	token, err := tokenSource.Token()
	if err != nil {
		return nil, false, fmt.Errorf("failed to get token: %w", err)
	}

	// Save token using our storage method (respects --use-keychain flag)
	if tokenKey != "" {
		_ = SaveTokenForMethod(token, tokenKey)
	}

	// Determine if this was new authentication by comparing with what we loaded
	newAuth := existingToken == nil || existingToken.AccessToken != token.AccessToken

	return token, newAuth, nil
}

// GetAuthCodeConfiguration builds an authorization code authentication configuration from the CLI profile options
func GetAuthCodeConfiguration() (*config.Configuration, error) {
	cfg := config.NewConfiguration()

	// Get auth code client ID
	clientID, err := profiles.GetOptionValue(options.PingOneAuthenticationAuthCodeClientIDOption)
	if err != nil {
		return nil, fmt.Errorf("failed to get auth code client ID: %w", err)
	}
	if clientID == "" {
		return nil, &errs.PingCLIError{
			Prefix: "failed to get auth code configuration",
			Err:    ErrAuthCodeClientIDNotConfigured,
		}
	}

	// Get auth code environment ID
	environmentID, err := profiles.GetOptionValue(options.PingOneAuthenticationAuthCodeEnvironmentIDOption)
	if err != nil {
		return nil, fmt.Errorf("failed to get auth code environment ID: %w", err)
	}
	if environmentID == "" {
		return nil, &errs.PingCLIError{
			Prefix: "failed to get auth code configuration",
			Err:    ErrAuthCodeEnvironmentIDNotConfigured,
		}
	}

	// Get auth code redirect URI
	redirectURIPath, err := profiles.GetOptionValue(options.PingOneAuthenticationAuthCodeRedirectURIPathOption)
	if err != nil {
		return nil, fmt.Errorf("failed to get auth code redirect URI: %w", err)
	}
	if redirectURIPath == "" {
		return nil, &errs.PingCLIError{
			Prefix: "failed to get auth code configuration",
			Err:    ErrAuthCodeRedirectURINotConfigured,
		}
	}

	redirectURIPort, err := profiles.GetOptionValue(options.PingOneAuthenticationAuthCodeRedirectURIPortOption)
	if err != nil && redirectURIPort != "" {
		return nil, fmt.Errorf("failed to get auth code redirect URI port: %w", err)
	}

	redirectURI := config.AuthCodeRedirectURI{
		Port: redirectURIPort,
		Path: redirectURIPath,
	}

	// Get auth code scopes (optional)
	scopes, err := profiles.GetOptionValue(options.PingOneAuthenticationAuthCodeScopesOption)
	if err != nil {
		return nil, fmt.Errorf("failed to get auth code scopes: %w", err)
	}

	// Configure auth code settings
	cfg = cfg.WithAuthCodeClientID(clientID).
		WithAuthCodeEnvironmentID(environmentID).
		WithAuthCodeRedirectURI(redirectURI)

	scopesList := parseScopesList(scopes)
	if len(scopesList) > 0 {
		cfg = cfg.WithAuthCodeScopes(scopesList)
	}

	// Apply region configuration
	return applyRegionConfiguration(cfg)
}

func PerformClientCredentialsLogin(ctx context.Context) (*oauth2.Token, bool, error) {
	cfg, err := GetClientCredentialsConfiguration()
	if err != nil {
		return nil, false, fmt.Errorf("failed to get client credentials configuration: %w", err)
	}

	// Set grant type to client credentials
	cfg = cfg.WithGrantType(svcOAuth2.GrantTypeClientCredentials)

	// Load any existing cached token before calling SDK (checks both keychain and file storage)
	tokenKey, err := GetAuthMethodKeyFromConfig(cfg)
	var existingToken *oauth2.Token
	if err == nil {
		// Use LoadTokenForMethod which handles both keychain and file storage
		existingToken, _ = LoadTokenForMethod(tokenKey)

		// If we have a valid token, return it without performing new authentication
		if existingToken != nil && existingToken.Valid() {
			return existingToken, false, nil
		}
	}

	// Get token source to perform authentication or use cached token
	tokenSource, err := cfg.TokenSource(ctx)
	if err != nil {
		return nil, false, fmt.Errorf("failed to get token source: %w", err)
	}

	// Get token (SDK will return cached token if valid, or perform new authentication)
	token, err := tokenSource.Token()
	if err != nil {
		return nil, false, fmt.Errorf("failed to get token: %w", err)
	}

	// Save token using our storage method (respects --use-keychain flag)
	if tokenKey != "" {
		_ = SaveTokenForMethod(token, tokenKey)
	}

	// Determine if this was new authentication by comparing with what we loaded
	newAuth := existingToken == nil || existingToken.AccessToken != token.AccessToken

	return token, newAuth, nil
}

// GetClientCredentialsConfiguration builds a client credentials authentication configuration from the CLI profile options
func GetClientCredentialsConfiguration() (*config.Configuration, error) {
	cfg := config.NewConfiguration()

	// Get client credentials client ID
	clientID, err := profiles.GetOptionValue(options.PingOneAuthenticationClientCredentialsClientIDOption)
	if err != nil {
		return nil, fmt.Errorf("failed to get client credentials client ID: %w", err)
	}
	if clientID == "" {
		return nil, &errs.PingCLIError{
			Prefix: "failed to get client credentials configuration",
			Err:    ErrClientCredentialsClientIDNotConfigured,
		}
	}

	// Get client credentials client secret
	clientSecret, err := profiles.GetOptionValue(options.PingOneAuthenticationClientCredentialsClientSecretOption)
	if err != nil {
		return nil, fmt.Errorf("failed to get client credentials client secret: %w", err)
	}
	if clientSecret == "" {
		return nil, &errs.PingCLIError{
			Prefix: "failed to get client credentials configuration",
			Err:    ErrClientCredentialsClientSecretNotConfigured,
		}
	}

	// Get client credentials environment ID
	environmentID, err := profiles.GetOptionValue(options.PingOneAuthenticationClientCredentialsEnvironmentIDOption)
	if err != nil {
		return nil, fmt.Errorf("failed to get client credentials environment ID: %w", err)
	}
	if environmentID == "" {
		return nil, &errs.PingCLIError{
			Prefix: "failed to get client credentials configuration",
			Err:    ErrClientCredentialsEnvironmentIDNotConfigured,
		}
	}

	// Get client credentials scopes (optional)
	scopes, err := profiles.GetOptionValue(options.PingOneAuthenticationClientCredentialsScopesOption)
	if err != nil {
		return nil, fmt.Errorf("failed to get client credentials scopes: %w", err)
	}

	// Configure client credentials settings
	cfg = cfg.WithClientCredentialsClientID(clientID).
		WithClientCredentialsClientSecret(clientSecret).
		WithClientCredentialsEnvironmentID(environmentID)

	scopesList := parseScopesList(scopes)
	if len(scopesList) > 0 {
		cfg = cfg.WithClientCredentialsScopes(scopesList)
	}

	// Apply region configuration
	return applyRegionConfiguration(cfg)
}

// GetValidTokenSource returns a token source with a valid token, attempting to use cached tokens first
// and performing automatic authentication for client_credentials if needed
func GetValidTokenSource(ctx context.Context) (oauth2.TokenSource, error) {
	// Try to load cached token first
	cachedToken, err := LoadToken()
	if err == nil && cachedToken != nil && cachedToken.Valid() {
		// Return cached token if it's still valid
		return oauth2.StaticTokenSource(cachedToken), nil
	}

	// No valid cached token found - attempt automatic authentication based on configured method
	authType, err := profiles.GetOptionValue(options.PingOneAuthenticationTypeOption)
	if err != nil {
		return nil, fmt.Errorf("failed to get authentication type: %w", err)
	}

	// Automatically authenticate based on configured method
	var token *oauth2.Token
	var newAuth bool

	switch authType {
	case "device_code":
		token, newAuth, err = PerformDeviceCodeLogin(ctx)
		if err != nil {
			return nil, fmt.Errorf("automatic device code authentication failed: %w", err)
		}
	case "auth_code":
		token, newAuth, err = PerformAuthCodeLogin(ctx)
		if err != nil {
			return nil, fmt.Errorf("automatic authorization code authentication failed: %w", err)
		}
	case "client_credentials":
		token, newAuth, err = PerformClientCredentialsLogin(ctx)
		if err != nil {
			return nil, fmt.Errorf("automatic client credentials authentication failed: %w", err)
		}
	default:
		return nil, &errs.PingCLIError{
			Prefix: fmt.Sprintf("automatic authentication failed for type '%s'", authType),
			Err:    ErrUnsupportedAuthType,
		}
	}

	if newAuth {
		fmt.Printf("Successfully authenticated using %s authentication\n", authType)
	}

	// Return a static token source with the obtained token
	return oauth2.StaticTokenSource(token), nil
}
