// Copyright © 2026 Ping Identity Corporation

package platform_internal

import (
	"fmt"
	"os"
	"regexp"
	"testing"

	"github.com/pingidentity/pingcli/internal/configuration/options"
	"github.com/pingidentity/pingcli/internal/customtypes"
	"github.com/pingidentity/pingcli/internal/profiles"
	"github.com/pingidentity/pingcli/internal/testing/testutils"
	"github.com/pingidentity/pingcli/internal/testing/testutils_koanf"
	"github.com/stretchr/testify/require"
)

type testCase struct {
	name                       string
	services                   customtypes.ExportServices
	checkTfFiles               bool
	nilContext                 bool
	cACertPemFiles             customtypes.StringSlice
	pingOneAuthType            customtypes.PingOneAuthenticationType
	authStorageType            customtypes.StorageType
	pfAuthType                 customtypes.PingFederateAuthenticationType
	pfAccessToken              customtypes.String
	pfClientId                 *customtypes.String
	pfClientSecret             customtypes.String
	pfTokenURL                 customtypes.String
	outputDir                  customtypes.String
	overwriteOutputDirLocation bool
	changeWorkingDir           bool
	overwriteOnExport          customtypes.Bool
	expectedError              error
}

func Test_RunInternalExport(t *testing.T) {
	testutils_koanf.InitKoanfs(t)

	goldenCACertPemFile := createGoldenCACertPemFile(t)
	malformedCaCertPemFile := createMalformedCACertPemFile(t)
	unwriteableDir := createUnwriteableDir(t)
	unreadableDir := createUnreadableDir(t)
	nonEmptyDir := createNonEmptyDir(t)

	testCases := []testCase{
		{
			name: "Test Happy Path - All Services",
			services: []string{
				customtypes.ENUM_EXPORT_SERVICE_PINGONE_PROTECT,
				customtypes.ENUM_EXPORT_SERVICE_PINGFEDERATE,
				customtypes.ENUM_EXPORT_SERVICE_PINGONE_AUTHORIZE,
				customtypes.ENUM_EXPORT_SERVICE_PINGONE_MFA,
				customtypes.ENUM_EXPORT_SERVICE_PINGONE_PLATFORM,
				customtypes.ENUM_EXPORT_SERVICE_PINGONE_SSO,
			},
			pingOneAuthType: customtypes.PingOneAuthenticationType(customtypes.ENUM_PINGONE_AUTHENTICATION_TYPE_WORKER),
			authStorageType: customtypes.StorageType(customtypes.ENUM_STORAGE_TYPE_NONE),
			checkTfFiles:    true,
		},
		{
			name:     "Test export with no services selected",
			services: []string{},
			// No config means no interaction with PingOne auth type, so technically doesn't need it, but good practice if it starts doing verification
		},
		// TODO - The PF Container used for testing needs to support Access Token Auth
		// {
		// 	name:         "Test Happy Path - Access Token",
		// 	services:     []string{customtypes.ENUM_EXPORT_SERVICE_PINGFEDERATE},
		// 	checkTfFiles: true,
		// 	pfAuthType:   customtypes.PingFederateAuthenticationType(customtypes.ENUM_PINGFEDERATE_AUTHENTICATION_TYPE_ACCESS_TOKEN),
		// },
		// TODO - The PF Container used for testing needs to support Client Credentials Auth
		// {
		// 	name:         "Test Happy Path - PingFederate Client Credentials",
		// 	services:     []string{customtypes.ENUM_EXPORT_SERVICE_PINGFEDERATE},
		// 	checkTfFiles: true,
		// 	pfAuthType:   customtypes.PingFederateAuthenticationType(customtypes.ENUM_PINGFEDERATE_AUTHENTICATION_TYPE_CLIENT_CREDENTIALS),
		// },
		{
			name:          "Test with empty access token - PingFederate Access Token Auth",
			services:      []string{customtypes.ENUM_EXPORT_SERVICE_PINGFEDERATE},
			pfAuthType:    customtypes.PingFederateAuthenticationType(customtypes.ENUM_PINGFEDERATE_AUTHENTICATION_TYPE_ACCESS_TOKEN),
			pfAccessToken: "",
			expectedError: ErrAccessTokenEmpty,
		},
		{
			name:          "Test with invalid access token - PingFederate Access Token Auth",
			services:      []string{customtypes.ENUM_EXPORT_SERVICE_PINGFEDERATE},
			pfAuthType:    customtypes.PingFederateAuthenticationType(customtypes.ENUM_PINGFEDERATE_AUTHENTICATION_TYPE_ACCESS_TOKEN),
			pfAccessToken: "invalid-token",
			expectedError: ErrPingFederateInit,
		},
		{
			name:          "Test empty client credentials - PingFederate Client Credentials Auth",
			services:      []string{customtypes.ENUM_EXPORT_SERVICE_PINGFEDERATE},
			pfAuthType:    customtypes.PingFederateAuthenticationType(customtypes.ENUM_PINGFEDERATE_AUTHENTICATION_TYPE_CLIENT_CREDENTIALS),
			pfClientId:    customtypes.StringPtr(""),
			expectedError: ErrClientCredentialsEmpty,
		},
		{
			name:           "Test invalid client credentials - PingFederate Client Credentials Auth",
			services:       []string{customtypes.ENUM_EXPORT_SERVICE_PINGFEDERATE},
			pfAuthType:     customtypes.PingFederateAuthenticationType(customtypes.ENUM_PINGFEDERATE_AUTHENTICATION_TYPE_CLIENT_CREDENTIALS),
			pfClientId:     customtypes.StringPtr("invalid-client-id"),
			pfClientSecret: "invalid-client-secret",
			pfTokenURL:     "http://localhost:9031/pf-admin-api/v1/oauth/token",
			expectedError:  ErrPingFederateInit,
		},
		{
			name:           "Test Happy Path With PEM file - PingFederate",
			services:       []string{customtypes.ENUM_EXPORT_SERVICE_PINGFEDERATE},
			checkTfFiles:   true,
			cACertPemFiles: *goldenCACertPemFile,
		},
		{
			name:          "Test with nil context",
			nilContext:    true,
			expectedError: ErrNilContext,
		},
		{
			name:           "Test with invalid PEM filepath - PingFederate",
			services:       []string{customtypes.ENUM_EXPORT_SERVICE_PINGFEDERATE},
			cACertPemFiles: []string{"/invalid/file/path.pem"},
			expectedError:  ErrReadCaCertPemFile,
		},
		{
			name:           "Test with malformed PEM file - PingFederate",
			services:       []string{customtypes.ENUM_EXPORT_SERVICE_PINGFEDERATE},
			cACertPemFiles: *malformedCaCertPemFile,
			expectedError:  ErrPingFederateCACertParse,
		},
		{
			name:          "Test invalid PingFederate Auth Type",
			services:      []string{customtypes.ENUM_EXPORT_SERVICE_PINGFEDERATE},
			pfAuthType:    "invalid-auth-type",
			expectedError: ErrPingFederateAuthType,
		},
		{
			name:                       "Test empty output directory",
			services:                   []string{customtypes.ENUM_EXPORT_SERVICE_PINGONE_PROTECT},
			outputDir:                  "",
			overwriteOutputDirLocation: true,
			expectedError:              ErrOutputDirectoryEmpty,
		},
		{
			name:                       "Test non-writable output directory",
			services:                   []string{customtypes.ENUM_EXPORT_SERVICE_PINGONE_PROTECT},
			outputDir:                  customtypes.String(unwriteableDir),
			overwriteOutputDirLocation: true,
			expectedError:              ErrCreateOutputDirectory,
		},
		{
			name:                       "Test Happy Path with relative output directory",
			services:                   []string{customtypes.ENUM_EXPORT_SERVICE_PINGONE_PROTECT},
			outputDir:                  "relative-dir",
			overwriteOutputDirLocation: true,
			changeWorkingDir:           true,
			checkTfFiles:               true,
			pingOneAuthType:            customtypes.PingOneAuthenticationType(customtypes.ENUM_PINGONE_AUTHENTICATION_TYPE_WORKER),
			authStorageType:            customtypes.StorageType(customtypes.ENUM_STORAGE_TYPE_NONE),
		},
		{
			name:                       "Test unreable output directory",
			services:                   []string{customtypes.ENUM_EXPORT_SERVICE_PINGONE_PROTECT},
			outputDir:                  customtypes.String(unreadableDir),
			overwriteOutputDirLocation: true,
			expectedError:              ErrReadOutputDirectory,
		},
		{
			name:                       "Test non-empty output directory without overwrite",
			services:                   []string{customtypes.ENUM_EXPORT_SERVICE_PINGONE_PROTECT},
			outputDir:                  customtypes.String(nonEmptyDir),
			overwriteOutputDirLocation: true,
			expectedError:              ErrOutputDirectoryNotEmpty,
		},
		{
			name:                       "Test non-empty output directory with overwrite",
			services:                   []string{customtypes.ENUM_EXPORT_SERVICE_PINGONE_PROTECT},
			outputDir:                  customtypes.String(nonEmptyDir),
			overwriteOutputDirLocation: true,
			overwriteOnExport:          true,
			checkTfFiles:               true,
			pingOneAuthType:            customtypes.PingOneAuthenticationType(customtypes.ENUM_PINGONE_AUTHENTICATION_TYPE_WORKER),
			authStorageType:            customtypes.StorageType(customtypes.ENUM_STORAGE_TYPE_NONE),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			testutils_koanf.InitKoanfs(t)

			if tc.pingOneAuthType == customtypes.PingOneAuthenticationType(customtypes.ENUM_PINGONE_AUTHENTICATION_TYPE_WORKER) {
				if v := os.Getenv("TEST_PINGONE_WORKER_CLIENT_ID"); v != "" {
					t.Setenv("PINGCLI_PINGONE_WORKER_CLIENT_ID", v)
				}
				if v := os.Getenv("TEST_PINGONE_WORKER_CLIENT_SECRET"); v != "" {
					t.Setenv("PINGCLI_PINGONE_WORKER_CLIENT_SECRET", v)
				}
				// Use worker env ID if present, otherwise default to configured env ID
				if v := os.Getenv("TEST_PINGONE_WORKER_ENVIRONMENT_ID"); v != "" {
					t.Setenv("PINGCLI_PINGONE_WORKER_ENVIRONMENT_ID", v)
				} else if v := os.Getenv("TEST_PINGONE_ENVIRONMENT_ID"); v != "" {
					t.Setenv("PINGCLI_PINGONE_WORKER_ENVIRONMENT_ID", v)
				}
			}

			setupTestCase(t, tc)

			ctx := t.Context()
			if tc.nilContext {
				ctx = nil
			}

			err := RunInternalExport(ctx, "v1.2.3")

			if tc.expectedError != nil {
				require.Error(t, err)
				require.ErrorIs(t, err, tc.expectedError)
			} else {
				require.NoError(t, err)
			}

			if tc.checkTfFiles {
				outputDir, err := profiles.GetOptionValue(options.PlatformExportOutputDirectoryOption)
				require.NoError(t, err)

				files, err := os.ReadDir(outputDir)
				require.NoError(t, err)
				require.NotZero(t, len(files), "Expected non-zero number of files in output directory")

				re := regexp.MustCompile(`^.*\.tf$`)
				for _, file := range files {
					require.False(t, file.IsDir(), "Expected file, got directory: %v", file.Name())
					require.True(t, re.MatchString(file.Name()), "Expected .tf file, got: %v", file.Name())
				}
			}
		})
	}
}

