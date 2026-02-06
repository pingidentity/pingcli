// Copyright © 2026 Ping Identity Corporation

package customtypes_test

import (
	"testing"

	"github.com/pingidentity/pingcli/internal/customtypes"
	"github.com/pingidentity/pingcli/internal/testing/testutils_koanf"
	"github.com/pingidentity/pingcli/internal/utils"
	"github.com/stretchr/testify/require"
)

func Test_OutputFormat_Set(t *testing.T) {
	testutils_koanf.InitKoanfs(t)

	testCases := []struct {
		name          string
		cType         *customtypes.OutputFormat
		value         string
		expectedError error
	}{
		{
			name:  "Happy path - JSON",
			cType: new(customtypes.OutputFormat),
			value: customtypes.ENUM_OUTPUT_FORMAT_JSON,
		},
		{
			name:  "Happy path - empty",
			cType: new(customtypes.OutputFormat),
			value: "",
		},
		{
			name:          "Invalid value",
			cType:         new(customtypes.OutputFormat),
			value:         "invalid",
			expectedError: customtypes.ErrUnrecognizedOutputFormat,
		},
		{
			name:          "Nil custom type",
			cType:         nil,
			value:         customtypes.ENUM_OUTPUT_FORMAT_JSON,
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

func Test_OutputFormat_Type(t *testing.T) {
	testutils_koanf.InitKoanfs(t)

	testCases := []struct {
		name         string
		cType        *customtypes.OutputFormat
		expectedType string
	}{
		{
			name:         "Happy path",
			cType:        utils.Pointer(customtypes.OutputFormat(customtypes.ENUM_OUTPUT_FORMAT_JSON)),
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

func Test_OutputFormat_String(t *testing.T) {
	testutils_koanf.InitKoanfs(t)

	testCases := []struct {
		name        string
		cType       *customtypes.OutputFormat
		expectedStr string
	}{
		{
			name:        "Happy path",
			cType:       utils.Pointer(customtypes.OutputFormat(customtypes.ENUM_OUTPUT_FORMAT_JSON)),
			expectedStr: customtypes.ENUM_OUTPUT_FORMAT_JSON,
		},
		{
			name:        "Happy path - empty",
			cType:       utils.Pointer(customtypes.OutputFormat("")),
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
