// Copyright © 2025 Ping Identity Corporation

package auth_internal

import (
	"fmt"

	"github.com/pingidentity/pingcli/internal/configuration/options"
	"github.com/pingidentity/pingcli/internal/errs"
	"github.com/pingidentity/pingcli/internal/profiles"
	"github.com/pingidentity/pingone-go-client/config"
	svcOAuth2 "github.com/pingidentity/pingone-go-client/oauth2"
	"golang.org/x/oauth2"
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
		return "", fmt.Errorf("failed to get current auth method: %w", err)
	}

	if authMethod == "" {
		return "", ErrAuthMethodNotConfigured
	}

	return GetAuthMethodKey(authMethod)
}

// GetAuthMethodKey generates a unique keychain account name for the given authentication method
// using the environment ID and client ID from the profile configuration
func GetAuthMethodKey(authMethod string) (string, error) {
	// Get environment ID and client ID based on auth method
	var environmentID, clientID string
	var err error

	switch authMethod {
	case "device_code":
		environmentID, err = profiles.GetOptionValue(options.PingOneAuthenticationDeviceCodeEnvironmentIDOption)
		if err != nil {
			return "", fmt.Errorf("failed to get device code environment ID: %w", err)
		}
		clientID, err = profiles.GetOptionValue(options.PingOneAuthenticationDeviceCodeClientIDOption)
		if err != nil {
			return "", fmt.Errorf("failed to get device code client ID: %w", err)
		}
	case "auth_code", "authorization_code":
		environmentID, err = profiles.GetOptionValue(options.PingOneAuthenticationAuthCodeEnvironmentIDOption)
		if err != nil {
			return "", fmt.Errorf("failed to get auth code environment ID: %w", err)
		}
		clientID, err = profiles.GetOptionValue(options.PingOneAuthenticationAuthCodeClientIDOption)
		if err != nil {
			return "", fmt.Errorf("failed to get auth code client ID: %w", err)
		}
	case "client_credentials":
		environmentID, err = profiles.GetOptionValue(options.PingOneAuthenticationClientCredentialsEnvironmentIDOption)
		if err != nil {
			return "", fmt.Errorf("failed to get client credentials environment ID: %w", err)
		}
		clientID, err = profiles.GetOptionValue(options.PingOneAuthenticationClientCredentialsClientIDOption)
		if err != nil {
			return "", fmt.Errorf("failed to get client credentials client ID: %w", err)
		}
	default:
		return "", &errs.PingCLIError{
			Prefix: fmt.Sprintf("failed to generate token key for auth method '%s'", authMethod),
			Err:    ErrUnsupportedAuthMethod,
		}
	}

	if environmentID == "" || clientID == "" {
		return "", &errs.PingCLIError{
			Prefix: "failed to generate token key",
			Err:    ErrTokenKeyGenerationRequirements,
		}
	}

	// Use the SDK's GenerateKeychainAccountName for consistency
	return svcOAuth2.GenerateKeychainAccountName(environmentID, clientID, authMethod), nil
}

// GetAuthMethodKeyFromConfig generates a unique keychain account name from a configuration object
func GetAuthMethodKeyFromConfig(cfg *config.Configuration) (string, error) {
	if cfg == nil || cfg.Auth.GrantType == nil {
		return "", ErrGrantTypeNotSet
	}

	// Convert GrantType to string
	grantType := string(*cfg.Auth.GrantType)

	return GetAuthMethodKey(grantType)
}

// SaveToken saves a token to the keychain for the currently configured authentication method
func (tm *DefaultTokenManager) SaveToken(token *oauth2.Token) error {
	authMethod, err := GetCurrentAuthMethod()
	if err != nil {
		return fmt.Errorf("failed to get current auth method: %w", err)
	}

	_, err = SaveTokenForMethod(token, authMethod)
	return err
}

// LoadToken loads a token from the keychain for the currently configured authentication method
func (tm *DefaultTokenManager) LoadToken() (*oauth2.Token, error) {
	authMethod, err := GetCurrentAuthMethod()
	if err != nil {
		return nil, fmt.Errorf("failed to get current auth method: %w", err)
	}

	return LoadTokenForMethod(authMethod)
}

// ClearToken clears the token from the keychain for the currently configured authentication method
func (tm *DefaultTokenManager) ClearToken() error {
	authMethod, err := GetCurrentAuthMethod()
	if err != nil {
		return fmt.Errorf("failed to get current auth method: %w", err)
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
