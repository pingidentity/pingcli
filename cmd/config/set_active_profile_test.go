package config_test

import (
	"testing"

	"github.com/pingidentity/pingcli/internal/testing/testutils"
	"github.com/pingidentity/pingcli/internal/testing/testutils_cobra"
)

// Test Config set-active-profile Command Executes without issue
func TestConfigSetActiveProfileCmd_Execute(t *testing.T) {
	err := testutils_cobra.ExecutePingcli(t, "config", "set-active-profile", "--profile", "production")
	testutils.CheckExpectedError(t, err, nil)
}

// Test Config set-active-profile Command fails when provided too many arguments
func TestConfigSetActiveProfileCmd_TooManyArgs(t *testing.T) {
	expectedErrorPattern := `^failed to execute 'pingcli config set-active-profile': command accepts 0 arg\(s\), received 1$`
	err := testutils_cobra.ExecutePingcli(t, "config", "set-active-profile", "extra-arg")
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}

// Test Config set-active-profile Command fails when provided an invalid flag
func TestConfigSetActiveProfileCmd_InvalidFlag(t *testing.T) {
	expectedErrorPattern := `^unknown flag: --invalid$`
	err := testutils_cobra.ExecutePingcli(t, "config", "set-active-profile", "--invalid")
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}

// Test Config set-active-profile Command fails when provided an non-existent profile name
func TestConfigSetActiveProfileCmd_NonExistentProfileName(t *testing.T) {
	expectedErrorPattern := `^failed to set active profile: invalid profile name: '.*' profile does not exist$`
	err := testutils_cobra.ExecutePingcli(t, "config", "set-active-profile", "--profile", "nonexistent")
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}

// Test Config set-active-profile Command succeeds when provided the active profile
func TestConfigSetActiveProfileCmd_ActiveProfile(t *testing.T) {
	err := testutils_cobra.ExecutePingcli(t, "config", "set-active-profile", "--profile", "default")
	testutils.CheckExpectedError(t, err, nil)
}

// Test Config set-active-profile Command fails when provided an invalid profile name
func TestConfigSetActiveProfileCmd_InvalidProfileName(t *testing.T) {
	expectedErrorPattern := `^failed to set active profile: invalid profile name: '.*'\. name must contain only alphanumeric characters, underscores, and dashes$`
	err := testutils_cobra.ExecutePingcli(t, "config", "set-active-profile", "--profile", "pname&*^*&^$&@!")
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}

// Test Config set-active-profile Command --help, -h flag
func TestConfigSetActiveProfileCmd_HelpFlag(t *testing.T) {
	err := testutils_cobra.ExecutePingcli(t, "config", "set-active-profile", "--help")
	testutils.CheckExpectedError(t, err, nil)

	err = testutils_cobra.ExecutePingcli(t, "config", "set-active-profile", "-h")
	testutils.CheckExpectedError(t, err, nil)
}
