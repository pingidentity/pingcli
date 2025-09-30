// Copyright © 2025 Ping Identity Corporation

package config_test

import (
	"slices"
	"strings"
	"testing"

	"github.com/pingidentity/pingcli/cmd/common"
	"github.com/pingidentity/pingcli/internal/configuration/options"
	"github.com/pingidentity/pingcli/internal/testing/testutils"
	"github.com/pingidentity/pingcli/internal/testing/testutils_cobra"
	"github.com/pingidentity/pingcli/internal/testing/testutils_koanf"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_ConfigListKeysCommand(t *testing.T) {
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
			args:      []string{},
			expectErr: false,
		},
		{
			name:      "Happy Path - help",
			args:      []string{"--help"},
			expectErr: false,
		},
		{
			name:      "Happy Path - yaml flag",
			args:      []string{"--yaml"},
			expectErr: false,
		},
		{
			name:          "Too many arguments",
			args:          []string{"extra-arg"},
			expectErr:     true,
			expectedErrIs: common.ErrExactArgs,
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

			output := testutils.CaptureStdout(func() {
				err := testutils_cobra.ExecutePingcli(t, append([]string{"config", "list-keys"}, tc.args...)...)

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

			if !tc.expectErr && !slices.Contains(tc.args, "--help") {
				for _, option := range options.Options() {
					if option == options.RootActiveProfileOption {
						continue
					}

					if slices.Contains(tc.args, "--yaml") {
						assert.Contains(t, output, option.KoanfKey[strings.LastIndex(option.KoanfKey, ".")+1:])
					} else {
						assert.Contains(t, output, option.KoanfKey)
					}
				}
			}
		})
	}
}