func setupTestCase(t *testing.T, tc testCase) {
	t.Helper()

	if tc.services != nil {
		options.PlatformExportServiceOption.Flag.Changed = true
		options.PlatformExportServiceOption.CobraParamValue = &tc.services
		t.Cleanup(func() {
			options.PlatformExportServiceOption.Flag.Changed = false
			options.PlatformExportServiceOption.CobraParamValue = nil
		})
	}

	if tc.pingOneAuthType != "" {
		// Set runtime PingOne Auth Type (and clear via Cleanup = logout)
		options.PingOneAuthenticationTypeOption.Flag.Changed = true
		options.PingOneAuthenticationTypeOption.CobraParamValue = &tc.pingOneAuthType
		t.Cleanup(func() {
			options.PingOneAuthenticationTypeOption.Flag.Changed = false
			options.PingOneAuthenticationTypeOption.CobraParamValue = nil
		})
	}

	if tc.authStorageType != "" {
		options.AuthStorageOption.Flag.Changed = true
		options.AuthStorageOption.CobraParamValue = &tc.authStorageType
		t.Cleanup(func() {
			options.AuthStorageOption.Flag.Changed = false
			options.AuthStorageOption.CobraParamValue = nil
		})
	}

	if tc.cACertPemFiles != nil {
		require.Contains(t, tc.services, customtypes.ENUM_EXPORT_SERVICE_PINGFEDERATE, "cACertPemFiles is only applicable to PingFederate service export")
		options.PingFederateCACertificatePemFilesOption.Flag.Changed = true
		options.PingFederateCACertificatePemFilesOption.CobraParamValue = &tc.cACertPemFiles
		t.Cleanup(func() {
			options.PingFederateCACertificatePemFilesOption.Flag.Changed = false
			options.PingFederateCACertificatePemFilesOption.CobraParamValue = nil
		})
	}

	if tc.pfAuthType != "" {
		require.Contains(t, tc.services, customtypes.ENUM_EXPORT_SERVICE_PINGFEDERATE, "pfAuthType is only applicable to PingFederate service export")
		options.PingFederateAuthenticationTypeOption.Flag.Changed = true
		options.PingFederateAuthenticationTypeOption.CobraParamValue = &tc.pfAuthType
		t.Cleanup(func() {
			options.PingFederateAuthenticationTypeOption.Flag.Changed = false
			options.PingFederateAuthenticationTypeOption.CobraParamValue = nil
		})
	}

	if tc.pfAccessToken != "" {
		require.Contains(t, tc.services, customtypes.ENUM_EXPORT_SERVICE_PINGFEDERATE, "pfAccessToken is only applicable to PingFederate service export")
		options.PingFederateAccessTokenAuthAccessTokenOption.Flag.Changed = true
		options.PingFederateAccessTokenAuthAccessTokenOption.CobraParamValue = &tc.pfAccessToken
		t.Cleanup(func() {
			options.PingFederateAccessTokenAuthAccessTokenOption.Flag.Changed = false
			options.PingFederateAccessTokenAuthAccessTokenOption.CobraParamValue = nil
		})
	}

	if tc.pfClientId != nil {
		require.Contains(t, tc.services, customtypes.ENUM_EXPORT_SERVICE_PINGFEDERATE, "pfClientId is only applicable to PingFederate service export")
		options.PingFederateClientCredentialsAuthClientIDOption.Flag.Changed = true
		options.PingFederateClientCredentialsAuthClientIDOption.CobraParamValue = tc.pfClientId
		t.Cleanup(func() {
			options.PingFederateClientCredentialsAuthClientIDOption.Flag.Changed = false
			options.PingFederateClientCredentialsAuthClientIDOption.CobraParamValue = nil
		})
	}

	if tc.pfClientSecret != "" {
		require.Contains(t, tc.services, customtypes.ENUM_EXPORT_SERVICE_PINGFEDERATE, "pfClientSecret is only applicable to PingFederate service export")
		options.PingFederateClientCredentialsAuthClientSecretOption.Flag.Changed = true
		options.PingFederateClientCredentialsAuthClientSecretOption.CobraParamValue = &tc.pfClientSecret
		t.Cleanup(func() {
			options.PingFederateClientCredentialsAuthClientSecretOption.Flag.Changed = false
			options.PingFederateClientCredentialsAuthClientSecretOption.CobraParamValue = nil
		})
	}

	if tc.pfTokenURL != "" {
		require.Contains(t, tc.services, customtypes.ENUM_EXPORT_SERVICE_PINGFEDERATE, "pfTokenURL is only applicable to PingFederate service export")
		options.PingFederateClientCredentialsAuthTokenURLOption.Flag.Changed = true
		options.PingFederateClientCredentialsAuthTokenURLOption.CobraParamValue = &tc.pfTokenURL
		t.Cleanup(func() {
			options.PingFederateClientCredentialsAuthTokenURLOption.Flag.Changed = false
			options.PingFederateClientCredentialsAuthTokenURLOption.CobraParamValue = nil
		})
	}

	if tc.overwriteOutputDirLocation {
		options.PlatformExportOutputDirectoryOption.Flag.Changed = true
		options.PlatformExportOutputDirectoryOption.CobraParamValue = &tc.outputDir
		t.Cleanup(func() {
			options.PlatformExportOutputDirectoryOption.Flag.Changed = false
			options.PlatformExportOutputDirectoryOption.CobraParamValue = nil
		})
	}

	if tc.changeWorkingDir {
		originalWd, err := os.Getwd()
		require.NoError(t, err)

		t.Chdir(t.TempDir())

		t.Cleanup(func() {
			t.Chdir(originalWd)
		})
	}

	if tc.overwriteOnExport {
		options.PlatformExportOverwriteOption.Flag.Changed = true
		options.PlatformExportOverwriteOption.CobraParamValue = &tc.overwriteOnExport
		t.Cleanup(func() {
			options.PlatformExportOverwriteOption.Flag.Changed = false
			options.PlatformExportOverwriteOption.CobraParamValue = nil
		})
	}
}

