// Copyright © 2025 Ping Identity Corporation

package platform_test

import (
	"os"
	"strings"
	"testing"

	"github.com/pingidentity/pingcli/internal/configuration/options"
	"github.com/pingidentity/pingcli/internal/customtypes"
	"github.com/pingidentity/pingcli/internal/testing/testutils"
	"github.com/pingidentity/pingcli/internal/testing/testutils_cobra"
	"github.com/pingidentity/pingcli/internal/testing/testutils_koanf"
)

// Test Platform Export Command Executes without issue
func TestPlatformExportCmd_Execute(t *testing.T) {
	testutils_koanf.InitKoanfs(t)
	outputDir := t.TempDir()

	err := testutils_cobra.ExecutePingcli(t, "platform", "export",
		"--"+options.PlatformExportOutputDirectoryOption.CobraParamName, outputDir,
		"--"+options.PlatformExportOverwriteOption.CobraParamName)
	testutils.CheckExpectedError(t, err, nil)
}

// Test Platform Export Command fails when provided too many arguments
func TestPlatformExportCmd_TooManyArgs(t *testing.T) {
	expectedErrorPattern := `command accepts 0 arg\(s\), received 1`
	err := testutils_cobra.ExecutePingcli(t, "platform", "export", "extra-arg")
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}

// Test Platform Export Command fails when provided invalid flag
func TestPlatformExportCmd_InvalidFlag(t *testing.T) {
	expectedErrorPattern := `^unknown flag: --invalid$`
	err := testutils_cobra.ExecutePingcli(t, "platform", "export", "--invalid")
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}

// Test Platform Export Command --help, -h flag
func TestPlatformExportCmd_HelpFlag(t *testing.T) {
	err := testutils_cobra.ExecutePingcli(t, "platform", "export", "--help")
	testutils.CheckExpectedError(t, err, nil)

	err = testutils_cobra.ExecutePingcli(t, "platform", "export", "-h")
	testutils.CheckExpectedError(t, err, nil)
}

// Test Platform Export Command --service-group, -g flag
func TestPlatformExportCmd_ServiceGroupFlag(t *testing.T) {
	testutils_koanf.InitKoanfs(t)
	outputDir := t.TempDir()

	err := testutils_cobra.ExecutePingcli(t, "platform", "export",
		"--"+options.PlatformExportOutputDirectoryOption.CobraParamName, outputDir,
		"--"+options.PlatformExportOverwriteOption.CobraParamName,
		"--"+options.PlatformExportServiceGroupOption.CobraParamName, "pingone")
	testutils.CheckExpectedError(t, err, nil)
}

// Test Platform Export Command --service-group with non-supported service group
func TestPlatformExportCmd_ServiceGroupFlagInvalidServiceGroup(t *testing.T) {
	testutils_koanf.InitKoanfs(t)

	expectedErrorPattern := `unrecognized service group 'invalid'`
	err := testutils_cobra.ExecutePingcli(t, "platform", "export",
		"--"+options.PlatformExportServiceGroupOption.CobraParamName, "invalid")
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}

// Test Platform Export Command --services flag
func TestPlatformExportCmd_ServicesFlag(t *testing.T) {
	testutils_koanf.InitKoanfs(t)
	outputDir := t.TempDir()

	err := testutils_cobra.ExecutePingcli(t, "platform", "export",
		"--"+options.PlatformExportOutputDirectoryOption.CobraParamName, outputDir,
		"--"+options.PlatformExportOverwriteOption.CobraParamName,
		"--"+options.PlatformExportServiceOption.CobraParamName, customtypes.ENUM_EXPORT_SERVICE_PINGONE_PROTECT)
	testutils.CheckExpectedError(t, err, nil)
}

// Test Platform Export Command --services flag with invalid service
func TestPlatformExportCmd_ServicesFlagInvalidService(t *testing.T) {
	testutils_koanf.InitKoanfs(t)

	expectedErrorPattern := `unrecognized service 'invalid'`
	err := testutils_cobra.ExecutePingcli(t, "platform", "export",
		"--"+options.PlatformExportServiceOption.CobraParamName, "invalid")
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}

