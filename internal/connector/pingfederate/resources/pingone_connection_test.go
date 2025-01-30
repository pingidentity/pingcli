package resources_test

import (
	"testing"

	"github.com/patrickcping/pingone-go-sdk-v2/management"
	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/common"
	"github.com/pingidentity/pingcli/internal/connector/pingfederate/resources"
	"github.com/pingidentity/pingcli/internal/testing/testutils"
	"github.com/pingidentity/pingcli/internal/testing/testutils_resource"
	"github.com/pingidentity/pingcli/internal/utils"
	client "github.com/pingidentity/pingfederate-go-client/v1210/configurationapi"
)

func TestableResource_PingFederatePingoneConnection(t *testing.T) *testutils_resource.TestableResource {
	t.Helper()

	pingfederateClientInfo := testutils.GetPingFederateClientInfo(t)
	return &testutils_resource.TestableResource{
		ClientInfo:         pingfederateClientInfo,
		ExportableResource: resources.AuthenticationApiApplication(pingfederateClientInfo),
		TestResource: testutils_resource.TestResource{
			Dependencies: []testutils_resource.TestResource{
				{
					Dependencies: []testutils_resource.TestResource{
						{
							Dependencies: nil,
							CreateFunc:   createPingOnePingFederateGateway,
							DeleteFunc:   deletePingOnePingFederateGateway,
						},
					},
					CreateFunc: createPingOnePingFederateGatewayCredential,
					DeleteFunc: nil,
				},
			},
			CreateFunc: createPingoneConnection,
			DeleteFunc: deletePingoneConnection,
		},
	}
}

func Test_PingFederatePingoneConnection(t *testing.T) {
	tr := TestableResource_PingFederatePingoneConnection(t)

	creationInfo := tr.CreateResource(t, tr.TestResource)
	defer tr.DeleteResource(t, tr.TestResource)

	expectedImportBlocks := []connector.ImportBlock{
		{
			ResourceType: tr.ExportableResource.ResourceType(),
			ResourceName: creationInfo[testutils_resource.ENUM_NAME],
			ResourceID:   creationInfo[testutils_resource.ENUM_ID],
		},
	}

	testutils.ValidateImportBlocks(t, tr.ExportableResource, &expectedImportBlocks)
}

func createPingoneConnection(t *testing.T, clientInfo *connector.ClientInfo, strArgs ...string) testutils_resource.ResourceCreationInfo {
	t.Helper()

	if len(strArgs) != 2 {
		t.Fatalf("Unexpected number of arguments provided to createPingoneConnection(): %v", strArgs)
	}
	resourceType := strArgs[0]
	credential := strArgs[1]

	request := clientInfo.PingFederateApiClient.PingOneConnectionsAPI.CreatePingOneConnection(clientInfo.Context)
	result := client.PingOneConnection{
		Credential: &credential,
		Id:         utils.Pointer("TestPingoneConnectionId"),
		Name:       "TestPingoneConnectionName",
	}

	request = request.Body(result)

	resource, response, err := request.Execute()
	err = common.HandleClientResponse(response, err, "CreateApplication", resourceType)
	if err != nil {
		t.Fatalf("Failed to create test %s: %v", resourceType, err)
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
	err = common.HandleClientResponse(response, err, "DeletePingOneConnection", resourceType)
	if err != nil {
		t.Errorf("Failed to delete test %s: %v", resourceType, err)
	}
}

func createPingOnePingFederateGateway(t *testing.T, pingOneClientInfo *connector.ClientInfo, strArgs ...string) testutils_resource.ResourceCreationInfo {
	t.Helper()

	if len(strArgs) != 1 {
		t.Fatalf("Unexpected number of arguments provided to createPingoneConnection(): %v", strArgs)
	}
	resourceType := strArgs[0]

	result := management.CreateGatewayRequest{
		Gateway: &management.Gateway{
			Enabled: true,
			Name:    "TestPingFederateGateway",
			Type:    management.ENUMGATEWAYTYPE_PING_FEDERATE,
		},
	}

	createGateway201Response, response, err := pingOneClientInfo.PingOneApiClient.ManagementAPIClient.GatewaysApi.CreateGateway(pingOneClientInfo.Context, testutils.GetEnvironmentID()).CreateGatewayRequest(result).Execute()
	err = common.HandleClientResponse(response, err, "CreateGateway", resourceType)
	if err != nil {
		t.Fatalf("Failed to create test %s: %v", resourceType, err)
	}

	if createGateway201Response == nil || createGateway201Response.Gateway == nil {
		t.Fatalf("Failed to create test %s: %v", resourceType, err)
	}

	gatewayId, gatewayIdOk := createGateway201Response.Gateway.GetIdOk()
	if !gatewayIdOk {
		t.Fatalf("Failed to create test %s: %v", resourceType, err)
	}

	return testutils_resource.ResourceCreationInfo{
		testutils_resource.ENUM_ID: *gatewayId,
	}
}

func deletePingOnePingFederateGateway(t *testing.T, pingOneClientInfo *connector.ClientInfo, resourceType, gatewayId string) {
	t.Helper()

	response, err := pingOneClientInfo.PingOneApiClient.ManagementAPIClient.GatewaysApi.DeleteGateway(pingOneClientInfo.Context, testutils.GetEnvironmentID(), gatewayId).Execute()
	err = common.HandleClientResponse(response, err, "DeleteGateway", resourceType)
	if err != nil {
		t.Errorf("Failed to delete test %s: %v", resourceType, err)
	}
}

func createPingOnePingFederateGatewayCredential(t *testing.T, pingOneClientInfo *connector.ClientInfo, strArgs ...string) testutils_resource.ResourceCreationInfo {
	t.Helper()

	if len(strArgs) != 2 {
		t.Fatalf("Unexpected number of arguments provided to createPingoneConnection(): %v", strArgs)
	}
	resourceType := strArgs[0]
	gatewayId := strArgs[1]

	gatewayCredential, response, err := pingOneClientInfo.PingOneApiClient.ManagementAPIClient.GatewayCredentialsApi.CreateGatewayCredential(pingOneClientInfo.Context, testutils.GetEnvironmentID(), gatewayId).Execute()
	err = common.HandleClientResponse(response, err, "CreateGatewayCredential", resourceType)
	if err != nil {
		t.Fatalf("Failed to create test %s: %v", resourceType, err)
	}

	if gatewayCredential == nil {
		t.Fatalf("Failed to create test %s: %v", resourceType, err)
	}

	credential, credentialOk := gatewayCredential.GetCredentialOk()
	if !credentialOk {
		t.Fatalf("Failed to create test %s: %v", resourceType, err)
	}

	return testutils_resource.ResourceCreationInfo{
		testutils_resource.ENUM_CREDENTIAL: *credential,
	}
}
