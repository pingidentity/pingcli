// Copyright © 2026 Ping Identity Corporation

package config_test

import (
	"testing"

	"github.com/pingidentity/pingcli/cmd/common"
	"github.com/pingidentity/pingcli/internal/profiles"
	"github.com/pingidentity/pingcli/internal/testing/testutils_cobra"
	"github.com/pingidentity/pingcli/internal/testing/testutils_koanf"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_ConfigViewProfileCommand(t *testing.T) {
	testutils_koanf.InitKoanfs(t)

	testCases := []struct {
		name                string
		args                []string
		expectErr           bool
		expectedErrIs       error
		expectedErrContains string
	}{
		{
			name:      "Happy Path - view active profile",
			args:      []string{},
			expectErr: false,
		},
		{
			name:      "Happy Path - view specified profile",
			args:      []string{"production"},
			expectErr: false,
		},
		{
			name:      "Happy Path - with unmask-values flag",
			args:      []string{"--unmask-values"},
			expectErr: false,
		},
		{
			name:          "Too many arguments",
			args:          []string{"profile1", "profile2"},
			expectErr:     true,
			expectedErrIs: common.ErrRangeArgs,
		},
		{
			name:                "Invalid flag",
			args:                []string{"--invalid-flag"},
			expectErr:           true,
			expectedErrContains: "unknown flag",
		},
		{
			name:          "Non-existent profile",
			args:          []string{"non-existent"},
			expectErr:     true,
			expectedErrIs: profiles.ErrProfileNameNotExist,
		},
		{
			name:          "Invalid profile name format",
			args:          []string{"(*&*(#))"},
			expectErr:     true,
			expectedErrIs: profiles.ErrProfileNameNotExist,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			testutils_koanf.InitKoanfs(t)

			err := testutils_cobra.ExecutePingcli(t, append([]string{"config", "view-profile"}, tc.args...)...)

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
