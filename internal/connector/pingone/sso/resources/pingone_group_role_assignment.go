package resources

import (
	"fmt"

	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/common"
	"github.com/pingidentity/pingcli/internal/logger"
)

// Verify that the resource satisfies the exportable resource interface
var (
	_ connector.ExportableResource = &PingoneGroupRoleAssignmentResource{}
)

type PingoneGroupRoleAssignmentResource struct {
	clientInfo *connector.PingOneClientInfo
}

// Utility method for creating a PingoneGroupRoleAssignmentResource
func GroupRoleAssignment(clientInfo *connector.PingOneClientInfo) *PingoneGroupRoleAssignmentResource {
	return &PingoneGroupRoleAssignmentResource{
		clientInfo: clientInfo,
	}
}

func (r *PingoneGroupRoleAssignmentResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()

	l.Debug().Msgf("Fetching all %s resources...", r.ResourceType())

	apiExecuteFunc := r.clientInfo.ApiClient.ManagementAPIClient.GroupsApi.ReadAllGroups(r.clientInfo.Context, r.clientInfo.ExportEnvironmentID).Execute
	apiFunctionName := "ReadAllGroups"

	embedded, err := common.GetManagementEmbedded(apiExecuteFunc, apiFunctionName, r.ResourceType())
	if err != nil {
		return nil, err
	}

	importBlocks := []connector.ImportBlock{}

	l.Debug().Msgf("Generating Import Blocks for all %s resources...", r.ResourceType())

	for _, group := range embedded.GetGroups() {
		groupId, groupIdOk := group.GetIdOk()
		groupName, groupNameOk := group.GetNameOk()

		if groupIdOk && groupNameOk {
			apiGroupRoleAssignmentExecuteFunc := r.clientInfo.ApiClient.ManagementAPIClient.GroupRoleAssignmentsApi.ReadGroupRoleAssignments(r.clientInfo.Context, r.clientInfo.ExportEnvironmentID, *groupId).Execute
			apiGroupRoleAssignmentFunctionName := "ReadGroupRoleAssignments"

			embeddedGroupRoleAssignments, err := common.GetManagementEmbedded(apiGroupRoleAssignmentExecuteFunc, apiGroupRoleAssignmentFunctionName, r.ResourceType())
			if err != nil {
				return nil, err
			}

			for groupRoleAssignmentIndex, groupRoleAssignment := range embeddedGroupRoleAssignments.GetRoleAssignments() {
				groupRoleAssignmentId, groupRoleAssignmentIdOk := groupRoleAssignment.GetIdOk()
				groupRoleAssignmentRole, groupRoleAssignmentRoleOk := groupRoleAssignment.GetRoleOk()

				var (
					groupRoleAssignmentRoleId   *string
					groupRoleAssignmentRoleIdOk bool
				)

				if groupRoleAssignmentRoleOk {
					groupRoleAssignmentRoleId, groupRoleAssignmentRoleIdOk = groupRoleAssignmentRole.GetIdOk()
				}

				if groupRoleAssignmentIdOk && groupRoleAssignmentRoleOk && groupRoleAssignmentRoleIdOk {
					role, response, err := r.clientInfo.ApiClient.ManagementAPIClient.RolesApi.ReadOneRole(r.clientInfo.Context, *groupRoleAssignmentRoleId).Execute()
					err = common.HandleClientResponse(response, err, "ReadOneRole", r.ResourceType())
					if err != nil {
						return nil, err
					}

					if role != nil {
						roleName, roleNameOk := role.GetNameOk()

						if roleNameOk {
							commentData := map[string]string{
								"Resource Type":            r.ResourceType(),
								"Group Name":               *groupName,
								"Role Name":                string(*roleName),
								"Export Environment ID":    r.clientInfo.ExportEnvironmentID,
								"Group ID":                 *groupId,
								"Group Role Assignment ID": *groupRoleAssignmentId,
							}

							importBlocks = append(importBlocks, connector.ImportBlock{
								ResourceType:       r.ResourceType(),
								ResourceName:       fmt.Sprintf("%s_%s_%d", *groupName, *roleName, (groupRoleAssignmentIndex + 1)),
								ResourceID:         fmt.Sprintf("%s/%s/%s", r.clientInfo.ExportEnvironmentID, *groupId, *groupRoleAssignmentId),
								CommentInformation: common.GenerateCommentInformation(commentData),
							})
						}
					}
				}
			}
		}
	}

	return &importBlocks, nil
}

func (r *PingoneGroupRoleAssignmentResource) ResourceType() string {
	return "pingone_group_role_assignment"
}
