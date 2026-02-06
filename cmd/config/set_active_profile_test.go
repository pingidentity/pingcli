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

func Test_ConfigSetActiveProfileCommand(t *testing.T) {
	testutils_koanf.InitKoanfs(t)

	testCases := []struct {
		name                string
		args                []string
		expectErr           bool
		expectedErrIs       error
		expectedErrContains string
	}{
		{
			name:      "Happy Path",
			args:      []string{"production"},
			expectErr: false,
		},
		{
			name:      "Happy Path - set active to current active",
			args:      []string{"default"},
			expectErr: false,
		},
		{
			name:      "Happy Path - help",
			args:      []string{"--help"},
			expectErr: false,
		},
		{
			name:          "Too many arguments",
			args:          []string{"extra-arg", "extra-arg2"},
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
			args:          []string{"nonexistent"},
			expectErr:     true,
			expectedErrIs: profiles.ErrProfileNameNotExist,
		},
		{
			name:          "Invalid profile name format",
			args:          []string{"pname&*^*&^$&@!"},
			expectErr:     true,
			expectedErrIs: profiles.ErrProfileNameNotExist,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			testutils_koanf.InitKoanfs(t)

			err := testutils_cobra.ExecutePingcli(t, append([]string{"config", "set-active-profile"}, tc.args...)...)

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
