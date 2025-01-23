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

func Test_PingFederateOauthAccessTokenMapping_Export(t *testing.T) {
	PingFederateClientInfo := testutils.GetPingFederateClientInfo(t)
	resource := resources.OauthAccessTokenMapping(PingFederateClientInfo)

	oauthAccessTokenMappingId, oauthAccessTokenMappingContextType := createOauthAccessTokenMapping(t, PingFederateClientInfo, resource.ResourceType())
	defer deleteOauthAccessTokenMapping(t, PingFederateClientInfo, resource.ResourceType(), oauthAccessTokenMappingId)

	expectedImportBlocks := []connector.ImportBlock{
		{
			ResourceType: resource.ResourceType(),
			ResourceName: fmt.Sprintf("%s_%s", oauthAccessTokenMappingContextType, oauthAccessTokenMappingId),
			ResourceID:   oauthAccessTokenMappingId,
		},
	}

	testutils.ValidateImportBlocks(t, resource, &expectedImportBlocks)
}

func createOauthAccessTokenMapping(t *testing.T, clientInfo *connector.PingFederateClientInfo, resourceType string) (string, string) {
	t.Helper()

	request := clientInfo.ApiClient.OauthAccessTokenMappingsAPI.CreateMapping(clientInfo.Context)
	result := client.AccessTokenMapping{}
	result.Id = utils.Pointer("TestAccessTokenMappingId")
	result.AccessTokenManagerRef = client.ResourceLink{}
	result.AccessTokenManagerRef.Id = ""
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
	result.Context = client.AccessTokenMappingContext{}
	result.Context.ContextRef = client.ResourceLink{}
	result.Context.ContextRef.Id = ""

	request = request.Body(result)

	resource, response, err := request.Execute()
	err = common.HandleClientResponse(response, err, "CreateMapping", resourceType)
	if err != nil {
		t.Fatalf("Failed to create test %s: %v", resourceType, err)
	}

	return *resource.Id, resource.Context.Type
}

func deleteOauthAccessTokenMapping(t *testing.T, clientInfo *connector.PingFederateClientInfo, resourceType, id string) {
	t.Helper()

	request := clientInfo.ApiClient.OauthAccessTokenMappingsAPI.DeleteMapping(clientInfo.Context, id)

	response, err := request.Execute()
	err = common.HandleClientResponse(response, err, "DeleteMapping", resourceType)
	if err != nil {
		t.Errorf("Failed to delete test %s: %v", resourceType, err)
	}
}
