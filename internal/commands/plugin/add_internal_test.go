// Copyright © 2025 Ping Identity Corporation

package plugin_internal

import (
	"os"
	"testing"

	"github.com/pingidentity/pingcli/internal/testing/testutils_koanf"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_RunInternalPluginAdd(t *testing.T) {
	testutils_koanf.InitKoanfs(t)

	goldenPluginFileName := createGoldenPlugin(t)

	testCases := []struct {
		name          string
		pluginName    string
		expectedError error
	}{
		{
			name:       "Happy path - Add plugin",
			pluginName: goldenPluginFileName,
		},
		{
			name:          "Test non-existent plugin",
			pluginName:    "non-existent-plugin",
			expectedError: ErrPluginNotFound,
		},
		{
			name:          "Test empty plugin name",
			pluginName:    "",
			expectedError: ErrPluginNameEmpty,
		},
		// TODO - In testutils_koanf.InitKoanfs(t), create a valid plugin executable and add it to the config and path similar to below
		// {
		// 	name:          "Test adding a plugin that already exists",
		// 	pluginName:    "existing-plugin",
		// 	expectedError: ErrPluginAlreadyExists,
		// },
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			testutils_koanf.InitKoanfs(t)

			err := RunInternalPluginAdd(tc.pluginName)

			if tc.expectedError != nil {
				require.Error(t, err)
				assert.ErrorIs(t, err, tc.expectedError)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func createGoldenPlugin(t *testing.T) string {
	t.Helper()

	pathDir := t.TempDir()
	t.Setenv("PATH", pathDir)

	testPlugin, err := os.CreateTemp(pathDir, "test-plugin-*.sh")
	require.NoError(t, err)

	_, err = testPlugin.WriteString("#!/usr/bin/env sh\necho \"Hello, world!\"\nexit 0\n")
	require.NoError(t, err)

	err = testPlugin.Chmod(0755)
	require.NoError(t, err)

	err = testPlugin.Close()
	require.NoError(t, err)

	return testPlugin.Name()
}
