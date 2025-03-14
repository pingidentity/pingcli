package pingfederate

import (
	"testing"

	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/common"
	"github.com/pingidentity/pingcli/internal/connector/pingfederate/resources"
	"github.com/pingidentity/pingcli/internal/testing/testutils_resource"
	client "github.com/pingidentity/pingfederate-go-client/v1210/configurationapi"
)

func TestableResource_PingFederateOauthAuthenticationPolicyContractMapping(t *testing.T, clientInfo *connector.ClientInfo) testutils_resource.TestableResource {
	t.Helper()

	return testutils_resource.TestableResource{
		ClientInfo: clientInfo,
		CreateFunc: createOauthAuthenticationPolicyContractMapping,
		DeleteFunc: deleteOauthAuthenticationPolicyContractMapping,
		Dependencies: []testutils_resource.TestableResource{
			TestableResource_PingFederateAuthenticationPolicyContract(t, clientInfo),
		},
		ExportableResource: resources.OauthAuthenticationPolicyContractMapping(clientInfo),
	}
}

func createOauthAuthenticationPolicyContractMapping(t *testing.T, clientInfo *connector.ClientInfo, strArgs ...string) testutils_resource.ResourceCreationInfo {
	t.Helper()

	if len(strArgs) != 2 {
		t.Fatalf("Unexpected number of arguments provided to createOauthAuthenticationPolicyContractMapping(): %v", strArgs)
	}
	resourceType := strArgs[0]
	testApcId := strArgs[1]

	request := clientInfo.PingFederateApiClient.OauthAuthenticationPolicyContractMappingsAPI.CreateApcMapping(clientInfo.Context)
	clientStruct := client.ApcToPersistentGrantMapping{
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

	request = request.Body(clientStruct)

	resource, response, err := request.Execute()
	ok, err := common.HandleClientResponse(response, err, "CreateApcMapping", resourceType)
	if err != nil {
		t.Fatalf("Failed to create test %s: %v", resourceType, err)
	}
	if !ok {
		t.Fatalf("Failed to create test %s: non-ok response", resourceType)
	}

	return testutils_resource.ResourceCreationInfo{
		testutils_resource.ENUM_ID: resource.Id,
	}
}

func deleteOauthAuthenticationPolicyContractMapping(t *testing.T, clientInfo *connector.ClientInfo, resourceType, id string) {
	t.Helper()

	request := clientInfo.PingFederateApiClient.OauthAuthenticationPolicyContractMappingsAPI.DeleteApcMapping(clientInfo.Context, id)

	response, err := request.Execute()
	ok, err := common.HandleClientResponse(response, err, "DeleteApcMapping", resourceType)
	if err != nil {
		t.Errorf("Failed to delete test %s: %v", resourceType, err)
	}
	if !ok {
		t.Fatalf("Failed to delete test %s: non-ok response", resourceType)
	}
}
