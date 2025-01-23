package resources_test

import (
	"testing"

	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/common"
	"github.com/pingidentity/pingcli/internal/connector/pingfederate/resources"
	"github.com/pingidentity/pingcli/internal/testing/testutils"
	client "github.com/pingidentity/pingfederate-go-client/v1210/configurationapi"
)

func Test_PingFederateOauthAccessTokenManager_Export(t *testing.T) {
	PingFederateClientInfo := testutils.GetPingFederateClientInfo(t)
	resource := resources.OauthAccessTokenManager(PingFederateClientInfo)

	oauthAccessTokenManagerId, oauthAccessTokenManagerName := createOauthAccessTokenManager(t, PingFederateClientInfo, resource.ResourceType())
	defer deleteOauthAccessTokenManager(t, PingFederateClientInfo, resource.ResourceType(), oauthAccessTokenManagerId)

	expectedImportBlocks := []connector.ImportBlock{
		{
			ResourceType: resource.ResourceType(),
			ResourceName: oauthAccessTokenManagerName,
			ResourceID:   oauthAccessTokenManagerId,
		},
	}

	testutils.ValidateImportBlocks(t, resource, &expectedImportBlocks)
}

func createOauthAccessTokenManager(t *testing.T, clientInfo *connector.PingFederateClientInfo, resourceType string) (string, string) {
	t.Helper()

	request := clientInfo.ApiClient.OauthAccessTokenManagersAPI.CreateTokenManager(clientInfo.Context)
	result := client.AccessTokenManager{}
	result.Id = "TestAccessTokenManagerId"
	result.Name = "TestAccessTokenManagerName"
	result.Configuration = client.PluginConfiguration{}
	result.PluginDescriptorRef = client.ResourceLink{}
	result.PluginDescriptorRef.Id = ""

	request = request.Body(result)

	resource, response, err := request.Execute()
	err = common.HandleClientResponse(response, err, "CreateTokenManager", resourceType)
	if err != nil {
		t.Fatalf("Failed to create test %s: %v", resourceType, err)
	}

	return resource.Id, resource.Name
}

func deleteOauthAccessTokenManager(t *testing.T, clientInfo *connector.PingFederateClientInfo, resourceType, id string) {
	t.Helper()

	request := clientInfo.ApiClient.OauthAccessTokenManagersAPI.DeleteTokenManager(clientInfo.Context, id)

	response, err := request.Execute()
	err = common.HandleClientResponse(response, err, "DeleteTokenManager", resourceType)
	if err != nil {
		t.Errorf("Failed to delete test %s: %v", resourceType, err)
	}
}
