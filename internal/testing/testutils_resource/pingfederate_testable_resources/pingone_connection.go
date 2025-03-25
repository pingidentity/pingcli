// Copyright © 2025 Ping Identity Corporation
// Code generated by ping-cli-generator

package pingfederate_testable_resources

import (
	"testing"

	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/common"
	"github.com/pingidentity/pingcli/internal/connector/pingfederate/resources"
	"github.com/pingidentity/pingcli/internal/testing/testutils_resource"
	"github.com/pingidentity/pingcli/internal/testing/testutils_resource/pingone_testable_resources"
	"github.com/pingidentity/pingcli/internal/utils"
	client "github.com/pingidentity/pingfederate-go-client/v1220/configurationapi"
)

func PingoneConnection(t *testing.T, clientInfo *connector.ClientInfo) *testutils_resource.TestableResource {
	t.Helper()

	return &testutils_resource.TestableResource{
		ClientInfo: clientInfo,
		CreateFunc: createPingoneConnection,
		DeleteFunc: deletePingoneConnection,
		Dependencies: []*testutils_resource.TestableResource{
			pingone_testable_resources.GatewayCredential(t, clientInfo),
		},
		ExportableResource: resources.PingoneConnection(clientInfo),
	}
}

func createPingoneConnection(t *testing.T, clientInfo *connector.ClientInfo, resourceType string, strArgs ...string) testutils_resource.ResourceInfo {
	t.Helper()

	if len(strArgs) != 1 {
		t.Fatalf("Unexpected number of arguments provided to createPingoneConnection(): %v", strArgs)
	}
	credential := strArgs[0]

	request := clientInfo.PingFederateApiClient.PingOneConnectionsAPI.CreatePingOneConnection(clientInfo.PingFederateContext)
	clientStruct := client.PingOneConnection{
		Credential: &credential,
		Id:         utils.Pointer("TestPingoneConnectionId"),
		Name:       "TestPingoneConnectionName",
	}

	request = request.Body(clientStruct)

	resource, response, err := request.Execute()
	ok, err := common.HandleClientResponse(response, err, "CreatePingOneConnection", resourceType)
	if err != nil {
		t.Fatalf("Failed to execute client function\nResponse Status: %s\nResponse Body: %s\nError: %v", response.Status, response.Body, err)
	}
	if !ok {
		t.Fatalf("Failed to execute client function\nResponse Status: %s\nResponse Body: %s", response.Status, response.Body)
	}

	return testutils_resource.ResourceInfo{
		DeletionIds: []string{},
		CreationInfo: map[testutils_resource.ResourceCreationInfoType]string{
			testutils_resource.ENUM_ID:   *resource.Id,
			testutils_resource.ENUM_NAME: resource.Name,
		},
	}
}

func deletePingoneConnection(t *testing.T, clientInfo *connector.ClientInfo, resourceType string, ids ...string) {
	t.Helper()

	if len(ids) != 1 {
		t.Fatalf("Unexpected number of arguments provided to deletePingoneConnection(): %v", ids)
	}
	id := ids[0]

	request := clientInfo.PingFederateApiClient.PingOneConnectionsAPI.DeletePingOneConnection(clientInfo.PingFederateContext, id)

	response, err := request.Execute()
	ok, err := common.HandleClientResponse(response, err, "DeletePingOneConnection", resourceType)
	if err != nil {
		t.Fatalf("Failed to execute client function\nResponse Status: %s\nResponse Body: %s\nError: %v", response.Status, response.Body, err)
	}
	if !ok {
		t.Fatalf("Failed to execute client function\nResponse Status: %s\nResponse Body: %s", response.Status, response.Body)
	}
}
