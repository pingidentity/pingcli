package pingfederate

import (
	"testing"

	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/common"
	"github.com/pingidentity/pingcli/internal/connector/pingfederate/resources"
	"github.com/pingidentity/pingcli/internal/testing/testutils_resource"
	"github.com/pingidentity/pingcli/internal/testing/testutils_resource/pingone"
	"github.com/pingidentity/pingcli/internal/utils"
	client "github.com/pingidentity/pingfederate-go-client/v1210/configurationapi"
)

func TestableResource_PingFederatePingoneConnection(t *testing.T, clientInfo *connector.ClientInfo) testutils_resource.TestableResource {
	t.Helper()

	return testutils_resource.TestableResource{
		ClientInfo: clientInfo,
		CreateFunc: createPingoneConnection,
		DeleteFunc: deletePingoneConnection,
		Dependencies: []testutils_resource.TestableResource{
			pingone.TestableResource_PingOnePingFederateGatewayCredential(t, clientInfo),
		},
		ExportableResource: resources.PingoneConnection(clientInfo),
	}
}

func createPingoneConnection(t *testing.T, clientInfo *connector.ClientInfo, strArgs ...string) testutils_resource.ResourceCreationInfo {
	t.Helper()

	if len(strArgs) != 1 { //TODO
		t.Fatalf("Unexpected number of arguments provided to createPingoneConnection(): %v", strArgs)
	}
	resourceType := strArgs[0]
	credential := strArgs[1]

	request := clientInfo.PingFederateApiClient.PingOneConnectionsAPI.CreatePingOneConnection(clientInfo.Context)
	clientStruct := client.PingOneConnection{
		Credential: &credential,
		Id:         utils.Pointer("TestPingoneConnectionId"),
		Name:       "TestPingoneConnectionName",
	}

	request = request.Body(clientStruct)

	resource, response, err := request.Execute()
	ok, err := common.HandleClientResponse(response, err, "CreatePingOneConnection", resourceType)
	if err != nil {
		t.Fatalf("Failed to create test %s: %v", resourceType, err)
	}
	if !ok {
		t.Fatalf("Failed to create test %s: non-ok response", resourceType)
	}

	return testutils_resource.ResourceCreationInfo{
		testutils_resource.ENUM_ID:   *resource.Id,
		testutils_resource.ENUM_NAME: resource.Name,
	}
}

func deletePingoneConnection(t *testing.T, clientInfo *connector.ClientInfo, resourceType, id string) {
	t.Helper()

	request := clientInfo.PingFederateApiClient.PingOneConnectionsAPI.DeletePingOneConnection(clientInfo.Context, id)

	response, err := request.Execute()
	ok, err := common.HandleClientResponse(response, err, "DeletePingOneConnection", resourceType)
	if err != nil {
		t.Errorf("Failed to delete test %s: %v", resourceType, err)
	}
	if !ok {
		t.Fatalf("Failed to delete test %s: non-ok response", resourceType)
	}
}
