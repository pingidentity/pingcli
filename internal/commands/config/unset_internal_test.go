// Copyright © 2026 Ping Identity Corporation

package config_internal

import (
	"testing"

	"github.com/pingidentity/pingcli/internal/configuration"
	"github.com/pingidentity/pingcli/internal/configuration/options"
	"github.com/pingidentity/pingcli/internal/customtypes"
	"github.com/pingidentity/pingcli/internal/profiles"
	"github.com/pingidentity/pingcli/internal/testing/testutils_koanf"
	"github.com/stretchr/testify/require"
)

func Test_RunInternalConfigUnset(t *testing.T) {
	testutils_koanf.InitKoanfs(t)

	testCases := []struct {
		name          string
		profileName   customtypes.String
		koanfKey      string
		checkOption   *options.Option
		expectedError error
	}{
		{
			name:        "Unset noColor",
			koanfKey:    options.RootColorOption.KoanfKey,
			checkOption: &options.RootColorOption,
		},
		{
			name:          "Unset on non-existent key",
			koanfKey:      "nonExistantKey",
			expectedError: configuration.ErrInvalidConfigurationKey,
		},
		{
			name:        "Unset key on a different profile",
			profileName: customtypes.String("production"),
			koanfKey:    options.RootColorOption.KoanfKey,
			checkOption: &options.RootColorOption,
		},
		{
			name:          "Unset key with a non-existent profile",
			profileName:   customtypes.String("nonExistant"),
			koanfKey:      options.RootColorOption.KoanfKey,
			expectedError: profiles.ErrProfileNameNotExist,
		},
		{
			name:          "Run Unset with no key provided",
			koanfKey:      "",
			expectedError: configuration.ErrInvalidConfigurationKey,
		},
		{
			name:        "Unset with case-insensitive key",
			koanfKey:    "nOcOlOr",
			checkOption: &options.RootColorOption,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			testutils_koanf.InitKoanfs(t)

			if tc.profileName != "" {
				options.RootProfileOption.Flag.Changed = true
				options.RootProfileOption.CobraParamValue = &tc.profileName
			}

			err := RunInternalConfigUnset(tc.koanfKey)

			if tc.expectedError != nil {
				require.Error(t, err)
				require.ErrorIs(t, err, tc.expectedError)
			} else {
				require.NoError(t, err)
			}

			if tc.checkOption != nil {
				vVal, err := profiles.GetOptionValue(*tc.checkOption)
				if err != nil {
					require.Fail(t, "GetOptionValue returned error: %v", err)
				}

				if vVal != tc.checkOption.DefaultValue.String() {
					require.Fail(t, "Expected %s to be %s, got %v", tc.checkOption.KoanfKey, tc.checkOption.DefaultValue.String(), vVal)
				}
			}
		})
	}
}
