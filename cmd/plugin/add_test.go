// Copyright © 2026 Ping Identity Corporation

package plugin_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/pingidentity/pingcli/cmd/common"
	plugin_internal "github.com/pingidentity/pingcli/internal/commands/plugin"
	"github.com/pingidentity/pingcli/internal/testing/testutils_cobra"
	"github.com/pingidentity/pingcli/internal/testing/testutils_koanf"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_PluginAddCommand(t *testing.T) {
	testutils_koanf.InitKoanfs(t)

	goldenPlugin := createGoldenPluginExecutable(t)
	pluginFilename := filepath.Base(goldenPlugin)
	require.FileExists(t, goldenPlugin, "Test plugin executable does not exist")

	t.Cleanup(func() {
		err := os.Remove(goldenPlugin)
		require.NoError(t, err, "Failed to remove test plugin executable")
	})

	testCases := []struct {
		name                string
		args                []string
		expectErr           bool
		expectedErrIs       error
		expectedErrContains string
	}{
		{
			name:      "Happy Path",
			args:      []string{pluginFilename},
			expectErr: false,
		},
		{
			name:      "Happy Path - help",
			args:      []string{"--help"},
			expectErr: false,
		},
		{
			name: "Non-existent plugin",
			args: []string{
				"non-existent-plugin",
			},
			expectErr:     true,
			expectedErrIs: plugin_internal.ErrPluginNotFound,
		},
		{
			name:          "Too many arguments",
			args:          []string{"arg", "extra-arg"},
			expectErr:     true,
			expectedErrIs: common.ErrExactArgs,
		},
		{
			name:                "Invalid flag",
			args:                []string{"test-plugin-name", "--invalid-flag"},
			expectErr:           true,
			expectedErrContains: "unknown flag",
		},
		{
			name:          "Too few arguments",
			args:          []string{},
			expectErr:     true,
			expectedErrIs: common.ErrExactArgs,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			testutils_koanf.InitKoanfs(t)

			err := testutils_cobra.ExecutePingcli(t, append([]string{"plugin", "add"}, tc.args...)...)

			if !tc.expectErr {
				require.NoError(t, err)

				return
			}

			assert.Error(t, err)
			if tc.expectedErrIs != nil {
				assert.ErrorIs(t, err, tc.expectedErrIs)
			}
			if tc.expectedErrContains != "" {
				assert.Contains(t, err.Error(), tc.expectedErrContains)
			}
		})
	}
}

func createGoldenPluginExecutable(t *testing.T) string {
	t.Helper()

	pathDir := t.TempDir()
	t.Setenv("PATH", pathDir)

	testPlugin, err := os.CreateTemp(pathDir, "test-plugin-*.sh")
	require.NoError(t, err, "Failed to create temporary plugin file")

	_, err = testPlugin.WriteString("#!/usr/bin/env sh\necho \"Hello, world!\"\nexit 0\n")
	require.NoError(t, err, "Failed to write to temporary plugin file")

	err = testPlugin.Chmod(0755)
	require.NoError(t, err, "Failed to set permissions on temporary plugin file")

	err = testPlugin.Close()
	require.NoError(t, err, "Failed to close temporary plugin file")

	return testPlugin.Name()
}
