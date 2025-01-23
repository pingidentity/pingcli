package resources_test

import (
	"testing"

	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/common"
	"github.com/pingidentity/pingcli/internal/connector/pingfederate/resources"
	"github.com/pingidentity/pingcli/internal/testing/testutils"
	client "github.com/pingidentity/pingfederate-go-client/v1210/configurationapi"
)

func Test_PingFederatePasswordCredentialValidator_Export(t *testing.T) {
	PingFederateClientInfo := testutils.GetPingFederateClientInfo(t)
	resource := resources.PasswordCredentialValidator(PingFederateClientInfo)

	passwordCredentialValidatorId, passwordCredentialValidatorName := createPasswordCredentialValidator(t, PingFederateClientInfo, resource.ResourceType())
	defer deletePasswordCredentialValidator(t, PingFederateClientInfo, resource.ResourceType(), passwordCredentialValidatorId)

	expectedImportBlocks := []connector.ImportBlock{
		{
			ResourceType: resource.ResourceType(),
			ResourceName: passwordCredentialValidatorName,
			ResourceID:   passwordCredentialValidatorId,
		},
	}

	testutils.ValidateImportBlocks(t, resource, &expectedImportBlocks)
}

func createPasswordCredentialValidator(t *testing.T, clientInfo *connector.PingFederateClientInfo, resourceType string) (string, string) {
	t.Helper()

	request := clientInfo.ApiClient.PasswordCredentialValidatorsAPI.CreatePasswordCredentialValidator(clientInfo.Context)
	result := client.PasswordCredentialValidator{
		Configuration: client.PluginConfiguration{},
		Id:            "TestPasswordCredentialValidatorId",
		Name:          "TestPasswordCredentialValidatorName",
		PluginDescriptorRef: client.ResourceLink{
			Id: "",
		},
	}

	request = request.Body(result)

	resource, response, err := request.Execute()
	err = common.HandleClientResponse(response, err, "CreatePasswordCredentialValidator", resourceType)
	if err != nil {
		t.Fatalf("Failed to create test %s: %v", resourceType, err)
	}

	return resource.Id, resource.Name
}

func deletePasswordCredentialValidator(t *testing.T, clientInfo *connector.PingFederateClientInfo, resourceType, id string) {
	t.Helper()

	request := clientInfo.ApiClient.PasswordCredentialValidatorsAPI.DeletePasswordCredentialValidator(clientInfo.Context, id)

	response, err := request.Execute()
	err = common.HandleClientResponse(response, err, "DeletePasswordCredentialValidator", resourceType)
	if err != nil {
		t.Errorf("Failed to delete test %s: %v", resourceType, err)
	}
}
