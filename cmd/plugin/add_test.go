// Copyright © 2025 Ping Identity Corporation

package plugin_test

import (
	"os"
	"testing"

	"github.com/pingidentity/pingcli/internal/testing/testutils"
	"github.com/pingidentity/pingcli/internal/testing/testutils_cobra"
)

// Test Plugin add Command Executes without issue
func TestPluginAddCmd_Execute(t *testing.T) {
	// Create a temporary PATH for a test plugin
	pathDir := t.TempDir()
	t.Setenv("PATH", pathDir)

	testPlugin, err := os.CreateTemp(pathDir, "test-plugin-*.sh")
	if err != nil {
		t.Fatalf("Failed to create temporary plugin file: %v", err)
	}
	defer os.Remove(testPlugin.Name())

	_, err = testPlugin.WriteString("#!/usr/bin/env sh\necho \"Hello, world!\"\nexit 0\n")
	if err != nil {
		t.Fatalf("Failed to write to temporary plugin file: %v", err)
	}

	err = testPlugin.Chmod(0755)
	if err != nil {
		t.Fatalf("Failed to set permissions on temporary plugin file: %v", err)
	}

	testPlugin.Close()

	err = testutils_cobra.ExecutePingcli(t, "plugin", "add", testPlugin.Name())
	testutils.CheckExpectedError(t, err, nil)
}

// Test Plugin add Command fails when provided a non-existent plugin
func TestPluginAddCmd_NonExistentPlugin(t *testing.T) {
	expectedErrorPattern := `^failed to add plugin: exec: .*: executable file not found in \$PATH$`
	err := testutils_cobra.ExecutePingcli(t, "plugin", "add", "non-existent-plugin")
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}

// Test Plugin add Command fails when provided too many arguments
func TestPluginAddCmd_TooManyArgs(t *testing.T) {
	expectedErrorPattern := `^failed to execute 'pingcli plugin add': command accepts 1 arg\(s\), received 2$`
	err := testutils_cobra.ExecutePingcli(t, "plugin", "add", "test-plugin-name", "extra-arg")
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}

// Test Plugin add Command fails when provided too few arguments
func TestPluginAddCmd_TooFewArgs(t *testing.T) {
	expectedErrorPattern := `^failed to execute 'pingcli plugin add': command accepts 1 arg\(s\), received 0$`
	err := testutils_cobra.ExecutePingcli(t, "plugin", "add")
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}

// Test Plugin add Command fails when provided an invalid flag
func TestPluginAddCmd_InvalidFlag(t *testing.T) {
	expectedErrorPattern := `^unknown flag: --invalid$`
	err := testutils_cobra.ExecutePingcli(t, "plugin", "add", "test-plugin-name", "--invalid")
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}

// Test Plugin add Command --help, -h flag
func TestPluginAddCmd_HelpFlag(t *testing.T) {
	err := testutils_cobra.ExecutePingcli(t, "plugin", "add", "--help")
	testutils.CheckExpectedError(t, err, nil)

	err = testutils_cobra.ExecutePingcli(t, "plugin", "add", "-h")
	testutils.CheckExpectedError(t, err, nil)
}
