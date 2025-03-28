// Copyright © 2025 Ping Identity Corporation

package cmd_test

import (
	"testing"

	"github.com/pingidentity/pingcli/internal/configuration/options"
	"github.com/pingidentity/pingcli/internal/customtypes"
	"github.com/pingidentity/pingcli/internal/output"
	"github.com/pingidentity/pingcli/internal/testing/testutils"
	"github.com/pingidentity/pingcli/internal/testing/testutils_cobra"
	"github.com/pingidentity/pingcli/internal/testing/testutils_viper"
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
	err := testutils_cobra.ExecutePingcli(t, "--"+options.RootConfigOption.CobraParamName, "config.yaml")
	testutils.CheckExpectedError(t, err, nil)
}

// Test Root Command fails when provided no value for the --config flag
func TestRootCmd_NoValueConfigFlag(t *testing.T) {
	expectedErrorPattern := `^flag needs an argument: --config$`
	err := testutils_cobra.ExecutePingcli(t, "--"+options.RootConfigOption.CobraParamName)
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

// // Test Root Command Detailed Exit Code Flag
func TestRootCmd_DetailedExitCodeFlag(t *testing.T) {
	err := testutils_cobra.ExecutePingcli(t, "--"+options.RootDetailedExitCodeOption.CobraParamName)
	testutils.CheckExpectedError(t, err, nil)

	err = testutils_cobra.ExecutePingcli(t, "-"+options.RootDetailedExitCodeOption.Flag.Shorthand)
	testutils.CheckExpectedError(t, err, nil)
}

// // Test Root Command Detailed Exit Code Flag with output Warn
func TestRootCmd_DetailedExitCodeWarnLoggedFunc(t *testing.T) {
	testutils_viper.InitVipers(t)
	t.Setenv(options.RootDetailedExitCodeOption.EnvVar, "true")
	output.Warn("test warning", nil)

	warnLogged, err := output.DetailedExitCodeWarnLogged()
	testutils.CheckExpectedError(t, err, nil)
	if !warnLogged {
		t.Errorf("Expected DetailedExitCodeWarnLogged to return true")
	}
}
