package resources

import (
	"fmt"

	"github.com/pingidentity/pingctl/internal/connector"
	"github.com/pingidentity/pingctl/internal/connector/common"
	"github.com/pingidentity/pingctl/internal/logger"
)

// Verify that the resource satisfies the exportable resource interface
var (
	_ connector.ExportableResource = &PingoneAuthorizeApplicationRoleResource{}
)

type PingoneAuthorizeApplicationRoleResource struct {
	clientInfo *connector.PingOneClientInfo
}

// Utility method for creating a PingoneAuthorizeApplicationRoleResource
func AuthorizeApplicationRole(clientInfo *connector.PingOneClientInfo) *PingoneAuthorizeApplicationRoleResource {
	return &PingoneAuthorizeApplicationRoleResource{
		clientInfo: clientInfo,
	}
}

func (r *PingoneAuthorizeApplicationRoleResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()

	l.Debug().Msgf("Fetching all %s resources...", r.ResourceType())

	apiExecuteFunc := r.clientInfo.ApiClient.AuthorizeAPIClient.ApplicationRolesApi.ReadApplicationRoles(r.clientInfo.Context, r.clientInfo.ExportEnvironmentID).Execute
	apiFunctionName := "ReadApplicationRoles"

	embedded, err := common.GetAuthorizeEmbedded(apiExecuteFunc, apiFunctionName, r.ResourceType())
	if err != nil {
		return nil, err
	}

	importBlocks := []connector.ImportBlock{}

	l.Debug().Msgf("Generating Import Blocks for all %s resources...", r.ResourceType())

	for _, appRole := range embedded.GetRoles() {
		appRoleName, appRoleNameOk := appRole.GetNameOk()
		appRoleId, appRoleIdOk := appRole.GetIdOk()

		if appRoleNameOk && appRoleIdOk {
			commentData := map[string]string{
				"Resource Type":                   r.ResourceType(),
				"Authorize Application Role Name": *appRoleName,
				"Export Environment ID":           r.clientInfo.ExportEnvironmentID,
				"Authorize Application Role ID":   *appRoleId,
			}

			importBlocks = append(importBlocks, connector.ImportBlock{
				ResourceType:       r.ResourceType(),
				ResourceName:       *appRoleName,
				ResourceID:         fmt.Sprintf("%s/%s", r.clientInfo.ExportEnvironmentID, *appRoleId),
				CommentInformation: common.GenerateCommentInformation(commentData),
			})
		}
	}

	return &importBlocks, nil
}

func (r *PingoneAuthorizeApplicationRoleResource) ResourceType() string {
	return "pingone_authorize_application_role"
}