func createCaCertPemFile(t *testing.T, certStr string) *customtypes.StringSlice {
	t.Helper()

	testCACertPemFiles := new(customtypes.StringSlice)

	caCertFile, err := os.CreateTemp(t.TempDir(), "caCert-*.pem")
	require.NoError(t, err)

	_, err = caCertFile.WriteString(certStr)
	require.NoError(t, err)

	err = caCertFile.Close()
	require.NoError(t, err)

	err = testCACertPemFiles.Set(caCertFile.Name())
	require.NoError(t, err)

	return testCACertPemFiles
}

func createGoldenCACertPemFile(t *testing.T) *customtypes.StringSlice {
	t.Helper()

	certStr, err := testutils.CreateX509Certificate()
	require.NoError(t, err)

	return createCaCertPemFile(t, certStr)
}

func createMalformedCACertPemFile(t *testing.T) *customtypes.StringSlice {
	t.Helper()

	return createCaCertPemFile(t, "malformed-cert")
}

func createUnwriteableDir(t *testing.T) string {
	t.Helper()

	dir := t.TempDir()
	err := os.Chmod(dir, 0400) // read-only
	require.NoError(t, err)

	return fmt.Sprintf("%s/subdir", dir)
}

func createUnreadableDir(t *testing.T) string {
	t.Helper()

	dir := t.TempDir()
	err := os.Chmod(dir, 0000) // no permissions
	require.NoError(t, err)

	return dir
}

func createNonEmptyDir(t *testing.T) string {
	t.Helper()

	dir := t.TempDir()
	file, err := os.CreateTemp(dir, "file-*.tf") // #nosec G304
	require.NoError(t, err)

	err = file.Close()
	require.NoError(t, err)

	return dir
}
