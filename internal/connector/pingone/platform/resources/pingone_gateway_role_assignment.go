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

func (r *PingOneGatewayRoleAssignmentResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()

	l.Debug().Msgf("Fetching all %s resources...", r.ResourceType())

	apiExecuteFunc := r.clientInfo.ApiClient.ManagementAPIClient.GatewaysApi.ReadAllGateways(r.clientInfo.Context, r.clientInfo.ExportEnvironmentID).Execute
	apiFunctionName := "ReadAllGateways"

	gatewaysEmbedded, err := common.GetManagementEmbedded(apiExecuteFunc, apiFunctionName, r.ResourceType())
	if err != nil {
		return nil, err
	}

	importBlocks := []connector.ImportBlock{}

	l.Debug().Msgf("Generating Import Blocks for all %s resources...", r.ResourceType())

	for _, gatewayInner := range gatewaysEmbedded.GetGateways() {
		// Only PingFederate Connections have role assignments
		if gatewayInner.Gateway != nil {
			gatewayType, gatewayTypeOk := gatewayInner.Gateway.GetTypeOk()
			if gatewayTypeOk && *gatewayType == management.ENUMGATEWAYTYPE_PING_FEDERATE {
				gatewayId, gatewayIdOk := gatewayInner.Gateway.GetIdOk()
				gatewayName, gatewayNameOk := gatewayInner.Gateway.GetNameOk()

				if gatewayIdOk && gatewayNameOk {
					apiExecuteFunc := r.clientInfo.ApiClient.ManagementAPIClient.GatewayRoleAssignmentsApi.ReadGatewayRoleAssignments(r.clientInfo.Context, r.clientInfo.ExportEnvironmentID, *gatewayId).Execute
					apiFunctionName := "ReadGatewayRoleAssignments"

					gatewayRoleAssignmentsEmbedded, err := common.GetManagementEmbedded(apiExecuteFunc, apiFunctionName, r.ResourceType())
					if err != nil {
						return nil, err
					}

					for _, roleAssignment := range gatewayRoleAssignmentsEmbedded.GetRoleAssignments() {
						roleAssignmentId, roleAssignmentIdOk := roleAssignment.GetIdOk()
						roleAssignmentRole, roleAssignmentRoleOk := roleAssignment.GetRoleOk()
						if roleAssignmentIdOk && roleAssignmentRoleOk {
							roleAssignmentRoleId, roleAssignmentRoleIdOk := roleAssignmentRole.GetIdOk()
							if roleAssignmentRoleIdOk {
								role, resp, err := r.clientInfo.ApiClient.ManagementAPIClient.RolesApi.ReadOneRole(r.clientInfo.Context, *roleAssignmentRoleId).Execute()
								err = common.HandleClientResponse(resp, err, "ReadOneRole", r.ResourceType())
								if err != nil {
									return nil, err
								}
								if role != nil {
									roleName, roleNameOk := role.GetNameOk()
									if roleNameOk {
										commentData := map[string]string{
											"Resource Type":         r.ResourceType(),
											"Gateway Name":          *gatewayName,
											"Role Name":             string(*roleName),
											"Export Environment ID": r.clientInfo.ExportEnvironmentID,
											"Gateway ID":            *gatewayId,
											"Role Assignment ID":    *roleAssignmentId,
										}

										importBlocks = append(importBlocks, connector.ImportBlock{
											ResourceType:       r.ResourceType(),
											ResourceName:       fmt.Sprintf("%s_%s_%s", *gatewayName, *roleName, *roleAssignmentId),
											ResourceID:         fmt.Sprintf("%s/%s/%s", r.clientInfo.ExportEnvironmentID, *gatewayId, *roleAssignmentId),
											CommentInformation: common.GenerateCommentInformation(commentData),
										})
									}
								}
							}
						}
					}
				}
			}
		}
	}

	return &importBlocks, nil
}

func (r *PingOneGatewayRoleAssignmentResource) ResourceType() string {
	return "pingone_gateway_role_assignment"
}