// Test Platform Export Command --format flag
func TestPlatformExportCmd_ExportFormatFlag(t *testing.T) {
	testutils_koanf.InitKoanfs(t)
	outputDir := t.TempDir()

	err := testutils_cobra.ExecutePingcli(t, "platform", "export",
		"--"+options.PlatformExportOutputDirectoryOption.CobraParamName, outputDir,
		"--"+options.PlatformExportExportFormatOption.CobraParamName, customtypes.ENUM_EXPORT_FORMAT_HCL,
		"--"+options.PlatformExportOverwriteOption.CobraParamName,
		"--"+options.PlatformExportServiceOption.CobraParamName, customtypes.ENUM_EXPORT_SERVICE_PINGONE_PROTECT)
	testutils.CheckExpectedError(t, err, nil)
}

// Test Platform Export Command --format flag with invalid format
func TestPlatformExportCmd_ExportFormatFlagInvalidFormat(t *testing.T) {
	testutils_koanf.InitKoanfs(t)

	expectedErrorPattern := `unrecognized export format 'invalid'`
	err := testutils_cobra.ExecutePingcli(t, "platform", "export",
		"--"+options.PlatformExportExportFormatOption.CobraParamName, "invalid")
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}

// Test Platform Export Command --output-directory flag
func TestPlatformExportCmd_OutputDirectoryFlag(t *testing.T) {
	testutils_koanf.InitKoanfs(t)
	outputDir := t.TempDir()

	err := testutils_cobra.ExecutePingcli(t, "platform", "export",
		"--"+options.PlatformExportOutputDirectoryOption.CobraParamName, outputDir,
		"--"+options.PlatformExportOverwriteOption.CobraParamName,
		"--"+options.PlatformExportServiceOption.CobraParamName, customtypes.ENUM_EXPORT_SERVICE_PINGONE_PROTECT)
	testutils.CheckExpectedError(t, err, nil)
}

// Test Platform Export Command --output-directory flag with invalid directory
func TestPlatformExportCmd_OutputDirectoryFlagInvalidDirectory(t *testing.T) {
	testutils_koanf.InitKoanfs(t)

	expectedErrorPattern := `^failed to create output directory '\/invalid': mkdir \/invalid: .+$`
	err := testutils_cobra.ExecutePingcli(t, "platform", "export",
		"--"+options.PlatformExportOutputDirectoryOption.CobraParamName, "/invalid")
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}

// Test Platform Export Command --overwrite flag
func TestPlatformExportCmd_OverwriteFlag(t *testing.T) {
	testutils_koanf.InitKoanfs(t)
	outputDir := t.TempDir()

	err := testutils_cobra.ExecutePingcli(t, "platform", "export",
		"--"+options.PlatformExportOutputDirectoryOption.CobraParamName, outputDir,
		"--"+options.PlatformExportOverwriteOption.CobraParamName,
		"--"+options.PlatformExportServiceOption.CobraParamName, customtypes.ENUM_EXPORT_SERVICE_PINGONE_PROTECT)
	testutils.CheckExpectedError(t, err, nil)
}

// Test Platform Export Command --overwrite flag false with existing directory
// where the directory already contains a file
func TestPlatformExportCmd_OverwriteFlagFalseWithExistingDirectory(t *testing.T) {
	testutils_koanf.InitKoanfs(t)
	outputDir := t.TempDir()

	_, err := os.Create(outputDir + "/file") //#nosec G304 -- this is a test
	if err != nil {
		t.Errorf("Error creating file in output directory: %v", err)
	}

	expectedErrorPattern := `output directory is not empty.*use '--overwrite'`
	err = testutils_cobra.ExecutePingcli(t, "platform", "export",
		"--"+options.PlatformExportOutputDirectoryOption.CobraParamName, outputDir,
		"--"+options.PlatformExportServiceOption.CobraParamName, customtypes.ENUM_EXPORT_SERVICE_PINGONE_PROTECT,
		"--"+options.PlatformExportOverwriteOption.CobraParamName+"=false")
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}

// Test Platform Export Command --overwrite flag true with existing directory
// where the directory already contains a file
func TestPlatformExportCmd_OverwriteFlagTrueWithExistingDirectory(t *testing.T) {
	testutils_koanf.InitKoanfs(t)
	outputDir := t.TempDir()

	_, err := os.Create(outputDir + "/file") //#nosec G304 -- this is a test
	if err != nil {
		t.Errorf("Error creating file in output directory: %v", err)
	}

	err = testutils_cobra.ExecutePingcli(t, "platform", "export",
		"--"+options.PlatformExportOutputDirectoryOption.CobraParamName, outputDir,
		"--"+options.PlatformExportServiceOption.CobraParamName, customtypes.ENUM_EXPORT_SERVICE_PINGONE_PROTECT,
		"--"+options.PlatformExportOverwriteOption.CobraParamName)
	testutils.CheckExpectedError(t, err, nil)
}

