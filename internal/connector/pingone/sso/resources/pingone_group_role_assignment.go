package resources

import (
	"fmt"

	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/common"
	"github.com/pingidentity/pingcli/internal/logger"
)

// Verify that the resource satisfies the exportable resource interface
var (
	_ connector.ExportableResource = &PingOneGroupRoleAssignmentResource{}
)

type PingOneGroupRoleAssignmentResource struct {
	clientInfo   *connector.PingOneClientInfo
	importBlocks *[]connector.ImportBlock
}

// Utility method for creating a PingOneGroupRoleAssignmentResource
func GroupRoleAssignment(clientInfo *connector.PingOneClientInfo) *PingOneGroupRoleAssignmentResource {
	return &PingOneGroupRoleAssignmentResource{
		clientInfo:   clientInfo,
		importBlocks: &[]connector.ImportBlock{},
	}
}

func (r *PingOneGroupRoleAssignmentResource) ResourceType() string {
	return "pingone_group_role_assignment"
}

func (r *PingOneGroupRoleAssignmentResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()
	l.Debug().Msgf("Exporting all '%s' Resources...", r.ResourceType())

	err := r.exportGroupRoleAssignments()
	if err != nil {
		return nil, err
	}

	return r.importBlocks, nil
}

func (r *PingOneGroupRoleAssignmentResource) exportGroupRoleAssignments() error {
	iter := r.clientInfo.ApiClient.ManagementAPIClient.GroupsApi.ReadAllGroups(r.clientInfo.Context, r.clientInfo.ExportEnvironmentID).Execute()

	for cursor, err := range iter {
		err = common.HandleClientResponse(cursor.HTTPResponse, err, "ReadAllGroups", r.ResourceType())
		if err != nil {
			return err
		}

		if cursor.EntityArray == nil {
			return common.DataNilError(r.ResourceType(), cursor.HTTPResponse)
		}

		embedded, embeddedOk := cursor.EntityArray.GetEmbeddedOk()
		if !embeddedOk {
			return common.DataNilError(r.ResourceType(), cursor.HTTPResponse)
		}

		for _, group := range embedded.GetGroups() {
			groupId, groupIdOk := group.GetIdOk()
			groupName, groupNameOk := group.GetNameOk()

			if groupIdOk && groupNameOk {
				err := r.exportGroupRoleAssignmentsByGroup(*groupId, *groupName)
				if err != nil {
					return err
				}
			}
		}
	}

	return nil
}

func (r *PingOneGroupRoleAssignmentResource) exportGroupRoleAssignmentsByGroup(groupId string, groupName string) error {
	iter := r.clientInfo.ApiClient.ManagementAPIClient.GroupRoleAssignmentsApi.ReadGroupRoleAssignments(r.clientInfo.Context, r.clientInfo.ExportEnvironmentID, groupId).Execute()

	for cursor, err := range iter {
		err = common.HandleClientResponse(cursor.HTTPResponse, err, "ReadGroupRoleAssignments", r.ResourceType())
		if err != nil {
			return err
		}

		if cursor.EntityArray == nil {
			return common.DataNilError(r.ResourceType(), cursor.HTTPResponse)
		}

		embedded, embeddedOk := cursor.EntityArray.GetEmbeddedOk()
		if !embeddedOk {
			return common.DataNilError(r.ResourceType(), cursor.HTTPResponse)
		}

		for _, groupRoleAssignment := range embedded.GetRoleAssignments() {
			groupRoleAssignmentId, groupRoleAssignmentIdOk := groupRoleAssignment.GetIdOk()
			groupRoleAssignmentRole, groupRoleAssignmentRoleOk := groupRoleAssignment.GetRoleOk()

			if groupRoleAssignmentIdOk && groupRoleAssignmentRoleOk {
				groupRoleAssignmentRoleId, groupRoleAssignmentRoleIdOk := groupRoleAssignmentRole.GetIdOk()
				if groupRoleAssignmentRoleIdOk {
					err := r.exportGroupRoleAssignmentsByRole(groupId, groupName, *groupRoleAssignmentId, *groupRoleAssignmentRoleId)
					if err != nil {
						return err
					}
				}
			}
		}
	}

	return nil
}

func (r *PingOneGroupRoleAssignmentResource) exportGroupRoleAssignmentsByRole(groupId, groupName, groupRoleAssignmentId, roleId string) error {
	apiRole, resp, err := r.clientInfo.ApiClient.ManagementAPIClient.RolesApi.ReadOneRole(r.clientInfo.Context, roleId).Execute()
	err = common.HandleClientResponse(resp, err, "ReadOneRole", r.ResourceType())
	if err != nil {
		return err
	}

	if apiRole != nil {
		apiRoleName, apiRoleNameOk := apiRole.GetNameOk()
		if apiRoleNameOk {
			r.addImportBlock(groupId, groupName, groupRoleAssignmentId, string(*apiRoleName))
		}
	}

	return nil
}

func (r *PingOneGroupRoleAssignmentResource) addImportBlock(groupId, groupName, groupRoleAssignmentId, roleName string) {
	commentData := map[string]string{
		"Export Environment ID":    r.clientInfo.ExportEnvironmentID,
		"Group ID":                 groupId,
		"Group Name":               groupName,
		"Group Role Assignment ID": groupRoleAssignmentId,
		"Group Role Name":          roleName,
		"Resource Type":            r.ResourceType(),
	}

	importBlock := connector.ImportBlock{
		ResourceType:       r.ResourceType(),
		ResourceName:       fmt.Sprintf("%s_%s_%s", groupName, roleName, groupRoleAssignmentId),
		ResourceID:         fmt.Sprintf("%s/%s/%s", r.clientInfo.ExportEnvironmentID, groupId, groupRoleAssignmentId),
		CommentInformation: common.GenerateCommentInformation(commentData),
	}

	*r.importBlocks = append(*r.importBlocks, importBlock)
}
