// Copyright © 2025 Ping Identity Corporation

package auth_internal

import (
	"fmt"
	"strings"

	"github.com/pingidentity/pingcli/internal/configuration/options"
	"github.com/pingidentity/pingcli/internal/errs"
	"github.com/pingidentity/pingcli/internal/profiles"
	"github.com/pingidentity/pingone-go-client/config"
	svcOAuth2 "github.com/pingidentity/pingone-go-client/oauth2"
	"golang.org/x/oauth2"
)

var (
	tokenManagerErrorPrefix = "failed to manage token"
)

// TokenManager defines the interface for managing OAuth2 tokens in the keychain
type TokenManager interface {
	SaveToken(token *oauth2.Token) error
	LoadToken() (*oauth2.Token, error)
	ClearToken() error
	HasToken() bool
}

// DefaultTokenManager implements the TokenManager interface using the default pingcli keychain service
type DefaultTokenManager struct {
	serviceName string
}

// NewDefaultTokenManager creates a new DefaultTokenManager instance
func NewDefaultTokenManager() TokenManager {
	return &DefaultTokenManager{
		serviceName: "pingcli",
	}
}

// GetCurrentAuthMethod returns the configured authentication method key for the active profile
func GetCurrentAuthMethod() (string, error) {
	authMethod, err := profiles.GetOptionValue(options.PingOneAuthenticationTypeOption)
	if err != nil {
		return "", fmt.Errorf("failed to get current grant type: %w", err)
	}

	if authMethod == "" {
		return "", ErrAuthMethodNotConfigured
	}

	return GetAuthMethodKey(authMethod)
}

// GetAuthMethodKey generates a unique keychain account name for the given authentication method
// using the environment ID and client ID from the profile configuration
func GetAuthMethodKey(authMethod string) (string, error) {
	// Get configuration for the grant type to extract environment ID and client ID
	var cfg *config.Configuration
	var err error
	var grantType svcOAuth2.GrantType

	switch authMethod {
	case "device_code":
		cfg, err = GetDeviceCodeConfiguration()
		if err != nil {
			return "", fmt.Errorf("failed to get device code configuration: %w", err)
		}
		grantType = svcOAuth2.GrantTypeDeviceCode
	case "authorization_code":
		cfg, err = GetAuthorizationCodeConfiguration()
		if err != nil {
			return "", fmt.Errorf("failed to get auth code configuration: %w", err)
		}
		grantType = svcOAuth2.GrantTypeAuthorizationCode
	case "client_credentials":
		cfg, err = GetClientCredentialsConfiguration()
		if err != nil {
			return "", fmt.Errorf("failed to get client credentials configuration: %w", err)
		}
		grantType = svcOAuth2.GrantTypeClientCredentials
	case "worker":
		cfg, err = GetWorkerConfiguration()
		if err != nil {
			return "", fmt.Errorf("failed to get worker configuration: %w", err)
		}
		grantType = svcOAuth2.GrantTypeClientCredentials
	default:
		return "", &errs.PingCLIError{
			Prefix: tokenManagerErrorPrefix,
			Err:    fmt.Errorf("%w: %s", ErrUnsupportedAuthMethod, authMethod),
		}
	}

	// Set the grant type before generating the token key
	cfg = cfg.WithGrantType(grantType)

	// Extract environment ID and client ID from configuration
	environmentID := ""
	if cfg.Endpoint.EnvironmentID != nil {
		environmentID = *cfg.Endpoint.EnvironmentID
	}

	clientID := ""
	switch grantType {
	case svcOAuth2.GrantTypeDeviceCode:
		if cfg.Auth.DeviceCode != nil && cfg.Auth.DeviceCode.DeviceCodeClientID != nil {
			clientID = *cfg.Auth.DeviceCode.DeviceCodeClientID
		}
	case svcOAuth2.GrantTypeAuthorizationCode:
		if cfg.Auth.AuthorizationCode != nil && cfg.Auth.AuthorizationCode.AuthorizationCodeClientID != nil {
			clientID = *cfg.Auth.AuthorizationCode.AuthorizationCodeClientID
		}
	case svcOAuth2.GrantTypeClientCredentials:
		if cfg.Auth.ClientCredentials != nil && cfg.Auth.ClientCredentials.ClientCredentialsClientID != nil {
			clientID = *cfg.Auth.ClientCredentials.ClientCredentialsClientID
		}
	}

	// Build suffix to disambiguate across provider/grant/profile for both keychain and files
	profileName, _ := profiles.GetOptionValue(options.RootActiveProfileOption)
	if profileName == "" {
		profileName = "default"
	}
	providerName, _ := profiles.GetOptionValue(options.AuthProviderOption)
	if strings.TrimSpace(providerName) == "" {
		providerName = "pingone"
	}
	suffix := fmt.Sprintf("_%s_%s_%s", providerName, string(grantType), profileName)
	// Use the SDK's GenerateKeychainAccountName with optional suffix
	tokenKey := svcOAuth2.GenerateKeychainAccountNameWithSuffix(environmentID, clientID, string(grantType), suffix)
	if tokenKey == "" || tokenKey == "default-token" {
		return "", &errs.PingCLIError{
			Prefix: tokenManagerErrorPrefix,
			Err:    ErrTokenKeyGenerationRequirements,
		}
	}

	return tokenKey, nil
}