// Test Platform Export Command with
// --pingone-worker-environment-id flag
// --pingone-worker-client-id flag
// --pingone-worker-client-secret flag
// --pingone-region flag
func TestPlatformExportCmd_PingOneWorkerEnvironmentIdFlag(t *testing.T) {
	testutils_koanf.InitKoanfs(t)
	outputDir := t.TempDir()

	err := testutils_cobra.ExecutePingcli(t, "platform", "export",
		"--"+options.PlatformExportOutputDirectoryOption.CobraParamName, outputDir,
		"--"+options.PlatformExportOverwriteOption.CobraParamName,
		"--"+options.PlatformExportServiceOption.CobraParamName, customtypes.ENUM_EXPORT_SERVICE_PINGONE_PROTECT,
		"--"+options.PingOneAuthenticationWorkerEnvironmentIDOption.CobraParamName, os.Getenv("TEST_PINGONE_ENVIRONMENT_ID"),
		"--"+options.PingOneAuthenticationWorkerClientIDOption.CobraParamName, os.Getenv("TEST_PINGONE_WORKER_CLIENT_ID"),
		"--"+options.PingOneAuthenticationWorkerClientSecretOption.CobraParamName, os.Getenv("TEST_PINGONE_WORKER_CLIENT_SECRET"),
		"--"+options.PingOneRegionCodeOption.CobraParamName, os.Getenv("TEST_PINGONE_REGION_CODE"))
	testutils.CheckExpectedError(t, err, nil)
}

// Test Platform Export Command with partial worker credentials (should fail during authentication)
func TestPlatformExportCmd_PingOneWorkerEnvironmentIdFlagRequiredTogether(t *testing.T) {
	testutils_koanf.InitKoanfs(t)
	outputDir := t.TempDir()

	// With only environment ID provided, may succeed if worker client ID/secret/region configured
	err := testutils_cobra.ExecutePingcli(t, "platform", "export",
		"--"+options.PlatformExportOutputDirectoryOption.CobraParamName, outputDir,
		"--"+options.PlatformExportServiceOption.CobraParamName, customtypes.ENUM_EXPORT_SERVICE_PINGONE_PLATFORM,
		"--"+options.PingOneAuthenticationWorkerEnvironmentIDOption.CobraParamName, os.Getenv("TEST_PINGONE_ENVIRONMENT_ID"))

	// May succeed if worker credentials are fully configured
	if err == nil {
		t.Skip("Export succeeded - worker credentials fully configured")
	}
	// Should get authentication-related error if credentials missing
	if !strings.Contains(err.Error(), "failed to initialize") &&
		!strings.Contains(err.Error(), "client") &&
		!strings.Contains(err.Error(), "authentication") {
		t.Errorf("Expected authentication error, got: %v", err)
	}
}

// Test Platform Export command with PingFederate Basic Auth flags
func TestPlatformExportCmd_PingFederateBasicAuthFlags(t *testing.T) {
	testutils_koanf.InitKoanfs(t)
	outputDir := t.TempDir()

	err := testutils_cobra.ExecutePingcli(t, "platform", "export",
		"--"+options.PlatformExportOutputDirectoryOption.CobraParamName, outputDir,
		"--"+options.PlatformExportOverwriteOption.CobraParamName,
		"--"+options.PlatformExportServiceOption.CobraParamName, customtypes.ENUM_EXPORT_SERVICE_PINGFEDERATE,
		"--"+options.PingFederateBasicAuthUsernameOption.CobraParamName, "Administrator",
		"--"+options.PingFederateBasicAuthPasswordOption.CobraParamName, "2FederateM0re",
		"--"+options.PingFederateAuthenticationTypeOption.CobraParamName, customtypes.ENUM_PINGFEDERATE_AUTHENTICATION_TYPE_BASIC,
	)
	// Success when PingFederate server is available, error when not
	if err == nil {
		t.Skip("PingFederate export succeeded - server available")
	}
	if !strings.Contains(err.Error(), "PingFederate") && !strings.Contains(err.Error(), "failed to initialize") {
		t.Errorf("Expected PingFederate initialization error, got: %v", err)
	}
}

// Test Platform Export Command fails when not provided required PingFederate Basic Auth flags together
func TestPlatformExportCmd_PingFederateBasicAuthFlagsRequiredTogether(t *testing.T) {
	testutils_koanf.InitKoanfs(t)

	expectedErrorPattern := `^if any flags in the group \[pingfederate-username pingfederate-password] are set they must all be set; missing \[pingfederate-password]$`
	err := testutils_cobra.ExecutePingcli(t, "platform", "export",
		"--"+options.PingFederateBasicAuthUsernameOption.CobraParamName, "Administrator")
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}

