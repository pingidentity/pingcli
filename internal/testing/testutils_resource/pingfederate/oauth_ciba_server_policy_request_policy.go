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

func TestableResource_PingFederateOauthCibaServerPolicyRequestPolicy(t *testing.T, clientInfo *connector.ClientInfo) *testutils_resource.TestableResource {
	t.Helper()

	return &testutils_resource.TestableResource{
		ClientInfo: clientInfo,
		CreateFunc: createOauthCibaServerPolicyRequestPolicy,
		DeleteFunc: deleteOauthCibaServerPolicyRequestPolicy,
		Dependencies: []*testutils_resource.TestableResource{
			TestableResource_PingFederateOutOfBandAuthPlugins(t, clientInfo),
		},
		ExportableResource: resources.OauthCibaServerPolicyRequestPolicy(clientInfo),
	}
}

func createOauthCibaServerPolicyRequestPolicy(t *testing.T, clientInfo *connector.ClientInfo, strArgs ...string) testutils_resource.ResourceCreationInfo {
	t.Helper()

	if len(strArgs) != 2 {
		t.Fatalf("Unexpected number of arguments provided to createOauthCibaServerPolicyRequestPolicy(): %v", strArgs)
	}
	resourceType := strArgs[0]
	testAuthenticatorId := strArgs[1]

	request := clientInfo.PingFederateApiClient.OauthCibaServerPolicyAPI.CreateCibaServerPolicy(clientInfo.PingFederateContext)
	clientStruct := client.RequestPolicy{
		AllowUnsignedLoginHintToken: utils.Pointer(false),
		AuthenticatorRef: client.ResourceLink{
			Id: testAuthenticatorId,
		},
		Id: "TestRequestPolicyId",
		IdentityHintContract: client.IdentityHintContract{
			CoreAttributes: []client.IdentityHintAttribute{
				{
					Name: "IDENTITY_HINT_SUBJECT",
				},
			},
		},
		IdentityHintContractFulfillment: &client.AttributeMapping{
			AttributeContractFulfillment: map[string]client.AttributeFulfillmentValue{
				"IDENTITY_HINT_SUBJECT": {
					Source: client.SourceTypeIdKey{
						Type: "REQUEST",
					},
					Value: "IDENTITY_HINT_SUBJECT",
				},
			},
		},
		IdentityHintMapping: &client.AttributeMapping{
			AttributeContractFulfillment: map[string]client.AttributeFulfillmentValue{
				"subject": {
					Source: client.SourceTypeIdKey{
						Type: "NO_MAPPING",
					},
				},
				"USER_KEY": {
					Source: client.SourceTypeIdKey{
						Type: "NO_MAPPING",
					},
				},
			},
		},
		Name:                        "TestRequestPolicyName",
		RequireTokenForIdentityHint: utils.Pointer(false),
		TransactionLifetime:         utils.Pointer(int64(120)),
	}

	request = request.Body(clientStruct)

	resource, response, err := request.Execute()
	ok, err := common.HandleClientResponse(response, err, "CreateCibaServerPolicy", resourceType)
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

func deleteOauthCibaServerPolicyRequestPolicy(t *testing.T, clientInfo *connector.ClientInfo, resourceType, id string) {
	t.Helper()

	request := clientInfo.PingFederateApiClient.OauthCibaServerPolicyAPI.DeleteCibaServerPolicy(clientInfo.PingFederateContext, id)

	response, err := request.Execute()
	ok, err := common.HandleClientResponse(response, err, "DeleteCibaServerPolicy", resourceType)
	if err != nil {
		t.Fatalf("Failed to execute client function\nResponse Status: %s\nResponse Body: %s\nError: %v", response.Status, response.Body, err)
	}
	if !ok {
		t.Fatalf("Failed to execute client function\nResponse Status: %s\nResponse Body: %s", response.Status, response.Body)
	}
}
