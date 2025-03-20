// Copyright © 2025 Ping Identity Corporation
// Code generated by ping-cli-generator

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
	clientInfo *connector.ClientInfo
}

// Utility method for creating a PingOneGatewayRoleAssignmentResource
func GatewayRoleAssignment(clientInfo *connector.ClientInfo) *PingOneGatewayRoleAssignmentResource {
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

	gatewayRoleAssignmentData, err := r.getGatewayRoleAssignmentData()
	if err != nil {
		return nil, err
	}

	for gatewayRoleAssignmentId, gatewayRoleAssignmentName := range gatewayRoleAssignmentData {
		commentData := map[string]string{
			"Gateway Role Assignment ID":   gatewayRoleAssignmentId,
			"Gateway Role Assignment Name": gatewayRoleAssignmentName,
			"Export Environment ID":        r.clientInfo.PingOneExportEnvironmentID,
			"Resource Type":                r.ResourceType(),
		}

		importBlock := connector.ImportBlock{
			ResourceType:       r.ResourceType(),
			ResourceName:       gatewayRoleAssignmentName,
			ResourceID:         fmt.Sprintf("%s/%s", r.clientInfo.PingOneExportEnvironmentID, gatewayRoleAssignmentId),
			CommentInformation: common.GenerateCommentInformation(commentData),
		}

		importBlocks = append(importBlocks, importBlock)
	}

	return &importBlocks, nil
}

func (r *PingOneGatewayRoleAssignmentResource) getGatewayData() (map[string]string, error) {
	//TODO
}

func (r *PingOneGatewayRoleAssignmentResource) getGatewayRoleAssignmentData() (map[string]string, error) {
	gatewayRoleAssignmentData := make(map[string]string)

	iter := r.clientInfo.PingOneApiClient.ManagementAPIClient.GatewayRoleAssignmentsApi.ReadGatewayRoleAssignments(r.clientInfo.PingOneContext, r.clientInfo.PingOneExportEnvironmentID).Execute()
	apiObjs, err := pingone.GetManagementAPIObjectsFromIterator[management.RoleAssignment](iter, "ReadGatewayRoleAssignments", "GetRoleAssignments", r.ResourceType())
	if err != nil {
		return nil, err
	}

	for _, gatewayRoleAssignment := range apiObjs {
		gatewayRoleAssignmentId, gatewayRoleAssignmentIdOk := gatewayRoleAssignment.GetIdOk()
		gatewayRoleAssignmentName, gatewayRoleAssignmentNameOk := gatewayRoleAssignment.GetNameOk()

		if gatewayRoleAssignmentIdOk && gatewayRoleAssignmentNameOk {
			gatewayRoleAssignmentData[*gatewayRoleAssignmentId] = *gatewayRoleAssignmentName
		}
	}

	return gatewayRoleAssignmentData, nil
}
