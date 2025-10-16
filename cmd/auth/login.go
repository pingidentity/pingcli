// Copyright © 2025 Ping Identity Corporation

package auth

import (
	"context"
	"fmt"

	"github.com/pingidentity/pingcli/cmd/common"
	auth_internal "github.com/pingidentity/pingcli/internal/auth"
	"github.com/pingidentity/pingcli/internal/configuration/options"
	"github.com/pingidentity/pingcli/internal/profiles"
	"github.com/spf13/cobra"
)

func NewLoginCommand() *cobra.Command {
	cmd := &cobra.Command{
		Args:                  common.ExactArgs(0),
		DisableFlagsInUseLine: true, // We write our own flags in @Use attribute
		Long:                  "Login user to the CLI using one of the supported authentication flows: device code, authorization code, or client credentials",
		RunE:                  authLoginRunE,
		Short:                 "Login user to the CLI",
		Use:                   "login [flags]",
	}

	cmd.Flags().AddFlag(options.AuthMethodAuthCodeOption.Flag)
	cmd.Flags().AddFlag(options.AuthMethodClientCredentialsOption.Flag)
	cmd.Flags().AddFlag(options.AuthMethodDeviceCodeOption.Flag)

	return cmd
}

func authLoginRunE(cmd *cobra.Command, args []string) error {
	deviceCodeStr, err := profiles.GetOptionValue(options.AuthMethodDeviceCodeOption)
	if err != nil {
		return fmt.Errorf("failed to get device-code flag: %w", err)
	}

	clientCredentialsStr, err := profiles.GetOptionValue(options.AuthMethodClientCredentialsOption)
	if err != nil {
		return fmt.Errorf("failed to get client-credentials flag: %w", err)
	}

	authCodeStr, err := profiles.GetOptionValue(options.AuthMethodAuthCodeOption)
	if err != nil {
		return fmt.Errorf("failed to get auth-code flag: %w", err)
	}

	// Check that exactly one authentication method is specified
	authMethods := []string{deviceCodeStr, clientCredentialsStr, authCodeStr}
	selectedCount := 0
	for _, method := range authMethods {
		if method == "true" {
			selectedCount++
		}
	}

	if selectedCount == 0 {
		return fmt.Errorf("please specify an authentication method: --auth-code, --client-credentials, or --device-code")
	}

	if selectedCount > 1 {
		return fmt.Errorf("please specify only one authentication method")
	}

	ctx := context.Background()

	// Get current profile name for messaging
	profileName, err := profiles.GetOptionValue(options.RootActiveProfileOption)
	if err != nil {
		return fmt.Errorf("failed to get active profile: %w", err)
	}

	if deviceCodeStr == "true" {
		// Perform device code authentication
		token, newAuth, err := auth_internal.PerformDeviceCodeLogin(ctx)
		if err != nil {
			return fmt.Errorf("device code login failed: %w", err)
		}

		if newAuth {
			fmt.Printf("Successfully logged in using device_code authentication. Credentials saved to Keychain for profile '%s'.\n", profileName)
		} else {
			fmt.Printf("Already authenticated with valid device_code token for profile '%s'.\n", profileName)
		}
		fmt.Printf("Access token expires: %s\n", token.Expiry.Format("2006-01-02 15:04:05 MST"))
		if token.RefreshToken != "" {
			fmt.Printf("Refresh token available for automatic renewal.\n")
		}

		return nil
	}

	if clientCredentialsStr == "true" {
		// Perform client credentials authentication
		token, newAuth, err := auth_internal.PerformClientCredentialsLogin(ctx)
		if err != nil {
			return fmt.Errorf("client credentials login failed: %w", err)
		}

		if newAuth {
			fmt.Printf("Successfully logged in using client_credentials authentication. Credentials saved to Keychain for profile '%s'.\n", profileName)
		} else {
			fmt.Printf("Already authenticated with valid client_credentials token for profile '%s'.\n", profileName)
		}
		fmt.Printf("Access token expires: %s\n", token.Expiry.Format("2006-01-02 15:04:05 MST"))
		if token.RefreshToken != "" {
			fmt.Printf("Refresh token available for automatic renewal.\n")
		}

		return nil
	}

	if authCodeStr == "true" {
		// Perform authorization code authentication
		token, newAuth, err := auth_internal.PerformAuthCodeLogin(ctx)
		if err != nil {
			return fmt.Errorf("authorization code login failed: %w", err)
		}

		if newAuth {
			fmt.Printf("Successfully logged in using auth_code authentication. Credentials saved to Keychain for profile '%s'.\n", profileName)
		} else {
			fmt.Printf("Already authenticated with valid auth_code token for profile '%s'.\n", profileName)
		}
		fmt.Printf("Access token expires: %s\n", token.Expiry.Format("2006-01-02 15:04:05 MST"))
		if token.RefreshToken != "" {
			fmt.Printf("Refresh token available for automatic renewal.\n")
		}

		return nil
	}

	return fmt.Errorf("no valid authentication method selected")
}
