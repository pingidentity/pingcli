// Copyright © 2025 Ping Identity Corporation

package resources

import (
	"fmt"

	"github.com/patrickcping/pingone-go-sdk-v2/authorize"
	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/common"
	"github.com/pingidentity/pingcli/internal/connector/pingone"
	"github.com/pingidentity/pingcli/internal/logger"
)

// Verify that the resource satisfies the exportable resource interface
var (
	_ connector.ExportableResource = &PingoneAuthorizeApplicationRoleResource{}
)

type PingoneAuthorizeApplicationRoleResource struct {
	clientInfo *connector.ClientInfo
}

// Utility method for creating a PingoneAuthorizeApplicationRoleResource
func AuthorizeApplicationRole(clientInfo *connector.ClientInfo) *PingoneAuthorizeApplicationRoleResource {
	return &PingoneAuthorizeApplicationRoleResource{
		clientInfo: clientInfo,
	}
}

func (r *PingoneAuthorizeApplicationRoleResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()
	l.Debug().Msgf("Exporting all '%s' Resources...", r.ResourceType())

	importBlocks := []connector.ImportBlock{}

	ApplicationRoleData, err := r.getApplicationRoleData()
	if err != nil {
		return nil, err
	}

	for applicationRoleId, applicationRoleName := range ApplicationRoleData {
		commentData := map[string]string{
			"Export Environment ID": r.clientInfo.PingOneExportEnvironmentID,
			"Application Role ID":   applicationRoleId,
			"Application Role Name": applicationRoleName,
			"Resource Type":         r.ResourceType(),
		}

		importBlock := connector.ImportBlock{
			ResourceType:       r.ResourceType(),
			ResourceName:       applicationRoleName,
			ResourceID:         fmt.Sprintf("%s/%s", r.clientInfo.PingOneExportEnvironmentID, applicationRoleId),
			CommentInformation: common.GenerateCommentInformation(commentData),
		}

		importBlocks = append(importBlocks, importBlock)
	}

	return &importBlocks, nil
}

func (r *PingoneAuthorizeApplicationRoleResource) getApplicationRoleData() (map[string]string, error) {
	applicationRoleData := make(map[string]string)

	iter := r.clientInfo.PingOneApiClient.AuthorizeAPIClient.ApplicationRolesApi.ReadApplicationRoles(r.clientInfo.PingOneContext, r.clientInfo.PingOneExportEnvironmentID).Execute()
	applicationRoles, err := pingone.GetAuthorizeAPIObjectsFromIterator[authorize.ApplicationRole](iter, "ReadApplicationRoles", "GetRoles", r.ResourceType())
	if err != nil {
		return nil, err
	}

	for _, applicationRole := range applicationRoles {
		applicationRoleId, applicationRoleIdOk := applicationRole.GetIdOk()
		applicationRoleName, applicationRoleNameOk := applicationRole.GetNameOk()

		if applicationRoleIdOk && applicationRoleNameOk {
			applicationRoleData[*applicationRoleId] = *applicationRoleName
		}
	}

	return applicationRoleData, nil
}

func (r *PingoneAuthorizeApplicationRoleResource) ResourceType() string {
	return "pingone_authorize_application_role"
}
