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
	"strings"

	auth_internal "github.com/pingidentity/pingcli/internal/auth"
	"github.com/pingidentity/pingcli/internal/configuration/options"
	"github.com/pingidentity/pingcli/internal/customtypes"
	"github.com/pingidentity/pingcli/internal/output"
	"github.com/pingidentity/pingcli/internal/profiles"
)

// GetAPIURLForRegion builds the correct API URL based on region configuration
func GetAPIURLForRegion(uri string) (string, error) {
	regionCode, err := profiles.GetOptionValue(options.PingOneRegionCodeOption)
	if err != nil {
		return "", fmt.Errorf("failed to get region code: %w", err)
	}

	var tld string
	switch regionCode {
	case customtypes.ENUM_PINGONE_REGION_CODE_AP:
		tld = "asia"
	case customtypes.ENUM_PINGONE_REGION_CODE_AU:
		tld = "com.au"
	case customtypes.ENUM_PINGONE_REGION_CODE_CA:
		tld = "ca"
	case customtypes.ENUM_PINGONE_REGION_CODE_EU:
		tld = "eu"
	case customtypes.ENUM_PINGONE_REGION_CODE_NA:
		tld = "com"
	case customtypes.ENUM_PINGONE_REGION_CODE_SG:
		tld = "asia"
	default:
		tld = "com" // default to NA
	}

	return fmt.Sprintf("https://api.pingone.%s/v1/%s", tld, uri), nil
}

func runInternalPingOneRequest(uri string) (err error) {
	var accessToken string
	var ctx = context.Background()

	// Use the unified authentication system with OAuth2 token source
	tokenSource, err := auth_internal.GetValidTokenSource(ctx)
	if err != nil {
		return fmt.Errorf("failed to get valid token source: %w", err)
	}

	// Get access token from the token source (handles caching and refresh)
	token, err := tokenSource.Token()
	if err != nil {
		return fmt.Errorf("failed to get access token: %w", err)
	}

	accessToken = token.AccessToken

	// Build API URL using proper region configuration
	apiURL, err := GetAPIURLForRegion(uri)
	if err != nil {
		return fmt.Errorf("failed to build API URL: %w", err)
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
		return ErrHttpMethodEmpty
	}

	data, err := GetDataRaw()
	if err != nil {
		return err
	}

	if data == "" {
		data, err = GetDataFile()
		if err != nil {
			return err
		}
	}

	payload := strings.NewReader(data)

	// Create a simple HTTP client (not OAuth2-managed)
	client := &http.Client{}

	req, err := http.NewRequestWithContext(ctx, httpMethod, apiURL, payload)
	if err != nil {
		return err
	}

	// Manually add Authorization header like curl command
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", accessToken))

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
