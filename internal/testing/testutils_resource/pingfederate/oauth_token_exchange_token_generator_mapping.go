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

func TestableResource_PingFederateOauthTokenExchangeTokenGeneratorMapping(t *testing.T, clientInfo *connector.ClientInfo) *testutils_resource.TestableResource {
	t.Helper()

	return &testutils_resource.TestableResource{
		ClientInfo: clientInfo,
		CreateFunc: createOauthTokenExchangeTokenGeneratorMapping,
		DeleteFunc: deleteOauthTokenExchangeTokenGeneratorMapping,
		Dependencies: []*testutils_resource.TestableResource{
			TestableResource_PingFederateOauthTokenExchangeProcessorPolicy(t, clientInfo),
			TestableResource_PingFederateSpTokenGenerator(t, clientInfo),
		},
		ExportableResource: resources.OauthTokenExchangeTokenGeneratorMapping(clientInfo),
	}
}

func createOauthTokenExchangeTokenGeneratorMapping(t *testing.T, clientInfo *connector.ClientInfo, strArgs ...string) testutils_resource.ResourceCreationInfo {
	t.Helper()

	if len(strArgs) != 3 {
		t.Fatalf("Unexpected number of arguments provided to createOauthTokenExchangeTokenGeneratorMapping(): %v", strArgs)
	}
	resourceType := strArgs[0]
	testProcessorPolicyId := strArgs[1]
	testTokenGeneratorId := strArgs[2]

	request := clientInfo.PingFederateApiClient.OauthTokenExchangeTokenGeneratorMappingsAPI.CreateTokenGeneratorMapping(clientInfo.PingFederateContext)
	clientStruct := client.ProcessorPolicyToGeneratorMapping{
		AttributeContractFulfillment: map[string]client.AttributeFulfillmentValue{
			"SAML_SUBJECT": {
				Source: client.SourceTypeIdKey{
					Type: "NO_MAPPING",
				},
			},
		},
		Id:       utils.Pointer(testProcessorPolicyId + "|" + testTokenGeneratorId),
		SourceId: testProcessorPolicyId,
		TargetId: testTokenGeneratorId,
	}

	request = request.Body(clientStruct)

	resource, response, err := request.Execute()
	ok, err := common.HandleClientResponse(response, err, "CreateTokenGeneratorMapping", resourceType)
	if err != nil {
		t.Fatalf("Failed to execute client function\nResponse Status: %s\nResponse Body: %s\nError: %v", response.Status, response.Body, err)
	}
	if !ok {
		t.Fatalf("Failed to execute client function\nResponse Status: %s\nResponse Body: %s", response.Status, response.Body)
	}

	return testutils_resource.ResourceCreationInfo{
		testutils_resource.ENUM_ID:        *resource.Id,
		testutils_resource.ENUM_SOURCE_ID: testProcessorPolicyId,
		testutils_resource.ENUM_TARGET_ID: testTokenGeneratorId,
	}
}

func deleteOauthTokenExchangeTokenGeneratorMapping(t *testing.T, clientInfo *connector.ClientInfo, resourceType, id string) {
	t.Helper()

	request := clientInfo.PingFederateApiClient.OauthTokenExchangeTokenGeneratorMappingsAPI.DeleteTokenGeneratorMappingById(clientInfo.PingFederateContext, id)

	response, err := request.Execute()
	ok, err := common.HandleClientResponse(response, err, "DeleteTokenGeneratorMappingById", resourceType)
	if err != nil {
		t.Fatalf("Failed to execute client function\nResponse Status: %s\nResponse Body: %s\nError: %v", response.Status, response.Body, err)
	}
	if !ok {
		t.Fatalf("Failed to execute client function\nResponse Status: %s\nResponse Body: %s", response.Status, response.Body)
	}
}
