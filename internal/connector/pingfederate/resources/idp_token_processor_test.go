package resources_test

import (
	"testing"

	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/common"
	"github.com/pingidentity/pingcli/internal/connector/pingfederate/resources"
	"github.com/pingidentity/pingcli/internal/testing/testutils"
	client "github.com/pingidentity/pingfederate-go-client/v1210/configurationapi"
)

func Test_PingFederateIdpTokenProcessor_Export(t *testing.T) {
	PingFederateClientInfo := testutils.GetPingFederateClientInfo(t)
	resource := resources.IdpTokenProcessor(PingFederateClientInfo)

	idpTokenProcessorId, idpTokenProcessorName := createIdpTokenProcessor(t, PingFederateClientInfo, resource.ResourceType())
	defer deleteIdpTokenProcessor(t, PingFederateClientInfo, resource.ResourceType(), idpTokenProcessorId)

	expectedImportBlocks := []connector.ImportBlock{
		{
			ResourceType: resource.ResourceType(),
			ResourceName: idpTokenProcessorName,
			ResourceID:   idpTokenProcessorId,
		},
	}

	testutils.ValidateImportBlocks(t, resource, &expectedImportBlocks)
}

func createIdpTokenProcessor(t *testing.T, clientInfo *connector.PingFederateClientInfo, resourceType string) (string, string) {
	t.Helper()

	request := clientInfo.ApiClient.IdpTokenProcessorsAPI.CreateTokenProcessor(clientInfo.Context)
	result := client.TokenProcessor{}
	result.Id = "TestTokenProcessorId"
	result.Name = "TestTokenProcessorName"
	result.Configuration = client.PluginConfiguration{}
	result.PluginDescriptorRef = client.ResourceLink{
		Id: "",
	}

	request = request.Body(result)

	resource, response, err := request.Execute()
	err = common.HandleClientResponse(response, err, "CreateTokenProcessor", resourceType)
	if err != nil {
		t.Fatalf("Failed to create test %s: %v", resourceType, err)
	}

	return resource.Id, resource.Name
}

func deleteIdpTokenProcessor(t *testing.T, clientInfo *connector.PingFederateClientInfo, resourceType, id string) {
	t.Helper()

	request := clientInfo.ApiClient.IdpTokenProcessorsAPI.DeleteTokenProcessor(clientInfo.Context, id)

	response, err := request.Execute()
	err = common.HandleClientResponse(response, err, "DeleteTokenProcessor", resourceType)
	if err != nil {
		t.Errorf("Failed to delete test %s: %v", resourceType, err)
	}
}
