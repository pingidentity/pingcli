// Copyright © 2025 Ping Identity Corporation

package customtypes_test

import (
	"slices"
	"testing"

	"github.com/pingidentity/pingcli/internal/customtypes"
	"github.com/pingidentity/pingcli/internal/testing/testutils_koanf"
	"github.com/pingidentity/pingcli/internal/utils"
	"github.com/stretchr/testify/require"
)

func Test_HTTPMethod_Set(t *testing.T) {
	testutils_koanf.InitKoanfs(t)

	testCases := []struct {
		name          string
		cType         *customtypes.HTTPMethod
		value         string
		expectedError error
	}{
		{
			name:  "Happy path - GET",
			cType: new(customtypes.HTTPMethod),
			value: customtypes.ENUM_HTTP_METHOD_GET,
		},
		{
			name:  "Happy path - empty",
			cType: new(customtypes.HTTPMethod),
			value: "",
		},
		{
			name:          "Invalid value",
			cType:         new(customtypes.HTTPMethod),
			value:         "invalid",
			expectedError: customtypes.ErrUnrecognizedMethod,
		},
		{
			name:          "Nil custom type",
			cType:         nil,
			value:         customtypes.ENUM_HTTP_METHOD_GET,
			expectedError: customtypes.ErrCustomTypeNil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			testutils_koanf.InitKoanfs(t)

			err := tc.cType.Set(tc.value)

			if tc.expectedError != nil {
				require.Error(t, err)
				require.ErrorIs(t, err, tc.expectedError)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func Test_HTTPMethod_Type(t *testing.T) {
	testutils_koanf.InitKoanfs(t)

	testCases := []struct {
		name         string
		cType        *customtypes.HTTPMethod
		expectedType string
	}{
		{
			name:         "Happy path",
			cType:        utils.Pointer(customtypes.HTTPMethod(customtypes.ENUM_HTTP_METHOD_GET)),
			expectedType: "string",
		},
		{
			name:         "Nil custom type",
			cType:        nil,
			expectedType: "string",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			testutils_koanf.InitKoanfs(t)

			actualType := tc.cType.Type()

			require.Equal(t, tc.expectedType, actualType)
		})
	}
}

func Test_HTTPMethod_String(t *testing.T) {
	testutils_koanf.InitKoanfs(t)

	testCases := []struct {
		name        string
		cType       *customtypes.HTTPMethod
		expectedStr string
	}{
		{
			name:        "Happy path",
			cType:       utils.Pointer(customtypes.HTTPMethod(customtypes.ENUM_HTTP_METHOD_GET)),
			expectedStr: customtypes.ENUM_HTTP_METHOD_GET,
		},
		{
			name:        "Happy path - empty",
			cType:       utils.Pointer(customtypes.HTTPMethod("")),
			expectedStr: "",
		},
		{
			name:        "Nil custom type",
			cType:       nil,
			expectedStr: "",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			testutils_koanf.InitKoanfs(t)

			actualStr := tc.cType.String()

			require.Equal(t, tc.expectedStr, actualStr)
		})
	}
}

func Test_HTTPMethodValidValues(t *testing.T) {
	testutils_koanf.InitKoanfs(t)

	expectedValues := []string{
		customtypes.ENUM_HTTP_METHOD_DELETE,
		customtypes.ENUM_HTTP_METHOD_GET,
		customtypes.ENUM_HTTP_METHOD_PATCH,
		customtypes.ENUM_HTTP_METHOD_POST,
		customtypes.ENUM_HTTP_METHOD_PUT,
	}

	slices.Sort(expectedValues)

	actualValues := customtypes.HTTPMethodValidValues()

	require.Equal(t, expectedValues, actualValues)
}
