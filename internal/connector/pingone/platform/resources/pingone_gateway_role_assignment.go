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
	clientInfo   *connector.PingOneClientInfo
	importBlocks *[]connector.ImportBlock
}

// Utility method for creating a PingOneGatewayRoleAssignmentResource
func GatewayRoleAssignment(clientInfo *connector.PingOneClientInfo) *PingOneGatewayRoleAssignmentResource {
	return &PingOneGatewayRoleAssignmentResource{
		clientInfo:   clientInfo,
		importBlocks: &[]connector.ImportBlock{},
	}
}

func (r *PingOneGatewayRoleAssignmentResource) ResourceType() string {
	return "pingone_gateway_role_assignment"
}

func (r *PingOneGatewayRoleAssignmentResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()
	l.Debug().Msgf("Exporting all '%s' Resources...", r.ResourceType())

	err := r.exportGatewayRoleAssignments()
	if err != nil {
		return nil, err
	}

	return r.importBlocks, nil
}

func (r *PingOneGatewayRoleAssignmentResource) exportGatewayRoleAssignments() error {
	iter := r.clientInfo.ApiClient.ManagementAPIClient.GatewaysApi.ReadAllGateways(r.clientInfo.Context, r.clientInfo.ExportEnvironmentID).Execute()

	for cursor, err := range iter {
		err = common.HandleClientResponse(cursor.HTTPResponse, err, "ReadAllGateways", r.ResourceType())
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

		for _, gatewayInner := range embedded.GetGateways() {
			// Only PingFederate Connections have role assignments
			if gatewayInner.Gateway != nil {
				gatewayType, gatewayTypeOk := gatewayInner.Gateway.GetTypeOk()

				if gatewayTypeOk && *gatewayType == management.ENUMGATEWAYTYPE_PING_FEDERATE {
					gatewayId, gatewayIdOk := gatewayInner.Gateway.GetIdOk()
					gatewayName, gatewayNameOk := gatewayInner.Gateway.GetNameOk()

					if gatewayIdOk && gatewayNameOk {
						err := r.exportGatewayRoleAssignmentsByGateway(*gatewayId, *gatewayName)
						if err != nil {
							return err
						}
					}
				}
			}
		}
	}

	return nil
}

func (r *PingOneGatewayRoleAssignmentResource) exportGatewayRoleAssignmentsByGateway(gatewayId, gatewayName string) error {
	iter := r.clientInfo.ApiClient.ManagementAPIClient.GatewayRoleAssignmentsApi.ReadGatewayRoleAssignments(r.clientInfo.Context, r.clientInfo.ExportEnvironmentID, gatewayId).Execute()

	for cursor, err := range iter {
		err = common.HandleClientResponse(cursor.HTTPResponse, err, "ReadGatewayRoleAssignments", r.ResourceType())
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

		for _, roleAssignment := range embedded.GetRoleAssignments() {
			roleAssignmentId, roleAssignmentIdOk := roleAssignment.GetIdOk()
			roleAssignmentRole, roleAssignmentRoleOk := roleAssignment.GetRoleOk()

			if roleAssignmentIdOk && roleAssignmentRoleOk {
				roleAssignmentRoleId, roleAssignmentRoleIdOk := roleAssignmentRole.GetIdOk()
				if roleAssignmentRoleIdOk {
					err := r.exportGatewayRoleAssignmentsByRole(gatewayId, gatewayName, *roleAssignmentId, *roleAssignmentRoleId)
					if err != nil {
						return err
					}
				}
			}
		}
	}

	return nil
}

func (r *PingOneGatewayRoleAssignmentResource) exportGatewayRoleAssignmentsByRole(gatewayId, gatewayName, roleAssignmentId, roleId string) error {
	role, resp, err := r.clientInfo.ApiClient.ManagementAPIClient.RolesApi.ReadOneRole(r.clientInfo.Context, roleId).Execute()
	err = common.HandleClientResponse(resp, err, "ReadOneRole", r.ResourceType())
	if err != nil {
		return err
	}

	if role != nil {
		roleName, roleNameOk := role.GetNameOk()
		if roleNameOk {
			r.addImportBlock(gatewayId, gatewayName, roleAssignmentId, string(*roleName))
		}
	}

	return nil
}

func (r *PingOneGatewayRoleAssignmentResource) addImportBlock(gatewayId, gatewayName, roleAssignmentId, roleName string) {
	commentData := map[string]string{
		"Export Environment ID": r.clientInfo.ExportEnvironmentID,
		"Gateway ID":            gatewayId,
		"Gateway Name":          gatewayName,
		"Resource Type":         r.ResourceType(),
		"Role Assignment ID":    roleAssignmentId,
		"Role Name":             roleName,
	}

	importBlock := connector.ImportBlock{
		ResourceType:       r.ResourceType(),
		ResourceName:       fmt.Sprintf("%s_%s_%s", gatewayName, roleName, roleAssignmentId),
		ResourceID:         fmt.Sprintf("%s/%s/%s", r.clientInfo.ExportEnvironmentID, gatewayId, roleAssignmentId),
		CommentInformation: common.GenerateCommentInformation(commentData),
	}

	*r.importBlocks = append(*r.importBlocks, importBlock)
}
