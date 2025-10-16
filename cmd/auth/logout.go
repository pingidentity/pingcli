// Copyright © 2025 Ping Identity Corporation

package auth

import (
	"fmt"

	"github.com/pingidentity/pingcli/cmd/common"
	auth_internal "github.com/pingidentity/pingcli/internal/auth"
	"github.com/pingidentity/pingcli/internal/configuration/options"
	"github.com/pingidentity/pingcli/internal/profiles"
	"github.com/spf13/cobra"
)

func NewLogoutCommand() *cobra.Command {
	cmd := &cobra.Command{
		Args:                  common.ExactArgs(0),
		DisableFlagsInUseLine: true, // We write our own flags in @Use attribute
		Long:                  "Logout user from the CLI by clearing stored credentials from Keychain",
		RunE:                  authLogoutRunE,
		Short:                 "Logout user from the CLI",
		Use:                   "logout [flags]",
	}

	return cmd
}

func authLogoutRunE(cmd *cobra.Command, args []string) error {
	// Get the currently configured authentication type
	authType, err := profiles.GetOptionValue(options.PingOneAuthenticationTypeOption)
	if err != nil {
		return fmt.Errorf("failed to get authentication type: %w", err)
	}

	// Support original "worker" type as alias for "client_credentials"
	if authType == "worker" {
		authType = "client_credentials"
	}

	// Generate the hash-based token key using the same logic as login/token manager
	tokenKey, err := auth_internal.GetCurrentAuthMethod()
	if err != nil {
		return fmt.Errorf("failed to generate token key for %s: %w", authType, err)
	}

	// Clear only the token for the currently configured authentication method
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
