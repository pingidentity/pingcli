package resources_test

import (
	"fmt"
	"testing"

	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/common"
	"github.com/pingidentity/pingcli/internal/connector/pingfederate/resources"
	"github.com/pingidentity/pingcli/internal/testing/testutils"
	"github.com/pingidentity/pingcli/internal/utils"
	client "github.com/pingidentity/pingfederate-go-client/v1210/configurationapi"
)

func Test_PingFederateSpAuthenticationPolicyContractMapping_Export(t *testing.T) {
	PingFederateClientInfo := testutils.GetPingFederateClientInfo(t)
	resource := resources.SpAuthenticationPolicyContractMapping(PingFederateClientInfo)

	testAPCId, _ := createAuthenticationPolicyContract(t, PingFederateClientInfo, resource.ResourceType())
	defer deleteAuthenticationPolicyContract(t, PingFederateClientInfo, resource.ResourceType(), testAPCId)

	testSPAdapterId, _ := createSpAdapter(t, PingFederateClientInfo, resource.ResourceType())
	defer deleteSpAdapter(t, PingFederateClientInfo, resource.ResourceType(), testSPAdapterId)

	spAuthenticationPolicyContractMappingId, spAuthenticationPolicyContractMappingSourceId, spAuthenticationPolicyContractMappingTargetId := createSpAuthenticationPolicyContractMapping(t, PingFederateClientInfo, resource.ResourceType(), testAPCId, testSPAdapterId)
	defer deleteSpAuthenticationPolicyContractMapping(t, PingFederateClientInfo, resource.ResourceType(), spAuthenticationPolicyContractMappingId)

	expectedImportBlocks := []connector.ImportBlock{
		{
			ResourceType: resource.ResourceType(),
			ResourceName: fmt.Sprintf("%s_to_%s", spAuthenticationPolicyContractMappingSourceId, spAuthenticationPolicyContractMappingTargetId),
			ResourceID:   spAuthenticationPolicyContractMappingId,
		},
	}

	testutils.ValidateImportBlocks(t, resource, &expectedImportBlocks)
}

func createSpAuthenticationPolicyContractMapping(t *testing.T, clientInfo *connector.PingFederateClientInfo, resourceType, testAPCId, testSPAdapterId string) (string, string, string) {
	t.Helper()

	request := clientInfo.ApiClient.SpAuthenticationPolicyContractMappingsAPI.CreateApcToSpAdapterMapping(clientInfo.Context)
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
	err = common.HandleClientResponse(response, err, "CreateApcToSpAdapterMapping", resourceType)
	if err != nil {
		t.Fatalf("Failed to create test %s: %v", resourceType, err)
	}

	return *resource.Id, resource.SourceId, resource.TargetId
}

func deleteSpAuthenticationPolicyContractMapping(t *testing.T, clientInfo *connector.PingFederateClientInfo, resourceType, id string) {
	t.Helper()

	request := clientInfo.ApiClient.SpAuthenticationPolicyContractMappingsAPI.DeleteApcToSpAdapterMappingById(clientInfo.Context, id)

	response, err := request.Execute()
	err = common.HandleClientResponse(response, err, "DeleteApcToSpAdapterMappingById", resourceType)
	if err != nil {
		t.Errorf("Failed to delete test %s: %v", resourceType, err)
	}
}
