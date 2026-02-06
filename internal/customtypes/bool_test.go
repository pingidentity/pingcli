// Copyright © 2026 Ping Identity Corporation

package customtypes_test

import (
	"testing"

	"github.com/pingidentity/pingcli/internal/customtypes"
	"github.com/pingidentity/pingcli/internal/testing/testutils_koanf"
	"github.com/pingidentity/pingcli/internal/utils"
	"github.com/stretchr/testify/require"
)

func Test_Bool_Set(t *testing.T) {
	testutils_koanf.InitKoanfs(t)

	testCases := []struct {
		name          string
		cType         *customtypes.Bool
		boolStr       string
		expectedError error
	}{
		{
			name:    "Happy path - true",
			cType:   new(customtypes.Bool),
			boolStr: "true",
		},
		{
			name:    "Happy path - false",
			cType:   new(customtypes.Bool),
			boolStr: "false",
		},
		{
			name:          "Invalid value",
			cType:         new(customtypes.Bool),
			boolStr:       "invalid",
			expectedError: customtypes.ErrParseBool,
		},
		{
			name:          "Empty value",
			cType:         new(customtypes.Bool),
			boolStr:       "",
			expectedError: customtypes.ErrParseBool,
		},
		{
			name:          "Nil custom bool type",
			cType:         nil,
			boolStr:       "true",
			expectedError: customtypes.ErrCustomTypeNil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			testutils_koanf.InitKoanfs(t)

			err := tc.cType.Set(tc.boolStr)

			if tc.expectedError != nil {
				require.Error(t, err)
				require.ErrorIs(t, err, tc.expectedError)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func Test_Bool_Type(t *testing.T) {
	testutils_koanf.InitKoanfs(t)

	testCases := []struct {
		name         string
		cType        *customtypes.Bool
		expectedType string
	}{
		{
			name:         "Happy path - true",
			cType:        utils.Pointer(customtypes.Bool(true)),
			expectedType: "bool",
		},
		{
			name:         "Happy path - false",
			cType:        utils.Pointer(customtypes.Bool(false)),
			expectedType: "bool",
		},
		{
			name:         "Nil custom bool type",
			cType:        nil,
			expectedType: "bool",
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

func Test_Bool_String(t *testing.T) {
	testutils_koanf.InitKoanfs(t)

	testCases := []struct {
		name        string
		cType       *customtypes.Bool
		expectedStr string
	}{
		{
			name:        "Happy path - true",
			cType:       utils.Pointer(customtypes.Bool(true)),
			expectedStr: "true",
		},
		{
			name:        "Happy path - false",
			cType:       utils.Pointer(customtypes.Bool(false)),
			expectedStr: "false",
		},
		{
			name:        "Nil custom bool type",
			cType:       nil,
			expectedStr: "false",
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

func Test_Bool_Bool(t *testing.T) {
	testutils_koanf.InitKoanfs(t)

	testCases := []struct {
		name         string
		cType        *customtypes.Bool
		expectedBool bool
	}{
		{
			name:         "Happy path - true",
			cType:        utils.Pointer(customtypes.Bool(true)),
			expectedBool: true,
		},
		{
			name:         "Happy path - false",
			cType:        utils.Pointer(customtypes.Bool(false)),
			expectedBool: false,
		},
		{
			name:         "Nil custom bool type",
			cType:        nil,
			expectedBool: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			testutils_koanf.InitKoanfs(t)

			actualBool := tc.cType.Bool()

			require.Equal(t, tc.expectedBool, actualBool)
		})
	}
}
