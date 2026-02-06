// Copyright © 2026 Ping Identity Corporation

package customtypes_test

import (
	"testing"

	"github.com/pingidentity/pingcli/internal/customtypes"
	"github.com/pingidentity/pingcli/internal/testing/testutils_koanf"
	"github.com/pingidentity/pingcli/internal/utils"
	"github.com/stretchr/testify/require"
)

func Test_ExportFormat_Set(t *testing.T) {
	testutils_koanf.InitKoanfs(t)

	testCases := []struct {
		name          string
		cType         *customtypes.ExportFormat
		formatStr     string
		expectedError error
	}{
		{
			name:      "Happy path - HCL",
			cType:     new(customtypes.ExportFormat),
			formatStr: customtypes.ENUM_EXPORT_FORMAT_HCL,
		},
		{
			name:      "Happy path - empty",
			cType:     new(customtypes.ExportFormat),
			formatStr: "",
		},
		{
			name:          "Invalid value",
			cType:         new(customtypes.ExportFormat),
			formatStr:     "invalid",
			expectedError: customtypes.ErrUnrecognizedFormat,
		},
		{
			name:          "Nil custom type",
			cType:         nil,
			formatStr:     customtypes.ENUM_EXPORT_FORMAT_HCL,
			expectedError: customtypes.ErrCustomTypeNil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			testutils_koanf.InitKoanfs(t)

			err := tc.cType.Set(tc.formatStr)

			if tc.expectedError != nil {
				require.Error(t, err)
				require.ErrorIs(t, err, tc.expectedError)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func Test_ExportFormat_Type(t *testing.T) {
	testutils_koanf.InitKoanfs(t)

	testCases := []struct {
		name         string
		cType        *customtypes.ExportFormat
		expectedType string
	}{
		{
			name:         "Happy path - HCL",
			cType:        utils.Pointer(customtypes.ExportFormat(customtypes.ENUM_EXPORT_FORMAT_HCL)),
			expectedType: "string",
		},
		{
			name:         "Happy path - empty",
			cType:        utils.Pointer(customtypes.ExportFormat("")),
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

func Test_ExportFormat_String(t *testing.T) {
	testutils_koanf.InitKoanfs(t)

	testCases := []struct {
		name        string
		cType       *customtypes.ExportFormat
		expectedStr string
	}{
		{
			name:        "Happy path - HCL",
			cType:       utils.Pointer(customtypes.ExportFormat(customtypes.ENUM_EXPORT_FORMAT_HCL)),
			expectedStr: customtypes.ENUM_EXPORT_FORMAT_HCL,
		},
		{
			name:        "Happy path - Empty",
			cType:       utils.Pointer(customtypes.ExportFormat("")),
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

func Test_ExportFormatValidValues(t *testing.T) {
	testutils_koanf.InitKoanfs(t)

	expectedValues := []string{
		customtypes.ENUM_EXPORT_FORMAT_HCL,
	}

	actualValues := customtypes.ExportFormatValidValues()

	require.Equal(t, expectedValues, actualValues)
}