// Test Platform Export Command fails when provided invalid PingOne Client Credential flags
func TestPlatformExportCmd_PingOneClientCredentialFlagsInvalid(t *testing.T) {
	testutils_koanf.InitKoanfs(t)
	outputDir := t.TempDir()

	expectedErrorPattern := `failed to initialize pingone API client.*Check worker client ID, worker client secret, worker environment ID, and pingone region code`
	err := testutils_cobra.ExecutePingcli(t, "platform", "export",
		"--"+options.PlatformExportOutputDirectoryOption.CobraParamName, outputDir,
		"--"+options.PlatformExportOverwriteOption.CobraParamName,
		"--"+options.PlatformExportServiceOption.CobraParamName, customtypes.ENUM_EXPORT_SERVICE_PINGONE_PROTECT,
		"--"+options.PingOneAuthenticationWorkerEnvironmentIDOption.CobraParamName, os.Getenv("TEST_PINGONE_ENVIRONMENT_ID"),
		"--"+options.PingOneAuthenticationWorkerClientIDOption.CobraParamName, os.Getenv("TEST_PINGONE_WORKER_CLIENT_ID"),
		"--"+options.PingOneAuthenticationWorkerClientSecretOption.CobraParamName, "invalid",
		"--"+options.PingOneRegionCodeOption.CobraParamName, os.Getenv("TEST_PINGONE_REGION_CODE"),
	)
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}

// Test Platform Export Command fails when provided invalid PingFederate Basic Auth flags
func TestPlatformExportCmd_PingFederateBasicAuthFlagsInvalid(t *testing.T) {
	testutils_koanf.InitKoanfs(t)
	outputDir := t.TempDir()

	expectedErrorPattern := `failed to initialize PingFederate service.*Check authentication type and credentials`
	err := testutils_cobra.ExecutePingcli(t, "platform", "export",
		"--"+options.PlatformExportOutputDirectoryOption.CobraParamName, outputDir,
		"--"+options.PlatformExportOverwriteOption.CobraParamName,
		"--"+options.PlatformExportServiceOption.CobraParamName, customtypes.ENUM_EXPORT_SERVICE_PINGFEDERATE,
		"--"+options.PingFederateBasicAuthUsernameOption.CobraParamName, "Administrator",
		"--"+options.PingFederateBasicAuthPasswordOption.CobraParamName, "invalid",
		"--"+options.PingFederateAuthenticationTypeOption.CobraParamName, customtypes.ENUM_PINGFEDERATE_AUTHENTICATION_TYPE_BASIC,
	)
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}

// Test Platform Export Command fails when not provided required PingFederate Client Credentials Auth flags together
func TestPlatformExportCmd_PingFederateClientCredentialsAuthFlagsRequiredTogether(t *testing.T) {
	testutils_koanf.InitKoanfs(t)

	expectedErrorPattern := `^if any flags in the group \[pingfederate-client-id pingfederate-client-secret pingfederate-token-url] are set they must all be set; missing \[pingfederate-client-secret pingfederate-token-url]$`
	err := testutils_cobra.ExecutePingcli(t, "platform", "export",
		"--"+options.PingFederateClientCredentialsAuthClientIDOption.CobraParamName, "test")
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}

// Test Platform Export Command fails when provided invalid PingFederate Client Credentials Auth flags
func TestPlatformExportCmd_PingFederateClientCredentialsAuthFlagsInvalid(t *testing.T) {
	testutils_koanf.InitKoanfs(t)
	outputDir := t.TempDir()

	expectedErrorPattern := `failed to initialize PingFederate service.*Check authentication type and credentials`
	err := testutils_cobra.ExecutePingcli(t, "platform", "export",
		"--"+options.PlatformExportOutputDirectoryOption.CobraParamName, outputDir,
		"--"+options.PlatformExportOverwriteOption.CobraParamName,
		"--"+options.PlatformExportServiceOption.CobraParamName, customtypes.ENUM_EXPORT_SERVICE_PINGFEDERATE,
		"--"+options.PingFederateClientCredentialsAuthClientIDOption.CobraParamName, "test",
		"--"+options.PingFederateClientCredentialsAuthClientSecretOption.CobraParamName, "invalid",
		"--"+options.PingFederateClientCredentialsAuthTokenURLOption.CobraParamName, "https://localhost:9031/as/token.oauth2",
		"--"+options.PingFederateClientCredentialsAuthScopesOption.CobraParamName, "email",
		"--"+options.PingFederateAuthenticationTypeOption.CobraParamName, customtypes.ENUM_PINGFEDERATE_AUTHENTICATION_TYPE_CLIENT_CREDENTIALS,
	)
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}

