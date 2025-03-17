package pingfederate

import (
	"testing"

	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/common"
	"github.com/pingidentity/pingcli/internal/connector/pingfederate/resources"
	"github.com/pingidentity/pingcli/internal/testing/testutils_resource"
	"github.com/pingidentity/pingcli/internal/utils"
	client "github.com/pingidentity/pingfederate-go-client/v1220/configurationapi"
)

func TestableResource_PingFederateMetadataUrl(t *testing.T, clientInfo *connector.ClientInfo) *testutils_resource.TestableResource {
	t.Helper()

	return &testutils_resource.TestableResource{
		ClientInfo:         clientInfo,
		CreateFunc:         createMetadataUrl,
		DeleteFunc:         deleteMetadataUrl,
		Dependencies:       nil,
		ExportableResource: resources.MetadataUrl(clientInfo),
	}
}

func createMetadataUrl(t *testing.T, clientInfo *connector.ClientInfo, strArgs ...string) testutils_resource.ResourceCreationInfo {
	t.Helper()

	if len(strArgs) != 1 {
		t.Fatalf("Unexpected number of arguments provided to createMetadataUrl(): %v", strArgs)
	}
	resourceType := strArgs[0]

	request := clientInfo.PingFederateApiClient.MetadataUrlsAPI.AddMetadataUrl(clientInfo.PingFederateContext)
	clientStruct := client.MetadataUrl{
		Id:   utils.Pointer("TestMetadataUrlId"),
		Name: "TestMetadataUrlName",
		Url:  "https://www.example.com",
	}

	request = request.Body(clientStruct)

	resource, response, err := request.Execute()
	ok, err := common.HandleClientResponse(response, err, "AddMetadataUrl", resourceType)
	if err != nil {
		t.Fatalf("Failed to execute client function\nResponse Status: %s\nResponse Body: %s\nError: %v", response.Status, response.Body, err)
	}
	if !ok {
		t.Fatalf("Failed to execute client function\nResponse Status: %s\nResponse Body: %s", response.Status, response.Body)
	}

	return testutils_resource.ResourceCreationInfo{
		testutils_resource.ENUM_ID:   *resource.Id,
		testutils_resource.ENUM_NAME: resource.Name,
	}
}

func deleteMetadataUrl(t *testing.T, clientInfo *connector.ClientInfo, resourceType, id string) {
	t.Helper()

	request := clientInfo.PingFederateApiClient.MetadataUrlsAPI.DeleteMetadataUrl(clientInfo.PingFederateContext, id)

	response, err := request.Execute()
	ok, err := common.HandleClientResponse(response, err, "DeleteMetadataUrl", resourceType)
	if err != nil {
		t.Fatalf("Failed to execute client function\nResponse Status: %s\nResponse Body: %s\nError: %v", response.Status, response.Body, err)
	}
	if !ok {
		t.Fatalf("Failed to execute client function\nResponse Status: %s\nResponse Body: %s", response.Status, response.Body)
	}
}
