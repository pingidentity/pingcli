package auth_internal

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/pingidentity/pingcli/internal/configuration/options"
	"github.com/pingidentity/pingcli/internal/customtypes"
	"github.com/pingidentity/pingcli/internal/errs"
	"github.com/pingidentity/pingcli/internal/output"
	"github.com/pingidentity/pingcli/internal/profiles"
	"github.com/pingidentity/pingone-go-client/config"
	pingoneoauth2 "github.com/pingidentity/pingone-go-client/oauth2"
	"github.com/pingidentity/pingone-go-client/pingone"
)

var (
	pingoneAPIClient *pingone.APIClient
)

// ClearPingOneClientCache clears the cached PingOne API client instance, forcing re-initialization on next use
func ClearPingOneClientCache() {
	pingoneAPIClient = nil
}

// GetAuthenticatedPingOneClient returns a PingOne API client instance with valid authentication credentials
func GetAuthenticatedPingOneClient(ctx context.Context) (*pingone.APIClient, error) {
	// Get a valid token source (will handle caching and refresh)
	tokenSource, err := GetValidTokenSource(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get valid token source: %w", err)
	}

	// Get a valid token
	token, err := tokenSource.Token()
	if err != nil {
		return nil, fmt.Errorf("failed to get valid token: %w", err)
	}

	// Create configuration with the access token
	configConfiguration, err := getConfigConfigurationWithToken(token.AccessToken)
	if err != nil {
		return nil, fmt.Errorf("failed to create config configuration: %w", err)
	}

	pingoneConfiguration := pingone.NewConfiguration(configConfiguration)

	apiClient, err := pingone.NewAPIClient(pingoneConfiguration)
	if err != nil {
		return nil, fmt.Errorf("failed to create PingOne API client: %w", err)
	}

	return apiClient, nil
}

// GetPingOneClient returns the cached PingOne API client instance or creates a new one if not already initialized
func GetPingOneClient() (*pingone.APIClient, error) {
	if pingoneAPIClient != nil {
		return pingoneAPIClient, nil
	}

	configConfiguration, err := getConfigConfiguration()
	if err != nil {
		return nil, err
	}
	pingoneConfiguration := pingone.NewConfiguration(configConfiguration)

	pingoneAPIClient, err = pingone.NewAPIClient(pingoneConfiguration)
	if err != nil {
		return nil, err
	}

	return pingoneAPIClient, nil
}

// GetPingOneAccessToken retrieves a valid access token, either from cache or by performing new authentication
func GetPingOneAccessToken() (accessToken string, err error) {
	// Check if existing access token is available
	accessToken, err = profiles.GetOptionValue(options.RequestAccessTokenOption)
	if err != nil {
		return "", err
	}

	if accessToken != "" {
		accessTokenExpiry, err := profiles.GetOptionValue(options.RequestAccessTokenExpiryOption)
		if err != nil {
			return "", err
		}

		if accessTokenExpiry == "" {
			accessTokenExpiry = "0"
		}

		// convert expiry string to int
		tokenExpiryInt, err := strconv.ParseInt(accessTokenExpiry, 10, 64)
		if err != nil {
			return "", err
		}

		// Get current Unix epoch time in seconds
		currentEpochSeconds := time.Now().Unix()

		// Return access token if it is still valid
		if currentEpochSeconds < tokenExpiryInt {
			return accessToken, nil
		}
	}

	output.Message("PingOne access token does not exist or is expired, requesting a new token...", nil)

	// If no valid access token is available, login and get a new one
	return pingOneAuth()
}

