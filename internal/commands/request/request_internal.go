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
	"slices"
	"strconv"
	"strings"
	"time"

	"github.com/pingidentity/pingcli/internal/configuration/options"
	"github.com/pingidentity/pingcli/internal/customtypes"
	"github.com/pingidentity/pingcli/internal/errs"
	"github.com/pingidentity/pingcli/internal/output"
	"github.com/pingidentity/pingcli/internal/profiles"
)

var (
	requestErrorPrefix = "failed to send custom request"
)

type PingOneAuthResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int64  `json:"expires_in"`
}

func RunInternalRequest(uri string) (err error) {
	service, err := profiles.GetOptionValue(options.RequestServiceOption)
	if err != nil {
		return &errs.PingCLIError{Prefix: requestErrorPrefix, Err: err}
	}

	if service == "" {
		return &errs.PingCLIError{Prefix: requestErrorPrefix, Err: ErrServiceEmpty}
	}

	switch service {
	case customtypes.ENUM_REQUEST_SERVICE_PINGONE:
		err = runInternalPingOneRequest(uri)
		if err != nil {
			return &errs.PingCLIError{Prefix: requestErrorPrefix, Err: err}
		}
	default:
		return &errs.PingCLIError{Prefix: requestErrorPrefix, Err: fmt.Errorf("%w: '%s'", ErrUnrecognizedService, service)}
	}

	return nil
}

func runInternalPingOneRequest(uri string) (err error) {
	accessToken, err := pingoneAccessToken()
	if err != nil {
		return &errs.PingCLIError{Prefix: requestErrorPrefix, Err: err}
	}

	topLevelDomain, err := getTopLevelDomain()
	if err != nil {
		return &errs.PingCLIError{Prefix: requestErrorPrefix, Err: err}
	}

	failOption, err := profiles.GetOptionValue(options.RequestFailOption)
	if err != nil {
		return &errs.PingCLIError{Prefix: requestErrorPrefix, Err: err}
	}

	apiURL := fmt.Sprintf("https://api.pingone.%s/v1/%s", topLevelDomain, uri)

	httpMethod, err := profiles.GetOptionValue(options.RequestHTTPMethodOption)
	if err != nil {
		return &errs.PingCLIError{Prefix: requestErrorPrefix, Err: err}
	}

	if httpMethod == "" {
		return &errs.PingCLIError{Prefix: requestErrorPrefix, Err: ErrHttpMethodEmpty}
	}

	if !slices.Contains(customtypes.HTTPMethodValidValues(), httpMethod) {
		return &errs.PingCLIError{Prefix: requestErrorPrefix, Err: fmt.Errorf("%w: '%s'", ErrUnrecognizedHttpMethod, httpMethod)}
	}

	data, err := getDataRaw()
	if err != nil {
		return &errs.PingCLIError{Prefix: requestErrorPrefix, Err: err}
	}

	if data == "" {
		data, err = getDataFile()
		if err != nil {
			return &errs.PingCLIError{Prefix: requestErrorPrefix, Err: err}
		}
	}

	payload := strings.NewReader(data)

	client := &http.Client{}
	req, err := http.NewRequestWithContext(context.Background(), httpMethod, apiURL, payload)
	if err != nil {
		return &errs.PingCLIError{Prefix: requestErrorPrefix, Err: err}
	}

	headers, err := profiles.GetOptionValue(options.RequestHeaderOption)
	if err != nil {
		return &errs.PingCLIError{Prefix: requestErrorPrefix, Err: err}
	}

	requestHeaders := new(customtypes.HeaderSlice)
	err = requestHeaders.Set(headers)
	if err != nil {
		return &errs.PingCLIError{Prefix: requestErrorPrefix, Err: err}
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
		return &errs.PingCLIError{Prefix: requestErrorPrefix, Err: err}
	}

	defer func() {
		cErr := res.Body.Close()
		if cErr != nil {
			err = errors.Join(err, cErr)
			err = &errs.PingCLIError{Prefix: requestErrorPrefix, Err: err}
		}
	}()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return &errs.PingCLIError{Prefix: requestErrorPrefix, Err: err}
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
		return topLevelDomain, &errs.PingCLIError{Prefix: requestErrorPrefix, Err: err}
	}

	if pingoneRegionCode == "" {
		return topLevelDomain, &errs.PingCLIError{Prefix: requestErrorPrefix, Err: ErrPingOneRegionCodeEmpty}
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
		return topLevelDomain, &errs.PingCLIError{Prefix: requestErrorPrefix, Err: fmt.Errorf("%w: '%s'", ErrUnrecognizedPingOneRegionCode, pingoneRegionCode)}
	}

	return topLevelDomain, nil
}

