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

func Test_PingFederateOpenidConnectPolicy_Export(t *testing.T) {
	PingFederateClientInfo := testutils.GetPingFederateClientInfo(t)
	resource := resources.OpenidConnectPolicy(PingFederateClientInfo)

	testKeyPairId, _, _ := createKeypairsSigningKey(t, PingFederateClientInfo, resource.ResourceType())
	defer deleteKeypairsSigningKey(t, PingFederateClientInfo, resource.ResourceType(), testKeyPairId)

	testAccessTokenManagerId, _ := createOauthAccessTokenManager(t, PingFederateClientInfo, resource.ResourceType(), testKeyPairId)
	defer deleteOauthAccessTokenManager(t, PingFederateClientInfo, resource.ResourceType(), testAccessTokenManagerId)

	openidConnectPolicyId, openidConnectPolicyName := createOpenidConnectPolicy(t, PingFederateClientInfo, resource.ResourceType(), testAccessTokenManagerId)
	defer deleteOpenidConnectPolicy(t, PingFederateClientInfo, resource.ResourceType(), openidConnectPolicyId)

	expectedImportBlocks := []connector.ImportBlock{
		{
			ResourceType: resource.ResourceType(),
			ResourceName: openidConnectPolicyName,
			ResourceID:   openidConnectPolicyId,
		},
	}

	testutils.ValidateImportBlocks(t, resource, &expectedImportBlocks)
}

func createOpenidConnectPolicy(t *testing.T, clientInfo *connector.PingFederateClientInfo, resourceType, testAccessTokenManagerId string) (string, string) {
	t.Helper()

	request := clientInfo.ApiClient.OauthOpenIdConnectAPI.CreateOIDCPolicy(clientInfo.Context)
	result := client.OpenIdConnectPolicy{
		AccessTokenManagerRef: client.ResourceLink{
			Id: testAccessTokenManagerId,
		},
		AttributeContract: client.OpenIdConnectAttributeContract{
			CoreAttributes: []client.OpenIdConnectAttribute{
				{
					MultiValued: utils.Pointer(false),
					Name:        "sub",
				},
			},
		},
		AttributeMapping: client.AttributeMapping{
			AttributeContractFulfillment: map[string]client.AttributeFulfillmentValue{
				"sub": {
					Source: client.SourceTypeIdKey{
						Type: "NO_MAPPING",
					},
				},
			},
		},
		Id:   "TestOIDCPolicyId",
		Name: "TestOIDCPolicyName",
	}

	request = request.Body(result)

	resource, response, err := request.Execute()
	err = common.HandleClientResponse(response, err, "CreateOIDCPolicy", resourceType)
	if err != nil {
		t.Fatalf("Failed to create test %s: %v", resourceType, err)
	}

	return resource.Id, resource.Name
}

func deleteOpenidConnectPolicy(t *testing.T, clientInfo *connector.PingFederateClientInfo, resourceType, id string) {
	t.Helper()

	request := clientInfo.ApiClient.OauthOpenIdConnectAPI.DeleteOIDCPolicy(clientInfo.Context, id)

	response, err := request.Execute()
	err = common.HandleClientResponse(response, err, "DeleteOIDCPolicy", resourceType)
	if err != nil {
		t.Errorf("Failed to delete test %s: %v", resourceType, err)
	}
}
