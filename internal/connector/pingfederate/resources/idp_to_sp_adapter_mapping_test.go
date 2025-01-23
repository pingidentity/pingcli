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

func Test_PingFederateIdpToSpAdapterMapping_Export(t *testing.T) {
	PingFederateClientInfo := testutils.GetPingFederateClientInfo(t)
	resource := resources.IdpToSpAdapterMapping(PingFederateClientInfo)

	testPCVId, _ := createPasswordCredentialValidator(t, PingFederateClientInfo, resource.ResourceType())
	defer deletePasswordCredentialValidator(t, PingFederateClientInfo, resource.ResourceType(), testPCVId)

	testIdpAdapterId, _ := createIdpAdapter(t, PingFederateClientInfo, resource.ResourceType(), testPCVId)
	defer deleteIdpAdapter(t, PingFederateClientInfo, resource.ResourceType(), testIdpAdapterId)

	testSpAdapterId, _ := createSpAdapter(t, PingFederateClientInfo, resource.ResourceType())
	defer deleteSpAdapter(t, PingFederateClientInfo, resource.ResourceType(), testSpAdapterId)

	idpToSpAdapterMappingId, idpToSpAdapterMappingSourceId, idpToSpAdapterMappingTargetId := createIdpToSpAdapterMapping(t, PingFederateClientInfo, resource.ResourceType(), testIdpAdapterId, testSpAdapterId)
	defer deleteIdpToSpAdapterMapping(t, PingFederateClientInfo, resource.ResourceType(), idpToSpAdapterMappingId)

	expectedImportBlocks := []connector.ImportBlock{
		{
			ResourceType: resource.ResourceType(),
			ResourceName: fmt.Sprintf("%s_to_%s", idpToSpAdapterMappingSourceId, idpToSpAdapterMappingTargetId),
			ResourceID:   idpToSpAdapterMappingId,
		},
	}

	testutils.ValidateImportBlocks(t, resource, &expectedImportBlocks)
}

func createIdpToSpAdapterMapping(t *testing.T, clientInfo *connector.PingFederateClientInfo, resourceType, testIdpAdapterId, testSpAdapterId string) (string, string, string) {
	t.Helper()

	request := clientInfo.ApiClient.IdpToSpAdapterMappingAPI.CreateIdpToSpAdapterMapping(clientInfo.Context)
	result := client.IdpToSpAdapterMapping{
		AttributeContractFulfillment: map[string]client.AttributeFulfillmentValue{
			"subject": {
				Source: client.SourceTypeIdKey{
					Type: "NO_MAPPING",
				},
			},
		},
		Id:       utils.Pointer(testIdpAdapterId + "|" + testSpAdapterId),
		SourceId: testIdpAdapterId,
		TargetId: testSpAdapterId,
	}

	request = request.Body(result)

	resource, response, err := request.Execute()
	err = common.HandleClientResponse(response, err, "CreateIdpToSpAdapterMapping", resourceType)
	if err != nil {
		t.Fatalf("Failed to create test %s: %v", resourceType, err)
	}

	return *resource.Id, resource.SourceId, resource.TargetId
}

func deleteIdpToSpAdapterMapping(t *testing.T, clientInfo *connector.PingFederateClientInfo, resourceType, id string) {
	t.Helper()

	request := clientInfo.ApiClient.IdpToSpAdapterMappingAPI.DeleteIdpToSpAdapterMappingsById(clientInfo.Context, id)

	response, err := request.Execute()
	err = common.HandleClientResponse(response, err, "DeleteIdpToSpAdapterMappingsById", resourceType)
	if err != nil {
		t.Errorf("Failed to delete test %s: %v", resourceType, err)
	}
}
