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

func Test_PingFederateMetadataUrl_Export(t *testing.T) {
	PingFederateClientInfo := testutils.GetPingFederateClientInfo(t)
	resource := resources.MetadataUrl(PingFederateClientInfo)

	metadataUrlId, metadataUrlName := createMetadataUrl(t, PingFederateClientInfo, resource.ResourceType())
	defer deleteMetadataUrl(t, PingFederateClientInfo, resource.ResourceType(), metadataUrlId)

	expectedImportBlocks := []connector.ImportBlock{
		{
			ResourceType: resource.ResourceType(),
			ResourceName: metadataUrlName,
			ResourceID:   metadataUrlId,
		},
	}

	testutils.ValidateImportBlocks(t, resource, &expectedImportBlocks)
}

func createMetadataUrl(t *testing.T, clientInfo *connector.PingFederateClientInfo, resourceType string) (string, string) {
	t.Helper()

	request := clientInfo.ApiClient.MetadataUrlsAPI.AddMetadataUrl(clientInfo.Context)
	result := client.MetadataUrl{}
	result.Id = utils.Pointer("TestMetadataUrlId")
	result.Name = "TestMetadataUrlName"
	result.Url = "https://www.example.com"

	request = request.Body(result)

	resource, response, err := request.Execute()
	err = common.HandleClientResponse(response, err, "AddMetadataUrl", resourceType)
	if err != nil {
		t.Fatalf("Failed to create test %s: %v", resourceType, err)
	}

	return *resource.Id, resource.Name
}

func deleteMetadataUrl(t *testing.T, clientInfo *connector.PingFederateClientInfo, resourceType, id string) {
	t.Helper()

	request := clientInfo.ApiClient.MetadataUrlsAPI.DeleteMetadataUrl(clientInfo.Context, id)

	response, err := request.Execute()
	err = common.HandleClientResponse(response, err, "DeleteMetadataUrl", resourceType)
	if err != nil {
		t.Errorf("Failed to delete test %s: %v", resourceType, err)
	}
}