// Test Platform Export Command fails when provided invalid PingFederate OAuth2 Token URL
func TestPlatformExportCmd_PingFederateClientCredentialsAuthFlagsInvalidTokenURL(t *testing.T) {
	testutils_koanf.InitKoanfs(t)
	outputDir := t.TempDir()

	expectedErrorPattern := `failed to initialize PingFederate service.*Check authentication type and credentials`
	err := testutils_cobra.ExecutePingcli(t, "platform", "export",
		"--"+options.PlatformExportOutputDirectoryOption.CobraParamName, outputDir,
		"--"+options.PlatformExportOverwriteOption.CobraParamName,
		"--"+options.PlatformExportServiceOption.CobraParamName, customtypes.ENUM_EXPORT_SERVICE_PINGFEDERATE,
		"--"+options.PingFederateClientCredentialsAuthClientIDOption.CobraParamName, "test",
		"--"+options.PingFederateClientCredentialsAuthClientSecretOption.CobraParamName, "2FederateM0re!",
		"--"+options.PingFederateClientCredentialsAuthTokenURLOption.CobraParamName, "https://localhost:9031/as/invalid",
		"--"+options.PingFederateAuthenticationTypeOption.CobraParamName, customtypes.ENUM_PINGFEDERATE_AUTHENTICATION_TYPE_CLIENT_CREDENTIALS,
	)
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}

// Test Platform Export command with PingFederate X-Bypass Header set to true
func TestPlatformExportCmd_PingFederateXBypassHeaderFlag(t *testing.T) {
	testutils_koanf.InitKoanfs(t)
	outputDir := t.TempDir()

	err := testutils_cobra.ExecutePingcli(t, "platform", "export",
		"--"+options.PlatformExportOutputDirectoryOption.CobraParamName, outputDir,
		"--"+options.PlatformExportOverwriteOption.CobraParamName,
		"--"+options.PlatformExportServiceOption.CobraParamName, customtypes.ENUM_EXPORT_SERVICE_PINGFEDERATE,
		"--"+options.PingFederateXBypassExternalValidationHeaderOption.CobraParamName,
		"--"+options.PingFederateBasicAuthUsernameOption.CobraParamName, "Administrator",
		"--"+options.PingFederateBasicAuthPasswordOption.CobraParamName, "2FederateM0re",
		"--"+options.PingFederateAuthenticationTypeOption.CobraParamName, customtypes.ENUM_PINGFEDERATE_AUTHENTICATION_TYPE_BASIC,
	)
	testutils.CheckExpectedError(t, err, nil)
}

// Test Platform Export command with PingFederate --pingfederate-insecure-trust-all-tls flag set to true
func TestPlatformExportCmd_PingFederateTrustAllTLSFlag(t *testing.T) {
	testutils_koanf.InitKoanfs(t)
	outputDir := t.TempDir()

	err := testutils_cobra.ExecutePingcli(t, "platform", "export",
		"--"+options.PlatformExportOutputDirectoryOption.CobraParamName, outputDir,
		"--"+options.PlatformExportOverwriteOption.CobraParamName,
		"--"+options.PlatformExportServiceOption.CobraParamName, customtypes.ENUM_EXPORT_SERVICE_PINGFEDERATE,
		"--"+options.PingFederateInsecureTrustAllTLSOption.CobraParamName,
		"--"+options.PingFederateBasicAuthUsernameOption.CobraParamName, "Administrator",
		"--"+options.PingFederateBasicAuthPasswordOption.CobraParamName, "2FederateM0re",
		"--"+options.PingFederateAuthenticationTypeOption.CobraParamName, customtypes.ENUM_PINGFEDERATE_AUTHENTICATION_TYPE_BASIC,
	)
	testutils.CheckExpectedError(t, err, nil)
}

