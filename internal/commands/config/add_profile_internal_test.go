// Copyright © 2025 Ping Identity Corporation

package config_internal

import (
	"errors"
	"os"
	"testing"

	"github.com/pingidentity/pingcli/internal/configuration/options"
	"github.com/pingidentity/pingcli/internal/customtypes"
	"github.com/pingidentity/pingcli/internal/profiles"
	"github.com/pingidentity/pingcli/internal/testing/testutils_koanf"
	"github.com/stretchr/testify/assert"
)

func Test_RunInternalConfigAddProfile(t *testing.T) {
	testCases := []struct {
		name          string
		profileName   customtypes.String
		description   customtypes.String
		setActive     customtypes.Bool
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
			expectedError: ErrProfileNameNotProvided,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			testutils_koanf.InitKoanfs(t)

			options.ConfigAddProfileNameOption.Flag.Changed = true
			options.ConfigAddProfileNameOption.CobraParamValue = &tc.profileName

			options.ConfigAddProfileDescriptionOption.Flag.Changed = true
			options.ConfigAddProfileDescriptionOption.CobraParamValue = &tc.description

			options.ConfigAddProfileSetActiveOption.Flag.Changed = true
			options.ConfigAddProfileSetActiveOption.CobraParamValue = &tc.setActive

			err := RunInternalConfigAddProfile(os.Stdin)

			if tc.expectedError != nil {
				assert.Error(t, err)
				var addProfileErr *AddProfileError
				if errors.As(err, &addProfileErr) {
					assert.ErrorIs(t, addProfileErr.Unwrap(), tc.expectedError)
				} else {
					assert.Fail(t, "Expected error to be of type AddProfileError")
				}
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
