// Copyright © 2025 Ping Identity Corporation

package customtypes_test

import (
	"testing"

	"github.com/pingidentity/pingcli/internal/customtypes"
	"github.com/pingidentity/pingcli/internal/testing/testutils_koanf"
	"github.com/pingidentity/pingcli/internal/utils"
	"github.com/stretchr/testify/require"
)

func Test_AuthServices_GetServices(t *testing.T) {
	testutils_koanf.InitKoanfs(t)

	testCases := []struct {
		name         string
		cType        *customtypes.AuthServices
		expectedStrs []string
	}{
		{
			name:         "Happy path",
			cType:        new(customtypes.AuthServices),
			expectedStrs: nil,
		},
		{
			name:  "Happy path - multiple services",
			cType: utils.Pointer(customtypes.AuthServices([]string{customtypes.ENUM_AUTH_SERVICE_PINGONE, customtypes.ENUM_AUTH_SERVICE_PINGFEDERATE})),
			expectedStrs: []string{
				customtypes.ENUM_AUTH_SERVICE_PINGONE,
				customtypes.ENUM_AUTH_SERVICE_PINGFEDERATE,
			},
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

			actualStrs := tc.cType.GetServices()

			require.Equal(t, tc.expectedStrs, actualStrs)
		})
	}
}

func Test_AuthServices_Set(t *testing.T) {
	testutils_koanf.InitKoanfs(t)

	testCases := []struct {
		name                string
		cType               *customtypes.AuthServices
		servicesStrs        []string
		expectedNumServices int
		expectedError       error
	}{
		{
			name:                "Happy path - pingone",
			cType:               new(customtypes.AuthServices),
			servicesStrs:        []string{customtypes.ENUM_AUTH_SERVICE_PINGONE},
			expectedNumServices: 1,
		},
		{
			name:                "Happy path - pingfederate",
			cType:               new(customtypes.AuthServices),
			servicesStrs:        []string{customtypes.ENUM_AUTH_SERVICE_PINGFEDERATE},
			expectedNumServices: 1,
		},
		{
			name:                "Happy path - multiple services",
			cType:               new(customtypes.AuthServices),
			servicesStrs:        []string{customtypes.ENUM_AUTH_SERVICE_PINGONE, customtypes.ENUM_AUTH_SERVICE_PINGFEDERATE},
			expectedNumServices: 2,
		},
		{
			name:                "Happy path - duplicate services",
			cType:               new(customtypes.AuthServices),
			servicesStrs:        []string{customtypes.ENUM_AUTH_SERVICE_PINGFEDERATE, customtypes.ENUM_AUTH_SERVICE_PINGONE, customtypes.ENUM_AUTH_SERVICE_PINGFEDERATE},
			expectedNumServices: 2,
		},
		{
			name:                "Happy path - empty",
			cType:               new(customtypes.AuthServices),
			servicesStrs:        []string{""},
			expectedNumServices: 0,
		},
		{
			name:                "Happy path - comma separated",
			cType:               new(customtypes.AuthServices),
			servicesStrs:        []string{"pingone,pingfederate"},
			expectedNumServices: 2,
		},
		{
			name:                "Happy path - case insensitive",
			cType:               new(customtypes.AuthServices),
			servicesStrs:        []string{"PingOne", "PINGFEDERATE"},
			expectedNumServices: 2,
		},
		{
			name:                "Invalid value",
			cType:               new(customtypes.AuthServices),
			servicesStrs:        []string{"invalid"},
			expectedNumServices: 0,
			expectedError:       customtypes.ErrUnrecognizedAuthService,
		},
		{
			name:                "Nil custom type",
			cType:               nil,
			servicesStrs:        []string{customtypes.ENUM_AUTH_SERVICE_PINGONE},
			expectedNumServices: 0,
			expectedError:       customtypes.ErrCustomTypeNil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			testutils_koanf.InitKoanfs(t)

			for _, servicesStr := range tc.servicesStrs {
				err := tc.cType.Set(servicesStr)

				if tc.expectedError != nil {
					require.Error(t, err)
					require.ErrorIs(t, err, tc.expectedError)
				} else {
					require.NoError(t, err)
				}
			}

			require.Equal(t, tc.expectedNumServices, len(tc.cType.GetServices()))
		})
	}
}

