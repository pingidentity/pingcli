// Copyright © 2025 Ping Identity Corporation

package auth

import (
	"fmt"
	"io"
	"strings"

	auth_internal "github.com/pingidentity/pingcli/internal/auth"
	"github.com/pingidentity/pingcli/internal/configuration/options"
	"github.com/pingidentity/pingcli/internal/customtypes"
	"github.com/pingidentity/pingcli/internal/errs"
	"github.com/pingidentity/pingcli/internal/input"
	"github.com/pingidentity/pingcli/internal/output"
	"github.com/pingidentity/pingcli/internal/profiles"
	"github.com/pingidentity/pingone-go-client/config"
)

const (
	defaultRedirectURI = config.DefaultAuthCodeRedirectURI
)

var (
	loginInteractiveErrorPrefix = "failed to configure authentication"
)

// PromptForAuthType prompts the user to select an authentication type
// If showStatus is true, it will show (configured) or (not configured) status next to each option
func PromptForAuthType(rc io.ReadCloser, showStatus bool) (string, error) {
	authTypes := []string{
		customtypes.ENUM_PINGONE_AUTHENTICATION_TYPE_AUTH_CODE,
		customtypes.ENUM_PINGONE_AUTHENTICATION_TYPE_DEVICE_CODE,
		customtypes.ENUM_PINGONE_AUTHENTICATION_TYPE_CLIENT_CREDENTIALS,
	}

	// If showStatus is true, check which methods are configured and append status
	displayOptions := authTypes
	if showStatus {
		configStatus, err := getAuthMethodsConfigurationStatus()
		if err != nil {
			return "", err
		}

		displayOptions = make([]string, len(authTypes))
		for i, authType := range authTypes {
			if configStatus[authType] {
				displayOptions[i] = fmt.Sprintf("%s (configured)", authType)
			} else {
				displayOptions[i] = fmt.Sprintf("%s (not configured)", authType)
			}
		}
	}

	selectedOption, err := input.RunPromptSelect(
		"Select authentication type for this profile",
		displayOptions,
		rc,
	)
	if err != nil {
		return "", &errs.PingCLIError{Prefix: loginInteractiveErrorPrefix, Err: err}
	}

	// Extract the actual auth type from the display option (remove status text)
	selectedType := selectedOption
	if showStatus {
		// Find the matching auth type from the original list
		for i, displayOpt := range displayOptions {
			if displayOpt == selectedOption {
				selectedType = authTypes[i]

				break
			}
		}
	}

	return selectedType, nil
}

// PromptForAuthCodeConfig prompts for auth code configuration
func PromptForAuthCodeConfig(rc io.ReadCloser) (clientID, environmentID, redirectURIPath, redirectURIPort, scopes string, err error) {
	// Client ID (required)
	clientID, err = input.RunPrompt(
		"Authorization Code Client ID",
		func(s string) error {
			if strings.TrimSpace(s) == "" {
				return ErrClientIDRequired
			}

			return nil
		},
		rc,
	)
	if err != nil {
		return "", "", "", "", "", &errs.PingCLIError{Prefix: loginInteractiveErrorPrefix, Err: err}
	}

	// Environment ID (required)
	environmentID, err = input.RunPrompt(
		"PingOne Environment ID",
		func(s string) error {
			if strings.TrimSpace(s) == "" {
				return ErrEnvironmentIDRequired
			}

			return nil
		},
		rc,
	)
	if err != nil {
		return "", "", "", "", "", &errs.PingCLIError{Prefix: loginInteractiveErrorPrefix, Err: err}
	}

	// Redirect URI (optional, has default)
	output.Message(fmt.Sprintf("Redirect URI (press Enter for default: %s)", defaultRedirectURI), nil)
	redirectURIPath, err = input.RunPrompt(
		"Redirect URI path",
		nil, // No validation - optional
		rc,
	)

	if err != nil {
		return "", "", "", "", "", &errs.PingCLIError{Prefix: loginInteractiveErrorPrefix, Err: err}
	}
	if strings.TrimSpace(redirectURIPath) == "" {
		redirectURIPath = defaultRedirectURI
	}

	// Scopes (optional)
	output.Message("Scopes (optional, comma-separated)", nil)
	scopes, err = input.RunPrompt(
		"Scopes",
		nil, // No validation - optional
		rc,
	)
	if err != nil {
		return "", "", "", "", "", &errs.PingCLIError{Prefix: loginInteractiveErrorPrefix, Err: err}
	}

	return clientID, environmentID, redirectURIPath, redirectURIPort, scopes, nil
}

