// Copyright © 2025 Ping Identity Corporation
// Code generated by ping-cli-generator

package pingfederate_testable_resources

import (
	"testing"

	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/common"
	"github.com/pingidentity/pingcli/internal/connector/pingfederate/resources"
	"github.com/pingidentity/pingcli/internal/testing/testutils_resource"
	"github.com/pingidentity/pingcli/internal/utils"
	client "github.com/pingidentity/pingfederate-go-client/v1220/configurationapi"
)

func OauthAccessTokenManager(t *testing.T, clientInfo *connector.ClientInfo) *testutils_resource.TestableResource {
	t.Helper()

	return &testutils_resource.TestableResource{
		ClientInfo: clientInfo,
		CreateFunc: createOauthAccessTokenManager,
		DeleteFunc: deleteOauthAccessTokenManager,
		Dependencies: []*testutils_resource.TestableResource{
			KeypairsSigningKey(t, clientInfo),
		},
		ExportableResource: resources.OauthAccessTokenManager(clientInfo),
	}
}

func createOauthAccessTokenManager(t *testing.T, clientInfo *connector.ClientInfo, resourceType string, strArgs ...string) testutils_resource.ResourceInfo {
	t.Helper()

	if len(strArgs) != 1 {
		t.Errorf("Unexpected number of arguments provided to createOauthAccessTokenManager(): %v", strArgs)

		return testutils_resource.ResourceInfo{}
	}
	testKeyPairId := strArgs[0]

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
		t.Errorf("Failed to execute client function\nResponse Status: %s\nResponse Body: %s\nError: %v", response.Status, response.Body, err)

		return testutils_resource.ResourceInfo{}
	}
	if !ok {
		t.Errorf("Failed to execute client function\nResponse Status: %s\nResponse Body: %s", response.Status, response.Body)

		return testutils_resource.ResourceInfo{}
	}

	return testutils_resource.ResourceInfo{
		DeletionIds: []string{
			resource.Id,
		},
		CreationInfo: map[testutils_resource.ResourceCreationInfoType]string{
			testutils_resource.ENUM_ID:   resource.Id,
			testutils_resource.ENUM_NAME: resource.Name,
		},
	}
}

func deleteOauthAccessTokenManager(t *testing.T, clientInfo *connector.ClientInfo, resourceType string, ids ...string) {
	t.Helper()

	if len(ids) != 1 {
		t.Errorf("Unexpected number of arguments provided to deleteOauthAccessTokenManager(): %v", ids)

		return
	}

	request := clientInfo.PingFederateApiClient.OauthAccessTokenManagersAPI.DeleteTokenManager(clientInfo.PingFederateContext, ids[0])

	response, err := request.Execute()
	ok, err := common.HandleClientResponse(response, err, "DeleteTokenManager", resourceType)
	if err != nil {
		t.Errorf("Failed to execute client function\nResponse Status: %s\nResponse Body: %s\nError: %v", response.Status, response.Body, err)

		return
	}
	if !ok {
		t.Errorf("Failed to execute client function\nResponse Status: %s\nResponse Body: %s", response.Status, response.Body)

		return
	}
}
