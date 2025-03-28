// Copyright © 2025 Ping Identity Corporation
// Code generated by ping-cli-generator

package pingone_platform_testable_resources

import (
	"testing"

	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/common"
	"github.com/pingidentity/pingcli/internal/connector/pingone/platform/resources"
	"github.com/pingidentity/pingcli/internal/testing/testutils_resource"
)

func GatewayCredential(t *testing.T, clientInfo *connector.ClientInfo) *testutils_resource.TestableResource {
	t.Helper()

	return &testutils_resource.TestableResource{
		ClientInfo: clientInfo,
		CreateFunc: createGatewayCredential,
		DeleteFunc: deleteGatewayCredential,
		Dependencies: []*testutils_resource.TestableResource{
			Gateway(t, clientInfo),
		},
		ExportableResource: resources.GatewayCredential(clientInfo),
	}
}

func createGatewayCredential(t *testing.T, clientInfo *connector.ClientInfo, resourceType string, strArgs ...string) testutils_resource.ResourceInfo {
	t.Helper()

	if len(strArgs) != 1 {
		t.Errorf("Unexpected number of arguments provided to createGatewayCredential(): %v", strArgs)

		return testutils_resource.ResourceInfo{}
	}
	gatewayId := strArgs[0]

	request := clientInfo.PingOneApiClient.ManagementAPIClient.GatewayCredentialsApi.CreateGatewayCredential(clientInfo.PingOneContext, clientInfo.PingOneExportEnvironmentID, gatewayId)

	resource, response, err := request.Execute()
	ok, err := common.HandleClientResponse(response, err, "CreateGatewayCredential", resourceType)
	if err != nil {
		t.Errorf("Failed to execute PingOne client function\nResponse Status: %s\nResponse Body: %s\nError: %v", response.Status, response.Body, err)

		return testutils_resource.ResourceInfo{}
	}
	if !ok {
		t.Errorf("Failed to execute PingOne client function\nResponse Status: %s\nResponse Body: %s", response.Status, response.Body)

		return testutils_resource.ResourceInfo{}
	}

	return testutils_resource.ResourceInfo{
		DeletionIds: []string{
			gatewayId,
			*resource.Id,
		},
		CreationInfo: map[testutils_resource.ResourceCreationInfoType]string{
			testutils_resource.ENUM_ID:         *resource.Credential,
			testutils_resource.ENUM_CREDENTIAL: *resource.Id,
		},
	}
}

func deleteGatewayCredential(t *testing.T, clientInfo *connector.ClientInfo, resourceType string, ids ...string) {
	t.Helper()

	if len(ids) != 2 {
		t.Errorf("Unexpected number of arguments provided to deleteGatewayCredential(): %v", ids)

		return
	}

	request := clientInfo.PingOneApiClient.ManagementAPIClient.GatewayCredentialsApi.DeleteGatewayCredential(clientInfo.PingOneContext, clientInfo.PingOneExportEnvironmentID, ids[0], ids[1])

	response, err := request.Execute()
	ok, err := common.HandleClientResponse(response, err, "DeleteGatewayCredential", resourceType)
	if err != nil {
		t.Errorf("Failed to execute PingOne client function\nResponse Status: %s\nResponse Body: %s\nError: %v", response.Status, response.Body, err)

		return
	}
	if !ok {
		t.Errorf("Failed to execute PingOne client function\nResponse Status: %s\nResponse Body: %s", response.Status, response.Body)

		return
	}
}
