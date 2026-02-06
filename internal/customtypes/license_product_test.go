// Copyright © 2026 Ping Identity Corporation

package customtypes_test

import (
	"testing"

	"github.com/pingidentity/pingcli/internal/customtypes"
	"github.com/pingidentity/pingcli/internal/testing/testutils_koanf"
	"github.com/pingidentity/pingcli/internal/utils"
	"github.com/stretchr/testify/require"
)

func Test_LicenseProduct_Set(t *testing.T) {
	testutils_koanf.InitKoanfs(t)

	testCases := []struct {
		name          string
		cType         *customtypes.LicenseProduct
		value         string
		expectedError error
	}{
		{
			name:  "Happy path",
			cType: new(customtypes.LicenseProduct),
			value: customtypes.ENUM_LICENSE_PRODUCT_PING_ACCESS,
		},
		{
			name:  "Happy path - empty",
			cType: new(customtypes.LicenseProduct),
			value: "",
		},
		{
			name:          "Invalid value",
			cType:         new(customtypes.LicenseProduct),
			value:         "invalid",
			expectedError: customtypes.ErrUnrecognizedProduct,
		},
		{
			name:          "Nil custom type",
			cType:         nil,
			value:         customtypes.ENUM_LICENSE_PRODUCT_PING_ACCESS,
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

func Test_LicenseProduct_Type(t *testing.T) {
	testutils_koanf.InitKoanfs(t)

	testCases := []struct {
		name         string
		cType        *customtypes.LicenseProduct
		expectedType string
	}{
		{
			name:         "Happy path",
			cType:        utils.Pointer(customtypes.LicenseProduct(customtypes.ENUM_LICENSE_PRODUCT_PING_ACCESS)),
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

func Test_LicenseProduct_String(t *testing.T) {
	testutils_koanf.InitKoanfs(t)

	testCases := []struct {
		name        string
		cType       *customtypes.LicenseProduct
		expectedStr string
	}{
		{
			name:        "Happy path",
			cType:       utils.Pointer(customtypes.LicenseProduct(customtypes.ENUM_LICENSE_PRODUCT_PING_ACCESS)),
			expectedStr: customtypes.ENUM_LICENSE_PRODUCT_PING_ACCESS,
		},
		{
			name:        "Happy path - empty",
			cType:       utils.Pointer(customtypes.LicenseProduct("")),
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
