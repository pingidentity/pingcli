package common

import (
	"fmt"
	"maps"
	"net/http"
	"reflect"
	"slices"

	"github.com/patrickcping/pingone-go-sdk-v2/management"
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
		return false, fmt.Errorf("%s Request for resource '%s' was not successful. Response is nil.", apiFunctionName, resourceType)
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

func dataNilError(resourceType string, response *http.Response) error {
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

func CheckSingletonResource(response *http.Response, err error, apiFuncName, resourceType string) (bool, error) {
	ok, err := HandleClientResponse(response, err, "ReadBrandingSettings", resourceType)
	if err != nil {
		return false, err
	}
	if !ok {
		return false, nil
	}

	if response.StatusCode == 204 {
		output.Warn("API client 204 No Content response.", map[string]interface{}{
			"api_function": apiFuncName,
			"resource":     resourceType,
		})
		return false, nil
	}

	return true, nil
}

func GetManagementAPIObjectsFromIterator[T any](iter management.EntityArrayPagedIterator, clientFuncName, extractionFuncName, resourceType string) ([]T, error) {
	apiObjects := []T{}

	for cursor, err := range iter {
		ok, err := HandleClientResponse(cursor.HTTPResponse, err, clientFuncName, resourceType)
		if err != nil {
			return nil, err
		}
		// A warning was given when handling the client response. Return nil embeddeds to skip export of resource
		if !ok {
			return nil, nil
		}

		nilErr := dataNilError(resourceType, cursor.HTTPResponse)

		if cursor.EntityArray == nil {
			return nil, nilErr
		}

		embedded, embeddedOk := cursor.EntityArray.GetEmbeddedOk()
		if !embeddedOk {
			return nil, nilErr
		}

		reflectValues := reflect.ValueOf(embedded).MethodByName(extractionFuncName).Call(nil)
		for _, rValue := range reflectValues {
			apiObject, apiObjectOk := rValue.Interface().(T)
			if !apiObjectOk {
				output.SystemError(fmt.Sprintf("Failed to cast reflect value to %s", resourceType), nil)
			}

			apiObjects = append(apiObjects, apiObject)
		}
	}

	return apiObjects, nil
}
