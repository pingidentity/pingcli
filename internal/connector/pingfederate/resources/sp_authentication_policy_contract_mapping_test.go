package resources_test

import (
	"fmt"
	"testing"

	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/common"
	"github.com/pingidentity/pingcli/internal/connector/pingfederate/resources"
	"github.com/pingidentity/pingcli/internal/testing/testutils"
	"github.com/pingidentity/pingcli/internal/testing/testutils_resource"
	"github.com/pingidentity/pingcli/internal/utils"
	client "github.com/pingidentity/pingfederate-go-client/v1210/configurationapi"
)

func TestableResource_PingFederateSpAuthenticationPolicyContractMapping(t *testing.T) *testutils_resource.TestableResource {
	t.Helper()

	pingfederateClientInfo := testutils.GetPingFederateClientInfo(t)
	return &testutils_resource.TestableResource{
		ClientInfo:         pingfederateClientInfo,
		ExportableResource: resources.AuthenticationApiApplication(pingfederateClientInfo),
		TestResource: testutils_resource.TestResource{
			Dependencies: []testutils_resource.TestResource{
				{
					Dependencies: nil,
					CreateFunc:   createAuthenticationPolicyContract,
					DeleteFunc:   deleteAuthenticationPolicyContract,
				},
				{
					Dependencies: nil,
					CreateFunc:   createSpAdapter,
					DeleteFunc:   deleteSpAdapter,
				},
			},
			CreateFunc: createSpAuthenticationPolicyContractMapping,
			DeleteFunc: deleteSpAuthenticationPolicyContractMapping,
		},
	}
}

func Test_PingFederateSpAuthenticationPolicyContractMapping(t *testing.T) {
	tr := TestableResource_PingFederateSpAuthenticationPolicyContractMapping(t)

	creationInfo := tr.CreateResource(t, tr.TestResource)
	defer tr.DeleteResource(t, tr.TestResource)

	expectedImportBlocks := []connector.ImportBlock{
		{
			ResourceType: tr.ExportableResource.ResourceType(),
			ResourceName: fmt.Sprintf("%s_to_%s", creationInfo[testutils_resource.ENUM_SOURCE_ID], creationInfo[testutils_resource.ENUM_TARGET_ID]),
			ResourceID:   creationInfo[testutils_resource.ENUM_ID],
		},
	}

	testutils.ValidateImportBlocks(t, tr.ExportableResource, &expectedImportBlocks)
}

func createSpAuthenticationPolicyContractMapping(t *testing.T, clientInfo *connector.ClientInfo, strArgs ...string) testutils_resource.ResourceCreationInfo {
	t.Helper()

	if len(strArgs) != 3 {
		t.Fatalf("Unexpected number of arguments provided to createSpAuthenticationPolicyContractMapping(): %v", strArgs)
	}
	resourceType := strArgs[0]
	testAPCId := strArgs[1]
	testSPAdapterId := strArgs[2]

	request := clientInfo.PingFederateApiClient.SpAuthenticationPolicyContractMappingsAPI.CreateApcToSpAdapterMapping(clientInfo.Context)
	result := client.ApcToSpAdapterMapping{
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

	request = request.Body(result)

	resource, response, err := request.Execute()
	err = common.HandleClientResponse(response, err, "CreateApplication", resourceType)
	if err != nil {
		t.Fatalf("Failed to create test %s: %v", resourceType, err)
	}

	return testutils_resource.ResourceCreationInfo{
		testutils_resource.ENUM_ID:        *resource.Id,
		testutils_resource.ENUM_SOURCE_ID: resource.SourceId,
		testutils_resource.ENUM_TARGET_ID: resource.TargetId,
	}
}

func deleteSpAuthenticationPolicyContractMapping(t *testing.T, clientInfo *connector.ClientInfo, resourceType, id string) {
	t.Helper()

	request := clientInfo.PingFederateApiClient.SpAuthenticationPolicyContractMappingsAPI.DeleteApcToSpAdapterMappingById(clientInfo.Context, id)

	response, err := request.Execute()
	err = common.HandleClientResponse(response, err, "DeleteApcToSpAdapterMappingById", resourceType)
	if err != nil {
		t.Errorf("Failed to delete test %s: %v", resourceType, err)
	}
}