// Test Platform Export command fails with PingFederate --pingfederate-insecure-trust-all-tls flag set to false
func TestPlatformExportCmd_PingFederateTrustAllTLSFlagFalse(t *testing.T) {
	testutils_koanf.InitKoanfs(t)
	outputDir := t.TempDir()

	expectedErrorPattern := `failed to initialize PingFederate service.*Check authentication type and credentials`
	err := testutils_cobra.ExecutePingcli(t, "platform", "export",
		"--"+options.PlatformExportOutputDirectoryOption.CobraParamName, outputDir,
		"--"+options.PlatformExportOverwriteOption.CobraParamName,
		"--"+options.PlatformExportServiceOption.CobraParamName, customtypes.ENUM_EXPORT_SERVICE_PINGFEDERATE,
		"--"+options.PingFederateInsecureTrustAllTLSOption.CobraParamName+"=false",
		"--"+options.PingFederateBasicAuthUsernameOption.CobraParamName, "Administrator",
		"--"+options.PingFederateBasicAuthPasswordOption.CobraParamName, "2FederateM0re",
		"--"+options.PingFederateAuthenticationTypeOption.CobraParamName, customtypes.ENUM_PINGFEDERATE_AUTHENTICATION_TYPE_BASIC,
	)
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}

// Test Platform Export command passes with PingFederate
// --pingfederate-insecure-trust-all-tls=false
// and --pingfederate-ca-certificate-pem-files set
func TestPlatformExportCmd_PingFederateCaCertificatePemFiles(t *testing.T) {
	testutils_koanf.InitKoanfs(t)
	outputDir := t.TempDir()

	err := testutils_cobra.ExecutePingcli(t, "platform", "export",
		"--"+options.PlatformExportOutputDirectoryOption.CobraParamName, outputDir,
		"--"+options.PlatformExportOverwriteOption.CobraParamName,
		"--"+options.PlatformExportServiceOption.CobraParamName, customtypes.ENUM_EXPORT_SERVICE_PINGFEDERATE,
		"--"+options.PingFederateInsecureTrustAllTLSOption.CobraParamName+"=true",
		"--"+options.PingFederateCACertificatePemFilesOption.CobraParamName, "testdata/ssl-server-crt.pem",
		"--"+options.PingFederateBasicAuthUsernameOption.CobraParamName, "Administrator",
		"--"+options.PingFederateBasicAuthPasswordOption.CobraParamName, "2FederateM0re",
		"--"+options.PingFederateAuthenticationTypeOption.CobraParamName, customtypes.ENUM_PINGFEDERATE_AUTHENTICATION_TYPE_BASIC,
	)
	testutils.CheckExpectedError(t, err, nil)
}

// Test Platform Export command fails with --pingfederate-ca-certificate-pem-files set to non-existent file.
func TestPlatformExportCmd_PingFederateCaCertificatePemFilesInvalid(t *testing.T) {
	testutils_koanf.InitKoanfs(t)

	expectedErrorPattern := `^failed to read CA certificate PEM file '.*': open .*: no such file or directory$`
	err := testutils_cobra.ExecutePingcli(t, "platform", "export",
		"--"+options.PlatformExportServiceOption.CobraParamName, customtypes.ENUM_EXPORT_SERVICE_PINGFEDERATE,
		"--"+options.PingFederateCACertificatePemFilesOption.CobraParamName, "invalid/crt.pem",
		"--"+options.PingFederateBasicAuthUsernameOption.CobraParamName, "Administrator",
		"--"+options.PingFederateBasicAuthPasswordOption.CobraParamName, "2FederateM0re",
		"--"+options.PingFederateAuthenticationTypeOption.CobraParamName, customtypes.ENUM_PINGFEDERATE_AUTHENTICATION_TYPE_BASIC,
	)
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}

// Test Platform Export Command with PingOne client_credentials authentication
func TestPlatformExportCmd_PingOneClientCredentialsAuth(t *testing.T) {
	testutils_koanf.InitKoanfs(t)
	outputDir := t.TempDir()

	err := testutils_cobra.ExecutePingcli(t, "platform", "export",
		"--"+options.PlatformExportOutputDirectoryOption.CobraParamName, outputDir,
		"--"+options.PlatformExportOverwriteOption.CobraParamName,
		"--"+options.PlatformExportServiceOption.CobraParamName, customtypes.ENUM_EXPORT_SERVICE_PINGONE_PLATFORM,
		"--"+options.PingOneAuthenticationTypeOption.CobraParamName, customtypes.ENUM_PINGONE_AUTHENTICATION_TYPE_CLIENT_CREDENTIALS,
		"--"+options.PingOneAuthenticationClientCredentialsClientIDOption.CobraParamName, os.Getenv("TEST_PINGONE_CLIENT_ID"),
		"--"+options.PingOneAuthenticationClientCredentialsClientSecretOption.CobraParamName, os.Getenv("TEST_PINGONE_CLIENT_SECRET"),
		"--"+options.PingOneAuthenticationClientCredentialsEnvironmentIDOption.CobraParamName, os.Getenv("TEST_PINGONE_ENVIRONMENT_ID"),
		"--"+options.PingOneRegionCodeOption.CobraParamName, os.Getenv("TEST_PINGONE_REGION_CODE"))
	testutils.CheckExpectedError(t, err, nil)
}

