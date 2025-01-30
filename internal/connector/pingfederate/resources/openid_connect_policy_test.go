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

func TestableResource_PingFederateOpenidConnectPolicy(t *testing.T) *testutils_resource.TestableResource {
	t.Helper()

	pingfederateClientInfo := testutils.GetPingFederateClientInfo(t)
	return &testutils_resource.TestableResource{
		ClientInfo:         pingfederateClientInfo,
		ExportableResource: resources.AuthenticationApiApplication(pingfederateClientInfo),
		TestResource: testutils_resource.TestResource{
			Dependencies: []testutils_resource.TestResource{
				{
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
			},
			CreateFunc: createOpenidConnectPolicy,
			DeleteFunc: deleteOpenidConnectPolicy,
		},
	}
}

func Test_PingFederateOpenidConnectPolicy(t *testing.T) {
	tr := TestableResource_PingFederateOpenidConnectPolicy(t)

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

func createOpenidConnectPolicy(t *testing.T, clientInfo *connector.ClientInfo, strArgs ...string) testutils_resource.ResourceCreationInfo {
	t.Helper()

	if len(strArgs) != 2 {
		t.Fatalf("Unexpected number of arguments provided to createOpenidConnectPolicy(): %v", strArgs)
	}
	resourceType := strArgs[0]
	testAccessTokenManagerId := strArgs[1]

	request := clientInfo.PingFederateApiClient.OauthOpenIdConnectAPI.CreateOIDCPolicy(clientInfo.Context)
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
	err = common.HandleClientResponse(response, err, "CreateApplication", resourceType)
	if err != nil {
		t.Fatalf("Failed to create test %s: %v", resourceType, err)
	}

	return testutils_resource.ResourceCreationInfo{
		testutils_resource.ENUM_ID:   resource.Id,
		testutils_resource.ENUM_NAME: resource.Name,
	}
}

func deleteOpenidConnectPolicy(t *testing.T, clientInfo *connector.ClientInfo, resourceType, id string) {
	t.Helper()

	request := clientInfo.PingFederateApiClient.OauthOpenIdConnectAPI.DeleteOIDCPolicy(clientInfo.Context, id)

	response, err := request.Execute()
	err = common.HandleClientResponse(response, err, "DeleteOIDCPolicy", resourceType)
	if err != nil {
		t.Errorf("Failed to delete test %s: %v", resourceType, err)
	}
}
