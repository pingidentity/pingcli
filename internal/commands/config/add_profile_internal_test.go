// Copyright © 2025 Ping Identity Corporation

package config_internal

import (
	"os"
	"testing"

	"github.com/pingidentity/pingcli/internal/configuration/options"
	"github.com/pingidentity/pingcli/internal/customtypes"
	"github.com/pingidentity/pingcli/internal/profiles"
	"github.com/pingidentity/pingcli/internal/testing/testutils_koanf"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_RunInternalConfigAddProfile(t *testing.T) {
	testCases := []struct {
		name          string
		profileName   customtypes.String
		description   customtypes.String
		setActive     customtypes.Bool
		setKoanfNil   bool
		expectedError error
	}{
		{
			name:        "Create New Profile and Not Set it as the Active Profile",
			profileName: "test-profile",
			description: "test-description",
			setActive:   customtypes.Bool(false),
		},
		{
			name:        "Create New Profile and Set it as the Active Profile",
			profileName: "test-profile-active",
			description: "test-description-active",
			setActive:   customtypes.Bool(true),
		},
		{
			name:          "Invalid Profile Name: Already Exists",
			profileName:   "default",
			description:   "test-description",
			setActive:     customtypes.Bool(false),
			expectedError: profiles.ErrProfileNameAlreadyExists,
		},
		{
			name:          "Invalid Profile Name: None Provided",
			profileName:   "",
			description:   "test-description",
			setActive:     customtypes.Bool(false),
			expectedError: ErrNoProfileProvided,
		},
		{
			name:          "Koanf Not Initialized",
			profileName:   "new-profile",
			description:   "test-description",
			setActive:     customtypes.Bool(false),
			setKoanfNil:   true,
			expectedError: ErrKoanfNotInitialized,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			testutils_koanf.InitKoanfs(t)

			var koanfConfig *profiles.KoanfConfig
			if !tc.setKoanfNil {
				var err error
				koanfConfig, err = profiles.GetKoanfConfig()
				require.NoError(t, err)
			}

			options.ConfigAddProfileNameOption.Flag.Changed = true
			options.ConfigAddProfileNameOption.CobraParamValue = &tc.profileName

			options.ConfigAddProfileDescriptionOption.Flag.Changed = true
			options.ConfigAddProfileDescriptionOption.CobraParamValue = &tc.description

			options.ConfigAddProfileSetActiveOption.Flag.Changed = true
			options.ConfigAddProfileSetActiveOption.CobraParamValue = &tc.setActive

			err := RunInternalConfigAddProfile(os.Stdin, koanfConfig)

			if tc.expectedError != nil {
				require.Error(t, err)
				assert.ErrorIs(t, err, tc.expectedError)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
