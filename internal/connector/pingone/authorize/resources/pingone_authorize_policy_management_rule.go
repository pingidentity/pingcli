package resources

import (
	"fmt"

	"github.com/pingidentity/pingctl/internal/connector"
	"github.com/pingidentity/pingctl/internal/connector/common"
	"github.com/pingidentity/pingctl/internal/logger"
)

// Verify that the resource satisfies the exportable resource interface
var (
	_ connector.ExportableResource = &PingoneAuthorizePolicyManagementRuleResource{}
)

type PingoneAuthorizePolicyManagementRuleResource struct {
	clientInfo *connector.PingOneClientInfo
}

// Utility method for creating a PingoneAuthorizePolicyManagementRuleResource
func AuthorizePolicyManagementRule(clientInfo *connector.PingOneClientInfo) *PingoneAuthorizePolicyManagementRuleResource {
	return &PingoneAuthorizePolicyManagementRuleResource{
		clientInfo: clientInfo,
	}
}

func (r *PingoneAuthorizePolicyManagementRuleResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()

	l.Debug().Msgf("Fetching all %s resources...", r.ResourceType())

	apiExecuteFunc := r.clientInfo.ApiClient.AuthorizeAPIClient.AuthorizeEditorRulesApi.ListRules(r.clientInfo.Context, r.clientInfo.ExportEnvironmentID).Execute
	apiFunctionName := "ListRules"

	embedded, err := common.GetAuthorizeEmbedded(apiExecuteFunc, apiFunctionName, r.ResourceType())
	if err != nil {
		return nil, err
	}

	importBlocks := []connector.ImportBlock{}

	l.Debug().Msgf("Generating Import Blocks for all %s resources...", r.ResourceType())

	for _, authorizationRule := range embedded.GetAuthorizationRules() {
		authorizationRuleName, authorizationRuleNameOk := authorizationRule.GetNameOk()
		authorizationRuleId, authorizationRuleIdOk := authorizationRule.GetIdOk()

		if authorizationRuleNameOk && authorizationRuleIdOk {
			commentData := map[string]string{
				"Resource Type": r.ResourceType(),
				"Authorize Policy Management Authorization Rule Name": *authorizationRuleName,
				"Export Environment ID":                               r.clientInfo.ExportEnvironmentID,
				"Authorize Policy Management Authorization Rule ID":   *authorizationRuleId,
			}

			importBlocks = append(importBlocks, connector.ImportBlock{
				ResourceType:       r.ResourceType(),
				ResourceName:       *authorizationRuleName,
				ResourceID:         fmt.Sprintf("%s/%s", r.clientInfo.ExportEnvironmentID, *authorizationRuleId),
				CommentInformation: common.GenerateCommentInformation(commentData),
			})
		}
	}

	return &importBlocks, nil
}

func (r *PingoneAuthorizePolicyManagementRuleResource) ResourceType() string {
	return "pingone_authorize_policy_management_rule"
}
