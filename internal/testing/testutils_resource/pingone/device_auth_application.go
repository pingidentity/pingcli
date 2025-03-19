// Copyright © 2025 Ping Identity Corporation

package pingone

import (
	"testing"

	"github.com/patrickcping/pingone-go-sdk-v2/management"
	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/common"
	"github.com/pingidentity/pingcli/internal/testing/testutils_resource"
)

func TestableResource_PingOneDeviceAuthApplication(t *testing.T, clientInfo *connector.ClientInfo) *testutils_resource.TestableResource {
	t.Helper()

	return &testutils_resource.TestableResource{
		ClientInfo:         clientInfo,
		CreateFunc:         createDeviceAuthApplication,
		DeleteFunc:         deleteDeviceAuthApplication,
		Dependencies:       nil,
		ExportableResource: nil,
	}
}

func createDeviceAuthApplication(t *testing.T, clientInfo *connector.ClientInfo, strArgs ...string) testutils_resource.ResourceCreationInfo {
	t.Helper()

	if len(strArgs) != 1 {
		t.Fatalf("Unexpected number of arguments provided to createPingOneDeviceAuthApplication(): %v", strArgs)
	}
	resourceType := strArgs[0]

	result := management.CreateApplicationRequest{
		ApplicationOIDC: &management.ApplicationOIDC{
			Enabled: true,
			GrantTypes: []management.EnumApplicationOIDCGrantType{
				management.ENUMAPPLICATIONOIDCGRANTTYPE_DEVICE_CODE,
				management.ENUMAPPLICATIONOIDCGRANTTYPE_REFRESH_TOKEN,
			},
			Name:                    "TestDeviceAuthApplication",
			Protocol:                management.ENUMAPPLICATIONPROTOCOL_OPENID_CONNECT,
			TokenEndpointAuthMethod: management.ENUMAPPLICATIONOIDCTOKENAUTHMETHOD_NONE,
			Type:                    management.ENUMAPPLICATIONTYPE_CUSTOM_APP,
		},
	}

	createApplication201Response, response, err := clientInfo.PingOneApiClient.ManagementAPIClient.ApplicationsApi.CreateApplication(clientInfo.PingOneContext, clientInfo.PingOneExportEnvironmentID).CreateApplicationRequest(result).Execute()
	ok, err := common.HandleClientResponse(response, err, "CreateApplication", resourceType)
	if err != nil {
		t.Fatalf("Failed to execute client function\nResponse Status: %s\nResponse Body: %s\nError: %v", response.Status, response.Body, err)
	}
	if !ok {
		t.Fatalf("Failed to execute client function\nResponse Status: %s\nResponse Body: %s", response.Status, response.Body)
	}

	if createApplication201Response == nil || createApplication201Response.ApplicationOIDC == nil {
		t.Fatalf("Failed to create test %s: %v", resourceType, err)
	}

	appId, appIdOk := createApplication201Response.ApplicationOIDC.GetIdOk()
	if !appIdOk {
		t.Fatalf("Failed to create test %s: %v", resourceType, err)
	}

	return testutils_resource.ResourceCreationInfo{
		testutils_resource.ENUM_ID: *appId,
	}
}

func deleteDeviceAuthApplication(t *testing.T, clientInfo *connector.ClientInfo, resourceType, id string) {
	t.Helper()

	response, err := clientInfo.PingOneApiClient.ManagementAPIClient.ApplicationsApi.DeleteApplication(clientInfo.PingOneContext, clientInfo.PingOneExportEnvironmentID, id).Execute()
	ok, err := common.HandleClientResponse(response, err, "DeleteApplication", resourceType)
	if err != nil {
		t.Fatalf("Failed to execute client function\nResponse Status: %s\nResponse Body: %s\nError: %v", response.Status, response.Body, err)
	}
	if !ok {
		t.Fatalf("Failed to execute client function\nResponse Status: %s\nResponse Body: %s", response.Status, response.Body)
	}
}
