// Copyright © 2025 Ping Identity Corporation

package config_test

import (
	"testing"

	"github.com/pingidentity/pingcli/internal/configuration/options"
	"github.com/pingidentity/pingcli/internal/profiles"
	"github.com/pingidentity/pingcli/internal/testing/testutils"
	"github.com/pingidentity/pingcli/internal/testing/testutils_cobra"
	"github.com/pingidentity/pingcli/internal/testing/testutils_koanf"
)

// Test Config Unset Command Executes without issue
func TestConfigUnsetCmd_Execute(t *testing.T) {
	err := testutils_cobra.ExecutePingcli(t, "config", "unset", options.RootColorOption.KoanfKey)
	testutils.CheckExpectedError(t, err, nil)
}

// Test Config Set Command Fails when provided too few arguments
func TestConfigUnsetCmd_TooFewArgs(t *testing.T) {
	expectedErrorPattern := `^failed to execute 'pingcli config unset': command accepts 1 arg\(s\), received 0$`
	err := testutils_cobra.ExecutePingcli(t, "config", "unset")
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}

// Test Config Set Command Fails when provided too many arguments
func TestConfigUnsetCmd_TooManyArgs(t *testing.T) {
	expectedErrorPattern := `^failed to execute 'pingcli config unset': command accepts 1 arg\(s\), received 2$`
	err := testutils_cobra.ExecutePingcli(t, "config", "unset", options.RootColorOption.KoanfKey, options.RootOutputFormatOption.KoanfKey)
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}

// Test Config Unset Command Fails when an invalid key is provided
func TestConfigUnsetCmd_InvalidKey(t *testing.T) {
	expectedErrorPattern := `^failed to unset configuration: key '.*' is not recognized as a valid configuration key\.\s*Use 'pingcli config list-keys' to view all available keys`
	err := testutils_cobra.ExecutePingcli(t, "config", "unset", "pingcli.invalid")
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}

// Test Config Unset Command for key 'pingone.worker.clientId' updates koanf configuration
func TestConfigUnsetCmd_CheckKoanfConfig(t *testing.T) {
	testutils_koanf.InitKoanfs(t)

	koanfConfig := profiles.GetKoanfConfig().KoanfInstance()
	koanfKey := options.PingOneAuthenticationWorkerClientIDOption.KoanfKey
	profileKoanfKey := "default." + koanfKey

	koanfOldValue := koanfConfig.String(profileKoanfKey)

	err := testutils_cobra.ExecutePingcli(t, "config", "unset", koanfKey)
	testutils.CheckExpectedError(t, err, nil)

	koanfNewValue := koanfConfig.String(profileKoanfKey)
	if koanfOldValue == koanfNewValue {
		t.Errorf("Expected koanf configuration value to be updated. Old: %s, New: %s", koanfOldValue, koanfNewValue)
	}

	if koanfNewValue != "" {
		t.Errorf("Expected koanf configuration value to be empty. Got: %s", koanfNewValue)
	}
}

// Test Config Unset Command Fails when provided an invalid flag
func TestConfigUnsetCmd_InvalidFlag(t *testing.T) {
	expectedErrorPattern := `^unknown flag: --invalid$`
	err := testutils_cobra.ExecutePingcli(t, "config", "unset", "--invalid")
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}

// Test Config Unset Command --help, -h flag
func TestConfigUnsetCmd_HelpFlag(t *testing.T) {
	err := testutils_cobra.ExecutePingcli(t, "config", "unset", "--help")
	testutils.CheckExpectedError(t, err, nil)

	err = testutils_cobra.ExecutePingcli(t, "config", "unset", "-h")
	testutils.CheckExpectedError(t, err, nil)
}

// https://pkg.go.dev/testing#hdr-Examples
func Example_unsetMaskedValue() {
	t := testing.T{}
	_ = testutils_cobra.ExecutePingcli(&t, "config", "unset", options.PingFederateBasicAuthUsernameOption.KoanfKey)

	// Output:
	// SUCCESS: Configuration unset successfully:
	// service.pingFederate.authentication.basicAuth.username=
}

// https://pkg.go.dev/testing#hdr-Examples
func Example_unsetUnmaskedValue() {
	t := testing.T{}
	_ = testutils_cobra.ExecutePingcli(&t, "config", "unset", options.RootOutputFormatOption.KoanfKey)

	// Output:
	// SUCCESS: Configuration unset successfully:
	// outputFormat=text
}
