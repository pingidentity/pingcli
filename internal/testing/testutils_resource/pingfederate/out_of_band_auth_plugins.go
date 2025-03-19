// Copyright © 2025 Ping Identity Corporation

package pingfederate

import (
	"testing"

	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/common"
	"github.com/pingidentity/pingcli/internal/testing/testutils_resource"
	"github.com/pingidentity/pingcli/internal/testing/testutils_resource/pingone"
	"github.com/pingidentity/pingcli/internal/utils"
	client "github.com/pingidentity/pingfederate-go-client/v1220/configurationapi"
)

func TestableResource_PingFederateOutOfBandAuthPlugins(t *testing.T, clientInfo *connector.ClientInfo) *testutils_resource.TestableResource {
	t.Helper()

	return &testutils_resource.TestableResource{
		ClientInfo: clientInfo,
		CreateFunc: createOutOfBandAuthPlugins,
		DeleteFunc: deleteOutOfBandAuthPlugins,
		Dependencies: []*testutils_resource.TestableResource{
			TestableResource_PingFederatePingoneConnection(t, clientInfo),
			pingone.TestableResource_PingOneDeviceAuthApplication(t, clientInfo),
		},
		ExportableResource: nil,
	}
}

func createOutOfBandAuthPlugins(t *testing.T, clientInfo *connector.ClientInfo, strArgs ...string) testutils_resource.ResourceCreationInfo {
	t.Helper()

	if len(strArgs) != 3 {
		t.Fatalf("Unexpected number of arguments provided to createOutOfBandAuthPlugins(): %v", strArgs)
	}
	resourceType := strArgs[0]
	testPingOneConnectionId := strArgs[1]
	testDeviceAuthApplicationId := strArgs[2]

	request := clientInfo.PingFederateApiClient.OauthOutOfBandAuthPluginsAPI.CreateOOBAuthenticator(clientInfo.PingFederateContext)
	result := client.OutOfBandAuthenticator{
		AttributeContract: &client.OutOfBandAuthAttributeContract{
			CoreAttributes: []client.OutOfBandAuthAttribute{
				{
					Name: "subject",
				},
			},
		},
		Configuration: client.PluginConfiguration{
			Fields: []client.ConfigField{
				{
					Name:  "PingOne Environment",
					Value: utils.Pointer(testPingOneConnectionId + "|" + clientInfo.PingOneExportEnvironmentID),
				},
				{
					Name:  "Application",
					Value: &testDeviceAuthApplicationId,
				},
			},
		},
		Id:   "TestOOBAuthenticatorId",
		Name: "TestOOBAuthenticatorName",
		PluginDescriptorRef: client.ResourceLink{
			Id: "com.pingidentity.oobauth.pingone.mfa.PingOneMfaCibaAuthenticator",
		},
	}

	request = request.Body(result)

	resource, response, err := request.Execute()
	ok, err := common.HandleClientResponse(response, err, "CreateOOBAuthenticator", resourceType)
	if err != nil {
		t.Fatalf("Failed to execute client function\nResponse Status: %s\nResponse Body: %s\nError: %v", response.Status, response.Body, err)
	}
	if !ok {
		t.Fatalf("Failed to execute client function\nResponse Status: %s\nResponse Body: %s", response.Status, response.Body)
	}

	return testutils_resource.ResourceCreationInfo{
		testutils_resource.ENUM_ID: resource.Id,
	}
}

func deleteOutOfBandAuthPlugins(t *testing.T, clientInfo *connector.ClientInfo, resourceType, id string) {
	t.Helper()

	request := clientInfo.PingFederateApiClient.OauthOutOfBandAuthPluginsAPI.DeleteOOBAuthenticator(clientInfo.PingFederateContext, id)

	response, err := request.Execute()
	ok, err := common.HandleClientResponse(response, err, "DeleteOOBAuthenticator", resourceType)
	if err != nil {
		t.Fatalf("Failed to execute client function\nResponse Status: %s\nResponse Body: %s\nError: %v", response.Status, response.Body, err)
	}
	if !ok {
		t.Fatalf("Failed to execute client function\nResponse Status: %s\nResponse Body: %s", response.Status, response.Body)
	}
}
