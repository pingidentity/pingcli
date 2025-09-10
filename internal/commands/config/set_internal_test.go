// Copyright © 2025 Ping Identity Corporation

package config_internal

import (
	"errors"
	"fmt"
	"testing"

	"github.com/pingidentity/pingcli/internal/configuration"
	"github.com/pingidentity/pingcli/internal/configuration/options"
	"github.com/pingidentity/pingcli/internal/customtypes"
	"github.com/pingidentity/pingcli/internal/profiles"
	"github.com/pingidentity/pingcli/internal/testing/testutils_koanf"
	"github.com/stretchr/testify/assert"
)

func Test_RunInternalConfigSet(t *testing.T) {
	testutils_koanf.InitKoanfs(t)

	testCases := []struct {
		name          string
		profileName   customtypes.String
		kvPair        string
		expectedError error
	}{
		{
			name:   "Set noColor to True",
			kvPair: fmt.Sprintf("%s=true", options.RootColorOption.KoanfKey),
		},
		{
			name:          "Set active profile",
			kvPair:        fmt.Sprintf("%s=production", options.RootActiveProfileOption.KoanfKey),
			expectedError: ErrActiveProfileAssignment,
		},
		{
			name:          "Set non-existant key",
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
				assert.Error(t, err)
				var setError *SetError
				if errors.As(err, &setError) {
					assert.ErrorIs(t, setError.Unwrap(), tc.expectedError)
				} else {
					assert.Fail(t, "Expected error to be of type SetError")
				}
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

// Test Test_RunInternalConfigSet function succeeds with case-insensitive keys
func Test_RunInternalConfigSet_CaseInsensitiveKeys(t *testing.T) {
	testutils_koanf.InitKoanfs(t)

	err := RunInternalConfigSet("NoCoLoR=true")
	if err != nil {
		t.Errorf("RunInternalConfigSet returned error: %v", err)
	}

	// Make sure the actual correct key was set, not the case-insensitive one
	vVal, err := profiles.GetOptionValue(options.RootColorOption)
	if err != nil {
		t.Errorf("GetOptionValue returned error: %v", err)
	}

	if vVal != "true" {
		t.Errorf("Expected %s to be true, got %v", options.RootColorOption.KoanfKey, vVal)
	}
}
