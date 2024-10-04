package config_test

import (
	"fmt"
	"testing"

	"github.com/pingidentity/pingcli/internal/configuration/options"
	"github.com/pingidentity/pingcli/internal/profiles"
	"github.com/pingidentity/pingcli/internal/testing/testutils"
	"github.com/pingidentity/pingcli/internal/testing/testutils_cobra"
)

// Test Config Set Command Executes without issue
func TestConfigSetCmd_Execute(t *testing.T) {
	err := testutils_cobra.ExecutePingcli(t, "config", "set", fmt.Sprintf("%s=false", options.RootColorOption.ViperKey))
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
	err := testutils_cobra.ExecutePingcli(t, "config", "set", fmt.Sprintf("%s=false", options.RootColorOption.ViperKey), fmt.Sprintf("%s=true", options.RootColorOption.ViperKey))
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}

// Test Config Set Command Fails when an invalid key is provided
func TestConfigSetCmd_InvalidKey(t *testing.T) {
	expectedErrorPattern := `^failed to set configuration: key 'pingcli\.invalid' is not recognized as a valid configuration key\. Valid keys: [A-Za-z\.\s,]+$`
	err := testutils_cobra.ExecutePingcli(t, "config", "set", "pingcli.invalid=true")
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}

// Test Config Set Command Fails when an invalid value type is provided
func TestConfigSetCmd_InvalidValueType(t *testing.T) {
	expectedErrorPattern := `^failed to set configuration: value for key '.*' must be a boolean\. Allowed .*: strconv\.ParseBool: parsing ".*": invalid syntax$`
	err := testutils_cobra.ExecutePingcli(t, "config", "set", fmt.Sprintf("%s=invalid", options.RootColorOption.ViperKey))
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}

// Test Config Set Command Fails when no value is provided
func TestConfigSetCmd_NoValueProvided(t *testing.T) {
	expectedErrorPattern := `^failed to set configuration: value for key '.*' is empty\. Use 'pingcli config unset .*' to unset the key$`
	err := testutils_cobra.ExecutePingcli(t, "config", "set", fmt.Sprintf("%s=", options.RootColorOption.ViperKey))
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}

// Test Config Set Command for key 'pingone.worker.clientId' updates viper configuration
func TestConfigSetCmd_CheckViperConfig(t *testing.T) {
	viperKey := options.PingoneAuthenticationWorkerClientIDOption.ViperKey
	viperNewUUID := "12345678-1234-1234-1234-123456789012"

	err := testutils_cobra.ExecutePingcli(t, "config", "set", fmt.Sprintf("%s=%s", viperKey, viperNewUUID))
	testutils.CheckExpectedError(t, err, nil)

	mainViper := profiles.GetMainConfig().ViperInstance()
	profileViperKey := profiles.GetMainConfig().ActiveProfile().Name() + "." + viperKey

	viperNewValue := mainViper.GetString(profileViperKey)
	if viperNewValue != viperNewUUID {
		t.Errorf("Expected viper configuration value to be updated")
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
