package resources_test

import (
	"testing"

	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/common"
	"github.com/pingidentity/pingcli/internal/connector/pingfederate/resources"
	"github.com/pingidentity/pingcli/internal/testing/testutils"
	client "github.com/pingidentity/pingfederate-go-client/v1210/configurationapi"
)

func Test_PingFederateAuthenticationSelector_Export(t *testing.T) {
	PingFederateClientInfo := testutils.GetPingFederateClientInfo(t)
	resource := resources.AuthenticationSelector(PingFederateClientInfo)

	authenticationSelectorId, authenticationSelectorName := createAuthenticationSelector(t, PingFederateClientInfo, resource.ResourceType())
	defer deleteAuthenticationSelector(t, PingFederateClientInfo, resource.ResourceType(), authenticationSelectorId)

	expectedImportBlocks := []connector.ImportBlock{
		{
			ResourceType: resource.ResourceType(),
			ResourceName: authenticationSelectorName,
			ResourceID:   authenticationSelectorId,
		},
	}

	testutils.ValidateImportBlocks(t, resource, &expectedImportBlocks)
}

func createAuthenticationSelector(t *testing.T, clientInfo *connector.PingFederateClientInfo, resourceType string) (string, string) {
	t.Helper()

	request := clientInfo.ApiClient.AuthenticationSelectorsAPI.CreateAuthenticationSelector(clientInfo.Context)
	result := client.AuthenticationSelector{}
	result.Id = "TestAuthenticationSelectorId"
	result.Name = "TestAuthenticationSelectorName"
	result.Configuration = client.PluginConfiguration{}
	result.PluginDescriptorRef = client.ResourceLink{
		Id: "",
	}

	request = request.Body(result)

	resource, response, err := request.Execute()
	err = common.HandleClientResponse(response, err, "CreateAuthenticationSelector", resourceType)
	if err != nil {
		t.Fatalf("Failed to create test %s: %v", resourceType, err)
	}

	return resource.Id, resource.Name
}

func deleteAuthenticationSelector(t *testing.T, clientInfo *connector.PingFederateClientInfo, resourceType, id string) {
	t.Helper()

	request := clientInfo.ApiClient.AuthenticationSelectorsAPI.DeleteAuthenticationSelector(clientInfo.Context, id)

	response, err := request.Execute()
	err = common.HandleClientResponse(response, err, "DeleteAuthenticationSelector", resourceType)
	if err != nil {
		t.Errorf("Failed to delete test %s: %v", resourceType, err)
	}
}
