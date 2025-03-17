package pingfederate

import (
	"testing"

	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/common"
	"github.com/pingidentity/pingcli/internal/connector/pingfederate/resources"
	"github.com/pingidentity/pingcli/internal/testing/testutils_resource"
	"github.com/pingidentity/pingcli/internal/utils"
	client "github.com/pingidentity/pingfederate-go-client/v1220/configurationapi"
)

func TestableResource_PingFederateKeypairsSigningKey(t *testing.T, clientInfo *connector.ClientInfo) *testutils_resource.TestableResource {
	t.Helper()

	return &testutils_resource.TestableResource{
		ClientInfo:         clientInfo,
		CreateFunc:         createKeypairsSigningKey,
		DeleteFunc:         deleteKeypairsSigningKey,
		Dependencies:       nil,
		ExportableResource: resources.IdpAdapter(clientInfo),
	}
}

func createKeypairsSigningKey(t *testing.T, clientInfo *connector.ClientInfo, strArgs ...string) testutils_resource.ResourceCreationInfo {
	t.Helper()

	if len(strArgs) != 1 {
		t.Fatalf("Unexpected number of arguments provided to createIdpAdapter(): %v", strArgs)
	}
	resourceType := strArgs[0]

	request := clientInfo.PingFederateApiClient.KeyPairsSigningAPI.CreateSigningKeyPair(clientInfo.PingFederateContext)
	result := client.NewKeyPairSettings{
		City:               utils.Pointer("Denver"),
		CommonName:         "*.pingidentity.com",
		Country:            "US",
		Id:                 utils.Pointer("testkeypairid"),
		KeyAlgorithm:       "RSA",
		KeySize:            utils.Pointer(int64(2048)),
		Organization:       "Ping Identity Corporation",
		SignatureAlgorithm: utils.Pointer("SHA256withRSA"),
		State:              utils.Pointer("CO"),
		ValidDays:          365,
	}

	request = request.Body(result)

	resource, response, err := request.Execute()
	ok, err := common.HandleClientResponse(response, err, "CreateKeyPair", resourceType)
	if err != nil {
		t.Fatalf("Failed to execute client function\nResponse Status: %s\nResponse Body: %s\nError: %v", response.Status, response.Body, err)
	}
	if !ok {
		t.Fatalf("Failed to execute client function\nResponse Status: %s\nResponse Body: %s", response.Status, response.Body)
	}

	return testutils_resource.ResourceCreationInfo{
		testutils_resource.ENUM_ID:            *resource.Id,
		testutils_resource.ENUM_ISSUER_DN:     *resource.IssuerDN,
		testutils_resource.ENUM_SERIAL_NUMBER: *resource.SerialNumber,
	}
}

func deleteKeypairsSigningKey(t *testing.T, clientInfo *connector.ClientInfo, resourceType, id string) {
	t.Helper()

	request := clientInfo.PingFederateApiClient.KeyPairsSigningAPI.DeleteSigningKeyPair(clientInfo.PingFederateContext, id)

	response, err := request.Execute()
	ok, err := common.HandleClientResponse(response, err, "DeleteSigningKeyPair", resourceType)
	if err != nil {
		t.Fatalf("Failed to execute client function\nResponse Status: %s\nResponse Body: %s\nError: %v", response.Status, response.Body, err)
	}
	if !ok {
		t.Fatalf("Failed to execute client function\nResponse Status: %s\nResponse Body: %s", response.Status, response.Body)
	}
}
