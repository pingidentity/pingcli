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

func Test_PingFederateSpIdpConnection_Export(t *testing.T) {
	PingFederateClientInfo := testutils.GetPingFederateClientInfo(t)
	resource := resources.SpIdpConnection(PingFederateClientInfo)

	spIdpConnectionId, spIdpConnectionName := createSpIdpConnection(t, PingFederateClientInfo, resource.ResourceType())
	defer deleteSpIdpConnection(t, PingFederateClientInfo, resource.ResourceType(), spIdpConnectionId)

	expectedImportBlocks := []connector.ImportBlock{
		{
			ResourceType: resource.ResourceType(),
			ResourceName: spIdpConnectionName,
			ResourceID:   spIdpConnectionId,
		},
	}

	testutils.ValidateImportBlocks(t, resource, &expectedImportBlocks)
}

func createSpIdpConnection(t *testing.T, clientInfo *connector.PingFederateClientInfo, resourceType string) (string, string) {
	t.Helper()

	request := clientInfo.ApiClient.SpIdpConnectionsAPI.CreateConnection(clientInfo.Context)
	result := client.IdpConnection{
		Connection: client.Connection{
			EntityId: "TestEntityId",
			Id:       utils.Pointer("TestIdpConnectionId"),
			Name:     "TestIdpConnectionName",
		},
	}

	request = request.Body(result)

	resource, response, err := request.Execute()
	err = common.HandleClientResponse(response, err, "CreateConnection", resourceType)
	if err != nil {
		t.Fatalf("Failed to create test %s: %v", resourceType, err)
	}

	return *resource.Id, resource.Name
}

func deleteSpIdpConnection(t *testing.T, clientInfo *connector.PingFederateClientInfo, resourceType, id string) {
	t.Helper()

	request := clientInfo.ApiClient.SpIdpConnectionsAPI.DeleteConnection(clientInfo.Context, id)

	response, err := request.Execute()
	err = common.HandleClientResponse(response, err, "DeleteConnection", resourceType)
	if err != nil {
		t.Errorf("Failed to delete test %s: %v", resourceType, err)
	}
}
