// Copyright © 2025 Ping Identity Corporation

package auth_internal

import (
	"context"
	"errors"
	"fmt"

	"github.com/pingidentity/pingcli/internal/configuration/options"
	"github.com/pingidentity/pingcli/internal/profiles"
	"github.com/pingidentity/pingone-go-client/config"
	svcOAuth2 "github.com/pingidentity/pingone-go-client/oauth2"
	"golang.org/x/oauth2"
)

// No init() function needed - direct integration is cleaner

// Token storage keys for different authentication methods
const (
	deviceCodeTokenKey        = "device-code-token"
	authCodeTokenKey          = "auth-code-token" // #nosec G101 -- This is a keychain identifier, not a credential
	clientCredentialsTokenKey = "client-credentials-token"
)

// getTokenStorage returns the appropriate token storage for the given auth method
func getTokenStorage(authMethod string) *svcOAuth2.KeychainStorage {
	return svcOAuth2.NewKeychainStorage("pingcli", authMethod)
}

// SaveToken saves an OAuth2 token using the SDK keychain storage for the specific auth method
func SaveTokenForMethod(token *oauth2.Token, authMethod string) error {
	storage := getTokenStorage(authMethod)

	return storage.SaveToken(token)
}

// LoadTokenForMethod loads an OAuth2 token using the SDK keychain storage for the specific auth method
func LoadTokenForMethod(authMethod string) (*oauth2.Token, error) {
	storage := getTokenStorage(authMethod)

	return storage.LoadToken()
}

// SaveToken saves an OAuth2 token using device code storage (for backward compatibility)
func SaveToken(token *oauth2.Token) error {
	return SaveTokenForMethod(token, deviceCodeTokenKey)
}

// LoadToken loads an OAuth2 token using device code storage (for backward compatibility)
func LoadToken() (*oauth2.Token, error) {
	// Try to load from all auth methods, starting with device code for backward compatibility
	methods := []string{deviceCodeTokenKey, authCodeTokenKey, clientCredentialsTokenKey}

	for _, method := range methods {
		token, err := LoadTokenForMethod(method)
		if err == nil && token != nil {
			return token, nil
		}
	}

	return nil, fmt.Errorf("no token found for any authentication method")
}

// GetTokenSource returns an OAuth2 token source from cached token
func GetTokenSource(ctx context.Context) (oauth2.TokenSource, error) {
	// Try to load cached token
	cachedToken, err := LoadToken()
	if err != nil {
		return nil, fmt.Errorf("failed to load cached token: %w", err)
	}

	if cachedToken == nil {
		return nil, fmt.Errorf("no cached token available")
	}

	// Return a simple static token source - SDK handles complex refresh during login
	return oauth2.StaticTokenSource(cachedToken), nil
}

// ClearToken removes the cached token using the SDK keychain storage
func ClearToken() error {
	// Clear tokens from all auth methods
	methods := []string{deviceCodeTokenKey, authCodeTokenKey, clientCredentialsTokenKey}

	var errs []error
	for _, method := range methods {
		storage := getTokenStorage(method)
		if err := storage.ClearToken(); err != nil {
			errs = append(errs, err)
		}
	}

	// Also clear the cached PingOne API client to force re-initialization
	ClearPingOneClientCache()

	return errors.Join(errs...)
}

// ClearTokenForMethod removes the cached token for a specific authentication method
func ClearTokenForMethod(authMethod string) error {
	storage := getTokenStorage(authMethod)
	if err := storage.ClearToken(); err != nil {
		return err
	}

	// Also clear the cached PingOne API client to force re-initialization
	ClearPingOneClientCache()

	return nil
}

