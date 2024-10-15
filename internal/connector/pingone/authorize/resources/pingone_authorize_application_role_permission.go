package resources

import (
	"fmt"

	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/common"
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

	l.Debug().Msgf("Fetching all %s resources...", r.ResourceType())

	apiExecuteApplicationRoleFunc := r.clientInfo.ApiClient.AuthorizeAPIClient.ApplicationRolesApi.ReadApplicationRoles(r.clientInfo.Context, r.clientInfo.ExportEnvironmentID).Execute
	apiApplicationRoleFunctionName := "ReadApplicationRoles"

	embedded, err := common.GetAuthorizeEmbedded(apiExecuteApplicationRoleFunc, apiApplicationRoleFunctionName, r.ResourceType())
	if err != nil {
		return nil, err
	}

	importBlocks := []connector.ImportBlock{}

	l.Debug().Msgf("Generating Import Blocks for all %s resources...", r.ResourceType())

	for _, applicationRole := range embedded.GetRoles() {
		var (
			applicationRoleId     *string
			applicationRoleIdOk   bool
			applicationRoleName   *string
			applicationRoleNameOk bool
		)

		applicationRoleId, applicationRoleIdOk = applicationRole.GetIdOk()
		applicationRoleName, applicationRoleNameOk = applicationRole.GetNameOk()

		if applicationRoleIdOk && applicationRoleNameOk {
			apiExecutePermissionsFunc := r.clientInfo.ApiClient.AuthorizeAPIClient.ApplicationResourcePermissionsApi.ReadApplicationPermissions(r.clientInfo.Context, r.clientInfo.ExportEnvironmentID, *applicationRoleId).Execute
			apiPermissionsFunctionName := "ReadApplicationPermissions"

			permissionsEmbedded, err := common.GetAuthorizeEmbedded(apiExecutePermissionsFunc, apiPermissionsFunctionName, r.ResourceType())
			if err != nil {
				return nil, err
			}

			for _, applicationRolePermission := range permissionsEmbedded.GetPermissions() {
				if v := applicationRolePermission.ApplicationRolePermission; v != nil {
					applicationRolePermissionId, applicationRolePermissionIdOk := v.GetIdOk()

					if applicationRolePermissionIdOk {

						commentData := map[string]string{
							"Resource Type":                            r.ResourceType(),
							"Authorize Application Role Name":          *applicationRoleName,
							"Authorize Application Role ID":            *applicationRoleId,
							"Export Environment ID":                    r.clientInfo.ExportEnvironmentID,
							"Authorize Application Role Permission ID": *applicationRolePermissionId,
						}

						importBlocks = append(importBlocks, connector.ImportBlock{
							ResourceType:       r.ResourceType(),
							ResourceName:       fmt.Sprintf("%s_%s", *applicationRoleName, *applicationRolePermissionId),
							ResourceID:         fmt.Sprintf("%s/%s/%s", r.clientInfo.ExportEnvironmentID, *applicationRoleId, *applicationRolePermissionId),
							CommentInformation: common.GenerateCommentInformation(commentData),
						})
					}
				}
			}
		}
	}

	return &importBlocks, nil
}

func (r *PingoneAuthorizeApplicationRolePermissionResource) ResourceType() string {
	return "pingone_authorize_api_service_operation"
}
