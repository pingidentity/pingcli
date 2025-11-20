// Copyright © 2025 Ping Identity Corporation

package customtypes_test

import (
	"testing"

	"github.com/pingidentity/pingcli/internal/customtypes"
	"github.com/pingidentity/pingcli/internal/testing/testutils_koanf"
	"github.com/pingidentity/pingcli/internal/utils"
	"github.com/stretchr/testify/require"
)

func Test_AuthProvider_Set(t *testing.T) {
	testutils_koanf.InitKoanfs(t)

	testCases := []struct {
		name            string
		cType           *customtypes.AuthProvider
		providerStr     string
		expectedService string
		expectedError   error
	}{
		{
			name:            "Happy path - pingone",
			cType:           new(customtypes.AuthProvider),
			providerStr:     customtypes.ENUM_AUTH_PROVIDER_PINGONE,
			expectedService: customtypes.ENUM_AUTH_PROVIDER_PINGONE,
		},
		{
			name:            "Happy path - case insensitive uppercase",
			cType:           new(customtypes.AuthProvider),
			providerStr:     "PINGONE",
			expectedService: customtypes.ENUM_AUTH_PROVIDER_PINGONE,
		},
		{
			name:            "Happy path - case insensitive mixed",
			cType:           new(customtypes.AuthProvider),
			providerStr:     "PingOne",
			expectedService: customtypes.ENUM_AUTH_PROVIDER_PINGONE,
		},
		{
			name:            "Happy path - with whitespace",
			cType:           new(customtypes.AuthProvider),
			providerStr:     "  pingone  ",
			expectedService: customtypes.ENUM_AUTH_PROVIDER_PINGONE,
		},
		{
			name:            "Happy path - empty string",
			cType:           new(customtypes.AuthProvider),
			providerStr:     "",
			expectedService: "",
		},
		{
			name:          "Invalid value",
			cType:         new(customtypes.AuthProvider),
			providerStr:   "invalid",
			expectedError: customtypes.ErrUnrecognizedAuthProvider,
		},
		{
			name:          "Invalid value - pingfederate not yet supported",
			cType:         new(customtypes.AuthProvider),
			providerStr:   "pingfederate",
			expectedError: customtypes.ErrUnrecognizedAuthProvider,
		},
		{
			name:          "Nil custom type",
			cType:         nil,
			providerStr:   customtypes.ENUM_AUTH_PROVIDER_PINGONE,
			expectedError: customtypes.ErrCustomTypeNil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			testutils_koanf.InitKoanfs(t)

			err := tc.cType.Set(tc.providerStr)

			if tc.expectedError != nil {
				require.Error(t, err)
				require.ErrorIs(t, err, tc.expectedError)
			} else {
				require.NoError(t, err)
				if tc.cType != nil {
					require.Equal(t, tc.expectedService, string(*tc.cType))
				}
			}
		})
	}
}

func Test_AuthProvider_ContainsPingOne(t *testing.T) {
	testutils_koanf.InitKoanfs(t)

	testCases := []struct {
		name         string
		cType        *customtypes.AuthProvider
		expectedBool bool
	}{
		{
			name:         "Happy path - pingone",
			cType:        utils.Pointer(customtypes.AuthProvider(customtypes.ENUM_AUTH_PROVIDER_PINGONE)),
			expectedBool: true,
		},
		{
			name:         "Happy path - empty",
			cType:        utils.Pointer(customtypes.AuthProvider("")),
			expectedBool: false,
		},
		{
			name:         "Nil custom type",
			cType:        nil,
			expectedBool: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			testutils_koanf.InitKoanfs(t)

			actualBool := tc.cType.ContainsPingOne()

			require.Equal(t, tc.expectedBool, actualBool)
		})
	}
}

func Test_AuthProvider_Type(t *testing.T) {
	testutils_koanf.InitKoanfs(t)

	testCases := []struct {
		name         string
		cType        *customtypes.AuthProvider
		expectedType string
	}{
		{
			name:         "Happy path",
			cType:        utils.Pointer(customtypes.AuthProvider(customtypes.ENUM_AUTH_PROVIDER_PINGONE)),
			expectedType: "string",
		},
		{
			name:         "Happy path - empty",
			cType:        utils.Pointer(customtypes.AuthProvider("")),
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

func Test_AuthProvider_String(t *testing.T) {
	testutils_koanf.InitKoanfs(t)

	testCases := []struct {
		name        string
		cType       *customtypes.AuthProvider
		expectedStr string
	}{
		{
			name:        "Happy path - pingone",
			cType:       utils.Pointer(customtypes.AuthProvider(customtypes.ENUM_AUTH_PROVIDER_PINGONE)),
			expectedStr: customtypes.ENUM_AUTH_PROVIDER_PINGONE,
		},
		{
			name:        "Happy path - empty",
			cType:       utils.Pointer(customtypes.AuthProvider("")),
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

func Test_AuthProviderValidValues(t *testing.T) {
	expectedServices := []string{
		customtypes.ENUM_AUTH_PROVIDER_PINGONE,
	}

	actualServices := customtypes.AuthProviderValidValues()
	require.Equal(t, expectedServices, actualServices)
	require.Equal(t, len(expectedServices), len(actualServices))
}

func Test_AuthProviderValidValuesMap(t *testing.T) {
	expectedMap := map[string]string{
		"pingone": customtypes.ENUM_AUTH_PROVIDER_PINGONE,
	}

	actualMap := customtypes.AuthProviderValidValuesMap()
	require.Equal(t, expectedMap, actualMap)
	require.Equal(t, len(expectedMap), len(actualMap))
}
