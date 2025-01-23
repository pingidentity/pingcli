package resources_test

import (
	"testing"

	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/common"
	"github.com/pingidentity/pingcli/internal/connector/pingfederate/resources"
	"github.com/pingidentity/pingcli/internal/testing/testutils"
	"github.com/pingidentity/pingcli/internal/utils"
	client "github.com/pingidentity/pingfederate-go-client/v1210/configurationapi"
)

func Test_PingFederateKeypairsOauthOpenidConnectAdditionalKeySet_Export(t *testing.T) {
	PingFederateClientInfo := testutils.GetPingFederateClientInfo(t)
	resource := resources.KeypairsOauthOpenidConnectAdditionalKeySet(PingFederateClientInfo)

	testKeyPairId, _, _ := createKeypairsSigningKey(t, PingFederateClientInfo, resource.ResourceType())
	defer deleteKeypairsSigningKey(t, PingFederateClientInfo, resource.ResourceType(), testKeyPairId)

	testOauthIssuerId, _ := createOauthIssuer(t, PingFederateClientInfo, resource.ResourceType())
	defer deleteOauthIssuer(t, PingFederateClientInfo, resource.ResourceType(), testOauthIssuerId)

	keypairsOauthOpenidConnectAdditionalKeySetId, keypairsOauthOpenidConnectAdditionalKeySetName := createKeypairsOauthOpenidConnectAdditionalKeySet(t, PingFederateClientInfo, resource.ResourceType(), testKeyPairId, testOauthIssuerId)
	defer deleteKeypairsOauthOpenidConnectAdditionalKeySet(t, PingFederateClientInfo, resource.ResourceType(), keypairsOauthOpenidConnectAdditionalKeySetId)

	expectedImportBlocks := []connector.ImportBlock{
		{
			ResourceType: resource.ResourceType(),
			ResourceName: keypairsOauthOpenidConnectAdditionalKeySetName,
			ResourceID:   keypairsOauthOpenidConnectAdditionalKeySetId,
		},
	}

	testutils.ValidateImportBlocks(t, resource, &expectedImportBlocks)
}

func createKeypairsOauthOpenidConnectAdditionalKeySet(t *testing.T, clientInfo *connector.PingFederateClientInfo, resourceType, testKeyPairId, testOauthIssuerId string) (string, string) {
	t.Helper()

	request := clientInfo.ApiClient.KeyPairsOauthOpenIdConnectAPI.CreateKeySet(clientInfo.Context)
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
	err = common.HandleClientResponse(response, err, "CreateKeySet", resourceType)
	if err != nil {
		t.Fatalf("Failed to create test %s: %v", resourceType, err)
	}

	return *resource.Id, resource.Name
}

func deleteKeypairsOauthOpenidConnectAdditionalKeySet(t *testing.T, clientInfo *connector.PingFederateClientInfo, resourceType, id string) {
	t.Helper()

	request := clientInfo.ApiClient.KeyPairsOauthOpenIdConnectAPI.DeleteKeySet(clientInfo.Context, id)

	response, err := request.Execute()
	err = common.HandleClientResponse(response, err, "DeleteKeySet", resourceType)
	if err != nil {
		t.Errorf("Failed to delete test %s: %v", resourceType, err)
	}
}
