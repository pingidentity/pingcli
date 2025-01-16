package resources_test

import (
	"testing"

	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/common"
	"github.com/pingidentity/pingcli/internal/connector/pingfederate/resources"
	"github.com/pingidentity/pingcli/internal/testing/testutils"
	client "github.com/pingidentity/pingfederate-go-client/v1210/configurationapi"
)

func Test_PingFederateAuthenticationApiApplication_Export(t *testing.T) {
	PingFederateClientInfo := testutils.GetPingFederateClientInfo(t)
	resource := resources.AuthenticationApiApplication(PingFederateClientInfo)

	authenticationApiApplicationId, authenticationApiApplicationName := createAuthenticationApiApplication(t, PingFederateClientInfo, resource.ResourceType())
	defer deleteAuthenticationApiApplication(t, PingFederateClientInfo, resource.ResourceType(), authenticationApiApplicationId)

	expectedImportBlocks := []connector.ImportBlock{
		{
			ResourceType: resource.ResourceType(),
			ResourceName: authenticationApiApplicationName,
			ResourceID:   authenticationApiApplicationId,
		},
	}

	testutils.ValidateImportBlocks(t, resource, &expectedImportBlocks)
}

func createAuthenticationApiApplication(t *testing.T, clientInfo *connector.PingFederateClientInfo, resourceType string) (string, string) {
	t.Helper()

	request := clientInfo.ApiClient.AuthenticationApiAPI.CreateApplication(clientInfo.Context)
	result := client.AuthnApiApplication{}
	result.Id = "TestAuthnApiApplicationId"
	result.Name = "TestAuthnApiApplicationName"
	result.Url = "https://www.example.com"

	request = request.Body(result)

	resource, response, err := request.Execute()
	err = common.HandleClientResponse(response, err, "CreateApplication", resourceType)
	if err != nil {
		t.Fatalf("Failed to create test %s: %v", resourceType, err)
	}

	return resource.Id, resource.Name
}

func deleteAuthenticationApiApplication(t *testing.T, clientInfo *connector.PingFederateClientInfo, resourceType, id string) {
	t.Helper()

	request := clientInfo.ApiClient.AuthenticationApiAPI.DeleteApplication(clientInfo.Context, id)

	response, err := request.Execute()
	err = common.HandleClientResponse(response, err, "DeleteApplication", resourceType)
	if err != nil {
		t.Errorf("Failed to delete test %s: %v", resourceType, err)
	}
}
