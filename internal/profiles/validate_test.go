// Copyright © 2025 Ping Identity Corporation

package profiles_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/pingidentity/pingcli/internal/configuration/options"
	"github.com/pingidentity/pingcli/internal/profiles"
	"github.com/pingidentity/pingcli/internal/testing/testutils_koanf"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_Validate(t *testing.T) {
	testutils_koanf.InitKoanfs(t)

	testCases := []struct {
		name          string
		fileContents  string
		expectedError error
	}{
		{
			name:          "Happy path - Default",
			fileContents:  testutils_koanf.GetDefaultConfigFileContents(),
			expectedError: nil,
		},
		// { // validate() does not support case insensitive profile names or koanf keys
		//	name:         "Happy path - Legacy",
		//	fileContents: testutils_koanf.GetDefaultLegacyConfigFileContents(),
		// },
		{
			name:          "Invalid uuid",
			fileContents:  getInvalidUUIDFileContents(t),
			expectedError: profiles.ErrValidateUUID,
		},
		{
			name:          "Invalid region",
			fileContents:  getInvalidRegionFileContents(t),
			expectedError: profiles.ErrValidatePingOneRegionCode,
		},
		{
			name:          "Invalid bool",
			fileContents:  getInvalidBoolFileContents(t),
			expectedError: profiles.ErrValidateBoolean,
		},
		{
			name:          "Invalid output format",
			fileContents:  getInvalidOutputFormatFileContents(t),
			expectedError: profiles.ErrValidateOutputFormat,
		},
		{
			name:          "Invalid profile name",
			fileContents:  getInvalidProfileNameFileContents(t),
			expectedError: profiles.ErrProfileNameFormat,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			testutils_koanf.InitKoanfsCustomFile(t, tc.fileContents)

			err := profiles.Validate()

			if tc.expectedError != nil {
				assert.Error(t, err)
				assert.ErrorIs(t, err, tc.expectedError)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func getInvalidUUIDFileContents(t *testing.T) string {
	t.Helper()

	pingoneEnvIdKeyParts := strings.Split(options.PlatformExportPingOneEnvironmentIDOption.KoanfKey, ".")
	require.Equal(t, 3, len(pingoneEnvIdKeyParts))

	invalidUUIDFileContents := fmt.Sprintf(`%s: default
default:
    %s: "default description"
    %s:
        %s:
            %s: "invalid"`,
		options.RootActiveProfileOption.KoanfKey,
		options.ProfileDescriptionOption.KoanfKey,
		pingoneEnvIdKeyParts[0],
		pingoneEnvIdKeyParts[1],
		pingoneEnvIdKeyParts[2],
	)

	return invalidUUIDFileContents
}

func getInvalidRegionFileContents(t *testing.T) string {
	t.Helper()

	pingoneRegionCodeKeyParts := strings.Split(options.PingOneRegionCodeOption.KoanfKey, ".")
	require.Equal(t, 3, len(pingoneRegionCodeKeyParts))

	invalidRegionFileContents := fmt.Sprintf(`%s: default
default:
    %s: "default description"
    %s:
        %s:
            %s: "invalid"`,
		options.RootActiveProfileOption.KoanfKey,
		options.ProfileDescriptionOption.KoanfKey,
		pingoneRegionCodeKeyParts[0],
		pingoneRegionCodeKeyParts[1],
		pingoneRegionCodeKeyParts[2],
	)

	return invalidRegionFileContents
}

func getInvalidBoolFileContents(t *testing.T) string {
	t.Helper()

	invalidBoolFileContents := fmt.Sprintf(`%s: default
default:
    %s: "default description"
    %s: invalid`,
		options.RootActiveProfileOption.KoanfKey,
		options.ProfileDescriptionOption.KoanfKey,
		options.RootColorOption.KoanfKey,
	)

	return invalidBoolFileContents
}

func getInvalidOutputFormatFileContents(t *testing.T) string {
	t.Helper()

	invalidOutputFormatFileContents := fmt.Sprintf(`%s: default
default:
    %s: "default description"
    %s: invalid`,
		options.RootActiveProfileOption.KoanfKey,
		options.ProfileDescriptionOption.KoanfKey,
		options.RootOutputFormatOption.KoanfKey,
	)

	return invalidOutputFormatFileContents
}

func getInvalidProfileNameFileContents(t *testing.T) string {
	t.Helper()

	invalidProfileNameFileContents := fmt.Sprintf(`%s: default
default:
    %s: "default description"
invalid(&*^&*^&*^**$):
    %s: "default description"`,
		options.RootActiveProfileOption.KoanfKey,
		options.ProfileDescriptionOption.KoanfKey,
		options.ProfileDescriptionOption.KoanfKey,
	)

	return invalidProfileNameFileContents
}
