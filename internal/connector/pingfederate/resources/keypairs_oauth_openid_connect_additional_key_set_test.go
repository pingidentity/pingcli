package resources_test

import (
	"testing"

	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/common"
	"github.com/pingidentity/pingcli/internal/connector/pingfederate/resources"
	"github.com/pingidentity/pingcli/internal/testing/testutils"
	"github.com/pingidentity/pingcli/internal/testing/testutils_resource"
	"github.com/pingidentity/pingcli/internal/utils"
	client "github.com/pingidentity/pingfederate-go-client/v1210/configurationapi"
)

func TestableResource_PingFederateKeypairsOauthOpenidConnectAdditionalKeySet(t *testing.T) *testutils_resource.TestableResource {
	t.Helper()

	pingfederateClientInfo := testutils.GetPingFederateClientInfo(t)
	return &testutils_resource.TestableResource{
		ClientInfo:         pingfederateClientInfo,
		ExportableResource: resources.AuthenticationApiApplication(pingfederateClientInfo),
		TestResource: testutils_resource.TestResource{
			Dependencies: []testutils_resource.TestResource{
				{
					Dependencies: nil,
					CreateFunc:   createOauthIssuer,
					DeleteFunc:   deleteOauthIssuer,
				},
				{
					Dependencies: nil,
					CreateFunc:   createKeypairsSigningKey,
					DeleteFunc:   deleteKeypairsSigningKey,
				},
			},
			CreateFunc: createKeypairsOauthOpenidConnectAdditionalKeySet,
			DeleteFunc: deleteKeypairsOauthOpenidConnectAdditionalKeySet,
		},
	}
}

func Test_PingFederateKeypairsOauthOpenidConnectAdditionalKeySet(t *testing.T) {
	tr := TestableResource_PingFederateKeypairsOauthOpenidConnectAdditionalKeySet(t)

	creationInfo := tr.CreateResource(t, tr.TestResource)
	defer tr.DeleteResource(t, tr.TestResource)

	expectedImportBlocks := []connector.ImportBlock{
		{
			ResourceType: tr.ExportableResource.ResourceType(),
			ResourceName: creationInfo[testutils_resource.ENUM_NAME],
			ResourceID:   creationInfo[testutils_resource.ENUM_ID],
		},
	}

	testutils.ValidateImportBlocks(t, tr.ExportableResource, &expectedImportBlocks)
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
	result := client.AdditionalKeySet{
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

	request = request.Body(result)

	resource, response, err := request.Execute()
	err = common.HandleClientResponse(response, err, "CreateApplication", resourceType)
	if err != nil {
		t.Fatalf("Failed to create test %s: %v", resourceType, err)
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
	err = common.HandleClientResponse(response, err, "DeleteKeySet", resourceType)
	if err != nil {
		t.Errorf("Failed to delete test %s: %v", resourceType, err)
	}
}
