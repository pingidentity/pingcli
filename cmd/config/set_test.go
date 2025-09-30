// Copyright © 2025 Ping Identity Corporation

package config_test

import (
	"fmt"
	"testing"

	"github.com/pingidentity/pingcli/cmd/common"
	config_internal "github.com/pingidentity/pingcli/internal/commands/config"
	"github.com/pingidentity/pingcli/internal/configuration"
	"github.com/pingidentity/pingcli/internal/configuration/options"
	"github.com/pingidentity/pingcli/internal/customtypes"
	"github.com/pingidentity/pingcli/internal/profiles"
	"github.com/pingidentity/pingcli/internal/testing/testutils_cobra"
	"github.com/pingidentity/pingcli/internal/testing/testutils_koanf"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_ConfigSetCommand(t *testing.T) {
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
			args:      []string{fmt.Sprintf("%s=false", options.RootColorOption.KoanfKey)},
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
			args:          []string{fmt.Sprintf("%s=false", options.RootColorOption.KoanfKey), fmt.Sprintf("%s=true", options.RootColorOption.KoanfKey)},
			expectErr:     true,
			expectedErrIs: common.ErrExactArgs,
		},
		{
			name:          "Invalid key",
			args:          []string{"pingcli.invalid=true"},
			expectErr:     true,
			expectedErrIs: configuration.ErrInvalidConfigurationKey,
		},
		{
			name:          "Invalid value type for key",
			args:          []string{fmt.Sprintf("%s=invalid", options.RootColorOption.KoanfKey)},
			expectErr:     true,
			expectedErrIs: customtypes.ErrParseBool,
		},
		{
			name:          "No value provided",
			args:          []string{fmt.Sprintf("%s=", options.RootColorOption.KoanfKey)},
			expectErr:     true,
			expectedErrIs: config_internal.ErrEmptyValue,
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

			err := testutils_cobra.ExecutePingcli(t, append([]string{"config", "set"}, tc.args...)...)

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

// TestConfigSetCmd_CheckKoanfConfig verifies that the 'config set' command correctly updates the underlying koanf configuration.
func TestConfigSetCmd_CheckKoanfConfig(t *testing.T) {
	testutils_koanf.InitKoanfs(t)

	koanfKey := options.PingOneAuthenticationWorkerClientIDOption.KoanfKey
	koanfNewUUID := "12345678-1234-1234-1234-123456789012"

	err := testutils_cobra.ExecutePingcli(t, "config", "set", fmt.Sprintf("%s=%s", koanfKey, koanfNewUUID))
	require.NoError(t, err)

	koanfConfig, err := profiles.GetKoanfConfig()
	require.NoError(t, err, "Error getting koanf configuration")

	koanfInstance := koanfConfig.KoanfInstance()
	profileKoanfKey := "default." + koanfKey

	koanfNewValue, ok := koanfInstance.Get(profileKoanfKey).(*customtypes.UUID)
	require.True(t, ok, "Koanf value is not of the expected type *customtypes.UUID")
	assert.Equal(t, koanfNewUUID, koanfNewValue.String(), "Expected koanf configuration value to be updated")
}

// https://pkg.go.dev/testing#hdr-Examples
func Example_setMaskedValue() {
	t := testing.T{}
	testutils_koanf.InitKoanfs(&t)
	_ = testutils_cobra.ExecutePingcli(&t, "config", "set", fmt.Sprintf("%s=%s", options.PingFederateBasicAuthPasswordOption.KoanfKey, "1234"))

	// Output:
	// SUCCESS: Configuration set successfully:
	// service.pingFederate.authentication.basicAuth.password=********
}

// https://pkg.go.dev/testing#hdr-Examples
func Example_set_UnmaskedValuesFlag() {
	t := testing.T{}
	testutils_koanf.InitKoanfs(&t)
	_ = testutils_cobra.ExecutePingcli(&t, "config", "set", "--"+options.ConfigUnmaskSecretValueOption.CobraParamName, fmt.Sprintf("%s=%s", options.PingFederateBasicAuthPasswordOption.KoanfKey, "1234"))

	// Output:
	// SUCCESS: Configuration set successfully:
	// service.pingFederate.authentication.basicAuth.password=1234
}

// https://pkg.go.dev/testing#hdr-Examples
func Example_setUnmaskedValue() {
	t := testing.T{}
	testutils_koanf.InitKoanfs(&t)
	_ = testutils_cobra.ExecutePingcli(&t, "config", "set", fmt.Sprintf("%s=%s", options.RootColorOption.KoanfKey, "true"))

	// Output:
	// SUCCESS: Configuration set successfully:
	// noColor=true
}
