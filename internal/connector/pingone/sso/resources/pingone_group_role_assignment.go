package resources

import (
	"fmt"

	"github.com/patrickcping/pingone-go-sdk-v2/management"
	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/common"
	"github.com/pingidentity/pingcli/internal/connector/pingone"
	"github.com/pingidentity/pingcli/internal/logger"
)

// Verify that the resource satisfies the exportable resource interface
var (
	_ connector.ExportableResource = &PingOneGroupRoleAssignmentResource{}
)

type PingOneGroupRoleAssignmentResource struct {
	clientInfo *connector.ClientInfo
}

// Utility method for creating a PingOneGroupRoleAssignmentResource
func GroupRoleAssignment(clientInfo *connector.ClientInfo) *PingOneGroupRoleAssignmentResource {
	return &PingOneGroupRoleAssignmentResource{
		clientInfo: clientInfo,
	}
}

func (r *PingOneGroupRoleAssignmentResource) ResourceType() string {
	return "pingone_group_role_assignment"
}

func (r *PingOneGroupRoleAssignmentResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()
	l.Debug().Msgf("Exporting all '%s' Resources...", r.ResourceType())

	importBlocks := []connector.ImportBlock{}

	groupData, err := r.getGroupData()
	if err != nil {
		return nil, err
	}

	for groupId, groupName := range groupData {
		groupRoleAssignmentData, err := r.getGroupRoleAssignmentData(groupId)
		if err != nil {
			return nil, err
		}

		for groupRoleAssignmentId, roleId := range groupRoleAssignmentData {
			roleName, err := r.getRoleName(roleId)
			if err != nil {
				return nil, err
			}
			if roleName == nil {
				continue
			}

			commentData := map[string]string{
				"Export Environment ID":    r.clientInfo.PingOneExportEnvironmentID,
				"Group ID":                 groupId,
				"Group Name":               groupName,
				"Group Role Assignment ID": groupRoleAssignmentId,
				"Group Role Name":          string(*roleName),
				"Resource Type":            r.ResourceType(),
			}

			importBlock := connector.ImportBlock{
				ResourceType:       r.ResourceType(),
				ResourceName:       fmt.Sprintf("%s_%s_%s", groupName, string(*roleName), groupRoleAssignmentId),
				ResourceID:         fmt.Sprintf("%s/%s/%s", r.clientInfo.PingOneExportEnvironmentID, groupId, groupRoleAssignmentId),
				CommentInformation: common.GenerateCommentInformation(commentData),
			}

			importBlocks = append(importBlocks, importBlock)
		}
	}

	return &importBlocks, nil
}

func (r *PingOneGroupRoleAssignmentResource) getGroupData() (map[string]string, error) {
	groupData := make(map[string]string)

	iter := r.clientInfo.PingOneApiClient.ManagementAPIClient.GroupsApi.ReadAllGroups(r.clientInfo.PingOneContext, r.clientInfo.PingOneExportEnvironmentID).Execute()
	groups, err := pingone.GetManagementAPIObjectsFromIterator[management.Group](iter, "ReadAllGroups", "GetGroups", r.ResourceType())
	if err != nil {
		return nil, err
	}

	for _, group := range groups {
		groupId, groupIdOk := group.GetIdOk()
		groupName, groupNameOk := group.GetNameOk()

		if groupIdOk && groupNameOk {
			groupData[*groupId] = *groupName
		}
	}

	return groupData, nil
}

func (r *PingOneGroupRoleAssignmentResource) getGroupRoleAssignmentData(groupId string) (map[string]string, error) {
	groupRoleAssignmentData := make(map[string]string)

	iter := r.clientInfo.PingOneApiClient.ManagementAPIClient.GroupRoleAssignmentsApi.ReadGroupRoleAssignments(r.clientInfo.PingOneContext, r.clientInfo.PingOneExportEnvironmentID, groupId).Execute()
	roleAssignments, err := pingone.GetManagementAPIObjectsFromIterator[management.RoleAssignment](iter, "ReadGroupRoleAssignments", "GetRoleAssignments", r.ResourceType())
	if err != nil {
		return nil, err
	}

	for _, groupRoleAssignment := range roleAssignments {
		groupRoleAssignmentId, groupRoleAssignmentIdOk := groupRoleAssignment.GetIdOk()
		groupRoleAssignmentRole, groupRoleAssignmentRoleOk := groupRoleAssignment.GetRoleOk()

		if groupRoleAssignmentIdOk && groupRoleAssignmentRoleOk {
			groupRoleAssignmentRoleId, groupRoleAssignmentRoleIdOk := groupRoleAssignmentRole.GetIdOk()

			if groupRoleAssignmentRoleIdOk {
				groupRoleAssignmentData[*groupRoleAssignmentId] = *groupRoleAssignmentRoleId
			}
		}
	}

	return groupRoleAssignmentData, nil
}

func (r *PingOneGroupRoleAssignmentResource) getRoleName(roleId string) (*management.EnumRoleName, error) {
	apiRole, resp, err := r.clientInfo.PingOneApiClient.ManagementAPIClient.RolesApi.ReadOneRole(r.clientInfo.PingOneContext, roleId).Execute()
	ok, err := common.HandleClientResponse(resp, err, "ReadOneRole", r.ResourceType())
	if err != nil {
		return nil, err
	}
	if !ok {
		return nil, nil
	}

	if apiRole != nil {
		apiRoleName, apiRoleNameOk := apiRole.GetNameOk()
		if apiRoleNameOk {
			return apiRoleName, nil
		}
	}

	return nil, fmt.Errorf("unable to get role name for role ID: %s", roleId)
}
