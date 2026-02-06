// Copyright © 2026 Ping Identity Corporation

package customtypes_test

import (
	"testing"

	"github.com/pingidentity/pingcli/internal/customtypes"
	"github.com/pingidentity/pingcli/internal/testing/testutils_koanf"
	"github.com/pingidentity/pingcli/internal/utils"
	"github.com/stretchr/testify/require"
)

func Test_Int_Set(t *testing.T) {
	testutils_koanf.InitKoanfs(t)

	testCases := []struct {
		name          string
		cType         *customtypes.Int
		value         string
		expectedError error
	}{
		{
			name:  "Happy path",
			cType: new(customtypes.Int),
			value: "42",
		},
		{
			name:          "Invalid value",
			cType:         new(customtypes.Int),
			value:         "invalid",
			expectedError: customtypes.ErrParseInt,
		},
		{
			name:          "Empty value",
			cType:         new(customtypes.Int),
			value:         "",
			expectedError: customtypes.ErrParseInt,
		},
		{
			name:          "Nil custom type",
			cType:         nil,
			value:         "42",
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

func Test_Int_Type(t *testing.T) {
	testutils_koanf.InitKoanfs(t)

	testCases := []struct {
		name         string
		cType        *customtypes.Int
		expectedType string
	}{
		{
			name:         "Happy path",
			cType:        utils.Pointer(customtypes.Int(42)),
			expectedType: "int64",
		},
		{
			name:         "Nil custom type",
			cType:        nil,
			expectedType: "int64",
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

func Test_Int_String(t *testing.T) {
	testutils_koanf.InitKoanfs(t)

	testCases := []struct {
		name        string
		cType       *customtypes.Int
		expectedStr string
	}{
		{
			name:        "Happy path",
			cType:       utils.Pointer(customtypes.Int(42)),
			expectedStr: "42",
		},
		{
			name:        "Nil custom type",
			cType:       nil,
			expectedStr: "0",
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

func Test_Int_Int64(t *testing.T) {
	testutils_koanf.InitKoanfs(t)

	testCases := []struct {
		name        string
		cType       *customtypes.Int
		expectedInt int64
	}{
		{
			name:        "Happy path",
			cType:       utils.Pointer(customtypes.Int(42)),
			expectedInt: 42,
		},
		{
			name:        "Nil custom type",
			cType:       nil,
			expectedInt: 0,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			testutils_koanf.InitKoanfs(t)

			actualInt := tc.cType.Int64()

			require.Equal(t, tc.expectedInt, actualInt)
		})
	}
}
