// Copyright © 2025 Ping Identity Corporation

package customtypes_test

import (
	"fmt"
	"testing"

	"github.com/pingidentity/pingcli/internal/customtypes"
	"github.com/pingidentity/pingcli/internal/testing/testutils_koanf"
	"github.com/pingidentity/pingcli/internal/utils"
	"github.com/stretchr/testify/require"
)

func Test_ExportServices_GetServices(t *testing.T) {
	testutils_koanf.InitKoanfs(t)

	testCases := []struct {
		name         string
		cType        *customtypes.ExportServices
		expectedStrs []string
	}{
		{
			name:         "Happy path",
			cType:        new(customtypes.ExportServices),
			expectedStrs: nil,
		},
		{
			name:  "Happy path - multiple services",
			cType: utils.Pointer(customtypes.ExportServices([]string{customtypes.ENUM_EXPORT_SERVICE_PINGONE_MFA, customtypes.ENUM_EXPORT_SERVICE_PINGONE_SSO})),
			expectedStrs: []string{
				customtypes.ENUM_EXPORT_SERVICE_PINGONE_MFA,
				customtypes.ENUM_EXPORT_SERVICE_PINGONE_SSO,
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

func Test_ExportServices_Set(t *testing.T) {
	testutils_koanf.InitKoanfs(t)

	testCases := []struct {
		name                string
		cType               *customtypes.ExportServices
		servicesStrs        []string
		expectedNumServices int
		expectedError       error
	}{
		{
			name:                "Happy path",
			cType:               new(customtypes.ExportServices),
			servicesStrs:        []string{customtypes.ENUM_EXPORT_SERVICE_PINGONE_MFA},
			expectedNumServices: 1,
		},
		{
			name:                "Happy path - multiple services",
			cType:               new(customtypes.ExportServices),
			servicesStrs:        []string{customtypes.ENUM_EXPORT_SERVICE_PINGONE_MFA, customtypes.ENUM_EXPORT_SERVICE_PINGONE_SSO},
			expectedNumServices: 2,
		},
		{
			name:                "Happy path - duplicate services",
			cType:               new(customtypes.ExportServices),
			servicesStrs:        []string{customtypes.ENUM_EXPORT_SERVICE_PINGFEDERATE, customtypes.ENUM_EXPORT_SERVICE_PINGONE_MFA, customtypes.ENUM_EXPORT_SERVICE_PINGFEDERATE},
			expectedNumServices: 2,
		},
		{
			name:                "Happy path - empty",
			cType:               new(customtypes.ExportServices),
			servicesStrs:        []string{""},
			expectedNumServices: 0,
		},
		{
			name:                "Invalid value",
			cType:               new(customtypes.ExportServices),
			servicesStrs:        []string{"invalid"},
			expectedNumServices: 0,
			expectedError:       customtypes.ErrUnrecognisedExportService,
		},
		{
			name:                "Nil custom type",
			cType:               nil,
			servicesStrs:        []string{customtypes.ENUM_EXPORT_SERVICE_PINGONE_MFA},
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

func Test_ExportServices_SetServicesByServiceGroup(t *testing.T) {
	testutils_koanf.InitKoanfs(t)

	testCases := []struct {
		name                string
		cType               *customtypes.ExportServices
		serviceGroup        *customtypes.ExportServiceGroup
		expectedNumServices int
		expectedError       error
	}{
		{
			name:                "Happy path - pingone",
			cType:               new(customtypes.ExportServices),
			serviceGroup:        utils.Pointer(customtypes.ExportServiceGroup(customtypes.ENUM_EXPORT_SERVICE_GROUP_PINGONE)),
			expectedNumServices: 5,
		},
		{
			name:                "Happy path - empty service group",
			cType:               new(customtypes.ExportServices),
			serviceGroup:        utils.Pointer(customtypes.ExportServiceGroup("")),
			expectedNumServices: 0,
		},
		{
			name:                "Nil custom type",
			cType:               nil,
			serviceGroup:        utils.Pointer(customtypes.ExportServiceGroup(customtypes.ENUM_EXPORT_SERVICE_GROUP_PINGONE)),
			expectedNumServices: 0,
			expectedError:       customtypes.ErrCustomTypeNil,
		},
		{
			name:                "Nil service group",
			cType:               new(customtypes.ExportServices),
			serviceGroup:        nil,
			expectedNumServices: 0,
		},
		{
			name:                "Invalid service group",
			cType:               new(customtypes.ExportServices),
			serviceGroup:        utils.Pointer(customtypes.ExportServiceGroup("invalid")),
			expectedNumServices: 0,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			testutils_koanf.InitKoanfs(t)

			err := tc.cType.SetServicesByServiceGroup(tc.serviceGroup)

			if tc.expectedError != nil {
				require.Error(t, err)
				require.ErrorIs(t, err, tc.expectedError)
			} else {
				require.NoError(t, err)
			}

			require.Equal(t, tc.expectedNumServices, len(tc.cType.GetServices()))
		})
	}
}

func Test_ExportServices_ContainsPingOneService(t *testing.T) {
	testutils_koanf.InitKoanfs(t)

	testCases := []struct {
		name         string
		cType        *customtypes.ExportServices
		expectedBool bool
	}{
		{
			name:         "Happy path - pingone mfa",
			cType:        utils.Pointer(customtypes.ExportServices([]string{customtypes.ENUM_EXPORT_SERVICE_PINGONE_MFA})),
			expectedBool: true,
		},
		{
			name:         "Happy path - pingone sso",
			cType:        utils.Pointer(customtypes.ExportServices([]string{customtypes.ENUM_EXPORT_SERVICE_PINGONE_SSO})),
			expectedBool: true,
		},
		{
			name:         "Happy path - pingone platform",
			cType:        utils.Pointer(customtypes.ExportServices([]string{customtypes.ENUM_EXPORT_SERVICE_PINGONE_PLATFORM})),
			expectedBool: true,
		},
		{
			name:         "Happy path - pingone authorize",
			cType:        utils.Pointer(customtypes.ExportServices([]string{customtypes.ENUM_EXPORT_SERVICE_PINGONE_AUTHORIZE})),
			expectedBool: true,
		},
		{
			name:         "Happy path - pingone protect",
			cType:        utils.Pointer(customtypes.ExportServices([]string{customtypes.ENUM_EXPORT_SERVICE_PINGONE_PROTECT})),
			expectedBool: true,
		},
		{
			name:         "Happy path - pingfederate",
			cType:        utils.Pointer(customtypes.ExportServices([]string{customtypes.ENUM_EXPORT_SERVICE_PINGFEDERATE})),
			expectedBool: false,
		},
		{
			name:         "Happy path - empty",
			cType:        utils.Pointer(customtypes.ExportServices([]string{""})),
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

			actualBool := tc.cType.ContainsPingOneService()

			require.Equal(t, tc.expectedBool, actualBool)
		})
	}
}

func Test_ExportServices_ContainsPingFederateService(t *testing.T) {
	testutils_koanf.InitKoanfs(t)

	testCases := []struct {
		name         string
		cType        *customtypes.ExportServices
		expectedBool bool
	}{
		{
			name:         "Happy path - pingfederate",
			cType:        utils.Pointer(customtypes.ExportServices([]string{customtypes.ENUM_EXPORT_SERVICE_PINGFEDERATE})),
			expectedBool: true,
		},
		{
			name:         "Happy path - pingone mfa",
			cType:        utils.Pointer(customtypes.ExportServices([]string{customtypes.ENUM_EXPORT_SERVICE_PINGONE_MFA})),
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

			actualBool := tc.cType.ContainsPingFederateService()

			require.Equal(t, tc.expectedBool, actualBool)
		})
	}
}

func Test_ExportServices_Type(t *testing.T) {
	testutils_koanf.InitKoanfs(t)

	testCases := []struct {
		name         string
		cType        *customtypes.ExportServices
		expectedType string
	}{
		{
			name:         "Happy path",
			cType:        utils.Pointer(customtypes.ExportServices([]string{customtypes.ENUM_EXPORT_SERVICE_PINGONE_MFA})),
			expectedType: "[]string",
		},
		{
			name:         "Happy path - empty",
			cType:        utils.Pointer(customtypes.ExportServices([]string{""})),
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

func Test_ExportServices_String(t *testing.T) {
	testutils_koanf.InitKoanfs(t)

	testCases := []struct {
		name        string
		cType       *customtypes.ExportServices
		expectedStr string
	}{
		{
			name: "Happy path",
			cType: utils.Pointer(customtypes.ExportServices([]string{
				customtypes.ENUM_EXPORT_SERVICE_PINGONE_MFA,
			})),
			expectedStr: customtypes.ENUM_EXPORT_SERVICE_PINGONE_MFA,
		},
		{
			name: "Test ordering",
			cType: utils.Pointer(customtypes.ExportServices([]string{
				customtypes.ENUM_EXPORT_SERVICE_PINGONE_SSO,
				customtypes.ENUM_EXPORT_SERVICE_PINGONE_MFA,
			})),
			expectedStr: fmt.Sprintf("%s,%s", customtypes.ENUM_EXPORT_SERVICE_PINGONE_MFA, customtypes.ENUM_EXPORT_SERVICE_PINGONE_SSO),
		},
		{
			name:        "Nil custom type",
			cType:       nil,
			expectedStr: "",
		},
		{
			name:        "Happy path - empty",
			cType:       utils.Pointer(customtypes.ExportServices([]string{""})),
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

func Test_ExportServices_Merge(t *testing.T) {
	testutils_koanf.InitKoanfs(t)

	testCases := []struct {
		name                string
		cType               *customtypes.ExportServices
		es2                 *customtypes.ExportServices
		expectedNumServices int
		expectedServices    []string
		expectedError       error
	}{
		{
			name: "Happy path",
			cType: utils.Pointer(customtypes.ExportServices([]string{
				customtypes.ENUM_EXPORT_SERVICE_PINGONE_MFA,
			})),
			es2: utils.Pointer(customtypes.ExportServices([]string{
				customtypes.ENUM_EXPORT_SERVICE_PINGONE_SSO,
			})),
			expectedNumServices: 2,
			expectedServices: []string{
				customtypes.ENUM_EXPORT_SERVICE_PINGONE_MFA,
				customtypes.ENUM_EXPORT_SERVICE_PINGONE_SSO,
			},
		},
		{
			name: "Test ordering",
			cType: utils.Pointer(customtypes.ExportServices([]string{
				customtypes.ENUM_EXPORT_SERVICE_PINGONE_SSO,
			})),
			es2: utils.Pointer(customtypes.ExportServices([]string{
				customtypes.ENUM_EXPORT_SERVICE_PINGONE_MFA,
			})),
			expectedNumServices: 2,
			expectedServices: []string{
				customtypes.ENUM_EXPORT_SERVICE_PINGONE_MFA,
				customtypes.ENUM_EXPORT_SERVICE_PINGONE_SSO,
			},
		},
		{
			name:                "Happy path - empty",
			cType:               utils.Pointer(customtypes.ExportServices([]string{})),
			es2:                 utils.Pointer(customtypes.ExportServices([]string{})),
			expectedNumServices: 0,
			expectedServices:    []string{},
		},
		{
			name:  "Nil custom type",
			cType: nil,
			es2: utils.Pointer(customtypes.ExportServices([]string{
				customtypes.ENUM_EXPORT_SERVICE_PINGONE_SSO,
			})),
			expectedNumServices: 0,
			expectedServices:    []string{},
			expectedError:       customtypes.ErrCustomTypeNil,
		},
		{
			name: "Nil es2",
			cType: utils.Pointer(customtypes.ExportServices([]string{
				customtypes.ENUM_EXPORT_SERVICE_PINGONE_SSO,
			})),
			es2:                 nil,
			expectedNumServices: 1,
			expectedServices: []string{
				customtypes.ENUM_EXPORT_SERVICE_PINGONE_SSO,
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			testutils_koanf.InitKoanfs(t)

			err := tc.cType.Merge(tc.es2)

			if tc.expectedError != nil {
				require.Error(t, err)
				require.ErrorIs(t, err, tc.expectedError)
			} else {
				require.NoError(t, err)
			}

			require.Equal(t, tc.expectedNumServices, len(tc.cType.GetServices()))
			require.Equal(t, tc.expectedServices, tc.cType.GetServices())
		})
	}
}

func Test_ExportServicesValidValues(t *testing.T) {
	expectedServices := []string{
		customtypes.ENUM_EXPORT_SERVICE_PINGFEDERATE,
		customtypes.ENUM_EXPORT_SERVICE_PINGONE_AUTHORIZE,
		customtypes.ENUM_EXPORT_SERVICE_PINGONE_MFA,
		customtypes.ENUM_EXPORT_SERVICE_PINGONE_PLATFORM,
		customtypes.ENUM_EXPORT_SERVICE_PINGONE_PROTECT,
		customtypes.ENUM_EXPORT_SERVICE_PINGONE_SSO,
	}

	actualServices := customtypes.ExportServicesValidValues()
	require.Equal(t, actualServices, expectedServices)
	require.Equal(t, len(actualServices), len(expectedServices))
}
