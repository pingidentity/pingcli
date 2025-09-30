// Copyright © 2025 Ping Identity Corporation

package platform_test

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/pingidentity/pingcli/cmd/common"
	platform_internal "github.com/pingidentity/pingcli/internal/commands/platform"
	"github.com/pingidentity/pingcli/internal/configuration/options"
	"github.com/pingidentity/pingcli/internal/customtypes"
	"github.com/pingidentity/pingcli/internal/testing/testutils_cobra"
	"github.com/pingidentity/pingcli/internal/testing/testutils_koanf"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_PlatformExportCommand(t *testing.T) {
	testutils_koanf.InitKoanfs(t)

	testCases := []struct {
		name                string
		args                []string
		setup               func(t *testing.T, tempDir string)
		expectErr           bool
		expectedErrIs       error
		expectedErrContains string
	}{
		{
			name: "Happy Path - minimal flags",
			args: []string{
				"--" + options.PlatformExportOutputDirectoryOption.CobraParamName, "{{tempdir}}",
				"--" + options.PlatformExportOverwriteOption.CobraParamName,
			},
			expectErr: false,
		},
		{
			name:          "Too many arguments",
			args:          []string{"extra-arg"},
			expectErr:     true,
			expectedErrIs: common.ErrExactArgs,
		},
		{
			name:                "Invalid flag",
			args:                []string{"--invalid-flag"},
			expectErr:           true,
			expectedErrContains: "unknown flag",
		},
		{
			name:      "Happy path - help",
			args:      []string{"--help"},
			expectErr: false,
		},
		{
			name: "Happy Path - with service group",
			args: []string{
				"--" + options.PlatformExportOutputDirectoryOption.CobraParamName, "{{tempdir}}",
				"--" + options.PlatformExportOverwriteOption.CobraParamName,
				"--" + options.PlatformExportServiceGroupOption.CobraParamName, "pingone",
			},
			expectErr: false,
		},
		{
			name: "Invalid service group",
			args: []string{
				"--" + options.PlatformExportServiceGroupOption.CobraParamName, "invalid",
			},
			expectErr:     true,
			expectedErrIs: customtypes.ErrUnrecognisedServiceGroup,
		},
		{
			name: "Happy Path - with specific service",
			args: []string{
				"--" + options.PlatformExportOutputDirectoryOption.CobraParamName, "{{tempdir}}",
				"--" + options.PlatformExportOverwriteOption.CobraParamName,
				"--" + options.PlatformExportServiceOption.CobraParamName, customtypes.ENUM_EXPORT_SERVICE_PINGONE_PROTECT,
			},
			expectErr: false,
		},
		{
			name: "Happy Path - with specific service and format",
			args: []string{
				"--" + options.PlatformExportOutputDirectoryOption.CobraParamName, "{{tempdir}}",
				"--" + options.PlatformExportOverwriteOption.CobraParamName,
				"--" + options.PlatformExportServiceOption.CobraParamName, customtypes.ENUM_EXPORT_SERVICE_PINGONE_PROTECT,
				"--" + options.PlatformExportExportFormatOption.CobraParamName, customtypes.ENUM_EXPORT_FORMAT_HCL,
			},
			expectErr: false,
		},
		{
			name: "Invalid service",
			args: []string{
				"--" + options.PlatformExportServiceOption.CobraParamName, "invalid",
			},
			expectErr:     true,
			expectedErrIs: customtypes.ErrUnrecognisedExportService,
		},
		{
			name: "Invalid format",
			args: []string{
				"--" + options.PlatformExportExportFormatOption.CobraParamName, "invalid",
			},
			expectErr:     true,
			expectedErrIs: customtypes.ErrUnrecognisedFormat,
		},
		{
			name: "Invalid output directory",
			args: []string{
				"--" + options.PlatformExportOutputDirectoryOption.CobraParamName, "/invalid-dir",
			},
			expectErr:     true,
			expectedErrIs: platform_internal.ErrCreateOutputDirectory,
		},
		{
			name: "Overwrite false on non-empty directory",
			args: []string{
				"--" + options.PlatformExportOutputDirectoryOption.CobraParamName, "{{tempdir}}",
				"--" + options.PlatformExportOverwriteOption.CobraParamName + "=false",
			},
			setup: func(t *testing.T, tempDir string) {
				t.Helper()

				_, err := os.Create(filepath.Join(tempDir, "file")) // #nosec G304
				require.NoError(t, err)
			},
			expectErr:     true,
			expectedErrIs: platform_internal.ErrOutputDirectoryNotEmpty,
		},
		{
			name: "Happy Path - overwrite non-empty directory",
			args: []string{
				"--" + options.PlatformExportOutputDirectoryOption.CobraParamName, "{{tempdir}}",
				"--" + options.PlatformExportOverwriteOption.CobraParamName,
			},
			setup: func(t *testing.T, tempDir string) {
				t.Helper()

				_, err := os.Create(filepath.Join(tempDir, "file")) // #nosec G304
				require.NoError(t, err)
			},
			expectErr: false,
		},
		{
			name: "Happy Path - with pingone service and all required flags",
			args: []string{
				"--" + options.PlatformExportOutputDirectoryOption.CobraParamName, "{{tempdir}}",
				"--" + options.PlatformExportOverwriteOption.CobraParamName,
				"--" + options.PlatformExportServiceOption.CobraParamName, customtypes.ENUM_EXPORT_SERVICE_PINGONE_PROTECT,
				"--" + options.PingOneAuthenticationWorkerEnvironmentIDOption.CobraParamName, os.Getenv("TEST_PINGONE_ENVIRONMENT_ID"),
				"--" + options.PingOneAuthenticationWorkerClientIDOption.CobraParamName, os.Getenv("TEST_PINGONE_WORKER_CLIENT_ID"),
				"--" + options.PingOneAuthenticationWorkerClientSecretOption.CobraParamName, os.Getenv("TEST_PINGONE_WORKER_CLIENT_SECRET"),
				"--" + options.PingOneRegionCodeOption.CobraParamName, os.Getenv("TEST_PINGONE_REGION_CODE"),
			},
			expectErr: false,
		},
		{
			name: "PingOne flags not together",
			args: []string{
				"--" + options.PingOneAuthenticationWorkerEnvironmentIDOption.CobraParamName, os.Getenv("TEST_PINGONE_ENVIRONMENT_ID"),
			},
			expectErr:           true,
			expectedErrContains: "if any flags in the group [pingone-worker-environment-id pingone-worker-client-id pingone-worker-client-secret pingone-region-code] are set they must all be set",
		},
		{
			name: "Happy Path - with pingfederate service and all required flags for Basic Auth",
			args: []string{
				"--" + options.PlatformExportOutputDirectoryOption.CobraParamName, "{{tempdir}}",
				"--" + options.PlatformExportOverwriteOption.CobraParamName,
				"--" + options.PlatformExportServiceOption.CobraParamName, customtypes.ENUM_EXPORT_SERVICE_PINGFEDERATE,
				"--" + options.PingFederateBasicAuthUsernameOption.CobraParamName, "Administrator",
				"--" + options.PingFederateBasicAuthPasswordOption.CobraParamName, "2FederateM0re",
				"--" + options.PingFederateAuthenticationTypeOption.CobraParamName, customtypes.ENUM_PINGFEDERATE_AUTHENTICATION_TYPE_BASIC,
			},
			expectErr: false,
		},
		{
			name: "PingFederate Basic Auth flags not together",
			args: []string{
				"--" + options.PingFederateBasicAuthUsernameOption.CobraParamName, "Administrator",
			},
			expectErr:           true,
			expectedErrContains: "if any flags in the group [pingfederate-username pingfederate-password] are set they must all be set",
		},
		{
			name: "Pingone export fails with invalid credentials",
			args: []string{
				"--" + options.PlatformExportOutputDirectoryOption.CobraParamName, "{{tempdir}}",
				"--" + options.PlatformExportOverwriteOption.CobraParamName,
				"--" + options.PlatformExportServiceOption.CobraParamName, customtypes.ENUM_EXPORT_SERVICE_PINGONE_PROTECT,
				"--" + options.PingOneAuthenticationWorkerEnvironmentIDOption.CobraParamName, os.Getenv("TEST_PINGONE_ENVIRONMENT_ID"),
				"--" + options.PingOneAuthenticationWorkerClientIDOption.CobraParamName, os.Getenv("TEST_PINGONE_WORKER_CLIENT_ID"),
				"--" + options.PingOneAuthenticationWorkerClientSecretOption.CobraParamName, "invalid",
				"--" + options.PingOneRegionCodeOption.CobraParamName, os.Getenv("TEST_PINGONE_REGION_CODE"),
			},
			expectErr:     true,
			expectedErrIs: platform_internal.ErrPingOneInit,
		},
		{
			name: "Pingfederate export fails with invalid credentials",
			args: []string{
				"--" + options.PlatformExportOutputDirectoryOption.CobraParamName, "{{tempdir}}",
				"--" + options.PlatformExportOverwriteOption.CobraParamName,
				"--" + options.PlatformExportServiceOption.CobraParamName, customtypes.ENUM_EXPORT_SERVICE_PINGFEDERATE,
				"--" + options.PingFederateBasicAuthUsernameOption.CobraParamName, "Administrator",
				"--" + options.PingFederateBasicAuthPasswordOption.CobraParamName, "invalid",
				"--" + options.PingFederateAuthenticationTypeOption.CobraParamName, customtypes.ENUM_PINGFEDERATE_AUTHENTICATION_TYPE_BASIC,
			},
			expectErr:     true,
			expectedErrIs: platform_internal.ErrPingFederateInit,
		},
		{
			name: "Pingfederate Client Credentials Auth flags not together",
			args: []string{
				"--" + options.PingFederateClientCredentialsAuthClientIDOption.CobraParamName, "test",
			},
			expectErr:           true,
			expectedErrContains: "if any flags in the group [pingfederate-client-id pingfederate-client-secret pingfederate-token-url] are set they must all be set",
		},
		{
			name: "Pignfederate export fails with invalid Client Credentials Auth credentials",
			args: []string{
				"--" + options.PlatformExportOutputDirectoryOption.CobraParamName, "{{tempdir}}",
				"--" + options.PlatformExportOverwriteOption.CobraParamName,
				"--" + options.PlatformExportServiceOption.CobraParamName, customtypes.ENUM_EXPORT_SERVICE_PINGFEDERATE,
				"--" + options.PingFederateClientCredentialsAuthClientIDOption.CobraParamName, "test",
				"--" + options.PingFederateClientCredentialsAuthClientSecretOption.CobraParamName, "invalid",
				"--" + options.PingFederateClientCredentialsAuthTokenURLOption.CobraParamName, "https://localhost:9031/as/token.oauth2",
				"--" + options.PingFederateClientCredentialsAuthScopesOption.CobraParamName, "email",
				"--" + options.PingFederateAuthenticationTypeOption.CobraParamName, customtypes.ENUM_PINGFEDERATE_AUTHENTICATION_TYPE_CLIENT_CREDENTIALS,
			},
			expectErr:     true,
			expectedErrIs: platform_internal.ErrPingFederateInit,
		},
		{
			name: "Pingfederate export fails with invalid client credentials auth token URL",
			args: []string{
				"--" + options.PlatformExportOutputDirectoryOption.CobraParamName, "{{tempdir}}",
				"--" + options.PlatformExportOverwriteOption.CobraParamName,
				"--" + options.PlatformExportServiceOption.CobraParamName, customtypes.ENUM_EXPORT_SERVICE_PINGFEDERATE,
				"--" + options.PingFederateClientCredentialsAuthClientIDOption.CobraParamName, "test",
				"--" + options.PingFederateClientCredentialsAuthClientSecretOption.CobraParamName, "2FederateM0re!",
				"--" + options.PingFederateClientCredentialsAuthTokenURLOption.CobraParamName, "https://localhost:9031/as/invalid",
				"--" + options.PingFederateClientCredentialsAuthScopesOption.CobraParamName, "email",
				"--" + options.PingFederateAuthenticationTypeOption.CobraParamName, customtypes.ENUM_PINGFEDERATE_AUTHENTICATION_TYPE_CLIENT_CREDENTIALS,
			},
			expectErr:     true,
			expectedErrIs: platform_internal.ErrPingFederateInit,
		},
		{
			name: "Happy path - pingfederate with X-Bypass Header flag set to true",
			args: []string{
				"--" + options.PlatformExportOutputDirectoryOption.CobraParamName, "{{tempdir}}",
				"--" + options.PlatformExportOverwriteOption.CobraParamName,
				"--" + options.PlatformExportServiceOption.CobraParamName, customtypes.ENUM_EXPORT_SERVICE_PINGFEDERATE,
				"--" + options.PingFederateXBypassExternalValidationHeaderOption.CobraParamName,
				"--" + options.PingFederateBasicAuthUsernameOption.CobraParamName, "Administrator",
				"--" + options.PingFederateBasicAuthPasswordOption.CobraParamName, "2FederateM0re",
				"--" + options.PingFederateAuthenticationTypeOption.CobraParamName, customtypes.ENUM_PINGFEDERATE_AUTHENTICATION_TYPE_BASIC,
			},
			expectErr: false,
		},
		{
			name: "Happy path - pingfederate with Trust All TLS flag set to true",
			args: []string{
				"--" + options.PlatformExportOutputDirectoryOption.CobraParamName, "{{tempdir}}",
				"--" + options.PlatformExportOverwriteOption.CobraParamName,
				"--" + options.PlatformExportServiceOption.CobraParamName, customtypes.ENUM_EXPORT_SERVICE_PINGFEDERATE,
				"--" + options.PingFederateInsecureTrustAllTLSOption.CobraParamName,
				"--" + options.PingFederateBasicAuthUsernameOption.CobraParamName, "Administrator",
				"--" + options.PingFederateBasicAuthPasswordOption.CobraParamName, "2FederateM0re",
				"--" + options.PingFederateAuthenticationTypeOption.CobraParamName, customtypes.ENUM_PINGFEDERATE_AUTHENTICATION_TYPE_BASIC,
			},
			expectErr: false,
		},
		{
			name: "Pingfederate export fails with Trust All TLS flag set to false",
			args: []string{
				"--" + options.PlatformExportOutputDirectoryOption.CobraParamName, "{{tempdir}}",
				"--" + options.PlatformExportOverwriteOption.CobraParamName,
				"--" + options.PlatformExportServiceOption.CobraParamName, customtypes.ENUM_EXPORT_SERVICE_PINGFEDERATE,
				"--" + options.PingFederateInsecureTrustAllTLSOption.CobraParamName + "=false",
				"--" + options.PingFederateBasicAuthUsernameOption.CobraParamName, "Administrator",
				"--" + options.PingFederateBasicAuthPasswordOption.CobraParamName, "2FederateM0re",
				"--" + options.PingFederateAuthenticationTypeOption.CobraParamName, customtypes.ENUM_PINGFEDERATE_AUTHENTICATION_TYPE_BASIC,
			},
			expectErr:     true,
			expectedErrIs: platform_internal.ErrPingFederateInit,
		},
		{
			name: "Happy path - pingfederate with CA certificate PEM files flag set",
			args: []string{
				"--" + options.PlatformExportOutputDirectoryOption.CobraParamName, "{{tempdir}}",
				"--" + options.PlatformExportOverwriteOption.CobraParamName,
				"--" + options.PlatformExportServiceOption.CobraParamName, customtypes.ENUM_EXPORT_SERVICE_PINGFEDERATE,
				"--" + options.PingFederateCACertificatePemFilesOption.CobraParamName, "testdata/ssl-server-crt.pem",
				"--" + options.PingFederateBasicAuthUsernameOption.CobraParamName, "Administrator",
				"--" + options.PingFederateBasicAuthPasswordOption.CobraParamName, "2FederateM0re",
				"--" + options.PingFederateAuthenticationTypeOption.CobraParamName, customtypes.ENUM_PINGFEDERATE_AUTHENTICATION_TYPE_BASIC,
			},
			expectErr: false,
		},
		{
			name: "Pingfederate export fails with CA certificate PEM files flag set to invalid file",
			args: []string{
				"--" + options.PlatformExportOutputDirectoryOption.CobraParamName, "{{tempdir}}",
				"--" + options.PlatformExportOverwriteOption.CobraParamName,
				"--" + options.PlatformExportServiceOption.CobraParamName, customtypes.ENUM_EXPORT_SERVICE_PINGFEDERATE,
				"--" + options.PingFederateCACertificatePemFilesOption.CobraParamName, "invalid/crt.pem",
				"--" + options.PingFederateBasicAuthUsernameOption.CobraParamName, "Administrator",
				"--" + options.PingFederateBasicAuthPasswordOption.CobraParamName, "2FederateM0re",
				"--" + options.PingFederateAuthenticationTypeOption.CobraParamName, customtypes.ENUM_PINGFEDERATE_AUTHENTICATION_TYPE_BASIC,
			},
			expectErr:     true,
			expectedErrIs: platform_internal.ErrReadCaCertPemFile,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			testutils_koanf.InitKoanfs(t)

			tempDir := t.TempDir()
			finalArgs := make([]string, len(tc.args))
			for i, arg := range tc.args {
				finalArgs[i] = strings.ReplaceAll(arg, "{{tempdir}}", tempDir)
			}

			if tc.setup != nil {
				tc.setup(t, tempDir)
			}

			err := testutils_cobra.ExecutePingcli(t, append([]string{"platform", "export"}, finalArgs...)...)

			if !tc.expectErr {
				require.NoError(t, err)

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
