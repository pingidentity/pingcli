package platform_internal

import (
	"context"
	"crypto/tls"
	"net/http"
	"os"
	"regexp"
	"testing"

	"github.com/pingidentity/pingcli/internal/configuration/options"
	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/customtypes"
	"github.com/pingidentity/pingcli/internal/profiles"
	"github.com/pingidentity/pingcli/internal/testing/testutils"
	"github.com/pingidentity/pingcli/internal/testing/testutils_viper"
	pingfederateGoClient "github.com/pingidentity/pingfederate-go-client/v1220/configurationapi"
)

// Test RunInternalExport function
func TestRunInternalExport(t *testing.T) {
	testutils_viper.InitVipers(t)

	err := RunInternalExport(context.Background(), "v1.2.3")
	testutils.CheckExpectedError(t, err, nil)

	// Check if there are terraform files in the export directory
	outputDir, err := profiles.GetOptionValue(options.PlatformExportOutputDirectoryOption)
	if err != nil {
		t.Fatalf("profiles.GetOptionValue() error = %v", err)
	}

	files, err := os.ReadDir(outputDir)
	if err != nil {
		t.Fatalf("os.ReadDir() error = %v", err)
	}

	// Check the number of files in the directory
	if len(files) == 0 {
		t.Errorf("RunInternalExport() num files = %v, want non-zero", len(files))
	}

	// Check the file type is .tf
	re := regexp.MustCompile(`^.*\.tf$`)
	for _, file := range files {
		if file.IsDir() {
			t.Errorf("RunInternalExport() file = %v, want file", file)
		}

		if !re.MatchString(file.Name()) {
			t.Errorf("RunInternalExport() file = %v, want .tf file", file.Name())
		}
	}
}

// Test RunInternalExport function fails with nil context
func TestRunInternalExportNilContext(t *testing.T) {
	expectedErrorPattern := `^failed to run 'platform export' command\. context is nil$`
	err := RunInternalExport(nil, "v1.2.3") //nolint:staticcheck // SA1012 this is a test
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}

// Test initPingFederateServices function
func TestInitPingFederateServices(t *testing.T) {
	testutils_viper.InitVipers(t)

	err := initPingFederateServices(context.Background(), "v1.2.3")
	testutils.CheckExpectedError(t, err, nil)

	// make sure pf context is not nil
	if pingfederateContext == nil {
		t.Errorf("initPingFederateServices() pingfederateContext = %v, want non-nil", pingfederateContext)
	}

	// check pf context has auth values included
	if pingfederateContext.Value(pingfederateGoClient.ContextBasicAuth) == nil {
		t.Errorf("initPingFederateServices() pingfederateContext.Value = %v, want non-nil", pingfederateContext.Value(pingfederateGoClient.ContextBasicAuth))
	}
}

// Test initPingFederateServices function fails with nil context
func TestInitPingFederateServicesNilContext(t *testing.T) {
	expectedErrorPattern := `^failed to initialize PingFederate services\. context is nil$`
	err := initPingFederateServices(nil, "v1.2.3") //nolint:staticcheck // SA1012 this is a test
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}

// Test initPingOneServices function
func TestInitPingOneServices(t *testing.T) {
	testutils_viper.InitVipers(t)

	err := initPingOneServices(context.Background(), "v1.2.3")
	testutils.CheckExpectedError(t, err, nil)

	// make sure po context is not nil
	if pingoneContext == nil {
		t.Errorf("initPingOneServices() pingoneContext = %v, want non-nil", pingoneContext)
	}
}

// Test initPingFederateApiClient function
func TestInitPingFederateApiClient(t *testing.T) {
	testutils_viper.InitVipers(t)

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true, //#nosec G402 -- This is a test
		},
	}

	err := initPingFederateApiClient(tr, "v1.2.3")
	testutils.CheckExpectedError(t, err, nil)

	// make sure pf client is not nil
	if pingfederateApiClient == nil {
		t.Errorf("initPingFederateApiClient() pingfederateApiClient = %v, want non-nil", pingfederateApiClient)
	}
}

// Test initPingFederateApiClient function fails with nil transport
func TestInitPingFederateApiClientNilTransport(t *testing.T) {
	expectedErrorPattern := `^failed to initialize pingfederate API client\. http transport is nil$`
	err := initPingFederateApiClient(nil, "v1.2.3")
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}

