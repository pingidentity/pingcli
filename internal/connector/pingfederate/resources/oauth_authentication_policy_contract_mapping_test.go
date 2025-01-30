package resources_test

import (
	"fmt"
	"testing"

	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/common"
	"github.com/pingidentity/pingcli/internal/connector/pingfederate/resources"
	"github.com/pingidentity/pingcli/internal/testing/testutils"
	"github.com/pingidentity/pingcli/internal/testing/testutils_resource"
	client "github.com/pingidentity/pingfederate-go-client/v1210/configurationapi"
)

func TestableResource_PingFederateOauthAuthenticationPolicyContractMapping(t *testing.T) *testutils_resource.TestableResource {
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
			},
			CreateFunc: createOauthAuthenticationPolicyContractMapping,
			DeleteFunc: deleteOauthAuthenticationPolicyContractMapping,
		},
	}
}

func Test_PingFederateOauthAuthenticationPolicyContractMapping(t *testing.T) {
	tr := TestableResource_PingFederateOauthAuthenticationPolicyContractMapping(t)

	creationInfo := tr.CreateResource(t, tr.TestResource)
	defer tr.DeleteResource(t, tr.TestResource)

	expectedImportBlocks := []connector.ImportBlock{
		{
			ResourceType: tr.ExportableResource.ResourceType(),
			ResourceName: fmt.Sprintf("%s_mapping", creationInfo[testutils_resource.ENUM_ID]),
			ResourceID:   creationInfo[testutils_resource.ENUM_ID],
		},
	}

	testutils.ValidateImportBlocks(t, tr.ExportableResource, &expectedImportBlocks)
}

func createOauthAuthenticationPolicyContractMapping(t *testing.T, clientInfo *connector.ClientInfo, strArgs ...string) testutils_resource.ResourceCreationInfo {
	t.Helper()

	if len(strArgs) != 2 {
		t.Fatalf("Unexpected number of arguments provided to createOauthAuthenticationPolicyContractMapping(): %v", strArgs)
	}
	resourceType := strArgs[0]
	testApcId := strArgs[1]

	request := clientInfo.PingFederateApiClient.OauthAuthenticationPolicyContractMappingsAPI.CreateApcMapping(clientInfo.Context)
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
	err = common.HandleClientResponse(response, err, "CreateApplication", resourceType)
	if err != nil {
		t.Fatalf("Failed to create test %s: %v", resourceType, err)
	}

	return testutils_resource.ResourceCreationInfo{
		testutils_resource.ENUM_ID: resource.Id,
	}
}

func deleteOauthAuthenticationPolicyContractMapping(t *testing.T, clientInfo *connector.ClientInfo, resourceType, id string) {
	t.Helper()

	request := clientInfo.PingFederateApiClient.OauthAuthenticationPolicyContractMappingsAPI.DeleteApcMapping(clientInfo.Context, id)

	response, err := request.Execute()
	err = common.HandleClientResponse(response, err, "DeleteApcMapping", resourceType)
	if err != nil {
		t.Errorf("Failed to delete test %s: %v", resourceType, err)
	}
}
