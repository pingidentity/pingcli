// Copyright © 2025 Ping Identity Corporation

package config_test

import (
	"fmt"
	"testing"

	"github.com/pingidentity/pingcli/internal/configuration/options"
	"github.com/pingidentity/pingcli/internal/customtypes"
	"github.com/pingidentity/pingcli/internal/profiles"
	"github.com/pingidentity/pingcli/internal/testing/testutils"
	"github.com/pingidentity/pingcli/internal/testing/testutils_cobra"
)

// Test Config Set Command Executes without issue
func TestConfigSetCmd_Execute(t *testing.T) {
	err := testutils_cobra.ExecutePingcli(t, "config", "set", fmt.Sprintf("%s=false", options.RootColorOption.KoanfKey))
	testutils.CheckExpectedError(t, err, nil)
}

// Test Config Set Command Fails when provided too few arguments
func TestConfigSetCmd_TooFewArgs(t *testing.T) {
	expectedErrorPattern := `^failed to execute 'pingcli config set': command accepts 1 arg\(s\), received 0$`
	err := testutils_cobra.ExecutePingcli(t, "config", "set")
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}

// Test Config Set Command Fails when provided too many arguments
func TestConfigSetCmd_TooManyArgs(t *testing.T) {
	expectedErrorPattern := `^failed to execute 'pingcli config set': command accepts 1 arg\(s\), received 2$`
	err := testutils_cobra.ExecutePingcli(t, "config", "set", fmt.Sprintf("%s=false", options.RootColorOption.KoanfKey), fmt.Sprintf("%s=true", options.RootColorOption.KoanfKey))
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}

// Test Config Set Command Fails when an invalid key is provided
func TestConfigSetCmd_InvalidKey(t *testing.T) {
	expectedErrorPattern := `^failed to set configuration: key 'pingcli\.invalid' is not recognized as a valid configuration key\.\s*Use 'pingcli config list-keys' to view all available keys`
	err := testutils_cobra.ExecutePingcli(t, "config", "set", "pingcli.invalid=true")
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}

// Test Config Set Command Fails when an invalid value type is provided
func TestConfigSetCmd_InvalidValueType(t *testing.T) {
	expectedErrorPattern := `^failed to set configuration: value for key '.*' must be a boolean\. Allowed .*: strconv\.ParseBool: parsing ".*": invalid syntax$`
	err := testutils_cobra.ExecutePingcli(t, "config", "set", fmt.Sprintf("%s=invalid", options.RootColorOption.KoanfKey))
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}

// Test Config Set Command Fails when no value is provided
func TestConfigSetCmd_NoValueProvided(t *testing.T) {
	expectedErrorPattern := `^failed to set configuration: value for key '.*' is empty\. Use 'pingcli config unset .*' to unset the key$`
	err := testutils_cobra.ExecutePingcli(t, "config", "set", fmt.Sprintf("%s=", options.RootColorOption.KoanfKey))
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}

// Test Config Set Command for key 'pingone.worker.clientId' updates koanf configuration
func TestConfigSetCmd_CheckKoanfConfig(t *testing.T) {
	koanfKey := options.PingOneAuthenticationWorkerClientIDOption.KoanfKey
	koanfNewUUID := "12345678-1234-1234-1234-123456789012"

	err := testutils_cobra.ExecutePingcli(t, "config", "set", fmt.Sprintf("%s=%s", koanfKey, koanfNewUUID))
	testutils.CheckExpectedError(t, err, nil)

	koanf := profiles.GetKoanfConfig().KoanfInstance()
	profileKoanfKey := "default." + koanfKey

	koanfNewValue, ok := koanf.Get(profileKoanfKey).(*customtypes.UUID)
	if ok && koanfNewValue.String() != koanfNewUUID {
		t.Errorf("Expected koanf configuration value to be updated")
	}
}

// Test Config Set Command --help, -h flag
func TestConfigSetCmd_HelpFlag(t *testing.T) {
	err := testutils_cobra.ExecutePingcli(t, "config", "set", "--help")
	testutils.CheckExpectedError(t, err, nil)

	err = testutils_cobra.ExecutePingcli(t, "config", "set", "-h")
	testutils.CheckExpectedError(t, err, nil)
}

// Test Config Set Command Fails when provided an invalid flag
func TestConfigSetCmd_InvalidFlag(t *testing.T) {
	expectedErrorPattern := `^unknown flag: --invalid$`
	err := testutils_cobra.ExecutePingcli(t, "config", "set", "--invalid")
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}

// https://pkg.go.dev/testing#hdr-Examples
func Example_setMaskedValue() {
	t := testing.T{}
	_ = testutils_cobra.ExecutePingcli(&t, "config", "set", fmt.Sprintf("%s=%s", options.PingFederateBasicAuthPasswordOption.KoanfKey, "1234"))

	// Output:
	// SUCCESS: Configuration set successfully:
	// service.pingFederate.authentication.basicAuth.password=********
}

// https://pkg.go.dev/testing#hdr-Examples
func Example_set_UnmaskedValuesFlag() {
	t := testing.T{}
	_ = testutils_cobra.ExecutePingcli(&t, "config", "set", "--unmask-values", fmt.Sprintf("%s=%s", options.PingFederateBasicAuthPasswordOption.KoanfKey, "1234"))

	// Output:
	// SUCCESS: Configuration set successfully:
	// service.pingFederate.authentication.basicAuth.password=1234
}

// https://pkg.go.dev/testing#hdr-Examples
func Example_setUnmaskedValue() {
	t := testing.T{}
	_ = testutils_cobra.ExecutePingcli(&t, "config", "set", fmt.Sprintf("%s=%s", options.RootColorOption.KoanfKey, "true"))

	// Output:
	// SUCCESS: Configuration set successfully:
	// noColor=true
}
