package resources_test

import (
	"testing"

	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/common"
	"github.com/pingidentity/pingcli/internal/connector/pingfederate/resources"
	"github.com/pingidentity/pingcli/internal/testing/testutils"
	client "github.com/pingidentity/pingfederate-go-client/v1210/configurationapi"
)

func Test_PingFederateOauthClientRegistrationPolicy_Export(t *testing.T) {
	PingFederateClientInfo := testutils.GetPingFederateClientInfo(t)
	resource := resources.OauthClientRegistrationPolicy(PingFederateClientInfo)

	oauthClientRegistrationPolicyId, oauthClientRegistrationPolicyName := createOauthClientRegistrationPolicy(t, PingFederateClientInfo, resource.ResourceType())
	defer deleteOauthClientRegistrationPolicy(t, PingFederateClientInfo, resource.ResourceType(), oauthClientRegistrationPolicyId)

	expectedImportBlocks := []connector.ImportBlock{
		{
			ResourceType: resource.ResourceType(),
			ResourceName: oauthClientRegistrationPolicyName,
			ResourceID:   oauthClientRegistrationPolicyId,
		},
	}

	testutils.ValidateImportBlocks(t, resource, &expectedImportBlocks)
}

func createOauthClientRegistrationPolicy(t *testing.T, clientInfo *connector.PingFederateClientInfo, resourceType string) (string, string) {
	t.Helper()

	request := clientInfo.ApiClient.OauthClientRegistrationPoliciesAPI.CreateDynamicClientRegistrationPolicy(clientInfo.Context)
	result := client.ClientRegistrationPolicy{
		Id:   "TestClientRegistrationPolicyId",
		Name: "TestClientRegistrationPolicyName",
		PluginDescriptorRef: client.ResourceLink{
			Id: "com.pingidentity.pf.client.registration.ResponseTypesConstraintsPlugin",
		},
	}

	request = request.Body(result)

	resource, response, err := request.Execute()
	err = common.HandleClientResponse(response, err, "CreateDynamicClientRegistrationPolicy", resourceType)
	if err != nil {
		t.Fatalf("Failed to create test %s: %v", resourceType, err)
	}

	return resource.Id, resource.Name
}

func deleteOauthClientRegistrationPolicy(t *testing.T, clientInfo *connector.PingFederateClientInfo, resourceType, id string) {
	t.Helper()

	request := clientInfo.ApiClient.OauthClientRegistrationPoliciesAPI.DeleteDynamicClientRegistrationPolicy(clientInfo.Context, id)

	response, err := request.Execute()
	err = common.HandleClientResponse(response, err, "DeleteDynamicClientRegistrationPolicy", resourceType)
	if err != nil {
		t.Errorf("Failed to delete test %s: %v", resourceType, err)
	}
}
