package auth_internal

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/pingidentity/pingcli/internal/configuration/options"
	"github.com/pingidentity/pingcli/internal/customtypes"
	"github.com/pingidentity/pingcli/internal/output"
	"github.com/pingidentity/pingcli/internal/profiles"
	"github.com/pingidentity/pingone-go-client/config"
	pingoneoauth2 "github.com/pingidentity/pingone-go-client/oauth2"
	"github.com/pingidentity/pingone-go-client/pingone"
)

var (
	pingoneAPIClient *pingone.APIClient
)

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

	//accessToken := pingoneAPIClientConfig.GetAccessToken()
	//if accessToken == "" {
	//	return nil, fmt.Errorf("failed to create PingOne client: access token is empty")
	//
	//accessTokenExpiry := pingoneAPIClientConfig.GetAccessTokenExpiry()
	//accessTokenExpiryStr := strconv.Itoa(accessTokenExpir//
	//if err := savePingOneKoanfValues(accessToken, accessTokenExpiryStr); err != nil {
	//	return nil, fmt.Errorf("failed to save PingOne access token: %w", err)
	//}

	return pingoneAPIClient, nil
}

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

		authCodeScopesList := strings.Split(authCodeScopes, ",")

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

		deviceCodeScopesList := strings.Split(deviceCodeScopes, ",")

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

		// Scopes are mandatory - no defaults allowed
		// if len(clientCredentialsScopes) == 0 {
		// 	return nil, fmt.Errorf("client credentials scopes are required but not provided - user must supply scopes")
		// }

		clientCredentialsScopesList := strings.Split(clientCredentialsScopes, ",")
		// Trim whitespace from each scope
		for i, scope := range clientCredentialsScopesList {
			clientCredentialsScopesList[i] = strings.TrimSpace(scope)
		}
		// Filter out empty scopes
		filteredScopes := make([]string, 0)
		for _, scope := range clientCredentialsScopesList {
			if scope != "" {
				filteredScopes = append(filteredScopes, scope)
			}
		}

		configConfiguration.WithClientCredentialsScopes(filteredScopes)

		configConfiguration.WithClientCredentialsClientID(clientID)
		configConfiguration.WithClientCredentialsClientSecret(clientSecret)
		configConfiguration.WithClientCredentialsEnvironmentID(clientCredentialsEnvironmentID)

	default:
		return nil, fmt.Errorf("unrecognized or unsupported PingOne authentication type: '%s'", authType)
	}

	configConfiguration.WithGrantType(pingoneoauth2.GrantType(authType))

	pingOneRegionCode, err := profiles.GetOptionValue(options.PingOneRegionCodeOption)
	if err != nil {
		return nil, err
	}

	switch pingOneRegionCode {
	case customtypes.ENUM_PINGONE_REGION_CODE_AP:
		configConfiguration.WithTopLevelDomain(pingoneoauth2.TopLevelDomainAPAC)
	case customtypes.ENUM_PINGONE_REGION_CODE_AU:
		configConfiguration.WithTopLevelDomain(pingoneoauth2.TopLevelDomainAU)
	case customtypes.ENUM_PINGONE_REGION_CODE_CA:
		configConfiguration.WithTopLevelDomain(pingoneoauth2.TopLevelDomainCA)
	case customtypes.ENUM_PINGONE_REGION_CODE_EU:
		configConfiguration.WithTopLevelDomain(pingoneoauth2.TopLevelDomainEU)
	case customtypes.ENUM_PINGONE_REGION_CODE_NA:
		configConfiguration.WithTopLevelDomain(pingoneoauth2.TopLevelDomainNA)
	case customtypes.ENUM_PINGONE_REGION_CODE_SG:
		configConfiguration.WithTopLevelDomain(pingoneoauth2.TopLevelDomainSG)
	default:
		return nil, fmt.Errorf("PingOne region code is required and must be valid.")
	}

	return configConfiguration, nil
}

func pingOneAuth() (accessToken string, err error) {
	pingOneClient, err := GetPingOneClient()
	if err != nil {
		return "", err
	}

	pingOneClientConfig := pingOneClient.GetConfig()
	if pingOneClientConfig == nil {
		return "", fmt.Errorf("PingOne client configuration is nil")
	}

	accessToken, err = pingOneClientConfig.Service.GetAccessToken(context.Background())
	if err != nil {
		return "", err
	}

	return accessToken, nil
}
