// Copyright © 2026 Ping Identity Corporation

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

func Test_RunInternalConfigDeleteProfile(t *testing.T) {
	testCases := []struct {
		name              string
		profileName       string
		autoConfirmDelete customtypes.Bool
		expectedError     error
	}{
		{
			name:              "Delete Existing Profile",
			profileName:       "production",
			autoConfirmDelete: true,
		},
		{
			name:              "Invalid Profile Name: Active Profile",
			profileName:       "default",
			autoConfirmDelete: true,
			expectedError:     profiles.ErrDeleteActiveProfile,
		},
		{
			name:              "Invalid Profile Name: Non-Existent Profile",
			profileName:       "non-existent",
			autoConfirmDelete: false,
			expectedError:     profiles.ErrProfileNameNotExist,
		},
		{
			name:              "Invalid Profile Name: Empty Profile Name",
			profileName:       "",
			autoConfirmDelete: false,
			expectedError:     profiles.ErrProfileNameEmpty,
		},
		{
			name:              "Invalid Profile Name: Special Characters",
			profileName:       "(*#&)",
			autoConfirmDelete: false,
			expectedError:     profiles.ErrProfileNameNotExist,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			testutils_koanf.InitKoanfs(t)

			options.ConfigDeleteAutoAcceptOption.Flag.Changed = true
			options.ConfigDeleteAutoAcceptOption.CobraParamValue = &tc.autoConfirmDelete

			err := RunInternalConfigDeleteProfile([]string{tc.profileName}, os.Stdin)

			if tc.expectedError != nil {
				require.Error(t, err)
				assert.ErrorIs(t, err, tc.expectedError)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
