// Copyright © 2025 Ping Identity Corporation

package plugin_internal

import (
	"os"
	"testing"

	"github.com/pingidentity/pingcli/internal/testing/testutils"
	"github.com/pingidentity/pingcli/internal/testing/testutils_koanf"
)

// Test RunInternalPluginAdd function
func Test_RunInternalPluginAdd(t *testing.T) {
	testutils_koanf.InitKoanfs(t)

	// Create a temporary $PATH for a test plugin
	pathDir := t.TempDir()
	t.Setenv("PATH", pathDir)

	testPlugin, err := os.CreateTemp(pathDir, "test-plugin-*.sh")
	if err != nil {
		t.Fatalf("Failed to create temporary plugin file: %v", err)
	}

	defer func() {
		err = os.Remove(testPlugin.Name())
		if err != nil {
			t.Fatalf("Failed to remove temporary plugin file: %v", err)
		}
	}()

	_, err = testPlugin.WriteString("#!/usr/bin/env sh\necho \"Hello, world!\"\nexit 0\n")
	if err != nil {
		t.Fatalf("Failed to write to temporary plugin file: %v", err)
	}

	err = testPlugin.Chmod(0755)
	if err != nil {
		t.Fatalf("Failed to set permissions on temporary plugin file: %v", err)
	}

	err = testPlugin.Close()
	if err != nil {
		t.Fatalf("Failed to close temporary plugin file: %v", err)
	}

	err = RunInternalPluginAdd(testPlugin.Name())
	if err != nil {
		t.Errorf("RunInternalPluginAdd returned error: %v", err)
	}
}

// Test RunInternalPluginAdd function fails with non-existent plugin
func Test_RunInternalPluginAdd_NonExistentPlugin(t *testing.T) {
	expectedErrorPattern := `^failed to add plugin: exec: .*: executable file not found in \$PATH$`
	err := RunInternalPluginAdd("non-existent-plugin")
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}
