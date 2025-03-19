// Copyright © 2025 Ping Identity Corporation

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

func TestableResource_PingFederateOauthAccessTokenManager(t *testing.T, clientInfo *connector.ClientInfo) *testutils_resource.TestableResource {
	t.Helper()

	return &testutils_resource.TestableResource{
		ClientInfo: clientInfo,
		CreateFunc: createOauthAccessTokenManager,
		DeleteFunc: deleteOauthAccessTokenManager,
		Dependencies: []*testutils_resource.TestableResource{
			TestableResource_PingFederateKeypairsSigningKey(t, clientInfo),
		},
		ExportableResource: resources.OauthAccessTokenManager(clientInfo),
	}
}

func createOauthAccessTokenManager(t *testing.T, clientInfo *connector.ClientInfo, strArgs ...string) testutils_resource.ResourceCreationInfo {
	t.Helper()

	if len(strArgs) != 2 {
		t.Fatalf("Unexpected number of arguments provided to createOauthAccessTokenManager(): %v", strArgs)
	}
	resourceType := strArgs[0]
	testKeyPairId := strArgs[1]

	request := clientInfo.PingFederateApiClient.OauthAccessTokenManagersAPI.CreateTokenManager(clientInfo.PingFederateContext)
	clientStruct := client.AccessTokenManager{
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

	request = request.Body(clientStruct)

	resource, response, err := request.Execute()
	ok, err := common.HandleClientResponse(response, err, "CreateTokenManager", resourceType)
	if err != nil {
		t.Fatalf("Failed to execute client function\nResponse Status: %s\nResponse Body: %s\nError: %v", response.Status, response.Body, err)
	}
	if !ok {
		t.Fatalf("Failed to execute client function\nResponse Status: %s\nResponse Body: %s", response.Status, response.Body)
	}

	return testutils_resource.ResourceCreationInfo{
		testutils_resource.ENUM_ID:   resource.Id,
		testutils_resource.ENUM_NAME: resource.Name,
	}
}

func deleteOauthAccessTokenManager(t *testing.T, clientInfo *connector.ClientInfo, resourceType, id string) {
	t.Helper()

	request := clientInfo.PingFederateApiClient.OauthAccessTokenManagersAPI.DeleteTokenManager(clientInfo.PingFederateContext, id)

	response, err := request.Execute()
	ok, err := common.HandleClientResponse(response, err, "DeleteTokenManager", resourceType)
	if err != nil {
		t.Fatalf("Failed to execute client function\nResponse Status: %s\nResponse Body: %s\nError: %v", response.Status, response.Body, err)
	}
	if !ok {
		t.Fatalf("Failed to execute client function\nResponse Status: %s\nResponse Body: %s", response.Status, response.Body)
	}
}
