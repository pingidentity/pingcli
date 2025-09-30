// Copyright © 2025 Ping Identity Corporation

package customtypes_test

import (
	"testing"

	"github.com/pingidentity/pingcli/internal/customtypes"
	"github.com/pingidentity/pingcli/internal/testing/testutils_koanf"
	"github.com/pingidentity/pingcli/internal/utils"
	"github.com/stretchr/testify/require"
)

func Test_RequestServices_Set(t *testing.T) {
	testutils_koanf.InitKoanfs(t)

	testCases := []struct {
		name          string
		cType         *customtypes.RequestService
		value         string
		expectedError error
	}{
		{
			name:  "Happy path",
			cType: new(customtypes.RequestService),
			value: customtypes.ENUM_REQUEST_SERVICE_PINGONE,
		},
		{
			name:  "Happy path - empty",
			cType: new(customtypes.RequestService),
			value: "",
		},
		{
			name:          "Invalid value",
			cType:         new(customtypes.RequestService),
			value:         "invalid",
			expectedError: customtypes.ErrUnrecognizedService,
		},
		{
			name:          "Nil custom type",
			cType:         nil,
			value:         customtypes.ENUM_REQUEST_SERVICE_PINGONE,
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

func Test_RequestServices_Type(t *testing.T) {
	testutils_koanf.InitKoanfs(t)

	testCases := []struct {
		name         string
		cType        *customtypes.RequestService
		expectedType string
	}{
		{
			name:         "Happy path",
			cType:        utils.Pointer(customtypes.RequestService(customtypes.ENUM_REQUEST_SERVICE_PINGONE)),
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

func Test_RequestServices_String(t *testing.T) {
	testutils_koanf.InitKoanfs(t)

	testCases := []struct {
		name        string
		cType       *customtypes.RequestService
		expectedStr string
	}{
		{
			name:        "Happy path",
			cType:       utils.Pointer(customtypes.RequestService(customtypes.ENUM_REQUEST_SERVICE_PINGONE)),
			expectedStr: customtypes.ENUM_REQUEST_SERVICE_PINGONE,
		},
		{
			name:        "Happy path - empty",
			cType:       utils.Pointer(customtypes.RequestService("")),
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
