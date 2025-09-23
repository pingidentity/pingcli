// Copyright © 2025 Ping Identity Corporation

package config_internal

import (
	"fmt"
	"testing"

	"github.com/pingidentity/pingcli/internal/configuration"
	"github.com/pingidentity/pingcli/internal/configuration/options"
	"github.com/pingidentity/pingcli/internal/customtypes"
	"github.com/pingidentity/pingcli/internal/profiles"
	"github.com/pingidentity/pingcli/internal/testing/testutils_koanf"
	"github.com/stretchr/testify/require"
)

func Test_RunInternalConfigSet(t *testing.T) {
	testutils_koanf.InitKoanfs(t)

	testCases := []struct {
		name          string
		profileName   customtypes.String
		checkOption   *options.Option
		checkValue    string
		kvPair        string
		expectedError error
	}{
		{
			name:        "Set noColor to True",
			checkOption: &options.RootColorOption,
			checkValue:  "true",
			kvPair:      fmt.Sprintf("%s=true", options.RootColorOption.KoanfKey),
		},
		{
			name:          "Set active profile",
			kvPair:        fmt.Sprintf("%s=production", options.RootActiveProfileOption.KoanfKey),
			expectedError: ErrActiveProfileAssignment,
		},
		{
			name:          "Set non-existent key",
			kvPair:        "nonExistantKey=true",
			expectedError: configuration.ErrInvalidConfigurationKey,
		},
		{
			name:          "Set boolean key with invalid variable type",
			kvPair:        fmt.Sprintf("%s=invalid", options.RootColorOption.KoanfKey),
			expectedError: ErrMustBeBoolean,
		},
		{
			name:          "Set key on non-existent profile",
			profileName:   "non-existent",
			kvPair:        fmt.Sprintf("%s=true", options.RootColorOption.KoanfKey),
			expectedError: profiles.ErrProfileNameNotExist,
		},
		{
			name:        "Set noColor to True on different profile",
			profileName: "production",
			checkOption: &options.RootColorOption,
			checkValue:  "true",
			kvPair:      fmt.Sprintf("%s=true", options.RootColorOption.KoanfKey),
		},
		{
			name:          "Set key on invalid profile name format",
			profileName:   "(*#&)",
			kvPair:        fmt.Sprintf("%s=true", options.RootColorOption.KoanfKey),
			expectedError: profiles.ErrProfileNameNotExist,
		},
		{
			name:          "Set key with empty value",
			kvPair:        fmt.Sprintf("%s=", options.RootColorOption.KoanfKey),
			expectedError: ErrEmptyValue,
		},
		{
			name:          "Run set command with no key-value pair provided",
			kvPair:        "",
			expectedError: ErrKeyAssignmentFormat,
		},
		{
			name:          "Run set with invalid key-value assignment format",
			kvPair:        "key::value",
			expectedError: ErrKeyAssignmentFormat,
		},
		{
			name:        "Set value with case-insensitive key",
			kvPair:      "nOcOlOr=true",
			checkOption: &options.RootColorOption,
			checkValue:  "true",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			testutils_koanf.InitKoanfs(t)

			if tc.profileName != "" {
				options.RootProfileOption.Flag.Changed = true
				options.RootProfileOption.CobraParamValue = &tc.profileName
			}

			err := RunInternalConfigSet(tc.kvPair)

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

				if vVal != tc.checkValue {
					require.Fail(t, "Expected %s to be %s, got %v", tc.checkOption.KoanfKey, tc.checkValue, vVal)
				}
			}
		})
	}
}
