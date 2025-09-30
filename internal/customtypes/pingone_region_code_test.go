// Copyright © 2025 Ping Identity Corporation

package customtypes_test

import (
	"testing"

	"github.com/pingidentity/pingcli/internal/customtypes"
	"github.com/pingidentity/pingcli/internal/testing/testutils_koanf"
	"github.com/pingidentity/pingcli/internal/utils"
	"github.com/stretchr/testify/require"
)

func Test_PingOneRegion_Set(t *testing.T) {
	testutils_koanf.InitKoanfs(t)

	testCases := []struct {
		name          string
		cType         *customtypes.PingOneRegionCode
		value         string
		expectedError error
	}{
		{
			name:  "Happy path",
			cType: new(customtypes.PingOneRegionCode),
			value: customtypes.ENUM_PINGONE_REGION_CODE_AP,
		},
		{
			name:  "Happy path - empty",
			cType: new(customtypes.PingOneRegionCode),
			value: "",
		},
		{
			name:          "Invalid value",
			cType:         new(customtypes.PingOneRegionCode),
			value:         "invalid",
			expectedError: customtypes.ErrUnrecognizedPingOneRegionCode,
		},
		{
			name:          "Nil custom type",
			cType:         nil,
			value:         customtypes.ENUM_PINGONE_REGION_CODE_AP,
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

func Test_PingOneRegion_Type(t *testing.T) {
	testutils_koanf.InitKoanfs(t)

	testCases := []struct {
		name         string
		cType        *customtypes.PingOneRegionCode
		expectedType string
	}{
		{
			name:         "Happy path",
			cType:        utils.Pointer(customtypes.PingOneRegionCode(customtypes.ENUM_PINGONE_REGION_CODE_AP)),
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

func Test_PingOneRegion_String(t *testing.T) {
	testutils_koanf.InitKoanfs(t)

	testCases := []struct {
		name        string
		cType       *customtypes.PingOneRegionCode
		expectedStr string
	}{
		{
			name:        "Happy path",
			cType:       utils.Pointer(customtypes.PingOneRegionCode(customtypes.ENUM_PINGONE_REGION_CODE_CA)),
			expectedStr: customtypes.ENUM_PINGONE_REGION_CODE_CA,
		},
		{
			name:        "Happy path - empty",
			cType:       utils.Pointer(customtypes.PingOneRegionCode("")),
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