// PromptForDeviceCodeConfig prompts for device code configuration
func PromptForDeviceCodeConfig(rc io.ReadCloser) (clientID, environmentID, scopes string, err error) {
	// Client ID (required)
	clientID, err = input.RunPrompt(
		"Device Code Client ID",
		func(s string) error {
			if strings.TrimSpace(s) == "" {
				return ErrClientIDRequired
			}

			return nil
		},
		rc,
	)
	if err != nil {
		return "", "", "", &errs.PingCLIError{Prefix: loginInteractiveErrorPrefix, Err: err}
	}

	// Environment ID (required)
	environmentID, err = input.RunPrompt(
		"PingOne Environment ID",
		func(s string) error {
			if strings.TrimSpace(s) == "" {
				return ErrEnvironmentIDRequired
			}

			return nil
		},
		rc,
	)
	if err != nil {
		return "", "", "", &errs.PingCLIError{Prefix: loginInteractiveErrorPrefix, Err: err}
	}

	// Scopes (optional)
	output.Message("Scopes (optional, comma-separated)", nil)
	scopes, err = input.RunPrompt(
		"Scopes",
		nil, // No validation - optional
		rc,
	)
	if err != nil {
		return "", "", "", &errs.PingCLIError{Prefix: loginInteractiveErrorPrefix, Err: err}
	}

	return clientID, environmentID, scopes, nil
}

// PromptForClientCredentialsConfig prompts for client credentials configuration
func PromptForClientCredentialsConfig(rc io.ReadCloser) (clientID, clientSecret, environmentID, scopes string, err error) {
	// Client ID (required)
	clientID, err = input.RunPrompt(
		"Client Credentials Client ID",
		func(s string) error {
			if strings.TrimSpace(s) == "" {
				return ErrClientIDRequired
			}

			return nil
		},
		rc,
	)
	if err != nil {
		return "", "", "", "", &errs.PingCLIError{Prefix: loginInteractiveErrorPrefix, Err: err}
	}

	// Client Secret (required)
	clientSecret, err = input.RunPrompt(
		"Client Credentials Client Secret",
		func(s string) error {
			if strings.TrimSpace(s) == "" {
				return ErrClientSecretRequired
			}

			return nil
		},
		rc,
	)
	if err != nil {
		return "", "", "", "", &errs.PingCLIError{Prefix: loginInteractiveErrorPrefix, Err: err}
	}

	// Environment ID (required)
	environmentID, err = input.RunPrompt(
		"PingOne Environment ID",
		func(s string) error {
			if strings.TrimSpace(s) == "" {
				return ErrEnvironmentIDRequired
			}

			return nil
		},
		rc,
	)
	if err != nil {
		return "", "", "", "", &errs.PingCLIError{Prefix: loginInteractiveErrorPrefix, Err: err}
	}

	// Scopes (optional)
	output.Message("Scopes (optional, comma-separated)", nil)
	scopes, err = input.RunPrompt(
		"Scopes",
		nil, // No validation - optional
		rc,
	)
	if err != nil {
		return "", "", "", "", &errs.PingCLIError{Prefix: loginInteractiveErrorPrefix, Err: err}
	}

	return clientID, clientSecret, environmentID, scopes, nil
}