func pingoneAccessToken() (accessToken string, err error) {
	// Check if existing access token is available
	accessToken, err = profiles.GetOptionValue(options.RequestAccessTokenOption)
	if err != nil {
		return accessToken, &errs.PingCLIError{Prefix: requestErrorPrefix, Err: err}
	}

	if accessToken != "" {
		accessTokenExpiry, err := profiles.GetOptionValue(options.RequestAccessTokenExpiryOption)
		if err != nil {
			return accessToken, &errs.PingCLIError{Prefix: requestErrorPrefix, Err: err}
		}

		if accessTokenExpiry == "" {
			accessTokenExpiry = "0"
		}

		// convert expiry string to int
		tokenExpiryInt, err := strconv.ParseInt(accessTokenExpiry, 10, 64)
		if err != nil {
			return accessToken, &errs.PingCLIError{Prefix: requestErrorPrefix, Err: err}
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
		return "", &errs.PingCLIError{Prefix: requestErrorPrefix, Err: err}
	}

	workerEnvId, err := profiles.GetOptionValue(options.PingOneAuthenticationWorkerEnvironmentIDOption)
	if err != nil {
		return "", &errs.PingCLIError{Prefix: requestErrorPrefix, Err: err}
	}

	if workerEnvId == "" {
		return "", &errs.PingCLIError{Prefix: requestErrorPrefix, Err: ErrPingOneWorkerEnvIDEmpty}
	}

	authURL := fmt.Sprintf("https://auth.pingone.%s/%s/as/token", topLevelDomain, workerEnvId)

	clientId, err := profiles.GetOptionValue(options.PingOneAuthenticationWorkerClientIDOption)
	if err != nil {
		return "", &errs.PingCLIError{Prefix: requestErrorPrefix, Err: err}
	}
	clientSecret, err := profiles.GetOptionValue(options.PingOneAuthenticationWorkerClientSecretOption)
	if err != nil {
		return "", &errs.PingCLIError{Prefix: requestErrorPrefix, Err: err}
	}

	if clientId == "" || clientSecret == "" {
		return "", &errs.PingCLIError{Prefix: requestErrorPrefix, Err: ErrPingOneClientIDAndSecretEmpty}
	}

	basicAuthBase64 := base64.StdEncoding.EncodeToString([]byte(clientId + ":" + clientSecret))

	payload := strings.NewReader("grant_type=client_credentials")

	client := &http.Client{}
	req, err := http.NewRequestWithContext(context.Background(), customtypes.ENUM_HTTP_METHOD_POST, authURL, payload)
	if err != nil {
		return "", &errs.PingCLIError{Prefix: requestErrorPrefix, Err: err}
	}

	req.Header.Add("Authorization", fmt.Sprintf("Basic %s", basicAuthBase64))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	res, err := client.Do(req)
	if err != nil {
		return "", &errs.PingCLIError{Prefix: requestErrorPrefix, Err: err}
	}

	defer func() {
		cErr := res.Body.Close()
		if cErr != nil {
			err = errors.Join(err, cErr)
			err = &errs.PingCLIError{Prefix: requestErrorPrefix, Err: err}
		}
	}()

	responseBodyBytes, err := io.ReadAll(res.Body)
	if err != nil {
		return "", &errs.PingCLIError{Prefix: requestErrorPrefix, Err: err}
	}

	if res.StatusCode < 200 || res.StatusCode >= 300 {
		return "", &errs.PingCLIError{
			Prefix: requestErrorPrefix,
			Err:    fmt.Errorf("%w: Response Status %s: Response Body %s", ErrPingOneAuthenticate, res.Status, string(responseBodyBytes)),
		}
	}

	pingoneAuthResponse := new(PingOneAuthResponse)
	err = json.Unmarshal(responseBodyBytes, pingoneAuthResponse)
	if err != nil {
		return "", &errs.PingCLIError{Prefix: requestErrorPrefix, Err: err}
	}

	currentTime := time.Now().Unix()
	tokenExpiry := currentTime + pingoneAuthResponse.ExpiresIn

	// Store access token and expiry
	pName, err := profiles.GetOptionValue(options.RootProfileOption)
	if err != nil {
		return "", &errs.PingCLIError{Prefix: requestErrorPrefix, Err: err}
	}

	if pName == "" {
		pName, err = profiles.GetOptionValue(options.RootActiveProfileOption)
		if err != nil {
			return "", &errs.PingCLIError{Prefix: requestErrorPrefix, Err: err}
		}
	}

	koanfConfig, err := profiles.GetKoanfConfig()
	if err != nil {
		return "", &errs.PingCLIError{Prefix: requestErrorPrefix, Err: err}
	}

	subKoanf, err := koanfConfig.GetProfileKoanf(pName)
	if err != nil {
		return "", &errs.PingCLIError{Prefix: requestErrorPrefix, Err: err}
	}

	err = subKoanf.Set(options.RequestAccessTokenOption.KoanfKey, pingoneAuthResponse.AccessToken)
	if err != nil {
		return "", &errs.PingCLIError{Prefix: requestErrorPrefix, Err: err}
	}
	err = subKoanf.Set(options.RequestAccessTokenExpiryOption.KoanfKey, tokenExpiry)
	if err != nil {
		return "", &errs.PingCLIError{Prefix: requestErrorPrefix, Err: err}
	}
	err = koanfConfig.SaveProfile(pName, subKoanf)
	if err != nil {
		return "", &errs.PingCLIError{Prefix: requestErrorPrefix, Err: err}
	}

	return pingoneAuthResponse.AccessToken, nil
}

func getDataFile() (data string, err error) {
	dataFilepath, err := profiles.GetOptionValue(options.RequestDataOption)
	if err != nil {
		return "", &errs.PingCLIError{Prefix: requestErrorPrefix, Err: err}
	}

	if dataFilepath != "" {
		dataFilepath = filepath.Clean(dataFilepath)
		contents, err := os.ReadFile(dataFilepath)
		if err != nil {
			return "", &errs.PingCLIError{Prefix: requestErrorPrefix, Err: err}
		}

		return string(contents), nil
	}

	return "", nil
}

func getDataRaw() (data string, err error) {
	data, err = profiles.GetOptionValue(options.RequestDataRawOption)
	if err != nil {
		return "", &errs.PingCLIError{Prefix: requestErrorPrefix, Err: err}
	}

	return data, nil
}
