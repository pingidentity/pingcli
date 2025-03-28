// Copyright © 2025 Ping Identity Corporation

package testutils

import (
	"bytes"
	"context"
	"crypto/rand"
	"crypto/rsa"
	"crypto/tls"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"errors"
	"fmt"
	"math/big"
	"net/http"
	"os"
	"regexp"
	"sync"
	"testing"
	"time"

	"github.com/patrickcping/pingone-go-sdk-v2/management"
	"github.com/patrickcping/pingone-go-sdk-v2/pingone"
	"github.com/pingidentity/pingcli/internal/configuration"
	"github.com/pingidentity/pingcli/internal/connector"
	pingfederateGoClient "github.com/pingidentity/pingfederate-go-client/v1220/configurationapi"
)

var (
	clientInfoOnce sync.Once
	clientInfo     *connector.ClientInfo = &connector.ClientInfo{}
)

func GetClientInfo(t *testing.T) *connector.ClientInfo {
	t.Helper()

	// Ensure that the client info is initialized only once
	clientInfoOnce.Do(func() {
		configuration.InitAllOptions()

		initPingFederateClientInfo(t, clientInfo)
		initPingOneClientInfo(t, clientInfo)
	})

	return clientInfo
}

func initPingFederateClientInfo(t *testing.T, clientInfo *connector.ClientInfo) {
	t.Helper()

	httpsHost := "https://localhost:9999"
	adminApiPath := "/pf-admin-api/v1"
	pfUsername := "Administrator"
	pfPassword := "2FederateM0re"

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

	clientInfo.PingFederateApiClient = pingfederateGoClient.NewAPIClient(pfClientConfig)
	clientInfo.PingFederateContext = context.WithValue(context.WithoutCancel(t.Context()), pingfederateGoClient.ContextBasicAuth, pingfederateGoClient.BasicAuth{
		UserName: pfUsername,
		Password: pfPassword,
	})
}

func initPingOneClientInfo(t *testing.T, clientInfo *connector.ClientInfo) {
	t.Helper()

	// Grab environment vars for initializing the API client.
	// These are set in GitHub Actions.
	clientID := os.Getenv("TEST_PINGONE_WORKER_CLIENT_ID")
	clientSecret := os.Getenv("TEST_PINGONE_WORKER_CLIENT_SECRET")
	environmentId := os.Getenv("TEST_PINGONE_ENVIRONMENT_ID")
	regionCode := os.Getenv("TEST_PINGONE_REGION_CODE")
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
	ctx := context.WithoutCancel(t.Context())

	// Initialize the API client
	client, err := apiConfig.APIClient(ctx)
	if err != nil {
		t.Fatal(err.Error())
	}

	clientInfo.PingOneApiClient = client
	clientInfo.PingOneContext = ctx
	clientInfo.PingOneApiClientId = clientID
	clientInfo.PingOneExportEnvironmentID = environmentId
}

func getValidatedActualImportBlocks(t *testing.T, resource connector.ExportableResource) *[]connector.ImportBlock {
	t.Helper()

	importBlocks, err := resource.ExportAll()
	if err != nil {
		t.Errorf("Failed to export %s: %s", resource.ResourceType(), err.Error())

		return nil
	}

	// Make sure the resource name and id in each import block is unique across all import blocks
	resourceNames := map[string]bool{}
	resourceIDs := map[string]bool{}
	for _, importBlock := range *importBlocks {
		if resourceNames[importBlock.ResourceName] {
			t.Errorf("Resource name %s is not unique", importBlock.ResourceName)

			return nil
		}
		resourceNames[importBlock.ResourceName] = true

		if resourceIDs[importBlock.ResourceID] {
			t.Errorf("Resource ID %s is not unique", importBlock.ResourceID)

			return nil
		}
		resourceIDs[importBlock.ResourceID] = true
	}

	return importBlocks
}

func getValidatedExpectedImportBlocks(t *testing.T, expectedImportBlocks *[]connector.ImportBlock) *[]connector.ImportBlock {
	t.Helper()

	// Check if provided pointer to expected import blocks is nil, and created an empty slice if so.
	if expectedImportBlocks == nil {
		expectedImportBlocks = &[]connector.ImportBlock{}
	}

	// Make sure the resource name and id in each import block is unique across all import blocks
	resourceNames := map[string]bool{}
	resourceIDs := map[string]bool{}
	for _, importBlock := range *expectedImportBlocks {
		if resourceNames[importBlock.ResourceName] {
			t.Errorf("Resource name %s is not unique", importBlock.ResourceName)
		}
		resourceNames[importBlock.ResourceName] = true

		if resourceIDs[importBlock.ResourceID] {
			t.Errorf("Resource ID %s is not unique", importBlock.ResourceID)
		}
		resourceIDs[importBlock.ResourceID] = true
	}

	return expectedImportBlocks
}

