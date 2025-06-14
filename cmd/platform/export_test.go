// Copyright © 2025 Ping Identity Corporation

package platform_test

import (
	"os"
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
	expectedErrorPattern := `^failed to execute 'pingcli platform export': command accepts 0 arg\(s\), received 1$`
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

	expectedErrorPattern := `^invalid argument ".*" for "-g, --service-group" flag: unrecognized service group '.*'\. Must be one of: .*$`
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

	expectedErrorPattern := `^invalid argument ".*" for "-s, --services" flag: failed to set ExportServices: Invalid service: .*\. Allowed services: .*$`
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

	expectedErrorPattern := `^invalid argument ".*" for "-f, --format" flag: unrecognized export format '.*'\. Must be one of: .*$`
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

	expectedErrorPattern := `^output directory '[A-Za-z0-9_\-\/]+' is not empty\. Use --overwrite to overwrite existing export data$`
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

// Test Platform Export Command fails when not provided required pingone flags together
func TestPlatformExportCmd_PingOneWorkerEnvironmentIdFlagRequiredTogether(t *testing.T) {
	testutils_koanf.InitKoanfs(t)

	expectedErrorPattern := `^if any flags in the group \[pingone-worker-environment-id pingone-worker-client-id pingone-worker-client-secret pingone-region-code] are set they must all be set; missing \[pingone-region-code pingone-worker-client-id pingone-worker-client-secret]$`
	err := testutils_cobra.ExecutePingcli(t, "platform", "export",
		"--"+options.PingOneAuthenticationWorkerEnvironmentIDOption.CobraParamName, os.Getenv("TEST_PINGONE_ENVIRONMENT_ID"))
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
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
	testutils.CheckExpectedError(t, err, nil)
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

	expectedErrorPattern := `^failed to initialize pingone API client\. Check worker client ID, worker client secret, worker environment ID, and pingone region code configuration values\. oauth2: \"invalid_client\" \"Request denied: Unsupported authentication method \(Correlation ID: .*\)\"$`
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

	expectedErrorPattern := `^failed to initialize PingFederate Go Client. Check authentication type and credentials$`
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

	expectedErrorPattern := `^failed to initialize PingFederate Go Client. Check authentication type and credentials$`
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

	expectedErrorPattern := `^failed to initialize PingFederate Go Client. Check authentication type and credentials$`
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

	expectedErrorPattern := `^failed to initialize PingFederate Go Client. Check authentication type and credentials$`
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
