// Copyright © 2025 Ping Identity Corporation

package auth_internal

import (
	"crypto/sha256"
	"fmt"

	"github.com/pingidentity/pingcli/internal/configuration/options"
	"github.com/pingidentity/pingcli/internal/profiles"
	"github.com/pingidentity/pingone-go-client/config"
	svcOAuth2 "github.com/pingidentity/pingone-go-client/oauth2"
	"golang.org/x/oauth2"
)

type TokenManager interface {
	SaveToken(token *oauth2.Token) error
	LoadToken() (*oauth2.Token, error)
	ClearToken() error
	HasToken() bool
}

type DefaultTokenManager struct {
	serviceName string
}

func NewDefaultTokenManager() TokenManager {
	return &DefaultTokenManager{
		serviceName: "pingcli",
	}
}

func GetCurrentAuthMethod() (string, error) {
	authMethod, err := profiles.GetOptionValue(options.PingOneAuthenticationTypeOption)
	if err != nil {
		return "", fmt.Errorf("failed to get current auth method: %w", err)
	}

	if authMethod == "" {
		return "", fmt.Errorf("auth method is not configured")
	}

	return GetAuthMethodKey(authMethod)
}

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
		return "", fmt.Errorf("unsupported auth method: %s", authMethod)
	}

	if environmentID == "" || clientID == "" {
		return "", fmt.Errorf("environment ID and client ID are required for token key generation (env: %s, client: %s)", environmentID, clientID)
	}

	// Create a hash of environment ID + client ID + auth method for uniqueness
	hash := sha256.Sum256([]byte(fmt.Sprintf("%s:%s:%s", environmentID, clientID, authMethod)))
	tokenKey := fmt.Sprintf("token-%x", hash[:8]) // Use first 8 bytes of hash for shorter key

	return tokenKey, nil
}

func GetAuthMethodKeyFromConfig(cfg *config.Configuration) (string, error) {
	if cfg == nil || cfg.Auth.GrantType == nil {
		return "", fmt.Errorf("configuration does not have grant type set")
	}

	// Convert GrantType to string
	grantType := string(*cfg.Auth.GrantType)

	return GetAuthMethodKey(grantType)
}

func (tm *DefaultTokenManager) SaveToken(token *oauth2.Token) error {
	authMethod, err := GetCurrentAuthMethod()
	if err != nil {
		return fmt.Errorf("failed to get current auth method: %w", err)
	}

	return SaveTokenForMethod(token, authMethod)
}

func (tm *DefaultTokenManager) LoadToken() (*oauth2.Token, error) {
	authMethod, err := GetCurrentAuthMethod()
	if err != nil {
		return nil, fmt.Errorf("failed to get current auth method: %w", err)
	}

	return LoadTokenForMethod(authMethod)
}

func (tm *DefaultTokenManager) ClearToken() error {
	authMethod, err := GetCurrentAuthMethod()
	if err != nil {
		return fmt.Errorf("failed to get current auth method: %w", err)
	}

	return ClearTokenForMethod(authMethod)
}

func (tm *DefaultTokenManager) HasToken() bool {
	tokenKey, err := GetCurrentAuthMethod()
	if err != nil {
		return false
	}

	storage := svcOAuth2.NewKeychainStorage("pingcli", tokenKey)
	hasToken, err := storage.HasToken()

	return err == nil && hasToken
}
