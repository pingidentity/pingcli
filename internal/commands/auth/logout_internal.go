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
// keychain and file storage. If no auth method flag is provided, clears all tokens.
// If a specific auth method flag is provided, clears only that method's token.
func AuthLogoutRunE(cmd *cobra.Command, args []string) error {
	// Check if any auth method flags were provided
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

		fmt.Printf("Successfully logged out from all authentication methods for service '%s'. All credentials cleared from storage for profile '%s'.\n", providerName, profileName)

		return nil
	}

	// Flag was provided - determine which auth method to clear
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

	// Generate token key for the selected auth method
	tokenKey, err := GetAuthMethodKey(authType)
	if err != nil {
		return fmt.Errorf("failed to generate token key for %s: %w", authType, err)
	}

	// Clear only the token for the specified authentication method
	location, err := ClearTokenForMethod(tokenKey)
	if err != nil {
		return fmt.Errorf("failed to clear %s credentials: %w", authType, err)
	}

	// Build storage location message
	var storageMsg string
	switch {
	case location.Keychain && location.File:
		storageMsg = "keychain and file storage"
	case location.Keychain:
		storageMsg = "keychain"
	case location.File:
		storageMsg = "file storage"
	default:
		storageMsg = "storage"
	}

	fmt.Printf("Successfully logged out from %s authentication for service '%s'. Credentials cleared from %s for profile '%s'.\n", authType, providerName, storageMsg, profileName)

	return nil
}
