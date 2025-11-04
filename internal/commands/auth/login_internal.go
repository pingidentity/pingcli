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
	ctx := context.Background()

	// Get current profile name for messaging
	profileName, err := profiles.GetOptionValue(options.RootActiveProfileOption)
	if err != nil {
		return fmt.Errorf("failed to get active profile: %w", err)
	}

	// Check if any authentication method flag was explicitly provided
	deviceCodeFlag := cmd.Flag(options.AuthMethodDeviceCodeOption.Flag.Name)
	clientCredentialsFlag := cmd.Flag(options.AuthMethodClientCredentialsOption.Flag.Name)
	authCodeFlag := cmd.Flag(options.AuthMethodAuthCodeOption.Flag.Name)

	flagProvided := (deviceCodeFlag != nil && deviceCodeFlag.Changed) ||
		(clientCredentialsFlag != nil && clientCredentialsFlag.Changed) ||
		(authCodeFlag != nil && authCodeFlag.Changed)

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
	deviceCodeStr, err := profiles.GetOptionValue(options.AuthMethodDeviceCodeOption)
	if err != nil {
		return fmt.Errorf("failed to get device-code flag: %w", err)
	}

	clientCredentialsStr, err := profiles.GetOptionValue(options.AuthMethodClientCredentialsOption)
	if err != nil {
		return fmt.Errorf("failed to get client-credentials flag: %w", err)
	}

	// Determine which authentication method was requested
	var selectedMethod string
	switch {
	case deviceCodeStr == "true":
		selectedMethod = string(svcOAuth2.GrantTypeDeviceCode)
	case clientCredentialsStr == "true":
		selectedMethod = string(svcOAuth2.GrantTypeClientCredentials)
	default:
		selectedMethod = string(svcOAuth2.GrantTypeAuthCode)
	}

	// Check if configuration exists for the selected method, prompt for setup if missing
	if err := ensureAuthConfigurationExists(selectedMethod); err != nil {
		// Configuration is missing - run interactive setup
		if configErr := RunInteractiveAuthConfig(os.Stdin); configErr != nil {
			return configErr
		}
	}

	// Perform authentication based on selected method flag
	var token *oauth2.Token
	var newAuth bool
	var location StorageLocation

	switch selectedMethod {
	case string(svcOAuth2.GrantTypeDeviceCode):
		token, newAuth, location, err = PerformDeviceCodeLogin(ctx)
		if err != nil {
			return handleLoginError(err, ctx, profileName)
		}
	case string(svcOAuth2.GrantTypeClientCredentials):
		token, newAuth, location, err = PerformClientCredentialsLogin(ctx)
		if err != nil {
			return handleLoginError(err, ctx, profileName)
		}
	default:
		token, newAuth, location, err = PerformAuthCodeLogin(ctx)
		if err != nil {
			return handleLoginError(err, ctx, profileName)
		}
	}

	// Display authentication result
	displayLoginSuccess(token, newAuth, location, selectedMethod, profileName)

	return nil
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
		return fmt.Errorf("%w: %s", ErrUnsupportedAuthMethod, authMethod)
	}
}

// performLoginByConfiguredType performs login using the configured authentication type
func performLoginByConfiguredType(ctx context.Context, authType, profileName string) error {
	var token *oauth2.Token
	var newAuth bool
	var location StorageLocation
	var err error
	var selectedMethod string

	switch authType {
	case customtypes.ENUM_PINGONE_AUTHENTICATION_TYPE_DEVICE_CODE:
		selectedMethod = string(svcOAuth2.GrantTypeDeviceCode)
		token, newAuth, location, err = PerformDeviceCodeLogin(ctx)
	case customtypes.ENUM_PINGONE_AUTHENTICATION_TYPE_CLIENT_CREDENTIALS:
		selectedMethod = string(svcOAuth2.GrantTypeClientCredentials)
		token, newAuth, location, err = PerformClientCredentialsLogin(ctx)
	case customtypes.ENUM_PINGONE_AUTHENTICATION_TYPE_AUTH_CODE:
		selectedMethod = string(svcOAuth2.GrantTypeAuthCode)
		token, newAuth, location, err = PerformAuthCodeLogin(ctx)
	default:
		return &errs.PingCLIError{
			Prefix: fmt.Sprintf("invalid authentication type: %s", authType),
			Err:    ErrInvalidAuthType,
		}
	}

	if err != nil {
		return handleLoginError(err, ctx, profileName)
	}

	displayLoginSuccess(token, newAuth, location, selectedMethod, profileName)

	return nil
}

// handleLoginError handles login errors and prompts for reconfiguration if needed
func handleLoginError(err error, ctx context.Context, profileName string) error {
	// Check if the error is related to missing configuration
	if isMissingConfigError(err) {
		output.Message(fmt.Sprintf("Authentication failed due to missing or invalid configuration: %v", err), nil)

		// Ask if user wants to reconfigure
		reconfigure, promptErr := PromptForReconfiguration(os.Stdin)
		if promptErr != nil {
			return &errs.PingCLIError{
				Prefix: "failed to prompt for reconfiguration",
				Err:    promptErr,
			}
		}

		if reconfigure {
			// Run interactive configuration
			if configErr := RunInteractiveAuthConfig(os.Stdin); configErr != nil {
				return configErr
			}

			// Retry login with new configuration
			authType, getErr := profiles.GetOptionValue(options.PingOneAuthenticationTypeOption)
			if getErr != nil {
				return &errs.PingCLIError{
					Prefix: "failed to read authentication type after reconfiguration",
					Err:    getErr,
				}
			}

			return performLoginByConfiguredType(ctx, authType, profileName)
		}

		return &errs.PingCLIError{
			Prefix: "authentication configuration required",
			Err:    ErrAuthConfigRequired,
		}
	}

	// Return the original error if it's not a configuration issue
	return err
}

// isMissingConfigError checks if an error is related to missing configuration
func isMissingConfigError(err error) bool {
	if err == nil {
		return false
	}

	// Check for common configuration-related errors
	errStr := strings.ToLower(err.Error())

	return strings.Contains(errStr, "missing") ||
		strings.Contains(errStr, "not found") ||
		strings.Contains(errStr, "not configured") ||
		strings.Contains(errStr, "required") ||
		strings.Contains(errStr, "no token")
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
