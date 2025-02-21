package common

import (
	"fmt"
	"maps"
	"net/http"
	"slices"

	"github.com/pingidentity/pingcli/internal/output"
)

const (
	SINGLETON_ID_COMMENT_DATA = "This resource is a singleton, so the value of 'ID' in the import block does not matter - it is just a placeholder and required by terraform."
)

func HandleClientResponse(response *http.Response, err error, apiFunctionName string, resourceType string) (bool, error) {
	if err != nil {
		// Only warn the user on client error and skip export of resource
		output.Warn("API client error.", map[string]interface{}{
			"api_function": apiFunctionName,
			"error":        err,
			"resource":     resourceType,
		})

		return false, nil
	}

	if response == nil {
		return false, fmt.Errorf("%s Request for resource '%s' was not successful. Response is nil", apiFunctionName, resourceType)
	}
	defer response.Body.Close()

	// When the client returns forbidden, warn user and skip export of resource
	if response.StatusCode == 403 {
		output.Warn("API client 403 forbidden response.", map[string]interface{}{
			"api_function":  apiFunctionName,
			"resource":      resourceType,
			"response_code": response.StatusCode,
			"response_body": response.Body,
		})

		return false, nil
	}

	// Error on any other non-200 response
	if response.StatusCode >= 300 || response.StatusCode < 200 {
		return false, fmt.Errorf("%s Request for resource '%s' was not successful. \nResponse Code: %s\nResponse Body: %s", apiFunctionName, resourceType, response.Status, response.Body)
	}

	return true, nil
}

func DataNilError(resourceType string, response *http.Response) error {
	return fmt.Errorf("failed to export resource '%s'.\n"+
		"API Client request for resource '%s' was not successful. response data is nil.\n"+
		"response code: %s\n"+
		"response body: %s",
		resourceType, resourceType, response.Status, response.Body)
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
