// Copyright © 2025 Ping Identity Corporation

package plugin_test

import (
	"testing"

	"github.com/pingidentity/pingcli/cmd/common"
	"github.com/pingidentity/pingcli/internal/testing/testutils_cobra"
	"github.com/pingidentity/pingcli/internal/testing/testutils_koanf"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_PluginRemoveCommand(t *testing.T) {
	testutils_koanf.InitKoanfs(t)

	testCases := []struct {
		name                string
		args                []string
		expectErr           bool
		expectedErrIs       error
		expectedErrContains string
	}{
		// { TODO: A test plugin that can be installed is needed to properly test removal.
		//	name:       "Happy Path - remove existing plugin",
		//	args:       []string{"existing-plugin"},
		//	expectErr:  false,
		// },
		{
			name:      "Happy Path - remove non-existent plugin",
			args:      []string{"non-existent-plugin"},
			expectErr: false,
		},
		{
			name:      "Happy Path - help",
			args:      []string{"--help"},
			expectErr: false,
		},
		{
			name:          "Too many arguments",
			args:          []string{"plugin-name", "extra-arg"},
			expectErr:     true,
			expectedErrIs: common.ErrExactArgs,
		},
		{
			name:          "Too few arguments",
			args:          []string{},
			expectErr:     true,
			expectedErrIs: common.ErrExactArgs,
		},
		{
			name:                "Invalid flag",
			args:                []string{"plugin-name", "--invalid-flag"},
			expectErr:           true,
			expectedErrContains: "unknown flag",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			testutils_koanf.InitKoanfs(t)

			err := testutils_cobra.ExecutePingcli(t, append([]string{"plugin", "remove"}, tc.args...)...)

			if !tc.expectErr {
				require.NoError(t, err)
				return
			}

			assert.Error(t, err)
			if tc.expectedErrIs != nil {
				assert.ErrorIs(t, err, tc.expectedErrIs)
			}
			if tc.expectedErrContains != "" {
				assert.ErrorContains(t, err, tc.expectedErrContains)
			}
		})
	}
}
