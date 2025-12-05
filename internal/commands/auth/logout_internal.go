// Copyright © 2025 Ping Identity Corporation
package auth_internal

import (
	"fmt"

	"github.com/pingidentity/pingcli/internal/configuration/options"
	"github.com/pingidentity/pingcli/internal/customtypes"
	"github.com/pingidentity/pingcli/internal/profiles"
	"github.com/spf13/cobra"
)

// AuthLogoutRunE implements the logout command logic, clearing credentials from both
// keychain and file storage. If no grant type flag is provided, clears all tokens.
// If a specific grant type flag is provided, clears only that method's token.
func AuthLogoutRunE(cmd *cobra.Command, args []string) error {
	// Check if any grant type flags were provided
	deviceCodeStr, _ := profiles.GetOptionValue(options.AuthMethodDeviceCodeOption)
	clientCredentialsStr, _ := profiles.GetOptionValue(options.AuthMethodClientCredentialsOption)
	authorizationCodeStr, _ := profiles.GetOptionValue(options.AuthMethodAuthorizationCodeOption)

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
		if err := ClearToken(); err != nil {
			return fmt.Errorf("failed to clear credentials: %w", err)
		}

		// Report the storage cleared using common formatter
		fmt.Printf("Successfully logged out from all methods for service '%s'. Credentials cleared from %s for profile '%s'.\n", providerName, formatFullLogoutStorageMessage(), profileName)

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
	default:
		authType = customtypes.ENUM_PINGONE_AUTHENTICATION_TYPE_AUTHORIZATION_CODE
	}

	// Generate token key for the selected grant type
	tokenKey, err := GetAuthMethodKey(authType)
	if err != nil {
		return fmt.Errorf("failed to generate token key for %s: %w", authType, err)
	}

	// Clear only the token for the specified grant type
	location, err := ClearTokenForMethod(tokenKey)
	if err != nil {
		return fmt.Errorf("failed to clear %s credentials: %w", authType, err)
	}

	// Build storage location message via common formatter
	storageMsg := formatStorageLocation(location)

	fmt.Printf("Successfully logged out from %s for service '%s'. Credentials cleared from %s for profile '%s'.\n", authType, providerName, storageMsg, profileName)

	return nil
}
