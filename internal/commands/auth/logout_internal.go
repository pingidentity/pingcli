// Copyright © 2026 Ping Identity Corporation
package auth_internal

import (
	"fmt"
	"strings"

	"github.com/pingidentity/pingcli/internal/configuration/options"
	"github.com/pingidentity/pingcli/internal/customtypes"
	"github.com/pingidentity/pingcli/internal/errs"
	"github.com/pingidentity/pingcli/internal/output"
	"github.com/pingidentity/pingcli/internal/profiles"
	"github.com/pingidentity/pingone-go-client/config"
	svcOAuth2 "github.com/pingidentity/pingone-go-client/oauth2"
	"github.com/spf13/cobra"
)

var (
	logoutErrorPrefix = "failed to logout"
)

// AuthLogoutRunE implements the logout command logic, clearing credentials from both
// keychain and file storage. If no grant type flag is provided, clears all tokens.
// If a specific grant type flag is provided, clears only that method's token.
func AuthLogoutRunE(cmd *cobra.Command, args []string) error {
	// Check if any grant type flags were provided
	deviceCodeStr, err := profiles.GetOptionValue(options.AuthMethodDeviceCodeOption)
	if err != nil {
		return &errs.PingCLIError{Prefix: logoutErrorPrefix, Err: err}
	}

	clientCredentialsStr, err := profiles.GetOptionValue(options.AuthMethodClientCredentialsOption)
	if err != nil {
		return &errs.PingCLIError{Prefix: logoutErrorPrefix, Err: err}
	}

	authorizationCodeStr, err := profiles.GetOptionValue(options.AuthMethodAuthorizationCodeOption)
	if err != nil {
		return &errs.PingCLIError{Prefix: logoutErrorPrefix, Err: err}
	}

	flagProvided := deviceCodeStr == "true" || clientCredentialsStr == "true" || authorizationCodeStr == "true"

	// Get current profile name for messages
	profileName, err := profiles.GetOptionValue(options.RootActiveProfileOption)
	if err != nil {
		profileName = "current profile"
	}

	// Get service name for token key generation
	providerName, err := profiles.GetOptionValue(options.AuthProviderOption)
	if err != nil || providerName == "" {
		providerName = customtypes.ENUM_AUTH_PROVIDER_PINGONE // Default to pingone
	}

	if !flagProvided {
		// No flag provided - clear ALL tokens (keychain and file storage)
		if err := ClearAllTokens(); err != nil {
			return fmt.Errorf("%s: %w", credentialsErrorPrefix, err)
		}
		// Report the storage cleared using common formatter
		output.Success(fmt.Sprintf("Successfully logged out and cleared credentials from all methods for service '%s' using profile '%s'.", providerName, profileName), nil)

		return nil
	}

	// Flag was provided - determine which grant type to clear
	// (deviceCodeStr, clientCredentialsStr, authCodeStr already retrieved above)

	var authType string
	switch {
	case deviceCodeStr == "true":
		authType = customtypes.ENUM_PINGONE_AUTHENTICATION_TYPE_DEVICE_CODE
	case clientCredentialsStr == "true":
		authType = customtypes.ENUM_PINGONE_AUTHENTICATION_TYPE_CLIENT_CREDENTIALS
	case authorizationCodeStr == "true":
		authType = customtypes.ENUM_PINGONE_AUTHENTICATION_TYPE_AUTHORIZATION_CODE
	default:
		return &errs.PingCLIError{Prefix: credentialsErrorPrefix, Err: ErrInvalidAuthMethod}
	}

	// Generate token key for the selected grant type
	tokenKey, err := GetAuthMethodKey(authType)
	if err != nil {
		return &errs.PingCLIError{Prefix: credentialsErrorPrefix, Err: err}
	}

	// Clear only the token for the specified grant type
	err = ClearToken(tokenKey)
	if err != nil {
		return &errs.PingCLIError{Prefix: credentialsErrorPrefix, Err: fmt.Errorf("failed to clear %s credentials: %w", authType, err)}
	}

	output.Success(fmt.Sprintf("Successfully logged out and cleared credentials from %s for service '%s' using profile '%s'.", authType, providerName, profileName), nil)

	return nil
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
	profileName, err := profiles.GetOptionValue(options.RootActiveProfileOption)
	if err != nil || profileName == "" {
		profileName = "default"
	}
	providerName, err := profiles.GetOptionValue(options.AuthProviderOption)
	if err != nil || strings.TrimSpace(providerName) == "" {
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