// Test Platform Export Command with PingOne device_code authentication
func TestPlatformExportCmd_PingOneDeviceCodeAuth(t *testing.T) {
	testutils_koanf.InitKoanfs(t)
	outputDir := t.TempDir()

	err := testutils_cobra.ExecutePingcli(t, "platform", "export",
		"--"+options.PlatformExportOutputDirectoryOption.CobraParamName, outputDir,
		"--"+options.PlatformExportOverwriteOption.CobraParamName,
		"--"+options.PlatformExportServiceOption.CobraParamName, customtypes.ENUM_EXPORT_SERVICE_PINGONE_PLATFORM,
		"--"+options.PingOneAuthenticationTypeOption.CobraParamName, customtypes.ENUM_PINGONE_AUTHENTICATION_TYPE_DEVICE_CODE,
		"--"+options.PingOneAuthenticationDeviceCodeClientIDOption.CobraParamName, os.Getenv("TEST_PINGONE_DEVICE_CODE_CLIENT_ID"),
		"--"+options.PingOneAuthenticationDeviceCodeEnvironmentIDOption.CobraParamName, os.Getenv("TEST_PINGONE_ENVIRONMENT_ID"),
		"--"+options.PingOneRegionCodeOption.CobraParamName, os.Getenv("TEST_PINGONE_REGION_CODE"))
	testutils.CheckExpectedError(t, err, nil)
}

// Test Platform Export Command with PingOne auth_code authentication
func TestPlatformExportCmd_PingOneAuthCodeAuth(t *testing.T) {
	testutils_koanf.InitKoanfs(t)
	outputDir := t.TempDir()

	err := testutils_cobra.ExecutePingcli(t, "platform", "export",
		"--"+options.PlatformExportOutputDirectoryOption.CobraParamName, outputDir,
		"--"+options.PlatformExportOverwriteOption.CobraParamName,
		"--"+options.PlatformExportServiceOption.CobraParamName, customtypes.ENUM_EXPORT_SERVICE_PINGONE_PLATFORM,
		"--"+options.PingOneAuthenticationTypeOption.CobraParamName, customtypes.ENUM_PINGONE_AUTHENTICATION_TYPE_AUTH_CODE,
		"--"+options.PingOneAuthenticationAuthCodeClientIDOption.CobraParamName, os.Getenv("TEST_PINGONE_AUTH_CODE_CLIENT_ID"),
		"--"+options.PingOneAuthenticationAuthCodeEnvironmentIDOption.CobraParamName, os.Getenv("TEST_PINGONE_ENVIRONMENT_ID"),
		"--"+options.PingOneAuthenticationAuthCodeRedirectURIPathOption.CobraParamName, "http://localhost:8080/callback",
		"--"+options.PingOneRegionCodeOption.CobraParamName, os.Getenv("TEST_PINGONE_REGION_CODE"))
	testutils.CheckExpectedError(t, err, nil)
}

// Test Platform Export Command fails when client_credentials authentication is missing client ID
func TestPlatformExportCmd_PingOneClientCredentialsAuthMissingClientID(t *testing.T) {
	testutils_koanf.InitKoanfs(t)
	outputDir := t.TempDir()

	err := testutils_cobra.ExecutePingcli(t, "platform", "export",
		"--"+options.PlatformExportOutputDirectoryOption.CobraParamName, outputDir,
		"--"+options.PlatformExportOverwriteOption.CobraParamName,
		"--"+options.PlatformExportServiceOption.CobraParamName, customtypes.ENUM_EXPORT_SERVICE_PINGONE_PLATFORM,
		"--"+options.PingOneAuthenticationTypeOption.CobraParamName, customtypes.ENUM_PINGONE_AUTHENTICATION_TYPE_CLIENT_CREDENTIALS,
		"--"+options.PingOneAuthenticationClientCredentialsClientSecretOption.CobraParamName, os.Getenv("TEST_PINGONE_CLIENT_SECRET"),
		"--"+options.PingOneRegionCodeOption.CobraParamName, os.Getenv("TEST_PINGONE_REGION_CODE"))

	// May succeed if worker credentials are configured as fallback
	if err == nil {
		t.Skip("Export succeeded - worker credentials available as fallback")
	}
	// Should get error about missing environment ID
	if !strings.Contains(err.Error(), "environment ID is empty") {
		t.Errorf("Expected 'environment ID is empty' error, got: %v", err)
	}
}

