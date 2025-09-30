// Copyright © 2025 Ping Identity Corporation

package license_test

import (
	"fmt"
	"testing"

	"github.com/pingidentity/pingcli/cmd/common"
	"github.com/pingidentity/pingcli/internal/configuration/options"
	"github.com/pingidentity/pingcli/internal/customtypes"
	"github.com/pingidentity/pingcli/internal/testing/testutils_cobra"
	"github.com/pingidentity/pingcli/internal/testing/testutils_koanf"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_LicenseCommand(t *testing.T) {
	testutils_koanf.InitKoanfs(t)

	testCases := []struct {
		name                string
		args                []string
		expectErr           bool
		expectedErrIs       error
		expectedErrContains string
	}{
		{
			name: "Happy Path",
			args: []string{
				"--" + options.LicenseProductOption.CobraParamName, customtypes.ENUM_LICENSE_PRODUCT_PING_FEDERATE,
				"--" + options.LicenseVersionOption.CobraParamName, "12.0",
			},
			expectErr: false,
		},
		{
			name: "Happy Path - shorthand flags",
			args: []string{
				"-" + options.LicenseProductOption.Flag.Shorthand, customtypes.ENUM_LICENSE_PRODUCT_PING_FEDERATE,
				"-" + options.LicenseVersionOption.Flag.Shorthand, "12.0",
			},
			expectErr: false,
		},
		{
			name: "Happy Path - with profile flag",
			args: []string{
				"--" + options.LicenseProductOption.CobraParamName, customtypes.ENUM_LICENSE_PRODUCT_PING_FEDERATE,
				"--" + options.LicenseVersionOption.CobraParamName, "12.0",
				"--" + options.RootProfileOption.CobraParamName, "default",
			},
			expectErr: false,
		},
		{
			name:      "Happy Path - help",
			args:      []string{"--help"},
			expectErr: false,
		},
		{
			name:          "Too many arguments",
			args:          []string{"extra-arg"},
			expectErr:     true,
			expectedErrIs: common.ErrExactArgs,
		},
		{
			name:                "Invalid flag",
			args:                []string{"--invalid-flag"},
			expectErr:           true,
			expectedErrContains: "unknown flag",
		},
		{
			name: "Missing required product flag",
			args: []string{
				"--" + options.LicenseVersionOption.CobraParamName, "12.0",
			},
			expectErr:           true,
			expectedErrContains: fmt.Sprintf(`required flag(s) "%s" not set`, options.LicenseProductOption.CobraParamName),
		},
		{
			name: "Missing required version flag",
			args: []string{
				"--" + options.LicenseProductOption.CobraParamName, customtypes.ENUM_LICENSE_PRODUCT_PING_FEDERATE,
			},
			expectErr:           true,
			expectedErrContains: fmt.Sprintf(`required flag(s) "%s" not set`, options.LicenseVersionOption.CobraParamName),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			testutils_koanf.InitKoanfs(t)

			err := testutils_cobra.ExecutePingcli(t, append([]string{"license"}, tc.args...)...)

			if !tc.expectErr {
				require.NoError(t, err)
				return
			}

			assert.Error(t, err)
			if tc.expectedErrIs != nil {
				assert.ErrorIs(t, err, tc.expectedErrIs)
			}
			if tc.expectedErrContains != "" {
				assert.ErrorContains(t, err, tc.expectedErrContains)
			}
		})
	}
}
