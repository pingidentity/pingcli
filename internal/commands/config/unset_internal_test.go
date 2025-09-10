// Copyright © 2025 Ping Identity Corporation

package config_internal

import (
	"errors"
	"testing"

	"github.com/pingidentity/pingcli/internal/configuration"
	"github.com/pingidentity/pingcli/internal/configuration/options"
	"github.com/pingidentity/pingcli/internal/customtypes"
	"github.com/pingidentity/pingcli/internal/profiles"
	"github.com/pingidentity/pingcli/internal/testing/testutils_koanf"
	"github.com/stretchr/testify/assert"
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
			name:          "Unset on non-existant key",
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
			name:          "Unset key with a non-existant profile",
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
				assert.Error(t, err)
				var unsetError *UnsetError
				if errors.As(err, &unsetError) {
					assert.ErrorIs(t, unsetError.Unwrap(), tc.expectedError)
				} else {
					assert.Fail(t, "Expected error to be of type UnsetError")
				}
			} else {
				assert.NoError(t, err)
			}

			if tc.checkOption != nil {
				vVal, err := profiles.GetOptionValue(*tc.checkOption)
				if err != nil {
					assert.Fail(t, "GetOptionValue returned error: %v", err)
				}

				if vVal != tc.checkOption.DefaultValue.String() {
					assert.Fail(t, "Expected %s to be %s, got %v", tc.checkOption.KoanfKey, tc.checkOption.DefaultValue.String(), vVal)
				}
			}
		})
	}
}
