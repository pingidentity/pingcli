package common

import (
	"fmt"
	"maps"
	"net/http"
	"slices"

	"github.com/pingidentity/pingcli/internal/logger"
)

const (
	SINGLETON_ID_COMMENT_DATA = "This resource is a singleton, so the value of 'ID' in the import block does not matter - it is just a placeholder and required by terraform."
)

func HandleClientResponse(response *http.Response, err error, apiFunctionName string, resourceType string) error {
	l := logger.Get()

	if response == nil {
		l.Error().Err(err).Msgf("%s Request for resource '%s' was not successful. Response is nil.", apiFunctionName, resourceType)
		return fmt.Errorf("%s Request for resource '%s' was not successful. Response is nil. Error: %v", apiFunctionName, resourceType, err)
	}

	defer response.Body.Close()

	if err != nil || response.StatusCode >= 300 || response.StatusCode < 200 {
		l.Error().Err(err).Msgf("%s Request for resource '%s' was not successful. \nResponse Code: %s\nResponse Body: %s", apiFunctionName, resourceType, response.Status, response.Body)
		return fmt.Errorf("%s Request for resource '%s' was not successful. \nResponse Code: %s\nResponse Body: %s\n Error: %v", apiFunctionName, resourceType, response.Status, response.Body, err)
	}

	return nil
}

func DataNilError(resourceType string, response *http.Response) error {
	return fmt.Errorf("failed to export resource '%s'.\n"+
		"PingOne API request for resource '%s' was not successful. response data is nil.\n"+
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
