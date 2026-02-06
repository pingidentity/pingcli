// Copyright © 2026 Ping Identity Corporation

package customtypes_test

import (
	"testing"

	"github.com/pingidentity/pingcli/internal/customtypes"
	"github.com/pingidentity/pingcli/internal/testing/testutils_koanf"
	"github.com/pingidentity/pingcli/internal/utils"
	"github.com/stretchr/testify/require"
)

func Test_UUID_Set(t *testing.T) {
	testutils_koanf.InitKoanfs(t)

	testCases := []struct {
		name          string
		cType         *customtypes.UUID
		value         string
		expectedError error
	}{
		{
			name:  "Happy path",
			cType: new(customtypes.UUID),
			value: "123e4567-e89b-12d3-a456-426614174000",
		},
		{
			name:  "Happy path - empty",
			cType: new(customtypes.UUID),
			value: "",
		},
		{
			name:          "Invalid value",
			cType:         new(customtypes.UUID),
			value:         "invalid",
			expectedError: customtypes.ErrInvalidUUID,
		},
		{
			name:          "Nil custom type",
			cType:         nil,
			value:         "123e4567-e89b-12d3-a456-426614174000",
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

func Test_UUID_Type(t *testing.T) {
	testutils_koanf.InitKoanfs(t)

	testCases := []struct {
		name         string
		cType        *customtypes.UUID
		expectedType string
	}{
		{
			name:         "Happy path",
			cType:        utils.Pointer(customtypes.UUID("123e4567-e89b-12d3-a456-426614174000")),
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

func Test_UUID_String(t *testing.T) {
	testutils_koanf.InitKoanfs(t)

	testCases := []struct {
		name        string
		cType       *customtypes.UUID
		expectedStr string
	}{
		{
			name:        "Happy path",
			cType:       utils.Pointer(customtypes.UUID("123e4567-e89b-12d3-a456-426614174000")),
			expectedStr: "123e4567-e89b-12d3-a456-426614174000",
		},
		{
			name:        "Happy path - empty",
			cType:       utils.Pointer(customtypes.UUID("")),
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
