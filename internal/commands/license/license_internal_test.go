// Copyright © 2025 Ping Identity Corporation

package license_internal

import (
	"testing"

	"github.com/pingidentity/pingcli/internal/configuration/options"
	"github.com/pingidentity/pingcli/internal/customtypes"
	"github.com/pingidentity/pingcli/internal/testing/testutils_koanf"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_RunInternalLicense(t *testing.T) {
	testCases := []struct {
		name          string
		product       customtypes.LicenseProduct
		version       customtypes.LicenseVersion
		devopsUser    customtypes.String
		devopsKey     customtypes.String
		expectedError error
	}{
		{
			name:    "Request PingFederate 13.0 License",
			product: customtypes.LicenseProduct(customtypes.ENUM_LICENSE_PRODUCT_PING_FEDERATE),
			version: "13.0",
		},
		{
			name:          "Request license with empty product",
			product:       "",
			version:       "13.0",
			expectedError: ErrRequiredValues,
		},
		{
			name:          "Request license with empty version",
			product:       customtypes.LicenseProduct(customtypes.ENUM_LICENSE_PRODUCT_PING_FEDERATE),
			version:       "",
			expectedError: ErrRequiredValues,
		},
		{
			name:          "Request license with invalid devops key",
			product:       customtypes.LicenseProduct(customtypes.ENUM_LICENSE_PRODUCT_PING_FEDERATE),
			version:       "13.0",
			devopsKey:     "invalid-key",
			expectedError: ErrLicenseRequest,
		},
		{
			name:          "Request license with invalid devops user",
			product:       customtypes.LicenseProduct(customtypes.ENUM_LICENSE_PRODUCT_PING_FEDERATE),
			version:       "13.0",
			devopsUser:    "invalid-user",
			expectedError: ErrLicenseRequest,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			testutils_koanf.InitKoanfs(t)

			options.LicenseProductOption.Flag.Changed = true
			options.LicenseProductOption.CobraParamValue = &tc.product

			options.LicenseVersionOption.Flag.Changed = true
			options.LicenseVersionOption.CobraParamValue = &tc.version

			if tc.devopsUser != "" {
				options.LicenseDevopsUserOption.Flag.Changed = true
				options.LicenseDevopsUserOption.CobraParamValue = &tc.devopsUser
			}

			if tc.devopsKey != "" {
				options.LicenseDevopsKeyOption.Flag.Changed = true
				options.LicenseDevopsKeyOption.CobraParamValue = &tc.devopsKey
			}

			err := RunInternalLicense()

			if tc.expectedError != nil {
				require.Error(t, err)
				assert.ErrorIs(t, err, tc.expectedError)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