func Test_AuthServices_ContainsPingOne(t *testing.T) {
	testutils_koanf.InitKoanfs(t)

	testCases := []struct {
		name         string
		cType        *customtypes.AuthServices
		expectedBool bool
	}{
		{
			name:         "Happy path - pingone",
			cType:        utils.Pointer(customtypes.AuthServices([]string{customtypes.ENUM_AUTH_SERVICE_PINGONE})),
			expectedBool: true,
		},
		{
			name:         "Happy path - pingfederate",
			cType:        utils.Pointer(customtypes.AuthServices([]string{customtypes.ENUM_AUTH_SERVICE_PINGFEDERATE})),
			expectedBool: false,
		},
		{
			name:         "Happy path - both services",
			cType:        utils.Pointer(customtypes.AuthServices([]string{customtypes.ENUM_AUTH_SERVICE_PINGONE, customtypes.ENUM_AUTH_SERVICE_PINGFEDERATE})),
			expectedBool: true,
		},
		{
			name:         "Happy path - empty",
			cType:        utils.Pointer(customtypes.AuthServices([]string{""})),
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

func Test_AuthServices_ContainsPingFederate(t *testing.T) {
	testutils_koanf.InitKoanfs(t)

	testCases := []struct {
		name         string
		cType        *customtypes.AuthServices
		expectedBool bool
	}{
		{
			name:         "Happy path - pingfederate",
			cType:        utils.Pointer(customtypes.AuthServices([]string{customtypes.ENUM_AUTH_SERVICE_PINGFEDERATE})),
			expectedBool: true,
		},
		{
			name:         "Happy path - pingone",
			cType:        utils.Pointer(customtypes.AuthServices([]string{customtypes.ENUM_AUTH_SERVICE_PINGONE})),
			expectedBool: false,
		},
		{
			name:         "Happy path - both services",
			cType:        utils.Pointer(customtypes.AuthServices([]string{customtypes.ENUM_AUTH_SERVICE_PINGONE, customtypes.ENUM_AUTH_SERVICE_PINGFEDERATE})),
			expectedBool: true,
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

			actualBool := tc.cType.ContainsPingFederate()

			require.Equal(t, tc.expectedBool, actualBool)
		})
	}
}

func Test_AuthServices_Type(t *testing.T) {
	testutils_koanf.InitKoanfs(t)

	testCases := []struct {
		name         string
		cType        *customtypes.AuthServices
		expectedType string
	}{
		{
			name:         "Happy path",
			cType:        utils.Pointer(customtypes.AuthServices([]string{customtypes.ENUM_AUTH_SERVICE_PINGONE})),
			expectedType: "[]string",
		},
		{
			name:         "Happy path - empty",
			cType:        utils.Pointer(customtypes.AuthServices([]string{""})),
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

func Test_AuthServices_String(t *testing.T) {
	testutils_koanf.InitKoanfs(t)

	testCases := []struct {
		name        string
		cType       *customtypes.AuthServices
		expectedStr string
	}{
		{
			name: "Happy path - pingone",
			cType: utils.Pointer(customtypes.AuthServices([]string{
				customtypes.ENUM_AUTH_SERVICE_PINGONE,
			})),
			expectedStr: customtypes.ENUM_AUTH_SERVICE_PINGONE,
		},
		{
			name: "Happy path - both services",
			cType: utils.Pointer(customtypes.AuthServices([]string{
				customtypes.ENUM_AUTH_SERVICE_PINGONE,
				customtypes.ENUM_AUTH_SERVICE_PINGFEDERATE,
			})),
			expectedStr: "pingfederate,pingone",
		},
		{
			name: "Test ordering - reverse input",
			cType: utils.Pointer(customtypes.AuthServices([]string{
				customtypes.ENUM_AUTH_SERVICE_PINGONE,
				customtypes.ENUM_AUTH_SERVICE_PINGFEDERATE,
			})),
			expectedStr: "pingfederate,pingone",
		},
		{
			name:        "Nil custom type",
			cType:       nil,
			expectedStr: "",
		},
		{
			name:        "Happy path - empty",
			cType:       utils.Pointer(customtypes.AuthServices([]string{""})),
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

func Test_AuthServicesValidValues(t *testing.T) {
	expectedServices := []string{
		customtypes.ENUM_AUTH_SERVICE_PINGFEDERATE,
		customtypes.ENUM_AUTH_SERVICE_PINGONE,
	}

	actualServices := customtypes.AuthServicesValidValues()
	require.Equal(t, actualServices, expectedServices)
	require.Equal(t, len(actualServices), len(expectedServices))
}

func Test_AuthServicesValidValuesMap(t *testing.T) {
	expectedMap := map[string]string{
		"pingfederate": customtypes.ENUM_AUTH_SERVICE_PINGFEDERATE,
		"pingone":      customtypes.ENUM_AUTH_SERVICE_PINGONE,
	}

	actualMap := customtypes.AuthServicesValidValuesMap()
	require.Equal(t, expectedMap, actualMap)
	require.Equal(t, len(expectedMap), len(actualMap))
}
