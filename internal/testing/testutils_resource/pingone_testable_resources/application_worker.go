// Copyright © 2025 Ping Identity Corporation
// Code generated by ping-cli-generator

package pingone_testable_resources

import (
	"testing"

	"github.com/hashicorp/go-uuid"
	"github.com/patrickcping/pingone-go-sdk-v2/management"
	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/common"
	"github.com/pingidentity/pingcli/internal/connector/pingone/sso/resources"
	"github.com/pingidentity/pingcli/internal/testing/testutils_resource"
	"github.com/pingidentity/pingcli/internal/utils"
)

func ApplicationWorker(t *testing.T, clientInfo *connector.ClientInfo) *testutils_resource.TestableResource {
	t.Helper()

	return &testutils_resource.TestableResource{
		ClientInfo:         clientInfo,
		CreateFunc:         createApplicationWorker,
		DeleteFunc:         deleteApplicationWorker,
		Dependencies:       nil,
		ExportableResource: resources.Application(clientInfo),
	}
}

func createApplicationWorker(t *testing.T, clientInfo *connector.ClientInfo, resourceType string, strArgs ...string) testutils_resource.ResourceInfo {
	t.Helper()

	if len(strArgs) != 0 {
		t.Errorf("Unexpected number of arguments provided to createApplicationWorker(): %v", strArgs)
		return testutils_resource.ResourceInfo{}
	}

	// Give unique name to application to avoid collisions in dependency creations
	applicationName, err := uuid.GenerateUUID()
	if err != nil {
		t.Errorf("Failed to generate UUID for application name: %v", err)
		return testutils_resource.ResourceInfo{}
	}

	request := clientInfo.PingOneApiClient.ManagementAPIClient.ApplicationsApi.CreateApplication(clientInfo.PingOneContext, clientInfo.PingOneExportEnvironmentID)
	clientStruct := management.CreateApplicationRequest{
		ApplicationOIDC: &management.ApplicationOIDC{
			Enabled:     true,
			Name:        applicationName,
			Description: utils.Pointer("Test Worker Application"),
			Type:        management.ENUMAPPLICATIONTYPE_WORKER,
			Protocol:    management.ENUMAPPLICATIONPROTOCOL_OPENID_CONNECT,
			GrantTypes: []management.EnumApplicationOIDCGrantType{
				management.ENUMAPPLICATIONOIDCGRANTTYPE_CLIENT_CREDENTIALS,
			},
			AssignActorRoles:        utils.Pointer(false),
			TokenEndpointAuthMethod: management.ENUMAPPLICATIONOIDCTOKENAUTHMETHOD_CLIENT_SECRET_BASIC,
		},
	}

	request = request.CreateApplicationRequest(clientStruct)

	resource, response, err := request.Execute()
	ok, err := common.HandleClientResponse(response, err, "CreateApplication", resourceType)
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
			*resource.ApplicationOIDC.Id,
		},
		CreationInfo: map[testutils_resource.ResourceCreationInfoType]string{
			testutils_resource.ENUM_ID:   *resource.ApplicationOIDC.Id,
			testutils_resource.ENUM_NAME: resource.ApplicationOIDC.Name,
		},
	}
}

func deleteApplicationWorker(t *testing.T, clientInfo *connector.ClientInfo, resourceType string, ids ...string) {
	t.Helper()

	if len(ids) != 1 {
		t.Errorf("Unexpected number of arguments provided to deleteApplicationWorker(): %v", ids)
		return
	}

	request := clientInfo.PingOneApiClient.ManagementAPIClient.ApplicationsApi.DeleteApplication(clientInfo.PingOneContext, clientInfo.PingOneExportEnvironmentID, ids[0])

	response, err := request.Execute()
	ok, err := common.HandleClientResponse(response, err, "DeleteApplication", resourceType)
	if err != nil {
		t.Errorf("Failed to execute PingOne client function\nResponse Status: %s\nResponse Body: %s\nError: %v", response.Status, response.Body, err)
		return
	}
	if !ok {
		t.Errorf("Failed to execute PingOne client function\nResponse Status: %s\nResponse Body: %s", response.Status, response.Body)
		return
	}
}
