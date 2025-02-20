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
	_ connector.ExportableResource = &PingOneGatewayRoleAssignmentResource{}
)

type PingOneGatewayRoleAssignmentResource struct {
	clientInfo *connector.PingOneClientInfo
}

// Utility method for creating a PingOneGatewayRoleAssignmentResource
func GatewayRoleAssignment(clientInfo *connector.PingOneClientInfo) *PingOneGatewayRoleAssignmentResource {
	return &PingOneGatewayRoleAssignmentResource{
		clientInfo: clientInfo,
	}
}

func (r *PingOneGatewayRoleAssignmentResource) ResourceType() string {
	return "pingone_gateway_role_assignment"
}

func (r *PingOneGatewayRoleAssignmentResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()
	l.Debug().Msgf("Exporting all '%s' Resources...", r.ResourceType())

	importBlocks := []connector.ImportBlock{}

	gatewayData, err := getGatewayData(r.clientInfo, r.ResourceType())
	if err != nil {
		return nil, err
	}

	//TODO: Only PingFederate Connections have role assignments

	for gatewayId, gatewayName := range gatewayData {
		gatewayRoleAssignmentData, err := getGatewayRoleAssignmentData(r.clientInfo, r.ResourceType(), gatewayId)
		if err != nil {
			return nil, err
		}

		for roleAssignmentId, roleId := range gatewayRoleAssignmentData {
			roleName, err := getRoleAssignmentRoleName(r.clientInfo, r.ResourceType(), roleId)
			if err != nil {
				return nil, err
			}

			commentData := map[string]string{
				"Export Environment ID": r.clientInfo.ExportEnvironmentID,
				"Gateway ID":            gatewayId,
				"Gateway Name":          gatewayName,
				"Resource Type":         r.ResourceType(),
				"Role Assignment ID":    roleAssignmentId,
				"Role Name":             string(*roleName),
			}

			importBlock := connector.ImportBlock{
				ResourceType:       r.ResourceType(),
				ResourceName:       fmt.Sprintf("%s_%s_%s", gatewayName, string(*roleName), roleAssignmentId),
				ResourceID:         fmt.Sprintf("%s/%s/%s", r.clientInfo.ExportEnvironmentID, gatewayId, roleAssignmentId),
				CommentInformation: common.GenerateCommentInformation(commentData),
			}

			importBlocks = append(importBlocks, importBlock)
		}
	}

	return &importBlocks, nil
}

func getGatewayRoleAssignmentData(clientInfo *connector.PingOneClientInfo, resourceType, gatewayId string) (map[string]string, error) {
	gatewayRoleAssignmentData := make(map[string]string)

	iter := clientInfo.ApiClient.ManagementAPIClient.GatewayRoleAssignmentsApi.ReadGatewayRoleAssignments(clientInfo.Context, clientInfo.ExportEnvironmentID, gatewayId).Execute()
	gatewayRoleAssignments, err := common.GetManagementAPIObjectsFromIterator[management.RoleAssignment](iter, "ReadGatewayRoleAssignments", "GetRoleAssignments", resourceType)
	if err != nil {
		return nil, err
	}

	for _, roleAssignment := range gatewayRoleAssignments {
		roleAssignmentId, roleAssignmentIdOk := roleAssignment.GetIdOk()
		roleAssignmentRole, roleAssignmentRoleOk := roleAssignment.GetRoleOk()

		if roleAssignmentIdOk && roleAssignmentRoleOk {
			roleAssignmentRoleId, roleAssignmentRoleIdOk := roleAssignmentRole.GetIdOk()
			if roleAssignmentRoleIdOk {
				gatewayRoleAssignmentData[*roleAssignmentId] = *roleAssignmentRoleId
			}
		}
	}

	return gatewayRoleAssignmentData, nil
}

func getRoleAssignmentRoleName(clientInfo *connector.PingOneClientInfo, resourceType, roleId string) (*management.EnumRoleName, error) {
	role, resp, err := clientInfo.ApiClient.ManagementAPIClient.RolesApi.ReadOneRole(clientInfo.Context, roleId).Execute()
	ok, err := common.CheckSingletonResource(resp, err, "ReadOneRole", resourceType)
	if err != nil {
		return nil, err
	}
	if !ok {
		return nil, nil
	}

	if role != nil {
		roleName, roleNameOk := role.GetNameOk()
		if roleNameOk {
			return roleName, nil
		}
	}

	return nil, fmt.Errorf("failed to export resource '%s'. No role name found for Role ID '%s'.", resourceType, roleId)
}
