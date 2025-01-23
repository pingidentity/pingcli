package resources_test

import (
	"testing"

	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/common"
	"github.com/pingidentity/pingcli/internal/connector/pingfederate/resources"
	"github.com/pingidentity/pingcli/internal/testing/testutils"
	client "github.com/pingidentity/pingfederate-go-client/v1210/configurationapi"
)

func Test_PingFederateSecretManager_Export(t *testing.T) {
	PingFederateClientInfo := testutils.GetPingFederateClientInfo(t)
	resource := resources.SecretManager(PingFederateClientInfo)

	secretManagerId, secretManagerName := createSecretManager(t, PingFederateClientInfo, resource.ResourceType())
	defer deleteSecretManager(t, PingFederateClientInfo, resource.ResourceType(), secretManagerId)

	expectedImportBlocks := []connector.ImportBlock{
		{
			ResourceType: resource.ResourceType(),
			ResourceName: secretManagerName,
			ResourceID:   secretManagerId,
		},
	}

	testutils.ValidateImportBlocks(t, resource, &expectedImportBlocks)
}

func createSecretManager(t *testing.T, clientInfo *connector.PingFederateClientInfo, resourceType string) (string, string) {
	t.Helper()

	request := clientInfo.ApiClient.SecretManagersAPI.CreateSecretManager(clientInfo.Context)
	result := client.SecretManager{
		Configuration: client.PluginConfiguration{},
		Id:            "TestSecretManagerId",
		Name:          "TestSecretManagerName",
		PluginDescriptorRef: client.ResourceLink{
			Id: "",
		},
	}

	request = request.Body(result)

	resource, response, err := request.Execute()
	err = common.HandleClientResponse(response, err, "CreateSecretManager", resourceType)
	if err != nil {
		t.Fatalf("Failed to create test %s: %v", resourceType, err)
	}

	return resource.Id, resource.Name
}

func deleteSecretManager(t *testing.T, clientInfo *connector.PingFederateClientInfo, resourceType, id string) {
	t.Helper()

	request := clientInfo.ApiClient.SecretManagersAPI.DeleteSecretManager(clientInfo.Context, id)

	response, err := request.Execute()
	err = common.HandleClientResponse(response, err, "DeleteSecretManager", resourceType)
	if err != nil {
		t.Errorf("Failed to delete test %s: %v", resourceType, err)
	}
}
