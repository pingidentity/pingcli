package resources

import (
	"fmt"

	"github.com/patrickcping/pingone-go-sdk-v2/management"
	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/common"
	"github.com/pingidentity/pingcli/internal/logger"
)

// Verify that the resource satisfies the exportable resource interface
var (
	_ connector.ExportableResource = &PingOneSignOnPolicyActionResource{}
)

type PingOneSignOnPolicyActionResource struct {
	clientInfo *connector.PingOneClientInfo
}

// Utility method for creating a PingOneSignOnPolicyActionResource
func SignOnPolicyAction(clientInfo *connector.PingOneClientInfo) *PingOneSignOnPolicyActionResource {
	return &PingOneSignOnPolicyActionResource{
		clientInfo: clientInfo,
	}
}

func (r *PingOneSignOnPolicyActionResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()

	l.Debug().Msgf("Fetching all %s resources...", r.ResourceType())

	apiExecuteFunc := r.clientInfo.ApiClient.ManagementAPIClient.SignOnPoliciesApi.ReadAllSignOnPolicies(r.clientInfo.Context, r.clientInfo.ExportEnvironmentID).Execute
	apiFunctionName := "ReadAllSignOnPolicies"

	embedded, err := common.GetManagementEmbedded(apiExecuteFunc, apiFunctionName, r.ResourceType())
	if err != nil {
		return nil, err
	}

	importBlocks := []connector.ImportBlock{}

	l.Debug().Msgf("Generating Import Blocks for all %s resources...", r.ResourceType())

	for _, signOnPolicy := range embedded.GetSignOnPolicies() {
		signOnPolicyId, signOnPolicyIdOk := signOnPolicy.GetIdOk()
		signOnPolicyName, signOnPolicyNameOk := signOnPolicy.GetNameOk()

		if signOnPolicyIdOk && signOnPolicyNameOk {

			apiExecuteFunc := r.clientInfo.ApiClient.ManagementAPIClient.SignOnPolicyActionsApi.ReadAllSignOnPolicyActions(r.clientInfo.Context, r.clientInfo.ExportEnvironmentID, *signOnPolicyId).Execute
			apiFunctionName := "ReadAllSignOnPolicyActions"

			actionEmbedded, err := common.GetManagementEmbedded(apiExecuteFunc, apiFunctionName, r.ResourceType())
			if err != nil {
				return nil, err
			}

			var (
				actionId     *string
				actionIdOk   bool
				actionType   *management.EnumSignOnPolicyType
				actionTypeOk bool
			)

			for _, action := range actionEmbedded.GetActions() {
				switch {
				case action.SignOnPolicyActionAgreement != nil:
					actionId, actionIdOk = action.SignOnPolicyActionAgreement.GetIdOk()
					actionType, actionTypeOk = action.SignOnPolicyActionAgreement.GetTypeOk()
				case action.SignOnPolicyActionCommon != nil:
					actionId, actionIdOk = action.SignOnPolicyActionCommon.GetIdOk()
					actionType, actionTypeOk = action.SignOnPolicyActionCommon.GetTypeOk()
				case action.SignOnPolicyActionIDFirst != nil:
					actionId, actionIdOk = action.SignOnPolicyActionIDFirst.GetIdOk()
					actionType, actionTypeOk = action.SignOnPolicyActionIDFirst.GetTypeOk()
				case action.SignOnPolicyActionIDP != nil:
					actionId, actionIdOk = action.SignOnPolicyActionIDP.GetIdOk()
					actionType, actionTypeOk = action.SignOnPolicyActionIDP.GetTypeOk()
				case action.SignOnPolicyActionLogin != nil:
					actionId, actionIdOk = action.SignOnPolicyActionLogin.GetIdOk()
					actionType, actionTypeOk = action.SignOnPolicyActionLogin.GetTypeOk()
				case action.SignOnPolicyActionMFA != nil:
					actionId, actionIdOk = action.SignOnPolicyActionMFA.GetIdOk()
					actionType, actionTypeOk = action.SignOnPolicyActionMFA.GetTypeOk()
				case action.SignOnPolicyActionPingIDWinLoginPasswordless != nil:
					actionId, actionIdOk = action.SignOnPolicyActionPingIDWinLoginPasswordless.GetIdOk()
					actionType, actionTypeOk = action.SignOnPolicyActionPingIDWinLoginPasswordless.GetTypeOk()
				case action.SignOnPolicyActionProgressiveProfiling != nil:
					actionId, actionIdOk = action.SignOnPolicyActionProgressiveProfiling.GetIdOk()
					actionType, actionTypeOk = action.SignOnPolicyActionProgressiveProfiling.GetTypeOk()
				default:
					continue
				}

				if actionIdOk && actionTypeOk {
					commentData := map[string]string{
						"Resource Type":            r.ResourceType(),
						"Sign On Policy Name":      *signOnPolicyName,
						"Action Type":              string(*actionType),
						"Export Environment ID":    r.clientInfo.ExportEnvironmentID,
						"Sign On Policy ID":        *signOnPolicyId,
						"Sign On Policy Action ID": *actionId,
					}

					importBlocks = append(importBlocks, connector.ImportBlock{
						ResourceType:       r.ResourceType(),
						ResourceName:       fmt.Sprintf("%s_%s", *signOnPolicyName, *actionType),
						ResourceID:         fmt.Sprintf("%s/%s/%s", r.clientInfo.ExportEnvironmentID, *signOnPolicyId, *actionId),
						CommentInformation: common.GenerateCommentInformation(commentData),
					})
				}
			}
		}
	}

	return &importBlocks, nil
}

func (r *PingOneSignOnPolicyActionResource) ResourceType() string {
	return "pingone_sign_on_policy_action"
}
