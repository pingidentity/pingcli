package resources_test

import (
	"testing"

	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/common"
	"github.com/pingidentity/pingcli/internal/connector/pingfederate/resources"
	"github.com/pingidentity/pingcli/internal/testing/testutils"
	client "github.com/pingidentity/pingfederate-go-client/v1210/configurationapi"
)

func Test_PingFederateSpAdapter_Export(t *testing.T) {
	PingFederateClientInfo := testutils.GetPingFederateClientInfo(t)
	resource := resources.SpAdapter(PingFederateClientInfo)

	spAdapterId, spAdapterName := createSpAdapter(t, PingFederateClientInfo, resource.ResourceType())
	defer deleteSpAdapter(t, PingFederateClientInfo, resource.ResourceType(), spAdapterId)

	expectedImportBlocks := []connector.ImportBlock{
		{
			ResourceType: resource.ResourceType(),
			ResourceName: spAdapterName,
			ResourceID:   spAdapterId,
		},
	}

	testutils.ValidateImportBlocks(t, resource, &expectedImportBlocks)
}

func createSpAdapter(t *testing.T, clientInfo *connector.PingFederateClientInfo, resourceType string) (string, string) {
	t.Helper()

	request := clientInfo.ApiClient.SpAdaptersAPI.CreateSpAdapter(clientInfo.Context)
	result := client.SpAdapter{
		Configuration: client.PluginConfiguration{},
		Id:            "TestSpAdapterId",
		Name:          "TestSpAdapterName",
		PluginDescriptorRef: client.ResourceLink{
			Id: "",
		},
	}

	request = request.Body(result)

	resource, response, err := request.Execute()
	err = common.HandleClientResponse(response, err, "CreateSpAdapter", resourceType)
	if err != nil {
		t.Fatalf("Failed to create test %s: %v", resourceType, err)
	}

	return resource.Id, resource.Name
}

func deleteSpAdapter(t *testing.T, clientInfo *connector.PingFederateClientInfo, resourceType, id string) {
	t.Helper()

	request := clientInfo.ApiClient.SpAdaptersAPI.DeleteSpAdapter(clientInfo.Context, id)

	response, err := request.Execute()
	err = common.HandleClientResponse(response, err, "DeleteSpAdapter", resourceType)
	if err != nil {
		t.Errorf("Failed to delete test %s: %v", resourceType, err)
	}
}