// Test initPingOneApiClient function
func TestInitPingOneApiClient(t *testing.T) {
	testutils_viper.InitVipers(t)

	err := initPingOneApiClient(context.Background(), "v1.2.3")
	testutils.CheckExpectedError(t, err, nil)

	// make sure po client is not nil
	if pingoneApiClient == nil {
		t.Errorf("initPingOneApiClient() pingoneApiClient = %v, want non-nil", pingoneApiClient)
	}
}

// Test initPingOneApiClient function fails with nil context
func TestInitPingOneApiClientNilContext(t *testing.T) {
	expectedErrorPattern := `^failed to initialize pingone API client\. context is nil$`
	err := initPingOneApiClient(nil, "v1.2.3") //nolint:staticcheck // SA1012 this is a test
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}

// Test createOrValidateOutputDir function with non-existent directory
func TestCreateOrValidateOutputDir(t *testing.T) {
	testutils_viper.InitVipers(t)

	outputDir := os.TempDir() + "/nonexistantdir"

	_, err := createOrValidateOutputDir(outputDir, false)
	testutils.CheckExpectedError(t, err, nil)
}

// Test createOrValidateOutputDir function with existent directory
func TestCreateOrValidateOutputDirExistentDir(t *testing.T) {
	testutils_viper.InitVipers(t)

	outputDir := t.TempDir()

	_, err := createOrValidateOutputDir(outputDir, false)
	testutils.CheckExpectedError(t, err, nil)
}

// Test createOrValidateOutputDir function with existent directory and overwrite flag
// when there is a file in the directory
func TestCreateOrValidateOutputDirExistentDirWithFile(t *testing.T) {
	testutils_viper.InitVipers(t)

	outputDir := t.TempDir()

	file, err := os.Create(outputDir + "/file")
	if err != nil {
		t.Fatalf("os.Create() error = %v", err)
	}
	file.Close()

	_, err = createOrValidateOutputDir(outputDir, true)
	testutils.CheckExpectedError(t, err, nil)
}

// Test createOrValidateOutputDir function fails with existent directory and no overwrite flag
// when there is a file in the directory
func TestCreateOrValidateOutputDirExistentDirWithFileNoOverwrite(t *testing.T) {
	testutils_viper.InitVipers(t)

	outputDir := t.TempDir()

	file, err := os.Create(outputDir + "/file")
	if err != nil {
		t.Fatalf("os.Create() error = %v", err)
	}
	file.Close()

	expectedErrorPattern := `^output directory '.*' is not empty\. Use --overwrite to overwrite existing export data$`
	_, err = createOrValidateOutputDir(outputDir, false)
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}

// Test getPingOneExportEnvID function
func TestGetPingOneExportEnvID(t *testing.T) {
	testutils_viper.InitVipers(t)

	if err := getPingOneExportEnvID(); err != nil {
		t.Errorf("getPingOneExportEnvID() error = %v, want nil", err)
	}

	// Check pingoneExportEnvID is not empty
	if pingoneExportEnvID == "" {
		t.Errorf("getPingOneExportEnvID() pingoneExportEnvID = %v, want non-empty", pingoneExportEnvID)
	}
}

// Test validatePingOneExportEnvID function
func TestValidatePingOneExportEnvID(t *testing.T) {
	testutils_viper.InitVipers(t)

	if err := initPingOneApiClient(context.Background(), "v1.2.3"); err != nil {
		t.Errorf("initPingOneApiClient() error = %v, want nil", err)
	}

	if err := getPingOneExportEnvID(); err != nil {
		t.Errorf("getPingOneExportEnvID() error = %v, want nil", err)
	}

	err := validatePingOneExportEnvID(context.Background())
	testutils.CheckExpectedError(t, err, nil)
}

// Test validatePingOneExportEnvID function fails with nil context
func TestValidatePingOneExportEnvIDNilContext(t *testing.T) {
	expectedErrorPattern := `^failed to validate pingone environment ID '.*'\. context is nil$`
	err := validatePingOneExportEnvID(nil) //nolint:staticcheck // SA1012 this is a test
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}

