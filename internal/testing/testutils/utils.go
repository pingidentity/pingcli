package testutils

import (
	"context"
	"crypto/tls"
	"net/http"
	"os"
	"regexp"
	"sync"
	"testing"

	"github.com/patrickcping/pingone-go-sdk-v2/management"
	"github.com/patrickcping/pingone-go-sdk-v2/pingone"
	"github.com/pingidentity/pingcli/internal/configuration"
	"github.com/pingidentity/pingcli/internal/configuration/options"
	"github.com/pingidentity/pingcli/internal/connector"
	pingfederateGoClient "github.com/pingidentity/pingfederate-go-client/v1210/configurationapi"
)

var (
	envIdOnce         sync.Once
	apiClientOnce     sync.Once
	PingOneClientInfo *connector.PingOneClientInfo
	environmentId     string
)

func GetEnvironmentID() string {
	envIdOnce.Do(func() {
		environmentId = os.Getenv(options.PlatformExportPingOneEnvironmentIDOption.EnvVar)
	})

	return environmentId
}

// Utility method to initialize a PingOne SDK client for testing
func GetPingOneClientInfo(t *testing.T) *connector.PingOneClientInfo {
	t.Helper()

	apiClientOnce.Do(func() {
		configuration.InitAllOptions()
		// Grab environment vars for initializing the API client.
		// These are set in GitHub Actions.
		clientID := os.Getenv(options.PingOneAuthenticationWorkerClientIDOption.EnvVar)
		clientSecret := os.Getenv(options.PingOneAuthenticationWorkerClientSecretOption.EnvVar)
		environmentId := GetEnvironmentID()
		regionCode := os.Getenv(options.PingOneRegionCodeOption.EnvVar)
		sdkRegionCode := management.EnumRegionCode(regionCode)

		if clientID == "" || clientSecret == "" || environmentId == "" || regionCode == "" {
			t.Fatalf("Unable to retrieve env var value for one or more of clientID, clientSecret, environmentID, regionCode.")
		}

		apiConfig := &pingone.Config{
			ClientID:      &clientID,
			ClientSecret:  &clientSecret,
			EnvironmentID: &environmentId,
			RegionCode:    &sdkRegionCode,
		}

		// Make empty context for testing
		ctx := context.Background()

		// Initialize the API client
		client, err := apiConfig.APIClient(ctx)
		if err != nil {
			t.Fatal(err.Error())
		}

		PingOneClientInfo = &connector.PingOneClientInfo{
			Context:             ctx,
			ApiClient:           client,
			ApiClientId:         &clientID,
			ExportEnvironmentID: environmentId,
		}
	})

	return PingOneClientInfo
}

func GetPingFederateClientInfo(t *testing.T) *connector.PingFederateClientInfo {
	t.Helper()

	configuration.InitAllOptions()

	httpsHost := os.Getenv(options.PingFederateHTTPSHostOption.EnvVar)
	adminApiPath := os.Getenv(options.PingFederateAdminAPIPathOption.EnvVar)
	pfUsername := os.Getenv(options.PingFederateBasicAuthUsernameOption.EnvVar)
	pfPassword := os.Getenv(options.PingFederateBasicAuthPasswordOption.EnvVar)

	if httpsHost == "" || adminApiPath == "" || pfUsername == "" || pfPassword == "" {
		t.Fatalf("Unable to retrieve env var value for one or more of httpsHost, adminApiPath, pfUsername, pfPassword.")
	}

	pfClientConfig := pingfederateGoClient.NewConfiguration()
	pfClientConfig.DefaultHeader["X-Xsrf-Header"] = "PingFederate"
	pfClientConfig.Servers = pingfederateGoClient.ServerConfigurations{
		{
			URL: httpsHost + adminApiPath,
		},
	}
	httpClient := &http.Client{Transport: &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true, //#nosec G402 -- This is a test
		}}}
	pfClientConfig.HTTPClient = httpClient

	apiClient := pingfederateGoClient.NewAPIClient(pfClientConfig)

	return &connector.PingFederateClientInfo{
		ApiClient: apiClient,
		Context: context.WithValue(context.Background(), pingfederateGoClient.ContextBasicAuth, pingfederateGoClient.BasicAuth{
			UserName: pfUsername,
			Password: pfPassword,
		}),
	}
}

