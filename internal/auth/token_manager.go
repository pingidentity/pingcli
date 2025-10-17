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
	// Map worker to client_credentials
	if authMethod == "worker" {
		authMethod = "client_credentials"
	}

	// Get environment ID and client ID based on auth method
	var environmentID, clientID string

	switch authMethod {
	case "device_code":
		environmentID, _ = profiles.GetOptionValue(options.PingOneAuthenticationDeviceCodeEnvironmentIDOption)
		clientID, _ = profiles.GetOptionValue(options.PingOneAuthenticationDeviceCodeClientIDOption)
	case "auth_code", "authorization_code":
		environmentID, _ = profiles.GetOptionValue(options.PingOneAuthenticationAuthCodeEnvironmentIDOption)
		clientID, _ = profiles.GetOptionValue(options.PingOneAuthenticationAuthCodeClientIDOption)
	case "client_credentials":
		environmentID, _ = profiles.GetOptionValue(options.PingOneAuthenticationClientCredentialsEnvironmentIDOption)
		clientID, _ = profiles.GetOptionValue(options.PingOneAuthenticationClientCredentialsClientIDOption)
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
