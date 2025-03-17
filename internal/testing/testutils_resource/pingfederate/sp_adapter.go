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

func TestableResource_PingFederateSpAdapter(t *testing.T, clientInfo *connector.ClientInfo) *testutils_resource.TestableResource {
	t.Helper()

	return &testutils_resource.TestableResource{
		ClientInfo:         clientInfo,
		CreateFunc:         createSpAdapter,
		DeleteFunc:         deleteSpAdapter,
		Dependencies:       nil,
		ExportableResource: resources.SpAdapter(clientInfo),
	}
}

func createSpAdapter(t *testing.T, clientInfo *connector.ClientInfo, strArgs ...string) testutils_resource.ResourceCreationInfo {
	t.Helper()

	if len(strArgs) != 1 {
		t.Fatalf("Unexpected number of arguments provided to createSpAdapter(): %v", strArgs)
	}
	resourceType := strArgs[0]

	request := clientInfo.PingFederateApiClient.SpAdaptersAPI.CreateSpAdapter(clientInfo.PingFederateContext)
	clientStruct := client.SpAdapter{
		Configuration: client.PluginConfiguration{
			Fields: []client.ConfigField{
				{
					Name:  "Password",
					Value: utils.Pointer("TestPassword1"),
				},
				{
					Name:  "Confirm Password",
					Value: utils.Pointer("TestPassword1"),
				},
			},
		},
		Id:   "TestSpAdapterId",
		Name: "TestSpAdapterName",
		PluginDescriptorRef: client.ResourceLink{
			Id: "com.pingidentity.adapters.opentoken.SpAuthnAdapter",
		},
	}

	request = request.Body(clientStruct)

	resource, response, err := request.Execute()
	ok, err := common.HandleClientResponse(response, err, "CreateSpAdapter", resourceType)
	if err != nil {
		t.Fatalf("Failed to execute client function\nResponse Status: %s\nResponse Body: %s\nError: %v", response.Status, response.Body, err)
	}
	if !ok {
		t.Fatalf("Failed to execute client function\nResponse Status: %s\nResponse Body: %s", response.Status, response.Body)
	}

	return testutils_resource.ResourceCreationInfo{
		testutils_resource.ENUM_ID:   resource.Id,
		testutils_resource.ENUM_NAME: resource.Name,
	}
}

func deleteSpAdapter(t *testing.T, clientInfo *connector.ClientInfo, resourceType, id string) {
	t.Helper()

	request := clientInfo.PingFederateApiClient.SpAdaptersAPI.DeleteSpAdapter(clientInfo.PingFederateContext, id)

	response, err := request.Execute()
	ok, err := common.HandleClientResponse(response, err, "DeleteSpAdapter", resourceType)
	if err != nil {
		t.Fatalf("Failed to execute client function\nResponse Status: %s\nResponse Body: %s\nError: %v", response.Status, response.Body, err)
	}
	if !ok {
		t.Fatalf("Failed to execute client function\nResponse Status: %s\nResponse Body: %s", response.Status, response.Body)
	}
}
