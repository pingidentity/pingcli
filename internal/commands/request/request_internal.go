// Copyright © 2025 Ping Identity Corporation

package request_internal

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/pingidentity/pingcli/internal/configuration/options"
	"github.com/pingidentity/pingcli/internal/customtypes"
	"github.com/pingidentity/pingcli/internal/output"
	"github.com/pingidentity/pingcli/internal/profiles"
)

type PingOneAuthResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int64  `json:"expires_in"`
}

func RunInternalRequest(uri string) (err error) {
	service, err := profiles.GetOptionValue(options.RequestServiceOption)
	if err != nil {
		return fmt.Errorf("failed to send custom request: %w", err)
	}

	if service == "" {
		return fmt.Errorf("failed to send custom request: service is required")
	}

	switch service {
	case customtypes.ENUM_REQUEST_SERVICE_PINGONE:
		err = runInternalPingOneRequest(uri)
		if err != nil {
			return fmt.Errorf("failed to send custom request: %w", err)
		}
	default:
		return fmt.Errorf("failed to send custom request: unrecognized service '%s'", service)
	}

	return nil
}

func runInternalPingOneRequest(uri string) (err error) {
	accessToken, err := pingoneAccessToken()
	if err != nil {
		return err
	}

	topLevelDomain, err := getTopLevelDomain()
	if err != nil {
		return err
	}

	failOption, err := profiles.GetOptionValue(options.RequestFailOption)
	if err != nil {
		return err
	}

	apiURL := fmt.Sprintf("https://api.pingone.%s/v1/%s", topLevelDomain, uri)

	httpMethod, err := profiles.GetOptionValue(options.RequestHTTPMethodOption)
	if err != nil {
		return err
	}

	if httpMethod == "" {
		return fmt.Errorf("http method is required")
	}

	data, err := getDataRaw()
	if err != nil {
		return err
	}

	if data == "" {
		data, err = getDataFile()
		if err != nil {
			return err
		}
	}

	payload := strings.NewReader(data)

	client := &http.Client{}
	req, err := http.NewRequestWithContext(context.Background(), httpMethod, apiURL, payload)
	if err != nil {
		return err
	}

	headers, err := profiles.GetOptionValue(options.RequestHeaderOption)
	if err != nil {
		return err
	}

	requestHeaders := new(customtypes.HeaderSlice)
	err = requestHeaders.Set(headers)
	if err != nil {
		return err
	}

	requestHeaders.SetHttpRequestHeaders(req)

	// Set default content type if not provided
	if req.Header.Get("Content-Type") == "" {
		req.Header.Add("Content-Type", "application/json")
	}

	// Set default authorization header
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", accessToken))

	res, err := client.Do(req)
	if err != nil {
		return err
	}

	defer func() {
		cErr := res.Body.Close()
		if cErr != nil {
			err = errors.Join(err, cErr)
		}
	}()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	fields := map[string]any{
		"response": json.RawMessage(body),
		"status":   res.StatusCode,
	}

	if res.StatusCode < 200 || res.StatusCode >= 300 {
		output.UserError("Failed Custom Request", fields)
		if failOption == "true" {
			// Allow response body to clean up before exiting
			defer os.Exit(1)

			return nil
		}
	} else {
		output.Success("Custom request successful", fields)
	}

	return nil
}

func getTopLevelDomain() (topLevelDomain string, err error) {
	pingoneRegionCode, err := profiles.GetOptionValue(options.PingOneRegionCodeOption)
	if err != nil {
		return "", err
	}

	if pingoneRegionCode == "" {
		return "", fmt.Errorf("PingOne region code is required")
	}

	switch pingoneRegionCode {
	case customtypes.ENUM_PINGONE_REGION_CODE_AP:
		topLevelDomain = customtypes.ENUM_PINGONE_TLD_AP
	case customtypes.ENUM_PINGONE_REGION_CODE_AU:
		topLevelDomain = customtypes.ENUM_PINGONE_TLD_AU
	case customtypes.ENUM_PINGONE_REGION_CODE_CA:
		topLevelDomain = customtypes.ENUM_PINGONE_TLD_CA
	case customtypes.ENUM_PINGONE_REGION_CODE_EU:
		topLevelDomain = customtypes.ENUM_PINGONE_TLD_EU
	case customtypes.ENUM_PINGONE_REGION_CODE_NA:
		topLevelDomain = customtypes.ENUM_PINGONE_TLD_NA
	default:
		return "", fmt.Errorf("unrecognized PingOne region code: '%s'", pingoneRegionCode)
	}

	return topLevelDomain, nil
}

