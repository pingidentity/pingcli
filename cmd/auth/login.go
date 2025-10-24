// Copyright © 2025 Ping Identity Corporation

package auth

import (
	"context"
	"fmt"

	"github.com/pingidentity/pingcli/cmd/common"
	auth_internal "github.com/pingidentity/pingcli/internal/auth"
	"github.com/pingidentity/pingcli/internal/configuration/options"
	"github.com/pingidentity/pingcli/internal/profiles"
	svcOAuth2 "github.com/pingidentity/pingone-go-client/oauth2"
	"github.com/spf13/cobra"
	"golang.org/x/oauth2"
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

	// Enforce that exactly one authentication method must be specified
	cmd.MarkFlagsMutuallyExclusive(
		options.AuthMethodAuthCodeOption.Flag.Name,
		options.AuthMethodClientCredentialsOption.Flag.Name,
		options.AuthMethodDeviceCodeOption.Flag.Name,
	)

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

	ctx := context.Background()

	// Get current profile name for messaging
	profileName, err := profiles.GetOptionValue(options.RootActiveProfileOption)
	if err != nil {
		return fmt.Errorf("failed to get active profile: %w", err)
	}

	// Perform authentication based on selected method (auth_code is default if none specified)
	var token *oauth2.Token
	var newAuth bool
	var selectedMethod string

	switch {
	case deviceCodeStr == "true":
		selectedMethod = string(svcOAuth2.GrantTypeDeviceCode)
		token, newAuth, err = auth_internal.PerformDeviceCodeLogin(ctx)
		if err != nil {
			return fmt.Errorf("device code login failed: %w", err)
		}
	case clientCredentialsStr == "true":
		selectedMethod = string(svcOAuth2.GrantTypeClientCredentials)
		token, newAuth, err = auth_internal.PerformClientCredentialsLogin(ctx)
		if err != nil {
			return fmt.Errorf("client credentials login failed: %w", err)
		}
	default:
		// Default to auth_code if no method flag is specified
		selectedMethod = string(svcOAuth2.GrantTypeAuthCode)
		token, newAuth, err = auth_internal.PerformAuthCodeLogin(ctx)
		if err != nil {
			return fmt.Errorf("authorization code login failed: %w", err)
		}
	}

	// Display authentication result
	if newAuth {
		fmt.Printf("Successfully logged in using %s authentication. Credentials saved to Keychain for profile '%s'.\n", selectedMethod, profileName)
	} else {
		fmt.Printf("Already authenticated with valid %s token for profile '%s'.\n", selectedMethod, profileName)
	}
	fmt.Printf("Access token expires: %s\n", token.Expiry.Format("2006-01-02 15:04:05 MST"))
	if token.RefreshToken != "" {
		fmt.Printf("Refresh token available for automatic renewal.\n")
	}

	return nil
}
