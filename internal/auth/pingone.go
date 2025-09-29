package auth_internal

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/pingidentity/pingcli/internal/configuration/options"
	"github.com/pingidentity/pingcli/internal/output"
	"github.com/pingidentity/pingcli/internal/profiles"
	"github.com/pingidentity/pingone-go-client/config"
	pingoneoauth2 "github.com/pingidentity/pingone-go-client/oauth2"
	"github.com/pingidentity/pingone-go-client/pingone"
)

var (
	pingoneAPIClient *pingone.APIClient
)

func savePingOneKoanfValues(accessToken, accessTokenExpiry string) error {
	pName, err := profiles.GetOptionValue(options.RootProfileOption)
	if err != nil {
		return err
	}
	if pName == "" {
		pName, err = profiles.GetOptionValue(options.RootActiveProfileOption)
		if err != nil {
			return err
		}
	}
	subKoanf, err := profiles.GetKoanfConfig().GetProfileKoanf(pName)
	if err != nil {
		return err
	}
	if err = subKoanf.Set(options.RequestAccessTokenOption.KoanfKey, accessToken); err != nil {
		return err
	}

	tokenExpiryInt, err := strconv.ParseInt(accessTokenExpiry, 10, 64)
	if err != nil {
		return err
	}

	if err = subKoanf.Set(options.RequestAccessTokenExpiryOption.KoanfKey, tokenExpiryInt); err != nil {
		return err
	}
	if err = profiles.GetKoanfConfig().SaveProfile(pName, subKoanf); err != nil {
		return err
	}
	return nil
}

func GetPingOneClient() (*pingone.APIClient, error) {
	pingoneAPIClientConfig, err := getPingOneConfiguration()
	if err != nil {
		return nil, err
	}

	if pingoneAPIClient == nil {
		pingOneNewConfig := pingone.NewConfiguration(pingoneAPIClientConfig)

		pingOneRegionCode, err := profiles.GetOptionValue(options.PingOneRegionCodeOption)
		if err != nil {
			return nil, err
		}

		isValidRegion := pingone.EnvironmentRegionCode(pingOneRegionCode).IsValid()

		if pingOneRegionCode == "" || !isValidRegion {
			return nil, fmt.Errorf("PingOne region code is required and must be valid.")
		}

		pingOneNewConfig.Service.WithRootDomain(pingOneRegionCode)

		pingoneAPIClient, err = pingone.NewAPIClient(pingOneNewConfig)
		if err != nil {
			return nil, err
		}
	}

	pingOneAPIClientConfig := pingoneAPIClient.GetConfig()

	accessToken := pingOneAPIClientConfig.Service.GetAccessToken()
	if accessToken == "" {
		return nil, fmt.Errorf("failed to create PingOne client: access token is empty")
	}

	accessTokenExpiry := pingOneAPIClientConfig.Service.GetAccessTokenExpiry()
	accessTokenExpiryStr := strconv.Itoa(accessTokenExpiry)

	if err := savePingOneKoanfValues(accessToken, accessTokenExpiryStr); err != nil {
		return nil, fmt.Errorf("failed to save PingOne access token: %w", err)
	}

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

func getPingOneConfiguration() (*config.Configuration, error) {
	pingOneConfig := config.NewConfiguration()

	authType, err := profiles.GetOptionValue(options.PingOneAuthenticationTypeOption)
	if err != nil {
		return nil, err
	}

	if !pingoneoauth2.IsValidGrantType(authType) && authType != "worker" {
		return nil, fmt.Errorf("unrecognized or unsupported PingOne authentication type: '%s'", authType)
	}

	switch pingoneoauth2.GrantType(authType) {
	case pingoneoauth2.GrantTypeAuthCode:
		authCodeClientID, err := profiles.GetOptionValue(options.PingOneAuthenticationAuthCodeClientIDOption)
		if err != nil {
			return nil, err
		}

		authCodeEnvId, err := profiles.GetOptionValue(options.PingOneAuthenticationAuthCodeEnvironmentIDOption)
		if err != nil {
			return nil, err
		}

		authCodePortStr, err := profiles.GetOptionValue(options.PingOneAuthenticationAuthCodePortOption)
		if err != nil {
			return nil, err
		}

		authCodeScopes, err := profiles.GetOptionValue(options.PingOneAuthenticationAuthCodeScopesOption)
		if err != nil {
			return nil, err
		}

		// Set the PingOne Configuration values
		authCodeScopesList := strings.Split(authCodeScopes, ",")

		pingOneConfig.Auth.AuthCode = &config.AuthCode{
			AuthCodeClientID:      &authCodeClientID,
			AuthCodeEnvironmentID: &authCodeEnvId,
			AuthCodePort:          &authCodePortStr,
			AuthCodeScopes:        &authCodeScopesList,
		}

	case pingoneoauth2.GrantTypeDeviceCode:
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

		*pingOneConfig.Auth.GrantType = pingoneoauth2.GrantType(authType)

		pingOneConfig.Auth.DeviceCode = &config.DeviceCode{
			DeviceCodeClientID:      &deviceCodeClientID,
			DeviceCodeEnvironmentID: &deviceCodeEnvId,
			DeviceCodeScopes:        &deviceCodeScopesList,
		}
	// Support original "worker" type as alias for "client_credentials"
	case pingoneoauth2.GrantTypeClientCredentials, "worker":
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

		environmentID, err := profiles.GetOptionValue(options.PingOneAuthenticationClientCredentialsEnvironmentIDOption)
		if err != nil {
			return nil, err
		}

		*pingOneConfig.Auth.GrantType = pingoneoauth2.GrantType(authType)

		pingOneConfig.Auth.ClientCredentials = &config.ClientCredentials{
			ClientCredentialsClientID:      &clientID,
			ClientCredentialsClientSecret:  &clientSecret,
			ClientCredentialsEnvironmentID: &environmentID,
		}
	default:
		return nil, fmt.Errorf("unrecognized or unsupported PingOne authentication type: '%s'", authType)
	}

	return pingOneConfig, nil
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

	return *pingOneClientConfig.Service.Auth.AccessToken, nil
}
