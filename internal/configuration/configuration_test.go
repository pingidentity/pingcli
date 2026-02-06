// Copyright © 2026 Ping Identity Corporation

package configuration_test

import (
	"strings"
	"testing"

	"github.com/pingidentity/pingcli/internal/configuration"
	"github.com/pingidentity/pingcli/internal/configuration/options"
	"github.com/pingidentity/pingcli/internal/testing/testutils_koanf"
	"github.com/stretchr/testify/require"
)

func Test_ValidateKoanfKey(t *testing.T) {
	testutils_koanf.InitKoanfs(t)

	testCases := []struct {
		name          string
		koanfKey      string
		expectedError error
	}{
		{
			name:     "Happy path - valid key",
			koanfKey: options.RootColorOption.KoanfKey,
		},
		{
			name:          "Invalid key",
			koanfKey:      "invalid-key",
			expectedError: configuration.ErrInvalidConfigurationKey,
		},
		{
			name:          "Empty key",
			koanfKey:      "",
			expectedError: configuration.ErrInvalidConfigurationKey,
		},
		{
			name:     "Happy Path - case insensitive key",
			koanfKey: strings.ToUpper(options.RootColorOption.KoanfKey),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			testutils_koanf.InitKoanfs(t)

			err := configuration.ValidateKoanfKey(tc.koanfKey)

			if tc.expectedError != nil {
				require.Error(t, err)
				require.ErrorIs(t, err, tc.expectedError)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func Test_ValidateParentKoanfKey(t *testing.T) {
	testutils_koanf.InitKoanfs(t)

	testCases := []struct {
		name          string
		koanfKey      string
		expectedError error
	}{
		{
			name:     "Happy path - valid parent key",
			koanfKey: strings.SplitN(options.PingOneAuthenticationTypeOption.KoanfKey, ".", 2)[0],
		},
		{
			name:          "Invalid key",
			koanfKey:      "invalid-parent-key",
			expectedError: configuration.ErrInvalidConfigurationKey,
		},
		{
			name:          "Empty key",
			koanfKey:      "",
			expectedError: configuration.ErrInvalidConfigurationKey,
		},
		{
			name:     "Happy Path - case insensitive parent key",
			koanfKey: strings.ToUpper(strings.SplitN(options.PingOneAuthenticationTypeOption.KoanfKey, ".", 2)[0]),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			testutils_koanf.InitKoanfs(t)

			err := configuration.ValidateParentKoanfKey(tc.koanfKey)

			if tc.expectedError != nil {
				require.Error(t, err)
				require.ErrorIs(t, err, tc.expectedError)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func Test_OptionFromKoanfKey(t *testing.T) {
	testutils_koanf.InitKoanfs(t)

	testCases := []struct {
		name           string
		koanfKey       string
		expectedOption options.Option
		expectedError  error
	}{
		{
			name:           "Happy path - valid key",
			koanfKey:       options.RootColorOption.KoanfKey,
			expectedOption: options.RootColorOption,
		},
		{
			name:           "Happy path - case insensitive key",
			koanfKey:       strings.ToUpper(options.RootColorOption.KoanfKey),
			expectedOption: options.RootColorOption,
		},
		{
			name:          "Invalid key",
			koanfKey:      "invalid-key",
			expectedError: configuration.ErrNoOptionForKey,
		},
		{
			name:          "Empty key",
			koanfKey:      "",
			expectedError: configuration.ErrEmptyKeyForOptionSearch,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			testutils_koanf.InitKoanfs(t)

			opt, err := configuration.OptionFromKoanfKey(tc.koanfKey)

			if tc.expectedError != nil {
				require.Error(t, err)
				require.ErrorIs(t, err, tc.expectedError)
			} else {
				require.NoError(t, err)
			}

			require.Equal(t, tc.expectedOption, opt)
		})
	}
}
