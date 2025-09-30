// Copyright © 2025 Ping Identity Corporation

package cmd_test

import (
	"fmt"
	"testing"

	"github.com/pingidentity/pingcli/cmd"
	"github.com/pingidentity/pingcli/internal/configuration/options"
	"github.com/pingidentity/pingcli/internal/customtypes"
	"github.com/pingidentity/pingcli/internal/output"
	"github.com/pingidentity/pingcli/internal/testing/testutils_cobra"
	"github.com/pingidentity/pingcli/internal/testing/testutils_koanf"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_RootCommand_Validation(t *testing.T) {
	testutils_koanf.InitKoanfs(t)

	testCases := []struct {
		name                string
		args                []string
		expectErr           bool
		expectedErrIs       error
		expectedErrContains string
	}{
		// Basic Command Structure
		{
			name:      "Happy Path - no args",
			args:      []string{},
			expectErr: false,
		},
		{
			name:      "Happy Path - help",
			args:      []string{"--help"},
			expectErr: false,
		},
		{
			name:      "Happy Path - version",
			args:      []string{"--version"},
			expectErr: false,
		},
		{
			name:                "Invalid command",
			args:                []string{"invalid-command"},
			expectErr:           true,
			expectedErrContains: "unknown command \"invalid-command\" for \"pingcli\"",
		},
		{
			name:                "Invalid flag",
			args:                []string{"--invalid-flag"},
			expectErr:           true,
			expectedErrContains: "unknown flag: --invalid-flag",
		},
		{
			name:      "Happy Path - output-format flag",
			args:      []string{"--" + options.RootOutputFormatOption.CobraParamName, "json"},
			expectErr: false,
		},
		{
			name:          "Invalid output-format",
			args:          []string{"--" + options.RootOutputFormatOption.CobraParamName, "invalid"},
			expectErr:     true,
			expectedErrIs: customtypes.ErrUnrecognizedOutputFormat,
		},
		{
			name:                "No value for output-format",
			args:                []string{"--" + options.RootOutputFormatOption.CobraParamName},
			expectErr:           true,
			expectedErrContains: "flag needs an argument: --" + options.RootOutputFormatOption.CobraParamName,
		},
		{
			name:      "Happy Path - no-color flag",
			args:      []string{"--" + options.RootColorOption.CobraParamName},
			expectErr: false,
		},
		{
			name:          "Invalid no-color value",
			args:          []string{"--" + options.RootColorOption.CobraParamName + "=invalid"},
			expectErr:     true,
			expectedErrIs: customtypes.ErrParseBool,
		},
		{
			name:      "Happy Path - config flag",
			args:      []string{"--" + options.RootConfigOption.CobraParamName, "config.yaml"},
			expectErr: false,
		},
		{
			name:                "No value for config",
			args:                []string{"--" + options.RootConfigOption.CobraParamName},
			expectErr:           true,
			expectedErrContains: "flag needs an argument: --" + options.RootConfigOption.CobraParamName,
		},
		{
			name:      "Happy Path - profile flag",
			args:      []string{"--" + options.RootProfileOption.CobraParamName, "default"},
			expectErr: false,
		},
		{
			name:                "No value for profile",
			args:                []string{"--" + options.RootProfileOption.CobraParamName},
			expectErr:           true,
			expectedErrContains: "flag needs an argument: --" + options.RootProfileOption.CobraParamName,
		},
		{
			name:      "Happy Path - detailed-exit-code flag",
			args:      []string{"--" + options.RootDetailedExitCodeOption.CobraParamName},
			expectErr: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			testutils_koanf.InitKoanfs(t)

			err := testutils_cobra.ExecutePingcli(t, tc.args...)

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

func Test_RootCommand_OutputComparison(t *testing.T) {
	textOutput, err := testutils_cobra.ExecutePingcliCaptureCobraOutput(t, "--output-format", "text")
	require.NoError(t, err)

	jsonOutput, err := testutils_cobra.ExecutePingcliCaptureCobraOutput(t, "--output-format", "json")
	require.NoError(t, err)

	assert.Equal(t, textOutput, jsonOutput, "Expected text and json output to be the same for the root command")
}

func Test_DetailedExitCodeWarnLogged(t *testing.T) {
	testutils_koanf.InitKoanfs(t)
	t.Setenv(options.RootDetailedExitCodeOption.EnvVar, "true")

	output.Warn("test warning", nil)

	warnLogged, err := output.DetailedExitCodeWarnLogged()
	require.NoError(t, err)
	assert.True(t, warnLogged, "Expected DetailedExitCodeWarnLogged to return true")
}

func TestParseArgsForConfigFile(t *testing.T) {
	testutils_koanf.InitKoanfs(t)
	defaultCfgFile := options.RootConfigOption.DefaultValue.String()

	testCases := []struct {
		name     string
		args     []string
		envVar   string
		expected string
	}{
		{
			name:     "no flags or env var",
			args:     []string{"pingcli"},
			expected: defaultCfgFile,
		},
		{
			name:     "config flag with equals",
			args:     []string{"pingcli", fmt.Sprintf("--%s=test.yaml", options.RootConfigOption.CobraParamName)},
			expected: "test.yaml",
		},
		{
			name:     "config flag with space",
			args:     []string{"pingcli", fmt.Sprintf("--%s", options.RootConfigOption.CobraParamName), "test2.yaml"},
			expected: "test2.yaml",
		},
		{
			name:     "short config flag with equals",
			args:     []string{"pingcli", fmt.Sprintf("-%s=test3.yaml", options.RootConfigOption.Flag.Shorthand)},
			expected: "test3.yaml",
		},
		{
			name:     "short config flag with space",
			args:     []string{"pingcli", fmt.Sprintf("-%s", options.RootConfigOption.Flag.Shorthand), "test4.yaml"},
			expected: "test4.yaml",
		},
		{
			name:     "env var",
			args:     []string{"pingcli"},
			envVar:   "test5.yaml",
			expected: "test5.yaml",
		},
		{
			name:     "flag overrides env var",
			args:     []string{"pingcli", fmt.Sprintf("--%s", options.RootConfigOption.CobraParamName), "flag.yaml"},
			envVar:   "env.yaml",
			expected: "flag.yaml",
		},
		{
			name:     "invalid format defaults to default",
			args:     []string{"pingcli", fmt.Sprintf("--%s::test.yaml", options.RootConfigOption.CobraParamName)},
			expected: defaultCfgFile,
		},
		{
			name:     "invalid flag name is ignored",
			args:     []string{"pingcli", "--confi=test.yaml"},
			expected: defaultCfgFile,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.envVar != "" {
				t.Setenv(options.RootConfigOption.EnvVar, tc.envVar)
			}

			result := cmd.ParseArgsForConfigFile(tc.args)
			assert.Equal(t, tc.expected, result)
		})
	}
}
