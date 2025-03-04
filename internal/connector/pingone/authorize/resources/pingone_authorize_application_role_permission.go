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
	clientInfo *connector.PingOneClientInfo
}

// Utility method for creating a PingoneAuthorizeApplicationRolePermissionResource
func AuthorizeApplicationRolePermission(clientInfo *connector.PingOneClientInfo) *PingoneAuthorizeApplicationRolePermissionResource {
	return &PingoneAuthorizeApplicationRolePermissionResource{
		clientInfo: clientInfo,
	}
}

func (r *PingoneAuthorizeApplicationRolePermissionResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()
	l.Debug().Msgf("Exporting all '%s' Resources...", r.ResourceType())

	importBlocks := []connector.ImportBlock{}

	applicationRoleData, err := r.getApplicationRoleData()
	if err != nil {
		return nil, err
	}

	for appRoleId, appRoleName := range applicationRoleData {
		appRolePermissionData, err := r.getApplicationRolePermissionData(appRoleId)
		if err != nil {
			return nil, err
		}

		for appRolePermissionId, appRolePermissionKey := range appRolePermissionData {
			commentData := map[string]string{
				"Application Role ID":             appRoleId,
				"Application Role Name":           appRoleName,
				"Application Role Permission ID":  appRolePermissionId,
				"Application Role Permission Key": appRolePermissionKey,
				"Export Environment ID":           r.clientInfo.ExportEnvironmentID,
				"Resource Type":                   r.ResourceType(),
			}

			importBlock := connector.ImportBlock{
				ResourceType:       r.ResourceType(),
				ResourceName:       fmt.Sprintf("%s_%s", appRoleName, appRolePermissionKey),
				ResourceID:         fmt.Sprintf("%s/%s/%s", r.clientInfo.ExportEnvironmentID, appRoleId, appRolePermissionId),
				CommentInformation: common.GenerateCommentInformation(commentData),
			}

			importBlocks = append(importBlocks, importBlock)
		}
	}

	return &importBlocks, nil
}

func (r *PingoneAuthorizeApplicationRolePermissionResource) getApplicationRoleData() (map[string]string, error) {
	applicationRoleData := make(map[string]string)

	iter := r.clientInfo.ApiClient.AuthorizeAPIClient.ApplicationRolesApi.ReadApplicationRoles(r.clientInfo.Context, r.clientInfo.ExportEnvironmentID).Execute()
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

func (r *PingoneAuthorizeApplicationRolePermissionResource) getApplicationRolePermissionData(appRoleId string) (map[string]string, error) {
	applicationRolePermissionData := make(map[string]string)

	iter := r.clientInfo.ApiClient.AuthorizeAPIClient.ApplicationRolePermissionsApi.ReadApplicationRolePermissions(r.clientInfo.Context, r.clientInfo.ExportEnvironmentID, appRoleId).Execute()
	applicationRolePermissions, err := pingone.GetAuthorizeAPIObjectsFromIterator[authorize.EntityArrayEmbeddedPermissionsInner](iter, "ReadApplicationRolePermissions", "GetPermissions", r.ResourceType())
	if err != nil {
		return nil, err
	}

	for _, applicationRolePermission := range applicationRolePermissions {

		var (
			applicationRolePermissionId    *string
			applicationRolePermissionIdOk  bool
			applicationRolePermissionKey   *string
			applicationRolePermissionKeyOk bool
		)

		switch t := applicationRolePermission.GetActualInstance().(type) {
		case *authorize.ApplicationRolePermission:
			applicationRolePermissionId, applicationRolePermissionIdOk = t.GetIdOk()
			applicationRolePermissionKey, applicationRolePermissionKeyOk = t.GetKeyOk()
		default:
			continue
		}

		if applicationRolePermissionIdOk && applicationRolePermissionKeyOk {
			applicationRolePermissionData[*applicationRolePermissionId] = *applicationRolePermissionKey
		}
	}

	return applicationRolePermissionData, nil
}

func (r *PingoneAuthorizeApplicationRolePermissionResource) ResourceType() string {
	return "pingone_authorize_application_role_permission"
}
