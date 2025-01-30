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

func TestableResource_PingFederateOauthAccessTokenManager(t *testing.T) *testutils_resource.TestableResource {
	t.Helper()

	pingfederateClientInfo := testutils.GetPingFederateClientInfo(t)
	return &testutils_resource.TestableResource{
		ClientInfo:         pingfederateClientInfo,
		ExportableResource: resources.AuthenticationApiApplication(pingfederateClientInfo),
		TestResource: testutils_resource.TestResource{
			Dependencies: []testutils_resource.TestResource{
				{
					Dependencies: nil,
					CreateFunc:   createKeypairsSigningKey,
					DeleteFunc:   deleteKeypairsSigningKey,
				},
			},
			CreateFunc: createOauthAccessTokenManager,
			DeleteFunc: deleteOauthAccessTokenManager,
		},
	}
}

func Test_PingFederateOauthAccessTokenManager(t *testing.T) {
	tr := TestableResource_PingFederateOauthAccessTokenManager(t)

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

func createOauthAccessTokenManager(t *testing.T, clientInfo *connector.ClientInfo, strArgs ...string) testutils_resource.ResourceCreationInfo {
	t.Helper()

	if len(strArgs) != 2 {
		t.Fatalf("Unexpected number of arguments provided to createOauthAccessTokenManager(): %v", strArgs)
	}
	resourceType := strArgs[0]
	testKeyPairId := strArgs[1]

	request := clientInfo.PingFederateApiClient.OauthAccessTokenManagersAPI.CreateTokenManager(clientInfo.Context)
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
	err = common.HandleClientResponse(response, err, "CreateApplication", resourceType)
	if err != nil {
		t.Fatalf("Failed to create test %s: %v", resourceType, err)
	}

	return testutils_resource.ResourceCreationInfo{
		testutils_resource.ENUM_ID:   resource.Id,
		testutils_resource.ENUM_NAME: resource.Name,
	}
}

func deleteOauthAccessTokenManager(t *testing.T, clientInfo *connector.ClientInfo, resourceType, id string) {
	t.Helper()

	request := clientInfo.PingFederateApiClient.OauthAccessTokenManagersAPI.DeleteTokenManager(clientInfo.Context, id)

	response, err := request.Execute()
	err = common.HandleClientResponse(response, err, "DeleteTokenManager", resourceType)
	if err != nil {
		t.Errorf("Failed to delete test %s: %v", resourceType, err)
	}
}
