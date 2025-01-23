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

func Test_PingFederatePingoneConnection_Export(t *testing.T) {
	PingFederateClientInfo := testutils.GetPingFederateClientInfo(t)
	resource := resources.PingoneConnection(PingFederateClientInfo)

	pingoneConnectionId, pingoneConnectionName := createPingoneConnection(t, PingFederateClientInfo, resource.ResourceType())
	defer deletePingoneConnection(t, PingFederateClientInfo, resource.ResourceType(), pingoneConnectionId)

	expectedImportBlocks := []connector.ImportBlock{
		{
			ResourceType: resource.ResourceType(),
			ResourceName: pingoneConnectionName,
			ResourceID:   pingoneConnectionId,
		},
	}

	testutils.ValidateImportBlocks(t, resource, &expectedImportBlocks)
}

func createPingoneConnection(t *testing.T, clientInfo *connector.PingFederateClientInfo, resourceType string) (string, string) {
	t.Helper()

	request := clientInfo.ApiClient.PingOneConnectionsAPI.CreatePingOneConnection(clientInfo.Context)
	result := client.PingOneConnection{
		Id:   utils.Pointer("TestPingoneConnectionId"),
		Name: "TestPingoneConnectionName",
	}

	request = request.Body(result)

	resource, response, err := request.Execute()
	err = common.HandleClientResponse(response, err, "CreatePingOneConnection", resourceType)
	if err != nil {
		t.Fatalf("Failed to create test %s: %v", resourceType, err)
	}

	return *resource.Id, resource.Name
}

func deletePingoneConnection(t *testing.T, clientInfo *connector.PingFederateClientInfo, resourceType, id string) {
	t.Helper()

	request := clientInfo.ApiClient.PingOneConnectionsAPI.DeletePingOneConnection(clientInfo.Context, id)

	response, err := request.Execute()
	err = common.HandleClientResponse(response, err, "DeletePingOneConnection", resourceType)
	if err != nil {
		t.Errorf("Failed to delete test %s: %v", resourceType, err)
	}
}
