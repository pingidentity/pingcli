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

func TestableResource_PingFederateTokenProcessorToTokenGeneratorMapping(t *testing.T, clientInfo *connector.ClientInfo) *testutils_resource.TestableResource {
	t.Helper()

	return &testutils_resource.TestableResource{
		ClientInfo: clientInfo,
		CreateFunc: createTokenProcessorToTokenGeneratorMapping,
		DeleteFunc: deleteTokenProcessorToTokenGeneratorMapping,
		Dependencies: []*testutils_resource.TestableResource{
			TestableResource_PingFederateIdpTokenProcessor(t, clientInfo),
			TestableResource_PingFederateSpTokenGenerator(t, clientInfo),
		},
		ExportableResource: resources.TokenProcessorToTokenGeneratorMapping(clientInfo),
	}
}

func createTokenProcessorToTokenGeneratorMapping(t *testing.T, clientInfo *connector.ClientInfo, strArgs ...string) testutils_resource.ResourceCreationInfo {
	t.Helper()

	if len(strArgs) != 3 {
		t.Fatalf("Unexpected number of arguments provided to createTokenProcessorToTokenGeneratorMapping(): %v", strArgs)
	}
	resourceType := strArgs[0]
	testTokenProcessorId := strArgs[1]
	testTokenGeneratorId := strArgs[2]

	request := clientInfo.PingFederateApiClient.TokenProcessorToTokenGeneratorMappingsAPI.CreateTokenToTokenMapping(clientInfo.PingFederateContext)
	clientStruct := client.TokenToTokenMapping{
		AttributeContractFulfillment: map[string]client.AttributeFulfillmentValue{
			"SAML_SUBJECT": {
				Source: client.SourceTypeIdKey{
					Type: "NO_MAPPING",
				},
			},
		},
		Id:       utils.Pointer(testTokenProcessorId + "|" + testTokenGeneratorId),
		SourceId: testTokenProcessorId,
		TargetId: testTokenGeneratorId,
	}

	request = request.Body(clientStruct)

	resource, response, err := request.Execute()
	ok, err := common.HandleClientResponse(response, err, "CreateTokenToTokenMapping", resourceType)
	if err != nil {
		t.Fatalf("Failed to execute client function\nResponse Status: %s\nResponse Body: %s\nError: %v", response.Status, response.Body, err)
	}
	if !ok {
		t.Fatalf("Failed to execute client function\nResponse Status: %s\nResponse Body: %s", response.Status, response.Body)
	}

	return testutils_resource.ResourceCreationInfo{
		testutils_resource.ENUM_ID:        *resource.Id,
		testutils_resource.ENUM_SOURCE_ID: resource.SourceId,
		testutils_resource.ENUM_TARGET_ID: resource.TargetId,
	}
}

func deleteTokenProcessorToTokenGeneratorMapping(t *testing.T, clientInfo *connector.ClientInfo, resourceType, id string) {
	t.Helper()

	request := clientInfo.PingFederateApiClient.TokenProcessorToTokenGeneratorMappingsAPI.DeleteTokenToTokenMappingById(clientInfo.PingFederateContext, id)

	response, err := request.Execute()
	ok, err := common.HandleClientResponse(response, err, "DeleteTokenToTokenMappingById", resourceType)
	if err != nil {
		t.Fatalf("Failed to execute client function\nResponse Status: %s\nResponse Body: %s\nError: %v", response.Status, response.Body, err)
	}
	if !ok {
		t.Fatalf("Failed to execute client function\nResponse Status: %s\nResponse Body: %s", response.Status, response.Body)
	}
}
