package config_test

import (
	"testing"

	"github.com/pingidentity/pingcli/internal/testing/testutils"
	"github.com/pingidentity/pingcli/internal/testing/testutils_cobra"
)

// Test config add profile command executes without issue
func TestConfigAddProfileCmd_Execute(t *testing.T) {
	err := testutils_cobra.ExecutePingcli(t, "config", "add-profile",
		"--name", "test-profile",
		"--description", "test description",
		"--set-active=false")
	testutils.CheckExpectedError(t, err, nil)
}

// Test config add profile command fails when provided too many arguments
func TestConfigAddProfileCmd_TooManyArgs(t *testing.T) {
	expectedErrorPattern := `^failed to execute 'pingcli config add-profile': command accepts 0 arg\(s\), received 1$`
	err := testutils_cobra.ExecutePingcli(t, "config", "add-profile", "extra-arg")
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}

// Test config add profile command fails when provided an invalid flag
func TestConfigAddProfileCmd_InvalidFlag(t *testing.T) {
	expectedErrorPattern := `^unknown flag: --invalid-flag$`
	err := testutils_cobra.ExecutePingcli(t, "config", "add-profile", "--invalid-flag")
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}

// Test config add profile command fails when provided an invalid value for a flag
func TestConfigAddProfileCmd_InvalidFlagValue(t *testing.T) {
	expectedErrorPattern := `^invalid argument ".*" for ".*" flag: strconv\.ParseBool: parsing ".*": invalid syntax$`
	err := testutils_cobra.ExecutePingcli(t, "config", "add-profile", "--set-active", "invalid-value")
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}

// Test config add profile command fails when provided a duplicate profile name
func TestConfigAddProfileCmd_DuplicateProfileName(t *testing.T) {
	expectedErrorPattern := `^failed to add profile: invalid profile name: '.*'\. profile already exists$`
	err := testutils_cobra.ExecutePingcli(t, "config", "add-profile",
		"--name", "default",
		"--description", "test description",
		"--set-active=false")
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}

// Test config add profile command fails when provided an invalid profile name
func TestConfigAddProfileCmd_InvalidProfileName(t *testing.T) {
	expectedErrorPattern := `^failed to add profile: invalid profile name: '.*'\. name must contain only alphanumeric characters, underscores, and dashes$`
	err := testutils_cobra.ExecutePingcli(t, "config", "add-profile",
		"--name", "pname&*^*&^$&@!",
		"--description", "test description",
		"--set-active=false")
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}

// Test config add profile command fails when provided an invalid set-active value
func TestConfigAddProfileCmd_InvalidSetActiveValue(t *testing.T) {
	expectedErrorPattern := `^invalid argument ".*" for "-s, --set-active" flag: strconv\.ParseBool: parsing ".*": invalid syntax$`
	err := testutils_cobra.ExecutePingcli(t, "config", "add-profile",
		"--name", "test-profile",
		"--description", "test description",
		"--set-active", "invalid-value")
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}
