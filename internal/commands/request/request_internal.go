// Copyright © 2025 Ping Identity Corporation

package request_internal

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	auth_internal "github.com/pingidentity/pingcli/internal/auth"
	"github.com/pingidentity/pingcli/internal/configuration/options"
	"github.com/pingidentity/pingcli/internal/customtypes"
	"github.com/pingidentity/pingcli/internal/output"
	"github.com/pingidentity/pingcli/internal/profiles"
)

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
	pingOneClient, err := auth_internal.GetPingOneClient()
	if err != nil {
		return err
	}

	failOption, err := profiles.GetOptionValue(options.RequestFailOption)
	if err != nil {
		return err
	}

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

	apiURL, err := pingOneClient.GetConfig().Servers.)
	if err != nil {
		return err
	}

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
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", *pingOneClient.GetConfig().Service.Auth.AccessToken))

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
