// Copyright © 2025 Ping Identity Corporation

package config_internal

import (
	"testing"

	"github.com/pingidentity/pingcli/internal/profiles"
	"github.com/pingidentity/pingcli/internal/testing/testutils_koanf"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_RunInternalConfigViewProfile(t *testing.T) {
	testCases := []struct {
		name          string
		profileName   []string
		expectedError error
	}{
		{
			name:        "View active profile by providing no profile",
			profileName: []string{},
		},
		{
			name:          "View non-existent profile",
			profileName:   []string{"nonexistent"},
			expectedError: profiles.ErrProfileNameNotExist,
		},
		{
			name:        "View profile by providing one",
			profileName: []string{"production"},
		},
		{
			name:          "View empty name profile",
			profileName:   []string{""},
			expectedError: profiles.ErrProfileNameEmpty,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			testutils_koanf.InitKoanfs(t)

			err := RunInternalConfigViewProfile(tc.profileName)

			if tc.expectedError != nil {
				require.Error(t, err)
				assert.ErrorIs(t, err, tc.expectedError)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
