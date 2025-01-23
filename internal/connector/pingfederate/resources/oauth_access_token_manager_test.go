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

func Test_PingFederateOauthAccessTokenManager_Export(t *testing.T) {
	PingFederateClientInfo := testutils.GetPingFederateClientInfo(t)
	resource := resources.OauthAccessTokenManager(PingFederateClientInfo)

	testKeyPairId, _, _ := createKeypairsSigningKey(t, PingFederateClientInfo, resource.ResourceType())
	defer deleteKeypairsSigningKey(t, PingFederateClientInfo, resource.ResourceType(), testKeyPairId)

	oauthAccessTokenManagerId, oauthAccessTokenManagerName := createOauthAccessTokenManager(t, PingFederateClientInfo, resource.ResourceType(), testKeyPairId)
	defer deleteOauthAccessTokenManager(t, PingFederateClientInfo, resource.ResourceType(), oauthAccessTokenManagerId)

	expectedImportBlocks := []connector.ImportBlock{
		{
			ResourceType: resource.ResourceType(),
			ResourceName: oauthAccessTokenManagerName,
			ResourceID:   oauthAccessTokenManagerId,
		},
	}

	testutils.ValidateImportBlocks(t, resource, &expectedImportBlocks)
}

func createOauthAccessTokenManager(t *testing.T, clientInfo *connector.PingFederateClientInfo, resourceType, testKeyPairId string) (string, string) {
	t.Helper()

	request := clientInfo.ApiClient.OauthAccessTokenManagersAPI.CreateTokenManager(clientInfo.Context)
	result := client.AccessTokenManager{
		AttributeContract: &client.AccessTokenAttributeContract{
			ExtendedAttributes: []client.AccessTokenAttribute{
				{
					MultiValued: utils.Pointer(false),
					Name:        "testAttribute",
				},
			},
		},
		Configuration: client.PluginConfiguration{
			Fields: []client.ConfigField{
				{
					Name:  "Active Signing Certificate Key ID",
					Value: utils.Pointer("testKeyId"),
				},
				{
					Name:  "JWS Algorithm",
					Value: utils.Pointer("RS256"),
				},
			},
			Tables: []client.ConfigTable{
				{
					Name: "Certificates",
					Rows: []client.ConfigRow{
						{
							DefaultRow: utils.Pointer(false),
							Fields: []client.ConfigField{
								{
									Name:  "Key ID",
									Value: utils.Pointer("testKeyId"),
								},
								{
									Name:  "Certificate",
									Value: &testKeyPairId,
								},
							},
						},
					},
				},
			},
		},
		Id:   "TestOauthAccessTokenManagerId",
		Name: "TestOauthAccessTokenManagerName",
		PluginDescriptorRef: client.ResourceLink{
			Id: "com.pingidentity.pf.access.token.management.plugins.JwtBearerAccessTokenManagementPlugin",
		},
	}

	request = request.Body(result)

	resource, response, err := request.Execute()
	err = common.HandleClientResponse(response, err, "CreateTokenManager", resourceType)
	if err != nil {
		t.Fatalf("Failed to create test %s: %v", resourceType, err)
	}

	return resource.Id, resource.Name
}

func deleteOauthAccessTokenManager(t *testing.T, clientInfo *connector.PingFederateClientInfo, resourceType, id string) {
	t.Helper()

	request := clientInfo.ApiClient.OauthAccessTokenManagersAPI.DeleteTokenManager(clientInfo.Context, id)

	response, err := request.Execute()
	err = common.HandleClientResponse(response, err, "DeleteTokenManager", resourceType)
	if err != nil {
		t.Errorf("Failed to delete test %s: %v", resourceType, err)
	}
}
