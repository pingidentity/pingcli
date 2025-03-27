// Copyright © 2025 Ping Identity Corporation
// Code generated by ping-cli-generator

package pingone_authorize_testable_resources

import (
	"testing"

	"github.com/patrickcping/pingone-go-sdk-v2/authorize"
	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/common"
	"github.com/pingidentity/pingcli/internal/connector/pingone/authorize/resources"
	"github.com/pingidentity/pingcli/internal/testing/testutils_resource"
)

func AuthorizeApplicationRolePermission(t *testing.T, clientInfo *connector.ClientInfo) *testutils_resource.TestableResource {
	t.Helper()

	return &testutils_resource.TestableResource{
		ClientInfo: clientInfo,
		CreateFunc: createAuthorizeApplicationRolePermission,
		DeleteFunc: deleteAuthorizeApplicationRolePermission,
		Dependencies: []*testutils_resource.TestableResource{
			AuthorizeApplicationRole(t, clientInfo),
			ApplicationResourcePermission(t, clientInfo),
		},
		ExportableResource: resources.AuthorizeApplicationRolePermission(clientInfo),
	}
}

func createAuthorizeApplicationRolePermission(t *testing.T, clientInfo *connector.ClientInfo, resourceType string, strArgs ...string) testutils_resource.ResourceInfo {
	t.Helper()

	if len(strArgs) != 4 {
		t.Errorf("Unexpected number of arguments provided to createAuthorizeApplicationRolePermission(): %v", strArgs)
		return testutils_resource.ResourceInfo{}
	}
	applicationRoleId := strArgs[0]
	applicationResourcePermissionId := strArgs[3]

	request := clientInfo.PingOneApiClient.AuthorizeAPIClient.ApplicationRolePermissionsApi.CreateApplicationRolePermission(clientInfo.PingOneContext, clientInfo.PingOneExportEnvironmentID, applicationRoleId)
	clientStruct := authorize.ApplicationRolePermission{
		Id: applicationResourcePermissionId,
	}

	request = request.ApplicationRolePermission(clientStruct)

	resource, response, err := request.Execute()
	ok, err := common.HandleClientResponse(response, err, "CreateApplicationRolePermission", resourceType)
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
			applicationRoleId,
			resource.Id,
		},
		CreationInfo: map[testutils_resource.ResourceCreationInfoType]string{
			testutils_resource.ENUM_ID:   resource.Id,
			testutils_resource.ENUM_NAME: *resource.Key,
		},
	}
}

func deleteAuthorizeApplicationRolePermission(t *testing.T, clientInfo *connector.ClientInfo, resourceType string, ids ...string) {
	t.Helper()

	if len(ids) != 2 {
		t.Errorf("Unexpected number of arguments provided to deleteAuthorizeApplicationRolePermission(): %v", ids)
		return
	}

	request := clientInfo.PingOneApiClient.AuthorizeAPIClient.ApplicationRolePermissionsApi.DeleteApplicationRolePermission(clientInfo.PingOneContext, clientInfo.PingOneExportEnvironmentID, ids[0], ids[1])

	response, err := request.Execute()
	ok, err := common.HandleClientResponse(response, err, "DeleteApplicationRolePermission", resourceType)
	if err != nil {
		t.Errorf("Failed to execute PingOne client function\nResponse Status: %s\nResponse Body: %s\nError: %v", response.Status, response.Body, err)
		return
	}
	if !ok {
		t.Errorf("Failed to execute PingOne client function\nResponse Status: %s\nResponse Body: %s", response.Status, response.Body)
		return
	}
}
