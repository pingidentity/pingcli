package config_test

import (
	"os"
	"testing"

	"github.com/pingidentity/pingcli/internal/configuration/options"
	"github.com/pingidentity/pingcli/internal/profiles"
	"github.com/pingidentity/pingcli/internal/testing/testutils"
	"github.com/pingidentity/pingcli/internal/testing/testutils_cobra"
)

// Test Config Unset Command Executes without issue
func TestConfigUnsetCmd_Execute(t *testing.T) {
	err := testutils_cobra.ExecutePingcli(t, "config", "unset", options.RootColorOption.ViperKey)
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
	err := testutils_cobra.ExecutePingcli(t, "config", "unset", options.RootColorOption.ViperKey, options.RootOutputFormatOption.ViperKey)
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}

// Test Config Unset Command Fails when an invalid key is provided
func TestConfigUnsetCmd_InvalidKey(t *testing.T) {
	expectedErrorPattern := `^failed to unset configuration: key '.*' is not recognized as a valid configuration key\. Valid keys: [A-Za-z\.\s,]+$`
	err := testutils_cobra.ExecutePingcli(t, "config", "unset", "pingcli.invalid")
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}

// Test Config Unset Command for key 'pingone.worker.clientId' updates viper configuration
func TestConfigUnsetCmd_CheckViperConfig(t *testing.T) {
	viperKey := options.PingoneAuthenticationWorkerClientIDOption.ViperKey
	viperOldValue := os.Getenv(options.PingoneAuthenticationWorkerClientIDOption.EnvVar)

	err := testutils_cobra.ExecutePingcli(t, "config", "unset", viperKey)
	testutils.CheckExpectedError(t, err, nil)

	mainViper := profiles.GetMainConfig().ViperInstance()
	profileViperKey := profiles.GetMainConfig().ActiveProfile().Name() + "." + viperKey
	viperNewValue := mainViper.GetString(profileViperKey)
	if viperOldValue == viperNewValue {
		t.Errorf("Expected viper configuration value to be updated. Old: %s, New: %s", viperOldValue, viperNewValue)
	}

	if viperNewValue != "" {
		t.Errorf("Expected viper configuration value to be empty. Got: %s", viperNewValue)
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
