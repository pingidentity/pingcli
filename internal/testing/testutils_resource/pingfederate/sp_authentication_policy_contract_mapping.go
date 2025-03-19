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

func TestableResource_PingFederateSpAuthenticationPolicyContractMapping(t *testing.T, clientInfo *connector.ClientInfo) *testutils_resource.TestableResource {
	t.Helper()

	return &testutils_resource.TestableResource{
		ClientInfo: clientInfo,
		CreateFunc: createSpAuthenticationPolicyContractMapping,
		DeleteFunc: deleteSpAuthenticationPolicyContractMapping,
		Dependencies: []*testutils_resource.TestableResource{
			TestableResource_PingFederateAuthenticationPolicyContract(t, clientInfo),
			TestableResource_PingFederateSpAdapter(t, clientInfo),
		},
		ExportableResource: resources.SpAuthenticationPolicyContractMapping(clientInfo),
	}
}

func createSpAuthenticationPolicyContractMapping(t *testing.T, clientInfo *connector.ClientInfo, strArgs ...string) testutils_resource.ResourceCreationInfo {
	t.Helper()

	if len(strArgs) != 3 {
		t.Fatalf("Unexpected number of arguments provided to createSpAuthenticationPolicyContractMapping(): %v", strArgs)
	}
	resourceType := strArgs[0]
	testAPCId := strArgs[1]
	testSPAdapterId := strArgs[2]

	request := clientInfo.PingFederateApiClient.SpAuthenticationPolicyContractMappingsAPI.CreateApcToSpAdapterMapping(clientInfo.PingFederateContext)
	clientStruct := client.ApcToSpAdapterMapping{
		AttributeContractFulfillment: map[string]client.AttributeFulfillmentValue{
			"subject": {
				Source: client.SourceTypeIdKey{
					Type: "NO_MAPPING",
				},
			},
		},
		Id:       utils.Pointer(testAPCId + "|" + testSPAdapterId),
		SourceId: testAPCId,
		TargetId: testSPAdapterId,
	}

	request = request.Body(clientStruct)

	resource, response, err := request.Execute()
	ok, err := common.HandleClientResponse(response, err, "CreateApcToSpAdapterMapping", resourceType)
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

func deleteSpAuthenticationPolicyContractMapping(t *testing.T, clientInfo *connector.ClientInfo, resourceType, id string) {
	t.Helper()

	request := clientInfo.PingFederateApiClient.SpAuthenticationPolicyContractMappingsAPI.DeleteApcToSpAdapterMappingById(clientInfo.PingFederateContext, id)

	response, err := request.Execute()
	ok, err := common.HandleClientResponse(response, err, "DeleteApcToSpAdapterMappingById", resourceType)
	if err != nil {
		t.Fatalf("Failed to execute client function\nResponse Status: %s\nResponse Body: %s\nError: %v", response.Status, response.Body, err)
	}
	if !ok {
		t.Fatalf("Failed to execute client function\nResponse Status: %s\nResponse Body: %s", response.Status, response.Body)
	}
}