func PerformDeviceCodeLogin(ctx context.Context) (*oauth2.Token, bool, error) {
	cfg, err := GetDeviceCodeConfiguration()
	if err != nil {
		return nil, false, fmt.Errorf("failed to get device code configuration: %w", err)
	}

	// Set grant type to device code
	cfg = cfg.WithGrantType(svcOAuth2.GrantTypeDeviceCode)

	// Load any existing cached token before calling SDK
	tokenKey, err := GetAuthMethodKeyFromConfig(cfg)
	var existingToken *oauth2.Token
	if err == nil {
		keychainStorage := svcOAuth2.NewKeychainStorage("pingcli", tokenKey)
		existingToken, _ = keychainStorage.LoadToken()
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

	// Determine if this was new authentication by comparing with what we loaded
	newAuth := existingToken == nil || existingToken.AccessToken != token.AccessToken

	return token, newAuth, nil
}

// GetDeviceCodeConfiguration builds device code configuration from CLI options
func GetDeviceCodeConfiguration() (*config.Configuration, error) {
	cfg := config.NewConfiguration()

	// Get device code client ID
	clientID, err := profiles.GetOptionValue(options.PingOneAuthenticationDeviceCodeClientIDOption)
	if err != nil {
		return nil, fmt.Errorf("failed to get device code client ID: %w", err)
	}
	if clientID == "" {
		return nil, fmt.Errorf("device code client ID is not configured. Please run 'pingcli config set service.pingone.authentication.deviceCode.clientID=<your-client-id>'")
	}

	// Get device code environment ID
	environmentID, err := profiles.GetOptionValue(options.PingOneAuthenticationDeviceCodeEnvironmentIDOption)
	if err != nil {
		return nil, fmt.Errorf("failed to get device code environment ID: %w", err)
	}
	if environmentID == "" {
		return nil, fmt.Errorf("device code environment ID is not configured. Please run 'pingcli config set service.pingone.authentication.deviceCode.environmentID=<your-env-id>'")
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

	// Check if we already have a valid cached token by generating the same key the SDK would use
	tokenKey, err := GetAuthMethodKeyFromConfig(cfg)
	var existingToken *oauth2.Token
	if err == nil {
		// Try to load existing token using the hash-based key
		keychainStorage := svcOAuth2.NewKeychainStorage("pingcli", tokenKey)
		existingToken, _ = keychainStorage.LoadToken()
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

	// Determine if this was new authentication by comparing with what we loaded
	newAuth := existingToken == nil || existingToken.AccessToken != token.AccessToken

	return token, newAuth, nil
}

// GetAuthCodeConfiguration builds auth code configuration from CLI options
func GetAuthCodeConfiguration() (*config.Configuration, error) {
	cfg := config.NewConfiguration()

	// Get auth code client ID
	clientID, err := profiles.GetOptionValue(options.PingOneAuthenticationAuthCodeClientIDOption)
	if err != nil {
		return nil, fmt.Errorf("failed to get auth code client ID: %w", err)
	}
	if clientID == "" {
		return nil, fmt.Errorf("auth code client ID is not configured. Please run 'pingcli config set service.pingone.authentication.authCode.clientID=<your-client-id>'")
	}

	// Get auth code environment ID
	environmentID, err := profiles.GetOptionValue(options.PingOneAuthenticationAuthCodeEnvironmentIDOption)
	if err != nil {
		return nil, fmt.Errorf("failed to get auth code environment ID: %w", err)
	}
	if environmentID == "" {
		return nil, fmt.Errorf("auth code environment ID is not configured. Please run 'pingcli config set service.pingone.authentication.authCode.environmentID=<your-env-id>'")
	}

	// Get auth code redirect URI
	redirectURI, err := profiles.GetOptionValue(options.PingOneAuthenticationAuthCodeRedirectURIOption)
	if err != nil {
		return nil, fmt.Errorf("failed to get auth code redirect URI: %w", err)
	}
	if redirectURI == "" {
		return nil, fmt.Errorf("auth code redirect URI is not configured. Please run 'pingcli config set service.pingone.authentication.authCode.redirectURI=<your-redirect-uri>'")
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

	// Load any existing cached token before calling SDK
	tokenKey, err := GetAuthMethodKeyFromConfig(cfg)
	var existingToken *oauth2.Token
	if err == nil {
		keychainStorage := svcOAuth2.NewKeychainStorage("pingcli", tokenKey)
		existingToken, _ = keychainStorage.LoadToken()
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

	// Determine if this was new authentication by comparing with what we loaded
	newAuth := existingToken == nil || existingToken.AccessToken != token.AccessToken

	return token, newAuth, nil
}

// GetClientCredentialsConfiguration builds client credentials configuration from CLI options
func GetClientCredentialsConfiguration() (*config.Configuration, error) {
	cfg := config.NewConfiguration()

	// Get client credentials client ID
	clientID, err := profiles.GetOptionValue(options.PingOneAuthenticationClientCredentialsClientIDOption)
	if err != nil {
		return nil, fmt.Errorf("failed to get client credentials client ID: %w", err)
	}
	if clientID == "" {
		return nil, fmt.Errorf("client credentials client ID is not configured. Please run 'pingcli config set service.pingone.authentication.clientCredentials.clientID=<your-client-id>'")
	}

	// Get client credentials client secret
	clientSecret, err := profiles.GetOptionValue(options.PingOneAuthenticationClientCredentialsClientSecretOption)
	if err != nil {
		return nil, fmt.Errorf("failed to get client credentials client secret: %w", err)
	}
	if clientSecret == "" {
		return nil, fmt.Errorf("client credentials client secret is not configured. Please run 'pingcli config set service.pingone.authentication.clientCredentials.clientSecret=<your-client-secret>'")
	}

	// Get client credentials environment ID
	environmentID, err := profiles.GetOptionValue(options.PingOneAuthenticationClientCredentialsEnvironmentIDOption)
	if err != nil {
		return nil, fmt.Errorf("failed to get client credentials environment ID: %w", err)
	}
	if environmentID == "" {
		return nil, fmt.Errorf("client credentials environment ID is not configured. Please run 'pingcli config set service.pingone.authentication.clientCredentials.environmentID=<your-env-id>'")
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

// GetValidTokenSource returns a token source with valid tokens, using cache and refresh as needed
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
		return nil, fmt.Errorf("unsupported authentication type '%s'. Please run 'pingcli login' to authenticate", authType)
	}

	if newAuth {
		fmt.Printf("Successfully authenticated using %s authentication\n", authType)
	}

	// Return a static token source with the obtained token
	return oauth2.StaticTokenSource(token), nil
}