// getConfigConfiguration builds a PingOne SDK configuration from the active profile settings
func getConfigConfiguration() (*config.Configuration, error) {
	configConfiguration := config.NewConfiguration()

	authType, err := profiles.GetOptionValue(options.PingOneAuthenticationTypeOption)
	if err != nil {
		return nil, err
	}

	switch authType {
	case customtypes.ENUM_PINGONE_AUTHENTICATION_TYPE_AUTH_CODE:
		authCodeClientID, err := profiles.GetOptionValue(options.PingOneAuthenticationAuthCodeClientIDOption)
		if err != nil {
			return nil, err
		}

		authCodeEnvId, err := profiles.GetOptionValue(options.PingOneAuthenticationAuthCodeEnvironmentIDOption)
		if err != nil {
			return nil, err
		}

		authCodeScopes, err := profiles.GetOptionValue(options.PingOneAuthenticationAuthCodeScopesOption)
		if err != nil {
			return nil, err
		}

		redirectURI, err := profiles.GetOptionValue(options.PingOneAuthenticationAuthCodeRedirectURIOption)
		if err != nil {
			return nil, err
		}

		authCodeScopesList := parseScopesList(authCodeScopes)

		configConfiguration.WithAuthCodeClientID(authCodeClientID)
		configConfiguration.WithAuthCodeEnvironmentID(authCodeEnvId)
		configConfiguration.WithAuthCodeScopes(authCodeScopesList)
		configConfiguration.WithAuthCodeRedirectURI(redirectURI)

	case customtypes.ENUM_PINGONE_AUTHENTICATION_TYPE_DEVICE_CODE:
		deviceCodeClientID, err := profiles.GetOptionValue(options.PingOneAuthenticationDeviceCodeClientIDOption)
		if err != nil {
			return nil, err
		}

		deviceCodeEnvId, err := profiles.GetOptionValue(options.PingOneAuthenticationDeviceCodeEnvironmentIDOption)
		if err != nil {
			return nil, err
		}

		deviceCodeScopes, err := profiles.GetOptionValue(options.PingOneAuthenticationDeviceCodeScopesOption)
		if err != nil {
			return nil, err
		}

		deviceCodeScopesList := parseScopesList(deviceCodeScopes)

		configConfiguration.WithDeviceCodeClientID(deviceCodeClientID)
		configConfiguration.WithDeviceCodeEnvironmentID(deviceCodeEnvId)
		configConfiguration.WithDeviceCodeScopes(deviceCodeScopesList)

	// Support original "worker" type as alias for "client_credentials"
	case customtypes.ENUM_PINGONE_AUTHENTICATION_TYPE_CLIENT_CREDENTIALS, customtypes.ENUM_PINGONE_AUTHENTICATION_TYPE_WORKER:
		if authType == "worker" {
			authType = "client_credentials"
		}

		clientID, err := profiles.GetOptionValue(options.PingOneAuthenticationClientCredentialsClientIDOption)
		if err != nil {
			return nil, err
		}

		clientSecret, err := profiles.GetOptionValue(options.PingOneAuthenticationClientCredentialsClientSecretOption)
		if err != nil {
			return nil, err
		}

		clientCredentialsEnvironmentID, err := profiles.GetOptionValue(options.PingOneAuthenticationClientCredentialsEnvironmentIDOption)
		if err != nil {
			return nil, err
		}

		clientCredentialsScopes, err := profiles.GetOptionValue(options.PingOneAuthenticationClientCredentialsScopesOption)
		if err != nil {
			return nil, err
		}

		clientCredentialsScopesList := parseScopesList(clientCredentialsScopes)

		configConfiguration.WithClientCredentialsScopes(clientCredentialsScopesList)

		configConfiguration.WithClientCredentialsClientID(clientID)
		configConfiguration.WithClientCredentialsClientSecret(clientSecret)
		configConfiguration.WithClientCredentialsEnvironmentID(clientCredentialsEnvironmentID)

	default:
		return nil, &errs.PingCLIError{
			Prefix: fmt.Sprintf("failed to get configuration for authentication type '%s'", authType),
			Err:    ErrPingOneUnrecognizedAuthType,
		}
	}

	configConfiguration.WithGrantType(pingoneoauth2.GrantType(authType))

	// Apply region configuration
	configConfiguration, err = applyRegionConfiguration(configConfiguration)
	if err != nil {
		return nil, err
	}

	return configConfiguration, nil
}

// getConfigConfigurationWithToken creates a PingOne SDK configuration with a specific access token
func getConfigConfigurationWithToken(accessToken string) (*config.Configuration, error) {
	configConfiguration := config.NewConfiguration()

	// Set the access token directly
	configConfiguration.WithAccessToken(accessToken)

	// Apply region configuration
	configConfiguration, err := applyRegionConfiguration(configConfiguration)
	if err != nil {
		return nil, err
	}

	return configConfiguration, nil
}

// pingOneAuth performs PingOne authentication and returns the access token
func pingOneAuth() (accessToken string, err error) {
	pingOneClient, err := GetPingOneClient()
	if err != nil {
		return "", err
	}

	pingOneClientConfig := pingOneClient.GetConfig()
	if pingOneClientConfig == nil {
		return "", ErrPingOneClientConfigNil
	}

	accessToken, err = pingOneClientConfig.Service.GetAccessToken(context.Background())
	if err != nil {
		return "", err
	}

	return accessToken, nil
}