// SaveAuthConfigToProfile saves the authentication configuration to the active profile
func SaveAuthConfigToProfile(authType, clientID, clientSecret, environmentID, redirectURIPath, redirectURIport, scopes string) error {
	koanfConfig, err := profiles.GetKoanfConfig()
	if err != nil {
		return &errs.PingCLIError{Prefix: loginInteractiveErrorPrefix, Err: err}
	}

	profileName, err := profiles.GetOptionValue(options.RootActiveProfileOption)
	if err != nil {
		return &errs.PingCLIError{Prefix: loginInteractiveErrorPrefix, Err: err}
	}

	subKoanf, err := koanfConfig.GetProfileKoanf(profileName)
	if err != nil {
		return &errs.PingCLIError{Prefix: loginInteractiveErrorPrefix, Err: err}
	}

	// Set the authentication type
	if err = subKoanf.Set(options.PingOneAuthenticationTypeOption.KoanfKey, authType); err != nil {
		return &errs.PingCLIError{Prefix: loginInteractiveErrorPrefix, Err: err}
	}

	// Save type-specific configuration
	switch authType {
	case customtypes.ENUM_PINGONE_AUTHENTICATION_TYPE_AUTH_CODE:
		if err = subKoanf.Set(options.PingOneAuthenticationAuthCodeClientIDOption.KoanfKey, clientID); err != nil {
			return &errs.PingCLIError{Prefix: loginInteractiveErrorPrefix, Err: err}
		}
		if err = subKoanf.Set(options.PingOneAuthenticationAuthCodeEnvironmentIDOption.KoanfKey, environmentID); err != nil {
			return &errs.PingCLIError{Prefix: loginInteractiveErrorPrefix, Err: err}
		}
		if redirectURIPath != "" {
			if err = subKoanf.Set(options.PingOneAuthenticationAuthCodeRedirectURIPathOption.KoanfKey, redirectURIPath); err != nil {
				return &errs.PingCLIError{Prefix: loginInteractiveErrorPrefix, Err: err}
			}
		}
		if redirectURIport != "" {
			if err = subKoanf.Set(options.PingOneAuthenticationAuthCodeRedirectURIPortOption.KoanfKey, redirectURIport); err != nil {
				return &errs.PingCLIError{Prefix: loginInteractiveErrorPrefix, Err: err}
			}
		}
		if scopes != "" {
			if err = subKoanf.Set(options.PingOneAuthenticationAuthCodeScopesOption.KoanfKey, scopes); err != nil {
				return &errs.PingCLIError{Prefix: loginInteractiveErrorPrefix, Err: err}
			}
		}

	case customtypes.ENUM_PINGONE_AUTHENTICATION_TYPE_DEVICE_CODE:
		if err = subKoanf.Set(options.PingOneAuthenticationDeviceCodeClientIDOption.KoanfKey, clientID); err != nil {
			return &errs.PingCLIError{Prefix: loginInteractiveErrorPrefix, Err: err}
		}
		if err = subKoanf.Set(options.PingOneAuthenticationDeviceCodeEnvironmentIDOption.KoanfKey, environmentID); err != nil {
			return &errs.PingCLIError{Prefix: loginInteractiveErrorPrefix, Err: err}
		}
		if scopes != "" {
			if err = subKoanf.Set(options.PingOneAuthenticationDeviceCodeScopesOption.KoanfKey, scopes); err != nil {
				return &errs.PingCLIError{Prefix: loginInteractiveErrorPrefix, Err: err}
			}
		}

	case customtypes.ENUM_PINGONE_AUTHENTICATION_TYPE_CLIENT_CREDENTIALS:
		if err = subKoanf.Set(options.PingOneAuthenticationClientCredentialsClientIDOption.KoanfKey, clientID); err != nil {
			return &errs.PingCLIError{Prefix: loginInteractiveErrorPrefix, Err: err}
		}
		if err = subKoanf.Set(options.PingOneAuthenticationClientCredentialsClientSecretOption.KoanfKey, clientSecret); err != nil {
			return &errs.PingCLIError{Prefix: loginInteractiveErrorPrefix, Err: err}
		}
		if err = subKoanf.Set(options.PingOneAuthenticationClientCredentialsEnvironmentIDOption.KoanfKey, environmentID); err != nil {
			return &errs.PingCLIError{Prefix: loginInteractiveErrorPrefix, Err: err}
		}
		if scopes != "" {
			if err = subKoanf.Set(options.PingOneAuthenticationClientCredentialsScopesOption.KoanfKey, scopes); err != nil {
				return &errs.PingCLIError{Prefix: loginInteractiveErrorPrefix, Err: err}
			}
		}
	}

	// Save the profile
	if err = koanfConfig.SaveProfile(profileName, subKoanf); err != nil {
		return &errs.PingCLIError{Prefix: loginInteractiveErrorPrefix, Err: err}
	}

	output.Success(fmt.Sprintf("Authentication configuration saved to profile '%s'", profileName), nil)

	return nil
}

