// Copyright © 2025 Ping Identity Corporation

package auth

import (
	"fmt"

	"github.com/pingidentity/pingcli/cmd/common"
	auth_internal "github.com/pingidentity/pingcli/internal/auth"
	"github.com/pingidentity/pingcli/internal/configuration/options"
	"github.com/pingidentity/pingcli/internal/customtypes"
	"github.com/pingidentity/pingcli/internal/profiles"
	"github.com/spf13/cobra"
)

func NewLogoutCommand() *cobra.Command {
	cmd := &cobra.Command{
		Args:                  common.ExactArgs(0),
		DisableFlagsInUseLine: true, // We write our own flags in @Use attribute
		Long:                  "Logout user from the CLI by clearing stored credentials from Keychain. By default, uses the authentication method configured in the active profile. You can specify a different authentication method using the auth method flags.",
		RunE:                  authLogoutRunE,
		Short:                 "Logout user from the CLI",
		Use:                   "logout [flags]",
	}

	// Add the same auth method flags as login command
	cmd.Flags().AddFlag(options.AuthMethodAuthCodeOption.Flag)
	cmd.Flags().AddFlag(options.AuthMethodClientCredentialsOption.Flag)
	cmd.Flags().AddFlag(options.AuthMethodDeviceCodeOption.Flag)

	// These flags are mutually exclusive - only one can be specified
	cmd.MarkFlagsMutuallyExclusive(
		options.AuthMethodAuthCodeOption.Flag.Name,
		options.AuthMethodClientCredentialsOption.Flag.Name,
		options.AuthMethodDeviceCodeOption.Flag.Name,
	)

	return cmd
}

func authLogoutRunE(cmd *cobra.Command, args []string) error {
	// Check if any auth method flags were provided
	deviceCodeStr, err := profiles.GetOptionValue(options.AuthMethodDeviceCodeOption)
	if err != nil {
		return fmt.Errorf("failed to get device-code flag: %w", err)
	}

	clientCredentialsStr, err := profiles.GetOptionValue(options.AuthMethodClientCredentialsOption)
	if err != nil {
		return fmt.Errorf("failed to get client-credentials flag: %w", err)
	}

	var authType string
	var tokenKey string

	// Determine which auth method to use
	switch {
	case deviceCodeStr == "true":
		authType = customtypes.ENUM_PINGONE_AUTHENTICATION_TYPE_DEVICE_CODE
		tokenKey, err = auth_internal.GetAuthMethodKey(authType)
		if err != nil {
			return fmt.Errorf("failed to generate token key for %s: %w", authType, err)
		}
	case clientCredentialsStr == "true":
		authType = customtypes.ENUM_PINGONE_AUTHENTICATION_TYPE_CLIENT_CREDENTIALS
		tokenKey, err = auth_internal.GetAuthMethodKey(authType)
		if err != nil {
			return fmt.Errorf("failed to generate token key for %s: %w", authType, err)
		}
	default:
		// No flags provided or --auth-code flag - default to auth_code
		authType = customtypes.ENUM_PINGONE_AUTHENTICATION_TYPE_AUTH_CODE
		tokenKey, err = auth_internal.GetAuthMethodKey(authType)
		if err != nil {
			return fmt.Errorf("failed to generate token key for %s: %w", authType, err)
		}
	}

	// Clear only the token for the specified or configured authentication method
	if err := auth_internal.ClearTokenForMethod(tokenKey); err != nil {
		return fmt.Errorf("failed to clear %s credentials: %w", authType, err)
	}

	// Get current profile name for the logout message
	profileName, err := profiles.GetOptionValue(options.RootActiveProfileOption)
	if err != nil {
		profileName = "unknown" // fallback if we can't get the profile name
	}

	fmt.Printf("Successfully logged out from %s authentication. Credentials cleared from Keychain for profile '%s'.\n", authType, profileName)

	return nil
}
