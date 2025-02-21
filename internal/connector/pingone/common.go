package pingone

import (
	"fmt"
	"net/http"
	"reflect"

	"github.com/patrickcping/pingone-go-sdk-v2/management"
	"github.com/patrickcping/pingone-go-sdk-v2/mfa"
	"github.com/patrickcping/pingone-go-sdk-v2/risk"
	"github.com/pingidentity/pingcli/internal/connector/common"
	"github.com/pingidentity/pingcli/internal/output"
)

func CheckSingletonResource(response *http.Response, err error, apiFuncName, resourceType string) (bool, error) {
	ok, err := common.HandleClientResponse(response, err, apiFuncName, resourceType)
	if err != nil {
		return false, err
	}
	if !ok {
		return false, nil
	}

	if response.StatusCode == 204 {
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

func GetManagementAPIObjectsFromIterator[T any](iter management.EntityArrayPagedIterator, clientFuncName, extractionFuncName, resourceType string) ([]T, error) {
	apiObjects := []T{}

	for cursor, err := range iter {
		ok, err := common.HandleClientResponse(cursor.HTTPResponse, err, clientFuncName, resourceType)
		if err != nil {
			return nil, err
		}
		// A warning was given when handling the client response. Return nil apiObjects to skip export of resource
		if !ok {
			return nil, nil
		}

		nilErr := common.DataNilError(resourceType, cursor.HTTPResponse)

		if cursor.EntityArray == nil {
			return nil, nilErr
		}

		embedded, embeddedOk := cursor.EntityArray.GetEmbeddedOk()
		if !embeddedOk {
			return nil, nilErr
		}

		apiObject, err := getAPIObjectFromEmbedded[T](reflect.ValueOf(embedded), extractionFuncName, resourceType)
		if err != nil {
			output.SystemError(err.Error(), nil)
		}

		apiObjects = append(apiObjects, apiObject...)
	}

	return apiObjects, nil
}

func GetMfaAPIObjectsFromIterator[T any](iter mfa.EntityArrayPagedIterator, clientFuncName, extractionFuncName, resourceType string) ([]T, error) {
	apiObjects := []T{}

	for cursor, err := range iter {
		ok, err := common.HandleClientResponse(cursor.HTTPResponse, err, clientFuncName, resourceType)
		if err != nil {
			return nil, err
		}
		// A warning was given when handling the client response. Return nil apiObjects to skip export of resource
		if !ok {
			return nil, nil
		}

		nilErr := common.DataNilError(resourceType, cursor.HTTPResponse)

		if cursor.EntityArray == nil {
			return nil, nilErr
		}

		embedded, embeddedOk := cursor.EntityArray.GetEmbeddedOk()
		if !embeddedOk {
			return nil, nilErr
		}

		apiObject, err := getAPIObjectFromEmbedded[T](reflect.ValueOf(embedded), extractionFuncName, resourceType)
		if err != nil {
			output.SystemError(err.Error(), nil)
		}

		apiObjects = append(apiObjects, apiObject...)
	}

	return apiObjects, nil
}

func GetRiskAPIObjectsFromIterator[T any](iter risk.EntityArrayPagedIterator, clientFuncName, extractionFuncName, resourceType string) ([]T, error) {
	apiObjects := []T{}

	for cursor, err := range iter {
		ok, err := common.HandleClientResponse(cursor.HTTPResponse, err, clientFuncName, resourceType)
		if err != nil {
			return nil, err
		}
		// A warning was given when handling the client response. Return nil apiObjects to skip export of resource
		if !ok {
			return nil, nil
		}

		nilErr := common.DataNilError(resourceType, cursor.HTTPResponse)

		if cursor.EntityArray == nil {
			return nil, nilErr
		}

		embedded, embeddedOk := cursor.EntityArray.GetEmbeddedOk()
		if !embeddedOk {
			return nil, nilErr
		}

		apiObject, err := getAPIObjectFromEmbedded[T](reflect.ValueOf(embedded), extractionFuncName, resourceType)
		if err != nil {
			output.SystemError(err.Error(), nil)
		}

		apiObjects = append(apiObjects, apiObject...)
	}

	return apiObjects, nil
}

func getAPIObjectFromEmbedded[T any](embedded reflect.Value, extractionFuncName, resourceType string) ([]T, error) {
	embeddedExtractionFunc := embedded.MethodByName(extractionFuncName)
	if !embeddedExtractionFunc.IsValid() {
		return nil, fmt.Errorf("failed to find extraction function '%s' for resource '%s'", extractionFuncName, resourceType)
	}

	reflectValues := embeddedExtractionFunc.Call(nil)
	if len(reflectValues) == 0 {
		return nil, fmt.Errorf("failed to get reflect value from embedded. embedded is empty")
	}

	rInterface := reflectValues[0].Interface()
	if rInterface == nil {
		return []T{}, nil
	}

	apiObject, apiObjectOk := rInterface.([]T)
	if !apiObjectOk {
		return nil, fmt.Errorf("failed to cast reflect value to %s", resourceType)
	}

	return apiObject, nil
}
