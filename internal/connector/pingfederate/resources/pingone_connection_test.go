package resources_test

import (
	"testing"

	"github.com/patrickcping/pingone-go-sdk-v2/management"
	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/common"
	"github.com/pingidentity/pingcli/internal/connector/pingfederate/resources"
	"github.com/pingidentity/pingcli/internal/testing/testutils"
	"github.com/pingidentity/pingcli/internal/utils"
	client "github.com/pingidentity/pingfederate-go-client/v1210/configurationapi"
)

func Test_PingFederatePingoneConnection_Export(t *testing.T) {
	PingFederateClientInfo := testutils.GetPingFederateClientInfo(t)
	pingOneClientInfo := testutils.GetPingOneClientInfo(t)
	resource := resources.PingoneConnection(PingFederateClientInfo)

	gatewayId := createPingOnePingFederateGateway(t, pingOneClientInfo, resource.ResourceType())
	defer deletePingOnePingFederateGateway(t, pingOneClientInfo, resource.ResourceType(), gatewayId)

	gatewayCredential := createPingOnePingFederateGatewayCredential(t, pingOneClientInfo, resource.ResourceType(), gatewayId)

	pingoneConnectionId, pingoneConnectionName := createPingoneConnection(t, PingFederateClientInfo, resource.ResourceType(), gatewayCredential)
	defer deletePingoneConnection(t, PingFederateClientInfo, resource.ResourceType(), pingoneConnectionId)

	expectedImportBlocks := []connector.ImportBlock{
		{
			ResourceType: resource.ResourceType(),
			ResourceName: pingoneConnectionName,
			ResourceID:   pingoneConnectionId,
		},
	}

	testutils.ValidateImportBlocks(t, resource, &expectedImportBlocks)
}

func createPingoneConnection(t *testing.T, clientInfo *connector.PingFederateClientInfo, resourceType, credential string) (string, string) {
	t.Helper()

	request := clientInfo.ApiClient.PingOneConnectionsAPI.CreatePingOneConnection(clientInfo.Context)
	result := client.PingOneConnection{
		Credential: &credential,
		Id:         utils.Pointer("TestPingoneConnectionId"),
		Name:       "TestPingoneConnectionName",
	}

	request = request.Body(result)

	resource, response, err := request.Execute()
	err = common.HandleClientResponse(response, err, "CreatePingOneConnection", resourceType)
	if err != nil {
		t.Fatalf("Failed to create test %s: %v", resourceType, err)
	}

	return *resource.Id, resource.Name
}

func deletePingoneConnection(t *testing.T, clientInfo *connector.PingFederateClientInfo, resourceType, id string) {
	t.Helper()

	request := clientInfo.ApiClient.PingOneConnectionsAPI.DeletePingOneConnection(clientInfo.Context, id)

	response, err := request.Execute()
	err = common.HandleClientResponse(response, err, "DeletePingOneConnection", resourceType)
	if err != nil {
		t.Errorf("Failed to delete test %s: %v", resourceType, err)
	}
}

func createPingOnePingFederateGateway(t *testing.T, pingOneClientInfo *connector.PingOneClientInfo, resourceType string) string {
	t.Helper()

	result := management.CreateGatewayRequest{
		Gateway: &management.Gateway{
			Enabled: true,
			Name:    "TestPingFederateGateway",
			Type:    management.ENUMGATEWAYTYPE_PING_FEDERATE,
		},
	}

	createGateway201Response, response, err := pingOneClientInfo.ApiClient.ManagementAPIClient.GatewaysApi.CreateGateway(pingOneClientInfo.Context, testutils.GetEnvironmentID()).CreateGatewayRequest(result).Execute()
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

	return *gatewayId
}

func deletePingOnePingFederateGateway(t *testing.T, pingOneClientInfo *connector.PingOneClientInfo, resourceType, gatewayId string) {
	t.Helper()

	response, err := pingOneClientInfo.ApiClient.ManagementAPIClient.GatewaysApi.DeleteGateway(pingOneClientInfo.Context, testutils.GetEnvironmentID(), gatewayId).Execute()
	err = common.HandleClientResponse(response, err, "DeleteGateway", resourceType)
	if err != nil {
		t.Errorf("Failed to delete test %s: %v", resourceType, err)
	}
}

func createPingOnePingFederateGatewayCredential(t *testing.T, pingOneClientInfo *connector.PingOneClientInfo, resourceType, gatewayId string) string {
	t.Helper()

	gatewayCredential, response, err := pingOneClientInfo.ApiClient.ManagementAPIClient.GatewayCredentialsApi.CreateGatewayCredential(pingOneClientInfo.Context, testutils.GetEnvironmentID(), gatewayId).Execute()
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

	return *credential
}