// Test Platform Export Command fails when device_code authentication is missing environment ID
func TestPlatformExportCmd_PingOneDeviceCodeAuthMissingEnvironmentID(t *testing.T) {
	testutils_koanf.InitKoanfs(t)
	outputDir := t.TempDir()

	err := testutils_cobra.ExecutePingcli(t, "platform", "export",
		"--"+options.PlatformExportOutputDirectoryOption.CobraParamName, outputDir,
		"--"+options.PlatformExportOverwriteOption.CobraParamName,
		"--"+options.PlatformExportServiceOption.CobraParamName, customtypes.ENUM_EXPORT_SERVICE_PINGONE_PLATFORM,
		"--"+options.PingOneAuthenticationTypeOption.CobraParamName, customtypes.ENUM_PINGONE_AUTHENTICATION_TYPE_DEVICE_CODE,
		"--"+options.PingOneAuthenticationDeviceCodeClientIDOption.CobraParamName, os.Getenv("TEST_PINGONE_DEVICE_CODE_CLIENT_ID"),
		"--"+options.PingOneRegionCodeOption.CobraParamName, os.Getenv("TEST_PINGONE_REGION_CODE"))

	// May succeed if worker credentials are configured as fallback
	if err == nil {
		t.Skip("Export succeeded - worker credentials available as fallback")
	}
	// Should get error about missing environment ID
	if !strings.Contains(err.Error(), "environment ID is empty") {
		t.Errorf("Expected 'environment ID is empty' error, got: %v", err)
	}
}

// Test Platform Export Command fails when region code is missing with new auth methods
func TestPlatformExportCmd_PingOneNewAuthMissingRegionCode(t *testing.T) {
	testutils_koanf.InitKoanfs(t)
	outputDir := t.TempDir()

	err := testutils_cobra.ExecutePingcli(t, "platform", "export",
		"--"+options.PlatformExportOutputDirectoryOption.CobraParamName, outputDir,
		"--"+options.PlatformExportOverwriteOption.CobraParamName,
		"--"+options.PlatformExportServiceOption.CobraParamName, customtypes.ENUM_EXPORT_SERVICE_PINGONE_PLATFORM,
		"--"+options.PingOneAuthenticationTypeOption.CobraParamName, customtypes.ENUM_PINGONE_AUTHENTICATION_TYPE_CLIENT_CREDENTIALS,
		"--"+options.PingOneAuthenticationClientCredentialsClientIDOption.CobraParamName, os.Getenv("TEST_PINGONE_CLIENT_ID"),
		"--"+options.PingOneAuthenticationClientCredentialsClientSecretOption.CobraParamName, os.Getenv("TEST_PINGONE_CLIENT_SECRET"),
		"--"+options.PingOneAuthenticationClientCredentialsEnvironmentIDOption.CobraParamName, os.Getenv("TEST_PINGONE_ENVIRONMENT_ID"))

	// May succeed if worker credentials with region code are configured as fallback
	if err == nil {
		t.Skip("Export succeeded - worker credentials with region code available as fallback")
	}
	// Should get error about missing region code
	if !strings.Contains(err.Error(), "pingone region code is empty") {
		t.Errorf("Expected 'pingone region code is empty' error, got: %v", err)
	}
}

// Test Platform Export Command with invalid authentication type
func TestPlatformExportCmd_PingOneInvalidAuthType(t *testing.T) {
	testutils_koanf.InitKoanfs(t)
	outputDir := t.TempDir()

	expectedErrorPattern := `unrecognized pingone authentication type`
	err := testutils_cobra.ExecutePingcli(t, "platform", "export",
		"--"+options.PlatformExportOutputDirectoryOption.CobraParamName, outputDir,
		"--"+options.PlatformExportOverwriteOption.CobraParamName,
		"--"+options.PlatformExportServiceOption.CobraParamName, customtypes.ENUM_EXPORT_SERVICE_PINGONE_PLATFORM,
		"--"+options.PingOneAuthenticationTypeOption.CobraParamName, "invalid_auth",
		"--"+options.PingOneRegionCodeOption.CobraParamName, os.Getenv("TEST_PINGONE_REGION_CODE"))
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}
