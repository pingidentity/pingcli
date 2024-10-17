package config_test

import (
	"testing"

	"github.com/pingidentity/pingcli/internal/testing/testutils"
	"github.com/pingidentity/pingcli/internal/testing/testutils_cobra"
)

// Test Config delete-profile Command Executes without issue
func TestConfigDeleteProfileCmd_Execute(t *testing.T) {
	err := testutils_cobra.ExecutePingcli(t, "config", "delete-profile", "--yes", "production")
	testutils.CheckExpectedError(t, err, nil)
}

// Test Config delete-profile Command fails when provided too many arguments
func TestConfigDeleteProfileCmd_TooManyArgs(t *testing.T) {
	expectedErrorPattern := `^failed to execute '.*': command accepts 0 to 1 arg\(s\), received 2$`
	err := testutils_cobra.ExecutePingcli(t, "config", "delete-profile", "extra-arg", "extra-arg2")
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}

// Test Config delete-profile Command fails when provided an invalid flag
func TestConfigDeleteProfileCmd_InvalidFlag(t *testing.T) {
	expectedErrorPattern := `^unknown flag: --invalid$`
	err := testutils_cobra.ExecutePingcli(t, "config", "delete-profile", "--invalid")
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}

// Test Config delete-profile Command fails when provided an non-existent profile name
func TestConfigDeleteProfileCmd_NonExistentProfileName(t *testing.T) {
	expectedErrorPattern := `^failed to delete profile: invalid profile name: '.*' profile does not exist$`
	err := testutils_cobra.ExecutePingcli(t, "config", "delete-profile", "--yes", "nonexistent")
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}

// Test Config delete-profile Command fails when provided the active profile
func TestConfigDeleteProfileCmd_ActiveProfile(t *testing.T) {
	expectedErrorPattern := `^failed to delete profile: '.*' is the active profile and cannot be deleted$`
	err := testutils_cobra.ExecutePingcli(t, "config", "delete-profile", "--yes", "default")
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}

// Test Config delete-profile Command fails when provided an invalid profile name
func TestConfigDeleteProfileCmd_InvalidProfileName(t *testing.T) {
	expectedErrorPattern := `^failed to delete profile: invalid profile name: '.*'\. name must contain only alphanumeric characters, underscores, and dashes$`
	err := testutils_cobra.ExecutePingcli(t, "config", "delete-profile", "--yes", "pname&*^*&^$&@!")
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}