// Test getExportableConnectors function
func TestGetExportableConnectors(t *testing.T) {
	testutils_viper.InitVipers(t)

	es := new(customtypes.ExportServices)
	err := es.Set(customtypes.ENUM_EXPORT_SERVICE_PINGONE_PROTECT)
	if err != nil {
		t.Fatalf("ms.Set() error = %v", err)
	}

	expectedConnectors := len(es.GetServices())

	exportableConnectors := getExportableConnectors(es)
	if len(*exportableConnectors) == 0 {
		t.Errorf("getExportableConnectors() exportableConnectors = %v, want non-empty", exportableConnectors)
	}

	if len(*exportableConnectors) != expectedConnectors {
		t.Errorf("getExportableConnectors() exportableConnectors = %v, want %v", len(*exportableConnectors), expectedConnectors)
	}
}

// Test getExportableConnectors function with nil MultiService
func TestGetExportableConnectorsNilMultiService(t *testing.T) {
	exportableConnectors := getExportableConnectors(nil)

	expectedConnectors := 0
	if len(*exportableConnectors) != expectedConnectors {
		t.Errorf("getExportableConnectors() exportableConnectors = %v, want %v", len(*exportableConnectors), expectedConnectors)
	}
}

// Test exportConnectors function
func TestExportConnectors(t *testing.T) {
	testutils_viper.InitVipers(t)

	err := initPingOneServices(context.Background(), "v1.2.3")
	if err != nil {
		t.Fatalf("initPingOneServices() error = %v", err)
	}

	es := new(customtypes.ExportServices)
	err = es.Set(customtypes.ENUM_EXPORT_SERVICE_PINGONE_PROTECT)
	if err != nil {
		t.Fatalf("ms.Set() error = %v", err)
	}

	exportableConnectors := getExportableConnectors(es)

	err = exportConnectors(exportableConnectors, customtypes.ENUM_EXPORT_FORMAT_HCL, t.TempDir(), true)
	testutils.CheckExpectedError(t, err, nil)
}

// Test exportConnectors function with nil exportable connectors
func TestExportConnectorsNilExportableConnectors(t *testing.T) {
	err := exportConnectors(nil, customtypes.ENUM_EXPORT_FORMAT_HCL, t.TempDir(), true)

	expectedErrorPattern := `^failed to export services\. exportable connectors list is nil$`
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}

// Test exportConnectors function with empty exportable connectors
func TestExportConnectorsEmptyExportableConnectors(t *testing.T) {
	exportableConnectors := &[]connector.Exportable{}

	err := exportConnectors(exportableConnectors, customtypes.ENUM_EXPORT_FORMAT_HCL, t.TempDir(), true)
	testutils.CheckExpectedError(t, err, nil)
}

// Test exportConnectors function with invalid export format
func TestExportConnectorsInvalidExportFormat(t *testing.T) {
	testutils_viper.InitVipers(t)

	err := initPingOneServices(context.Background(), "v1.2.3")
	if err != nil {
		t.Fatalf("initPingOneServices() error = %v", err)
	}

	es := new(customtypes.ExportServices)
	err = es.Set(customtypes.ENUM_EXPORT_SERVICE_PINGONE_PROTECT)
	if err != nil {
		t.Fatalf("ms.Set() error = %v", err)
	}

	exportableConnectors := getExportableConnectors(es)

	err = exportConnectors(exportableConnectors, "invalid", t.TempDir(), true)

	expectedErrorPattern := `^failed to export '.*' service: unrecognized export format ".*"\. Must be one of: .*$`
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}

// Test exportConnectors function with invalid output directory
func TestExportConnectorsInvalidOutputDir(t *testing.T) {
	testutils_viper.InitVipers(t)

	err := initPingOneServices(context.Background(), "v1.2.3")
	if err != nil {
		t.Fatalf("initPingOneServices() error = %v", err)
	}

	es := new(customtypes.ExportServices)
	err = es.Set(customtypes.ENUM_EXPORT_SERVICE_PINGONE_PROTECT)
	if err != nil {
		t.Fatalf("ms.Set() error = %v", err)
	}

	exportableConnectors := getExportableConnectors(es)

	err = exportConnectors(exportableConnectors, customtypes.ENUM_EXPORT_FORMAT_HCL, "/invalid", true)

	expectedErrorPattern := `^failed to export '.*' service: failed to create export file ".*". err: open .*: no such file or directory$`
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}
