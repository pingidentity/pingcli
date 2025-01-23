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

func Test_PingFederateOauthTokenExchangeTokenGeneratorMapping_Export(t *testing.T) {
	PingFederateClientInfo := testutils.GetPingFederateClientInfo(t)
	resource := resources.OauthTokenExchangeTokenGeneratorMapping(PingFederateClientInfo)

	oauthTokenExchangeTokenGeneratorMappingId, oauthTokenExchangeTokenGeneratorMappingSourceId, oauthTokenExchangeTokenGeneratorMappingTargetId := createOauthTokenExchangeTokenGeneratorMapping(t, PingFederateClientInfo, resource.ResourceType())
	defer deleteOauthTokenExchangeTokenGeneratorMapping(t, PingFederateClientInfo, resource.ResourceType(), oauthTokenExchangeTokenGeneratorMappingId)

	expectedImportBlocks := []connector.ImportBlock{
		{
			ResourceType: resource.ResourceType(),
			ResourceName: fmt.Sprintf("%s_to_%s", oauthTokenExchangeTokenGeneratorMappingSourceId, oauthTokenExchangeTokenGeneratorMappingTargetId),
			ResourceID:   oauthTokenExchangeTokenGeneratorMappingId,
		},
	}

	testutils.ValidateImportBlocks(t, resource, &expectedImportBlocks)
}

func createOauthTokenExchangeTokenGeneratorMapping(t *testing.T, clientInfo *connector.PingFederateClientInfo, resourceType string) (string, string, string) {
	t.Helper()

	request := clientInfo.ApiClient.OauthTokenExchangeTokenGeneratorMappingsAPI.CreateTokenGeneratorMapping(clientInfo.Context)
	result := client.ProcessorPolicyToGeneratorMapping{
		AttributeContractFulfillment: map[string]client.AttributeFulfillmentValue{},
		Id:                           utils.Pointer("TestProcessorPolicyToGeneratorMappingId"),
		SourceId:                     "TestSourceId",
		TargetId:                     "TestTargetId",
	}

	request = request.Body(result)

	resource, response, err := request.Execute()
	err = common.HandleClientResponse(response, err, "CreateTokenGeneratorMapping", resourceType)
	if err != nil {
		t.Fatalf("Failed to create test %s: %v", resourceType, err)
	}

	return *resource.Id, resource.SourceId, resource.TargetId
}

func deleteOauthTokenExchangeTokenGeneratorMapping(t *testing.T, clientInfo *connector.PingFederateClientInfo, resourceType, id string) {
	t.Helper()

	request := clientInfo.ApiClient.OauthTokenExchangeTokenGeneratorMappingsAPI.DeleteTokenGeneratorMappingById(clientInfo.Context, id)

	response, err := request.Execute()
	err = common.HandleClientResponse(response, err, "DeleteTokenGeneratorMappingById", resourceType)
	if err != nil {
		t.Errorf("Failed to delete test %s: %v", resourceType, err)
	}
}
