// Copyright © 2025 Ping Identity Corporation

package plugin_test

import (
	"testing"

	"github.com/pingidentity/pingcli/internal/testing/testutils"
	"github.com/pingidentity/pingcli/internal/testing/testutils_cobra"
)

// Test Plugin remove Command Executes without issue
func TestPluginRemoveCmd_Execute(t *testing.T) {
	t.SkipNow()

	// TODO: A test plugin that responds with a valid RPC configuration is needed
	// for pingcli to execute when the plugin is listed as used in pingcli.
	// We can probably use a future plugin for testing once made.
}

// Test Plugin remove Command succeeds when provided a non-existent plugin
func TestPluginRemoveCmd_NonExistentPlugin(t *testing.T) {
	err := testutils_cobra.ExecutePingcli(t, "plugin", "remove", "non-existent-plugin")
	testutils.CheckExpectedError(t, err, nil)
}

// Test Plugin remove Command fails when provided too many arguments
func TestPluginRemoveCmd_TooManyArgs(t *testing.T) {
	expectedErrorPattern := `^failed to execute 'pingcli plugin remove': command accepts 1 arg\(s\), received 2$`
	err := testutils_cobra.ExecutePingcli(t, "plugin", "remove", "test-plugin-name", "extra-arg")
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}

// Test Plugin remove Command fails when provided too few arguments
func TestPluginRemoveCmd_TooFewArgs(t *testing.T) {
	expectedErrorPattern := `^failed to execute 'pingcli plugin remove': command accepts 1 arg\(s\), received 0$`
	err := testutils_cobra.ExecutePingcli(t, "plugin", "remove")
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}

// Test Plugin remove Command fails when provided an invalid flag
func TestPluginRemoveCmd_InvalidFlag(t *testing.T) {
	expectedErrorPattern := `^unknown flag: --invalid$`
	err := testutils_cobra.ExecutePingcli(t, "plugin", "remove", "test-plugin-name", "--invalid")
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}

// Test Plugin remove Command --help, -h flag
func TestPluginRemoveCmd_HelpFlag(t *testing.T) {
	err := testutils_cobra.ExecutePingcli(t, "plugin", "remove", "--help")
	testutils.CheckExpectedError(t, err, nil)

	err = testutils_cobra.ExecutePingcli(t, "plugin", "remove", "-h")
	testutils.CheckExpectedError(t, err, nil)
}