func pingoneAccessToken() (accessToken string, err error) {
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
	return pingoneAuth()
}

func pingoneAuth() (accessToken string, err error) {
	topLevelDomain, err := getTopLevelDomain()
	if err != nil {
		return "", err
	}

	workerEnvId, err := profiles.GetOptionValue(options.PingOneAuthenticationWorkerEnvironmentIDOption)
	if err != nil {
		return "", err
	}

	if workerEnvId == "" {
		return "", fmt.Errorf("PingOne worker environment ID is required")
	}

	authURL := fmt.Sprintf("https://auth.pingone.%s/%s/as/token", topLevelDomain, workerEnvId)

	clientId, err := profiles.GetOptionValue(options.PingOneAuthenticationWorkerClientIDOption)
	if err != nil {
		return "", err
	}
	clientSecret, err := profiles.GetOptionValue(options.PingOneAuthenticationWorkerClientSecretOption)
	if err != nil {
		return "", err
	}

	if clientId == "" || clientSecret == "" {
		return "", fmt.Errorf("PingOne client ID and secret are required")
	}

	basicAuthBase64 := base64.StdEncoding.EncodeToString([]byte(clientId + ":" + clientSecret))

	payload := strings.NewReader("grant_type=client_credentials")

	client := &http.Client{}
	req, err := http.NewRequestWithContext(context.Background(), customtypes.ENUM_HTTP_METHOD_POST, authURL, payload)
	if err != nil {
		return "", err
	}

	req.Header.Add("Authorization", fmt.Sprintf("Basic %s", basicAuthBase64))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	res, err := client.Do(req)
	if err != nil {
		return "", err
	}

	defer func() {
		cErr := res.Body.Close()
		if cErr != nil {
			err = errors.Join(err, cErr)
		}
	}()

	responseBodyBytes, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	if res.StatusCode < 200 || res.StatusCode >= 300 {
		return "", fmt.Errorf("failed to authenticate with PingOne: Response Status %s: Response Body %s", res.Status, string(responseBodyBytes))
	}

	pingoneAuthResponse := new(PingOneAuthResponse)
	err = json.Unmarshal(responseBodyBytes, pingoneAuthResponse)
	if err != nil {
		return "", err
	}

	currentTime := time.Now().Unix()
	tokenExpiry := currentTime + pingoneAuthResponse.ExpiresIn

	// Store access token and expiry
	pName, err := profiles.GetOptionValue(options.RootProfileOption)
	if err != nil {
		return "", err
	}

	if pName == "" {
		pName, err = profiles.GetOptionValue(options.RootActiveProfileOption)
		if err != nil {
			return "", err
		}
	}

	subKoanf, err := profiles.GetKoanfConfig().GetProfileKoanf(pName)
	if err != nil {
		return "", err
	}

	err = subKoanf.Set(options.RequestAccessTokenOption.KoanfKey, pingoneAuthResponse.AccessToken)
	if err != nil {
		return "", err
	}
	err = subKoanf.Set(options.RequestAccessTokenExpiryOption.KoanfKey, tokenExpiry)
	if err != nil {
		return "", err
	}
	err = profiles.GetKoanfConfig().SaveProfile(pName, subKoanf)
	if err != nil {
		return "", err
	}

	return pingoneAuthResponse.AccessToken, nil
}

func getDataFile() (data string, err error) {
	dataFilepath, err := profiles.GetOptionValue(options.RequestDataOption)
	if err != nil {
		return "", err
	}

	if dataFilepath != "" {
		dataFilepath = filepath.Clean(dataFilepath)
		contents, err := os.ReadFile(dataFilepath)
		if err != nil {
			return "", err
		}

		return string(contents), nil
	}

	return "", nil
}

func getDataRaw() (data string, err error) {
	data, err = profiles.GetOptionValue(options.RequestDataRawOption)
	if err != nil {
		return "", err
	}

	return data, nil
}
