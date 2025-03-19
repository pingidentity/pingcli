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

func TestableResource_PingFederateOpenidConnectPolicy(t *testing.T, clientInfo *connector.ClientInfo) *testutils_resource.TestableResource {
	t.Helper()

	return &testutils_resource.TestableResource{
		ClientInfo: clientInfo,
		CreateFunc: createOpenidConnectPolicy,
		DeleteFunc: deleteOpenidConnectPolicy,
		Dependencies: []*testutils_resource.TestableResource{
			TestableResource_PingFederateOauthAccessTokenManager(t, clientInfo),
		},
		ExportableResource: resources.OpenidConnectPolicy(clientInfo),
	}
}

func createOpenidConnectPolicy(t *testing.T, clientInfo *connector.ClientInfo, strArgs ...string) testutils_resource.ResourceCreationInfo {
	t.Helper()

	if len(strArgs) != 2 {
		t.Fatalf("Unexpected number of arguments provided to createOpenidConnectPolicy(): %v", strArgs)
	}
	resourceType := strArgs[0]
	testAccessTokenManagerId := strArgs[1]

	request := clientInfo.PingFederateApiClient.OauthOpenIdConnectAPI.CreateOIDCPolicy(clientInfo.PingFederateContext)
	clientStruct := client.OpenIdConnectPolicy{
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

	request = request.Body(clientStruct)

	resource, response, err := request.Execute()
	ok, err := common.HandleClientResponse(response, err, "CreateOIDCPolicy", resourceType)
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

func deleteOpenidConnectPolicy(t *testing.T, clientInfo *connector.ClientInfo, resourceType, id string) {
	t.Helper()

	request := clientInfo.PingFederateApiClient.OauthOpenIdConnectAPI.DeleteOIDCPolicy(clientInfo.PingFederateContext, id)

	response, err := request.Execute()
	ok, err := common.HandleClientResponse(response, err, "DeleteOIDCPolicy", resourceType)
	if err != nil {
		t.Fatalf("Failed to execute client function\nResponse Status: %s\nResponse Body: %s\nError: %v", response.Status, response.Body, err)
	}
	if !ok {
		t.Fatalf("Failed to execute client function\nResponse Status: %s\nResponse Body: %s", response.Status, response.Body)
	}
}
