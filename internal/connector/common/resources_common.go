// Copyright © 2025 Ping Identity Corporation

package common

import (
	"errors"
	"fmt"
	"maps"
	"net/http"
	"slices"

	"github.com/pingidentity/pingcli/internal/errs"
	"github.com/pingidentity/pingcli/internal/output"
)

const (
	SINGLETON_ID_COMMENT_DATA = "This resource is a singleton, so the value of 'ID' in the import block does not matter - it is just a placeholder and required by terraform."
	resourceUtilsErrorPrefix  = "connector resource utils error"
)

func CheckSingletonResource(response *http.Response, err error, apiFuncName, resourceType string) (bool, error) {
	ok, err := HandleClientResponse(response, err, apiFuncName, resourceType)
	if err != nil {
		return false, &errs.PingCLIError{Prefix: resourceUtilsErrorPrefix, Err: err}
	}
	if !ok {
		return false, nil
	}

	if response.StatusCode == http.StatusNoContent {
		output.Warn("API client 204 No Content response.", map[string]interface{}{
			"API Function Name": apiFuncName,
			"Resource Type":     resourceType,
			"Response Code":     response.Status,
			"Response Body":     response.Body,
		})

		return false, nil
	}

	return true, nil
}

func HandleClientResponse(response *http.Response, err error, apiFunctionName string, resourceType string) (b bool, rErr error) {
	if err != nil {
		// Only warn the user on client error and skip export of resource
		output.Warn("API client error.", map[string]interface{}{
			"API Function Name": apiFunctionName,
			"Resource Type":     resourceType,
			"Client Error":      err,
		})

		return false, nil
	}

	if response == nil {
		return false, &errs.PingCLIError{Prefix: resourceUtilsErrorPrefix, Err: fmt.Errorf("%w: %q - %q. Response is nil", ErrResourceRequestFailed, apiFunctionName, resourceType)}
	}

	defer func() {
		cErr := response.Body.Close()
		if cErr != nil {
			rErr = errors.Join(rErr, cErr)
			rErr = &errs.PingCLIError{Prefix: resourceUtilsErrorPrefix, Err: rErr}
		}
	}()

	// When the client returns forbidden, warn user and skip export of resource
	if response.StatusCode == http.StatusForbidden {
		output.Warn("API client 403 forbidden response.", map[string]interface{}{
			"API Function Name": apiFunctionName,
			"Resource Type":     resourceType,
			"Response Code":     response.StatusCode,
			"Response Body":     response.Body,
		})

		return false, nil
	}

	// Error on any other non-200 response
	if response.StatusCode >= 300 || response.StatusCode < 200 {
		return false, &errs.PingCLIError{
			Prefix: resourceUtilsErrorPrefix,
			Err: fmt.Errorf(
				"%w: %q - %q. Response Code: %s, Response Body: %s",
				ErrResourceRequestFailed,
				apiFunctionName,
				resourceType,
				response.Status,
				response.Body),
		}
	}

	return true, nil
}

func DataNilError(resourceType string, response *http.Response) error {
	return &errs.PingCLIError{
		Prefix: resourceUtilsErrorPrefix,
		Err: fmt.Errorf(
			"%w: %q. API Client request for resource '%s' was not successful. response data is nil.\n"+
				"response code: %s\n"+
				"response body: %s",
			ErrExportResources,
			resourceType,
			resourceType,
			response.Status,
			response.Body,
		),
	}
}

func GenerateCommentInformation(data map[string]string) string {
	// Get a sorted slice of the keys
	keys := slices.Sorted(maps.Keys(data))

	commentInformation := "\n"
	for _, key := range keys {
		commentInformation += fmt.Sprintf("# %s: %s\n", key, data[key])
	}

	return commentInformation
}
