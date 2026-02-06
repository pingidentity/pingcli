// Copyright © 2026 Ping Identity Corporation

package config_internal

import (
	"io"
	"os"
	"strings"
	"testing"

	"github.com/pingidentity/pingcli/internal/configuration/options"
	"github.com/pingidentity/pingcli/internal/customtypes"
	"github.com/pingidentity/pingcli/internal/testing/testutils_koanf"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_RunInternalConfigListKeys(t *testing.T) {
	testutils_koanf.InitKoanfs(t)

	testCases := []struct {
		name          string
		contains      []string
		notContains   []string
		enableYaml    customtypes.Bool
		expectedError error
	}{
		{
			name: "Get List of Keys",
			contains: []string{
				options.RootColorOption.KoanfKey,
				options.RootOutputFormatOption.KoanfKey,
				options.ProfileDescriptionOption.KoanfKey,
				options.PlatformExportServiceGroupOption.KoanfKey,
				options.PingFederateAdminAPIPathOption.KoanfKey,
				options.PingOneAuthenticationWorkerClientIDOption.KoanfKey,
			},
			notContains: []string{
				options.RootActiveProfileOption.KoanfKey,
			},
		},
		{
			name:       "Get List of Keys in YAML format",
			enableYaml: true,
			contains: []string{
				strings.Split(options.PlatformExportServiceGroupOption.KoanfKey, ".")[0] + ":",
				strings.Split(options.PingFederateAdminAPIPathOption.KoanfKey, ".")[0] + ":",
				strings.Split(options.PingOneAuthenticationWorkerClientIDOption.KoanfKey, ".")[0] + ":",
			},
			notContains: []string{
				options.PlatformExportServiceGroupOption.KoanfKey,
				options.PingFederateAdminAPIPathOption.KoanfKey,
				options.PingOneAuthenticationWorkerClientIDOption.KoanfKey,
				options.RootActiveProfileOption.KoanfKey,
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			testutils_koanf.InitKoanfs(t)

			if tc.enableYaml {
				options.ConfigListKeysYamlOption.Flag.Changed = true
				options.ConfigListKeysYamlOption.CobraParamValue = &tc.enableYaml
			}

			originalStdout := os.Stdout
			r, w, _ := os.Pipe()
			os.Stdout = w

			t.Cleanup(func() {
				os.Stdout = originalStdout
			})

			err := RunInternalConfigListKeys()

			if tc.expectedError != nil {
				require.Error(t, err)
				require.ErrorIs(t, err, tc.expectedError)
			} else {
				require.NoError(t, err)
			}

			err = w.Close()
			require.NoError(t, err)

			capturedOutputBytes, _ := io.ReadAll(r)
			capturedOutput := string(capturedOutputBytes)

			for _, expected := range tc.contains {
				assert.Contains(t, capturedOutput, expected)
			}

			for _, notExpected := range tc.notContains {
				assert.NotContains(t, capturedOutput, notExpected)
			}
		})
	}
}
