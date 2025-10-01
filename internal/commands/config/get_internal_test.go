// Copyright © 2025 Ping Identity Corporation

package config_internal

import (
	"testing"

	"github.com/pingidentity/pingcli/internal/configuration"
	"github.com/pingidentity/pingcli/internal/configuration/options"
	"github.com/pingidentity/pingcli/internal/customtypes"
	"github.com/pingidentity/pingcli/internal/profiles"
	"github.com/pingidentity/pingcli/internal/testing/testutils_koanf"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_RunInternalConfigGet(t *testing.T) {
	testCases := []struct {
		name          string
		profileName   customtypes.String
		koanfKey      string
		expectedError error
	}{
		{
			name:        "Get configuration for existing key",
			profileName: "default",
			koanfKey:    "service",
		},
		{
			name:          "Get configuration for invalid key",
			profileName:   "default",
			koanfKey:      "invalid-key",
			expectedError: configuration.ErrInvalidConfigurationKey,
		},
		{
			name:        "Get configuration with a different profile",
			profileName: "production",
			koanfKey:    "service",
		},
		{
			name:          "Get configuration with a non-existent profile",
			profileName:   "non-existent",
			koanfKey:      "service",
			expectedError: profiles.ErrProfileNameNotExist,
		},
		{
			name:        "Get configuration with a case-insensitive key",
			profileName: "default",
			koanfKey:    "SeRvIcE",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			testutils_koanf.InitKoanfs(t)

			options.RootProfileOption.Flag.Changed = true
			options.RootProfileOption.CobraParamValue = &tc.profileName

			err := RunInternalConfigGet(tc.koanfKey)

			if tc.expectedError != nil {
				require.Error(t, err)
				assert.ErrorIs(t, err, tc.expectedError)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
