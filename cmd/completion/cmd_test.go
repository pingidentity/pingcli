// Copyright © 2025 Ping Identity Corporation

package completion_test

import (
	"testing"

	"github.com/pingidentity/pingcli/internal/testing/testutils_cobra"
	"github.com/pingidentity/pingcli/internal/testing/testutils_koanf"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_CompletionCommand(t *testing.T) {
	testutils_koanf.InitKoanfs(t)

	testCases := []struct {
		name                string
		args                []string
		expectErr           bool
		expectedErrIs       error
		expectedErrContains string
	}{
		{
			name:                "Too few arguments",
			args:                []string{},
			expectErr:           true,
			expectedErrContains: "accepts 1 arg(s), received 0",
		},
		{
			name:      "Happy Path - bash",
			args:      []string{"bash"},
			expectErr: false,
		},
		{
			name:      "Happy Path - help",
			args:      []string{"--help"},
			expectErr: false,
		},
		{
			name:                "Too many arguments",
			args:                []string{"bash", "extra-arg"},
			expectErr:           true,
			expectedErrContains: "accepts 1 arg(s), received 2",
		},
		{
			name:                "Invalid flag",
			args:                []string{"--invalid-flag"},
			expectErr:           true,
			expectedErrContains: "unknown flag",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			testutils_koanf.InitKoanfs(t)

			err := testutils_cobra.ExecutePingcli(t, append([]string{"completion"}, tc.args...)...)

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
