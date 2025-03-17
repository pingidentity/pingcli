package pingfederate

import (
	"testing"

	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/common"
	"github.com/pingidentity/pingcli/internal/testing/testutils_resource"
	"github.com/pingidentity/pingcli/internal/utils"
	client "github.com/pingidentity/pingfederate-go-client/v1220/configurationapi"
)

func TestableResource_PingFederateOauthTokenExchangeProcessorPolicy(t *testing.T, clientInfo *connector.ClientInfo) *testutils_resource.TestableResource {
	t.Helper()

	return &testutils_resource.TestableResource{
		ClientInfo: clientInfo,
		CreateFunc: createOauthTokenExchangeProcessorPolicy,
		DeleteFunc: deleteOauthTokenExchangeProcessorPolicy,
		Dependencies: []*testutils_resource.TestableResource{
			TestableResource_PingFederateIdpTokenProcessor(t, clientInfo),
		},
		ExportableResource: nil,
	}
}

func createOauthTokenExchangeProcessorPolicy(t *testing.T, clientInfo *connector.ClientInfo, strArgs ...string) testutils_resource.ResourceCreationInfo {
	t.Helper()

	if len(strArgs) != 2 {
		t.Fatalf("Unexpected number of arguments provided to createOauthTokenExchangeProcessorPolicy(): %v", strArgs)
	}
	resourceType := strArgs[0]
	testTokenProcessorId := strArgs[1]

	request := clientInfo.PingFederateApiClient.OauthTokenExchangeProcessorAPI.CreateOauthTokenExchangeProcessorPolicy(clientInfo.PingFederateContext)
	result := client.TokenExchangeProcessorPolicy{
		ActorTokenRequired: utils.Pointer(false),
		AttributeContract: client.TokenExchangeProcessorAttributeContract{
			CoreAttributes: []client.TokenExchangeProcessorAttribute{
				{
					Name: "subject",
				},
			},
		},
		Id:   "TestProcessorPolicyId",
		Name: "TestProcessorPolicyName",
		ProcessorMappings: []client.TokenExchangeProcessorMapping{
			{
				AttributeContractFulfillment: map[string]client.AttributeFulfillmentValue{
					"subject": {
						Source: client.SourceTypeIdKey{
							Type: "NO_MAPPING",
						},
					},
				},
				SubjectTokenType: "TestTokenType",
				SubjectTokenProcessor: client.ResourceLink{
					Id: testTokenProcessorId,
				},
			},
		},
	}

	request = request.Body(result)

	resource, response, err := request.Execute()
	ok, err := common.HandleClientResponse(response, err, "CreateOauthTokenExchangeProcessorPolicy", resourceType)
	if err != nil {
		t.Fatalf("Failed to execute client function\nResponse Status: %s\nResponse Body: %s\nError: %v", response.Status, response.Body, err)
	}
	if !ok {
		t.Fatalf("Failed to execute client function\nResponse Status: %s\nResponse Body: %s", response.Status, response.Body)
	}

	return testutils_resource.ResourceCreationInfo{
		testutils_resource.ENUM_ID: resource.Id,
	}
}

func deleteOauthTokenExchangeProcessorPolicy(t *testing.T, clientInfo *connector.ClientInfo, resourceType, id string) {
	t.Helper()

	request := clientInfo.PingFederateApiClient.OauthTokenExchangeProcessorAPI.DeleteOauthTokenExchangeProcessorPolicyy(clientInfo.PingFederateContext, id)

	response, err := request.Execute()
	ok, err := common.HandleClientResponse(response, err, "DeleteOauthTokenExchangeProcessorPolicy", resourceType)
	if err != nil {
		t.Fatalf("Failed to execute client function\nResponse Status: %s\nResponse Body: %s\nError: %v", response.Status, response.Body, err)
	}
	if !ok {
		t.Fatalf("Failed to execute client function\nResponse Status: %s\nResponse Body: %s", response.Status, response.Body)
	}
}
