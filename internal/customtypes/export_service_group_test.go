// Copyright © 2025 Ping Identity Corporation

package customtypes_test

import (
	"testing"

	"github.com/pingidentity/pingcli/internal/customtypes"
	"github.com/pingidentity/pingcli/internal/testing/testutils_koanf"
	"github.com/pingidentity/pingcli/internal/utils"
	"github.com/stretchr/testify/require"
)

func Test_ExportServiceGroup_Set(t *testing.T) {
	testutils_koanf.InitKoanfs(t)

	testCases := []struct {
		name          string
		cType         *customtypes.ExportServiceGroup
		value         string
		expectedError error
	}{
		{
			name:  "Happy path",
			cType: new(customtypes.ExportServiceGroup),
			value: customtypes.ENUM_EXPORT_SERVICE_GROUP_PINGONE,
		},
		{
			name:          "Invalid value",
			cType:         new(customtypes.ExportServiceGroup),
			value:         "invalid",
			expectedError: customtypes.ErrUnrecognisedServiceGroup,
		},
		{
			name:  "Happy path - empty",
			cType: new(customtypes.ExportServiceGroup),
			value: "",
		},
		{
			name:          "Nil custom type",
			cType:         nil,
			value:         customtypes.ENUM_EXPORT_SERVICE_GROUP_PINGONE,
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

func Test_ExportServiceGroup_Type(t *testing.T) {
	testutils_koanf.InitKoanfs(t)

	testCases := []struct {
		name         string
		cType        *customtypes.ExportServiceGroup
		expectedType string
	}{
		{
			name:         "Happy path",
			cType:        utils.Pointer(customtypes.ExportServiceGroup(customtypes.ENUM_EXPORT_SERVICE_GROUP_PINGONE)),
			expectedType: "string",
		},
		{
			name:         "Happy path - empty",
			cType:        utils.Pointer(customtypes.ExportServiceGroup("")),
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

func Test_ExportServiceGroup_String(t *testing.T) {
	testutils_koanf.InitKoanfs(t)

	testCases := []struct {
		name        string
		cType       *customtypes.ExportServiceGroup
		expectedStr string
	}{
		{
			name:        "Happy path",
			cType:       utils.Pointer(customtypes.ExportServiceGroup(customtypes.ENUM_EXPORT_SERVICE_GROUP_PINGONE)),
			expectedStr: customtypes.ENUM_EXPORT_SERVICE_GROUP_PINGONE,
		},
		{
			name:        "Happy path - empty",
			cType:       utils.Pointer(customtypes.ExportServiceGroup("")),
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

func Test_ExportServiceGroupValidValues(t *testing.T) {
	expectedServiceGroups := []string{
		customtypes.ENUM_EXPORT_SERVICE_GROUP_PINGONE,
	}

	actualServiceGroupValidValues := customtypes.ExportServiceGroupValidValues()
	require.Equal(t, actualServiceGroupValidValues, expectedServiceGroups)
}

func Test_ExportServiceGroup_GetServicesInGroup(t *testing.T) {
	testutils_koanf.InitKoanfs(t)

	testCases := []struct {
		name         string
		cType        *customtypes.ExportServiceGroup
		expectedStrs []string
	}{
		{
			name:  "Happy path - pingone",
			cType: utils.Pointer(customtypes.ExportServiceGroup(customtypes.ENUM_EXPORT_SERVICE_GROUP_PINGONE)),
			expectedStrs: []string{
				customtypes.ENUM_EXPORT_SERVICE_PINGONE_PLATFORM,
				customtypes.ENUM_EXPORT_SERVICE_PINGONE_AUTHORIZE,
				customtypes.ENUM_EXPORT_SERVICE_PINGONE_SSO,
				customtypes.ENUM_EXPORT_SERVICE_PINGONE_MFA,
				customtypes.ENUM_EXPORT_SERVICE_PINGONE_PROTECT,
			},
		},
		{
			name:         "non existant group",
			cType:        utils.Pointer(customtypes.ExportServiceGroup("non-existant")),
			expectedStrs: []string{},
		},
		{
			name:         "Nil custom type",
			cType:        nil,
			expectedStrs: []string{},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			testutils_koanf.InitKoanfs(t)

			actualStrs := tc.cType.GetServicesInGroup()

			require.Equal(t, tc.expectedStrs, actualStrs)
		})
	}
}
