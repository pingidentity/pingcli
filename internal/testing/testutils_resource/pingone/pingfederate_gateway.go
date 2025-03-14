package pingone

import (
	"testing"

	"github.com/patrickcping/pingone-go-sdk-v2/management"
	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/common"
	"github.com/pingidentity/pingcli/internal/testing/testutils"
	"github.com/pingidentity/pingcli/internal/testing/testutils_resource"
)

func TestableResource_PingOnePingFederateGateway(t *testing.T, clientInfo *connector.ClientInfo) testutils_resource.TestableResource {
	t.Helper()

	return testutils_resource.TestableResource{
		ClientInfo:         clientInfo,
		CreateFunc:         createPingFederateGateway,
		DeleteFunc:         deletePingFederateGateway,
		Dependencies:       nil,
		ExportableResource: nil,
	}
}

func createPingFederateGateway(t *testing.T, clientInfo *connector.ClientInfo, strArgs ...string) testutils_resource.ResourceCreationInfo {
	t.Helper()

	if len(strArgs) != 1 {
		t.Fatalf("Unexpected number of arguments provided to createPingFederateGateway(): %v", strArgs)
	}
	resourceType := strArgs[0]

	result := management.CreateGatewayRequest{
		Gateway: &management.Gateway{
			Enabled: true,
			Name:    "TestPingFederateGateway",
			Type:    management.ENUMGATEWAYTYPE_PING_FEDERATE,
		},
	}

	createGateway201Response, response, err := clientInfo.PingOneApiClient.ManagementAPIClient.GatewaysApi.CreateGateway(clientInfo.Context, testutils.GetEnvironmentID()).CreateGatewayRequest(result).Execute()
	ok, err := common.HandleClientResponse(response, err, "CreateGateway", resourceType)
	if err != nil {
		t.Fatalf("Failed to create test %s: %v", resourceType, err)
	}
	if !ok {
		t.Fatalf("Failed to create test %s: non-ok response", resourceType)
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

func deletePingFederateGateway(t *testing.T, clientInfo *connector.ClientInfo, resourceType, id string) {
	t.Helper()

	response, err := clientInfo.PingOneApiClient.ManagementAPIClient.GatewaysApi.DeleteGateway(clientInfo.Context, testutils.GetEnvironmentID(), id).Execute()
	ok, err := common.HandleClientResponse(response, err, "DeleteGateway", resourceType)
	if err != nil {
		t.Errorf("Failed to delete test %s: %v", resourceType, err)
	}
	if !ok {
		t.Fatalf("Failed to create test %s: non-ok response", resourceType)
	}
}
