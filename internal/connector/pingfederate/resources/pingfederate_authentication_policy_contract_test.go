package resources_test

import (
	"testing"

	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/common"
	"github.com/pingidentity/pingcli/internal/connector/pingfederate/resources"
	"github.com/pingidentity/pingcli/internal/testing/testutils"
	"github.com/pingidentity/pingcli/internal/utils"
	client "github.com/pingidentity/pingfederate-go-client/v1210/configurationapi"
)

func Test_PingFederateAuthenticationPolicyContract_Export(t *testing.T) {
	PingFederateClientInfo := testutils.GetPingFederateClientInfo(t)
	resource := resources.AuthenticationPolicyContract(PingFederateClientInfo)

	authenticationPolicyContractId, authenticationPolicyContractName := createAuthenticationPolicyContract(t, PingFederateClientInfo, resource.ResourceType())
	defer deleteAuthenticationPolicyContract(t, PingFederateClientInfo, resource.ResourceType(), authenticationPolicyContractId)

	expectedImportBlocks := []connector.ImportBlock{
		{
			ResourceType: resource.ResourceType(),
			ResourceName: authenticationPolicyContractName,
			ResourceID:   authenticationPolicyContractId,
		},
	}

	testutils.ValidateImportBlocks(t, resource, &expectedImportBlocks)
}

func createAuthenticationPolicyContract(t *testing.T, clientInfo *connector.PingFederateClientInfo, resourceType string) (string, string) {
	t.Helper()

	request := clientInfo.ApiClient.AuthenticationPolicyContractsAPI.CreateAuthenticationPolicyContract(clientInfo.Context)
	result := client.AuthenticationPolicyContract{}
	result.Id = utils.Pointer("TestAPCId")
	result.Name = utils.Pointer("TestAPCName")
	result.CoreAttributes = []client.AuthenticationPolicyContractAttribute{
		{
			Name: "subject",
		},
	}

	request = request.Body(result)

	resource, response, err := request.Execute()
	err = common.HandleClientResponse(response, err, "CreateAuthenticationPolicyContract", resourceType)
	if err != nil {
		t.Fatalf("Failed to create test %s: %v", resourceType, err)
	}

	return *resource.Id, *resource.Name
}

func deleteAuthenticationPolicyContract(t *testing.T, clientInfo *connector.PingFederateClientInfo, resourceType, id string) {
	t.Helper()

	request := clientInfo.ApiClient.AuthenticationPolicyContractsAPI.DeleteAuthenticationPolicyContract(clientInfo.Context, id)

	response, err := request.Execute()
	err = common.HandleClientResponse(response, err, "DeleteAuthenticationPolicyContract", resourceType)
	if err != nil {
		t.Errorf("Failed to delete test %s: %v", resourceType, err)
	}
}
