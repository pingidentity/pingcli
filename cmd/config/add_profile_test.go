// Copyright © 2025 Ping Identity Corporation

package config_test

import (
	"testing"

	"github.com/pingidentity/pingcli/cmd/common"
	config_internal "github.com/pingidentity/pingcli/internal/commands/config"
	"github.com/pingidentity/pingcli/internal/configuration/options"
	"github.com/pingidentity/pingcli/internal/customtypes"
	"github.com/pingidentity/pingcli/internal/profiles"
	"github.com/pingidentity/pingcli/internal/testing/testutils_cobra"
	"github.com/pingidentity/pingcli/internal/testing/testutils_koanf"
	"github.com/stretchr/testify/assert"
)

func Test_ConfigAddProfileCommand(t *testing.T) {
	testutils_koanf.InitKoanfs(t)

	testCases := []struct {
		name                string
		args                []string
		expectErr           bool
		expectedErrIs       error
		expectedErrContains string
	}{
		{
			name: "Happy Path",
			args: []string{
				"--" + options.ConfigAddProfileNameOption.CobraParamName, "test-profile",
				"--" + options.ConfigAddProfileDescriptionOption.CobraParamName, "test description",
				"--" + options.ConfigAddProfileSetActiveOption.CobraParamName + "=false",
			},
			expectErr: false,
		},
		{
			name: "Happy Path - Profile names are case insensitive",
			args: []string{
				"--" + options.ConfigAddProfileNameOption.CobraParamName, "DEfAuLt",
				"--" + options.ConfigAddProfileDescriptionOption.CobraParamName, "test description",
				"--" + options.ConfigAddProfileSetActiveOption.CobraParamName + "=false",
			},
			expectErr: false,
		},
		{
			name: "Too many arguments",
			args: []string{
				"extra-arg",
			},
			expectErr:     true,
			expectedErrIs: common.ErrExactArgs,
		},
		{
			name: "Invalid flag",
			args: []string{
				"--invalid-flag",
			},
			expectErr:           true,
			expectedErrContains: "unknown flag",
		},
		{
			name: "Invalid value for valid flag",
			args: []string{
				"--" + options.ConfigAddProfileSetActiveOption.CobraParamName + "=invalid-value",
			},
			expectErr:     true,
			expectedErrIs: customtypes.ErrParseBool,
		},
		{
			name: "Duplicate profile name",
			args: []string{
				"--" + options.ConfigAddProfileNameOption.CobraParamName, "default",
				"--" + options.ConfigAddProfileDescriptionOption.CobraParamName, "test description",
				"--" + options.ConfigAddProfileSetActiveOption.CobraParamName + "=false",
			},
			expectErr:     true,
			expectedErrIs: profiles.ErrProfileNameAlreadyExists,
		},
		{
			name: "Invalid profile name",
			args: []string{
				"--" + options.ConfigAddProfileNameOption.CobraParamName, "pname&*^*&^$&@!",
				"--" + options.ConfigAddProfileDescriptionOption.CobraParamName, "test description",
				"--" + options.ConfigAddProfileSetActiveOption.CobraParamName + "=false",
			},
			expectErr:     true,
			expectedErrIs: profiles.ErrProfileNameFormat,
		},
		{
			name: "Profile name is activeProfile",
			args: []string{
				"--" + options.ConfigAddProfileNameOption.CobraParamName, options.RootActiveProfileOption.KoanfKey,
				"--" + options.ConfigAddProfileDescriptionOption.CobraParamName, "test description",
				"--" + options.ConfigAddProfileSetActiveOption.CobraParamName + "=false",
			},
			expectErr:     true,
			expectedErrIs: profiles.ErrProfileNameSameAsActiveProfileKey,
		},
		{
			name: "Profile name is empty",
			args: []string{
				"--" + options.ConfigAddProfileNameOption.CobraParamName, "",
				"--" + options.ConfigAddProfileDescriptionOption.CobraParamName, "test description",
				"--" + options.ConfigAddProfileSetActiveOption.CobraParamName + "=false",
			},
			expectErr:     true,
			expectedErrIs: config_internal.ErrNoProfileProvided,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			testutils_koanf.InitKoanfs(t)

			err := testutils_cobra.ExecutePingcli(t, append([]string{"config", "add-profile"}, tc.args...)...)

			if !tc.expectErr {
				assert.NoError(t, err)
				return
			}

			assert.Error(t, err)
			if tc.expectedErrIs != nil {
				assert.ErrorIs(t, err, tc.expectedErrIs)
			}
			if tc.expectedErrContains != "" {
				assert.ErrorContains(t, err, tc.expectedErrContains)
			}
		})
	}
}
