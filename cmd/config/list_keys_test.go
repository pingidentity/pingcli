package config_test

import (
	"testing"

	"github.com/pingidentity/pingcli/internal/configuration/options"
	"github.com/pingidentity/pingcli/internal/testing/testutils"
	"github.com/pingidentity/pingcli/internal/testing/testutils_cobra"
)

// Test Config List Keys Command Executes without issue
func TestConfigListKeysCmd_Execute(t *testing.T) {
	err := testutils_cobra.ExecutePingcli(t, "config", "list-keys")
	testutils.CheckExpectedError(t, err, nil)
}

// Test Config List Keys YAML Command --help, -h flag
func TestConfigListKeysCmd_YAMLFlag(t *testing.T) {
	err := testutils_cobra.ExecutePingcli(t, "config", "list-keys", "--yaml")
	testutils.CheckExpectedError(t, err, nil)

	err = testutils_cobra.ExecutePingcli(t, "config", "list-keys", "-y")
	testutils.CheckExpectedError(t, err, nil)
}

// Test Config List Keys Command --help, -h flag
func TestConfigListKeysCmd_HelpFlag(t *testing.T) {
	err := testutils_cobra.ExecutePingcli(t, "config", "list-keys", "--help")
	testutils.CheckExpectedError(t, err, nil)

	err = testutils_cobra.ExecutePingcli(t, "config", "list-keys", "-h")
	testutils.CheckExpectedError(t, err, nil)
}

// Test Config List Keys Command fails when provided too many arguments
func TestConfigListKeysCmd_TooManyArgs(t *testing.T) {
	expectedErrorPattern := `^failed to execute 'pingcli config list-keys': command accepts 0 arg\(s\), received 1$`
	err := testutils_cobra.ExecutePingcli(t, "config", "list-keys", options.RootColorOption.ViperKey)
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}
