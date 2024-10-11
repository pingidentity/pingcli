package resources

import (
	"fmt"

	"github.com/pingidentity/pingctl/internal/connector"
	"github.com/pingidentity/pingctl/internal/connector/common"
	"github.com/pingidentity/pingctl/internal/logger"
)

// Verify that the resource satisfies the exportable resource interface
var (
	_ connector.ExportableResource = &PingoneAuthorizePolicyManagementStatementResource{}
)

type PingoneAuthorizePolicyManagementStatementResource struct {
	clientInfo *connector.PingOneClientInfo
}

// Utility method for creating a PingoneAuthorizePolicyManagementStatementResource
func AuthorizePolicyManagementStatement(clientInfo *connector.PingOneClientInfo) *PingoneAuthorizePolicyManagementStatementResource {
	return &PingoneAuthorizePolicyManagementStatementResource{
		clientInfo: clientInfo,
	}
}

func (r *PingoneAuthorizePolicyManagementStatementResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()

	l.Debug().Msgf("Fetching all %s resources...", r.ResourceType())

	apiExecuteFunc := r.clientInfo.ApiClient.AuthorizeAPIClient.AuthorizeEditorStatementsApi.ListStatements(r.clientInfo.Context, r.clientInfo.ExportEnvironmentID).Execute
	apiFunctionName := "ListStatements"

	embedded, err := common.GetAuthorizeEmbedded(apiExecuteFunc, apiFunctionName, r.ResourceType())
	if err != nil {
		return nil, err
	}

	importBlocks := []connector.ImportBlock{}

	l.Debug().Msgf("Generating Import Blocks for all %s resources...", r.ResourceType())

	for _, authorizationStatement := range embedded.GetAuthorizationStatements() {
		authorizationStatementName, authorizationStatementNameOk := authorizationStatement.GetNameOk()
		authorizationStatementId, authorizationStatementIdOk := authorizationStatement.GetIdOk()

		if authorizationStatementNameOk && authorizationStatementIdOk {
			commentData := map[string]string{
				"Resource Type": r.ResourceType(),
				"Authorize Policy Management Authorization Statement Name": *authorizationStatementName,
				"Export Environment ID":                                  r.clientInfo.ExportEnvironmentID,
				"Authorize Policy Management Authorization Statement ID": *authorizationStatementId,
			}

			importBlocks = append(importBlocks, connector.ImportBlock{
				ResourceType:       r.ResourceType(),
				ResourceName:       *authorizationStatementName,
				ResourceID:         fmt.Sprintf("%s/%s", r.clientInfo.ExportEnvironmentID, *authorizationStatementId),
				CommentInformation: common.GenerateCommentInformation(commentData),
			})
		}
	}

	return &importBlocks, nil
}

func (r *PingoneAuthorizePolicyManagementStatementResource) ResourceType() string {
	return "pingone_authorize_policy_management_statement"
}
