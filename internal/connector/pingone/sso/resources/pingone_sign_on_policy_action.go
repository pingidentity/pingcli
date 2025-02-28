package resources

import (
	"fmt"

	"github.com/patrickcping/pingone-go-sdk-v2/management"
	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/common"
	"github.com/pingidentity/pingcli/internal/connector/pingone"
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

func (r *PingOneSignOnPolicyActionResource) ResourceType() string {
	return "pingone_sign_on_policy_action"
}

func (r *PingOneSignOnPolicyActionResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()
	l.Debug().Msgf("Exporting all '%s' Resources...", r.ResourceType())

	importBlocks := []connector.ImportBlock{}

	signOnPolicyData, err := r.getSignOnPolicyData()
	if err != nil {
		return nil, err
	}

	for signOnPolicyId, signOnPolicyName := range signOnPolicyData {
		signOnPolicyActionData, err := r.getSignOnPolicyActionData(signOnPolicyId)
		if err != nil {
			return nil, err
		}

		for actionId, actionType := range signOnPolicyActionData {
			commentData := map[string]string{
				"Export Environment ID":      r.clientInfo.ExportEnvironmentID,
				"Resource Type":              r.ResourceType(),
				"Sign-On Policy Action ID":   actionId,
				"Sign-On Policy Action Type": actionType,
				"Sign-On Policy ID":          signOnPolicyId,
				"Sign-On Policy Name":        signOnPolicyName,
			}

			importBlock := connector.ImportBlock{
				ResourceType:       r.ResourceType(),
				ResourceName:       fmt.Sprintf("%s_%s", signOnPolicyName, actionType),
				ResourceID:         fmt.Sprintf("%s/%s/%s", r.clientInfo.ExportEnvironmentID, signOnPolicyId, actionId),
				CommentInformation: common.GenerateCommentInformation(commentData),
			}

			importBlocks = append(importBlocks, importBlock)
		}
	}

	return &importBlocks, nil
}

func (r *PingOneSignOnPolicyActionResource) getSignOnPolicyData() (map[string]string, error) {
	signOnPolicyData := make(map[string]string)

	iter := r.clientInfo.ApiClient.ManagementAPIClient.SignOnPoliciesApi.ReadAllSignOnPolicies(r.clientInfo.Context, r.clientInfo.ExportEnvironmentID).Execute()
	signOnPolicies, err := pingone.GetManagementAPIObjectsFromIterator[management.SignOnPolicy](iter, "ReadAllSignOnPolicies", "GetSignOnPolicies", r.ResourceType())
	if err != nil {
		return nil, err
	}

	for _, signOnPolicy := range signOnPolicies {
		signOnPolicyId, signOnPolicyIdOk := signOnPolicy.GetIdOk()
		signOnPolicyName, signOnPolicyNameOk := signOnPolicy.GetNameOk()

		if signOnPolicyIdOk && signOnPolicyNameOk {
			signOnPolicyData[*signOnPolicyId] = *signOnPolicyName
		}
	}

	return signOnPolicyData, nil
}

func (r *PingOneSignOnPolicyActionResource) getSignOnPolicyActionData(signOnPolicyId string) (map[string]string, error) {
	signOnPolicyActionData := make(map[string]string)

	iter := r.clientInfo.ApiClient.ManagementAPIClient.SignOnPolicyActionsApi.ReadAllSignOnPolicyActions(r.clientInfo.Context, r.clientInfo.ExportEnvironmentID, signOnPolicyId).Execute()
	signOnPolicyActions, err := pingone.GetManagementAPIObjectsFromIterator[management.SignOnPolicyAction](iter, "ReadAllSignOnPolicyActions", "GetActions", r.ResourceType())
	if err != nil {
		return nil, err
	}

	for _, action := range signOnPolicyActions {
		var (
			actionId     *string
			actionIdOk   bool
			actionType   *management.EnumSignOnPolicyType
			actionTypeOk bool
		)

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
		default:
			continue
		}

		if actionIdOk && actionTypeOk {
			signOnPolicyActionData[*actionId] = string(*actionType)
		}
	}

	return signOnPolicyActionData, nil
}
