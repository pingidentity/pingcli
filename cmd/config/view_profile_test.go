package config_test

import (
	"testing"

	"github.com/pingidentity/pingcli/internal/testing/testutils"
	"github.com/pingidentity/pingcli/internal/testing/testutils_cobra"
)

// Test Config Set Command Executes without issue
func TestConfigViewProfileCmd_Execute(t *testing.T) {
	err := testutils_cobra.ExecutePingcli(t, "config", "view-profile")
	testutils.CheckExpectedError(t, err, nil)
}

// Test Config Set Command Executes with --profile flag
func TestConfigViewProfileCmd_Execute_WithProfileFlag(t *testing.T) {
	err := testutils_cobra.ExecutePingcli(t, "config", "view-profile", "--profile", "production")
	testutils.CheckExpectedError(t, err, nil)
}

// Test Config Set Command fails with invalid flag
func TestConfigViewProfileCmd_Execute_InvalidFlag(t *testing.T) {
	expectedErrorPattern := `^unknown flag: --invalid$`
	err := testutils_cobra.ExecutePingcli(t, "config", "view-profile", "--invalid")
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}

// Test Config Set Command fails with non-existent profile
func TestConfigViewProfileCmd_Execute_NonExistentProfile(t *testing.T) {
	expectedErrorPattern := `^failed to view profile: invalid profile name: '.*' profile does not exist$`
	err := testutils_cobra.ExecutePingcli(t, "config", "view-profile", "--profile", "non-existent")
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}

// Test Config Set Command fails with invalid profile
func TestConfigViewProfileCmd_Execute_InvalidProfile(t *testing.T) {
	expectedErrorPattern := `^failed to view profile: invalid profile name: '.*'\. name must contain only alphanumeric characters, underscores, and dashes$`
	err := testutils_cobra.ExecutePingcli(t, "config", "view-profile", "--profile", "(*&*(#))")
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}
