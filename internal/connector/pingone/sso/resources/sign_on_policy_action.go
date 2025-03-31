// Copyright © 2025 Ping Identity Corporation
// Code generated by ping-cli-generator

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
	clientInfo *connector.ClientInfo
}

// Utility method for creating a PingOneSignOnPolicyActionResource
func SignOnPolicyAction(clientInfo *connector.ClientInfo) *PingOneSignOnPolicyActionResource {
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

	signOnPolicyData, err := r.getPolicyData()
	if err != nil {
		return nil, err
	}

	for signOnPolicyId, signOnPolicyName := range signOnPolicyData {
		signOnPolicyActionData, err := r.getSignOnPolicyActionData(signOnPolicyId)
		if err != nil {
			return nil, err
		}

		for signOnPolicysignOnPolicyActionId, signOnPolicyActionType := range signOnPolicyActionData {
			commentData := map[string]string{
				"Sign On Policy ID":          signOnPolicyId,
				"Sign On Policy Name":        signOnPolicyName,
				"Sign On Policy Action ID":   signOnPolicysignOnPolicyActionId,
				"Sign On Policy Action Type": signOnPolicyActionType,
				"Export Environment ID":      r.clientInfo.PingOneExportEnvironmentID,
				"Resource Type":              r.ResourceType(),
			}

			importBlock := connector.ImportBlock{
				ResourceType:       r.ResourceType(),
				ResourceName:       fmt.Sprintf("%s_%s", signOnPolicyName, signOnPolicyActionType),
				ResourceID:         fmt.Sprintf("%s/%s/%s", r.clientInfo.PingOneExportEnvironmentID, signOnPolicyId, signOnPolicysignOnPolicyActionId),
				CommentInformation: common.GenerateCommentInformation(commentData),
			}

			importBlocks = append(importBlocks, importBlock)
		}
	}

	return &importBlocks, nil
}

func (r *PingOneSignOnPolicyActionResource) getPolicyData() (map[string]string, error) {
	signOnPolicyData := make(map[string]string)

	iter := r.clientInfo.PingOneApiClient.ManagementAPIClient.SignOnPoliciesApi.ReadAllSignOnPolicies(r.clientInfo.PingOneContext, r.clientInfo.PingOneExportEnvironmentID).Execute()
	apiObjs, err := pingone.GetManagementAPIObjectsFromIterator[management.SignOnPolicy](iter, "ReadAllSignOnPolicies", "GetSignOnPolicies", r.ResourceType())
	if err != nil {
		return nil, err
	}

	for _, signOnPolicy := range apiObjs {
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

	iter := r.clientInfo.PingOneApiClient.ManagementAPIClient.SignOnPolicyActionsApi.ReadAllSignOnPolicyActions(r.clientInfo.PingOneContext, r.clientInfo.PingOneExportEnvironmentID, signOnPolicyId).Execute()
	apiObjs, err := pingone.GetManagementAPIObjectsFromIterator[management.SignOnPolicyAction](iter, "ReadAllSignOnPolicyActions", "GetActions", r.ResourceType())
	if err != nil {
		return nil, err
	}

	for _, innerObj := range apiObjs {
		var (
			signOnPolicyActionId     *string
			signOnPolicyActionIdOk   bool
			signOnPolicyActionType   *management.EnumSignOnPolicyType
			signOnPolicyActionTypeOk bool
		)

		switch {
		case innerObj.SignOnPolicyActionAgreement != nil:
			signOnPolicyActionId, signOnPolicyActionIdOk = innerObj.SignOnPolicyActionAgreement.GetIdOk()
			signOnPolicyActionType, signOnPolicyActionTypeOk = innerObj.SignOnPolicyActionAgreement.GetTypeOk()
		case innerObj.SignOnPolicyActionCommon != nil:
			signOnPolicyActionId, signOnPolicyActionIdOk = innerObj.SignOnPolicyActionCommon.GetIdOk()
			signOnPolicyActionType, signOnPolicyActionTypeOk = innerObj.SignOnPolicyActionCommon.GetTypeOk()
		case innerObj.SignOnPolicyActionIDFirst != nil:
			signOnPolicyActionId, signOnPolicyActionIdOk = innerObj.SignOnPolicyActionIDFirst.GetIdOk()
			signOnPolicyActionType, signOnPolicyActionTypeOk = innerObj.SignOnPolicyActionIDFirst.GetTypeOk()
		case innerObj.SignOnPolicyActionIDP != nil:
			signOnPolicyActionId, signOnPolicyActionIdOk = innerObj.SignOnPolicyActionIDP.GetIdOk()
			signOnPolicyActionType, signOnPolicyActionTypeOk = innerObj.SignOnPolicyActionIDP.GetTypeOk()
		case innerObj.SignOnPolicyActionLogin != nil:
			signOnPolicyActionId, signOnPolicyActionIdOk = innerObj.SignOnPolicyActionLogin.GetIdOk()
			signOnPolicyActionType, signOnPolicyActionTypeOk = innerObj.SignOnPolicyActionLogin.GetTypeOk()
		case innerObj.SignOnPolicyActionMFA != nil:
			signOnPolicyActionId, signOnPolicyActionIdOk = innerObj.SignOnPolicyActionMFA.GetIdOk()
			signOnPolicyActionType, signOnPolicyActionTypeOk = innerObj.SignOnPolicyActionMFA.GetTypeOk()
		case innerObj.SignOnPolicyActionPingIDWinLoginPasswordless != nil:
			signOnPolicyActionId, signOnPolicyActionIdOk = innerObj.SignOnPolicyActionPingIDWinLoginPasswordless.GetIdOk()
			signOnPolicyActionType, signOnPolicyActionTypeOk = innerObj.SignOnPolicyActionPingIDWinLoginPasswordless.GetTypeOk()
		case innerObj.SignOnPolicyActionProgressiveProfiling != nil:
			signOnPolicyActionId, signOnPolicyActionIdOk = innerObj.SignOnPolicyActionProgressiveProfiling.GetIdOk()
			signOnPolicyActionType, signOnPolicyActionTypeOk = innerObj.SignOnPolicyActionProgressiveProfiling.GetTypeOk()
		default:
			continue
		}

		if signOnPolicyActionIdOk && signOnPolicyActionTypeOk {
			signOnPolicyActionData[*signOnPolicyActionId] = string(*signOnPolicyActionType)
		}
	}

	return signOnPolicyActionData, nil
}
