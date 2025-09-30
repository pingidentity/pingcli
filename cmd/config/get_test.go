// Copyright © 2025 Ping Identity Corporation

package config_test

import (
	"testing"

	"github.com/pingidentity/pingcli/cmd/common"
	"github.com/pingidentity/pingcli/internal/configuration"
	"github.com/pingidentity/pingcli/internal/configuration/options"
	"github.com/pingidentity/pingcli/internal/testing/testutils_cobra"
	"github.com/pingidentity/pingcli/internal/testing/testutils_koanf"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_ConfigGetCommand(t *testing.T) {
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
			args:      []string{"export"},
			expectErr: false,
		},
		{
			name:          "Too many arguments",
			args:          []string{options.RootColorOption.KoanfKey, options.RootOutputFormatOption.KoanfKey},
			expectErr:     true,
			expectedErrIs: common.ErrExactArgs,
		},
		{
			name:      "Full key",
			args:      []string{options.PingOneAuthenticationWorkerClientIDOption.KoanfKey},
			expectErr: false,
		},
		{
			name:      "Partial key",
			args:      []string{"service.pingOne"},
			expectErr: false,
		},
		{
			name:          "Invalid key",
			args:          []string{"pingcli.invalid"},
			expectErr:     true,
			expectedErrIs: configuration.ErrInvalidConfigurationKey,
		},
		{
			name:                "Invalid flag",
			args:                []string{"--invalid-flag"},
			expectErr:           true,
			expectedErrContains: "unknown flag",
		},
		{
			name:          "No key",
			args:          []string{},
			expectErr:     true,
			expectedErrIs: common.ErrExactArgs,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			testutils_koanf.InitKoanfs(t)

			err := testutils_cobra.ExecutePingcli(t, append([]string{"config", "get"}, tc.args...)...)

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
