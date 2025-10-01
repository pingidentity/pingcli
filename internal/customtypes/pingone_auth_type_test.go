// Copyright © 2025 Ping Identity Corporation

package customtypes_test

import (
	"testing"

	"github.com/pingidentity/pingcli/internal/customtypes"
	"github.com/pingidentity/pingcli/internal/testing/testutils_koanf"
	"github.com/pingidentity/pingcli/internal/utils"
	"github.com/stretchr/testify/require"
)

func Test_PingOneAuthType_Set(t *testing.T) {
	testutils_koanf.InitKoanfs(t)

	testCases := []struct {
		name          string
		cType         *customtypes.PingOneAuthenticationType
		value         string
		expectedError error
	}{
		{
			name:  "Happy path",
			cType: new(customtypes.PingOneAuthenticationType),
			value: customtypes.ENUM_PINGONE_AUTHENTICATION_TYPE_WORKER,
		},
		{
			name:  "Happy path - empty",
			cType: new(customtypes.PingOneAuthenticationType),
			value: "",
		},
		{
			name:          "Invalid value",
			cType:         new(customtypes.PingOneAuthenticationType),
			value:         "invalid",
			expectedError: customtypes.ErrUnrecognizedPingOneAuth,
		},
		{
			name:          "Nil custom type",
			cType:         nil,
			value:         customtypes.ENUM_PINGONE_AUTHENTICATION_TYPE_WORKER,
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

func Test_PingOneAuthType_Type(t *testing.T) {
	testutils_koanf.InitKoanfs(t)

	testCases := []struct {
		name         string
		cType        *customtypes.PingOneAuthenticationType
		expectedType string
	}{
		{
			name:         "Happy path",
			cType:        utils.Pointer(customtypes.PingOneAuthenticationType(customtypes.ENUM_PINGONE_AUTHENTICATION_TYPE_WORKER)),
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

func Test_PingOneAuthType_String(t *testing.T) {
	testutils_koanf.InitKoanfs(t)

	testCases := []struct {
		name        string
		cType       *customtypes.PingOneAuthenticationType
		expectedStr string
	}{
		{
			name:        "Happy path",
			cType:       utils.Pointer(customtypes.PingOneAuthenticationType(customtypes.ENUM_PINGONE_AUTHENTICATION_TYPE_WORKER)),
			expectedStr: customtypes.ENUM_PINGONE_AUTHENTICATION_TYPE_WORKER,
		},
		{
			name:        "Happy path - empty",
			cType:       utils.Pointer(customtypes.PingOneAuthenticationType("")),
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
