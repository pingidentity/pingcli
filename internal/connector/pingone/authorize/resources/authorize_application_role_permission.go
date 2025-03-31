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
	_ connector.ExportableResource = &PingoneAuthorizeApplicationRolePermissionResource{}
)

type PingoneAuthorizeApplicationRolePermissionResource struct {
	clientInfo *connector.ClientInfo
}

// Utility method for creating a PingoneAuthorizeApplicationRolePermissionResource
func AuthorizeApplicationRolePermission(clientInfo *connector.ClientInfo) *PingoneAuthorizeApplicationRolePermissionResource {
	return &PingoneAuthorizeApplicationRolePermissionResource{
		clientInfo: clientInfo,
	}
}

func (r *PingoneAuthorizeApplicationRolePermissionResource) ResourceType() string {
	return "pingone_authorize_application_role_permission"
}

func (r *PingoneAuthorizeApplicationRolePermissionResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()
	l.Debug().Msgf("Exporting all '%s' Resources...", r.ResourceType())

	importBlocks := []connector.ImportBlock{}

	applicationRoleData, err := r.getApplicationRoleData()
	if err != nil {
		return nil, err
	}

	for applicationRoleId, applicationRoleName := range applicationRoleData {
		appRolePermissionData, err := r.getApplicationRolePermissionData(applicationRoleId)
		if err != nil {
			return nil, err
		}

		for applicationRolePermissionId, applicationRolePermissionKey := range appRolePermissionData {
			commentData := map[string]string{
				"Application Role ID":             applicationRoleId,
				"Application Role Name":           applicationRoleName,
				"Application Role Permission ID":  applicationRolePermissionId,
				"Application Role Permission Key": applicationRolePermissionKey,
				"Export Environment ID":           r.clientInfo.PingOneExportEnvironmentID,
				"Resource Type":                   r.ResourceType(),
			}

			importBlock := connector.ImportBlock{
				ResourceType:       r.ResourceType(),
				ResourceName:       fmt.Sprintf("%s_%s", applicationRoleName, applicationRolePermissionKey),
				ResourceID:         fmt.Sprintf("%s/%s/%s", r.clientInfo.PingOneExportEnvironmentID, applicationRoleId, applicationRolePermissionId),
				CommentInformation: common.GenerateCommentInformation(commentData),
			}

			importBlocks = append(importBlocks, importBlock)
		}
	}

	return &importBlocks, nil
}

func (r *PingoneAuthorizeApplicationRolePermissionResource) getApplicationRoleData() (map[string]string, error) {
	applicationRoleData := make(map[string]string)

	iter := r.clientInfo.PingOneApiClient.AuthorizeAPIClient.ApplicationRolesApi.ReadApplicationRoles(r.clientInfo.PingOneContext, r.clientInfo.PingOneExportEnvironmentID).Execute()
	apiObjs, err := pingone.GetAuthorizeAPIObjectsFromIterator[authorize.ApplicationRole](iter, "ReadApplicationRoles", "GetRoles", r.ResourceType())
	if err != nil {
		return nil, err
	}

	for _, applicationRole := range apiObjs {
		applicationRoleId, applicationRoleIdOk := applicationRole.GetIdOk()
		applicationRoleName, applicationRoleNameOk := applicationRole.GetNameOk()

		if applicationRoleIdOk && applicationRoleNameOk {
			applicationRoleData[*applicationRoleId] = *applicationRoleName
		}
	}

	return applicationRoleData, nil
}

func (r *PingoneAuthorizeApplicationRolePermissionResource) getApplicationRolePermissionData(appRoleId string) (map[string]string, error) {
	applicationRolePermissionData := make(map[string]string)

	iter := r.clientInfo.PingOneApiClient.AuthorizeAPIClient.ApplicationRolePermissionsApi.ReadApplicationRolePermissions(r.clientInfo.PingOneContext, r.clientInfo.PingOneExportEnvironmentID, appRoleId).Execute()
	apiObjs, err := pingone.GetAuthorizeAPIObjectsFromIterator[authorize.EntityArrayEmbeddedPermissionsInner](iter, "ReadApplicationRolePermissions", "GetPermissions", r.ResourceType())
	if err != nil {
		return nil, err
	}

	for _, innerObj := range apiObjs {
		if innerObj.ApplicationRolePermission != nil {
			applicationRolePermissionId, applicationRolePermissionIdOk := innerObj.ApplicationRolePermission.GetIdOk()
			applicationRolePermissionKey, applicationRolePermissionKeyOk := innerObj.ApplicationRolePermission.GetKeyOk()

			if applicationRolePermissionIdOk && applicationRolePermissionKeyOk {
				applicationRolePermissionData[*applicationRolePermissionId] = *applicationRolePermissionKey
			}
		}
	}

	return applicationRolePermissionData, nil
}
