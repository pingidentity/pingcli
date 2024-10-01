package cmd_test

import (
	"testing"

	"github.com/pingidentity/pingctl/internal/customtypes"
	"github.com/pingidentity/pingctl/internal/testing/testutils"
	"github.com/pingidentity/pingctl/internal/testing/testutils_cobra"
)

// Test Root Command Executes without issue
func TestRootCmd_Execute(t *testing.T) {
	err := testutils_cobra.ExecutePingctl(t)
	testutils.CheckExpectedError(t, err, nil)
}

// Test Root Command Executes fails when provided an invalid command
func TestRootCmd_InvalidCommand(t *testing.T) {
	expectedErrorPattern := `^unknown command "invalid" for "pingctl"$`
	err := testutils_cobra.ExecutePingctl(t, "invalid")
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}

// Test Root Command --help, -h flag
func TestRootCmd_HelpFlag(t *testing.T) {
	err := testutils_cobra.ExecutePingctl(t, "--help")
	testutils.CheckExpectedError(t, err, nil)

	err = testutils_cobra.ExecutePingctl(t, "-h")
	testutils.CheckExpectedError(t, err, nil)
}

// Test Root Command fails with invalid flag
func TestRootCmd_InvalidFlag(t *testing.T) {
	expectedErrorPattern := `^unknown flag: --invalid$`
	err := testutils_cobra.ExecutePingctl(t, "--invalid")
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}

// Test Root Command Executes when provided the --version, -v flag
func TestRootCmd_VersionFlag(t *testing.T) {
	err := testutils_cobra.ExecutePingctl(t, "--version")
	testutils.CheckExpectedError(t, err, nil)

	err = testutils_cobra.ExecutePingctl(t, "-v")
	testutils.CheckExpectedError(t, err, nil)
}

// Test Root Command Executes when provided the --output-format flag
func TestRootCmd_OutputFormatFlag(t *testing.T) {
	for _, outputFormat := range customtypes.OutputFormatValidValues() {
		err := testutils_cobra.ExecutePingctl(t, "--output-format", outputFormat)
		testutils.CheckExpectedError(t, err, nil)
	}
}

// Test Root Command fails when provided an invalid value for the --output-format flag
func TestRootCmd_InvalidOutputFlag(t *testing.T) {
	expectedErrorPattern := `^invalid argument "invalid" for "-O, --output-format" flag: unrecognized Output Format: 'invalid'\. Must be one of: [a-z\s,]+$`
	err := testutils_cobra.ExecutePingctl(t, "--output-format", "invalid")
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}

// Test Root Command fails when provided no value for the --output-format flag
func TestRootCmd_NoValueOutputFlag(t *testing.T) {
	expectedErrorPattern := `^flag needs an argument: --output-format$`
	err := testutils_cobra.ExecutePingctl(t, "--output-format")
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}

// Test Root Command Executes output does not change with output-format=text vs output-format=json
func TestRootCmd_OutputFlagTextVsJSON(t *testing.T) {
	textOutput, err := testutils_cobra.ExecutePingctlCaptureCobraOutput(t, "--output-format", "text")
	testutils.CheckExpectedError(t, err, nil)

	jsonOutput, err := testutils_cobra.ExecutePingctlCaptureCobraOutput(t, "--output-format", "json")
	testutils.CheckExpectedError(t, err, nil)

	if textOutput != jsonOutput {
		t.Errorf("Expected text and json output to be the same")
	}
}

// Test Root Command Executes when provided the --color flag
func TestRootCmd_ColorFlag(t *testing.T) {
	err := testutils_cobra.ExecutePingctl(t, "--color=true")
	testutils.CheckExpectedError(t, err, nil)

	err = testutils_cobra.ExecutePingctl(t, "--color=false")
	testutils.CheckExpectedError(t, err, nil)
}

// Test Root Command fails when provided an invalid value for the --color flag
func TestRootCmd_InvalidColorFlag(t *testing.T) {
	expectedErrorPattern := `^invalid argument "invalid" for "--color" flag: strconv\.ParseBool: parsing "invalid": invalid syntax$`
	err := testutils_cobra.ExecutePingctl(t, "--color=invalid")
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}

// Test Root Command Executes when provided the --config flag
func TestRootCmd_ConfigFlag(t *testing.T) {
	err := testutils_cobra.ExecutePingctl(t, "--config", "config.yaml")
	testutils.CheckExpectedError(t, err, nil)
}

// Test Root Command fails when provided no value for the --config flag
func TestRootCmd_NoValueConfigFlag(t *testing.T) {
	expectedErrorPattern := `^flag needs an argument: --config$`
	err := testutils_cobra.ExecutePingctl(t, "--config")
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}

// Test Root Command Executes when provided the --active-profile flag
func TestRootCmd_ActiveProfileFlag(t *testing.T) {
	err := testutils_cobra.ExecutePingctl(t, "--active-profile", "default")
	testutils.CheckExpectedError(t, err, nil)
}

// Test Root Command fails when provided no value for the --active-profile flag
func TestRootCmd_NoValueActiveProfileFlag(t *testing.T) {
	expectedErrorPattern := `^flag needs an argument: --active-profile$`
	err := testutils_cobra.ExecutePingctl(t, "--active-profile")
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}
