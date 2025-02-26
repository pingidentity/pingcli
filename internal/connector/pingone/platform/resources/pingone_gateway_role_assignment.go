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

	gatewayData, err := r.getGatewayData()
	if err != nil {
		return nil, err
	}

	for gatewayId, gatewayName := range gatewayData {
		gatewayRoleAssignmentData, err := r.getGatewayRoleAssignmentData(gatewayId)
		if err != nil {
			return nil, err
		}

		for roleAssignmentId, roleId := range gatewayRoleAssignmentData {
			roleName, err := r.getRoleAssignmentRoleName(roleId)
			if err != nil {
				return nil, err
			}
			if roleName == nil {
				continue
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

func (r *PingOneGatewayRoleAssignmentResource) getGatewayData() (map[string]string, error) {
	gatewayData := make(map[string]string)

	iter := r.clientInfo.ApiClient.ManagementAPIClient.GatewaysApi.ReadAllGateways(r.clientInfo.Context, r.clientInfo.ExportEnvironmentID).Execute()
	gatewayInners, err := pingone.GetManagementAPIObjectsFromIterator[management.EntityArrayEmbeddedGatewaysInner](iter, "ReadAllGateways", "GetGateways", r.ResourceType())
	if err != nil {
		return nil, err
	}

	for _, gatewayInner := range gatewayInners {
		// Only PingFederate Connections have role assignments
		if gatewayInner.Gateway != nil {
			gatewayType, gatewayTypeOk := gatewayInner.Gateway.GetTypeOk()

			if gatewayTypeOk && *gatewayType == management.ENUMGATEWAYTYPE_PING_FEDERATE {
				gatewayId, gatewayIdOk := gatewayInner.Gateway.GetIdOk()
				gatewayName, gatewayNameOk := gatewayInner.Gateway.GetNameOk()

				if gatewayIdOk && gatewayNameOk {
					gatewayData[*gatewayId] = *gatewayName
				}
			}
		}
	}

	return gatewayData, nil
}

func (r *PingOneGatewayRoleAssignmentResource) getGatewayRoleAssignmentData(gatewayId string) (map[string]string, error) {
	gatewayRoleAssignmentData := make(map[string]string)

	iter := r.clientInfo.ApiClient.ManagementAPIClient.GatewayRoleAssignmentsApi.ReadGatewayRoleAssignments(r.clientInfo.Context, r.clientInfo.ExportEnvironmentID, gatewayId).Execute()
	gatewayRoleAssignments, err := pingone.GetManagementAPIObjectsFromIterator[management.RoleAssignment](iter, "ReadGatewayRoleAssignments", "GetRoleAssignments", r.ResourceType())
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

func (r *PingOneGatewayRoleAssignmentResource) getRoleAssignmentRoleName(roleId string) (*management.EnumRoleName, error) {
	role, resp, err := r.clientInfo.ApiClient.ManagementAPIClient.RolesApi.ReadOneRole(r.clientInfo.Context, roleId).Execute()
	ok, err := pingone.CheckSingletonResource(resp, err, "ReadOneRole", r.ResourceType())
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

	return nil, fmt.Errorf("failed to export resource '%s'. No role name found for Role ID '%s'", r.ResourceType(), roleId)
}
