// Copyright © 2025 Ping Identity Corporation

package cmd_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/pingidentity/pingcli/cmd"
	"github.com/pingidentity/pingcli/internal/configuration/options"
	"github.com/pingidentity/pingcli/internal/customtypes"
	"github.com/pingidentity/pingcli/internal/output"
	"github.com/pingidentity/pingcli/internal/testing/testutils"
	"github.com/pingidentity/pingcli/internal/testing/testutils_cobra"
	"github.com/pingidentity/pingcli/internal/testing/testutils_koanf"
)

// Test Root Command Executes without issue
func TestRootCmd_Execute(t *testing.T) {
	err := testutils_cobra.ExecutePingcli(t)
	testutils.CheckExpectedError(t, err, nil)
}

// Test Root Command Executes fails when provided an invalid command
func TestRootCmd_InvalidCommand(t *testing.T) {
	expectedErrorPattern := `^unknown command "invalid" for "pingcli"$`
	err := testutils_cobra.ExecutePingcli(t, "invalid")
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}

// Test Root Command --help, -h flag
func TestRootCmd_HelpFlag(t *testing.T) {
	err := testutils_cobra.ExecutePingcli(t, "--help")
	testutils.CheckExpectedError(t, err, nil)

	err = testutils_cobra.ExecutePingcli(t, "-h")
	testutils.CheckExpectedError(t, err, nil)
}

// Test Root Command fails with invalid flag
func TestRootCmd_InvalidFlag(t *testing.T) {
	expectedErrorPattern := `^unknown flag: --invalid$`
	err := testutils_cobra.ExecutePingcli(t, "--invalid")
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}

// Test Root Command Executes when provided the --version, -v flag
func TestRootCmd_VersionFlag(t *testing.T) {
	err := testutils_cobra.ExecutePingcli(t, "--version")
	testutils.CheckExpectedError(t, err, nil)

	err = testutils_cobra.ExecutePingcli(t, "-v")
	testutils.CheckExpectedError(t, err, nil)
}

// Test Root Command Executes when provided the --output-format flag
func TestRootCmd_OutputFormatFlag(t *testing.T) {
	for _, outputFormat := range customtypes.OutputFormatValidValues() {
		err := testutils_cobra.ExecutePingcli(t, "--"+options.RootOutputFormatOption.CobraParamName, outputFormat)
		testutils.CheckExpectedError(t, err, nil)
	}
}

// Test Root Command fails when provided an invalid value for the --output-format flag
func TestRootCmd_InvalidOutputFlag(t *testing.T) {
	expectedErrorPattern := `^invalid argument "invalid" for "-O, --output-format" flag: unrecognized Output Format: 'invalid'\. Must be one of: [a-z\s,]+$`
	err := testutils_cobra.ExecutePingcli(t, "--"+options.RootOutputFormatOption.CobraParamName, "invalid")
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}

// Test Root Command fails when provided no value for the --output-format flag
func TestRootCmd_NoValueOutputFlag(t *testing.T) {
	expectedErrorPattern := `^flag needs an argument: --output-format$`
	err := testutils_cobra.ExecutePingcli(t, "--"+options.RootOutputFormatOption.CobraParamName)
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}

// Test Root Command Executes output does not change with output-format=text vs output-format=json
func TestRootCmd_OutputFlagTextVsJSON(t *testing.T) {
	textOutput, err := testutils_cobra.ExecutePingcliCaptureCobraOutput(t, "--"+options.RootOutputFormatOption.CobraParamName, "text")
	testutils.CheckExpectedError(t, err, nil)

	jsonOutput, err := testutils_cobra.ExecutePingcliCaptureCobraOutput(t, "--"+options.RootOutputFormatOption.CobraParamName, "json")
	testutils.CheckExpectedError(t, err, nil)

	if textOutput != jsonOutput {
		t.Errorf("Expected text and json output to be the same")
	}
}

// Test Root Command Executes when provided the --no-color flag
func TestRootCmd_ColorFlag(t *testing.T) {
	err := testutils_cobra.ExecutePingcli(t, "--"+options.RootColorOption.CobraParamName)
	testutils.CheckExpectedError(t, err, nil)

	err = testutils_cobra.ExecutePingcli(t, "--"+options.RootColorOption.CobraParamName+"=false")
	testutils.CheckExpectedError(t, err, nil)
}

// Test Root Command fails when provided an invalid value for the --no-color flag
func TestRootCmd_InvalidColorFlag(t *testing.T) {
	expectedErrorPattern := `^invalid argument "invalid" for ".*" flag: strconv\.ParseBool: parsing "invalid": invalid syntax$`
	err := testutils_cobra.ExecutePingcli(t, "--"+options.RootColorOption.CobraParamName+"=invalid")
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}

// Test Root Command Executes when provided the --config flag
func TestRootCmd_ConfigFlag(t *testing.T) {
	// Add the --config args to os.Args
	os.Args = append(os.Args, "--"+options.RootConfigOption.CobraParamName, "config.yaml")

	err := testutils_cobra.ExecutePingcli(t, "--"+options.RootConfigOption.CobraParamName, "config.yaml")
	testutils.CheckExpectedError(t, err, nil)
}

// Test Root Command fails when provided no value for the --config flag
func TestRootCmd_NoValueConfigFlag(t *testing.T) {
	expectedErrorPattern := `^flag needs an argument: --config$`
	err := testutils_cobra.ExecutePingcli(t, "--"+options.RootConfigOption.CobraParamName)
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}

// Test Root Command fails on non-existent configuration file
func TestRootCmd_NonExistentConfigFile(t *testing.T) {
	expectedErrorPattern := `^Configuration file '.*' does not exist. Use the default configuration file location or specify a valid configuration file location with the --config flag\.$`
	err := testutils_cobra.ExecutePingcli(t, "--"+options.RootConfigOption.CobraParamName, "non_existent.yaml")
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}

// Test Root Command Executes when provided the --profile flag
func TestRootCmd_ProfileFlag(t *testing.T) {
	err := testutils_cobra.ExecutePingcli(t, "--"+options.RootProfileOption.CobraParamName, "default")
	testutils.CheckExpectedError(t, err, nil)
}

// Test Root Command fails when provided no value for the --profile flag
func TestRootCmd_NoValueProfileFlag(t *testing.T) {
	expectedErrorPattern := `^flag needs an argument: --profile$`
	err := testutils_cobra.ExecutePingcli(t, "--"+options.RootProfileOption.CobraParamName)
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}

// Test Root Command Detailed Exit Code Flag
func TestRootCmd_DetailedExitCodeFlag(t *testing.T) {
	err := testutils_cobra.ExecutePingcli(t, "--"+options.RootDetailedExitCodeOption.CobraParamName)
	testutils.CheckExpectedError(t, err, nil)

	err = testutils_cobra.ExecutePingcli(t, "-"+options.RootDetailedExitCodeOption.Flag.Shorthand)
	testutils.CheckExpectedError(t, err, nil)
}

// Test Root Command Detailed Exit Code Flag with output Warn
func TestRootCmd_DetailedExitCodeWarnLoggedFunc(t *testing.T) {
	testutils_koanf.InitKoanfs(t)
	t.Setenv(options.RootDetailedExitCodeOption.EnvVar, "true")
	output.Warn("test warning", nil)

	warnLogged, err := output.DetailedExitCodeWarnLogged()
	testutils.CheckExpectedError(t, err, nil)
	if !warnLogged {
		t.Errorf("Expected DetailedExitCodeWarnLogged to return true")
	}
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
			if result != tc.expected {
				t.Errorf("expected %s, got %s", tc.expected, result)
			}
		})
	}
}
