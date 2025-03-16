package pingone

import (
	"testing"

	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/common"
	"github.com/pingidentity/pingcli/internal/testing/testutils_resource"
)

func TestableResource_PingOnePingFederateGatewayCredential(t *testing.T, clientInfo *connector.ClientInfo) *testutils_resource.TestableResource {
	t.Helper()

	return &testutils_resource.TestableResource{
		ClientInfo: clientInfo,
		CreateFunc: createPingFederateGatewayCredential,
		DeleteFunc: nil,
		Dependencies: []*testutils_resource.TestableResource{
			TestableResource_PingOnePingFederateGateway(t, clientInfo),
		},
		ExportableResource: nil,
	}
}

func createPingFederateGatewayCredential(t *testing.T, clientInfo *connector.ClientInfo, strArgs ...string) testutils_resource.ResourceCreationInfo {
	t.Helper()

	if len(strArgs) != 2 {
		t.Fatalf("Unexpected number of arguments provided to createPingoneConnection(): %v", strArgs)
	}
	resourceType := strArgs[0]
	gatewayId := strArgs[1]

	gatewayCredential, response, err := clientInfo.PingOneApiClient.ManagementAPIClient.GatewayCredentialsApi.CreateGatewayCredential(clientInfo.PingOneContext, clientInfo.PingOneExportEnvironmentID, gatewayId).Execute()
	ok, err := common.HandleClientResponse(response, err, "CreateGatewayCredential", resourceType)
	if err != nil {
		t.Fatalf("Failed to execute client function\nResponse Status: %s\nResponse Body: %s\nError: %v", response.Status, response.Body, err)
	}
	if !ok {
		t.Fatalf("Failed to execute client function\nResponse Status: %s\nResponse Body: %s", response.Status, response.Body)
	}

	if gatewayCredential == nil {
		t.Fatalf("Failed to create test %s: %v", resourceType, err)
	}

	credential, credentialOk := gatewayCredential.GetCredentialOk()
	if !credentialOk {
		t.Fatalf("Failed to create test %s: %v", resourceType, err)
	}

	return testutils_resource.ResourceCreationInfo{
		testutils_resource.ENUM_ID:         *credential, // the pingone_connection resource needs the credential instead of the id of the credential
		testutils_resource.ENUM_CREDENTIAL: *credential,
	}
}
