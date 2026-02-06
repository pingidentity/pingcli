// Copyright © 2026 Ping Identity Corporation

package customtypes_test

import (
	"testing"

	"github.com/pingidentity/pingcli/internal/customtypes"
	"github.com/pingidentity/pingcli/internal/testing/testutils_koanf"
	"github.com/pingidentity/pingcli/internal/utils"
	"github.com/stretchr/testify/require"
)

func Test_StringSlice_Set(t *testing.T) {
	testutils_koanf.InitKoanfs(t)

	testCases := []struct {
		name          string
		cType         *customtypes.StringSlice
		value         string
		expectedError error
	}{
		{
			name:  "Happy path",
			cType: new(customtypes.StringSlice),
			value: "value1,value2",
		},
		{
			name:  "Happy path - empty",
			cType: new(customtypes.StringSlice),
			value: "",
		},
		{
			name:          "Nil custom type",
			cType:         nil,
			value:         "value1,value2",
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

func Test_StringSlice_Type(t *testing.T) {
	testutils_koanf.InitKoanfs(t)

	testCases := []struct {
		name         string
		cType        *customtypes.StringSlice
		expectedType string
	}{
		{
			name:         "Happy path",
			cType:        utils.Pointer(customtypes.StringSlice([]string{"value1", "value2"})),
			expectedType: "[]string",
		},
		{
			name:         "Nil custom type",
			cType:        nil,
			expectedType: "[]string",
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

func Test_StringSlice_String(t *testing.T) {
	testutils_koanf.InitKoanfs(t)

	testCases := []struct {
		name        string
		cType       *customtypes.StringSlice
		expectedStr string
	}{
		{
			name:        "Happy path",
			cType:       utils.Pointer(customtypes.StringSlice([]string{"value1", "value2"})),
			expectedStr: "value1,value2",
		},
		{
			name:        "Happy path - empty",
			cType:       utils.Pointer(customtypes.StringSlice([]string{})),
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

func Test_StringSlice_StringSlice(t *testing.T) {
	testutils_koanf.InitKoanfs(t)

	testCases := []struct {
		name             string
		cType            *customtypes.StringSlice
		expectedStrSlice []string
	}{
		{
			name:             "Happy path",
			cType:            utils.Pointer(customtypes.StringSlice([]string{"value1", "value2"})),
			expectedStrSlice: []string{"value1", "value2"},
		},
		{
			name:             "Happy path - empty",
			cType:            utils.Pointer(customtypes.StringSlice([]string{})),
			expectedStrSlice: []string{},
		},
		{
			name:             "Nil custom type",
			cType:            nil,
			expectedStrSlice: []string{},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			testutils_koanf.InitKoanfs(t)

			actualStrSlice := tc.cType.StringSlice()

			require.Equal(t, tc.expectedStrSlice, actualStrSlice)
		})
	}
}

func Test_StringSlice_Remove(t *testing.T) {
	testutils_koanf.InitKoanfs(t)

	testCases := []struct {
		name          string
		cType         *customtypes.StringSlice
		value         string
		expectedBool  bool
		expectedError error
	}{
		{
			name:         "Happy path",
			cType:        utils.Pointer(customtypes.StringSlice([]string{"value1", "value2"})),
			value:        "value1",
			expectedBool: true,
		},
		{
			name:         "Happy path - not found",
			cType:        utils.Pointer(customtypes.StringSlice([]string{"value1", "value2"})),
			value:        "value3",
			expectedBool: false,
		},
		{
			name:         "Happy path - empty",
			cType:        utils.Pointer(customtypes.StringSlice([]string{"value1", "value2"})),
			value:        "",
			expectedBool: false,
		},
		{
			name:          "Nil custom type",
			cType:         nil,
			value:         "value1",
			expectedBool:  false,
			expectedError: customtypes.ErrCustomTypeNil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			testutils_koanf.InitKoanfs(t)

			actualBool, err := tc.cType.Remove(tc.value)

			if tc.expectedError != nil {
				require.Error(t, err)
				require.ErrorIs(t, err, tc.expectedError)
			} else {
				require.NoError(t, err)
			}

			require.Equal(t, tc.expectedBool, actualBool)
		})
	}
}
