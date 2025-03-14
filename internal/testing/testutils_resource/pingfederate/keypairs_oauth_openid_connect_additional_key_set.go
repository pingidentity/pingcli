package pingfederate

import (
	"testing"

	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/common"
	"github.com/pingidentity/pingcli/internal/connector/pingfederate/resources"
	"github.com/pingidentity/pingcli/internal/testing/testutils_resource"
	"github.com/pingidentity/pingcli/internal/utils"
	client "github.com/pingidentity/pingfederate-go-client/v1210/configurationapi"
)

func TestableResource_PingFederateKeypairsOauthOpenidConnectAdditionalKeySet(t *testing.T, clientInfo *connector.ClientInfo) testutils_resource.TestableResource {
	t.Helper()

	return testutils_resource.TestableResource{
		ClientInfo: clientInfo,
		CreateFunc: createKeypairsOauthOpenidConnectAdditionalKeySet,
		DeleteFunc: deleteKeypairsOauthOpenidConnectAdditionalKeySet,
		Dependencies: []testutils_resource.TestableResource{
			TestableResource_PingFederateOauthIssuer(t, clientInfo),
			TestableResource_PingFederateKeypairsSigningKey(t, clientInfo),
		},
		ExportableResource: resources.KeypairsOauthOpenidConnectAdditionalKeySet(clientInfo),
	}
}

func createKeypairsOauthOpenidConnectAdditionalKeySet(t *testing.T, clientInfo *connector.ClientInfo, strArgs ...string) testutils_resource.ResourceCreationInfo {
	t.Helper()

	if len(strArgs) != 3 {
		t.Fatalf("Unexpected number of arguments provided to createKeypairsOauthOpenidConnectAdditionalKeySet(): %v", strArgs)
	}
	resourceType := strArgs[0]
	testOauthIssuerId := strArgs[1]
	testKeyPairId := strArgs[2]

	request := clientInfo.PingFederateApiClient.KeyPairsOauthOpenIdConnectAPI.CreateKeySet(clientInfo.Context)
	clientStruct := client.AdditionalKeySet{
		Id: utils.Pointer("TestAdditionalKeySetId"),
		Issuers: []client.ResourceLink{
			{
				Id: testOauthIssuerId,
			},
		},
		Name: "TestAdditionalKeySetName",
		SigningKeys: client.SigningKeys{
			RsaActiveCertRef: &client.ResourceLink{
				Id: testKeyPairId,
			},
		},
	}

	request = request.Body(clientStruct)

	resource, response, err := request.Execute()
	ok, err := common.HandleClientResponse(response, err, "CreateKeySet", resourceType)
	if err != nil {
		t.Fatalf("Failed to create test %s: %v", resourceType, err)
	}
	if !ok {
		t.Fatalf("Failed to create test %s: non-ok response", resourceType)
	}

	return testutils_resource.ResourceCreationInfo{
		testutils_resource.ENUM_ID:   *resource.Id,
		testutils_resource.ENUM_NAME: resource.Name,
	}
}

func deleteKeypairsOauthOpenidConnectAdditionalKeySet(t *testing.T, clientInfo *connector.ClientInfo, resourceType, id string) {
	t.Helper()

	request := clientInfo.PingFederateApiClient.KeyPairsOauthOpenIdConnectAPI.DeleteKeySet(clientInfo.Context, id)

	response, err := request.Execute()
	ok, err := common.HandleClientResponse(response, err, "DeleteKeySet", resourceType)
	if err != nil {
		t.Errorf("Failed to delete test %s: %v", resourceType, err)
	}
	if !ok {
		t.Fatalf("Failed to delete test %s: non-ok response", resourceType)
	}
}
