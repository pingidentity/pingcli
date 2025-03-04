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
	l.Debug().Msgf("Exporting all '%s' Resources...", r.ResourceType())

	importBlocks := []connector.ImportBlock{}

	editorRuleData, err := r.getEditorRuleData()
	if err != nil {
		return nil, err
	}

	for editorRuleId, editorRuleName := range editorRuleData {
		commentData := map[string]string{
			"Export Environment ID": r.clientInfo.ExportEnvironmentID,
			"Editor Rule ID":        editorRuleId,
			"Editor Rule Name":      editorRuleName,
			"Resource Type":         r.ResourceType(),
		}

		importBlock := connector.ImportBlock{
			ResourceType:       r.ResourceType(),
			ResourceName:       editorRuleName,
			ResourceID:         fmt.Sprintf("%s/%s", r.clientInfo.ExportEnvironmentID, editorRuleId),
			CommentInformation: common.GenerateCommentInformation(commentData),
		}

		importBlocks = append(importBlocks, importBlock)
	}

	return &importBlocks, nil
}

func (r *PingoneAuthorizePolicyManagementRuleResource) getEditorRuleData() (map[string]string, error) {
	editorRuleData := make(map[string]string)

	iter := r.clientInfo.ApiClient.AuthorizeAPIClient.AuthorizeEditorRulesApi.ListRules(r.clientInfo.Context, r.clientInfo.ExportEnvironmentID).Execute()
	editorRules, err := pingone.GetAuthorizeAPIObjectsFromIterator[authorize.AuthorizeEditorDataRulesReferenceableRuleDTO](iter, "ListRules", "GetAuthorizationRules", r.ResourceType())
	if err != nil {
		return nil, err
	}

	for _, editorRule := range editorRules {

		editorRuleId, editorRuleIdOk := editorRule.GetIdOk()
		editorRuleName, editorRuleNameOk := editorRule.GetNameOk()

		if editorRuleIdOk && editorRuleNameOk {
			editorRuleData[*editorRuleId] = *editorRuleName
		}
	}

	return editorRuleData, nil
}

func (r *PingoneAuthorizePolicyManagementRuleResource) ResourceType() string {
	return "pingone_authorize_policy_management_rule"
}
