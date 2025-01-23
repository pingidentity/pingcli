package resources_test

import (
	"fmt"
	"testing"

	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/common"
	"github.com/pingidentity/pingcli/internal/connector/pingfederate/resources"
	"github.com/pingidentity/pingcli/internal/testing/testutils"
	client "github.com/pingidentity/pingfederate-go-client/v1210/configurationapi"
)

func Test_PingFederateOauthAuthenticationPolicyContractMapping_Export(t *testing.T) {
	PingFederateClientInfo := testutils.GetPingFederateClientInfo(t)
	resource := resources.OauthAuthenticationPolicyContractMapping(PingFederateClientInfo)

	testApcId, _ := createAuthenticationPolicyContract(t, PingFederateClientInfo, resource.ResourceType())
	defer deleteAuthenticationPolicyContract(t, PingFederateClientInfo, resource.ResourceType(), testApcId)

	oauthAuthenticationPolicyContractMappingId := createOauthAuthenticationPolicyContractMapping(t, PingFederateClientInfo, resource.ResourceType(), testApcId)
	defer deleteOauthAuthenticationPolicyContractMapping(t, PingFederateClientInfo, resource.ResourceType(), oauthAuthenticationPolicyContractMappingId)

	expectedImportBlocks := []connector.ImportBlock{
		{
			ResourceType: resource.ResourceType(),
			ResourceName: fmt.Sprintf("%s_mapping", oauthAuthenticationPolicyContractMappingId),
			ResourceID:   oauthAuthenticationPolicyContractMappingId,
		},
	}

	testutils.ValidateImportBlocks(t, resource, &expectedImportBlocks)
}

func createOauthAuthenticationPolicyContractMapping(t *testing.T, clientInfo *connector.PingFederateClientInfo, resourceType, testApcId string) string {
	t.Helper()

	request := clientInfo.ApiClient.OauthAuthenticationPolicyContractMappingsAPI.CreateApcMapping(clientInfo.Context)
	result := client.ApcToPersistentGrantMapping{
		AttributeContractFulfillment: map[string]client.AttributeFulfillmentValue{
			"USER_NAME": {
				Source: client.SourceTypeIdKey{
					Type: "NO_MAPPING",
				},
			},
			"USER_KEY": {
				Source: client.SourceTypeIdKey{
					Type: "NO_MAPPING",
				},
			},
		},
		AuthenticationPolicyContractRef: client.ResourceLink{
			Id: testApcId,
		},
		Id: "testApcToPersistentGrantMappingId",
	}

	request = request.Body(result)

	resource, response, err := request.Execute()
	err = common.HandleClientResponse(response, err, "CreateApcMapping", resourceType)
	if err != nil {
		t.Fatalf("Failed to create test %s: %v", resourceType, err)
	}

	return resource.Id
}

func deleteOauthAuthenticationPolicyContractMapping(t *testing.T, clientInfo *connector.PingFederateClientInfo, resourceType, id string) {
	t.Helper()

	request := clientInfo.ApiClient.OauthAuthenticationPolicyContractMappingsAPI.DeleteApcMapping(clientInfo.Context, id)

	response, err := request.Execute()
	err = common.HandleClientResponse(response, err, "DeleteApcMapping", resourceType)
	if err != nil {
		t.Errorf("Failed to delete test %s: %v", resourceType, err)
	}
}
