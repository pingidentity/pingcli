package resources_test

import (
	"testing"

	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/common"
	"github.com/pingidentity/pingcli/internal/connector/pingfederate/resources"
	"github.com/pingidentity/pingcli/internal/testing/testutils"
	client "github.com/pingidentity/pingfederate-go-client/v1210/configurationapi"
)

func Test_PingFederateOauthCibaServerPolicyRequestPolicy_Export(t *testing.T) {
	PingFederateClientInfo := testutils.GetPingFederateClientInfo(t)
	resource := resources.OauthCibaServerPolicyRequestPolicy(PingFederateClientInfo)

	oauthCibaServerPolicyRequestPolicyId, oauthCibaServerPolicyRequestPolicyName := createOauthCibaServerPolicyRequestPolicy(t, PingFederateClientInfo, resource.ResourceType())
	defer deleteOauthCibaServerPolicyRequestPolicy(t, PingFederateClientInfo, resource.ResourceType(), oauthCibaServerPolicyRequestPolicyId)

	expectedImportBlocks := []connector.ImportBlock{
		{
			ResourceType: resource.ResourceType(),
			ResourceName: oauthCibaServerPolicyRequestPolicyName,
			ResourceID:   oauthCibaServerPolicyRequestPolicyId,
		},
	}

	testutils.ValidateImportBlocks(t, resource, &expectedImportBlocks)
}

func createOauthCibaServerPolicyRequestPolicy(t *testing.T, clientInfo *connector.PingFederateClientInfo, resourceType string) (string, string) {
	t.Helper()

	request := clientInfo.ApiClient.OauthCibaServerPolicyAPI.CreateCibaServerPolicy(clientInfo.Context)
	result := client.RequestPolicy{
		AuthenticatorRef: client.ResourceLink{
			Id: "",
		},
		Id:                   "TestRequestPolicyId",
		IdentityHintContract: client.IdentityHintContract{},
		Name:                 "TestRequestPolicyName",
	}

	request = request.Body(result)

	resource, response, err := request.Execute()
	err = common.HandleClientResponse(response, err, "CreateCibaServerPolicy", resourceType)
	if err != nil {
		t.Fatalf("Failed to create test %s: %v", resourceType, err)
	}

	return resource.Id, resource.Name
}

func deleteOauthCibaServerPolicyRequestPolicy(t *testing.T, clientInfo *connector.PingFederateClientInfo, resourceType, id string) {
	t.Helper()

	request := clientInfo.ApiClient.OauthCibaServerPolicyAPI.DeleteCibaServerPolicy(clientInfo.Context, id)

	response, err := request.Execute()
	err = common.HandleClientResponse(response, err, "DeleteCibaServerPolicy", resourceType)
	if err != nil {
		t.Errorf("Failed to delete test %s: %v", resourceType, err)
	}
}
