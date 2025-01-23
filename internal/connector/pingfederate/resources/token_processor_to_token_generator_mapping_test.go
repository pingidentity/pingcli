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

func Test_PingFederateTokenProcessorToTokenGeneratorMapping_Export(t *testing.T) {
	PingFederateClientInfo := testutils.GetPingFederateClientInfo(t)
	resource := resources.TokenProcessorToTokenGeneratorMapping(PingFederateClientInfo)

	tokenProcessorToTokenGeneratorMappingId, tokenProcessorToTokenGeneratorMappingSourceId, tokenProcessorToTokenGeneratorMappingTargetId := createTokenProcessorToTokenGeneratorMapping(t, PingFederateClientInfo, resource.ResourceType())
	defer deleteTokenProcessorToTokenGeneratorMapping(t, PingFederateClientInfo, resource.ResourceType(), tokenProcessorToTokenGeneratorMappingId)

	expectedImportBlocks := []connector.ImportBlock{
		{
			ResourceType: resource.ResourceType(),
			ResourceName: fmt.Sprintf("%s_to_%s", tokenProcessorToTokenGeneratorMappingSourceId, tokenProcessorToTokenGeneratorMappingTargetId),
			ResourceID:   tokenProcessorToTokenGeneratorMappingId,
		},
	}

	testutils.ValidateImportBlocks(t, resource, &expectedImportBlocks)
}

func createTokenProcessorToTokenGeneratorMapping(t *testing.T, clientInfo *connector.PingFederateClientInfo, resourceType string) (string, string, string) {
	t.Helper()

	request := clientInfo.ApiClient.TokenProcessorToTokenGeneratorMappingsAPI.CreateTokenToTokenMapping(clientInfo.Context)
	result := client.TokenToTokenMapping{
		AttributeContractFulfillment: map[string]client.AttributeFulfillmentValue{},
		Id:                           utils.Pointer("TestTokenToTokenMappingId"),
		SourceId:                     "TestTokenToTokenMappingSourceId",
		TargetId:                     "TestTokenToTokenMappingTargetId",
	}

	request = request.Body(result)

	resource, response, err := request.Execute()
	err = common.HandleClientResponse(response, err, "CreateTokenToTokenMapping", resourceType)
	if err != nil {
		t.Fatalf("Failed to create test %s: %v", resourceType, err)
	}

	return *resource.Id, resource.SourceId, resource.TargetId
}

func deleteTokenProcessorToTokenGeneratorMapping(t *testing.T, clientInfo *connector.PingFederateClientInfo, resourceType, id string) {
	t.Helper()

	request := clientInfo.ApiClient.TokenProcessorToTokenGeneratorMappingsAPI.DeleteTokenToTokenMappingById(clientInfo.Context, id)

	response, err := request.Execute()
	err = common.HandleClientResponse(response, err, "DeleteTokenToTokenMappingById", resourceType)
	if err != nil {
		t.Errorf("Failed to delete test %s: %v", resourceType, err)
	}
}
