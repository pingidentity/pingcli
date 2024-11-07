package platform_test

import (
	"os"
	"testing"

	"github.com/pingidentity/pingcli/internal/configuration/options"
	"github.com/pingidentity/pingcli/internal/customtypes"
	"github.com/pingidentity/pingcli/internal/testing/testutils"
	"github.com/pingidentity/pingcli/internal/testing/testutils_cobra"
	"github.com/pingidentity/pingcli/internal/testing/testutils_viper"
)

// Test Platform Export Command Executes without issue
func TestPlatformExportCmd_Execute(t *testing.T) {
	outputDir := t.TempDir()

	err := testutils_cobra.ExecutePingcli(t, "platform", "export",
		"--output-directory", outputDir,
		"--overwrite")
	testutils.CheckExpectedError(t, err, nil)
}

// Test Platform Export Command fails when provided too many arguments
func TestPlatformExportCmd_TooManyArgs(t *testing.T) {
	testutils_viper.InitVipers(t)

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

// Test Platform Export Command --services flag
func TestPlatformExportCmd_ServiceFlag(t *testing.T) {
	outputDir := t.TempDir()

	err := testutils_cobra.ExecutePingcli(t, "platform", "export",
		"--output-directory", outputDir,
		"--overwrite",
		"--services", "pingone-protect")
	testutils.CheckExpectedError(t, err, nil)
}

// Test Platform Export Command --services flag with invalid service
func TestPlatformExportCmd_ServiceFlagInvalidService(t *testing.T) {
	expectedErrorPattern := `^invalid argument ".*" for "-s, --services" flag: failed to set ExportServices: Invalid service: .*\. Allowed services: .*$`
	err := testutils_cobra.ExecutePingcli(t, "platform", "export", "--services", "invalid")
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}

// Test Platform Export Command --format flag
func TestPlatformExportCmd_ExportFormatFlag(t *testing.T) {
	outputDir := t.TempDir()

	err := testutils_cobra.ExecutePingcli(t, "platform", "export",
		"--output-directory", outputDir,
		"--format", "HCL",
		"--overwrite",
		"--services", "pingone-protect")
	testutils.CheckExpectedError(t, err, nil)
}

// Test Platform Export Command --format flag with invalid format
func TestPlatformExportCmd_ExportFormatFlagInvalidFormat(t *testing.T) {
	expectedErrorPattern := `^invalid argument ".*" for "-f, --format" flag: unrecognized export format '.*'\. Must be one of: .*$`
	err := testutils_cobra.ExecutePingcli(t, "platform", "export", "--format", "invalid")
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}

// Test Platform Export Command --output-directory flag
func TestPlatformExportCmd_OutputDirectoryFlag(t *testing.T) {
	outputDir := t.TempDir()

	err := testutils_cobra.ExecutePingcli(t, "platform", "export",
		"--output-directory", outputDir,
		"--overwrite",
		"--services", "pingone-protect")
	testutils.CheckExpectedError(t, err, nil)
}

// Test Platform Export Command --output-directory flag with invalid directory
func TestPlatformExportCmd_OutputDirectoryFlagInvalidDirectory(t *testing.T) {
	expectedErrorPattern := `^failed to create output directory '\/invalid': mkdir \/invalid: .+$`
	err := testutils_cobra.ExecutePingcli(t, "platform", "export", "--output-directory", "/invalid")
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}

// Test Platform Export Command --overwrite flag
func TestPlatformExportCmd_OverwriteFlag(t *testing.T) {
	outputDir := t.TempDir()

	err := testutils_cobra.ExecutePingcli(t, "platform", "export",
		"--output-directory", outputDir,
		"--overwrite",
		"--services", "pingone-protect")
	testutils.CheckExpectedError(t, err, nil)
}

// Test Platform Export Command --overwrite flag false with existing directory
// where the directory already contains a file
func TestPlatformExportCmd_OverwriteFlagFalseWithExistingDirectory(t *testing.T) {
	outputDir := t.TempDir()

	_, err := os.Create(outputDir + "/file")
	if err != nil {
		t.Errorf("Error creating file in output directory: %v", err)
	}

	expectedErrorPattern := `^output directory '[A-Za-z0-9_\-\/]+' is not empty\. Use --overwrite to overwrite existing export data$`
	err = testutils_cobra.ExecutePingcli(t, "platform", "export",
		"--output-directory", outputDir,
		"--services", "pingone-protect",
		"--overwrite=false")
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}

// Test Platform Export Command --overwrite flag true with existing directory
// where the directory already contains a file
func TestPlatformExportCmd_OverwriteFlagTrueWithExistingDirectory(t *testing.T) {
	outputDir := t.TempDir()

	_, err := os.Create(outputDir + "/file")
	if err != nil {
		t.Errorf("Error creating file in output directory: %v", err)
	}

	err = testutils_cobra.ExecutePingcli(t, "platform", "export",
		"--output-directory", outputDir,
		"--services", "pingone-protect",
		"--overwrite")
	testutils.CheckExpectedError(t, err, nil)
}

// Test Platform Export Command with
// --pingone-worker-environment-id flag
// --pingone-worker-client-id flag
// --pingone-worker-client-secret flag
// --pingone-region flag
func TestPlatformExportCmd_PingOneWorkerEnvironmentIdFlag(t *testing.T) {
	outputDir := t.TempDir()

	err := testutils_cobra.ExecutePingcli(t, "platform", "export",
		"--output-directory", outputDir,
		"--overwrite",
		"--services", "pingone-protect",
		"--pingone-worker-environment-id", os.Getenv(options.PingOneAuthenticationWorkerEnvironmentIDOption.EnvVar),
		"--pingone-worker-client-id", os.Getenv(options.PingOneAuthenticationWorkerClientIDOption.EnvVar),
		"--pingone-worker-client-secret", os.Getenv(options.PingOneAuthenticationWorkerClientSecretOption.EnvVar),
		"--pingone-region-code", os.Getenv(options.PingOneRegionCodeOption.EnvVar))
	testutils.CheckExpectedError(t, err, nil)
}

// Test Platform Export Command fails when not provided required pingone flags together
func TestPlatformExportCmd_PingOneWorkerEnvironmentIdFlagRequiredTogether(t *testing.T) {
	expectedErrorPattern := `^if any flags in the group \[pingone-worker-environment-id pingone-worker-client-id pingone-worker-client-secret pingone-region-code] are set they must all be set; missing \[pingone-region-code pingone-worker-client-id pingone-worker-client-secret]$`
	err := testutils_cobra.ExecutePingcli(t, "platform", "export",
		"--pingone-worker-environment-id", os.Getenv(options.PingOneAuthenticationWorkerEnvironmentIDOption.EnvVar))
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}

// Test Platform Export command with PingFederate Basic Auth flags
func TestPlatformExportCmd_PingFederateBasicAuthFlags(t *testing.T) {
	outputDir := t.TempDir()

	err := testutils_cobra.ExecutePingcli(t, "platform", "export",
		"--output-directory", outputDir,
		"--overwrite",
		"--services", "pingfederate",
		"--pingfederate-username", os.Getenv(options.PingFederateBasicAuthUsernameOption.EnvVar),
		"--pingfederate-password", os.Getenv(options.PingFederateBasicAuthPasswordOption.EnvVar),
		"--pingfederate-authentication-type", "basicAuth",
	)
	testutils.CheckExpectedError(t, err, nil)
}

// Test Platform Export Command fails when not provided required PingFederate Basic Auth flags together
func TestPlatformExportCmd_PingFederateBasicAuthFlagsRequiredTogether(t *testing.T) {
	expectedErrorPattern := `^if any flags in the group \[pingfederate-username pingfederate-password] are set they must all be set; missing \[pingfederate-password]$`
	err := testutils_cobra.ExecutePingcli(t, "platform", "export",
		"--pingfederate-username", "Administrator")
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}

// Test Platform Export Command fails when provided invalid PingFederate Basic Auth flags
func TestPlatformExportCmd_PingFederateBasicAuthFlagsInvalid(t *testing.T) {
	outputDir := t.TempDir()

	expectedErrorPattern := `^failed to export 'pingfederate' service: failed to export resource .*\. err: .* Request for resource '.*' was not successful\.\s+Response Code: 401 Unauthorized\s+Response Body: {{"resultId":"invalid_credentials","message":"The credentials you provided were not recognized\."}}\s+Error: 401 Unauthorized$`
	err := testutils_cobra.ExecutePingcli(t, "platform", "export",
		"--output-directory", outputDir,
		"--overwrite",
		"--services", "pingfederate",
		"--pingfederate-username", "Administrator",
		"--pingfederate-password", "invalid",
		"--pingfederate-authentication-type", "basicAuth",
	)
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}

// Test Platform Export command with PingFederate Client Credentials Auth flags
func TestPlatformExportCmd_PingFederateClientCredentialsAuthFlags(t *testing.T) {
	outputDir := t.TempDir()

	err := testutils_cobra.ExecutePingcli(t, "platform", "export",
		"--output-directory", outputDir,
		"--overwrite",
		"--services", "pingfederate",
		"--pingfederate-client-id", os.Getenv(options.PingFederateClientCredentialsAuthClientIDOption.EnvVar),
		"--pingfederate-client-secret", os.Getenv(options.PingFederateClientCredentialsAuthClientSecretOption.EnvVar),
		"--pingfederate-scopes", os.Getenv(options.PingFederateClientCredentialsAuthScopesOption.EnvVar),
		"--pingfederate-token-url", os.Getenv(options.PingFederateClientCredentialsAuthTokenURLOption.EnvVar),
		"--pingfederate-authentication-type", "clientCredentialsAuth",
	)
	testutils.CheckExpectedError(t, err, nil)
}

// Test Platform Export Command fails when not provided required PingFederate Client Credentials Auth flags together
func TestPlatformExportCmd_PingFederateClientCredentialsAuthFlagsRequiredTogether(t *testing.T) {
	expectedErrorPattern := `^if any flags in the group \[pingfederate-client-id pingfederate-client-secret pingfederate-token-url] are set they must all be set; missing \[pingfederate-client-secret pingfederate-token-url]$`
	err := testutils_cobra.ExecutePingcli(t, "platform", "export",
		"--pingfederate-client-id", "test")
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}

// Test Platform Export Command fails when provided invalid PingFederate Client Credentials Auth flags
func TestPlatformExportCmd_PingFederateClientCredentialsAuthFlagsInvalid(t *testing.T) {
	outputDir := t.TempDir()

	expectedErrorPattern := `^failed to export 'pingfederate' service: failed to export resource .*\. err: .* Request for resource '.*' was not successful\. Response is nil\. Error: oauth2: "invalid_client" "Invalid client or client credentials\."$`
	err := testutils_cobra.ExecutePingcli(t, "platform", "export",
		"--output-directory", outputDir,
		"--overwrite",
		"--services", "pingfederate",
		"--pingfederate-client-id", "test",
		"--pingfederate-client-secret", "invalid",
		"--pingfederate-token-url", "https://localhost:9031/as/token.oauth2",
		"--pingfederate-authentication-type", "clientCredentialsAuth",
	)
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}

// Test Platform Export Command fails when provided invalid PingFederate OAuth2 Token URL
func TestPlatformExportCmd_PingFederateClientCredentialsAuthFlagsInvalidTokenURL(t *testing.T) {
	outputDir := t.TempDir()

	expectedErrorPattern := `(?s)^failed to export 'pingfederate' service: failed to export resource.*\. err:.*Request for resource '.*' was not successful\. Response is nil\. Error: oauth2: cannot fetch token: 404 Not Found\s+Response: \<!DOCTYPE html\>\s+.*$`
	err := testutils_cobra.ExecutePingcli(t, "platform", "export",
		"--output-directory", outputDir,
		"--overwrite",
		"--services", "pingfederate",
		"--pingfederate-client-id", os.Getenv(options.PingFederateClientCredentialsAuthClientIDOption.EnvVar),
		"--pingfederate-client-secret", os.Getenv(options.PingFederateClientCredentialsAuthClientSecretOption.EnvVar),
		"--pingfederate-token-url", "https://localhost:9031/as/invalid",
		"--pingfederate-authentication-type", "clientCredentialsAuth",
	)
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}

// Test Platform Export command with PingFederate X-Bypass Header set to true
func TestPlatformExportCmd_PingFederateXBypassHeaderFlag(t *testing.T) {
	outputDir := t.TempDir()

	err := testutils_cobra.ExecutePingcli(t, "platform", "export",
		"--output-directory", outputDir,
		"--overwrite",
		"--services", "pingfederate",
		"--pingfederate-x-bypass-external-validation-header",
		"--pingfederate-username", os.Getenv(options.PingFederateBasicAuthUsernameOption.EnvVar),
		"--pingfederate-password", os.Getenv(options.PingFederateBasicAuthPasswordOption.EnvVar),
		"--pingfederate-authentication-type", "basicAuth",
	)
	testutils.CheckExpectedError(t, err, nil)
}

// Test Platform Export command with PingFederate --pingfederate-insecure-trust-all-tls flag set to true
func TestPlatformExportCmd_PingFederateTrustAllTLSFlag(t *testing.T) {
	outputDir := t.TempDir()

	err := testutils_cobra.ExecutePingcli(t, "platform", "export",
		"--output-directory", outputDir,
		"--overwrite",
		"--services", "pingfederate",
		"--pingfederate-insecure-trust-all-tls",
		"--pingfederate-username", os.Getenv(options.PingFederateBasicAuthUsernameOption.EnvVar),
		"--pingfederate-password", os.Getenv(options.PingFederateBasicAuthPasswordOption.EnvVar),
		"--pingfederate-authentication-type", "basicAuth",
	)
	testutils.CheckExpectedError(t, err, nil)
}

// Test Platform Export command fails with PingFederate --pingfederate-insecure-trust-all-tls flag set to false
func TestPlatformExportCmd_PingFederateTrustAllTLSFlagFalse(t *testing.T) {
	outputDir := t.TempDir()

	expectedErrorPattern := `^failed to export '.*' service: failed to export resource .*\. err: .* Request for resource '.*' was not successful\. Response is nil\. Error: Get "https.*": tls: failed to verify certificate: x509: certificate signed by unknown authority$`
	err := testutils_cobra.ExecutePingcli(t, "platform", "export",
		"--output-directory", outputDir,
		"--overwrite",
		"--services", "pingfederate",
		"--pingfederate-insecure-trust-all-tls=false",
		"--pingfederate-username", os.Getenv(options.PingFederateBasicAuthUsernameOption.EnvVar),
		"--pingfederate-password", os.Getenv(options.PingFederateBasicAuthPasswordOption.EnvVar),
		"--pingfederate-authentication-type", "basicAuth",
	)
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}

// Test Platform Export command passes with PingFederate
// --pingfederate-insecure-trust-all-tls=false
// and --pingfederate-ca-certificate-pem-files set
func TestPlatformExportCmd_PingFederateCaCertificatePemFiles(t *testing.T) {
	outputDir := t.TempDir()

	err := testutils_cobra.ExecutePingcli(t, "platform", "export",
		"--output-directory", outputDir,
		"--overwrite",
		"--services", "pingfederate",
		"--pingfederate-insecure-trust-all-tls=false",
		"--pingfederate-ca-certificate-pem-files", "testdata/ssl-server-crt.pem",
		"--pingfederate-username", os.Getenv(options.PingFederateBasicAuthUsernameOption.EnvVar),
		"--pingfederate-password", os.Getenv(options.PingFederateBasicAuthPasswordOption.EnvVar),
		"--pingfederate-authentication-type", customtypes.ENUM_PINGFEDERATE_AUTHENTICATION_TYPE_BASIC,
	)
	testutils.CheckExpectedError(t, err, nil)
}

// Test Platform Export command fails with --pingfederate-ca-certificate-pem-files set to non-existent file.
func TestPlatformExportCmd_PingFederateCaCertificatePemFilesInvalid(t *testing.T) {
	expectedErrorPattern := `^failed to read CA certificate PEM file '.*': open .*: no such file or directory$`
	err := testutils_cobra.ExecutePingcli(t, "platform", "export",
		"--services", "pingfederate",
		"--pingfederate-ca-certificate-pem-files", "invalid/crt.pem",
		"--pingfederate-username", os.Getenv(options.PingFederateBasicAuthUsernameOption.EnvVar),
		"--pingfederate-password", os.Getenv(options.PingFederateBasicAuthPasswordOption.EnvVar),
		"--pingfederate-authentication-type", "basicAuth",
	)
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}
