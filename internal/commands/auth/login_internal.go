// Copyright © 2025 Ping Identity Corporation

package auth_internal

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/pingidentity/pingcli/internal/configuration/options"
	"github.com/pingidentity/pingcli/internal/customtypes"
	"github.com/pingidentity/pingcli/internal/errs"
	"github.com/pingidentity/pingcli/internal/output"
	"github.com/pingidentity/pingcli/internal/profiles"
	svcOAuth2 "github.com/pingidentity/pingone-go-client/oauth2"
	"github.com/spf13/cobra"
	"golang.org/x/oauth2"
)

// AuthLoginRunE implements the login command logic, handling authentication based on the selected
// method (auth code, device code, or client credentials) with support for interactive configuration
func AuthLoginRunE(cmd *cobra.Command, args []string) error {
	ctx := cmd.Context()

	// Get current profile name for messaging
	profileName, err := profiles.GetOptionValue(options.RootActiveProfileOption)
	if err != nil {
		return &errs.PingCLIError{
			Prefix: "failed to get active profile",
			Err:    err,
		}
	}

	// Determine desired authentication method
	deviceCodeStr, _ := profiles.GetOptionValue(options.AuthMethodDeviceCodeOption)
	clientCredentialsStr, _ := profiles.GetOptionValue(options.AuthMethodClientCredentialsOption)
	authorizationCodeStr, _ := profiles.GetOptionValue(options.AuthMethodAuthorizationCodeOption)

	flagProvided := deviceCodeStr == "true" || clientCredentialsStr == "true" || authorizationCodeStr == "true"

	// If no flag was provided, check if authentication type is configured
	if !flagProvided {
		authType, err := profiles.GetOptionValue(options.PingOneAuthenticationTypeOption)
		if err != nil || strings.TrimSpace(authType) == "" {
			// No authentication type configured - run interactive setup
			if err := RunInteractiveAuthConfig(os.Stdin); err != nil {
				return err
			}
			// After interactive setup, re-read the auth type
			authType, err = profiles.GetOptionValue(options.PingOneAuthenticationTypeOption)
			if err != nil {
				return &errs.PingCLIError{
					Prefix: "failed to read authentication type after configuration",
					Err:    err,
				}
			}
		}
	}

	// Determine which authentication method was requested and convert to auth type format
	var authType string
	switch {
	case deviceCodeStr == "true":
		authType = customtypes.ENUM_PINGONE_AUTHENTICATION_TYPE_DEVICE_CODE
	case clientCredentialsStr == "true":
		authType = customtypes.ENUM_PINGONE_AUTHENTICATION_TYPE_CLIENT_CREDENTIALS
	default:
		authType = customtypes.ENUM_PINGONE_AUTHENTICATION_TYPE_AUTHORIZATION_CODE
	}

	// Perform login based on auth type
	err = performLoginByConfiguredType(ctx, authType, profileName)
	if err != nil {
		return err
	}

	return nil
}

// performLoginByConfiguredType performs login using the configured authentication type
func performLoginByConfiguredType(ctx context.Context, authType, profileName string) error {
	var result *LoginResult
	var err error
	var selectedMethod string

	switch authType {
	case customtypes.ENUM_PINGONE_AUTHENTICATION_TYPE_DEVICE_CODE:
		selectedMethod = string(svcOAuth2.GrantTypeDeviceCode)
		result, err = PerformDeviceCodeLogin(ctx)
	case customtypes.ENUM_PINGONE_AUTHENTICATION_TYPE_AUTHORIZATION_CODE:
		selectedMethod = string(svcOAuth2.GrantTypeAuthorizationCode)
		result, err = PerformAuthorizationCodeLogin(ctx)
	case customtypes.ENUM_PINGONE_AUTHENTICATION_TYPE_CLIENT_CREDENTIALS:
		selectedMethod = string(svcOAuth2.GrantTypeClientCredentials)
		result, err = PerformClientCredentialsLogin(ctx)
	default:
		return &errs.PingCLIError{
			Prefix: fmt.Sprintf("invalid authentication type: %s", authType),
			Err:    ErrInvalidAuthType,
		}
	}

	if err != nil {
		return &errs.PingCLIError{
			Prefix: fmt.Sprintf("authentication failed for %s", authType),
			Err:    err,
		}
	}

	displayLoginSuccess(result.Token, result.NewAuth, result.Location, selectedMethod, profileName)

	return nil
}

// displayLoginSuccess displays the successful login message
func displayLoginSuccess(token *oauth2.Token, newAuth bool, location StorageLocation, selectedMethod, profileName string) {
	if newAuth {
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

		output.Success(fmt.Sprintf("Successfully logged in using %s authentication. Credentials saved to %s for profile '%s'.", selectedMethod, storageMsg, profileName), nil)
	} else {
		output.Message(fmt.Sprintf("Already authenticated with valid %s token for profile '%s'.", selectedMethod, profileName), nil)
	}
	output.Message(fmt.Sprintf("Access token expires: %s", token.Expiry.Format("2006-01-02 15:04:05 MST")), nil)
	if token.RefreshToken != "" {
		output.Message("Refresh token available for automatic renewal.", nil)
	}
}
