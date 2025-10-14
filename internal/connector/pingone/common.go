// Copyright © 2025 Ping Identity Corporation

package pingone

import (
	"fmt"
	"reflect"

	"github.com/patrickcping/pingone-go-sdk-v2/authorize"
	"github.com/patrickcping/pingone-go-sdk-v2/management"
	"github.com/patrickcping/pingone-go-sdk-v2/mfa"
	"github.com/patrickcping/pingone-go-sdk-v2/risk"
	"github.com/pingidentity/pingcli/internal/connector/common"
	"github.com/pingidentity/pingcli/internal/errs"
	"github.com/pingidentity/pingcli/internal/output"
)

var (
	pingoneConnectorCommonErrorPrefix = "pingone connector common utils error"
)

func GetAuthorizeAPIObjectsFromIterator[T any](iter authorize.EntityArrayPagedIterator, clientFuncName, extractionFuncName, resourceType string) ([]T, error) {
	apiObjects := []T{}

	for cursor, err := range iter {
		ok, err := common.HandleClientResponse(cursor.HTTPResponse, err, clientFuncName, resourceType)
		if err != nil {
			return nil, &errs.PingCLIError{Prefix: pingoneConnectorCommonErrorPrefix, Err: err}
		}
		// A warning was given when handling the client response. Return nil apiObjects to skip export of resource
		if !ok {
			return nil, nil
		}

		nilErr := &errs.PingCLIError{Prefix: pingoneConnectorCommonErrorPrefix, Err: common.DataNilError(resourceType, cursor.HTTPResponse)}

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

func GetManagementAPIObjectsFromIterator[T any](iter management.EntityArrayPagedIterator, clientFuncName, extractionFuncName, resourceType string) ([]T, error) {
	apiObjects := []T{}

	for cursor, err := range iter {
		ok, err := common.HandleClientResponse(cursor.HTTPResponse, err, clientFuncName, resourceType)
		if err != nil {
			return nil, &errs.PingCLIError{Prefix: pingoneConnectorCommonErrorPrefix, Err: err}
		}
		// A warning was given when handling the client response. Return nil apiObjects to skip export of resource
		if !ok {
			return nil, nil
		}

		nilErr := &errs.PingCLIError{Prefix: pingoneConnectorCommonErrorPrefix, Err: common.DataNilError(resourceType, cursor.HTTPResponse)}

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
			return nil, &errs.PingCLIError{Prefix: pingoneConnectorCommonErrorPrefix, Err: err}
		}
		// A warning was given when handling the client response. Return nil apiObjects to skip export of resource
		if !ok {
			return nil, nil
		}

		nilErr := &errs.PingCLIError{Prefix: pingoneConnectorCommonErrorPrefix, Err: common.DataNilError(resourceType, cursor.HTTPResponse)}

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
			return nil, &errs.PingCLIError{Prefix: pingoneConnectorCommonErrorPrefix, Err: err}
		}
		// A warning was given when handling the client response. Return nil apiObjects to skip export of resource
		if !ok {
			return nil, nil
		}

		nilErr := &errs.PingCLIError{Prefix: pingoneConnectorCommonErrorPrefix, Err: common.DataNilError(resourceType, cursor.HTTPResponse)}

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
		return nil, &errs.PingCLIError{
			Prefix: pingoneConnectorCommonErrorPrefix,
			Err:    fmt.Errorf("%w. Function %q. Resource %q", ErrUnknownExtractionFunction, extractionFuncName, resourceType),
		}
	}

	reflectValues := embeddedExtractionFunc.Call(nil)
	if len(reflectValues) == 0 {
		return nil, &errs.PingCLIError{Prefix: pingoneConnectorCommonErrorPrefix, Err: ErrEmbeddedEmpty}
	}

	rInterface := reflectValues[0].Interface()
	if rInterface == nil {
		return []T{}, nil
	}

	apiObject, apiObjectOk := rInterface.([]T)
	if !apiObjectOk {
		return nil, &errs.PingCLIError{Prefix: pingoneConnectorCommonErrorPrefix, Err: fmt.Errorf("%w. Resource Type %q", ErrCastReflectValue, resourceType)}
	}

	return apiObject, nil
}
