// Copyright © 2025 Ping Identity Corporation

package plugin_internal

import (
	"testing"

	"github.com/pingidentity/pingcli/internal/testing/testutils_koanf"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_RunInternalPluginRemove(t *testing.T) {
	testutils_koanf.InitKoanfs(t)

	goldenPluginFileName := createGoldenPlugin(t)

	testCases := []struct {
		name              string
		pluginName        string
		createPluginFirst bool
		expectedError     error
	}{
		{
			name:              "Happy path - List plugins",
			pluginName:        goldenPluginFileName,
			createPluginFirst: true,
		},
		{
			name:       "Test non-existent plugin",
			pluginName: "non-existent-plugin",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			testutils_koanf.InitKoanfs(t)

			if tc.createPluginFirst {
				err := RunInternalPluginAdd(tc.pluginName)
				require.NoError(t, err)
			}

			err := RunInternalPluginRemove(tc.pluginName)

			if tc.expectedError != nil {
				require.Error(t, err)
				assert.ErrorIs(t, err, tc.expectedError)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