// GetAuthMethodKeyFromConfig generates a unique keychain account name from a configuration object
// This uses the SDK's GenerateKeychainAccountName to ensure consistency with SDK token storage
func GetAuthMethodKeyFromConfig(cfg *config.Configuration) (string, error) {
	if cfg == nil || cfg.Auth.GrantType == nil {
		return "", ErrGrantTypeNotSet
	}

	// Extract environment ID from the config object
	environmentID := ""
	if cfg.Endpoint.EnvironmentID != nil {
		environmentID = *cfg.Endpoint.EnvironmentID
	}

	// Extract client ID based on grant type
	grantType := *cfg.Auth.GrantType
	clientID := ""
	switch grantType {
	case svcOAuth2.GrantTypeDeviceCode:
		if cfg.Auth.DeviceCode != nil && cfg.Auth.DeviceCode.DeviceCodeClientID != nil {
			clientID = *cfg.Auth.DeviceCode.DeviceCodeClientID
		}
	case svcOAuth2.GrantTypeAuthorizationCode:
		if cfg.Auth.AuthorizationCode != nil && cfg.Auth.AuthorizationCode.AuthorizationCodeClientID != nil {
			clientID = *cfg.Auth.AuthorizationCode.AuthorizationCodeClientID
		}
	case svcOAuth2.GrantTypeClientCredentials:
		if cfg.Auth.ClientCredentials != nil && cfg.Auth.ClientCredentials.ClientCredentialsClientID != nil {
			clientID = *cfg.Auth.ClientCredentials.ClientCredentialsClientID
		}
	}

	// Build suffix to disambiguate across provider/grant/profile for both keychain and files
	profileName, _ := profiles.GetOptionValue(options.RootActiveProfileOption)
	if profileName == "" {
		profileName = "default"
	}
	providerName, _ := profiles.GetOptionValue(options.AuthProviderOption)
	if strings.TrimSpace(providerName) == "" {
		providerName = "pingone"
	}
	suffix := fmt.Sprintf("_%s_%s_%s", providerName, string(grantType), profileName)
	// Use the SDK's GenerateKeychainAccountName with optional suffix
	tokenKey := svcOAuth2.GenerateKeychainAccountNameWithSuffix(environmentID, clientID, string(grantType), suffix)
	if tokenKey == "" || tokenKey == "default-token" {
		return "", &errs.PingCLIError{
			Prefix: tokenManagerErrorPrefix,
			Err:    ErrTokenKeyGenerationRequirements,
		}
	}

	return tokenKey, nil
}

// SaveToken saves a token to the keychain for the currently configured authentication method
func (tm *DefaultTokenManager) SaveToken(token *oauth2.Token) error {
	authMethod, err := GetCurrentAuthMethod()
	if err != nil {
		return fmt.Errorf("failed to get current grant type: %w", err)
	}

	_, err = SaveTokenForMethod(token, authMethod)

	return err
}

// LoadToken loads a token from the keychain for the currently configured authentication method
func (tm *DefaultTokenManager) LoadToken() (*oauth2.Token, error) {
	authMethod, err := GetCurrentAuthMethod()
	if err != nil {
		return nil, fmt.Errorf("failed to get current grant type: %w", err)
	}

	return LoadTokenForMethod(authMethod)
}

// ClearToken clears the token from the keychain for the currently configured authentication method
func (tm *DefaultTokenManager) ClearToken() error {
	authMethod, err := GetCurrentAuthMethod()
	if err != nil {
		return fmt.Errorf("failed to get current grant type: %w", err)
	}

	_, err = ClearTokenForMethod(authMethod)

	return err
}

// HasToken checks if a token exists in the keychain for the currently configured authentication method
func (tm *DefaultTokenManager) HasToken() bool {
	tokenKey, err := GetCurrentAuthMethod()
	if err != nil {
		return false
	}

	storage, err := svcOAuth2.NewKeychainStorage("pingcli", tokenKey)
	if err != nil {
		return false
	}
	hasToken, err := storage.HasToken()

	return err == nil && hasToken
}
