package config_test

import (
	"testing"

	"github.com/pingidentity/pingctl/internal/testing/testutils"
	"github.com/pingidentity/pingctl/internal/testing/testutils_cobra"
)

// Test Config Command Executes without issue
func TestConfigCmd_Execute(t *testing.T) {
	err := testutils_cobra.ExecutePingctl(t, "config",
		"--profile", "production",
		"--name", "myProfile",
		"--description", "hello")

	testutils.CheckExpectedError(t, err, nil)
}

// Test Config Command fails when provided invalid flag
func TestConfigCmd_InvalidFlag(t *testing.T) {
	expectedErrorPattern := `^unknown flag: --invalid$`
	err := testutils_cobra.ExecutePingctl(t, "config", "--invalid")
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}

// Test Config Command --help, -h flag
func TestConfigCmd_HelpFlag(t *testing.T) {
	err := testutils_cobra.ExecutePingctl(t, "config", "--help")
	testutils.CheckExpectedError(t, err, nil)

	err = testutils_cobra.ExecutePingctl(t, "config", "-h")
	testutils.CheckExpectedError(t, err, nil)
}

// Test Config Command fails when provided a profile name that does not exist
func TestConfigCmd_ProfileDoesNotExist(t *testing.T) {
	expectedErrorPattern := `^failed to update profile '.*' name to: .*\. invalid profile name: '.*' profile does not exist$`
	err := testutils_cobra.ExecutePingctl(t, "config",
		"--profile", "nonexistent",
		"--name", "myProfile",
		"--description", "hello")

	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}

// Test Config Command fails when attempting to update the active profile
func TestConfigCmd_UpdateActiveProfile(t *testing.T) {
	expectedErrorPattern := `^failed to update profile '.*' name to: .*\. '.*' is the active profile and cannot be deleted$`
	err := testutils_cobra.ExecutePingctl(t, "config",
		"--profile", "default",
		"--name", "myProfile",
		"--description", "hello")

	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}

// Test Config Command fails when provided an invalid profile name
func TestConfigCmd_InvalidProfileName(t *testing.T) {
	expectedErrorPattern := `^failed to update profile '.*' name to: .*\. invalid profile name: '.*'\. name must contain only alphanumeric characters, underscores, and dashes$`
	err := testutils_cobra.ExecutePingctl(t, "config",
		"--profile", "production",
		"--name", "pname&*^*&^$&@!",
		"--description", "hello")

	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}
