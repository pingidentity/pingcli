// Copyright © 2025 Ping Identity Corporation

package plugin_test

import (
	"testing"

	"github.com/pingidentity/pingcli/internal/testing/testutils"
	"github.com/pingidentity/pingcli/internal/testing/testutils_cobra"
)

// Test Plugin list Command Executes without issue
func TestPluginListCmd_Execute(t *testing.T) {
	err := testutils_cobra.ExecutePingcli(t, "plugin", "list")
	testutils.CheckExpectedError(t, err, nil)
}

// Test Plugin list Command fails when provided too many arguments
func TestPluginListCmd_TooManyArgs(t *testing.T) {
	expectedErrorPattern := `^failed to execute 'pingcli plugin list': command accepts 0 arg\(s\), received 1$`
	err := testutils_cobra.ExecutePingcli(t, "plugin", "list", "extra-arg")
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}

// Test Plugin list Command fails when provided an invalid flag
func TestPluginListCmd_InvalidFlag(t *testing.T) {
	expectedErrorPattern := `^unknown flag: --invalid$`
	err := testutils_cobra.ExecutePingcli(t, "plugin", "list", "--invalid")
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}

// Test Plugin list Command --help, -h flag
func TestPluginListCmd_HelpFlag(t *testing.T) {
	err := testutils_cobra.ExecutePingcli(t, "plugin", "list", "--help")
	testutils.CheckExpectedError(t, err, nil)

	err = testutils_cobra.ExecutePingcli(t, "plugin", "list", "-h")
	testutils.CheckExpectedError(t, err, nil)
}
