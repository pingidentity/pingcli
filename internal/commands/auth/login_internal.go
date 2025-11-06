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
	authCodeStr, _ := profiles.GetOptionValue(options.AuthMethodAuthCodeOption)

	flagProvided := deviceCodeStr != "" || clientCredentialsStr != "" || authCodeStr != ""

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

		// Now proceed with login using the configured authentication type
		return performLoginByConfiguredType(ctx, authType, profileName)
	}

	// Flag was provided - use the flag value to override any configuration
	// (deviceCodeStr, clientCredentialsStr, authCodeStr already retrieved above)

	// Determine which authentication method was requested and convert to auth type format
	var authType string
	switch {
	case deviceCodeStr == "true":
		authType = customtypes.ENUM_PINGONE_AUTHENTICATION_TYPE_DEVICE_CODE
	case clientCredentialsStr == "true":
		authType = customtypes.ENUM_PINGONE_AUTHENTICATION_TYPE_CLIENT_CREDENTIALS
	default:
		authType = customtypes.ENUM_PINGONE_AUTHENTICATION_TYPE_AUTH_CODE
	}

	// Use the common login flow
	return performLoginByConfiguredType(ctx, authType, profileName)
}

// ensureAuthConfigurationExists checks if the required configuration exists for the given auth method
// Returns an error if configuration is missing or invalid
func ensureAuthConfigurationExists(authMethod string) error {
	switch authMethod {
	case string(svcOAuth2.GrantTypeDeviceCode):
		_, err := GetDeviceCodeConfiguration()

		return err
	case string(svcOAuth2.GrantTypeClientCredentials):
		_, err := GetClientCredentialsConfiguration()

		return err
	case string(svcOAuth2.GrantTypeAuthCode):
		_, err := GetAuthCodeConfiguration()

		return err
	default:
		return &errs.PingCLIError{
			Prefix: fmt.Sprintf("unsupported authentication method: %s", authMethod),
			Err:    ErrUnsupportedAuthMethod,
		}
	}
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
	case customtypes.ENUM_PINGONE_AUTHENTICATION_TYPE_CLIENT_CREDENTIALS:
		selectedMethod = string(svcOAuth2.GrantTypeClientCredentials)
		result, err = PerformClientCredentialsLogin(ctx)
	case customtypes.ENUM_PINGONE_AUTHENTICATION_TYPE_AUTH_CODE:
		selectedMethod = string(svcOAuth2.GrantTypeAuthCode)
		result, err = PerformAuthCodeLogin(ctx)
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
		if location.Keychain && location.File {
			storageMsg = "keychain and file storage"
		} else if location.Keychain {
			storageMsg = "keychain"
		} else if location.File {
			storageMsg = "file storage"
		} else {
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
