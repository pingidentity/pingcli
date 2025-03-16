package config_test

import (
	"testing"

	"github.com/pingidentity/pingcli/internal/configuration/options"
	"github.com/pingidentity/pingcli/internal/testing/testutils"
	"github.com/pingidentity/pingcli/internal/testing/testutils_cobra"
)

// Test Config Get Command Executes without issue
func TestConfigGetCmd_Execute(t *testing.T) {
	err := testutils_cobra.ExecutePingcli(t, "config", "get", "export")
	testutils.CheckExpectedError(t, err, nil)
}

// Test Config Get Command fails when provided too many arguments
func TestConfigGetCmd_TooManyArgs(t *testing.T) {
	expectedErrorPattern := `^failed to execute 'pingcli config get': command accepts 1 arg\(s\), received 2$`
	err := testutils_cobra.ExecutePingcli(t, "config", "get", options.RootColorOption.ViperKey, options.RootOutputFormatOption.ViperKey)
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}

// Test Config Get Command Executes when provided a full key
func TestConfigGetCmd_FullKey(t *testing.T) {
	err := testutils_cobra.ExecutePingcli(t, "config", "get", options.PingOneAuthenticationWorkerClientIDOption.ViperKey)
	testutils.CheckExpectedError(t, err, nil)
}

// Test Config Get Command Executes when provided a partial key
func TestConfigGetCmd_PartialKey(t *testing.T) {
	err := testutils_cobra.ExecutePingcli(t, "config", "get", "service.pingone")
	testutils.CheckExpectedError(t, err, nil)
}

// Test Config Get Command fails when provided an invalid key
func TestConfigGetCmd_InvalidKey(t *testing.T) {
	expectedErrorPattern := `^failed to get configuration: key '.*' is not recognized as a valid configuration key.\s*Use 'pingcli config list-keys' to view all available keys`
	err := testutils_cobra.ExecutePingcli(t, "config", "get", "pingcli.invalid")
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}

// Test Config Get Command fails when provided an invalid flag
func TestConfigGetCmd_InvalidFlag(t *testing.T) {
	expectedErrorPattern := `^unknown flag: --invalid$`
	err := testutils_cobra.ExecutePingcli(t, "config", "get", "--invalid")
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}

// Test Config Get Command --help, -h flag
func TestConfigGetCmd_HelpFlag(t *testing.T) {
	err := testutils_cobra.ExecutePingcli(t, "config", "get", "--help")
	testutils.CheckExpectedError(t, err, nil)

	err = testutils_cobra.ExecutePingcli(t, "config", "get", "-h")
	testutils.CheckExpectedError(t, err, nil)
}

// Test Config Get Command fails when provided no key
func TestConfigGetCmd_NoKey(t *testing.T) {
	expectedErrorPattern := `^failed to execute '.*': command accepts 1 arg\(s\), received 0$`
	err := testutils_cobra.ExecutePingcli(t, "config", "get")
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}

// https://pkg.go.dev/testing#hdr-Examples
func Example_getEmptyMaskedValue() {
	t := testing.T{}
	_ = testutils_cobra.ExecutePingcli(&t, "config", "get", options.RequestAccessTokenOption.ViperKey)

	// Output:
	// Configuration values for profile 'default' and key 'request.accessToken':
	// request.accessToken=
	// request.accessTokenExpiry=0
}

// https://pkg.go.dev/testing#hdr-Examples
func Example_getMaskedValue() {
	t := testing.T{}
	_ = testutils_cobra.ExecutePingcli(&t, "config", "get", options.PingFederateBasicAuthPasswordOption.ViperKey)

	// Output:
	// Configuration values for profile 'default' and key 'service.pingfederate.authentication.basicAuth.password':
	// service.pingfederate.authentication.basicAuth.password=********
}

// https://pkg.go.dev/testing#hdr-Examples
func Example_getUnmaskedValue() {
	t := testing.T{}
	_ = testutils_cobra.ExecutePingcli(&t, "config", "get", options.RootColorOption.ViperKey)

	// Output:
	// Configuration values for profile 'default' and key 'noColor':
	// noColor=true
}

// https://pkg.go.dev/testing#hdr-Examples
func Example_get_UnmaskValuesFlag() {
	t := testing.T{}
	_ = testutils_cobra.ExecutePingcli(&t, "config", "get", "--unmask-values", options.RequestAccessTokenOption.ViperKey)

	// Output:
	// Configuration values for profile 'default' and key 'request.accessToken':
	// request.accessToken=
	// request.accessTokenExpiry=0
}