func ValidateImportBlocks(t *testing.T, resource connector.ExportableResource, expectedImportBlocks *[]connector.ImportBlock) {
	t.Helper()

	actualImportBlocks := getValidatedActualImportBlocks(t, resource)
	expectedImportBlocks = getValidatedExpectedImportBlocks(t, expectedImportBlocks)

	expectedImportBlocksMap := map[string]connector.ImportBlock{}
	for _, importBlock := range *expectedImportBlocks {
		expectedImportBlocksMap[importBlock.ResourceName] = importBlock
	}

	// Check number of export blocks
	expectedNumberOfBlocks := len(*expectedImportBlocks)
	actualNumberOfBlocks := len(*actualImportBlocks)
	if actualNumberOfBlocks != expectedNumberOfBlocks {
		t.Errorf("Expected %d import blocks, got %d", expectedNumberOfBlocks, actualNumberOfBlocks)

		return
	}

	// Make sure the import blocks match the expected import blocks
	for _, actualImportBlock := range *actualImportBlocks {
		expectedImportBlock, ok := expectedImportBlocksMap[actualImportBlock.ResourceName]

		if !ok {
			t.Errorf("No matching expected import block for generated import block:\n%s", actualImportBlock.String())

			continue
		}

		if !actualImportBlock.Equals(expectedImportBlock) {
			t.Errorf("Expected import block \n%s\n Got import block \n%s", expectedImportBlock.String(), actualImportBlock.String())
		}
	}
}

// Similar to ValidateImportBlocks, but only checks if the expectedImportBlocks are a subset of the actual import blocks.
// This is useful for resources that have pre-configured resources that are not created by the test.
func ValidateImportBlockSubset(t *testing.T, resource connector.ExportableResource, expectedImportBlocks *[]connector.ImportBlock) {
	t.Helper()

	actualImportBlocks := getValidatedActualImportBlocks(t, resource)
	expectedImportBlocks = getValidatedExpectedImportBlocks(t, expectedImportBlocks)

	actualImportBlocksMap := map[string]connector.ImportBlock{}
	for _, importBlock := range *actualImportBlocks {
		actualImportBlocksMap[importBlock.ResourceName] = importBlock
	}

	// Check number of export blocks
	expectedNumberOfBlocks := len(*expectedImportBlocks)
	actualNumberOfBlocks := len(*actualImportBlocks)
	if actualNumberOfBlocks < expectedNumberOfBlocks {
		t.Errorf("Expected import blocks count (%d) is greater than Actual import blocks count (%d)", expectedNumberOfBlocks, actualNumberOfBlocks)

		return
	}
	if expectedNumberOfBlocks == 0 {
		t.Errorf("Expected import blocks count is 0")

		return
	}

	// For each expected import block, make sure it matches an actual import block
	for _, expectedImportBlock := range *expectedImportBlocks {
		actualImportBlock, ok := actualImportBlocksMap[expectedImportBlock.ResourceName]

		if !ok {
			t.Errorf("No matching actual import block for expected import block:\n%s", expectedImportBlock.String())

			continue
		}

		if !actualImportBlock.Equals(expectedImportBlock) {
			t.Errorf("Expected import block \n%s\n Got import block \n%s", expectedImportBlock.String(), actualImportBlock.String())
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
func WriteStringToPipe(t *testing.T, str string) (reader *os.File) {
	t.Helper()

	reader, writer, err := os.Pipe()
	if err != nil {
		t.Fatal(err)
	}

	if _, err := writer.WriteString(str); err != nil {
		rcErr := reader.Close()
		wcErr := writer.Close()
		t.Fatal(errors.Join(err, rcErr, wcErr))
	}

	// Close the writer to simulate EOF
	if err = writer.Close(); err != nil {
		cErr := reader.Close()
		t.Fatal(errors.Join(err, cErr))
	}

	return reader
}

func CreateX509Certificate() (string, error) {
	serialNumberLimit := new(big.Int).Lsh(big.NewInt(1), 128)
	serialNumber, err := rand.Int(rand.Reader, serialNumberLimit)
	if err != nil {
		return "", fmt.Errorf("failed to generate serial number: %w", err)
	}

	certificateCA := &x509.Certificate{
		SerialNumber: serialNumber,
		Subject: pkix.Name{
			Organization:  []string{"Ping Identity Corporation"},
			Country:       []string{"US"},
			Province:      []string{"CO"},
			Locality:      []string{"Denver"},
			StreetAddress: []string{"1001 17th St"},
			PostalCode:    []string{"80202"},
			CommonName:    "*.pingidentity.com",
		},
		NotBefore:             time.Now(),
		NotAfter:              time.Now().AddDate(1, 0, 0),
		IsCA:                  true,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth, x509.ExtKeyUsageServerAuth},
		KeyUsage:              x509.KeyUsageDigitalSignature | x509.KeyUsageCertSign,
		BasicConstraintsValid: true,
	}

	caPrivKey, err := rsa.GenerateKey(rand.Reader, 4096)
	if err != nil {
		return "", fmt.Errorf("failed to generate private key: %w", err)
	}

	caBytes, err := x509.CreateCertificate(rand.Reader, certificateCA, certificateCA, &caPrivKey.PublicKey, caPrivKey)
	if err != nil {
		return "", fmt.Errorf("failed to create certificate: %w", err)
	}

	caPEM := new(bytes.Buffer)
	err = pem.Encode(caPEM, &pem.Block{
		Type:  "CERTIFICATE",
		Bytes: caBytes,
	})
	if err != nil {
		return "", fmt.Errorf("failed to encode certificate: %w", err)
	}

	return caPEM.String(), nil
}
