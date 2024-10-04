package common

import (
	"fmt"
	"net/http"

	"github.com/patrickcping/pingone-go-sdk-v2/management"
	"github.com/patrickcping/pingone-go-sdk-v2/mfa"
	"github.com/patrickcping/pingone-go-sdk-v2/risk"
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

	if err != nil || response.StatusCode == 404 || response.StatusCode >= 300 {
		l.Error().Err(err).Msgf("%s Request for resource '%s' was not successful. \nResponse Code: %s\nResponse Body: %s", apiFunctionName, resourceType, response.Status, response.Body)
		return fmt.Errorf("%s Request for resource '%s' was not successful. \nResponse Code: %s\nResponse Body: %s\n Error: %v", apiFunctionName, resourceType, response.Status, response.Body, err)
	}

	return nil
}

// Executes the function apiExecuteFunc for the ManagementAPIClient
// Handles err and response if Client call failed
// Returns embedded data if not nil
// Treats nil embedded data as an error
func GetManagementEmbedded(apiExecuteFunc func() (*management.EntityArray, *http.Response, error), apiFunctionName string, resourceType string) (*management.EntityArrayEmbedded, error) {
	l := logger.Get()

	entityArray, response, err := apiExecuteFunc()

	err = HandleClientResponse(response, err, apiFunctionName, resourceType)
	if err != nil {
		return nil, err
	}

	if entityArray == nil {
		l.Error().Msgf("Returned %s() entityArray is nil.", apiFunctionName)
		l.Error().Msgf("%s Response Code: %s\nResponse Body: %s", apiFunctionName, response.Status, response.Body)
		return nil, fmt.Errorf("failed to fetch %s resources via %s()", resourceType, apiFunctionName)
	}

	embedded, embeddedOk := entityArray.GetEmbeddedOk()
	if !embeddedOk {
		l.Error().Msgf("Returned %s() embedded data is nil.", apiFunctionName)
		l.Error().Msgf("%s Response Code: %s\nResponse Body: %s", apiFunctionName, response.Status, response.Body)
		return nil, fmt.Errorf("failed to fetch %s resources via %s()", resourceType, apiFunctionName)
	}

	return embedded, nil
}

// Executes the function apiExecuteFunc for the MFAAPIClient
// Handles err and response if Client call failed
// Returns embedded data if not nil
// Treats nil embedded data as an error
func GetMFAEmbedded(apiExecuteFunc func() (*mfa.EntityArray, *http.Response, error), apiFunctionName string, resourceType string) (*mfa.EntityArrayEmbedded, error) {
	l := logger.Get()

	entityArray, response, err := apiExecuteFunc()

	err = HandleClientResponse(response, err, apiFunctionName, resourceType)
	if err != nil {
		return nil, err
	}

	if entityArray == nil {
		l.Error().Msgf("Returned %s() entityArray is nil.", apiFunctionName)
		l.Error().Msgf("%s Response Code: %s\nResponse Body: %s", apiFunctionName, response.Status, response.Body)
		return nil, fmt.Errorf("failed to fetch %s resources via %s()", resourceType, apiFunctionName)
	}

	embedded, embeddedOk := entityArray.GetEmbeddedOk()
	if !embeddedOk {
		l.Error().Msgf("Returned %s() embedded data is nil.", apiFunctionName)
		l.Error().Msgf("%s Response Code: %s\nResponse Body: %s", apiFunctionName, response.Status, response.Body)
		return nil, fmt.Errorf("failed to fetch %s resources via %s()", resourceType, apiFunctionName)
	}

	return embedded, nil
}

// Executes the function apiExecuteFunc for the RiskAPIClient
// Handles err and response if Client call failed
// Returns embedded data if not nil
// Treats nil embedded data as an error
func GetProtectEmbedded(apiExecuteFunc func() (*risk.EntityArray, *http.Response, error), apiFunctionName string, resourceType string) (*risk.EntityArrayEmbedded, error) {
	l := logger.Get()

	entityArray, response, err := apiExecuteFunc()

	err = HandleClientResponse(response, err, apiFunctionName, resourceType)
	if err != nil {
		return nil, err
	}

	if entityArray == nil {
		l.Error().Msgf("Returned %s() entityArray is nil.", apiFunctionName)
		l.Error().Msgf("%s Response Code: %s\nResponse Body: %s", apiFunctionName, response.Status, response.Body)
		return nil, fmt.Errorf("failed to fetch %s resources via %s()", resourceType, apiFunctionName)
	}

	embedded, embeddedOk := entityArray.GetEmbeddedOk()
	if !embeddedOk {
		l.Error().Msgf("Returned %s() embedded data is nil.", apiFunctionName)
		l.Error().Msgf("%s Response Code: %s\nResponse Body: %s", apiFunctionName, response.Status, response.Body)
		return nil, fmt.Errorf("failed to fetch %s resources via %s()", resourceType, apiFunctionName)
	}

	return embedded, nil
}

func GenerateCommentInformation(data map[string]string) string {
	commentInformation := "\n"
	for key, value := range data {
		commentInformation += fmt.Sprintf("# %s: %s\n", key, value)
	}
	return commentInformation
}
