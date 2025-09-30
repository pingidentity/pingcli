// Copyright © 2025 Ping Identity Corporation

package config_test

import (
	"testing"

	"github.com/pingidentity/pingcli/cmd/common"
	"github.com/pingidentity/pingcli/internal/configuration"
	"github.com/pingidentity/pingcli/internal/configuration/options"
	"github.com/pingidentity/pingcli/internal/profiles"
	"github.com/pingidentity/pingcli/internal/testing/testutils_cobra"
	"github.com/pingidentity/pingcli/internal/testing/testutils_koanf"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_ConfigUnsetCommand(t *testing.T) {
	testutils_koanf.InitKoanfs(t)

	testCases := []struct {
		name                string
		args                []string
		expectErr           bool
		expectedErrIs       error
		expectedErrContains string
	}{
		{
			name:      "Happy Path",
			args:      []string{options.RootColorOption.KoanfKey},
			expectErr: false,
		},
		{
			name:      "Happy Path - help",
			args:      []string{"--help"},
			expectErr: false,
		},
		{
			name:          "Too few arguments",
			args:          []string{},
			expectErr:     true,
			expectedErrIs: common.ErrExactArgs,
		},
		{
			name:          "Too many arguments",
			args:          []string{options.RootColorOption.KoanfKey, options.RootOutputFormatOption.KoanfKey},
			expectErr:     true,
			expectedErrIs: common.ErrExactArgs,
		},
		{
			name:          "Invalid key",
			args:          []string{"pingcli.invalid"},
			expectErr:     true,
			expectedErrIs: configuration.ErrInvalidConfigurationKey,
		},
		{
			name:                "Invalid flag",
			args:                []string{"--invalid-flag"},
			expectErr:           true,
			expectedErrContains: "unknown flag",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			testutils_koanf.InitKoanfs(t)

			err := testutils_cobra.ExecutePingcli(t, append([]string{"config", "unset"}, tc.args...)...)

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

// TestConfigUnsetCmd_CheckKoanfConfig verifies that the 'config unset' command correctly updates the underlying koanf configuration.
func TestConfigUnsetCmd_CheckKoanfConfig(t *testing.T) {
	testutils_koanf.InitKoanfs(t)

	koanfConfig, err := profiles.GetKoanfConfig()
	require.NoError(t, err, "Error getting koanf configuration")

	koanfInstance := koanfConfig.KoanfInstance()
	koanfKey := options.PingOneAuthenticationWorkerClientIDOption.KoanfKey
	profileKoanfKey := "default." + koanfKey

	// Ensure there is a value to unset
	require.NotEmpty(t, koanfInstance.String(profileKoanfKey), "Precondition failed: koanf value is already empty")

	// Execute the unset command
	err = testutils_cobra.ExecutePingcli(t, "config", "unset", koanfKey)
	require.NoError(t, err)

	// Re-fetch the koanf instance to see the change
	koanfConfig, err = profiles.GetKoanfConfig()
	require.NoError(t, err, "Error getting koanf configuration")
	koanfInstance = koanfConfig.KoanfInstance()
	koanfNewValue := koanfInstance.String(profileKoanfKey)
	assert.Empty(t, koanfNewValue, "Expected koanf configuration value to be empty after unset")
}

// https://pkg.go.dev/testing#hdr-Examples
func Example_unsetMaskedValue() {
	t := testing.T{}
	testutils_koanf.InitKoanfs(&t)
	_ = testutils_cobra.ExecutePingcli(&t, "config", "unset", options.PingFederateBasicAuthUsernameOption.KoanfKey)

	// Output:
	// SUCCESS: Configuration unset successfully:
	// service.pingFederate.authentication.basicAuth.username=
}

// https://pkg.go.dev/testing#hdr-Examples
func Example_unsetUnmaskedValue() {
	t := testing.T{}
	testutils_koanf.InitKoanfs(&t)
	_ = testutils_cobra.ExecutePingcli(&t, "config", "unset", options.RootOutputFormatOption.KoanfKey)

	// Output:
	// SUCCESS: Configuration unset successfully:
	// outputFormat=text
}
