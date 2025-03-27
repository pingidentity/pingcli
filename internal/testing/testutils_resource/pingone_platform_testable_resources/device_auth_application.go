// Copyright © 2025 Ping Identity Corporation

package pingone_platform_testable_resources

import (
	"testing"

	"github.com/patrickcping/pingone-go-sdk-v2/management"
	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/common"
	"github.com/pingidentity/pingcli/internal/testing/testutils_resource"
)

func DeviceAuthApplication(t *testing.T, clientInfo *connector.ClientInfo) *testutils_resource.TestableResource {
	t.Helper()

	return &testutils_resource.TestableResource{
		ClientInfo:         clientInfo,
		CreateFunc:         createDeviceAuthApplication,
		DeleteFunc:         deleteDeviceAuthApplication,
		Dependencies:       nil,
		ExportableResource: nil,
	}
}

func createDeviceAuthApplication(t *testing.T, clientInfo *connector.ClientInfo, resourceType string, strArgs ...string) testutils_resource.ResourceInfo {
	t.Helper()

	if len(strArgs) != 0 {
		t.Errorf("Unexpected number of arguments provided to createPingOneDeviceAuthApplication(): %v", strArgs)

		return testutils_resource.ResourceInfo{}
	}

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
		t.Errorf("Failed to execute client function\nResponse Status: %s\nResponse Body: %s\nError: %v", response.Status, response.Body, err)

		return testutils_resource.ResourceInfo{}
	}
	if !ok {
		t.Errorf("Failed to execute client function\nResponse Status: %s\nResponse Body: %s", response.Status, response.Body)

		return testutils_resource.ResourceInfo{}
	}

	if createApplication201Response == nil || createApplication201Response.ApplicationOIDC == nil {
		t.Errorf("Failed to create test %s: %v", resourceType, err)

		return testutils_resource.ResourceInfo{}
	}

	appId, appIdOk := createApplication201Response.ApplicationOIDC.GetIdOk()
	if !appIdOk {
		t.Errorf("Failed to create test %s: %v", resourceType, err)

		return testutils_resource.ResourceInfo{}
	}

	return testutils_resource.ResourceInfo{
		DeletionIds: []string{
			*appId,
		},
		CreationInfo: map[testutils_resource.ResourceCreationInfoType]string{
			testutils_resource.ENUM_ID: *appId,
		},
	}
}

func deleteDeviceAuthApplication(t *testing.T, clientInfo *connector.ClientInfo, resourceType string, ids ...string) {
	t.Helper()

	if len(ids) != 1 {
		t.Errorf("Unexpected number of arguments provided to deleteDeviceAuthApplication(): %v", ids)

		return
	}

	response, err := clientInfo.PingOneApiClient.ManagementAPIClient.ApplicationsApi.DeleteApplication(clientInfo.PingOneContext, clientInfo.PingOneExportEnvironmentID, ids[0]).Execute()
	ok, err := common.HandleClientResponse(response, err, "DeleteApplication", resourceType)
	if err != nil {
		t.Errorf("Failed to execute client function\nResponse Status: %s\nResponse Body: %s\nError: %v", response.Status, response.Body, err)

		return
	}
	if !ok {
		t.Errorf("Failed to execute client function\nResponse Status: %s\nResponse Body: %s", response.Status, response.Body)

		return
	}
}
