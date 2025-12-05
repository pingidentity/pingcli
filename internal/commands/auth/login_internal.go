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

var (
	loginErrorPrefix = "failed to login"
)

// AuthLoginRunE implements the login command logic, handling authentication based on the selected
// method (auth code, device code, or client credentials) with support for interactive configuration
func AuthLoginRunE(cmd *cobra.Command, args []string) error {
	ctx := cmd.Context()

	// Get current profile name for messaging
	profileName, err := profiles.GetOptionValue(options.RootActiveProfileOption)
	if err != nil {
		return &errs.PingCLIError{
			Prefix: loginErrorPrefix,
			Err:    err,
		}
	}

	provider, err := profiles.GetOptionValue(options.AuthProviderOption)
	if err != nil || strings.TrimSpace(provider) == "" {
		// Default to pingone if no provider is specified
		provider = customtypes.ENUM_AUTH_PROVIDER_PINGONE
	}

	switch provider {
	case customtypes.ENUM_AUTH_PROVIDER_PINGONE:
		// Determine desired authentication method
		deviceCodeStr, _ := profiles.GetOptionValue(options.AuthMethodDeviceCodeOption)
		clientCredentialsStr, _ := profiles.GetOptionValue(options.AuthMethodClientCredentialsOption)
		authorizationCodeStr, _ := profiles.GetOptionValue(options.AuthMethodAuthorizationCodeOption)

		flagProvided := deviceCodeStr == "true" || clientCredentialsStr == "true" || authorizationCodeStr == "true"

		// If no flag was provided, check if authentication type is configured
		var authType string
		if !flagProvided {
			authType, err = profiles.GetOptionValue(options.PingOneAuthenticationTypeOption)
			if err != nil || strings.TrimSpace(authType) == "" {
				// No authentication type configured - run interactive setup
				if err := RunInteractiveAuthConfig(os.Stdin); err != nil {
					return &errs.PingCLIError{
						Prefix: loginErrorPrefix,
						Err:    err,
					}
				}
				// Reload auth type from profile after interactive setup
				authType, err = profiles.GetOptionValue(options.PingOneAuthenticationTypeOption)
				if err != nil || strings.TrimSpace(authType) == "" {
					return &errs.PingCLIError{
						Prefix: loginErrorPrefix,
						Err:    ErrInvalidAuthType,
					}
				}
			}
		}

		// Determine which authentication method was requested and convert to auth type format
		// If flags were provided, they take precedence. Otherwise, preserve configured authType (including legacy 'worker').
		if flagProvided {
			switch {
			case deviceCodeStr == "true":
				authType = customtypes.ENUM_PINGONE_AUTHENTICATION_TYPE_DEVICE_CODE
			case clientCredentialsStr == "true":
				authType = customtypes.ENUM_PINGONE_AUTHENTICATION_TYPE_CLIENT_CREDENTIALS
			default:
				authType = customtypes.ENUM_PINGONE_AUTHENTICATION_TYPE_AUTHORIZATION_CODE
			}
		}

		// Perform login based on auth type (ensure not empty)
		if strings.TrimSpace(authType) == "" {
			return &errs.PingCLIError{Prefix: loginErrorPrefix, Err: ErrInvalidAuthType}
		}
		err = performLoginByConfiguredType(ctx, authType, profileName)
		if err != nil {
			return err
		}
	default:
		return &errs.PingCLIError{
			Prefix: loginErrorPrefix,
			Err:    ErrInvalidAuthProvider,
		}
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
		// Pre-validate configuration; if missing, run interactive setup for device_code
		if _, cfgErr := GetDeviceCodeConfiguration(); cfgErr != nil {
			if interr := RunInteractiveAuthConfigForType(os.Stdin, customtypes.ENUM_PINGONE_AUTHENTICATION_TYPE_DEVICE_CODE); interr != nil {
				return &errs.PingCLIError{Prefix: loginErrorPrefix, Err: cfgErr}
			}
		}
		selectedMethod = string(svcOAuth2.GrantTypeDeviceCode)
		result, err = PerformDeviceCodeLogin(ctx)
	case customtypes.ENUM_PINGONE_AUTHENTICATION_TYPE_AUTHORIZATION_CODE:
		// Pre-validate configuration; if missing, run interactive setup for authorization_code
		if _, cfgErr := GetAuthorizationCodeConfiguration(); cfgErr != nil {
			if interr := RunInteractiveAuthConfigForType(os.Stdin, customtypes.ENUM_PINGONE_AUTHENTICATION_TYPE_AUTHORIZATION_CODE); interr != nil {
				return &errs.PingCLIError{Prefix: loginErrorPrefix, Err: cfgErr}
			}
		}
		selectedMethod = string(svcOAuth2.GrantTypeAuthorizationCode)
		result, err = PerformAuthorizationCodeLogin(ctx)
	case customtypes.ENUM_PINGONE_AUTHENTICATION_TYPE_CLIENT_CREDENTIALS:
		// Pre-validate configuration; if missing, run interactive setup for client_credentials
		if _, cfgErr := GetClientCredentialsConfiguration(); cfgErr != nil {
			if interr := RunInteractiveAuthConfigForType(os.Stdin, customtypes.ENUM_PINGONE_AUTHENTICATION_TYPE_CLIENT_CREDENTIALS); interr != nil {
				return &errs.PingCLIError{Prefix: loginErrorPrefix, Err: cfgErr}
			}
		}
		selectedMethod = string(svcOAuth2.GrantTypeClientCredentials)
		result, err = PerformClientCredentialsLogin(ctx)
	case customtypes.ENUM_PINGONE_AUTHENTICATION_TYPE_WORKER:
		// Legacy 'worker' type maps to client credentials flow
		if _, cfgErr := GetClientCredentialsConfiguration(); cfgErr != nil {
			if interr := RunInteractiveAuthConfigForType(os.Stdin, customtypes.ENUM_PINGONE_AUTHENTICATION_TYPE_CLIENT_CREDENTIALS); interr != nil {
				return &errs.PingCLIError{Prefix: loginErrorPrefix, Err: cfgErr}
			}
		}
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

		output.Success(fmt.Sprintf("Successfully logged in using %s. Credentials saved to %s for profile '%s'.", selectedMethod, storageMsg, profileName), nil)
		if token.RefreshToken != "" {
			output.Message("Refresh token available for automatic renewal.", nil)
		}
	} else {
		// Using cached token - SDK already logged the expiry
		output.Success(fmt.Sprintf("Using existing %s token for profile '%s'.", selectedMethod, profileName), nil)
		if token.RefreshToken != "" {
			output.Message("Token will be automatically refreshed when needed.", nil)
		}
	}
}
