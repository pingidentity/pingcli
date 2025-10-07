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

func runInternalPingOneRequest(uri string) (err error) {
	pingOneAPIClient, err := auth_internal.GetPingOneClient()
	if err != nil {
		return err
	}

	pingOneAPIClientConfig := pingOneAPIClient.GetConfig()
	if pingOneAPIClientConfig == nil {
		return fmt.Errorf("PingOne client configuration is nil")
	}

	tld := pingOneAPIClientConfig.Service.GetTopLevelDomain()
	if tld == "" {
		return fmt.Errorf("failed to determine PingOne environment top level domain")
	}

	failOption, err := profiles.GetOptionValue(options.RequestFailOption)
	if err != nil {
		return err
	}

	apiURL := fmt.Sprintf("https://api.pingone.%s/v1/%s", tld, uri)

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

	// Get access token once
	ctx := context.Background()
	accessToken, err := pingOneAPIClientConfig.Service.GetAccessToken(ctx)
	if err != nil {
		return fmt.Errorf("failed to get access token: %w", err)
	}

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