func ValidateImportBlocks(t *testing.T, resource connector.ExportableResource, expectedImportBlocks *[]connector.ImportBlock) {
	t.Helper()

	importBlocks, err := resource.ExportAll()
	if err != nil {
		t.Fatalf("Failed to export %s: %s", resource.ResourceType(), err.Error())
	}

	// Make sure the resource name and id in each import block is unique across all import blocks
	resourceNames := map[string]bool{}
	resourceIDs := map[string]bool{}
	for _, importBlock := range *importBlocks {
		if resourceNames[importBlock.ResourceName] {
			t.Errorf("Resource name %s is not unique", importBlock.ResourceName)
		}
		resourceNames[importBlock.ResourceName] = true

		if resourceIDs[importBlock.ResourceID] {
			t.Errorf("Resource ID %s is not unique", importBlock.ResourceID)
		}
		resourceIDs[importBlock.ResourceID] = true
	}

	// Check if provided pointer to expected import blocks is nil, and created an empty slice if so.
	if expectedImportBlocks == nil {
		expectedImportBlocks = &[]connector.ImportBlock{}
	}

	expectedImportBlocksMap := map[string]connector.ImportBlock{}
	for _, importBlock := range *expectedImportBlocks {
		expectedImportBlocksMap[importBlock.ResourceName] = importBlock
	}

	// Check number of export blocks
	expectedNumberOfBlocks := len(expectedImportBlocksMap)
	actualNumberOfBlocks := len(*importBlocks)
	if actualNumberOfBlocks != expectedNumberOfBlocks {
		t.Fatalf("Expected %d import blocks, got %d", expectedNumberOfBlocks, actualNumberOfBlocks)
	}

	// Make sure the import blocks match the expected import blocks
	for _, importBlock := range *importBlocks {
		expectedImportBlock, ok := expectedImportBlocksMap[importBlock.ResourceName]

		if !ok {
			t.Errorf("No matching expected import block for generated import block:\n%s", importBlock.String())
			continue
		}

		if !importBlock.Equals(expectedImportBlock) {
			t.Errorf("Expected import block \n%s\n Got import block \n%s", expectedImportBlock.String(), importBlock.String())
		}
	}
}

func CheckExpectedError(t *testing.T, err error, errMessagePattern *string) {
	t.Helper()

	if err == nil && errMessagePattern != nil {
		t.Errorf("Error message did not match expected regex\n\nerror message: '%v'\n\nregex pattern %s", err, *errMessagePattern)
		return
	}

	if err != nil && errMessagePattern == nil {
		t.Errorf("Expected no error, but got error: %v", err)
		return
	}

	if err != nil {
		regex := regexp.MustCompile(*errMessagePattern)
		if !regex.MatchString(err.Error()) {
			t.Errorf("Error message did not match expected regex\n\nerror message: '%v'\n\nregex pattern %s", err, *errMessagePattern)
		}
	}
}

// Get os.File with string written to it.
// The caller is responsible for closing the file.
func WriteStringToPipe(str string, t *testing.T) (reader *os.File) {
	t.Helper()

	reader, writer, err := os.Pipe()
	if err != nil {
		t.Fatal(err)
	}

	defer writer.Close()

	if _, err := writer.WriteString(str); err != nil {
		reader.Close()
		t.Fatal(err)
	}

	// Close the writer to simulate EOF
	if err = writer.Close(); err != nil {
		reader.Close()
		t.Fatal(err)
	}

	return reader
}
