// Copyright © 2026 Ping Identity Corporation

package config_internal

import (
	"os"
	"testing"

	"github.com/pingidentity/pingcli/internal/profiles"
	"github.com/pingidentity/pingcli/internal/testing/testutils_koanf"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_RunInternalConfigSetActiveProfile(t *testing.T) {
	testCases := []struct {
		name          string
		profileName   string
		expectedError error
	}{
		{
			name:        "Set different profile as active",
			profileName: "production",
		},
		{
			name:          "Set invalid profile name as active",
			profileName:   "(*#&)",
			expectedError: profiles.ErrProfileNameNotExist,
		},
		{
			name:          "Set non-existent profile as active",
			profileName:   "non-existent",
			expectedError: profiles.ErrProfileNameNotExist,
		},
		{
			name:          "Set empty profile name as active",
			profileName:   "",
			expectedError: profiles.ErrProfileNameEmpty,
		},
		{
			name:        "Set current active profile as active",
			profileName: "default",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			testutils_koanf.InitKoanfs(t)

			err := RunInternalConfigSetActiveProfile([]string{tc.profileName}, os.Stdin)

			if tc.expectedError != nil {
				require.Error(t, err)
				assert.ErrorIs(t, err, tc.expectedError)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