// RunInteractiveAuthConfig runs the full interactive authentication configuration flow
func RunInteractiveAuthConfig(rc io.ReadCloser) error {
	// Check if any authentication methods are already configured
	configStatus, err := getAuthMethodsConfigurationStatus()
	if err != nil {
		return err
	}

	// Count how many methods are configured
	configuredCount := 0
	for _, configured := range configStatus {
		if configured {
			configuredCount++
		}
	}

	// Determine if we should show status and what message to display
	showStatus := configuredCount > 0
	if showStatus {
		output.Message("Select an authentication method", nil)
	} else {
		output.Message("No authentication methods configured. Let's set one up!", nil)
	}

	// Step 1: Ask for auth type (with or without status indicators)
	authType, err := PromptForAuthType(rc, showStatus)
	if err != nil {
		return err
	}

	// Step 2: Check if this specific auth type has existing credentials
	hasExistingCredentials := configStatus[authType]

	if hasExistingCredentials {
		useExisting, err := input.RunPromptConfirm(
			fmt.Sprintf("Use existing %s credentials", authType),
			rc,
		)
		if err != nil {
			return &errs.PingCLIError{Prefix: loginInteractiveErrorPrefix, Err: err}
		}

		if useExisting {
			// Validate that the existing configuration is complete
			var validationErr error
			switch authType {
			case customtypes.ENUM_PINGONE_AUTHENTICATION_TYPE_AUTH_CODE:
				_, validationErr = auth_internal.GetAuthCodeConfiguration()
			case customtypes.ENUM_PINGONE_AUTHENTICATION_TYPE_DEVICE_CODE:
				_, validationErr = auth_internal.GetDeviceCodeConfiguration()
			case customtypes.ENUM_PINGONE_AUTHENTICATION_TYPE_CLIENT_CREDENTIALS:
				_, validationErr = auth_internal.GetClientCredentialsConfiguration()
			}

			if validationErr == nil {
				// Configuration is valid - just save the auth type and return
				return SaveAuthTypeOnly(authType)
			}

			// Configuration exists but is invalid/incomplete
			output.Message(fmt.Sprintf("Existing configuration is incomplete: %v", validationErr), nil)
			output.Message("Let's complete the configuration...", nil)
		} else {
			// User wants to reconfigure, continue with prompts
			output.Message("Let's reconfigure the credentials...", nil)
		}
	}

	var clientID, clientSecret, environmentID, redirectURIPath, redirectURIPort, scopes string

	// Step 3: Collect configuration based on selected type
	switch authType {
	case customtypes.ENUM_PINGONE_AUTHENTICATION_TYPE_AUTH_CODE:
		clientID, environmentID, redirectURIPath, redirectURIPort, scopes, err = PromptForAuthCodeConfig(rc)
		if err != nil {
			return err
		}

	case customtypes.ENUM_PINGONE_AUTHENTICATION_TYPE_DEVICE_CODE:
		clientID, environmentID, scopes, err = PromptForDeviceCodeConfig(rc)
		if err != nil {
			return err
		}

	case customtypes.ENUM_PINGONE_AUTHENTICATION_TYPE_CLIENT_CREDENTIALS:
		clientID, clientSecret, environmentID, scopes, err = PromptForClientCredentialsConfig(rc)
		if err != nil {
			return err
		}
	}

	// Step 4: Save configuration to profile
	return SaveAuthConfigToProfile(authType, clientID, clientSecret, environmentID, redirectURIPath, redirectURIPort, scopes)
}

// PromptForReconfiguration asks the user if they want to reconfigure authentication
func PromptForReconfiguration(rc io.ReadCloser) (bool, error) {
	return input.RunPromptConfirm("Do you want to reconfigure authentication", rc)
}

