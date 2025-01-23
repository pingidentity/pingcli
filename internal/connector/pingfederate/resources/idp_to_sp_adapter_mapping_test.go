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

	idpToSpAdapterMappingId, idpToSpAdapterMappingSourceId, idpToSpAdapterMappingTargetId := createIdpToSpAdapterMapping(t, PingFederateClientInfo, resource.ResourceType())
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

func createIdpToSpAdapterMapping(t *testing.T, clientInfo *connector.PingFederateClientInfo, resourceType string) (string, string, string) {
	t.Helper()

	request := clientInfo.ApiClient.IdpToSpAdapterMappingAPI.CreateIdpToSpAdapterMapping(clientInfo.Context)
	result := client.IdpToSpAdapterMapping{}
	result.Id = utils.Pointer("TestIdpToSpAdapterMappingId")
	result.SourceId = "TestIdpToSpAdapterMappingSourceId"
	result.TargetId = "TestIdpToSpAdapterMappingTargetId"
	result.AttributeContractFulfillment = map[string]client.AttributeFulfillmentValue{}
	// for key, attributeContractFulfillmentElement := range model.AttributeContractFulfillment.Elements() {
	// 	attributeContractFulfillmentValue := client.AttributeFulfillmentValue{}
	// 	attributeContractFulfillmentAttrs := attributeContractFulfillmentElement.(types.Object).Attributes()
	// 	attributeContractFulfillmentSourceValue := client.SourceTypeIdKey{}
	// 	attributeContractFulfillmentSourceAttrs := attributeContractFulfillmentAttrs["source"].(types.Object).Attributes()
	// 	attributeContractFulfillmentSourceValue.Id = attributeContractFulfillmentSourceAttrs["id"].(types.String).ValueStringPointer()
	// 	attributeContractFulfillmentSourceValue.Type = attributeContractFulfillmentSourceAttrs["type"].(types.String).ValueString()
	// 	attributeContractFulfillmentValue.Source = attributeContractFulfillmentSourceValue
	// 	attributeContractFulfillmentValue.Value = attributeContractFulfillmentAttrs["value"].(types.String).ValueString()
	// 	result.AttributeContractFulfillment[key] = attributeContractFulfillmentValue
	// }

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
