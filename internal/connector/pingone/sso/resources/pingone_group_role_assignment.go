package resources

import (
	"fmt"

	"github.com/patrickcping/pingone-go-sdk-v2/management"
	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/common"
	"github.com/pingidentity/pingcli/internal/logger"
)

// Verify that the resource satisfies the exportable resource interface
var (
	_ connector.ExportableResource = &PingOneGroupRoleAssignmentResource{}
)

type PingOneGroupRoleAssignmentResource struct {
	clientInfo *connector.PingOneClientInfo
}

// Utility method for creating a PingOneGroupRoleAssignmentResource
func GroupRoleAssignment(clientInfo *connector.PingOneClientInfo) *PingOneGroupRoleAssignmentResource {
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

	for groupId, groupName := range *groupData {
		groupRoleAssignmentData, err := r.getGroupRoleAssignmentData(groupId)
		if err != nil {
			return nil, err
		}

		for groupRoleAssignmentId, roleId := range *groupRoleAssignmentData {
			roleName, err := r.getRoleName(roleId)
			if err != nil {
				return nil, err
			}

			commentData := map[string]string{
				"Export Environment ID":    r.clientInfo.ExportEnvironmentID,
				"Group ID":                 groupId,
				"Group Name":               groupName,
				"Group Role Assignment ID": groupRoleAssignmentId,
				"Group Role Name":          string(*roleName),
				"Resource Type":            r.ResourceType(),
			}

			importBlock := connector.ImportBlock{
				ResourceType:       r.ResourceType(),
				ResourceName:       fmt.Sprintf("%s_%s_%s", groupName, string(*roleName), groupRoleAssignmentId),
				ResourceID:         fmt.Sprintf("%s/%s/%s", r.clientInfo.ExportEnvironmentID, groupId, groupRoleAssignmentId),
				CommentInformation: common.GenerateCommentInformation(commentData),
			}

			importBlocks = append(importBlocks, importBlock)
		}
	}

	return &importBlocks, nil
}

func (r *PingOneGroupRoleAssignmentResource) getGroupData() (*map[string]string, error) {
	groupData := make(map[string]string)

	iter := r.clientInfo.ApiClient.ManagementAPIClient.GroupsApi.ReadAllGroups(r.clientInfo.Context, r.clientInfo.ExportEnvironmentID).Execute()

	for cursor, err := range iter {
		err = common.HandleClientResponse(cursor.HTTPResponse, err, "ReadAllGroups", r.ResourceType())
		if err != nil {
			return nil, err
		}

		if cursor.EntityArray == nil {
			return nil, common.DataNilError(r.ResourceType(), cursor.HTTPResponse)
		}

		embedded, embeddedOk := cursor.EntityArray.GetEmbeddedOk()
		if !embeddedOk {
			return nil, common.DataNilError(r.ResourceType(), cursor.HTTPResponse)
		}

		for _, group := range embedded.GetGroups() {
			groupId, groupIdOk := group.GetIdOk()
			groupName, groupNameOk := group.GetNameOk()

			if groupIdOk && groupNameOk {
				groupData[*groupId] = *groupName
			}
		}
	}

	return &groupData, nil
}

func (r *PingOneGroupRoleAssignmentResource) getGroupRoleAssignmentData(groupId string) (*map[string]string, error) {
	groupRoleAssignmentData := make(map[string]string)

	iter := r.clientInfo.ApiClient.ManagementAPIClient.GroupRoleAssignmentsApi.ReadGroupRoleAssignments(r.clientInfo.Context, r.clientInfo.ExportEnvironmentID, groupId).Execute()

	for cursor, err := range iter {
		err = common.HandleClientResponse(cursor.HTTPResponse, err, "ReadGroupRoleAssignments", r.ResourceType())
		if err != nil {
			return nil, err
		}

		if cursor.EntityArray == nil {
			return nil, common.DataNilError(r.ResourceType(), cursor.HTTPResponse)
		}

		embedded, embeddedOk := cursor.EntityArray.GetEmbeddedOk()
		if !embeddedOk {
			return nil, common.DataNilError(r.ResourceType(), cursor.HTTPResponse)
		}

		for _, groupRoleAssignment := range embedded.GetRoleAssignments() {
			groupRoleAssignmentId, groupRoleAssignmentIdOk := groupRoleAssignment.GetIdOk()
			groupRoleAssignmentRole, groupRoleAssignmentRoleOk := groupRoleAssignment.GetRoleOk()

			if groupRoleAssignmentIdOk && groupRoleAssignmentRoleOk {
				groupRoleAssignmentRoleId, groupRoleAssignmentRoleIdOk := groupRoleAssignmentRole.GetIdOk()

				if groupRoleAssignmentRoleIdOk {
					groupRoleAssignmentData[*groupRoleAssignmentId] = *groupRoleAssignmentRoleId
				}
			}
		}
	}

	return &groupRoleAssignmentData, nil
}

func (r *PingOneGroupRoleAssignmentResource) getRoleName(roleId string) (*management.EnumRoleName, error) {
	apiRole, resp, err := r.clientInfo.ApiClient.ManagementAPIClient.RolesApi.ReadOneRole(r.clientInfo.Context, roleId).Execute()
	err = common.HandleClientResponse(resp, err, "ReadOneRole", r.ResourceType())
	if err != nil {
		return nil, err
	}

	if apiRole != nil {
		apiRoleName, apiRoleNameOk := apiRole.GetNameOk()
		if apiRoleNameOk {
			return apiRoleName, nil
		}
	}

	return nil, fmt.Errorf("unable to get role name for role ID: %s", roleId)
}
