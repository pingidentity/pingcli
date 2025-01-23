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

func Test_PingFederateIdpSpConnection_Export(t *testing.T) {
	PingFederateClientInfo := testutils.GetPingFederateClientInfo(t)
	resource := resources.IdpSpConnection(PingFederateClientInfo)

	idpSpConnectionId, idpSpConnectionName := createIdpSpConnection(t, PingFederateClientInfo, resource.ResourceType())
	defer deleteIdpSpConnection(t, PingFederateClientInfo, resource.ResourceType(), idpSpConnectionId)

	expectedImportBlocks := []connector.ImportBlock{
		{
			ResourceType: resource.ResourceType(),
			ResourceName: idpSpConnectionName,
			ResourceID:   idpSpConnectionId,
		},
	}

	testutils.ValidateImportBlocks(t, resource, &expectedImportBlocks)
}

func createIdpSpConnection(t *testing.T, clientInfo *connector.PingFederateClientInfo, resourceType string) (string, string) {
	t.Helper()

	request := clientInfo.ApiClient.IdpSpConnectionsAPI.CreateSpConnection(clientInfo.Context)
	result := client.SpConnection{}
	result.Id = utils.Pointer("TestSpConnectionId")
	result.Name = "TestSpConnectionName"
	result.EntityId = "TestSpConnectionEntityId"

	request = request.Body(result)

	resource, response, err := request.Execute()
	err = common.HandleClientResponse(response, err, "CreateSpConnection", resourceType)
	if err != nil {
		t.Fatalf("Failed to create test %s: %v", resourceType, err)
	}

	return *resource.Id, resource.Name
}

func deleteIdpSpConnection(t *testing.T, clientInfo *connector.PingFederateClientInfo, resourceType, id string) {
	t.Helper()

	request := clientInfo.ApiClient.IdpSpConnectionsAPI.DeleteSpConnection(clientInfo.Context, id)

	response, err := request.Execute()
	err = common.HandleClientResponse(response, err, "DeleteSpConnection", resourceType)
	if err != nil {
		t.Errorf("Failed to delete test %s: %v", resourceType, err)
	}
}
