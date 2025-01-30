package resources_test

import (
	"testing"

	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/common"
	"github.com/pingidentity/pingcli/internal/connector/pingfederate/resources"
	"github.com/pingidentity/pingcli/internal/testing/testutils"
	"github.com/pingidentity/pingcli/internal/testing/testutils_resource"
	"github.com/pingidentity/pingcli/internal/utils"
	client "github.com/pingidentity/pingfederate-go-client/v1210/configurationapi"
)

func TestableResource_PingFederateAuthenticationPolicyContract(t *testing.T) *testutils_resource.TestableResource {
	t.Helper()

	pingfederateClientInfo := testutils.GetPingFederateClientInfo(t)
	return &testutils_resource.TestableResource{
		ClientInfo:         pingfederateClientInfo,
		ExportableResource: resources.AuthenticationApiApplication(pingfederateClientInfo),
		TestResource: testutils_resource.TestResource{
			Dependencies: nil,
			CreateFunc:   createAuthenticationPolicyContract,
			DeleteFunc:   deleteAuthenticationPolicyContract,
		},
	}
}

func Test_PingFederateAuthenticationPolicyContract(t *testing.T) {
	tr := TestableResource_PingFederateAuthenticationPolicyContract(t)

	creationInfo := tr.CreateResource(t, tr.TestResource)
	defer tr.DeleteResource(t, tr.TestResource)

	expectedImportBlocks := []connector.ImportBlock{
		{
			ResourceType: tr.ExportableResource.ResourceType(),
			ResourceName: creationInfo[testutils_resource.ENUM_NAME],
			ResourceID:   creationInfo[testutils_resource.ENUM_ID],
		},
	}

	testutils.ValidateImportBlocks(t, tr.ExportableResource, &expectedImportBlocks)
}

func createAuthenticationPolicyContract(t *testing.T, clientInfo *connector.ClientInfo, strArgs ...string) testutils_resource.ResourceCreationInfo {
	t.Helper()

	if len(strArgs) != 1 {
		t.Fatalf("Unexpected number of arguments provided to createAuthenticationPolicyContract(): %v", strArgs)
	}
	resourceType := strArgs[0]

	request := clientInfo.PingFederateApiClient.AuthenticationPolicyContractsAPI.CreateAuthenticationPolicyContract(clientInfo.Context)
	result := client.AuthenticationPolicyContract{
		CoreAttributes: []client.AuthenticationPolicyContractAttribute{
			{
				Name: "subject",
			},
		},
		Id:   utils.Pointer("TestApcId"),
		Name: utils.Pointer("TestApcName"),
	}

	request = request.Body(result)

	resource, response, err := request.Execute()
	err = common.HandleClientResponse(response, err, "CreateApplication", resourceType)
	if err != nil {
		t.Fatalf("Failed to create test %s: %v", resourceType, err)
	}

	return testutils_resource.ResourceCreationInfo{
		testutils_resource.ENUM_ID:   *resource.Id,
		testutils_resource.ENUM_NAME: *resource.Name,
	}
}

func deleteAuthenticationPolicyContract(t *testing.T, clientInfo *connector.ClientInfo, resourceType, id string) {
	t.Helper()

	request := clientInfo.PingFederateApiClient.AuthenticationPolicyContractsAPI.DeleteAuthenticationPolicyContract(clientInfo.Context, id)

	response, err := request.Execute()
	err = common.HandleClientResponse(response, err, "DeleteAuthenticationPolicyContract", resourceType)
	if err != nil {
		t.Errorf("Failed to delete test %s: %v", resourceType, err)
	}
}
