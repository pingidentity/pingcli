package resources

import (
	"fmt"

	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/common"
	"github.com/pingidentity/pingcli/internal/logger"
)

// Verify that the resource satisfies the exportable resource interface
var (
	_ connector.ExportableResource = &PingoneAuthorizeTrustFrameworkConditionResource{}
)

type PingoneAuthorizeTrustFrameworkConditionResource struct {
	clientInfo *connector.PingOneClientInfo
}

// Utility method for creating a PingoneAuthorizeTrustFrameworkConditionResource
func AuthorizeTrustFrameworkCondition(clientInfo *connector.PingOneClientInfo) *PingoneAuthorizeTrustFrameworkConditionResource {
	return &PingoneAuthorizeTrustFrameworkConditionResource{
		clientInfo: clientInfo,
	}
}

func (r *PingoneAuthorizeTrustFrameworkConditionResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()

	l.Debug().Msgf("Fetching all %s resources...", r.ResourceType())

	apiExecuteFunc := r.clientInfo.ApiClient.AuthorizeAPIClient.AuthorizeEditorConditionsApi.ListConditions(r.clientInfo.Context, r.clientInfo.ExportEnvironmentID).Execute
	apiFunctionName := "ListConditions"

	embedded, err := common.GetAuthorizeEmbedded(apiExecuteFunc, apiFunctionName, r.ResourceType())
	if err != nil {
		return nil, err
	}

	importBlocks := []connector.ImportBlock{}

	l.Debug().Msgf("Generating Import Blocks for all %s resources...", r.ResourceType())

	for _, authorizationCondition := range embedded.GetAuthorizationConditions() {
		authorizationConditionName, authorizationConditionNameOk := authorizationCondition.GetFullNameOk()
		authorizationConditionId, authorizationConditionIdOk := authorizationCondition.GetIdOk()

		if authorizationConditionNameOk && authorizationConditionIdOk {
			commentData := map[string]string{
				"Resource Type": r.ResourceType(),
				"Authorize Trust Framework Condition Name": *authorizationConditionName,
				"Export Environment ID":                    r.clientInfo.ExportEnvironmentID,
				"Authorize Trust Framework Condition ID":   *authorizationConditionId,
			}

			importBlocks = append(importBlocks, connector.ImportBlock{
				ResourceType:       r.ResourceType(),
				ResourceName:       *authorizationConditionName,
				ResourceID:         fmt.Sprintf("%s/%s", r.clientInfo.ExportEnvironmentID, *authorizationConditionId),
				CommentInformation: common.GenerateCommentInformation(commentData),
			})
		}
	}

	return &importBlocks, nil
}

func (r *PingoneAuthorizeTrustFrameworkConditionResource) ResourceType() string {
	return "pingone_authorize_trust_framework_condition"
}
