// Copyright © 2025 Ping Identity Corporation
// Code generated by ping-cli-generator

package pingone_testable_resources

import (
	"testing"

	"github.com/patrickcping/pingone-go-sdk-v2/management"
	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/common"
	"github.com/pingidentity/pingcli/internal/connector/pingone"
	"github.com/pingidentity/pingcli/internal/connector/pingone/sso/resources"
	"github.com/pingidentity/pingcli/internal/testing/testutils_resource"
)

func GroupRoleAssignment(t *testing.T, clientInfo *connector.ClientInfo) *testutils_resource.TestableResource {
	t.Helper()

	return &testutils_resource.TestableResource{
		ClientInfo: clientInfo,
		CreateFunc: createGroupRoleAssignment,
		DeleteFunc: deleteGroupRoleAssignment,
		Dependencies: []*testutils_resource.TestableResource{
			Group(t, clientInfo),
		},
		ExportableResource: resources.GroupRoleAssignment(clientInfo),
	}
}

func createGroupRoleAssignment(t *testing.T, clientInfo *connector.ClientInfo, resourceType string, strArgs ...string) testutils_resource.ResourceInfo {
	t.Helper()

	if len(strArgs) != 1 {
		t.Fatalf("Unexpected number of arguments provided to createGroupRoleAssignment(): %v", strArgs)
	}
	groupId := strArgs[0]

	iter := clientInfo.PingOneApiClient.ManagementAPIClient.RolesApi.ReadAllRoles(clientInfo.PingOneContext).Execute()
	apiObjs, err := pingone.GetManagementAPIObjectsFromIterator[management.EntityArrayEmbeddedRolesInner](iter, "ReadAllRoles", "GetRoles", resourceType)
	if err != nil {
		t.Fatalf("Failed to execute PingOne client function\nError: %v", err)
	}
	if len(apiObjs) == 0 {
		t.Fatal("Failed to execute PingOne client function\n No built-in roles returned from ReadAllRoles()")
	}

	var (
		roleId   string
		roleName string
	)

	for _, role := range apiObjs {
		if role.Role != nil {
			if role.Role.Name != nil && *role.Role.Name == management.ENUMROLENAME_APPLICATION_OWNER {
				roleId = *role.Role.Id
				roleName = string(*role.Role.Name)
				break
			}
		}
	}

	request := clientInfo.PingOneApiClient.ManagementAPIClient.GroupRoleAssignmentsApi.CreateGroupRoleAssignment(clientInfo.PingOneContext, clientInfo.PingOneExportEnvironmentID, groupId)
	clientStruct := management.RoleAssignment{
		Role: management.RoleAssignmentRole{
			Id: roleId,
		},
		Scope: management.RoleAssignmentScope{
			Id:   clientInfo.PingOneExportEnvironmentID,
			Type: management.ENUMROLEASSIGNMENTSCOPETYPE_ENVIRONMENT,
		},
	}

	request = request.RoleAssignment(clientStruct)

	resource, response, err := request.Execute()
	ok, err := common.HandleClientResponse(response, err, "CreateGroupRoleAssignment", resourceType)
	if err != nil {
		t.Fatalf("Failed to execute PingOne client function\nResponse Status: %s\nResponse Body: %s\nError: %v", response.Status, response.Body, err)
	}
	if !ok {
		t.Fatalf("Failed to execute PingOne client function\nResponse Status: %s\nResponse Body: %s", response.Status, response.Body)
	}

	return testutils_resource.ResourceInfo{
		DeletionIds: []string{
			groupId,
			*resource.Id,
		},
		CreationInfo: map[testutils_resource.ResourceCreationInfoType]string{
			testutils_resource.ENUM_ID:   *resource.Id,
			testutils_resource.ENUM_NAME: roleName,
		},
	}
}

func deleteGroupRoleAssignment(t *testing.T, clientInfo *connector.ClientInfo, resourceType string, ids ...string) {
	t.Helper()

	if len(ids) != 2 {
		t.Fatalf("Unexpected number of arguments provided to deleteGroupRoleAssignment(): %v", ids)
	}

	request := clientInfo.PingOneApiClient.ManagementAPIClient.GroupRoleAssignmentsApi.DeleteGroupRoleAssignment(clientInfo.PingOneContext, clientInfo.PingOneExportEnvironmentID, ids[0], ids[1])

	response, err := request.Execute()
	ok, err := common.HandleClientResponse(response, err, "DeleteGroupRoleAssignment", resourceType)
	if err != nil {
		t.Fatalf("Failed to execute PingOne client function\nResponse Status: %s\nResponse Body: %s\nError: %v", response.Status, response.Body, err)
	}
	if !ok {
		t.Fatalf("Failed to execute PingOne client function\nResponse Status: %s\nResponse Body: %s", response.Status, response.Body)
	}
}
