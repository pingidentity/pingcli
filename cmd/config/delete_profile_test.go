// Copyright © 2025 Ping Identity Corporation

package config_test

import (
	"testing"

	"github.com/pingidentity/pingcli/cmd/common"
	"github.com/pingidentity/pingcli/internal/profiles"
	"github.com/pingidentity/pingcli/internal/testing/testutils_cobra"
	"github.com/pingidentity/pingcli/internal/testing/testutils_koanf"
	"github.com/stretchr/testify/assert"
)

func Test_ConfigDeleteProfileCommand(t *testing.T) {
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
			args:      []string{"--yes", "production"},
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
			args:          []string{"--yes", "non-existent"},
			expectErr:     true,
			expectedErrIs: profiles.ErrProfileNameNotExist,
		},
		{
			name:          "Active profile",
			args:          []string{"--yes", "default"},
			expectErr:     true,
			expectedErrIs: profiles.ErrDeleteActiveProfile,
		},
		{
			name:          "Invalid profile name",
			args:          []string{"--yes", "non-existent"},
			expectErr:     true,
			expectedErrIs: profiles.ErrProfileNameNotExist,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			testutils_koanf.InitKoanfs(t)

			err := testutils_cobra.ExecutePingcli(t, append([]string{"config", "delete-profile"}, tc.args...)...)

			if !tc.expectErr {
				assert.NoError(t, err)
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
