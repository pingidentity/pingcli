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
	deviceCodeFlag := cmd.Flag(options.AuthMethodDeviceCodeOption.Flag.Name)
	clientCredentialsFlag := cmd.Flag(options.AuthMethodClientCredentialsOption.Flag.Name)
	authCodeFlag := cmd.Flag(options.AuthMethodAuthCodeOption.Flag.Name)

	flagProvided := (deviceCodeFlag != nil && deviceCodeFlag.Changed) ||
		(clientCredentialsFlag != nil && clientCredentialsFlag.Changed) ||
		(authCodeFlag != nil && authCodeFlag.Changed)

	// Get current profile name for messages
	profileName, err := profiles.GetOptionValue(options.RootActiveProfileOption)
	if err != nil {
		profileName = "current profile"
	}

	if !flagProvided {
		// No flag provided - clear ALL tokens (keychain and file storage)
		if err := ClearToken(); err != nil {
			return fmt.Errorf("failed to clear credentials: %w", err)
		}

		fmt.Printf("Successfully logged out from all authentication methods. All credentials cleared from storage for profile '%s'.\n", profileName)
		return nil
	}

	// Flag was provided - determine which auth method to clear
	deviceCodeStr, err := profiles.GetOptionValue(options.AuthMethodDeviceCodeOption)
	if err != nil {
		return fmt.Errorf("failed to get device-code flag: %w", err)
	}

	clientCredentialsStr, err := profiles.GetOptionValue(options.AuthMethodClientCredentialsOption)
	if err != nil {
		return fmt.Errorf("failed to get client-credentials flag: %w", err)
	}

	var authType string
	switch {
	case deviceCodeStr == "true":
		authType = customtypes.ENUM_PINGONE_AUTHENTICATION_TYPE_DEVICE_CODE
	case clientCredentialsStr == "true":
		authType = customtypes.ENUM_PINGONE_AUTHENTICATION_TYPE_CLIENT_CREDENTIALS
	default:
		authType = customtypes.ENUM_PINGONE_AUTHENTICATION_TYPE_AUTH_CODE
	}

	// Check if configuration exists for this auth method before trying to generate token key
	var hasConfig bool
	switch authType {
	case customtypes.ENUM_PINGONE_AUTHENTICATION_TYPE_AUTH_CODE:
		clientID, _ := profiles.GetOptionValue(options.PingOneAuthenticationAuthCodeClientIDOption)
		envID, _ := profiles.GetOptionValue(options.PingOneAuthenticationAuthCodeEnvironmentIDOption)
		hasConfig = clientID != "" && envID != ""
	case customtypes.ENUM_PINGONE_AUTHENTICATION_TYPE_DEVICE_CODE:
		clientID, _ := profiles.GetOptionValue(options.PingOneAuthenticationDeviceCodeClientIDOption)
		envID, _ := profiles.GetOptionValue(options.PingOneAuthenticationDeviceCodeEnvironmentIDOption)
		hasConfig = clientID != "" && envID != ""
	case customtypes.ENUM_PINGONE_AUTHENTICATION_TYPE_CLIENT_CREDENTIALS:
		clientID, _ := profiles.GetOptionValue(options.PingOneAuthenticationClientCredentialsClientIDOption)
		envID, _ := profiles.GetOptionValue(options.PingOneAuthenticationClientCredentialsEnvironmentIDOption)
		hasConfig = clientID != "" && envID != ""
	}

	if !hasConfig {
		return fmt.Errorf("logout failed for %s in %s: %w", authType, profileName, ErrNoAuthConfiguration)
	}

	// Generate token key for the selected auth method
	tokenKey, err := GetAuthMethodKey(authType)
	if err != nil {
		return fmt.Errorf("failed to generate token key for %s: %w", authType, err)
	}

	// Clear only the token for the specified authentication method
	if err := ClearTokenForMethod(tokenKey); err != nil {
		return fmt.Errorf("failed to clear %s credentials: %w", authType, err)
	}

	fmt.Printf("Successfully logged out from %s authentication. Credentials cleared from storage for profile '%s'.\n", authType, profileName)

	return nil
}
