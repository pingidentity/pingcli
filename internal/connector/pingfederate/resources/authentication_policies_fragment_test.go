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

func Test_PingFederateAuthenticationPoliciesFragment_Export(t *testing.T) {
	PingFederateClientInfo := testutils.GetPingFederateClientInfo(t)
	resource := resources.AuthenticationPoliciesFragment(PingFederateClientInfo)

	authenticationPoliciesFragmentId, authenticationPoliciesFragmentName := createAuthenticationPoliciesFragment(t, PingFederateClientInfo, resource.ResourceType())
	defer deleteAuthenticationPoliciesFragment(t, PingFederateClientInfo, resource.ResourceType(), authenticationPoliciesFragmentId)

	expectedImportBlocks := []connector.ImportBlock{
		{
			ResourceType: resource.ResourceType(),
			ResourceName: authenticationPoliciesFragmentName,
			ResourceID:   authenticationPoliciesFragmentId,
		},
	}

	testutils.ValidateImportBlocks(t, resource, &expectedImportBlocks)
}

func createAuthenticationPoliciesFragment(t *testing.T, clientInfo *connector.PingFederateClientInfo, resourceType string) (string, string) {
	t.Helper()

	request := clientInfo.ApiClient.AuthenticationPoliciesAPI.CreateFragment(clientInfo.Context)
	result := client.AuthenticationPolicyFragment{}
	result.Id = utils.Pointer("TestAuthenticationPolicyFragmentId")
	result.Name = utils.Pointer("TestAuthenticationPolicyFragmentName")

	request = request.Body(result)

	resource, response, err := request.Execute()
	err = common.HandleClientResponse(response, err, "CreateFragment", resourceType)
	if err != nil {
		t.Fatalf("Failed to create test %s: %v", resourceType, err)
	}

	return *resource.Id, *resource.Name
}

func deleteAuthenticationPoliciesFragment(t *testing.T, clientInfo *connector.PingFederateClientInfo, resourceType, id string) {
	t.Helper()

	request := clientInfo.ApiClient.AuthenticationPoliciesAPI.DeleteFragment(clientInfo.Context, id)

	response, err := request.Execute()
	err = common.HandleClientResponse(response, err, "DeleteFragment", resourceType)
	if err != nil {
		t.Errorf("Failed to delete test %s: %v", resourceType, err)
	}
}