// checkExistingCredentials checks if credentials already exist for the given auth type
func checkExistingCredentials(authType string) (bool, error) {
	koanfConfig, err := profiles.GetKoanfConfig()
	if err != nil {
		return false, &errs.PingCLIError{Prefix: loginInteractiveErrorPrefix, Err: err}
	}

	profileName, err := profiles.GetOptionValue(options.RootActiveProfileOption)
	if err != nil {
		return false, &errs.PingCLIError{Prefix: loginInteractiveErrorPrefix, Err: err}
	}

	subKoanf, err := koanfConfig.GetProfileKoanf(profileName)
	if err != nil {
		return false, &errs.PingCLIError{Prefix: loginInteractiveErrorPrefix, Err: err}
	}

	// Check for type-specific required credentials
	switch authType {
	case customtypes.ENUM_PINGONE_AUTHENTICATION_TYPE_AUTH_CODE:
		clientID := subKoanf.String(options.PingOneAuthenticationAuthCodeClientIDOption.KoanfKey)
		environmentID := subKoanf.String(options.PingOneAuthenticationAuthCodeEnvironmentIDOption.KoanfKey)

		return clientID != "" && environmentID != "", nil

	case customtypes.ENUM_PINGONE_AUTHENTICATION_TYPE_DEVICE_CODE:
		clientID := subKoanf.String(options.PingOneAuthenticationDeviceCodeClientIDOption.KoanfKey)
		environmentID := subKoanf.String(options.PingOneAuthenticationDeviceCodeEnvironmentIDOption.KoanfKey)

		return clientID != "" && environmentID != "", nil

	case customtypes.ENUM_PINGONE_AUTHENTICATION_TYPE_CLIENT_CREDENTIALS:
		clientID := subKoanf.String(options.PingOneAuthenticationClientCredentialsClientIDOption.KoanfKey)
		clientSecret := subKoanf.String(options.PingOneAuthenticationClientCredentialsClientSecretOption.KoanfKey)
		environmentID := subKoanf.String(options.PingOneAuthenticationClientCredentialsEnvironmentIDOption.KoanfKey)

		return clientID != "" && clientSecret != "" && environmentID != "", nil
	}

	return false, nil
}

// getAuthMethodsConfigurationStatus returns a map of auth types to their configuration status
func getAuthMethodsConfigurationStatus() (map[string]bool, error) {
	authTypes := []string{
		customtypes.ENUM_PINGONE_AUTHENTICATION_TYPE_AUTH_CODE,
		customtypes.ENUM_PINGONE_AUTHENTICATION_TYPE_DEVICE_CODE,
		customtypes.ENUM_PINGONE_AUTHENTICATION_TYPE_CLIENT_CREDENTIALS,
	}

	status := make(map[string]bool)
	for _, authType := range authTypes {
		configured, err := checkExistingCredentials(authType)
		if err != nil {
			return nil, err
		}
		status[authType] = configured
	}

	return status, nil
}

// SaveAuthTypeOnly saves just the authentication type without modifying existing credentials
func SaveAuthTypeOnly(authType string) error {
	koanfConfig, err := profiles.GetKoanfConfig()
	if err != nil {
		return &errs.PingCLIError{Prefix: loginInteractiveErrorPrefix, Err: err}
	}

	profileName, err := profiles.GetOptionValue(options.RootActiveProfileOption)
	if err != nil {
		return &errs.PingCLIError{Prefix: loginInteractiveErrorPrefix, Err: err}
	}

	subKoanf, err := koanfConfig.GetProfileKoanf(profileName)
	if err != nil {
		return &errs.PingCLIError{Prefix: loginInteractiveErrorPrefix, Err: err}
	}

	// Set only the authentication type
	if err = subKoanf.Set(options.PingOneAuthenticationTypeOption.KoanfKey, authType); err != nil {
		return &errs.PingCLIError{Prefix: loginInteractiveErrorPrefix, Err: err}
	}

	// Save the profile
	if err = koanfConfig.SaveProfile(profileName, subKoanf); err != nil {
		return &errs.PingCLIError{Prefix: loginInteractiveErrorPrefix, Err: err}
	}

	output.Success(fmt.Sprintf("Authentication type set to '%s' for profile '%s'", authType, profileName), nil)

	return nil
}
